# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING, Any

from ..._types import Dispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects


class UploaderNamespace:
    """`uploader` endpoints."""

    def __init__(self, client: Dispatcher) -> None:
        self._client = client

    def complete_multipart_upload(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.CompleteMultipartUploadResponse]:
        """Initiate completion of a multipart upload after all parts have been uploaded

        POST /v3/multipart-uploads/{id}/complete

        Experimental: this endpoint is experimental upstream.

        Args:
            id: ID of the multipart upload to complete
        """
        return self._client._dispatch(
            namespace="uploader",
            operation="complete_multipart_upload",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def create_multipart_upload(
        self,
        *,
        body: dict[str, Any],
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.CreateMultipartUploadResponse]:
        """Create pre-signed URLs for multipart upload of large files

        POST /v3/multipart-uploads

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="uploader",
            operation="create_multipart_upload",
            params={
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def get_multipart_upload_status(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.GetMultipartUploadStatusResponse]:
        """Retrieve the current status of a multipart upload

        GET /v3/multipart-uploads/{id}/status

        Experimental: this endpoint is experimental upstream.

        Args:
            id: ID of the multipart upload to check status for
        """
        return self._client._dispatch(
            namespace="uploader",
            operation="get_multipart_upload_status",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    def get_upload_url(
        self,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.GetUploadUrlResponse]:
        """Get a presigned URL and related meta fields to upload a file. **The validity time for this URL is 60 minutes**.

        GET /upload-url
        """
        return self._client._dispatch(
            namespace="uploader",
            operation="get_upload_url",
            params={
                "headers": headers,
                "request": request,
            },
        )
