"""Exceptions raised by the SDK.

Nothing here inherits from a builtin, so no SDK error is an `OSError`.
"""

from __future__ import annotations

from typing import TYPE_CHECKING, Any

if TYPE_CHECKING:
    from ._types import Headers, RequestOptions

__all__ = [
    "APIConnectionError",
    "APIError",
    "APITimeoutError",
    "InvalidRequestError",
    "OptiCMPError",
]


class OptiCMPError(Exception):
    """Base for every error the SDK raises."""


class _RequestScopedError(OptiCMPError):
    """An error carrying the request it came from, but no response."""

    def __init__(self, message: str, *, request: RequestOptions | None = None) -> None:
        super().__init__(message)
        self.message = message
        self.request = request

    def __repr__(self) -> str:
        return f"{type(self).__name__}({self.message!r})"


class APIError(OptiCMPError):
    """The server responded, with a status outside the 2xx range."""

    def __init__(
        self,
        message: str,
        status: int,
        *,
        data: Any = None,
        headers: Headers | None = None,
        request: RequestOptions | None = None,
    ) -> None:
        super().__init__(message)
        self.message = message
        self.status = status
        self.data = data
        self.headers = headers
        self.request = request

    def __repr__(self) -> str:
        return f"APIError({self.message!r}, status={self.status})"


class APIConnectionError(_RequestScopedError):
    """The request never reached a response, so there is no status to report."""


class APITimeoutError(APIConnectionError):
    """The request timed out before a response arrived."""


class InvalidRequestError(_RequestScopedError):
    """The request could not be built or sent as written.

    Caller input, not a network fault — a malformed base URL, or a header value
    the HTTP client refused. Nothing was put on the wire.
    """
