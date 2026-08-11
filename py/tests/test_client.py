"""Client construction, plugins, dispatch and endpoint registration."""

from __future__ import annotations

from typing import Any

import pytest

from opti_cmp import AsyncOptiCMP, Headers, OptiCMP, TokenAuth
from opti_cmp._hook import SingularHook

from .conftest import json_transport, recording

OK = json_transport({})


class TestEndpointDefaults:
    def _defaults(self, **options: Any) -> dict[str, Any]:
        client = OptiCMP(transport=OK, **options)
        return client._endpoint_defaults(client.request_hook)

    def test_extracts_base_url(self) -> None:
        assert self._defaults(base_url="https://api.example.com")["base_url"] == (
            "https://api.example.com"
        )

    def test_extracts_headers(self) -> None:
        assert self._defaults(headers={"x-custom": "value"})["headers"] == {
            "x-custom": "value"
        }

    def test_extracts_request_options(self) -> None:
        assert self._defaults(request={"timeout": 5.0})["request"]["timeout"] == 5.0

    def test_attaches_the_hook(self) -> None:
        client = OptiCMP(transport=OK)
        defaults = client._endpoint_defaults(client.request_hook)
        assert defaults["request"]["hook"] is client.request_hook

    def test_drops_the_auth_option(self) -> None:
        defaults = self._defaults(auth=TokenAuth("secret"))
        assert "auth" not in defaults

    def test_defaults_to_empty_headers(self) -> None:
        assert self._defaults()["headers"] == {}


class TestConstructor:
    def test_exposes_request_and_hook(self) -> None:
        client = OptiCMP(transport=OK)
        assert callable(client.request)
        assert isinstance(client.request_hook, SingularHook)

    def test_runs_registered_plugins(self) -> None:
        calls: list[str] = []

        def plugin(_client: Any, _options: dict[str, Any]) -> None:
            calls.append("ran")

        OptiCMP.plugins(plugin)(transport=OK)
        assert calls == ["ran"]

    def test_passes_options_to_plugins(self) -> None:
        received: list[dict[str, Any]] = []

        def plugin(_client: Any, options: dict[str, Any]) -> None:
            received.append(options)

        OptiCMP.plugins(plugin)(transport=OK, base_url="https://api.example.com")
        assert received[0]["base_url"] == "https://api.example.com"

    def test_plugins_can_add_attributes(self) -> None:
        def plugin(client: Any, _options: dict[str, Any]) -> None:
            client.custom = lambda: "hello"

        # Plugin-added attributes are the untyped escape hatch.
        client: Any = OptiCMP.plugins(plugin)(transport=OK)
        assert client.custom() == "hello"

    def test_accepts_unknown_options_for_plugins(self) -> None:
        seen: list[Any] = []

        def plugin(_client: Any, options: dict[str, Any]) -> None:
            seen.append(options.get("my_option"))

        OptiCMP.plugins(plugin)(transport=OK, my_option=42)
        assert seen == [42]


class TestPluginFactory:
    def test_returns_a_new_class_and_leaves_the_original_alone(self) -> None:
        def plugin(_client: Any, _options: dict[str, Any]) -> None: ...

        extended = OptiCMP.plugins(plugin)
        assert extended is not OptiCMP
        assert plugin in extended._plugins
        assert plugin not in OptiCMP._plugins

    def test_skips_duplicate_plugins(self) -> None:
        def plugin(_client: Any, _options: dict[str, Any]) -> None: ...

        once = OptiCMP.plugins(plugin)
        twice = once.plugins(plugin)
        assert twice._plugins.count(plugin) == 1

    def test_is_chainable(self) -> None:
        order: list[str] = []

        def first(_c: Any, _o: dict[str, Any]) -> None:
            order.append("first")

        def second(_c: Any, _o: dict[str, Any]) -> None:
            order.append("second")

        OptiCMP.plugins(first).plugins(second)(transport=OK)
        assert order == ["first", "second"]

    def test_instances_remain_opti_cmp(self) -> None:
        def plugin(_c: Any, _o: dict[str, Any]) -> None: ...

        assert isinstance(OptiCMP.plugins(plugin)(transport=OK), OptiCMP)


class TestGeneratedNamespaces:
    def test_attaches_all_namespaces(self) -> None:
        client = OptiCMP(transport=OK)
        assert hasattr(client.campaign, "list_campaigns")
        assert hasattr(client.brand_compliance, "list_brand_compliance_categories")

    def test_builds_the_request_from_route_metadata(self) -> None:
        seen, ok = recording()
        with OptiCMP(transport=ok) as client:
            client.campaign.get_campaign("camp-1")

        assert seen[0].method == "GET"
        assert str(seen[0].url) == "https://api.cmp.optimizely.com/v3/campaigns/camp-1"

    def test_sends_a_body_for_write_endpoints(self) -> None:
        seen, ok = recording()
        with OptiCMP(transport=ok) as client:
            client.campaign.update_campaign("camp-1", body={"title": "renamed"})

        assert seen[0].method == "PATCH"
        assert seen[0].read() == b'{"title": "renamed"}'

    def test_per_call_headers_are_merged(self) -> None:
        seen, ok = recording()
        with OptiCMP(transport=ok) as client:
            client.campaign.get_campaign("camp-1", headers={"X-Trace": "t1"})

        assert seen[0].headers["x-trace"] == "t1"

    async def test_async_namespaces_are_awaitable(self) -> None:
        async with AsyncOptiCMP(transport=json_transport({"id": "camp-1"})) as client:
            response = await client.campaign.get_campaign("camp-1")

        assert response.status == 200
        assert response.data["id"] == "camp-1"


class TestRegisterEndpoints:
    def test_attaches_runtime_endpoints(self) -> None:
        seen, ok = recording()
        with OptiCMP(transport=ok) as typed_client:
            # Runtime-registered namespaces are untyped, like plugin attributes.
            client: Any = typed_client
            client.register_endpoints(
                {
                    "widget": {
                        "list_widgets": {
                            "method": "GET",
                            "url": "/widgets",
                        }
                    }
                }
            )
            client.widget.list_widgets(limit=5)

        assert str(seen[0].url) == "https://api.cmp.optimizely.com/v3/widgets?limit=5"

    async def test_attaches_runtime_endpoints_on_the_async_client(self) -> None:
        async with AsyncOptiCMP(transport=json_transport({"ok": True})) as typed_client:
            client: Any = typed_client
            client.register_endpoints(
                {
                    "widget": {
                        "get_widget": {
                            "method": "GET",
                            "url": "/widgets/{id}",
                        }
                    }
                }
            )
            response = await client.widget.get_widget(id="w1")

        assert response.data == {"ok": True}

    def test_adds_to_a_generated_namespace_without_dropping_it(self) -> None:
        seen, ok = recording()
        with OptiCMP(transport=ok) as typed_client:
            client: Any = typed_client
            client.register_endpoints(
                {"campaign": {"archive_campaign": {"method": "POST", "url": "/x"}}}
            )
            client.campaign.archive_campaign()
            client.campaign.get_campaign(id="camp-1")

        assert [str(request.url) for request in seen] == [
            "https://api.cmp.optimizely.com/v3/x",
            "https://api.cmp.optimizely.com/v3/campaigns/camp-1",
        ]


class TestPagination:
    NEXT = "https://api.cmp.optimizely.com/v3/campaigns?offset=50"

    def _page(self, *, next_: str | None = None, previous: str | None = None) -> Any:
        return {"data": [], "pagination": {"next": next_, "previous": previous}}

    def test_has_next_page(self) -> None:
        client = OptiCMP(transport=OK)
        response = _response(self._page(next_=self.NEXT))
        assert client.has_next_page(response) is True
        assert client.has_previous_page(response) is False

    def test_has_no_next_page_when_null(self) -> None:
        client = OptiCMP(transport=OK)
        assert client.has_next_page(_response(self._page())) is False

    def test_has_previous_page(self) -> None:
        client = OptiCMP(transport=OK)
        response = _response(self._page(previous=self.NEXT))
        assert client.has_previous_page(response) is True

    def test_tolerates_responses_without_pagination(self) -> None:
        client = OptiCMP(transport=OK)
        assert client.has_next_page(_response({"id": "x"})) is False

    def test_an_empty_next_url_is_not_a_page(self) -> None:
        # Treating "" as a page would GET the bare base URL.
        client = OptiCMP(transport=OK)
        response = _response(self._page(next_=""))
        assert client.has_next_page(response) is False
        with pytest.raises(ValueError, match="No next page available"):
            client.get_next_page(response)

    def test_get_next_page_follows_the_url(self) -> None:
        seen, ok = recording()
        with OptiCMP(transport=ok) as client:
            client.get_next_page(_response(self._page(next_=self.NEXT)))

        assert str(seen[0].url) == self.NEXT

    def test_get_next_page_raises_when_absent(self) -> None:
        # Caller misuse, so a ValueError rather than a fabricated HTTP 404.
        client = OptiCMP(transport=OK)
        with pytest.raises(ValueError, match="No next page available"):
            client.get_next_page(_response(self._page()))

    def test_get_previous_page_raises_when_absent(self) -> None:
        client = OptiCMP(transport=OK)
        with pytest.raises(ValueError, match="No previous page available"):
            client.get_previous_page(_response(self._page()))

    async def test_async_get_next_page(self) -> None:
        seen, ok = recording()
        async with AsyncOptiCMP(transport=ok) as client:
            await client.get_next_page(_response(self._page(next_=self.NEXT)))

        assert str(seen[0].url) == self.NEXT


def _response(data: Any) -> Any:
    from opti_cmp import Response

    return Response(data=data, headers=Headers(), status=200, url="")


class TestConstructorTyping:
    """The signature lives once on BaseClient and both clients inherit it."""

    def test_neither_client_overrides_init(self) -> None:
        # An override taking **options: Any would erase the typed signature,
        # which is how every constructor argument went unchecked before.
        assert "__init__" not in OptiCMP.__dict__
        assert "__init__" not in AsyncOptiCMP.__dict__

    def test_both_clients_share_the_signature(self) -> None:
        import inspect

        assert inspect.signature(OptiCMP) == inspect.signature(AsyncOptiCMP)

    def test_setup_runs_before_plugins(self) -> None:
        # Plugins read client.request, so the pipeline must exist by then.
        seen: list[bool] = []

        def plugin(client: Any, _options: dict[str, Any]) -> None:
            seen.append(hasattr(client, "request"))

        OptiCMP.plugins(plugin)(transport=OK)
        assert seen == [True]

    def test_base_client_requires_a_setup_implementation(self) -> None:
        from opti_cmp import BaseClient

        with pytest.raises(NotImplementedError):
            BaseClient()


class TestClientHeaders:
    def test_repeated_client_headers_survive(self) -> None:
        # A dict copy would collapse these into one "a, b" value.
        seen, ok = recording()
        with OptiCMP(
            headers=Headers([("x-flag", "a"), ("x-flag", "b")]),
            transport=ok,
        ) as client:
            client.campaign.list_campaigns()

        assert seen[0].headers.get_list("x-flag") == ["a", "b"]

    def test_a_plain_dict_is_still_accepted(self) -> None:
        seen, ok = recording()
        with OptiCMP(
            headers={"X-Org": "acme"},
            transport=ok,
        ) as client:
            client.campaign.list_campaigns()

        assert seen[0].headers["x-org"] == "acme"

    def test_a_sequence_of_pairs_is_accepted(self) -> None:
        # The form that expresses duplicates without constructing Headers.
        seen, ok = recording()
        with OptiCMP(
            headers=[("x-flag", "a"), ("x-flag", "b")],
            transport=ok,
        ) as client:
            client.campaign.list_campaigns()

        assert seen[0].headers.get_list("x-flag") == ["a", "b"]
