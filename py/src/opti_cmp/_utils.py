from __future__ import annotations

import re
from collections.abc import Mapping, Sequence
from typing import Any
from urllib.parse import quote

_URL_VARIABLE_RE = re.compile(r"\{([^}]+)\}")

# Every URL template in the CMP specification uses plain `{name}` expansion, so
# simple substitution covers it. Revisit if the spec ever adopts RFC 6570
# operators (`{?a,b}`, `{+path}`, ...).


def extract_url_variable_names(url: str) -> list[str]:
    # Same key `expand_url` looks up below, so the two cannot disagree about
    # what a template names.
    return [name.strip() for name in _URL_VARIABLE_RE.findall(url)]


def expand_url(url: str, params: Mapping[str, Any]) -> str:
    def replace(match: re.Match[str]) -> str:
        value = params.get(match.group(1).strip())
        # `safe=""` matches RFC 6570 simple expansion: everything outside the
        # unreserved set is percent-encoded.
        return "" if value is None else quote(str(value), safe="")

    return _URL_VARIABLE_RE.sub(replace, url)


def _copy_dicts(value: Any) -> Any:
    if isinstance(value, dict):
        return {key: _copy_dicts(item) for key, item in value.items()}
    return value


def deep_merge(target: Mapping[str, Any], source: Mapping[str, Any]) -> dict[str, Any]:
    """Recursively merge plain dicts; any other value in `source` replaces.

    Nested dicts are copied, never aliased. `options.request` is a documented
    hook-writable surface, so sharing it with the client's endpoint defaults
    would let one hook reconfigure every later request.
    """
    merged: dict[str, Any] = {key: _copy_dicts(value) for key, value in target.items()}
    for key, value in source.items():
        existing = merged.get(key)
        if isinstance(existing, dict) and isinstance(value, dict):
            merged[key] = deep_merge(existing, value)
        else:
            merged[key] = _copy_dicts(value)
    return merged


def pick(source: Mapping[str, Any], keys: Sequence[str]) -> dict[str, Any]:
    return {key: source[key] for key in keys if source.get(key) is not None}
