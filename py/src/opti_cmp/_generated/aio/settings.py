# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING

from ..._types import AsyncDispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects, schema


class SettingsNamespace:
    """`settings` endpoints."""

    def __init__(self, client: AsyncDispatcher) -> None:
        self._client = client

    async def get_settings(
        self,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.Settings]:
        """Get settings.

        GET /settings

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="settings",
            operation="get_settings",
            params={
                "headers": headers,
                "request": request,
            },
        )

    async def update_settings(
        self,
        *,
        execute: bool | None = None,
        overwrite_workflows: bool | None = None,
        body: schema.SettingsResources,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.SettingsUpdateResponse]:
        """Create or update settings. ## Import behavior: - Label: - If a label-group with the same name exists, the label-group is updated. - Labels are merged. - If a label with the same name exists, it is overwritten. - If no label-group with the same name exists, a new label-group is created.

        POST /settings

        Experimental: this endpoint is experimental upstream.

        Args:
            execute: If `execute=true` the settings are created or updated. Otherwise, the endpoint returns only the changeset.
            overwrite_workflows: If `overwrite_workflows=true` the existing workflows are overwritten. Otherwise, a new workflow is created where a prefix `Copy of` is added to the workflow's name.
        """
        return await self._client._dispatch(
            namespace="settings",
            operation="update_settings",
            params={
                "execute": execute,
                "overwrite_workflows": overwrite_workflows,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )
