"""Sends resolved requests over httpx."""

from __future__ import annotations

import json
from collections.abc import Awaitable, Callable
from typing import Any

import httpx

from ._endpoint import Endpoint
from ._errors import (
    APIConnectionError,
    APIError,
    APITimeoutError,
    InvalidRequestError,
    OptiCMPError,
)
from ._object import APIObject
from ._types import EndpointParams, RequestOptions, Response


class Request:
    """Synchronous request builder. Chainable via :meth:`defaults`."""

    def __init__(
        self,
        endpoint: Endpoint,
        send: Callable[[RequestOptions], Response[Any]],
    ) -> None:
        self.endpoint = endpoint
        self._send = send

    def defaults(self, params: EndpointParams) -> Request:
        return Request(self.endpoint.defaults(params), self._send)

    def __call__(
        self,
        route: EndpointParams | str | None = None,
        params: EndpointParams | None = None,
    ) -> Response[Any]:
        options = self.endpoint.merge(route, params)
        hook = (options.get("request") or {}).get("hook")

        if hook is None:
            return self._send(self.endpoint.parse(options))

        result: Response[Any] = hook(
            lambda opts: self._send(self.endpoint.parse(opts)), options
        )
        return result


class AsyncRequest:
    """Asynchronous request builder. Chainable via :meth:`defaults`."""

    def __init__(
        self,
        endpoint: Endpoint,
        send: Callable[[RequestOptions], Awaitable[Response[Any]]],
    ) -> None:
        self.endpoint = endpoint
        self._send = send

    def defaults(self, params: EndpointParams) -> AsyncRequest:
        return AsyncRequest(self.endpoint.defaults(params), self._send)

    async def __call__(
        self,
        route: EndpointParams | str | None = None,
        params: EndpointParams | None = None,
    ) -> Response[Any]:
        options = self.endpoint.merge(route, params)
        hook = (options.get("request") or {}).get("hook")

        if hook is None:
            return await self._send(self.endpoint.parse(options))

        result: Response[Any] = await hook(
            lambda opts: self._send(self.endpoint.parse(opts)), options
        )
        return result


def send(client: httpx.Client, options: RequestOptions) -> Response[Any]:
    try:
        response = client.request(**_httpx_kwargs(options))
    except (httpx.HTTPError, httpx.InvalidURL) as error:
        raise transport_error(error, options) from error

    return handle_response(response, options)


async def asend(client: httpx.AsyncClient, options: RequestOptions) -> Response[Any]:
    try:
        response = await client.request(**_httpx_kwargs(options))
    except (httpx.HTTPError, httpx.InvalidURL) as error:
        raise transport_error(error, options) from error

    return handle_response(response, options)


def _httpx_kwargs(options: RequestOptions) -> dict[str, Any]:
    kwargs: dict[str, Any] = {
        "method": options.method,
        "url": options.url,
        "headers": options.headers,
        "content": options.body,
    }
    # Seconds, the convention every Python HTTP client uses. Key presence, not
    # truthiness: an explicit `timeout: None` disables the deadline, while an
    # absent one leaves the client's own default in place.
    request = options.request or {}
    if "timeout" in request:
        kwargs["timeout"] = request["timeout"]
    return kwargs


def handle_response(response: httpx.Response, options: RequestOptions) -> Response[Any]:
    """Turn an httpx response into a `Response`, or raise `APIError`.

    Shared with the auth plugin, so token and userinfo calls report failures
    the same way endpoint calls do.
    """
    # Handed over as-is: a plain dict would collapse repeated headers such as
    # Set-Cookie into one comma-joined value.
    headers = response.headers
    data = _response_data(response)

    if not response.is_success:
        raise APIError(
            response.reason_phrase,
            response.status_code,
            data=data,
            headers=headers,
            request=options,
        )

    return Response(
        data=data,
        headers=headers,
        status=response.status_code,
        url=str(response.url),
    )


def _response_data(response: httpx.Response) -> Any:
    if response.status_code == 204:
        return None

    content_type = response.headers.get("content-type", "")

    if "application/json" in content_type:
        try:
            # `object_hook` builds the APIObject tree during the single C-level
            # parse, rather than walking a finished dict tree to copy it.
            return response.json(object_hook=APIObject)
        except json.JSONDecodeError:
            # A malformed body is more useful to the caller as text than as an
            # exception raised from inside the SDK.
            return response.text

    if content_type.startswith("text/") or "charset=utf-8" in content_type:
        return response.text

    return response.content


def transport_error(
    error: httpx.HTTPError | httpx.InvalidURL, options: RequestOptions
) -> OptiCMPError:
    """No response arrived, so there is no status to report."""
    if isinstance(error, httpx.TimeoutException):
        return APITimeoutError(str(error), request=options)
    # `InvalidURL` does not derive from `httpx.HTTPError`, and a
    # `LocalProtocolError` means the client refused to send what we built —
    # both are caller input, so neither is reported as a connection failure.
    if isinstance(error, httpx.InvalidURL | httpx.LocalProtocolError):
        return InvalidRequestError(str(error), request=options)
    return APIConnectionError(str(error), request=options)
