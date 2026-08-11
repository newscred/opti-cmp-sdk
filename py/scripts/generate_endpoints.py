"""Generate the typed endpoint surface from the shared specification.

Emits

  _generated/routes.py        route metadata for the dispatcher
  _generated/sync/<ns>.py     synchronous namespace classes
  _generated/aio/<ns>.py      asynchronous namespace classes

Code is assembled line by line rather than through a template engine: Python is
whitespace-sensitive and the output shape is regular enough that direct
emission is easier to keep correct.
"""

from __future__ import annotations

import html
import json
import re
from dataclasses import dataclass, field
from pathlib import Path
from typing import Any

from _spec import (
    GENERATED_DIR,
    GENERATED_MARKER,
    declared_names,
    iter_operations,
    json_response_schema,
    load_spec,
    pascal_case,
    ruff_format,
    safe_identifier,
    snake_case,
)

HEADER = f"{GENERATED_MARKER}\n"
INDENT = " " * 4


@dataclass
class Param:
    name: str  # wire name, e.g. "page_size"
    arg: str  # Python identifier
    type: str
    required: bool
    description: str | None = None


@dataclass
class Body:
    type: str
    required: bool


@dataclass
class Endpoint:
    #: The operationId, as the specification spells it.
    operation: str
    #: snake_case. Also the identifier used in ROUTES and on the wire to
    #: hooks and adapters, so there is one spelling throughout.
    method_name: str
    params: list[Param]
    return_type: str
    summary: str | None = None
    description: str | None = None
    deprecated: bool = False
    body: Body | None = None

    @property
    def positional(self) -> list[Param]:
        return [param for param in self.params if param.required]

    @property
    def keyword(self) -> list[Param]:
        return [param for param in self.params if not param.required]


@dataclass
class Namespace:
    #: The tag-derived name, camelCase, e.g. "brandCompliance".
    name: str
    endpoints: list[Endpoint] = field(default_factory=list)

    @property
    def module(self) -> str:
        return snake_case(self.name)

    @property
    def class_name(self) -> str:
        return f"{pascal_case(self.name)}Namespace"


# --------------------------------------------------------------------------
# Schema -> Python type
# --------------------------------------------------------------------------

_PRIMITIVES = {
    "boolean": "bool",
    "integer": "int",
    "number": "float",
    "string": "str",
}


def schema_to_type(
    schema: dict[str, Any] | None,
    known: set[str],
    inline_name: str | None = None,
    module: str = "schema",
) -> str:
    if not schema:
        return "Any"

    if "$ref" in schema:
        name = schema["$ref"].rsplit("/", 1)[-1]
        return f"{module}.{name}" if name in known else "dict[str, Any]"

    if "allOf" in schema:
        # An intersection has no Python spelling, and rendering one as a union
        # would type-check values that satisfy only half of it. Every allOf in
        # the CMP specification is a named base plus an inline refinement, so
        # drop the members that render as `dict[str, Any]` — they add no
        # checkable constraint — and keep the named base.
        members = [
            schema_to_type(member, known, module=module) for member in schema["allOf"]
        ]
        named = [name for name in dict.fromkeys(members) if name != "dict[str, Any]"]
        if len(named) == 1:
            return named[0]
        if not named:
            return "dict[str, Any]"
        raise SystemExit(
            f"allOf combining several named schemas has no Python type: {named}"
        )

    for key in ("oneOf", "anyOf"):
        if key in schema:
            members = [
                schema_to_type(member, known, module=module) for member in schema[key]
            ]
            unique = list(dict.fromkeys(members))
            return " | ".join(unique) if unique else "Any"

    kind = schema.get("type")

    if kind == "array":
        return f"list[{schema_to_type(schema.get('items'), known, module=module)}]"

    if kind == "object":
        if inline_name and inline_name in known:
            return f"{module}.{inline_name}"
        return "dict[str, Any]"

    if enum := schema.get("enum"):
        return literal_type(enum)

    if kind in _PRIMITIVES:
        return _PRIMITIVES[kind]

    return "Any"


def literal_type(values: list[Any]) -> str:
    rendered = ", ".join(repr(value) for value in values)
    return f"Literal[{rendered}]"


# --------------------------------------------------------------------------
# Operation -> Endpoint
# --------------------------------------------------------------------------


def build_params(
    path_parameters: list[dict[str, Any]] | None,
    operation_parameters: list[dict[str, Any]] | None,
    known: set[str],
) -> list[Param]:
    params: list[Param] = []
    seen: set[str] = set()

    for parameter in [*(path_parameters or []), *(operation_parameters or [])]:
        if parameter.get("in") == "cookie":
            continue
        name = parameter.get("name")
        if not name:
            continue
        # Dedupe on the Python argument, not the wire name: two wire names that
        # snake_case to one identifier would emit a duplicate keyword argument.
        arg = safe_identifier(name)
        if arg in seen:
            continue
        seen.add(arg)

        params.append(
            Param(
                name=name,
                arg=arg,
                # `schema_to_type` renders an enum as a Literal already.
                type=schema_to_type(parameter.get("schema") or {}, known),
                required=parameter.get("required", parameter.get("in") == "path"),
                description=parameter.get("description"),
            )
        )

    return params


def build_body(
    request_body: dict[str, Any] | None, known: set[str], endpoint_name: str
) -> Body | None:
    if not request_body:
        return None

    content = request_body.get("content") or {}
    entry = content.get("application/json") or content.get("multipart/form-data")
    schema = (entry or {}).get("schema")
    if not schema:
        return None

    inline_name = f"{pascal_case(endpoint_name)}Request"
    return Body(
        type=schema_to_type(schema, known, inline_name),
        required=request_body.get("required", False),
    )


def build_return_type(
    responses: dict[str, Any] | None, known: set[str], endpoint_name: str
) -> tuple[str, bool]:
    schema = json_response_schema(responses)
    if schema is None:
        return "Response[None]", False

    properties = schema.get("properties") or {}
    paginated = (properties.get("data") or {}).get(
        "type"
    ) == "array" and "pagination" in properties

    inline_name = f"{pascal_case(endpoint_name)}Response"
    body = schema_to_type(schema, known, inline_name, module="objects")
    return f"Response[{body}]", paginated


def add_pagination_params(params: list[Param]) -> None:
    existing = {param.name for param in params}
    for name in ("offset", "page_size"):
        if name not in existing:
            params.append(Param(name=name, arg=name, type="int", required=False))


# --------------------------------------------------------------------------
# Emission
# --------------------------------------------------------------------------


_HTML_TAG_RE = re.compile(r"<[^>]+>")


def clean_text(text: str | None, *, limit: int = 400) -> str:
    """First paragraph of an OpenAPI description, without its HTML decoration."""
    if not text:
        return ""

    paragraph = html.unescape(_HTML_TAG_RE.sub(" ", text)).split("\n\n")[0]
    collapsed = " ".join(paragraph.split())

    if len(collapsed) > limit:
        cut = collapsed.rfind(". ", 0, limit)
        collapsed = (
            collapsed[: cut + 1] if cut > 0 else collapsed[:limit].rstrip() + "..."
        )

    return collapsed.replace("\\", "\\\\").replace('"""', "'''")


def docstring_lines(endpoint: Endpoint, indent: str) -> list[str]:
    # `summary` in this specification is the HTTP line ("GET /campaigns"), and
    # `description` is the prose — so the prose leads and the HTTP line follows.
    http_line = clean_text(endpoint.summary)
    summary = clean_text(endpoint.description) or http_line

    # Upstream marks experimental endpoints with an inline HTML badge that
    # strips down to a bare "Experimental" prefix. Lift it out of the sentence.
    experimental = summary.lower().startswith("experimental ")
    if experimental:
        summary = summary[len("experimental ") :].strip()

    body: list[str] = [summary or f"Call the `{endpoint.operation}` endpoint."]

    if http_line:
        body.extend(["", http_line])

    if experimental:
        body.extend(["", "Experimental: this endpoint is experimental upstream."])

    if endpoint.deprecated:
        body.extend(["", "Deprecated: this endpoint is deprecated upstream."])

    documented = [param for param in endpoint.params if param.description]
    if documented:
        body.extend(["", "Args:"])
        for param in documented:
            body.append(f"{INDENT}{param.arg}: {clean_text(param.description)}")

    if len(body) == 1:
        return [f'{indent}"""{body[0]}"""']

    return [
        f'{indent}"""{body[0]}',
        *[f"{indent}{line}".rstrip() for line in body[1:]],
        f'{indent}"""',
    ]


def annotation_modules(namespace: Namespace) -> list[str]:
    """Which of `objects`/`schema` this namespace's annotations actually name."""
    annotations = " ".join(
        " ".join(
            [
                endpoint.return_type,
                *(param.type for param in endpoint.params),
                endpoint.body.type if endpoint.body else "",
            ]
        )
        for endpoint in namespace.endpoints
    )
    return [module for module in ("objects", "schema") if f"{module}." in annotations]


def emit_namespace(namespace: Namespace, *, is_async: bool) -> str:
    dispatcher = "AsyncDispatcher" if is_async else "Dispatcher"
    # `objects` and `schema` define ~500 classes each and are named only in
    # annotations, which `from __future__ import annotations` leaves as strings.
    # Importing them for real costs ~17ms at import time to build classes that
    # are never instantiated, so they go under `if TYPE_CHECKING` — and only
    # where referenced, because ruff cannot delete an unused import from inside
    # that block, it can only leave a bare `pass` behind.
    modules = annotation_modules(namespace)
    typing_names = "TYPE_CHECKING, Any, Literal" if modules else "Any, Literal"

    lines: list[str] = [
        HEADER.rstrip("\n"),
        "",
        "from __future__ import annotations",
        "",
        "from collections.abc import Mapping",
        f"from typing import {typing_names}",
        "",
        f"from ..._types import {dispatcher}, RequestConfig, Response",
    ]

    if modules:
        lines.extend(
            [
                "",
                "if TYPE_CHECKING:",
                f"{INDENT}from .. import {', '.join(modules)}",
            ]
        )

    lines.extend(
        [
            "",
            "",
            f"class {namespace.class_name}:",
            f'{INDENT}"""`{namespace.name}` endpoints."""',
            "",
            f"{INDENT}def __init__(self, client: {dispatcher}) -> None:",
            f"{INDENT * 2}self._client = client",
        ]
    )

    prefix = "async " if is_async else ""
    await_ = "await " if is_async else ""

    for endpoint in namespace.endpoints:
        lines.append("")
        lines.append(f"{INDENT}{prefix}def {endpoint.method_name}(")
        lines.append(f"{INDENT * 2}self,")

        for param in endpoint.positional:
            lines.append(f"{INDENT * 2}{param.arg}: {param.type},")

        lines.append(f"{INDENT * 2}*,")

        for param in endpoint.keyword:
            lines.append(f"{INDENT * 2}{param.arg}: {param.type} | None = None,")

        if endpoint.body:
            suffix = "" if endpoint.body.required else " | None = None"
            lines.append(f"{INDENT * 2}body: {endpoint.body.type}{suffix},")

        lines.append(f"{INDENT * 2}headers: Mapping[str, str] | None = None,")
        lines.append(f"{INDENT * 2}request: RequestConfig | None = None,")
        lines.append(f"{INDENT}) -> {endpoint.return_type}:")
        lines.extend(docstring_lines(endpoint, INDENT * 2))
        lines.append(f"{INDENT * 2}return {await_}self._client._dispatch(")
        lines.append(f'{INDENT * 3}namespace="{namespace.module}",')
        lines.append(f'{INDENT * 3}operation="{endpoint.method_name}",')
        lines.append(f"{INDENT * 3}params={{")

        for param in endpoint.params:
            lines.append(f'{INDENT * 4}"{param.name}": {param.arg},')
        if endpoint.body:
            lines.append(f'{INDENT * 4}"body": body,')
        lines.append(f'{INDENT * 4}"headers": headers,')
        lines.append(f'{INDENT * 4}"request": request,')

        lines.append(f"{INDENT * 3}}},")
        lines.append(f"{INDENT * 2})")

    return "\n".join(lines) + "\n"


def emit_package_init(namespaces: list[Namespace], *, is_async: bool) -> str:
    mixin = "AsyncNamespaces" if is_async else "SyncNamespaces"
    lines: list[str] = [
        HEADER.rstrip("\n"),
        "",
        "from __future__ import annotations",
        "",
        "from typing import Any",
        "",
    ]

    for namespace in namespaces:
        lines.append(f"from .{namespace.module} import {namespace.class_name}")

    lines.extend(["", "", "NAMESPACES: dict[str, type[Any]] = {"])
    for namespace in namespaces:
        lines.append(f'{INDENT}"{namespace.module}": {namespace.class_name},')
    lines.extend(["}", "", ""])

    lines.append(f"class {mixin}:")
    lines.append(
        f'{INDENT}"""Namespaces attached by `register_api_endpoints_plugin`."""'
    )
    lines.append("")
    for namespace in namespaces:
        lines.append(f"{INDENT}{namespace.module}: {namespace.class_name}")

    lines.extend(
        [
            "",
            "",
            "__all__ = [",
            f'{INDENT}"NAMESPACES",',
            f'{INDENT}"{mixin}",',
        ]
    )
    for namespace in namespaces:
        lines.append(f'{INDENT}"{namespace.class_name}",')
    lines.append("]")

    return "\n".join(lines) + "\n"


def emit_routes(routes: dict[str, dict[str, dict[str, Any]]]) -> str:
    table = json.dumps(routes, indent=4, sort_keys=True)
    return (
        f"{HEADER}\n"
        "from __future__ import annotations\n\n"
        "from .._types import Routes\n\n"
        f"ROUTES: Routes = {table}\n"
    )


# --------------------------------------------------------------------------


def known_schema_names() -> set[str]:
    source = GENERATED_DIR / "schema.py"
    if not source.exists():
        raise SystemExit(
            f"{source} not found. Run scripts/generate_schema_types.py first."
        )
    return set(declared_names(source.read_text()))


def build_namespaces(
    spec: Any, known: set[str]
) -> tuple[list[Namespace], dict[str, dict[str, dict[str, Any]]]]:
    by_name: dict[str, Namespace] = {}
    routes: dict[str, dict[str, dict[str, Any]]] = {}

    for path, method, operation, mapping in iter_operations(spec):
        name = mapping["name"]
        namespace_name = mapping["namespace"]

        params = build_params(
            spec.paths[path].get("parameters"), operation.get("parameters"), known
        )
        return_type, paginated = build_return_type(
            operation.get("responses"), known, name
        )
        if paginated:
            add_pagination_params(params)

        endpoint = Endpoint(
            operation=name,
            method_name=snake_case(name),
            params=params,
            return_type=return_type,
            summary=operation.get("summary"),
            description=operation.get("description"),
            deprecated=bool(mapping.get("deprecated") or operation.get("deprecated")),
            body=build_body(operation.get("requestBody"), known, name),
        )

        by_name.setdefault(namespace_name, Namespace(namespace_name)).endpoints.append(
            endpoint
        )

        # Keyed by the same snake_case identifiers the generated methods
        # pass to `_dispatch`, so the lookup, the client attribute and what
        # hooks see are all spelled the same way.
        namespace_routes = routes.setdefault(snake_case(namespace_name), {})
        if endpoint.method_name in namespace_routes:
            # Two operations collapsing to one method name would emit two
            # defs of the same name (last wins) and drop a route silently.
            raise SystemExit(
                f"Duplicate method {namespace_name}.{endpoint.method_name} "
                f"from operation {name!r}"
            )
        namespace_routes[endpoint.method_name] = {
            "method": method.upper(),
            "url": path,
        }

    namespaces = sorted(by_name.values(), key=lambda ns: ns.name.lower())
    for namespace in namespaces:
        namespace.endpoints.sort(key=lambda endpoint: endpoint.method_name)

    return namespaces, routes


def write(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content)


def main() -> None:
    print("Generating endpoints...")

    spec = load_spec()
    known = known_schema_names()
    namespaces, routes = build_namespaces(spec, known)

    write(GENERATED_DIR / "routes.py", emit_routes(routes))
    write(GENERATED_DIR / "__init__.py", HEADER)

    for flavor, is_async in (("sync", False), ("aio", True)):
        directory = GENERATED_DIR / flavor
        for namespace in namespaces:
            write(
                directory / f"{namespace.module}.py",
                emit_namespace(namespace, is_async=is_async),
            )
        write(
            directory / "__init__.py",
            emit_package_init(namespaces, is_async=is_async),
        )

    total = sum(len(namespace.endpoints) for namespace in namespaces)
    print(f"  Namespaces: {len(namespaces)}")
    print(f"  Total routes: {total}")

    ruff_format(GENERATED_DIR)
    print("Done!")


if __name__ == "__main__":
    main()
