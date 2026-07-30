# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING, Any, Literal

from ..._types import Dispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects, schema


class LibraryNamespace:
    """`library` endpoints."""

    def __init__(self, client: Dispatcher) -> None:
        self._client = client

    def add_asset_permissions(
        self,
        asset_id: str,
        *,
        body: schema.AssetPermissionBulkCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Grant asset access to users or teams

        POST /assets/{asset_id}/permissions

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="add_asset_permissions",
            params={
                "asset_id": asset_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def add_folder_permissions(
        self,
        id: str,
        *,
        body: schema.FolderPermissionBulkCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Grant folder access to users or teams

        POST /folders/{id}/permissions

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="add_folder_permissions",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def create_asset(
        self,
        *,
        body: schema.LibraryAssetCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.AssetResponse]:
        """Create a new asset. Supports only `images`, `videos`, and `raw files`. See [Upload assets](https://docs.developers.optimizely.com/content-marketing-platform/docs/upload-assets) to upload an asset.

        POST /assets
        """
        return self._client._dispatch(
            namespace="library",
            operation="create_asset",
            params={
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def create_asset_lineage(
        self,
        asset_id: str,
        *,
        body: schema.AssetLineageCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.AssetLineageResponse]:
        """Add a new external asset lineage.

        POST /assets/{asset_id}/lineages

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="create_asset_lineage",
            params={
                "asset_id": asset_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def create_asset_version(
        self,
        asset_id: str,
        *,
        body: schema.AssetVersionCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.LibraryAssetVersionResponse]:
        """Add a new version to a library asset. Supports adding versions to only `images`, `videos`, and `raw files` type assets. See [Upload assets](https://docs.developers.optimizely.com/content-marketing-platform/docs/upload-assets) to upload a version to a library asset.

        POST /assets/{asset_id}/versions
        """
        return self._client._dispatch(
            namespace="library",
            operation="create_asset_version",
            params={
                "asset_id": asset_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def create_file_urls(
        self,
        *,
        body: schema.FileUrlBulkCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.BatchFileUrlResponse]:
        """Generates download URLs of files given file guid.

        POST /file-urls
        """
        return self._client._dispatch(
            namespace="library",
            operation="create_file_urls",
            params={
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def create_folder(
        self,
        *,
        body: schema.FolderCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.FolderResponse]:
        """Create a folder. Use the `parent_folder_id` field to create a nested folder.

        POST /folders
        """
        return self._client._dispatch(
            namespace="library",
            operation="create_folder",
            params={
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def create_structured_content(
        self,
        *,
        body: schema.LibraryStructuredContentCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.LibraryStructuredContent]:
        """Creates a structured content.

        POST /structured-contents

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="create_structured_content",
            params={
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def delete_asset_lineage(
        self,
        asset_id: str,
        lineage_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Delete an asset lineage.

        DELETE /assets/{asset_id}/lineages/{lineage_id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="delete_asset_lineage",
            params={
                "asset_id": asset_id,
                "lineage_id": lineage_id,
                "headers": headers,
                "request": request,
            },
        )

    def delete_folder(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Delete a folder.

        DELETE /folders/{id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="delete_folder",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def delete_image(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Delete an image

        DELETE /images/{id}
        """
        return self._client._dispatch(
            namespace="library",
            operation="delete_image",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def delete_raw_file(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Delete a raw file

        DELETE /raw-files/{id}
        """
        return self._client._dispatch(
            namespace="library",
            operation="delete_raw_file",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def delete_video(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Delete a video

        DELETE /videos/{id}
        """
        return self._client._dispatch(
            namespace="library",
            operation="delete_video",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def get_article(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.LibraryArticle]:
        """Get an article

        GET /articles/{id}
        """
        return self._client._dispatch(
            namespace="library",
            operation="get_article",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def get_folder(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.FolderResponse]:
        """Get a folder

        GET /folders/{id}
        """
        return self._client._dispatch(
            namespace="library",
            operation="get_folder",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def get_image(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.LibraryImage]:
        """Get an image

        GET /images/{id}
        """
        return self._client._dispatch(
            namespace="library",
            operation="get_image",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def get_raw_file(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.LibraryRawFile]:
        """Get a raw file

        GET /raw-files/{id}
        """
        return self._client._dispatch(
            namespace="library",
            operation="get_raw_file",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def get_rendition(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.DetailedAssetRenditionResponse]:
        """Get a rendition given its `id`.

        GET /renditions/{id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="get_rendition",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def get_rendition_config(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.RenditionConfigResponse]:
        """Get a rendition configuration of an organization given its `id`.

        GET /rendition-configs/{id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="get_rendition_config",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def get_structured_content(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.LibraryStructuredContent]:
        """Get the structured content.

        GET /structured-contents/{id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="get_structured_content",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def get_video(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.LibraryVideo]:
        """Get a video

        GET /videos/{id}
        """
        return self._client._dispatch(
            namespace="library",
            operation="get_video",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def list_asset_fields(
        self,
        asset_id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListAssetFieldsResponse]:
        """Get the list of fields of an asset.

        GET assets/{asset_id}/fields

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="list_asset_fields",
            params={
                "asset_id": asset_id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    def list_asset_lineages(
        self,
        *,
        asset_id: str | None = None,
        used_in: Literal["external"] | None = None,
        created_at__from: str | None = None,
        created_at__to: str | None = None,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListAssetLineagesResponse]:
        """Get the list of asset lineage.

        GET /asset-lineages

        Experimental: this endpoint is experimental upstream.

        Args:
            created_at__from: Date and time as the lower limit to filter asset lineages by `created_at`, in ISO 8601 UTC format
            created_at__to: Date and time as the upper limit to filter asset lineages by `created_at`, in ISO 8601 UTC format
        """
        return self._client._dispatch(
            namespace="library",
            operation="list_asset_lineages",
            params={
                "asset_id": asset_id,
                "used_in": used_in,
                "created_at__from": created_at__from,
                "created_at__to": created_at__to,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    def list_asset_permissions(
        self,
        asset_id: str,
        *,
        access: Literal["view", "edit", "comment", "delete"] | None = None,
        max_access: Literal["view", "edit", "comment", "delete"] | None = None,
        min_access: Literal["view", "edit", "comment", "delete"] | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.AssetPermissionListResponseItem]:
        """List of entities that have permission to access the asset

        GET /assets/{asset_id}/permissions

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="list_asset_permissions",
            params={
                "asset_id": asset_id,
                "access": access,
                "max_access": max_access,
                "min_access": min_access,
                "headers": headers,
                "request": request,
            },
        )

    def list_asset_renditions(
        self,
        asset_id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListAssetRenditionsResponse]:
        """Get the renditions of an asset given its `id`.

        GET /assets/{asset_id}/renditions

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="list_asset_renditions",
            params={
                "asset_id": asset_id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    def list_assets(
        self,
        *,
        type: list[
            Literal["article", "image", "video", "raw_file", "structured_content"]
        ]
        | None = None,
        label: list[dict[str, Any]] | None = None,
        fields: list[dict[str, Any]] | None = None,
        created_at__from: str | None = None,
        created_at__to: str | None = None,
        modified_at__from: str | None = None,
        modified_at__to: str | None = None,
        folder_id: str | None = None,
        include_subfolder_assets: bool | None = None,
        search_text: str | None = None,
        campaign_id: str | None = None,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListAssetsResponse]:
        """Get the list of assets. Assets are sorted by `modified_at` in descending order. To get assets that are not inside any folder, pass `include_subfolder_assets=false` in the query param.

        GET /assets

        Args:
            type: Asset type to filter by. Example: `type=image&type=video`
            label: Label to filter by. **Must be passed as urlencoded string**. Labels that do not exist are ignored. If none of the provided labels exist, filtering is not applied. Example – `label=%7B%22group%22%3A%22ee63e3ee43925bb5cc8cd17b817d93ee%22%2C%22values%22%3A%5B%226706efc7828cd6aaedbc0434139cd3e1%22%2C%221f32651216cf2aefcaa08be1ea7dedf1%22%5D%7D`
            fields: List of fields to filter by. **Must be passed as base64 encoded string**. Fields that do not exist are ignored. If none of the provided fields exist, filtering is not applied. Example – `fields=Wwp7CiJpZCI6ICI2N2E4NDZhMWM3NzU1YTFwNThpNjh5MzVhIiwKInZhbHVlcyI6IFsiNjdhODQ2YTFjNzc1YWU1YWExYTE0YTA1Il0KfQpd=`
            created_at__from: Date and time as the lower limit to filter assets by `created_at`, in ISO 8601 UTC format
            created_at__to: Date and time as the upper limit to filter assets by `created_at`, in ISO 8601 UTC format
            modified_at__from: Date and time as the lower limit to filter assets by `modified_at`, in ISO 8601 UTC format
            modified_at__to: Date and time as the upper limit to filter assets by `modified_at`, in ISO 8601 UTC format
            folder_id: ID of the library folder to include assets from
            include_subfolder_assets: Indicates whether assets from subfolders need to be included
            search_text: Search assets by title or content description
            campaign_id: ID of the campaign to include assets from
        """
        return self._client._dispatch(
            namespace="library",
            operation="list_assets",
            params={
                "type": type,
                "label": label,
                "fields": fields,
                "created_at__from": created_at__from,
                "created_at__to": created_at__to,
                "modified_at__from": modified_at__from,
                "modified_at__to": modified_at__to,
                "folder_id": folder_id,
                "include_subfolder_assets": include_subfolder_assets,
                "search_text": search_text,
                "campaign_id": campaign_id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    def list_folder_permissions(
        self,
        id: str,
        *,
        access: Literal["view", "edit", "comment", "delete"] | None = None,
        max_access: Literal["view", "edit", "comment", "delete"] | None = None,
        min_access: Literal["view", "edit", "comment", "delete"] | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.FolderPermissionListResponseItem]:
        """List of entities that have permission to access the folder

        GET /folders/{id}/permissions

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="list_folder_permissions",
            params={
                "id": id,
                "access": access,
                "max_access": max_access,
                "min_access": min_access,
                "headers": headers,
                "request": request,
            },
        )

    def list_folders(
        self,
        *,
        parent_folder_id: str | None = None,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListFoldersResponse]:
        """Get the list of folders sorted by `modified_at` in descending order.

        GET /folders

        Args:
            parent_folder_id: ID of the parent folder to filter by
        """
        return self._client._dispatch(
            namespace="library",
            operation="list_folders",
            params={
                "parent_folder_id": parent_folder_id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    def list_related_assets(
        self,
        asset_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.RelatedAssetsListResponse]:
        """Returns a paginated list of related assets for the specified asset. Only supported for asset types: article, image, video, raw_file, and structured_content.

        GET /assets/{asset_id}/related-assets
        """
        return self._client._dispatch(
            namespace="library",
            operation="list_related_assets",
            params={
                "asset_id": asset_id,
                "headers": headers,
                "request": request,
            },
        )

    def remove_asset_permission(
        self,
        asset_id: str,
        accessor_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Remove accessor's access from an asset

        DELETE /asstes/{asset_id}/permissions/{accessor_id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="remove_asset_permission",
            params={
                "asset_id": asset_id,
                "accessor_id": accessor_id,
                "headers": headers,
                "request": request,
            },
        )

    def remove_folder_permission(
        self,
        id: str,
        accessor_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Remove accessor's access from a folder

        DELETE /folders/{id}/permissions/{accessor_id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="remove_folder_permission",
            params={
                "id": id,
                "accessor_id": accessor_id,
                "headers": headers,
                "request": request,
            },
        )

    def replace_related_assets(
        self,
        asset_id: str,
        *,
        body: schema.ReplaceRelatedAssetsRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ReplaceRelatedAssetsResponse]:
        """Replaces all related assets for the specified asset. Only supported for asset types: article, image, video, raw_file, and structured_content.

        PUT /assets/{asset_id}/related-assets
        """
        return self._client._dispatch(
            namespace="library",
            operation="replace_related_assets",
            params={
                "asset_id": asset_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def update_asset_field(
        self,
        asset_id: str,
        field_id: str,
        *,
        body: schema.AssetFieldUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.AssetFieldUpdateResponse]:
        """Update the field value of an asset.

        PUT /assets/{asset_id}/fields/{field_id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="update_asset_field",
            params={
                "asset_id": asset_id,
                "field_id": field_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def update_asset_fields(
        self,
        asset_id: str,
        *,
        body: schema.AssetFieldsUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.UpdateAssetFieldsResponse]:
        """Replace existing fields of an asset.

        PUT /assets/{asset_id}/fields

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="update_asset_fields",
            params={
                "asset_id": asset_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def update_asset_permission(
        self,
        asset_id: str,
        accessor_id: str,
        *,
        body: schema.AssetPermissionUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Update accessor's access level and ownership of asset

        PATCH /assets/{asset_id}/permissions/{accessor_id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="update_asset_permission",
            params={
                "asset_id": asset_id,
                "accessor_id": accessor_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def update_folder(
        self,
        id: str,
        *,
        body: schema.FolderUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.FolderResponse]:
        """Update a folder's name and/or parent.

        PATCH /folders/{id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="update_folder",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def update_folder_permission(
        self,
        id: str,
        accessor_id: str,
        *,
        body: schema.FolderPermissionUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Update accessor's access level and ownership of folder

        PATCH /folders/{id}/permissions/{accessor_id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="update_folder_permission",
            params={
                "id": id,
                "accessor_id": accessor_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def update_image(
        self,
        id: str,
        *,
        body: schema.LibraryImageUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.LibraryImage]:
        """Updates an image

        PATCH /images/{id}
        """
        return self._client._dispatch(
            namespace="library",
            operation="update_image",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def update_raw_file(
        self,
        id: str,
        *,
        body: schema.LibraryRawFileUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.LibraryRawFile]:
        """Updates a raw file

        PATCH /raw-files/{id}
        """
        return self._client._dispatch(
            namespace="library",
            operation="update_raw_file",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def update_structured_content(
        self,
        id: str,
        *,
        body: schema.LibraryStructuredContentUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.LibraryStructuredContent]:
        """Updates a structured content.

        PATCH /structured-contents/{id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="library",
            operation="update_structured_content",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def update_video(
        self,
        id: str,
        *,
        body: schema.LibraryVideoUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.LibraryVideo]:
        """Updates a video

        PATCH /videos/{id}
        """
        return self._client._dispatch(
            namespace="library",
            operation="update_video",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )
