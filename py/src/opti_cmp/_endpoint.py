"""Turns endpoint parameters into a concrete request."""

from __future__ import annotations

import json
from collections.abc import Mapping
from typing import Any
from urllib.parse import quote, urlencode

from ._types import EndpointOptions, EndpointParams, Headers, RequestOptions
from ._utils import deep_merge, expand_url, extract_url_variable_names
from .version import VERSION

JSON_CONTENT_TYPE = "application/json; charset=utf-8"

DEFAULTS: EndpointParams = {
    "base_url": "https://api.cmp.optimizely.com/v3",
    "headers": Headers(
        {
            "accept": "application/json",
            "user-agent": f"opti-cmp.py/{VERSION}",
        }
    ),
    "method": "GET",
    "url": "",
}

#: Keys that describe the request rather than being part of it. Everything else
#: in the options mapping becomes a query parameter or body field.
#: `namespace`/`operation` identify the generated endpoint for hooks;
#: verified against the specification, neither collides with a real CMP
#: parameter.
_RESERVED = frozenset(
    {
        "base_url",
        "body",
        "headers",
        "method",
        "namespace",
        "operation",
        "request",
        "url",
    }
)


class Endpoint:
    """Immutable holder of endpoint defaults, chainable via :meth:`defaults`."""

    def __init__(self, defaults: EndpointParams) -> None:
        self.DEFAULTS = defaults

    def defaults(self, params: EndpointParams) -> Endpoint:
        return Endpoint(merge(self.DEFAULTS, params))

    def merge(
        self,
        route: EndpointParams | str | None = None,
        params: EndpointParams | None = None,
    ) -> EndpointOptions:
        return merge(self.DEFAULTS, route, params)

    def parse(self, options: EndpointParams) -> RequestOptions:
        return parse(options)

    def __call__(
        self,
        route: EndpointParams | str | None = None,
        params: EndpointParams | None = None,
    ) -> RequestOptions:
        return parse(self.merge(route, params))


def merge(
    defaults: EndpointParams | None,
    route: EndpointParams | str | None = None,
    params: EndpointParams | None = None,
) -> EndpointOptions:
    options: EndpointParams

    if isinstance(route, str):
        method, _, url = route.partition(" ")
        options = dict(params or {})
        if url:
            options["method"] = method.upper()
            options["url"] = url
        else:
            options["url"] = method
    else:
        options = dict(route or {})

    merged = deep_merge(defaults, options) if defaults else dict(options)
    # Headers merge rather than replace, and `Headers` is not a dict, so
    # `deep_merge` cannot do it.
    headers = Headers()
    for source in ((defaults or {}).get("headers"), options.get("headers")):
        if source:
            headers.update(source)
    merged["headers"] = headers
    return EndpointOptions(merged)


def parse(options: EndpointParams) -> RequestOptions:
    base_url: str = options.get("base_url") or ""
    method: str = options["method"]
    # Copied, so `parse` adding content-type does not leak into the caller's
    # options.
    headers = Headers(options.get("headers") or {})
    request = options.get("request")
    raw_body = options.get("body")
    url_template: str = options.get("url") or ""

    params = {key: value for key, value in options.items() if key not in _RESERVED}

    url = expand_url(url_template, params)
    if not url.startswith("http"):
        url = f"{base_url}{url}"

    url_variable_names = extract_url_variable_names(url_template)
    remaining = {
        key: value
        for key, value in params.items()
        if key not in url_variable_names and value is not None
    }

    body: Any = None
    if method in ("DELETE", "GET"):
        url = _add_query_parameters(url, remaining)
    elif raw_body is not None:
        body = raw_body
    elif remaining:
        body = remaining

    encoded: bytes | None = None
    if body is not None:
        headers["content-type"] = JSON_CONTENT_TYPE
        encoded = json.dumps(body).encode()

    return RequestOptions(
        method=method,
        url=url,
        headers=headers,
        body=encoded,
        request=request,
    )


def _add_query_parameters(url: str, params: Mapping[str, Any]) -> str:
    query = urlencode(
        {name: _stringify(value) for name, value in params.items()},
        doseq=True,  # a sequence value repeats the key
        quote_via=quote,  # percent-encode spaces rather than using `+`
    )
    if not query:
        return url

    separator = "&" if "?" in url else "?"
    return f"{url}{separator}{query}"


def _stringify(value: Any) -> Any:
    """Render a parameter the way the API expects before it is encoded."""
    if isinstance(value, (list, tuple)):
        return [_stringify(item) for item in value]
    # JSON/JS spelling, not Python's `True`/`False`.
    if isinstance(value, bool):
        return "true" if value else "false"
    return str(value)


endpoint = Endpoint(merge(None, DEFAULTS))
