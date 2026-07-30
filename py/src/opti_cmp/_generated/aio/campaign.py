# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING

from ..._types import AsyncDispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects, schema


class CampaignNamespace:
    """`campaign` endpoints."""

    def __init__(self, client: AsyncDispatcher) -> None:
        self._client = client

    async def add_attachment_to_campaign(
        self,
        id: str,
        *,
        body: schema.AttachmentRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.AttachmentResponse]:
        """Add an attachment to the campaign.

        POST /campaigns/{id}/attachments

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="campaign",
            operation="add_attachment_to_campaign",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def add_comment_to_campaign(
        self,
        id: str,
        *,
        body: schema.CommentCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.CampaignCommentResponse]:
        """Post a comment on a campaign.

        POST /campaigns/{id}/comments

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="campaign",
            operation="add_comment_to_campaign",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def add_field_to_campaign(
        self,
        id: str,
        *,
        body: schema.ObjectFieldCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ObjectFieldCreateResponse]:
        """Add a field to a campaign.

        POST campaigns/{id}/fields

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="campaign",
            operation="add_field_to_campaign",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def create_campaign(
        self,
        *,
        body: schema.CampaignCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.CampaignResponse]:
        """Create a campaign

        POST /campaigns

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="campaign",
            operation="create_campaign",
            params={
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def get_campaign(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.CampaignResponse]:
        """Get a campaign

        GET /campaigns/{id}
        """
        return await self._client._dispatch(
            namespace="campaign",
            operation="get_campaign",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_campaign_brief(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.CampaignBriefResponse]:
        """Get brief of the campaign

        GET /campaigns/{id}/brief
        """
        return await self._client._dispatch(
            namespace="campaign",
            operation="get_campaign_brief",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    async def list_campaign_fields(
        self,
        id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListCampaignFieldsResponse]:
        """Get the list of fields of a campaign.

        GET campaigns/{id}/fields

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="campaign",
            operation="list_campaign_fields",
            params={
                "id": id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def list_campaigns(
        self,
        *,
        owner: str | None = None,
        start_date: str | None = None,
        end_date: str | None = None,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListCampaignsResponse]:
        """Get a list of campaigns.

        GET /campaigns

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="campaign",
            operation="list_campaigns",
            params={
                "owner": owner,
                "start_date": start_date,
                "end_date": end_date,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def update_campaign(
        self,
        id: str,
        *,
        body: schema.CampaignUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Update a campaign

        PATCH /campaigns/{id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="campaign",
            operation="update_campaign",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def update_campaign_field(
        self,
        campaign_id: str,
        field_id: str,
        *,
        body: schema.CampaignFieldUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.CampaignFieldUpdateResponse]:
        """Update the field value of a campaign.

        PUT /campaigns/{campaign_id}/fields/{field_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="campaign",
            operation="update_campaign_field",
            params={
                "campaign_id": campaign_id,
                "field_id": field_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )
