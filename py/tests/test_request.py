"""Sending requests and surfacing failures."""

from __future__ import annotations

from typing import Any

import httpx
import pytest

from opti_cmp import (
    APIConnectionError,
    APIError,
    APITimeoutError,
    AsyncOptiCMP,
    Headers,
    InvalidRequestError,
    OptiCMP,
    OptiCMPError,
)
from opti_cmp._request import transport_error
from opti_cmp._types import RequestOptions

from .conftest import json_transport, recording, transport


def _raise_connect_error(_request: httpx.Request) -> httpx.Response:
    raise httpx.ConnectError("connection refused")


class TestSuccessfulResponses:
    def test_parses_json_bodies(self) -> None:
        payload = {"id": 1, "name": "John"}
        with OptiCMP(
            transport=transport(lambda _r: httpx.Response(200, json=payload))
        ) as client:
            response = client.request("/users")

        assert response.data == payload
        assert response.status == 200
        assert response.url == "https://api.cmp.optimizely.com/v3/users"

    def test_returns_text_for_text_content_types(self) -> None:
        with OptiCMP(
            transport=transport(
                lambda _r: httpx.Response(
                    200, text="<html></html>", headers={"content-type": "text/html"}
                )
            )
        ) as client:
            assert client.request("/page").data == "<html></html>"

    def test_returns_none_for_204(self) -> None:
        with OptiCMP(transport=transport(lambda _r: httpx.Response(204))) as client:
            assert client.request("DELETE /users/1").data is None

    def test_returns_bytes_for_binary_content(self) -> None:
        with OptiCMP(
            transport=transport(
                lambda _r: httpx.Response(
                    200, content=b"\x89PNG", headers={"content-type": "image/png"}
                )
            )
        ) as client:
            assert client.request("/image").data == b"\x89PNG"

    def test_exposes_response_headers(self) -> None:
        with OptiCMP(
            transport=transport(
                lambda _r: httpx.Response(200, json={}, headers={"x-request-id": "abc"})
            )
        ) as client:
            assert client.request("/users").headers["x-request-id"] == "abc"


class TestErrors:
    def test_raises_http_error_with_status_and_body(self) -> None:
        with OptiCMP(
            transport=transport(
                lambda _r: httpx.Response(404, json={"message": "Not Found"})
            )
        ) as client:
            with pytest.raises(APIError) as info:
                client.request("/users/missing")

        assert info.value.status == 404
        assert info.value.data == {"message": "Not Found"}
        assert info.value.request is not None
        assert info.value.request.url.endswith("/users/missing")

    def test_connection_failure_is_not_an_http_error(self) -> None:
        # No response arrived, so there is no status to report. Previously this
        # was reported as APIError(status=500), indistinguishable from a real
        # server 500.
        def boom(_request: httpx.Request) -> httpx.Response:
            raise httpx.ConnectError("connection refused")

        with OptiCMP(transport=transport(boom)) as client:
            with pytest.raises(APIConnectionError) as info:
                client.request("/users")

        assert "connection refused" in info.value.message
        assert not hasattr(info.value, "status")
        assert info.value.request is not None
        assert info.value.request.url.endswith("/users")

    def test_timeout_is_its_own_error(self) -> None:
        def slow(_request: httpx.Request) -> httpx.Response:
            raise httpx.ReadTimeout("timed out")

        with OptiCMP(transport=transport(slow)) as client:
            with pytest.raises(APITimeoutError):
                client.request("/users")

    def test_a_server_500_stays_an_http_error(self) -> None:
        # The distinction that was impossible to make before.
        with OptiCMP(
            transport=transport(lambda _r: httpx.Response(500, json={"m": "boom"}))
        ) as client:
            with pytest.raises(APIError) as info:
                client.request("/users")

        assert info.value.status == 500
        assert not isinstance(info.value, APIConnectionError)

    @pytest.mark.parametrize(
        "transport_factory",
        [
            lambda: transport(lambda _r: httpx.Response(500, json={})),
            lambda: transport(_raise_connect_error),
        ],
        ids=["http-error", "connection-error"],
    )
    def test_every_failure_is_an_opti_cmp_error(self, transport_factory: Any) -> None:
        with OptiCMP(transport=transport_factory()) as client:
            with pytest.raises(OptiCMPError):
                client.request("/users")

    async def test_async_connection_failure(self) -> None:
        async with AsyncOptiCMP(transport=transport(_raise_connect_error)) as client:
            with pytest.raises(APIConnectionError):
                await client.request("/users")

    def test_http_error_carries_its_fields(self) -> None:
        error = APIError("Bad Request", 400, data={"a": 1}, headers=Headers({"x": "y"}))
        assert (error.status, error.data) == (400, {"a": 1})
        assert error.headers == Headers({"X": "y"})  # case-insensitive
        assert str(error) == "Bad Request"
        assert isinstance(error, Exception)


