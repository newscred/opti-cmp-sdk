"""The container every JSON response body is parsed into.

A `dict` subclass that also resolves keys as attributes, so both forms work::

    response.data.title
    response.data["title"]

Subclassing `dict` keeps responses usable anywhere a mapping is expected —
`json.dumps`, `**` unpacking, `isinstance(x, Mapping)`.

Response bodies are built by passing this class as `json.loads(object_hook=...)`,
so every nested object in a payload arrives as an `APIObject` already.

The generated classes in `_generated/objects.py` declare attributes for the type
checker but are never instantiated: every response is an `APIObject`, holding
whatever the API actually returned. Fields the API adds later still arrive and
are still reachable, they are simply untyped until the specification catches up.
"""

from __future__ import annotations

from typing import TYPE_CHECKING, Any


class AttrDict(dict[str, Any]):
    """A dict whose keys are also reachable as attributes.

    The attribute proxying is hidden from type checkers on purpose. A visible
    `__getattr__` makes every attribute valid, so `campaign.titel` would
    type-check as `Any` and the typo would survive to runtime. Hidden, the
    checker sees only the fields a subclass declares, and flags anything else.

    Subclasses that want a different miss policy override `__getattr__` and
    call back here for the default.
    """

    __slots__ = ()

    def __dir__(self) -> list[str]:
        # Makes tab completion work in a REPL and a debugger.
        return [*super().__dir__(), *self]

    if not TYPE_CHECKING:

        def __getattr__(self, name):
            try:
                return self[name]
            except KeyError:
                raise AttributeError(
                    f"{type(self).__name__} has no field {name!r}"
                ) from None

        def __setattr__(self, name, value):
            self[name] = value

        def __delattr__(self, name):
            try:
                del self[name]
            except KeyError:
                raise AttributeError(name) from None


class APIObject(AttrDict):
    """A JSON object from the API, readable by attribute or by key.

    A missing attribute raises for both typos and fields the API omitted. Use
    `.get(name)` when a field is genuinely optional.

    The cost of hiding the proxy from type checkers is that a field the
    specification does not describe yet — one the API started returning since
    the last regeneration — reads fine at runtime but needs `.get("field")` to
    satisfy the checker.
    """

    __slots__ = ()
