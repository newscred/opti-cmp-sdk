# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING, Literal

from ..._types import Dispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects, schema


class PublishingNamespace:
    """`publishing` endpoints."""

    def __init__(self, client: Dispatcher) -> None:
        self._client = client

    def bulk_create_publishing_event_metadata(
        self,
        publishing_event_id: str,
        *,
        body: schema.PublishingEventMetadataBulkCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.PublishingEventMetadataBulkCreateResponse]:
        """Create asset metadata in bulk.

        POST /v3/publishing-events/{publishing_event_id}/publishing-metadata

        Args:
            publishing_event_id: Unique identifier of the publishing event
        """
        return self._client._dispatch(
            namespace="publishing",
            operation="bulk_create_publishing_event_metadata",
            params={
                "publishing_event_id": publishing_event_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def get_publishing_event(
        self,
        publishing_event_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.PublishingEventResponse]:
        """Get the publishing event by ID

        GET /v3/publishing-events/{publishing_event_id}

        Args:
            publishing_event_id: Unique identifier of the publishing event
        """
        return self._client._dispatch(
            namespace="publishing",
            operation="get_publishing_event",
            params={
                "publishing_event_id": publishing_event_id,
                "headers": headers,
                "request": request,
            },
        )

    def get_publishing_event_asset_metadata(
        self,
        publishing_event_id: str,
        asset_id: str,
        publishing_metadata_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.PublishingEventMetadataResponse]:
        """Get publishing metadata.

        GET v3/publishing-events/{publishing_event_id}/assets/{asset_id}/publishing-metadata/{publishing_metadata_id}

        Args:
            publishing_event_id: Unique identifier of the publishing event
            asset_id: Unique identifier of the asset
            publishing_metadata_id: Unique identifier of the publishing metadata
        """
        return self._client._dispatch(
            namespace="publishing",
            operation="get_publishing_event_asset_metadata",
            params={
                "publishing_event_id": publishing_event_id,
                "asset_id": asset_id,
                "publishing_metadata_id": publishing_metadata_id,
                "headers": headers,
                "request": request,
            },
        )

    def list_publishing_channels(
        self,
        *,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.PublishingChannelListResponse]:
        """Get the list of publishing channels available in the organization. Use this to populate a channel picker before creating a publishing intent for a task.

        GET /publishing-channels

        Experimental: this endpoint is experimental upstream.

        Args:
            page_size: Number of results to return per page
        """
        return self._client._dispatch(
            namespace="publishing",
            operation="list_publishing_channels",
            params={
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    def list_publishing_event_metadata(
        self,
        publishing_event_id: str,
        *,
        status: Literal["published", "unpublished", "synced", "failed"] | None = None,
        asset_type: Literal[
            "article", "image", "video", "raw_file", "structured_content"
        ]
        | None = None,
        asset_id: str | None = None,
        locale: str | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.PublishingEventMetadataListResponse]:
        """Get the list of publishing metadata for each asset of a publishing event.

        GET /v3/publishing-events/{publishing_event_id}/publishing-metadata

        Args:
            publishing_event_id: Unique identifier of the publishing event
            status: Publishing status of the asset
            asset_type: Type of asset
            asset_id: Unique identifier of the asset.
            locale: The locale to which the asset is being published to.
        """
        return self._client._dispatch(
            namespace="publishing",
            operation="list_publishing_event_metadata",
            params={
                "publishing_event_id": publishing_event_id,
                "status": status,
                "asset_type": asset_type,
                "asset_id": asset_id,
                "locale": locale,
                "headers": headers,
                "request": request,
            },
        )
