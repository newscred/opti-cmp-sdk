"""Client construction, plugin registration and endpoint dispatch."""

from __future__ import annotations

from collections.abc import Callable, Mapping
from types import TracebackType
from typing import TYPE_CHECKING, Any, ClassVar

import httpx

from ._endpoint import endpoint
from ._generated.aio import AsyncNamespaces
from ._generated.routes import ROUTES
from ._generated.sync import SyncNamespaces
from ._hook import AsyncSingularHook, SingularHook
from ._request import AsyncRequest, Request, asend, send
from ._types import (
    EndpointParams,
    Headers,
    HeaderTypes,
    RequestConfig,
    Response,
    Routes,
)
from ._utils import deep_merge, pick
from .plugins.auth import auth_plugin
from .plugins.register_api_endpoints import register_api_endpoints_plugin

if TYPE_CHECKING:
    import sys

    from .plugins.auth import AsyncOAuthMethods, AuthOptions, OAuthMethods

    # Annotation-only, and `from __future__ import annotations` means these are
    # never evaluated — so `typing_extensions` stays out of the runtime import.
    if sys.version_info >= (3, 11):
        from typing import Self
    else:
        from typing_extensions import Self

Plugin = Callable[[Any, dict[str, Any]], None]

DEFAULT_PLUGINS: tuple[Plugin, ...] = (
    register_api_endpoints_plugin,
    auth_plugin,
)

#: httpx would otherwise apply its own 5s to every phase, which is too tight
#: for uploads and large list queries. Override per client or per call with
#: ``request={"timeout": seconds}``; `None` disables the deadline entirely.
DEFAULT_TIMEOUT = httpx.Timeout(30.0, connect=10.0)


class DynamicNamespace:
    """Holds endpoints attached at runtime via :meth:`BaseClient.register_endpoints`."""


class BaseClient:
    _plugins: ClassVar[tuple[Plugin, ...]] = DEFAULT_PLUGINS

    #: Which flavour of client this is. Plugins that need to build sync or async
    #: machinery read this rather than inspecting the hook or transport, so the
    #: plugin contract does not depend on either one's concrete type.
    _is_async: ClassVar[bool]

    #: Installed by the auth plugin.
    authenticate: Callable[..., None]

    def __init__(
        self,
        *,
        auth: AuthOptions | None = None,
        base_url: str | None = None,
        headers: HeaderTypes | None = None,
        request: RequestConfig | None = None,
        transport: httpx.BaseTransport | httpx.AsyncBaseTransport | None = None,
        **extra: Any,
    ) -> None:
        """Build a client.

        `extra` carries options for plugins; everything the SDK itself
        understands is named above, so a misspelled keyword lands there rather
        than being rejected. Subclasses do not override this — they implement
        :meth:`_setup` — so the signature is declared once and both the
        synchronous and asynchronous clients are checked against it.
        """
        options: dict[str, Any] = {
            "auth": auth,
            "base_url": base_url,
            # Copied into Headers, not a dict: a dict would collapse repeated
            # header values into one comma-joined string.
            "headers": Headers(headers) if headers else None,
            "request": dict(request) if request else None,
            "transport": transport,
            **extra,
        }
        self.options: dict[str, Any] = {
            key: value for key, value in options.items() if value is not None
        }
        #: Routes added through :meth:`register_endpoints`, checked before the
        #: generated ones so plugins can add endpoints and override them.
        self._runtime_routes: Routes = {}

        self._setup()
        self._apply_plugins()

    def _setup(self) -> None:
        """Build the transport, hook and request pipeline for this flavour."""
        raise NotImplementedError

    @classmethod
    def plugins(cls, *plugins: Plugin) -> type[Self]:
        """Return a subclass with `plugins` appended, skipping duplicates."""
        registered = list(cls._plugins)
        for plugin in plugins:
            if plugin not in registered:
                registered.append(plugin)
        return type(cls.__name__, (cls,), {"_plugins": tuple(registered)})

    def register_endpoints(self, routes: Routes) -> None:
        """Attach the endpoints in `routes` as attributes, keys used verbatim.

        The extension point for plugins that add endpoints the specification
        does not describe. Generated endpoints do not go through it — they are
        attached directly by `register_api_endpoints_plugin`.

        Keys become attribute names as written, so use snake_case, matching
        everything else the SDK exposes.
        """
        for namespace, operations in routes.items():
            self._runtime_routes.setdefault(namespace, {}).update(operations)
            target = self._namespace(namespace)
            for operation in operations:
                setattr(target, operation, self._bind_operation(namespace, operation))

    def _bind_operation(self, namespace: str, operation: str) -> Callable[..., Any]:
        raise NotImplementedError

    def _namespace(self, namespace: str) -> Any:
        """The object to hang endpoints on, whatever kind it already is.

        Registering into a generated namespace adds to it rather than
        replacing it, so its generated methods survive.
        """
        existing = getattr(self, namespace, None)
        if existing is None:
            existing = DynamicNamespace()
            setattr(self, namespace, existing)
        return existing

    def _apply_plugins(self) -> None:
        for plugin in type(self)._plugins:
            plugin(self, self.options)

    def _endpoint_defaults(
        self, hook: SingularHook | AsyncSingularHook
    ) -> EndpointParams:
        defaults = deep_merge(
            {"headers": Headers(), "request": {}},
            pick(self.options, ["base_url", "headers", "request"]),
        )
        defaults["request"]["hook"] = hook
        return defaults

    def _route_options(
        self, namespace: str, operation: str, params: Mapping[str, Any]
    ) -> EndpointParams:
        registered = self._runtime_routes.get(namespace, {})
        route = registered.get(operation) or ROUTES[namespace][operation]
        options: EndpointParams = {
            "method": route["method"],
            # Carried through so hooks can label by logical operation rather
            # than by the concrete URL. Stripped before the request is built.
            "namespace": namespace,
            "operation": operation,
            "url": route["url"],
        }
        options.update(
            {key: value for key, value in params.items() if value is not None}
        )
        return options

    @staticmethod
    def _page_url(response: Response[Any], direction: str) -> str | None:
        data = response.data
        if not isinstance(data, Mapping):
            return None
        pagination = data.get("pagination")
        if not isinstance(pagination, Mapping):
            return None
        url = pagination.get(direction)
        # An empty string is "no page", not a request against the bare base URL.
        return url if isinstance(url, str) and url else None

    def has_next_page(self, response: Response[Any]) -> bool:
        return self._page_url(response, "next") is not None

    def has_previous_page(self, response: Response[Any]) -> bool:
        return self._page_url(response, "previous") is not None

    def _require_page_url(self, response: Response[Any], direction: str) -> str:
        url = self._page_url(response, direction)
        if url is None:
            # Caller misuse, not an HTTP failure — check has_next_page first.
            raise ValueError(f"No {direction} page available")
        return url


class OptiCMP(BaseClient, SyncNamespaces):
    """Synchronous Optimizely CMP API client."""

    _is_async: ClassVar[bool] = False

    request: Request
    request_hook: SingularHook
    oauth: OAuthMethods

    def _setup(self) -> None:
        self._httpx = httpx.Client(
            transport=self.options.get("transport"),
            follow_redirects=True,
            timeout=DEFAULT_TIMEOUT,
        )
        self.request_hook = SingularHook()
        self.request = Request(
            endpoint.defaults(self._endpoint_defaults(self.request_hook)),
            lambda resolved: send(self._httpx, resolved),
        )

    def _dispatch(
        self, *, namespace: str, operation: str, params: dict[str, Any]
    ) -> Response[Any]:
        return self.request(self._route_options(namespace, operation, params))

    def _bind_operation(self, namespace: str, operation: str) -> Callable[..., Any]:
        def call(**params: Any) -> Response[Any]:
            return self._dispatch(
                namespace=namespace, operation=operation, params=params
            )

        call.__name__ = operation
        return call

    def get_next_page(self, response: Response[Any]) -> Response[Any]:
        return self.request(self._require_page_url(response, "next"))

    def get_previous_page(self, response: Response[Any]) -> Response[Any]:
        return self.request(self._require_page_url(response, "previous"))

    def close(self) -> None:
        self._httpx.close()

    def __enter__(self) -> Self:
        return self

    def __exit__(
        self,
        exc_type: type[BaseException] | None,
        exc: BaseException | None,
        tb: TracebackType | None,
    ) -> None:
        self.close()


class AsyncOptiCMP(BaseClient, AsyncNamespaces):
    """Asynchronous Optimizely CMP API client."""

    _is_async: ClassVar[bool] = True

    request: AsyncRequest
    request_hook: AsyncSingularHook
    oauth: AsyncOAuthMethods

    def _setup(self) -> None:
        self._httpx = httpx.AsyncClient(
            transport=self.options.get("transport"),
            follow_redirects=True,
            timeout=DEFAULT_TIMEOUT,
        )
        self.request_hook = AsyncSingularHook()
        self.request = AsyncRequest(
            endpoint.defaults(self._endpoint_defaults(self.request_hook)),
            lambda resolved: asend(self._httpx, resolved),
        )

    async def _dispatch(
        self, *, namespace: str, operation: str, params: dict[str, Any]
    ) -> Response[Any]:
        return await self.request(self._route_options(namespace, operation, params))

    def _bind_operation(self, namespace: str, operation: str) -> Callable[..., Any]:
        async def call(**params: Any) -> Response[Any]:
            return await self._dispatch(
                namespace=namespace, operation=operation, params=params
            )

        call.__name__ = operation
        return call

    async def get_next_page(self, response: Response[Any]) -> Response[Any]:
        return await self.request(self._require_page_url(response, "next"))

    async def get_previous_page(self, response: Response[Any]) -> Response[Any]:
        return await self.request(self._require_page_url(response, "previous"))

    async def aclose(self) -> None:
        await self._httpx.aclose()

    async def __aenter__(self) -> Self:
        return self

    async def __aexit__(
        self,
        exc_type: type[BaseException] | None,
        exc: BaseException | None,
        tb: TracebackType | None,
    ) -> None:
        await self.aclose()
