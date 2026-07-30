# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING

from ..._types import AsyncDispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects, schema


class BrandComplianceNamespace:
    """`brandCompliance` endpoints."""

    def __init__(self, client: AsyncDispatcher) -> None:
        self._client = client

    async def get_task_asset_draft_brand_compliance(
        self,
        task_id: str,
        asset_id: str,
        draft_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskAssetDraftBrandComplianceResponse]:
        """Get the brand compliance details of a draft of an asset of a task.

        GET /tasks/{task_id}/assets/{asset_id}/drafts/{draft_id}/brand-compliance

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="brand_compliance",
            operation="get_task_asset_draft_brand_compliance",
            params={
                "task_id": task_id,
                "asset_id": asset_id,
                "draft_id": draft_id,
                "headers": headers,
                "request": request,
            },
        )

    async def list_brand_compliance_categories(
        self,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListBrandComplianceCategoriesResponse]:
        """Get a list of brand compliance categories.

        GET /brand-compliance/categories

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="brand_compliance",
            operation="list_brand_compliance_categories",
            params={
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def update_task_asset_draft_brand_compliance(
        self,
        task_id: str,
        asset_id: str,
        draft_id: str,
        *,
        body: schema.TaskAssetDraftBrandComplianceRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskAssetDraftBrandComplianceResponse]:
        """Update the brand compliance details of a draft of an asset of a task.

        PUT /tasks/{task_id}/assets/{asset_id}/drafts/{draft_id}/brand-compliance

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="brand_compliance",
            operation="update_task_asset_draft_brand_compliance",
            params={
                "task_id": task_id,
                "asset_id": asset_id,
                "draft_id": draft_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )
