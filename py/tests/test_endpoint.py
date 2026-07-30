"""Merging endpoint options and parsing them into a request."""

from __future__ import annotations

import json

import pytest

from opti_cmp import Headers
from opti_cmp._endpoint import merge, parse

DEFAULTS = {
    "base_url": "https://api.example.com",
    "headers": {"accept": "application/json"},
    "method": "GET",
    "url": "/users",
}


def parsed(**overrides: object):  # type: ignore[no-untyped-def]
    return parse({**DEFAULTS, **overrides})


class TestParseUrl:
    def test_prepends_base_url_to_relative_urls(self) -> None:
        assert parsed().url == "https://api.example.com/users"

    def test_leaves_absolute_urls_alone(self) -> None:
        assert parsed(url="https://other.com/path").url == "https://other.com/path"

    def test_expands_a_template_variable(self) -> None:
        assert parsed(url="/users/{id}", id="123").url == (
            "https://api.example.com/users/123"
        )

    def test_expands_multiple_template_variables(self) -> None:
        result = parsed(
            url="/users/{user_id}/posts/{post_id}", user_id="u1", post_id="p2"
        )
        assert result.url == "https://api.example.com/users/u1/posts/p2"

    def test_percent_encodes_template_values(self) -> None:
        assert parsed(url="/users/{id}", id="a b/c").url == (
            "https://api.example.com/users/a%20b%2Fc"
        )


class TestParseQuery:
    def test_adds_query_parameters_for_get(self) -> None:
        url = parsed(page=1, limit=10).url
        assert "page=1" in url
        assert "limit=10" in url

    def test_repeats_key_for_list_values(self) -> None:
        url = parsed(tags=["a", "b", "c"]).url
        assert url.endswith("?tags=a&tags=b&tags=c")

    def test_url_encodes_values(self) -> None:
        assert "query=hello%20world" in parsed(query="hello world").url

    def test_serializes_booleans_as_json_not_python(self) -> None:
        # `str(True)` would produce "True", which the API does not accept.
        url = parsed(disabled=True, execute=False).url
        assert "disabled=true" in url
        assert "execute=false" in url

    def test_omits_unset_parameters(self) -> None:
        assert parsed(owner=None).url == "https://api.example.com/users"

    def test_adds_query_parameters_for_delete(self) -> None:
        assert "force=true" in parsed(method="DELETE", force=True).url


class TestParseBody:
    def test_collects_leftover_params_into_the_body(self) -> None:
        result = parsed(method="POST", name="John", email="john@example.com")
        assert result.body is not None
        assert json.loads(result.body) == {"name": "John", "email": "john@example.com"}
        assert result.headers["content-type"] == "application/json; charset=utf-8"

    def test_prefers_an_explicit_body(self) -> None:
        result = parsed(method="POST", body={"custom": "data"}, ignored="x")
        assert result.body is not None
        assert json.loads(result.body) == {"custom": "data"}

    def test_excludes_url_variables_from_the_body(self) -> None:
        result = parsed(method="POST", url="/users/{id}", id="123", name="John")
        assert result.url == "https://api.example.com/users/123"
        assert result.body is not None
        assert json.loads(result.body) == {"name": "John"}

    def test_no_body_when_nothing_is_left_over(self) -> None:
        result = parsed(method="POST", url="/users/{id}", id="123")
        assert result.body is None
        assert "content-type" not in result.headers


class TestMerge:
    def test_parses_method_and_path_route_strings(self) -> None:
        result = merge(DEFAULTS, "POST /users")
        assert result["method"] == "POST"
        assert result["url"] == "/users"

    def test_keeps_the_default_method_for_a_bare_path(self) -> None:
        result = merge(DEFAULTS, "/users")
        assert result["method"] == "GET"
        assert result["url"] == "/users"

    def test_merges_params(self) -> None:
        result = merge(DEFAULTS, "/users", {"page": 1, "limit": 10})
        assert result["page"] == 1
        assert result["limit"] == 10

    def test_merges_headers_with_defaults(self) -> None:
        result = merge(DEFAULTS, "/users", {"headers": {"authorization": "Bearer t"}})
        assert result["headers"]["accept"] == "application/json"
        assert result["headers"]["authorization"] == "Bearer t"

    def test_lowercases_header_names(self) -> None:
        result = merge(
            DEFAULTS,
            "/users",
            {
                "headers": {
                    "Authorization": "Bearer t",
                    "Content-Type": "application/json",
                }
            },
        )
        assert result["headers"]["authorization"] == "Bearer t"
        assert result["headers"]["content-type"] == "application/json"

    def test_params_override_defaults(self) -> None:
        result = merge(
            DEFAULTS,
            "/users",
            {"base_url": "https://other.com", "headers": {"accept": "text/plain"}},
        )
        assert result["base_url"] == "https://other.com"
        assert result["headers"]["accept"] == "text/plain"

    def test_accepts_an_options_mapping(self) -> None:
        result = merge(DEFAULTS, {"method": "POST", "url": "/items"})
        assert (result["method"], result["url"]) == ("POST", "/items")

    def test_preserves_base_url_from_defaults(self) -> None:
        assert merge(DEFAULTS, "/users")["base_url"] == "https://api.example.com"

    def test_does_not_mutate_the_defaults(self) -> None:
        merge(DEFAULTS, "/users", {"headers": {"x-custom": "v"}})
        assert DEFAULTS["headers"] == {"accept": "application/json"}


@pytest.mark.parametrize(
    ("url", "expected"),
    [
        ("/users", []),
        ("/users/{id}", ["id"]),
        ("/tasks/{task_id}/assets/{asset_id}", ["task_id", "asset_id"]),
    ],
)
def test_extract_url_variable_names(url: str, expected: list[str]) -> None:
    from opti_cmp._utils import extract_url_variable_names

    assert extract_url_variable_names(url) == expected


class TestHeaderMerging:
    def test_layers_merge_regardless_of_case(self) -> None:
        result = merge(
            {"headers": Headers({"accept": "application/json"}), "method": "GET"},
            {"headers": {"Accept": "text/plain", "X-Org": "acme"}},
        )
        assert result.headers["ACCEPT"] == "text/plain"
        assert result.headers["x-org"] == "acme"

    def test_parse_does_not_mutate_the_caller_s_headers(self) -> None:
        headers = Headers({"accept": "application/json"})
        parse({**DEFAULTS, "headers": headers, "method": "POST", "body": {"a": 1}})
        assert "content-type" not in headers
