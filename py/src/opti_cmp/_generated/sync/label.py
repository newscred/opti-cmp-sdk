# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING, Literal

from ..._types import Dispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects


class LabelNamespace:
    """`label` endpoints."""

    def __init__(self, client: Dispatcher) -> None:
        self._client = client

    def list_label_groups(
        self,
        *,
        source_org_type: Literal["current", "related"] | None = None,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListLabelGroupsResponse]:
        """Get the list of label groups. Label groups are sorted by `name` in ascending order.

        GET /label-groups

        Args:
            source_org_type: Source organization type to filter by
        """
        return self._client._dispatch(
            namespace="label",
            operation="list_label_groups",
            params={
                "source_org_type": source_org_type,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )
