# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING

from ..._types import Dispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects


class TeamNamespace:
    """`team` endpoints."""

    def __init__(self, client: Dispatcher) -> None:
        self._client = client

    def get_team(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TeamWithUsers]:
        """Get a team.

        GET /teams/{id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="team",
            operation="get_team",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def list_teams(
        self,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListTeamsResponse]:
        """Get a list of teams.

        GET /teams

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="team",
            operation="list_teams",
            params={
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )
