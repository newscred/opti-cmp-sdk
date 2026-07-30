# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING

from ..._types import Dispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects


class UserNamespace:
    """`user` endpoints."""

    def __init__(self, client: Dispatcher) -> None:
        self._client = client

    def find_user_by_email(
        self,
        email: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Find a user by email address

        GET /users
        """
        return self._client._dispatch(
            namespace="user",
            operation="find_user_by_email",
            params={
                "email": email,
                "headers": headers,
                "request": request,
            },
        )

    def get_user(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.UserResponse]:
        """Get a user

        GET /users/{id}
        """
        return self._client._dispatch(
            namespace="user",
            operation="get_user",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def list_users(
        self,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.UserListResponse]:
        """Get list of users.

        GET /userlist
        """
        return self._client._dispatch(
            namespace="user",
            operation="list_users",
            params={
                "headers": headers,
                "request": request,
            },
        )
