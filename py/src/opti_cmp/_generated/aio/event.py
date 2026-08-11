# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING

from ..._types import AsyncDispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects, schema


class EventNamespace:
    """`event` endpoints."""

    def __init__(self, client: AsyncDispatcher) -> None:
        self._client = client

    async def create_event(
        self,
        *,
        body: schema.EventCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.EventResponse]:
        """Create an event.

        POST /events

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="event",
            operation="create_event",
            params={
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def get_event(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.EventResponse]:
        """Get an event.

        GET /events/{id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="event",
            operation="get_event",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    async def list_event_fields(
        self,
        id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListEventFieldsResponse]:
        """Get the list of fields of an event.

        GET events/{id}/fields

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="event",
            operation="list_event_fields",
            params={
                "id": id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def list_events(
        self,
        *,
        start_date: str | None = None,
        end_date: str | None = None,
        campaign_id: str | None = None,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListEventsResponse]:
        """Get a list of events.

        GET /events

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="event",
            operation="list_events",
            params={
                "start_date": start_date,
                "end_date": end_date,
                "campaign_id": campaign_id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def update_event(
        self,
        id: str,
        *,
        body: schema.EventUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.EventResponse]:
        """Update an event.

        PATCH /events/{id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="event",
            operation="update_event",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def update_event_fields(
        self,
        id: str,
        *,
        body: schema.EventFieldsUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Replace existing fields of an event.

        PUT /events/{id}/fields

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="event",
            operation="update_event_fields",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )
