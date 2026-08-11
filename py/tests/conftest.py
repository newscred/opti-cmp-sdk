from __future__ import annotations

from collections.abc import Callable
from typing import Any

import httpx


def transport(
    handler: Callable[[httpx.Request], httpx.Response],
    recorder: list[httpx.Request] | None = None,
) -> httpx.MockTransport:
    def wrapped(request: httpx.Request) -> httpx.Response:
        if recorder is not None:
            recorder.append(request)
        return handler(request)

    return httpx.MockTransport(wrapped)


def json_transport(
    payload: Any,
    *,
    status: int = 200,
    recorder: list[httpx.Request] | None = None,
) -> httpx.MockTransport:
    return transport(lambda _request: httpx.Response(status, json=payload), recorder)


def recording() -> tuple[list[httpx.Request], httpx.MockTransport]:
    """An empty-200 transport and the list it records requests into."""
    seen: list[httpx.Request] = []
    return seen, json_transport({}, recorder=seen)
