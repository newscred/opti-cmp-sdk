# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING

from ..._types import AsyncDispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects


class WorkflowNamespace:
    """`workflow` endpoints."""

    def __init__(self, client: AsyncDispatcher) -> None:
        self._client = client

    async def get_workflow(
        self,
        workflow_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.WorkflowResponse]:
        """Get a workflow by ID.

        GET /workflows/{workflow_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="workflow",
            operation="get_workflow",
            params={
                "workflow_id": workflow_id,
                "headers": headers,
                "request": request,
            },
        )

    async def list_workflows(
        self,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListWorkflowsResponse]:
        """Get a list of workflows.

        GET /workflows

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="workflow",
            operation="list_workflows",
            params={
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )
