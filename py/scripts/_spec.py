"""Shared loading and naming helpers for the code generators.

Reads the language-neutral artifacts under `specification/`, the same ones
`js/scripts/*` consume.
"""

from __future__ import annotations

import json
import keyword
import re
import subprocess
import sys
from collections.abc import Iterator
from dataclasses import dataclass
from pathlib import Path
from typing import Any

ROOT = Path(__file__).resolve().parents[2]
SPEC_DIR = ROOT / "specification"
PACKAGE_DIR = Path(__file__).resolve().parents[1] / "src" / "opti_cmp"
GENERATED_DIR = PACKAGE_DIR / "_generated"

HTTP_METHODS = ("get", "post", "put", "patch", "delete")
SUCCESS_CODES = ("200", "201", "202", "204")

GENERATED_MARKER = "# Auto-generated - DO NOT EDIT"

# Names that would shadow the client's own endpoint-call parameters.
RESERVED_PARAM_NAMES = frozenset({"self", "body", "headers", "request"})


@dataclass
class Spec:
    paths: dict[str, Any]
    schemas: dict[str, Any]
    endpoint_names: dict[str, Any]


def load_spec() -> Spec:
    return Spec(
        paths=_read("paths.json"),
        schemas=_read("schemas.json"),
        endpoint_names=_read("endpoint-names.json"),
    )


def _read(name: str) -> dict[str, Any]:
    path = SPEC_DIR / name
    if not path.exists():
        raise SystemExit(f"{path} not found. Run `pnpm run generate` at the repo root.")
    data: dict[str, Any] = json.loads(path.read_text())
    return data


Operation = tuple[str, str, dict[str, Any], dict[str, Any]]


def iter_operations(spec: Spec) -> Iterator[Operation]:
    """Yield `(path, method, operation, mapping)` for every named endpoint.

    Sorted by path so both generators walk the specification in the same order
    and neither depends on the key order `paths.json` happens to have.
    """
    for path, path_item in sorted(spec.paths.items()):
        path_endpoints = spec.endpoint_names.get(path)
        if not path_endpoints:
            continue
        for method in HTTP_METHODS:
            operation = path_item.get(method)
            mapping = path_endpoints.get(method)
            if not operation or not mapping:
                continue
            yield path, method, operation, mapping


#: How datamodel-codegen spells a declaration. Object schemas render as classes,
#: union and enum schemas as `Name: TypeAlias = ...`, and shapes whose keys are
#: not identifiers as functional TypedDicts. Matching only classes would
#: silently degrade every `$ref` to one of the others to `dict[str, Any]`.
_DECLARATIONS = (
    r"^class (\w+)\(",
    r"^(\w+): TypeAlias =",
    r"^(\w+) = TypedDict\(",
)


def declared_names(source: str) -> list[str]:
    """Every type name `source` declares, in the order the patterns find them."""
    return [
        name
        for pattern in _DECLARATIONS
        for name in re.findall(pattern, source, flags=re.MULTILINE)
    ]


def ruff_format(*paths: Path) -> None:
    """Drop unused imports and normalise formatting in generated files.

    Applied by both generators, so running either on its own leaves the tree
    byte-identical to a full regeneration, and the generated output passes the
    same checks as hand-written code.
    """
    targets = [str(path) for path in paths]
    for argv in (["check", "--fix-only", "--quiet"], ["format", "--quiet"]):
        subprocess.run([sys.executable, "-m", "ruff", *argv, *targets], check=True)


def json_response_schema(
    responses: dict[str, Any] | None, codes: tuple[str, ...] = SUCCESS_CODES
) -> dict[str, Any] | None:
    for code in codes:
        response = (responses or {}).get(code)
        content = (response or {}).get("content") or {}
        schema = (content.get("application/json") or {}).get("schema")
        if schema:
            return schema  # type: ignore[no-any-return]
    return None


def pascal_case(name: str) -> str:
    cleaned = re.sub(r"[-_\s]+(.)?", lambda m: (m.group(1) or "").upper(), name)
    return cleaned[:1].upper() + cleaned[1:]


def snake_case(name: str) -> str:
    name = re.sub(r"[-\s]+", "_", name)
    name = re.sub(r"(?<=[a-z0-9])(?=[A-Z])", "_", name)
    return name.lower()


def safe_identifier(name: str) -> str:
    """Make `name` usable as a Python parameter name."""
    identifier = snake_case(name)
    if keyword.iskeyword(identifier) or identifier in RESERVED_PARAM_NAMES:
        identifier = f"{identifier}_"
    if not identifier.isidentifier():
        # snake_case only rewrites separators and case, so anything else the
        # spec puts in a parameter name (a dot, a leading digit) would emit a
        # module that does not compile. Fail the build instead.
        raise SystemExit(f"Parameter name {name!r} is not a Python identifier")
    return identifier
