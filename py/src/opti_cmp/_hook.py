"""The request hook: the SDK's public extension point.

A single extension point wrapped around each request. Plugins and callers
register callbacks; see :meth:`_HookRegistry.error` for the one rule worth
reading before you write one.

Callbacks can be registered by call or as decorators::

    @client.request_hook.before
    def add_tracing(options):
        options.headers["x-trace-id"] = new_id()
"""

from __future__ import annotations

from collections.abc import Awaitable, Callable
from functools import partial
from typing import TYPE_CHECKING, Any, TypeVar

if TYPE_CHECKING:
    # Import only for annotations: _types imports this module at runtime.
    from ._types import EndpointOptions

#: Receives the merged endpoint options and may mutate them in place.
BeforeHook = Callable[["EndpointOptions"], Any]
#: Receives ``(result, options)``. Return non-None to replace the result.
AfterHook = Callable[[Any, "EndpointOptions"], Any]
#: Receives ``(error, options)``. Return non-None to recover, None to re-raise.
ErrorHook = Callable[[BaseException, "EndpointOptions"], Any]
#: Receives ``(inner, options)`` and is responsible for calling ``inner``.
WrapHook = Callable[[Callable[["EndpointOptions"], Any], "EndpointOptions"], Any]

_C = TypeVar("_C", bound=Callable[..., Any])


class _HookRegistry:
    def __init__(self) -> None:
        self._before: list[Any] = []
        self._after: list[Any] = []
        self._error: list[Any] = []
        self._wrap: list[Any] = []

    def before(self, callback: _C) -> _C:
        """Run before the request. Mutate `options` in place to change it."""
        self._before.append(callback)
        return callback

    def after(self, callback: _C) -> _C:
        """Run on success. Return non-None to replace the response."""
        self._after.append(callback)
        return callback

    def error(self, callback: _C) -> _C:
        """Run on failure. Return non-None to recover, None to let it raise.

        Every registered error hook runs, so observers (logging, metrics) and a
        recovery hook can coexist. Returning None means "observed, no change",
        matching :meth:`after` — an error hook that only logs will not swallow
        the exception. Recovery is opt-in, and explicit.
        """
        self._error.append(callback)
        return callback

    def wrap(self, callback: _C) -> _C:
        """Wrap the whole call, before hooks included.

        Receives ``(inner, options)`` and must call `inner`. Before hooks run
        inside the wrap chain, so a retry wrap retries them too — which matters
        because the auth plugin fetches its token from a before hook.
        """
        self._wrap.append(callback)
        return callback

    def remove(self, callback: Any) -> None:
        for registry in (self._before, self._after, self._error, self._wrap):
            if callback in registry:
                registry.remove(callback)


class SingularHook(_HookRegistry):
    """Synchronous request hook, exposed as ``client.request_hook``."""

    def __call__(
        self, method: Callable[[EndpointOptions], Any], options: EndpointOptions
    ) -> Any:
        def call(options: EndpointOptions) -> Any:
            for before in self._before:
                before(options)
            return method(options)

        invoke: Callable[[EndpointOptions], Any] = call
        # Later-registered wrappers sit closest to the method.
        for wrapper in reversed(self._wrap):
            invoke = partial(wrapper, invoke)

        try:
            result = invoke(options)
        except Exception as error:
            recovered = None
            for handler in self._error:
                outcome = handler(error, options)
                if outcome is not None and recovered is None:
                    recovered = outcome
            if recovered is None:
                raise
            result = recovered

        for after in self._after:
            replacement = after(result, options)
            if replacement is not None:
                result = replacement
        return result


class AsyncSingularHook(_HookRegistry):
    """Asynchronous request hook. Callbacks may be sync or coroutine functions."""

    async def __call__(
        self,
        method: Callable[[EndpointOptions], Awaitable[Any]],
        options: EndpointOptions,
    ) -> Any:
        async def call(options: EndpointOptions) -> Any:
            for before in self._before:
                await _resolve(before(options))
            return await _resolve(method(options))

        invoke: Callable[[EndpointOptions], Any] = call
        for wrapper in reversed(self._wrap):
            invoke = _bind_async_wrap(wrapper, invoke)

        try:
            result = await _resolve(invoke(options))
        except Exception as error:
            recovered = None
            for handler in self._error:
                outcome = await _resolve(handler(error, options))
                if outcome is not None and recovered is None:
                    recovered = outcome
            if recovered is None:
                raise
            result = recovered

        for after in self._after:
            replacement = await _resolve(after(result, options))
            if replacement is not None:
                result = replacement
        return result


async def _resolve(value: Any) -> Any:
    if isinstance(value, Awaitable):
        return await value
    return value


def _bind_async_wrap(
    wrapper: Any, inner: Callable[[EndpointOptions], Any]
) -> Callable[[EndpointOptions], Any]:
    async def wrapped(options: EndpointOptions) -> Any:
        return await _resolve(wrapper(inner, options))

    return wrapped
