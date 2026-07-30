"""`client.request_hook` as an SDK user meets it: logging, metrics, retries."""

from __future__ import annotations

import time
from typing import Any, cast

import httpx
import pytest

from opti_cmp import (
    APIError,
    AsyncOptiCMP,
    EndpointOptions,
    Headers,
    OptiCMP,
    Response,
    TokenAuth,
)

from .conftest import json_transport, recording

OK = json_transport({"data": []})
FAILING = json_transport({"m": "boom"}, status=500)


class TestLogging:
    def test_before_sees_the_operation(self) -> None:
        records: list[str] = []

        with OptiCMP(transport=OK) as client:

            @client.request_hook.before
            def log_request(options: dict[str, Any]) -> None:
                records.append(f"{options['method']} {options['url']}")

            client.campaign.get_campaign("camp-1")

        assert records == ["GET /campaigns/{id}"]

    def test_after_sees_the_response(self) -> None:
        records: list[int] = []

        with OptiCMP(transport=OK) as client:

            @client.request_hook.after
            def log_response(response: Response[Any], _options: dict[str, Any]) -> None:
                records.append(response.status)

            client.campaign.list_campaigns()

        assert records == [200]

    def test_error_hook_logs_without_swallowing(self) -> None:
        records: list[APIError] = []

        with OptiCMP(transport=FAILING) as client:

            @client.request_hook.error
            def log_error(error: BaseException, _options: dict[str, Any]) -> None:
                records.append(error)  # type: ignore[arg-type]

            with pytest.raises(APIError) as info:
                client.campaign.list_campaigns()

        assert info.value.status == 500
        assert len(records) == 1


class TestSemanticLayer:
    """What the SDK hook sees that an httpx event hook cannot."""

    def test_before_sees_namespace_and_operation(self) -> None:
        seen: list[str] = []

        with OptiCMP(transport=OK) as client:
            client.request_hook.before(
                lambda o: seen.append(f"{o['namespace']}.{o['operation']}")
            )
            client.campaign.get_campaign("camp-abc-123")

        assert seen == ["campaign.get_campaign"]

    def test_url_stays_templated_for_metric_labels(self) -> None:
        # The concrete URL would be per-campaign, which is unusable as a metric
        # label. httpx event hooks only ever see the expanded URL.
        seen: list[str] = []

        with OptiCMP(transport=OK) as client:
            client.request_hook.before(lambda o: seen.append(o["url"]))
            client.campaign.get_campaign("camp-abc-123")

        assert seen == ["/campaigns/{id}"]

    def test_namespace_and_operation_never_reach_the_wire(self) -> None:
        seen, ok = recording()

        with OptiCMP(transport=ok) as client:
            client.campaign.list_campaigns(page_size=2)
            client.campaign.create_campaign(body={"title": "Q3"})

        query = str(seen[0].url)
        assert "namespace" not in query and "operation" not in query

        body = seen[1].read()
        assert b"namespace" not in body and b"operation" not in body


class TestMetrics:
    def test_wrap_can_time_the_whole_call(self) -> None:
        timings: list[float] = []

        with OptiCMP(transport=OK) as client:

            @client.request_hook.wrap
            def measure(inner: Any, options: dict[str, Any]) -> Any:
                started = time.perf_counter()
                try:
                    return inner(options)
                finally:
                    timings.append(time.perf_counter() - started)

            client.campaign.list_campaigns()

        assert len(timings) == 1
        assert timings[0] >= 0

    def test_counting_successes_and_failures(self) -> None:
        counts = {"ok": 0, "failed": 0}
        responses = iter([httpx.Response(200, json={}), httpx.Response(503, json={})])

        with OptiCMP(
            transport=httpx.MockTransport(lambda _r: next(responses))
        ) as client:
            client.request_hook.after(
                lambda _response, _options: counts.__setitem__("ok", counts["ok"] + 1)
            )
            client.request_hook.error(
                lambda _error, _options: counts.__setitem__(
                    "failed", counts["failed"] + 1
                )
            )

            client.campaign.list_campaigns()
            with pytest.raises(APIError):
                client.campaign.list_campaigns()

        assert counts == {"ok": 1, "failed": 1}


class TestCustomLogic:
    def test_before_can_inject_a_header(self) -> None:
        seen, ok = recording()

        with OptiCMP(transport=ok) as client:
            client.request_hook.before(
                lambda options: options["headers"].__setitem__("x-tenant", "acme")
            )
            client.campaign.list_campaigns()

        assert seen[0].headers["x-tenant"] == "acme"

    def test_error_hook_can_supply_a_fallback(self) -> None:
        fallback: Response[Any] = Response(
            data={"data": []}, headers=Headers(), status=200, url=""
        )

        with OptiCMP(transport=FAILING) as client:
            client.request_hook.error(lambda _error, _options: fallback)
            result = client.campaign.list_campaigns()

        assert cast(Any, result) is fallback

    def test_wrap_can_retry(self) -> None:
        attempts = {"n": 0}

        def flaky(_request: httpx.Request) -> httpx.Response:
            attempts["n"] += 1
            if attempts["n"] < 3:
                return httpx.Response(503, json={})
            return httpx.Response(200, json={"data": []})

        with OptiCMP(transport=httpx.MockTransport(flaky)) as client:

            @client.request_hook.wrap
            def retry(inner: Any, options: dict[str, Any]) -> Any:
                for attempt in range(3):
                    try:
                        return inner(options)
                    except APIError:
                        if attempt == 2:
                            raise
                return None

            result = client.campaign.list_campaigns()

        assert attempts["n"] == 3
        assert result.status == 200

    def test_after_can_replace_the_response(self) -> None:
        with OptiCMP(transport=OK) as client:
            client.request_hook.after(
                lambda response, _options: Response(
                    data={"unwrapped": True},
                    headers=response.headers,
                    status=response.status,
                    url=response.url,
                )
            )
            result = client.campaign.list_campaigns()

        assert cast(Any, result.data) == {"unwrapped": True}

    def test_hooks_compose_with_auth(self) -> None:
        seen, ok = recording()

        with OptiCMP(
            auth=TokenAuth("tok"),
            transport=ok,
        ) as client:
            client.request_hook.before(
                lambda options: options["headers"].__setitem__("x-tenant", "acme")
            )
            client.campaign.list_campaigns()

        assert seen[0].headers["authorization"] == "Bearer tok"
        assert seen[0].headers["x-tenant"] == "acme"

    def test_remove_unregisters_a_hook(self) -> None:
        calls: list[str] = []

        def observer(_options: dict[str, Any]) -> None:
            calls.append("x")

        with OptiCMP(transport=OK) as client:
            client.request_hook.before(observer)
            client.campaign.list_campaigns()
            client.request_hook.remove(observer)
            client.campaign.list_campaigns()

        assert len(calls) == 1


class TestAsyncHooks:
    async def test_async_before_hook(self) -> None:
        seen, ok = recording()

        async with AsyncOptiCMP(transport=ok) as client:

            @client.request_hook.before
            async def add_header(options: dict[str, Any]) -> None:
                options["headers"]["x-tenant"] = "acme"

            await client.campaign.list_campaigns()

        assert seen[0].headers["x-tenant"] == "acme"

    async def test_async_error_hook_logs_without_swallowing(self) -> None:
        records: list[BaseException] = []

        async with AsyncOptiCMP(transport=FAILING) as client:

            @client.request_hook.error
            async def log_error(error: BaseException, _options: dict[str, Any]) -> None:
                records.append(error)

            with pytest.raises(APIError):
                await client.campaign.list_campaigns()

        assert len(records) == 1

    async def test_async_after_hook_sees_the_response(self) -> None:
        statuses: list[int] = []

        async with AsyncOptiCMP(transport=OK) as client:

            @client.request_hook.after
            async def record(response: Response[Any], _options: dict[str, Any]) -> None:
                statuses.append(response.status)

            await client.campaign.list_campaigns()

        assert statuses == [200]


class TestOptionsObject:
    """Hook options read and write as attributes as well as keys."""

    def test_headers_are_reachable_as_an_attribute(self) -> None:
        seen, ok = recording()
        with OptiCMP(transport=ok) as client:

            @client.request_hook.before
            def add_tenant(options: EndpointOptions) -> None:
                options.headers["x-tenant"] = "acme"

            client.campaign.list_campaigns()

        assert seen[0].headers["x-tenant"] == "acme"

    def test_attribute_and_key_access_agree(self) -> None:
        captured: list[EndpointOptions] = []
        with OptiCMP(transport=OK) as client:
            client.request_hook.before(captured.append)
            client.campaign.get_campaign("camp-1")

        options = captured[0]
        assert options.namespace == "campaign"
        assert options.operation == "get_campaign"
        assert options.url == options["url"] == "/campaigns/{id}"
        assert options.headers is options["headers"]

    def test_declared_but_absent_fields_read_as_none(self) -> None:
        # A raw request has no operation; the annotation says `str | None`, so
        # reading it must not raise.
        captured: list[EndpointOptions] = []
        with OptiCMP(transport=OK) as client:
            client.request_hook.before(captured.append)
            client.request("/campaigns")

        assert captured[0].operation is None
        assert captured[0].namespace is None

    def test_unknown_fields_still_raise(self) -> None:
        captured: list[EndpointOptions] = []
        with OptiCMP(transport=OK) as client:
            client.request_hook.before(captured.append)
            client.campaign.list_campaigns()

        with pytest.raises(AttributeError, match="nonsense"):
            _ = cast(Any, captured[0]).nonsense

    def test_it_is_still_a_dict(self) -> None:
        captured: list[EndpointOptions] = []
        with OptiCMP(transport=OK) as client:
            client.request_hook.before(captured.append)
            client.campaign.list_campaigns()

        assert isinstance(captured[0], dict)
