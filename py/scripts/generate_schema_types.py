"""Generate the schema types from the shared spec.

Two modules come out of one datamodel-codegen run:

  `schema.py`   TypedDicts — used for request bodies, where the caller builds a
                dict literal and wants every key checked.
  `objects.py`  the same shapes as `APIObject` subclasses declaring attributes —
                used for responses, so `response.data.title` type-checks.

`objects.py` is derived from `schema.py` rather than generated separately, so
the two cannot drift: same names, same fields, same inheritance.

Component schemas come straight from `specification/schemas.json`. Inline
response schemas live under `paths.json`, so they are lifted into the same
namespace as `{PascalCaseEndpointName}Response` — the naming
`js/scripts/generate-schema-exports.ts` uses, so both SDKs expose the same type
names.
"""

from __future__ import annotations

import ast
import json
import subprocess
import sys
import tempfile
from pathlib import Path
from typing import Any

from _spec import (
    GENERATED_DIR,
    GENERATED_MARKER,
    Spec,
    declared_names,
    iter_operations,
    json_response_schema,
    load_spec,
    pascal_case,
    ruff_format,
)

OUTPUT = GENERATED_DIR / "schema.py"
OBJECTS_OUTPUT = GENERATED_DIR / "objects.py"

OBJECTS_HEADER = f'''{GENERATED_MARKER}
# mypy: disable-error-code="misc,assignment"
"""Response shapes, declaring attributes for the type checker.

Derived from `schema.py`. These classes are never instantiated: every response
body is an `APIObject` holding exactly what the API returned. Declaring the
fields here is what makes `response.data.title` type-check, while leaving fields
the specification does not yet know about reachable at runtime.

Optional fields are declared `T | None`. A field the API omits entirely raises
`AttributeError` rather than returning `None`; use `.get("field")` for those.

A few shapes keep the functional `TypedDict` form because their keys are not
valid Python identifiers (`x-amz-signature`). Those stay subscript-only, which
is the only thing that could work.
"""

from __future__ import annotations

from typing import Any, Literal, TypeAlias, TypedDict

from typing_extensions import NotRequired

from .._object import APIObject
'''

# Several upstream schemas compose `allOf` branches that disagree about whether
# a field is required, which lands as a TypedDict subclass relaxing a required
# field to NotRequired — something TypedDict inheritance forbids. The types are
# still the best available description of the payloads, so the check is
# suppressed for this generated module only.
HEADER = f'{GENERATED_MARKER}\n# mypy: disable-error-code="misc"\n'


#: A schema is public if it is something the API takes or returns. Everything
#: else is a nested anonymous object datamodel-code-generator had to invent a
#: name for — `Links2`, `Datum1`, `Choice12` — which is an artifact of how the
#: module was generated, not a concept in the CMP API. The same four suffixes
#: pick the JS SDK's public surface in
#: `js/scripts/generate-schema-exports.ts`, so both SDKs advertise one set.
PUBLIC_SUFFIXES = ("Request", "Response", "Payload", "ListResponseItem")


def public_names(source: str) -> list[str]:
    return sorted(
        name for name in declared_names(source) if name.endswith(PUBLIC_SUFFIXES)
    )


def emit_all(names: list[str]) -> str:
    """An `__all__` so the generated helper names stay out of the public API.

    They remain importable by name for anyone who needs one; they just are not
    advertised, do not answer to `import *`, and carry no stability promise.
    """
    entries = "".join(f'    "{name}",\n' for name in names)
    return f"\n__all__ = [\n{entries}]\n"


def collect_inline_responses(
    spec: Spec, taken: set[str]
) -> tuple[dict[str, Any], list[str]]:
    """Lift inline 2xx JSON response schemas into named component schemas."""
    inline: dict[str, Any] = {}
    skipped: list[str] = []

    for path, method, operation, mapping in iter_operations(spec):
        schema = json_response_schema(operation.get("responses"))
        if schema is None or "$ref" in schema:
            continue

        name = f"{pascal_case(mapping['name'])}Response"
        if name in taken or name in inline:
            skipped.append(f"{name} ({method.upper()} {path})")
            continue
        inline[name] = schema

    return inline, skipped


def _is_optional(annotation: ast.expr) -> bool:
    """True for `X | None`, which `--strict-nullable` already produces."""
    return (
        isinstance(annotation, ast.BinOp)
        and isinstance(annotation.op, ast.BitOr)
        and isinstance(annotation.right, ast.Constant)
        and annotation.right.value is None
    )


def _unwrap_not_required(annotation: ast.expr) -> ast.expr:
    """`NotRequired[T]` -> `T | None`; an omitted field reads as absent."""
    if not (
        isinstance(annotation, ast.Subscript)
        and isinstance(annotation.value, ast.Name)
        and annotation.value.id == "NotRequired"
    ):
        return annotation

    inner = annotation.slice
    # A field that is both optional and nullable is already `T | None`; adding
    # another would render as `T | None | None`.
    if _is_optional(inner):
        return inner
    return ast.BinOp(left=inner, op=ast.BitOr(), right=ast.Constant(value=None))


#: Everything OBJECTS_HEADER imports. The generated imports are dropped in
#: favour of it, so anything datamodel-codegen starts emitting that is not in
#: here would leave `objects.py` raising NameError on import.
OBJECTS_HEADER_NAMES = frozenset(
    {
        "annotations",
        "Any",
        "Literal",
        "NotRequired",
        "TypeAlias",
        "TypedDict",
        "APIObject",
    }
)


def derive_objects(schema_source: str) -> str:
    """Rewrite the TypedDict module as APIObject subclasses."""
    module = ast.parse(schema_source)
    body: list[ast.stmt] = []

    for node in module.body:
        if isinstance(node, ast.ImportFrom):
            # objects.py supplies its own imports; fail loudly rather than
            # emitting a module that cannot be imported.
            missing = {
                alias.asname or alias.name for alias in node.names
            } - OBJECTS_HEADER_NAMES
            if missing:
                raise SystemExit(
                    f"OBJECTS_HEADER does not import {sorted(missing)}, which "
                    f"datamodel-code-generator now emits. Add them to it."
                )
            continue
        if isinstance(node, ast.ClassDef):
            node.bases = [
                ast.Name(id="APIObject", ctx=ast.Load())
                if isinstance(base, ast.Name) and base.id == "TypedDict"
                else base
                for base in node.bases
            ]
            node.body = [
                ast.AnnAssign(
                    target=statement.target,
                    annotation=_unwrap_not_required(statement.annotation),
                    value=None,
                    simple=statement.simple,
                )
                if isinstance(statement, ast.AnnAssign)
                else statement
                for statement in node.body
            ]
        body.append(node)

    module.body = body
    return OBJECTS_HEADER + "\n\n" + ast.unparse(ast.fix_missing_locations(module))


def main() -> None:
    print("Generating schema types...")

    spec = load_spec()
    schemas = dict(spec.schemas)
    inline, skipped = collect_inline_responses(spec, set(schemas))
    schemas.update(inline)

    print(f"  Component schemas: {len(spec.schemas)}")
    print(f"  Inline response schemas: {len(inline)}")
    if skipped:
        print(f"  Skipped {len(skipped)} colliding inline names: {', '.join(skipped)}")

    document = {
        "openapi": "3.0.3",
        "info": {"title": "Optimizely CMP", "version": "0"},
        "paths": {},
        "components": {"schemas": schemas},
    }

    GENERATED_DIR.mkdir(parents=True, exist_ok=True)

    with tempfile.TemporaryDirectory() as tmp:
        source = Path(tmp) / "openapi.json"
        source.write_text(json.dumps(document))

        subprocess.run(
            [
                sys.executable,
                "-m",
                "datamodel_code_generator",
                "--input",
                str(source),
                "--input-file-type",
                "openapi",
                "--output-model-type",
                "typing.TypedDict",
                # PEP 728 `closed=True` is not supported by mypy.
                "--no-use-closed-typed-dict",
                "--target-python-version",
                "3.10",
                "--use-standard-collections",
                "--use-union-operator",
                "--disable-timestamp",
                # Without this, `nullable: true` is ignored for required
                # fields, so 120 fields that the API really does send as null
                # would be typed as never-None and blow up on attribute access.
                "--strict-nullable",
                # Same formatter the rest of the package uses; also keeps black
                # and isort out of the dependency tree.
                "--formatters",
                "ruff-format",
                "--output",
                str(OUTPUT),
            ],
            check=True,
        )

    schema_source = HEADER + OUTPUT.read_text()
    # Appended before objects.py is derived, so the `__all__` node is carried
    # across by `derive_objects` and both modules advertise the same names.
    names = public_names(schema_source)
    schema_source += emit_all(names)

    OUTPUT.write_text(schema_source)
    print(f"  Written: {OUTPUT}")

    OBJECTS_OUTPUT.write_text(derive_objects(schema_source))
    print(f"  Written: {OBJECTS_OUTPUT}")
    print(f"  Public names: {len(names)}")

    ruff_format(OUTPUT, OBJECTS_OUTPUT)
    print("Done!")


if __name__ == "__main__":
    main()
