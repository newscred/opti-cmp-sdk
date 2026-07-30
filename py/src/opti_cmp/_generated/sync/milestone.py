# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING

from ..._types import Dispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects, schema


class MilestoneNamespace:
    """`milestone` endpoints."""

    def __init__(self, client: Dispatcher) -> None:
        self._client = client

    def create_milestone(
        self,
        *,
        body: schema.MilestoneCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.MilestoneResponse]:
        """Creates a milestone. Defaults to the organization's campaign if campaign_id is omitted. Tasks are validated against the selected campaign and must belong to it.

        POST /milestones

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="milestone",
            operation="create_milestone",
            params={
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def get_milestone(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.MilestoneResponse]:
        """Get a milestone by ID.

        GET /milestones/{id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="milestone",
            operation="get_milestone",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def list_milestones(
        self,
        *,
        campaign_id: str | None = None,
        due_date__from: str | None = None,
        due_date__to: str | None = None,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListMilestonesResponse]:
        """Get a list of milestones.

        GET /milestones

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="milestone",
            operation="list_milestones",
            params={
                "campaign_id": campaign_id,
                "due_date__from": due_date__from,
                "due_date__to": due_date__to,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    def update_milestone(
        self,
        id: str,
        *,
        body: schema.MilestoneUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.MilestoneResponse]:
        """Updates a milestone. All fields are optional. Only provided fields will be updated. Note: If tasks array is empty, it will remove all task associations for that milestone.

        PATCH /milestones/{id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="milestone",
            operation="update_milestone",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )
