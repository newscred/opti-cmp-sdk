# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING, Literal

from ..._types import AsyncDispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects, schema


class WorkRequestNamespace:
    """`workRequest` endpoints."""

    def __init__(self, client: AsyncDispatcher) -> None:
        self._client = client

    async def add_attachment_to_work_request(
        self,
        id: str,
        *,
        body: schema.AttachmentRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.AttachmentResponse]:
        """Create attachments for a work request.

        POST /work-requests/{id}/attachments

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="add_attachment_to_work_request",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def add_comment_to_work_request(
        self,
        id: str,
        *,
        body: schema.CommentWithReplyCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.WorkRequestCommentResponse]:
        """Post a comment on a work request.

        POST /work-requests/{id}/comments

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="add_comment_to_work_request",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def create_campaign_from_work_request(
        self,
        id: str,
        *,
        body: schema.WorkRequestCampaignRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.WorkRequestCampaignResponse]:
        """Create a new campaign from a work request.

        POST /work-requests/{id}/campaigns

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="create_campaign_from_work_request",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def create_task_from_work_request(
        self,
        id: str,
        *,
        body: schema.WorkRequestTaskRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.WorkRequestTaskResponse]:
        """Create a new task from a work request.

        POST /work-requests/{id}/tasks

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="create_task_from_work_request",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def create_work_request(
        self,
        *,
        body: schema.WorkRequestCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.WorkRequestResponse]:
        """Create a work request.

        POST /work-requests

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="create_work_request",
            params={
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def create_work_request_creative_asset(
        self,
        id: str,
        *,
        body: schema.CreativeAssetRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.CreativeAssetResponse]:
        """Create creative assets for a work request.

        POST /work-requests/{id}/creative-assets

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="create_work_request_creative_asset",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def delete_work_request_attachment(
        self,
        work_request_id: str,
        attachment_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Delete a work request attachment.

        DELETE /work-requests/{work_request_id}/attachments/{attachment_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="delete_work_request_attachment",
            params={
                "work_request_id": work_request_id,
                "attachment_id": attachment_id,
                "headers": headers,
                "request": request,
            },
        )

    async def delete_work_request_creative_asset(
        self,
        work_request_id: str,
        creative_asset_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Delete a work request creative asset.

        DELETE /work-requests/{work_request_id}/creative-assets/{creative_asset_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="delete_work_request_creative_asset",
            params={
                "work_request_id": work_request_id,
                "creative_asset_id": creative_asset_id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_work_request(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.WorkRequestResponse]:
        """Get a work request by ID.

        GET /work-requests/{id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="get_work_request",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_work_request_comment(
        self,
        work_request_id: str,
        comment_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.WorkRequestCommentResponse]:
        """Get a work request comment by ID.

        GET /work-requests/{work_request_id}/comments/{comment_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="get_work_request_comment",
            params={
                "work_request_id": work_request_id,
                "comment_id": comment_id,
                "headers": headers,
                "request": request,
            },
        )

    async def list_work_request_approved_assets(
        self,
        id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListWorkRequestApprovedAssetsResponse]:
        """Get a list of approved assets of a work request.

        GET /work-requests/{id}/approved-assets

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="list_work_request_approved_assets",
            params={
                "id": id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def list_work_request_comments(
        self,
        id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListWorkRequestCommentsResponse]:
        """Get list of comments for a work request.

        GET /work-requests/{id}/comments

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="list_work_request_comments",
            params={
                "id": id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def list_work_request_related_resources(
        self,
        id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListWorkRequestRelatedResourcesResponse]:
        """Get a list of work request related resources.

        GET /work-requests/{id}/related-resources

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="list_work_request_related_resources",
            params={
                "id": id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def list_work_requests(
        self,
        *,
        created_by: str | None = None,
        status: Literal["Submitted", "Accepted", "Completed", "Declined"] | None = None,
        order_by: Literal["priority", "status", "created_at", "due_date"] | None = None,
        order_as: Literal["asc", "desc"] | None = None,
        template_ids: list[str] | None = None,
        created_at__from: str | None = None,
        created_at__to: str | None = None,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListWorkRequestsResponse]:
        """Get a list of work requests.

        GET /work-requests

        Experimental: this endpoint is experimental upstream.

        Args:
            template_ids: Unique identifiers of the templates. You can append multiple times, for example `template_ids=template1&template_ids=template2`.
            created_at__from: Date and time as the lower limit to filter work requests by `created_at`, in ISO 8601 UTC format
            created_at__to: Date and time as the upper limit to filter work requests by `created_at`, in ISO 8601 UTC format
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="list_work_requests",
            params={
                "created_by": created_by,
                "status": status,
                "order_by": order_by,
                "order_as": order_as,
                "template_ids": template_ids,
                "created_at__from": created_at__from,
                "created_at__to": created_at__to,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def update_work_request(
        self,
        id: str,
        *,
        body: schema.WorkRequestUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Updates a work request.

        PATCH /work-requests/{id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="update_work_request",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def update_work_request_form_field(
        self,
        work_request_id: str,
        form_field_identifier: str,
        *,
        body: schema.WorkRequestFormFieldUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.WorkRequestRequestFormFieldUpdateResponse]:
        """Updates a work request form field.

        PUT /work-requests/{work_request_id}/form-fields/{form_field_identifier}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="work_request",
            operation="update_work_request_form_field",
            params={
                "work_request_id": work_request_id,
                "form_field_identifier": form_field_identifier,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )
