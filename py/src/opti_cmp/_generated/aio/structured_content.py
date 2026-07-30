# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING, Any

from ..._types import AsyncDispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects, schema


class StructuredContentNamespace:
    """`structuredContent` endpoints."""

    def __init__(self, client: AsyncDispatcher) -> None:
        self._client = client

    async def acknowledge_sccontent_preview(
        self,
        content_id: str,
        version_id: str,
        preview_id: str,
        *,
        body: schema.SCContentPreviewAcknowledgeRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Acknowledge content preview. Content preview can be acknowledged only once. So make sure you are acknowledging only the content previews targeted for your integration. Otherwise it will stall acknowledgment from other integrations.

        POST /structured-content/contents/{content_id}/versions/{version_id}/previews/{preview_id}/acknowledge

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="acknowledge_sccontent_preview",
            params={
                "content_id": content_id,
                "version_id": version_id,
                "preview_id": preview_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def complete_sccontent_preview(
        self,
        content_id: str,
        version_id: str,
        preview_id: str,
        *,
        body: schema.SCContentPreviewCompleteRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Complete content preview.

        POST /structured-content/contents/{content_id}/versions/{version_id}/previews/{preview_id}/complete

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="complete_sccontent_preview",
            params={
                "content_id": content_id,
                "version_id": version_id,
                "preview_id": preview_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def create_sccontent_type(
        self,
        *,
        body: schema.SCContentTypeCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.SCContentTypeCreateResponse]:
        """Create content type.

        POST /structured-content/content-types

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="create_sccontent_type",
            params={
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def create_sccontent_type_managed_migration(
        self,
        content_type_id: str,
        *,
        body: schema.SCContentTypeManagedMigrationCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.CreateSCContentTypeManagedMigrationResponse]:
        """Create a new managed migration job.

        POST /structured-content/content-types/{content_type_id}/managed-migrations

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="create_sccontent_type_managed_migration",
            params={
                "content_type_id": content_type_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def create_sccontent_type_version(
        self,
        content_type_id: str,
        *,
        body: schema.SCContentTypeVersionCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.SCContentTypeCreateResponse]:
        """Add Content Type Version.

        POST /structured-content/content-types/{content_type_id}/versions

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="create_sccontent_type_version",
            params={
                "content_type_id": content_type_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def delete_sccontent_type_managed_migration(
        self,
        content_type_id: str,
        job_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Delete a managed migration job with not_started status.

        DELETE /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="delete_sccontent_type_managed_migration",
            params={
                "content_type_id": content_type_id,
                "job_id": job_id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_sccontent_type(
        self,
        content_type_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.SCContentType]:
        """Get content type.

        GET /structured-content/content-types/{content_type_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="get_sccontent_type",
            params={
                "content_type_id": content_type_id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_sccontent_type_managed_migration(
        self,
        content_type_id: str,
        job_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.SCContentTypeManagedMigrationResponse]:
        """Get details of a specific managed migration job.

        GET /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="get_sccontent_type_managed_migration",
            params={
                "content_type_id": content_type_id,
                "job_id": job_id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_sccontent_type_version(
        self,
        content_type_id: str,
        version_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.SCContentTypeVersion]:
        """Get content type version.

        GET /structured-content/content-types/{content_type_id}/versions/{version_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="get_sccontent_type_version",
            params={
                "content_type_id": content_type_id,
                "version_id": version_id,
                "headers": headers,
                "request": request,
            },
        )

    async def list_sccontent_type_managed_migrations(
        self,
        content_type_id: str,
        *,
        content_migration_summary: bool | None = None,
        offset: int | None = None,
        limit: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[list[objects.SCContentTypeManagedMigrationResponse]]:
        """Retrieves a list of managed migration jobs for a specific content type.

        GET /structured-content/content-types/{content_type_id}/managed-migrations

        Experimental: this endpoint is experimental upstream.

        Args:
            content_migration_summary: Whether include a summary of content migration status (total, not started, succeeded, errored).
            offset: Pagination offset (number of jobs to skip).
            limit: Pagination limit (number of jobs to return).
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="list_sccontent_type_managed_migrations",
            params={
                "content_type_id": content_type_id,
                "content_migration_summary": content_migration_summary,
                "offset": offset,
                "limit": limit,
                "headers": headers,
                "request": request,
            },
        )

    async def list_sccontent_type_versions(
        self,
        content_type_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[list[objects.BaseContentTypeVersionModel]]:
        """Get content type versions.

        GET /structured-content/content-types/{content_type_id}/versions

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="list_sccontent_type_versions",
            params={
                "content_type_id": content_type_id,
                "headers": headers,
                "request": request,
            },
        )

    async def list_sccontent_types(
        self,
        *,
        source: str | None = None,
        disabled: bool | None = None,
        list: schema.ContentTypeListingOption | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[list[objects.BaseContentTypeModel]]:
        """Get content types.

        GET /structured-content/content-types

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="list_sccontent_types",
            params={
                "source": source,
                "disabled": disabled,
                "list": list,
                "headers": headers,
                "request": request,
            },
        )

    async def migrate_sccontent(
        self,
        content_id: str,
        *,
        body: schema.SCContentMigrationCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.MigrateSCContentResponse]:
        """Migrate content to a specific content type version. Experimental

        POST /structured-content/contents/{content_id}/migration
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="migrate_sccontent",
            params={
                "content_id": content_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def start_sccontent_type_managed_migration(
        self,
        content_type_id: str,
        job_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.SCContentTypeManagedMigrationStartResponse]:
        """Start managed migration job.

        POST /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}/start

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="start_sccontent_type_managed_migration",
            params={
                "content_type_id": content_type_id,
                "job_id": job_id,
                "headers": headers,
                "request": request,
            },
        )

    async def update_sccontent_type(
        self,
        content_type_id: str,
        *,
        body: schema.SCContentTypeUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.SCContentTypeUpdateResponse]:
        """Update content type.

        POST /structured-content/content-types/{content_type_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="update_sccontent_type",
            params={
                "content_type_id": content_type_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def update_sccontent_type_managed_migration(
        self,
        content_type_id: str,
        job_id: str,
        *,
        body: dict[str, Any],
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.UpdateSCContentTypeManagedMigrationResponse]:
        """Update a managed migration job.

        PATCH /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="update_sccontent_type_managed_migration",
            params={
                "content_type_id": content_type_id,
                "job_id": job_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def validate_sccontent_type_managed_migration(
        self,
        content_type_id: str,
        *,
        body: schema.SCContentTypeManagedMigrationValidateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ValidateSCContentTypeManagedMigrationResponse]:
        """Check managed migration possible or not.

        POST /structured-content/content-types/{content_type_id}/managed-migrations/validate

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="structured_content",
            operation="validate_sccontent_type_managed_migration",
            params={
                "content_type_id": content_type_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )
