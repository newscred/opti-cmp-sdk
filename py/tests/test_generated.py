"""Guards the generated tree: parity with the JS SDK, and regeneration stability."""

from __future__ import annotations

import inspect
import json
import keyword
import os
import re
import subprocess
import sys
from pathlib import Path
from typing import Any

import pytest

from opti_cmp import VERSION
from opti_cmp._generated import routes
from opti_cmp._generated.aio import NAMESPACES as ASYNC_NAMESPACES
from opti_cmp._generated.sync import NAMESPACES as SYNC_NAMESPACES
from opti_cmp._utils import extract_url_variable_names

_CAMEL_BOUNDARY = re.compile(r"(?<=[a-z0-9])(?=[A-Z])")


def to_snake_case(name: str) -> str:
    """Bridge the JS SDK's camelCase identifiers to this SDK's."""
    return _CAMEL_BOUNDARY.sub("_", name).lower()


#: Restated rather than imported from `scripts/_spec.py`, so the test asserts
#: the naming rule independently instead of agreeing with itself.
_RESERVED_PARAM_NAMES = frozenset({"self", "body", "headers", "request"})


def to_argument_name(variable: str) -> str:
    name = to_snake_case(re.sub(r"[-\s]+", "_", variable))
    if keyword.iskeyword(name) or name in _RESERVED_PARAM_NAMES:
        return f"{name}_"
    return name


PY_DIR = Path(__file__).resolve().parents[1]
JS_ROUTES = (
    PY_DIR.parent / "js" / "src" / "plugins" / "register-api-endpoints" / "routes.json"
)


@pytest.fixture(scope="module")
def js_routes() -> dict[str, Any]:
    """The JS SDK's routes, re-keyed to this SDK's snake_case identifiers.

    The two SDKs spell namespaces and operations differently — camelCase there,
    snake_case here — so parity is checked after normalising. Both are derived
    from the same operationIds, and the conversion is collision-free across all
    19 namespaces and 170 operations, so this compares the same set of
    endpoints, not a lossy projection of it.
    """
    if not JS_ROUTES.exists():
        if os.environ.get("CI"):
            # Skipping here would drop the whole cross-SDK guarantee silently.
            pytest.fail(f"JS SDK routes.json is missing at {JS_ROUTES}")
        pytest.skip("JS SDK routes.json is not present")
    parsed: dict[str, Any] = json.loads(JS_ROUTES.read_text())
    return {
        to_snake_case(namespace): {
            to_snake_case(operation): route for operation, route in operations.items()
        }
        for namespace, operations in parsed.items()
    }


class TestParityWithJsSdk:
    def test_normalising_does_not_merge_distinct_names(
        self, js_routes: dict[str, Any]
    ) -> None:
        # Guards the comparison itself: if snake_case ever collapsed two
        # distinct operations, the parity checks below would silently weaken.
        source: dict[str, Any] = json.loads(JS_ROUTES.read_text())
        assert len(source) == len(js_routes)
        for namespace, operations in source.items():
            assert len(operations) == len(js_routes[to_snake_case(namespace)])

    def test_same_namespaces(self, js_routes: dict[str, Any]) -> None:
        assert set(js_routes) == set(routes.ROUTES)

    def test_same_operations_per_namespace(self, js_routes: dict[str, Any]) -> None:
        for namespace, operations in js_routes.items():
            assert set(operations) == set(routes.ROUTES[namespace]), namespace

    def test_same_method_and_url(self, js_routes: dict[str, Any]) -> None:
        for namespace, operations in js_routes.items():
            for operation, route in operations.items():
                generated = routes.ROUTES[namespace][operation]
                assert generated["method"] == route["method"], (namespace, operation)
                assert generated["url"] == route["url"], (namespace, operation)


class TestNamespaceModules:
    def test_sync_and_async_expose_the_same_namespaces(self) -> None:
        assert set(SYNC_NAMESPACES) == set(ASYNC_NAMESPACES)

    def test_every_route_has_a_generated_method(self) -> None:
        for namespace, operations in routes.ROUTES.items():
            for flavor in (SYNC_NAMESPACES, ASYNC_NAMESPACES):
                cls = flavor[namespace]
                for operation in operations:
                    assert hasattr(cls, operation), (namespace, operation)

    def test_signatures_cover_every_url_variable(self) -> None:
        """A method missing a `{var}` its URL needs would expand to an empty
        segment at runtime, so check all 170 rather than the handful the
        end-to-end tests exercise."""
        for namespace, operations in routes.ROUTES.items():
            for flavor in (SYNC_NAMESPACES, ASYNC_NAMESPACES):
                cls = flavor[namespace]
                for operation, route in operations.items():
                    parameters = inspect.signature(getattr(cls, operation)).parameters
                    for variable in extract_url_variable_names(route["url"]):
                        assert to_argument_name(variable) in parameters, (
                            namespace,
                            operation,
                            variable,
                        )

    def test_required_url_variables_are_positional(self) -> None:
        """A path variable has no sensible default, so it must be required."""
        for namespace, operations in routes.ROUTES.items():
            cls = SYNC_NAMESPACES[namespace]
            for operation, route in operations.items():
                parameters = inspect.signature(getattr(cls, operation)).parameters
                for variable in extract_url_variable_names(route["url"]):
                    parameter = parameters[to_argument_name(variable)]
                    assert parameter.default is inspect.Parameter.empty, (
                        namespace,
                        operation,
                        variable,
                    )

    def test_methods_carry_docstrings(self) -> None:
        list_campaigns = SYNC_NAMESPACES["campaign"].list_campaigns
        assert list_campaigns.__doc__
        assert "GET /campaigns" in list_campaigns.__doc__


def test_version_matches_pyproject() -> None:
    """release-please bumps pyproject.toml and version.py through two separate
    updaters with no cross-check. The JS side already shipped four releases
    reporting 0.0.0 in its User-Agent because one of them silently did nothing.
    """
    # Parsed with a regex rather than tomllib, which is 3.11+ and this package
    # supports 3.10.
    text = (PY_DIR / "pyproject.toml").read_text()
    match = re.search(r'^version = "([^"]+)"', text, flags=re.MULTILINE)
    assert match, "no version in pyproject.toml"
    assert match.group(1) == VERSION


@pytest.mark.slow
def test_regeneration_is_stable() -> None:
    """Regenerating must not change the committed tree.

    Compares file contents rather than `git diff`, so it still catches a stale
    tree when the generated files are not tracked yet.

    Note this runs the real generators, which **rewrite** the 44 files under
    `_generated/` in place. They are deterministic, so a passing run leaves the
    tree byte-identical — but a failing one leaves it modified.
    """
    generated = PY_DIR / "src" / "opti_cmp" / "_generated"
    before = _snapshot(generated)

    for script in ("generate_schema_types.py", "generate_endpoints.py"):
        subprocess.run([sys.executable, f"scripts/{script}"], cwd=PY_DIR, check=True)

    after = _snapshot(generated)

    assert set(before) == set(after), "generated file set changed"
    changed = sorted(name for name in before if before[name] != after[name])
    assert not changed, f"generated tree is stale: {changed}"


def _snapshot(directory: Path) -> dict[str, str]:
    return {
        str(path.relative_to(directory)): path.read_text()
        for path in sorted(directory.rglob("*.py"))
    }
