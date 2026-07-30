"""Mechanics of the request hook. See `test_hook_integration.py` for usage."""

from __future__ import annotations

from typing import Any

import pytest

from opti_cmp import EndpointOptions
from opti_cmp._hook import AsyncSingularHook, SingularHook


def double(options: EndpointOptions) -> int:
    value: int = options["value"]
    return value * 2


def _observe(sink: list[BaseException], error: BaseException) -> None:
    """An observer returns None, so the error must keep propagating."""
    sink.append(error)


def _boom(_options: EndpointOptions) -> int:
    raise RuntimeError("boom")


async def _aboom(_options: EndpointOptions) -> int:
    raise RuntimeError("boom")


class TestSingularHook:
    def test_calls_the_method_when_nothing_is_registered(self) -> None:
        assert SingularHook()(double, EndpointOptions({"value": 3})) == 6

    def test_before_hooks_can_mutate_the_options(self) -> None:
        hook = SingularHook()
        hook.before(lambda options: options.__setitem__("value", 10))
        assert hook(double, EndpointOptions({"value": 3})) == 20

    def test_before_hooks_run_in_registration_order(self) -> None:
        order: list[str] = []
        hook = SingularHook()
        hook.before(lambda _o: order.append("first"))
        hook.before(lambda _o: order.append("second"))
        hook(double, EndpointOptions({"value": 1}))
        assert order == ["first", "second"]

    def test_after_hooks_can_replace_the_result(self) -> None:
        hook = SingularHook()
        hook.after(lambda result, _options: result + 1)
        assert hook(double, EndpointOptions({"value": 3})) == 7

    def test_after_hooks_returning_none_keep_the_result(self) -> None:
        hook = SingularHook()
        hook.after(lambda _result, _options: None)
        assert hook(double, EndpointOptions({"value": 3})) == 6

    def test_errors_propagate_when_no_error_hook_is_registered(self) -> None:
        hook = SingularHook()

        with pytest.raises(RuntimeError, match="boom"):
            hook(_boom, EndpointOptions())

    def test_error_hooks_can_recover(self) -> None:
        hook = SingularHook()
        hook.error(lambda _error, _options: -1)
        assert hook(_boom, EndpointOptions()) == -1

    def test_observe_only_error_hooks_do_not_swallow(self) -> None:
        # The most likely hook an SDK user writes: log it and move on. Returning
        # None must not turn the exception into a None result.
        seen: list[BaseException] = []
        hook = SingularHook()
        hook.error(lambda error, _options: _observe(seen, error))

        with pytest.raises(RuntimeError, match="boom"):
            hook(_boom, EndpointOptions())

        assert len(seen) == 1

    def test_every_error_hook_runs(self) -> None:
        # Observers and a recovery hook coexist: a logger must not stop metrics
        # from firing, and neither must stop recovery.
        calls: list[str] = []
        hook = SingularHook()

        def recover(_error: BaseException, _options: EndpointOptions) -> int:
            calls.append("recovery")
            return -1

        hook.error(lambda _e, _o: calls.append("logger"))
        hook.error(lambda _e, _o: calls.append("metrics"))
        hook.error(recover)

        assert hook(_boom, EndpointOptions()) == -1
        assert calls == ["logger", "metrics", "recovery"]

    def test_first_recovering_hook_wins(self) -> None:
        hook = SingularHook()
        hook.error(lambda _e, _o: "first")
        hook.error(lambda _e, _o: "second")
        assert hook(_boom, EndpointOptions()) == "first"

    def test_wrap_hooks_surround_the_method(self) -> None:
        hook = SingularHook()
        hook.wrap(lambda inner, options: inner(options) + 100)
        assert hook(double, EndpointOptions({"value": 3})) == 106

    def test_wrap_encloses_before_hooks(self) -> None:
        """A retry wrap has to re-run them: OAuth fetches its token in one."""
        order: list[str] = []
        hook = SingularHook()
        hook.before(lambda _options: order.append("before"))

        def record(inner: Any, options: EndpointOptions) -> int:
            order.append("wrap-in")
            result: int = inner(options)
            order.append("wrap-out")
            return result

        hook.wrap(record)
        hook(double, EndpointOptions({"value": 3}))

        assert order == ["wrap-in", "before", "wrap-out"]

    def test_a_retry_wrap_reruns_before_hooks(self) -> None:
        attempts: list[int] = []
        hook = SingularHook()
        hook.before(lambda _options: attempts.append(len(attempts)))

        def retry_once(inner: Any, options: EndpointOptions) -> int:
            try:
                return inner(options)  # type: ignore[no-any-return]
            except RuntimeError:
                return inner(options)  # type: ignore[no-any-return]

        hook.wrap(retry_once)

        calls = 0

        def flaky(options: EndpointOptions) -> int:
            nonlocal calls
            calls += 1
            if calls == 1:
                raise RuntimeError("boom")
            return 42

        assert hook(flaky, EndpointOptions()) == 42
        assert len(attempts) == 2

    def test_remove_unregisters_a_callback(self) -> None:
        hook = SingularHook()

        def bump(options: EndpointOptions) -> None:
            options["value"] = 10

        hook.before(bump)
        hook.remove(bump)
        assert hook(double, EndpointOptions({"value": 3})) == 6


class TestAsyncSingularHook:
    async def adouble(self, options: EndpointOptions) -> int:
        value: int = options["value"]
        return value * 2

    async def test_calls_the_method(self) -> None:
        assert (
            await AsyncSingularHook()(self.adouble, EndpointOptions({"value": 3})) == 6
        )

    async def test_awaits_async_before_hooks(self) -> None:
        hook = AsyncSingularHook()

        async def bump(options: EndpointOptions) -> None:
            options["value"] = 10

        hook.before(bump)
        assert await hook(self.adouble, EndpointOptions({"value": 3})) == 20

    async def test_accepts_sync_before_hooks(self) -> None:
        hook = AsyncSingularHook()
        hook.before(lambda options: options.__setitem__("value", 5))
        assert await hook(self.adouble, EndpointOptions({"value": 3})) == 10

    async def test_after_hooks_can_replace_the_result(self) -> None:
        hook = AsyncSingularHook()
        hook.after(lambda result, _options: result + 1)
        assert await hook(self.adouble, EndpointOptions({"value": 3})) == 7

    async def test_error_hooks_can_recover(self) -> None:
        hook = AsyncSingularHook()
        hook.error(lambda _error, _options: -1)
        assert await hook(_aboom, EndpointOptions()) == -1

    async def test_errors_propagate_without_an_error_hook(self) -> None:
        hook = AsyncSingularHook()
        with pytest.raises(RuntimeError, match="boom"):
            await hook(_aboom, EndpointOptions())

    async def test_observe_only_error_hooks_do_not_swallow(self) -> None:
        seen: list[BaseException] = []
        hook = AsyncSingularHook()

        async def observe(error: BaseException, _options: EndpointOptions) -> None:
            seen.append(error)

        hook.error(observe)

        with pytest.raises(RuntimeError, match="boom"):
            await hook(_aboom, EndpointOptions())

        assert len(seen) == 1

    async def test_before_hooks_run_in_registration_order(self) -> None:
        order: list[str] = []
        hook = AsyncSingularHook()
        hook.before(lambda _options: order.append("first"))

        async def second(_options: EndpointOptions) -> None:
            order.append("second")

        hook.before(second)
        await hook(self.adouble, EndpointOptions({"value": 1}))

        assert order == ["first", "second"]

    async def test_every_error_hook_runs(self) -> None:
        seen: list[str] = []
        hook = AsyncSingularHook()
        hook.error(lambda _error, _options: seen.append("a"))
        hook.error(lambda _error, _options: seen.append("b"))

        with pytest.raises(RuntimeError, match="boom"):
            await hook(_aboom, EndpointOptions())

        assert seen == ["a", "b"]

    async def test_the_first_recovering_error_hook_wins(self) -> None:
        hook = AsyncSingularHook()
        hook.error(lambda _error, _options: "first")
        hook.error(lambda _error, _options: "second")
        assert await hook(_aboom, EndpointOptions()) == "first"

    async def test_wrap_hooks_surround_the_method(self) -> None:
        hook = AsyncSingularHook()

        async def add_hundred(inner: Any, options: EndpointOptions) -> int:
            result: int = await inner(options)
            return result + 100

        hook.wrap(add_hundred)
        assert await hook(self.adouble, EndpointOptions({"value": 3})) == 106

    async def test_wrap_encloses_before_hooks(self) -> None:
        """A retry wrap has to re-run them: OAuth fetches its token in one."""
        order: list[str] = []
        hook = AsyncSingularHook()
        hook.before(lambda _options: order.append("before"))

        async def record(inner: Any, options: EndpointOptions) -> int:
            order.append("wrap-in")
            result: int = await inner(options)
            order.append("wrap-out")
            return result

        hook.wrap(record)
        await hook(self.adouble, EndpointOptions({"value": 3}))

        assert order == ["wrap-in", "before", "wrap-out"]

    async def test_remove_unregisters_a_callback(self) -> None:
        hook = AsyncSingularHook()

        def bump(options: EndpointOptions) -> None:
            options["value"] = 10

        hook.before(bump)
        hook.remove(bump)
        assert await hook(self.adouble, EndpointOptions({"value": 3})) == 6