class TestRequestPipeline:
    def test_sends_method_url_and_body(self) -> None:
        seen: list[httpx.Request] = []
        with OptiCMP(transport=json_transport({}, status=201, recorder=seen)) as client:
            client.request("POST /campaigns", {"body": {"name": "Q3"}})

        assert seen[0].method == "POST"
        assert str(seen[0].url) == "https://api.cmp.optimizely.com/v3/campaigns"
        assert seen[0].read() == b'{"name": "Q3"}'

    def test_honours_a_custom_base_url(self) -> None:
        seen, ok = recording()
        with OptiCMP(
            base_url="https://staging.example.com/v3",
            transport=ok,
        ) as client:
            client.request("/campaigns")

        assert str(seen[0].url) == "https://staging.example.com/v3/campaigns"

    def test_client_headers_reach_the_request(self) -> None:
        seen, ok = recording()
        with OptiCMP(
            headers={"X-Org": "acme"},
            transport=ok,
        ) as client:
            client.request("/campaigns")

        assert seen[0].headers["x-org"] == "acme"

    def test_sends_a_user_agent(self) -> None:
        seen, ok = recording()
        with OptiCMP(transport=ok) as client:
            client.request("/campaigns")

        assert seen[0].headers["user-agent"].startswith("opti-cmp.py/")

    async def test_async_client_sends_the_same_request(self) -> None:
        seen: list[httpx.Request] = []
        async with AsyncOptiCMP(
            transport=json_transport({"ok": True}, recorder=seen)
        ) as client:
            response = await client.request("GET /campaigns")

        assert response.data == {"ok": True}
        assert seen[0].method == "GET"

    async def test_async_client_raises_http_error(self) -> None:
        async with AsyncOptiCMP(
            transport=transport(lambda _r: httpx.Response(500, json={"m": "boom"}))
        ) as client:
            with pytest.raises(APIError) as info:
                await client.request("/campaigns")

        assert info.value.status == 500


class TestResponseHeaders:
    def test_reads_are_case_insensitive(self) -> None:
        with OptiCMP(
            transport=transport(
                lambda _r: httpx.Response(200, json={}, headers={"X-Request-Id": "r1"})
            )
        ) as client:
            headers = client.request("/x").headers

        assert headers["X-Request-Id"] == headers["x-request-id"] == "r1"

    def test_repeated_headers_are_not_collapsed(self) -> None:
        # A plain dict would join these into "a=1, b=2" and lose the originals.
        raw = [("set-cookie", "a=1"), ("set-cookie", "b=2")]
        with OptiCMP(
            transport=transport(lambda _r: httpx.Response(200, json={}, headers=raw))
        ) as client:
            headers = client.request("/x").headers

        assert headers.get_list("set-cookie") == ["a=1", "b=2"]


class TestAwkwardBodies:
    """The fallback branches in `_response_data`."""

    def test_malformed_json_falls_back_to_text(self) -> None:
        broken = httpx.Response(
            200,
            content=b"{not json",
            headers={"content-type": "application/json"},
        )
        with OptiCMP(transport=transport(lambda _r: broken)) as client:
            assert client.request("/x").data == "{not json"

    def test_empty_json_body_falls_back_to_text(self) -> None:
        empty = httpx.Response(
            200, content=b"", headers={"content-type": "application/json"}
        )
        with OptiCMP(transport=transport(lambda _r: empty)) as client:
            assert client.request("/x").data == ""

    def test_charset_utf8_on_a_non_text_type_is_decoded(self) -> None:
        response = httpx.Response(
            200,
            content="héllo".encode(),
            headers={"content-type": "application/xml; charset=utf-8"},
        )
        with OptiCMP(transport=transport(lambda _r: response)) as client:
            assert client.request("/x").data == "héllo"

    def test_unknown_binary_type_stays_bytes(self) -> None:
        response = httpx.Response(
            200, content=b"\x89PNG", headers={"content-type": "image/png"}
        )
        with OptiCMP(transport=transport(lambda _r: response)) as client:
            assert client.request("/x").data == b"\x89PNG"


class TestTransportOptions:
    def test_default_timeout_is_not_httpx_five_seconds(self) -> None:
        with OptiCMP() as client:
            assert client._httpx.timeout == httpx.Timeout(30.0, connect=10.0)

    def test_client_timeout_reaches_httpx(self) -> None:
        seen, ok = recording()
        with OptiCMP(request={"timeout": 12.5}, transport=ok) as client:
            client.request("/campaigns")

        assert seen[0].extensions["timeout"] == {
            "connect": 12.5,
            "pool": 12.5,
            "read": 12.5,
            "write": 12.5,
        }

    def test_per_call_timeout_overrides_the_client(self) -> None:
        seen, ok = recording()
        with OptiCMP(request={"timeout": 12.5}, transport=ok) as client:
            client.request("/campaigns", {"request": {"timeout": 1.0}})

        assert seen[0].extensions["timeout"]["read"] == 1.0

    def test_an_explicit_none_timeout_disables_the_deadline(self) -> None:
        seen, ok = recording()
        with OptiCMP(request={"timeout": None}, transport=ok) as client:
            client.request("/campaigns")

        assert seen[0].extensions["timeout"] == {
            "connect": None,
            "pool": None,
            "read": None,
            "write": None,
        }

    def test_redirects_are_followed(self) -> None:
        def handler(request: httpx.Request) -> httpx.Response:
            if request.url.path.endswith("/old"):
                return httpx.Response(302, headers={"location": "/v3/new"})
            return httpx.Response(200, json={"landed": True})

        with OptiCMP(transport=transport(handler)) as client:
            response = client.request("/old")

        assert response.data == {"landed": True}
        assert response.url.endswith("/v3/new")

    def test_a_malformed_base_url_is_an_sdk_error(self) -> None:
        # httpx.InvalidURL does not derive from httpx.HTTPError, so it would
        # otherwise escape past `except OptiCMPError`.
        with OptiCMP(base_url="https://[::1") as client:
            with pytest.raises(InvalidRequestError):
                client.request("/campaigns")

    @pytest.mark.parametrize(
        ("error", "expected"),
        [
            (httpx.ConnectError("refused"), APIConnectionError),
            (httpx.ConnectTimeout("too slow"), APITimeoutError),
            (httpx.InvalidURL("bad port"), InvalidRequestError),
            # h11 raises this when a header value carries CRLF: the client
            # refused to send it, so it is caller input, not a network fault.
            (httpx.LocalProtocolError("illegal header value"), InvalidRequestError),
        ],
    )
    def test_httpx_failures_map_to_the_right_sdk_error(
        self, error: Exception, expected: type[OptiCMPError]
    ) -> None:
        options = RequestOptions(method="GET", url="/campaigns")
        assert type(transport_error(error, options)) is expected  # type: ignore[arg-type]
