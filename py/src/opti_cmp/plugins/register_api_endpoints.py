"""Attaches the generated namespaces to the client."""

from __future__ import annotations

from typing import Any


def register_api_endpoints_plugin(client: Any, _options: dict[str, Any]) -> None:
    if client._is_async:
        from .._generated.aio import NAMESPACES
    else:
        from .._generated.sync import NAMESPACES

    for attribute, namespace_class in NAMESPACES.items():
        setattr(client, attribute, namespace_class(client))
