from __future__ import annotations

from collections.abc import Mapping, Sequence
from dataclasses import dataclass, field
from typing import TYPE_CHECKING, Any, Generic, Protocol, TypedDict, TypeVar

import httpx

from ._hook import AsyncSingularHook, SingularHook
from ._object import AttrDict

#: Case-insensitive and multi-value aware, so a response's repeated headers
#: survive and `headers["X-Request-Id"]` works whatever case the server used.
Headers = httpx.Headers

#: Anything `Headers` can be built from. Mirrors httpx's own `HeaderTypes`,
#: which lives in a private module, so it is restated here rather than imported.
#: A sequence of pairs is how you set the same header more than once.
HeaderTypes = Headers | Mapping[str, str] | Sequence[tuple[str, str]]

#: Endpoint parameters are a flat mapping: reserved keys (``base_url``,
#: ``headers``, ``method``, ``url``, ``request``, ``body``) plus the endpoint's
#: own path/query parameters. Verified against the specification: no CMP
#: parameter collides with a reserved key.
EndpointParams = dict[str, Any]

T = TypeVar("T")


class RequestConfig(TypedDict, total=False):
    """Per-request transport options."""

    timeout: float | None
    #: Planted by the client; `Request.__call__` reads it back out.
    hook: SingularHook | AsyncSingularHook


class EndpointOptions(AttrDict):
    """The merged options a request is built from, as hooks receive them.

    A `dict`, so it still flows through merging and parsing unchanged, but the
    parts the SDK understands are also readable and writable as attributes::

        options.headers["authorization"] = f"Bearer {token}"
        options.namespace   # "campaign", or None for a raw client.request()

    Anything else in the mapping is the call's own parameters, reachable by key.
    """

    __slots__ = ()

    method: str
    url: str
    headers: Headers
    request: RequestConfig
    #: Set for generated endpoint calls only.
    namespace: str | None
    operation: str | None

    if not TYPE_CHECKING:
        _DECLARED = frozenset(
            {"method", "url", "headers", "request", "namespace", "operation"}
        )

        def __getattr__(self, name):
            # A declared field that is absent reads as None, so the annotations
            # above are honest. Anything else falls through to AttrDict, which
            # treats it as a mistake.
            if name not in self and name in EndpointOptions._DECLARED:
                return None
            return super().__getattr__(name)


@dataclass
class RequestOptions:
    """A fully resolved request, ready for the transport."""

    method: str
    url: str
    headers: Headers = field(default_factory=Headers)
    #: Out of `repr` because the OAuth token request carries `client_secret`
    #: in here; `Headers` redacts `authorization` on its own.
    body: bytes | None = field(default=None, repr=False)
    request: RequestConfig | None = None


@dataclass
class Response(Generic[T]):
    data: T
    headers: Headers
    status: int
    url: str


class Route(TypedDict):
    method: str
    url: str


Routes = dict[str, dict[str, Route]]


class Dispatcher(Protocol):
    """The seam every generated endpoint method calls."""

    def _dispatch(
        self, *, namespace: str, operation: str, params: dict[str, Any]
    ) -> Response[Any]: ...


class AsyncDispatcher(Protocol):
    async def _dispatch(
        self, *, namespace: str, operation: str, params: dict[str, Any]
    ) -> Response[Any]: ...
