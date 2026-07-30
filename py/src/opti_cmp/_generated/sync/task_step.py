# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING

from ..._types import Dispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects, schema


class TaskStepNamespace:
    """`taskStep` endpoints."""

    def __init__(self, client: Dispatcher) -> None:
        self._client = client

    def update_task_step(
        self,
        task_id: str,
        step_id: str,
        *,
        body: schema.TaskStepRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskStep]:
        """Update a step in a task

        PATCH /tasks/{task_id}/steps/{step_id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="task_step",
            operation="update_task_step",
            params={
                "task_id": task_id,
                "step_id": step_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )
