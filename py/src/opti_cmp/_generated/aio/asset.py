# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping

from ..._types import AsyncDispatcher, RequestConfig, Response


class AssetNamespace:
    """`asset` endpoints."""

    def __init__(self, client: AsyncDispatcher) -> None:
        self._client = client

    async def get_asset_url(
        self,
        asset_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Find URL of an asset by ID

        GET /asset-urls/{asset_id}
        """
        return await self._client._dispatch(
            namespace="asset",
            operation="get_asset_url",
            params={
                "asset_id": asset_id,
                "headers": headers,
                "request": request,
            },
        )
