# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING, Literal

from ..._types import AsyncDispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects


class TemplateNamespace:
    """`template` endpoints."""

    def __init__(self, client: AsyncDispatcher) -> None:
        self._client = client

    async def get_template(
        self,
        template_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TemplateResponse]:
        """Get a template by ID.

        GET /templates/{template_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="template",
            operation="get_template",
            params={
                "template_id": template_id,
                "headers": headers,
                "request": request,
            },
        )

    async def list_templates(
        self,
        *,
        search: str | None = None,
        applicable_to: Literal["work_request", "task_brief", "campaign_brief"]
        | None = None,
        include_inactive: bool | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TemplateListResponse]:
        """Get a list of templates.

        GET /templates

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="template",
            operation="list_templates",
            params={
                "search": search,
                "applicable_to": applicable_to,
                "include_inactive": include_inactive,
                "headers": headers,
                "request": request,
            },
        )
