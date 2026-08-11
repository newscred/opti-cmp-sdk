# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING

from ..._types import Dispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects, schema


class FieldNamespace:
    """`field` endpoints."""

    def __init__(self, client: Dispatcher) -> None:
        self._client = client

    def create_field(
        self,
        *,
        body: schema.SettingsFieldCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.CreateFieldResponse]:
        """Add a new field to an Organization.

        POST /fields

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="field",
            operation="create_field",
            params={
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def create_field_choices(
        self,
        id: str,
        *,
        body: schema.SettingsFieldChoiceCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.SettingsFieldChoiceCreateResponse]:
        """Create choices in a field.

        POST /fields/{id}/choices

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="field",
            operation="create_field_choices",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def delete_field_choice(
        self,
        field_id: str,
        choice_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Delete a choice from a field.

        DELETE /fields/{field_id}/choices/{choice_id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="field",
            operation="delete_field_choice",
            params={
                "field_id": field_id,
                "choice_id": choice_id,
                "headers": headers,
                "request": request,
            },
        )

    def list_fields(
        self,
        *,
        ids: str | None = None,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListFieldsResponse]:
        """Get the list of fields of an Organization.

        GET /fields

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="field",
            operation="list_fields",
            params={
                "ids": ids,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    def update_field(
        self,
        id: str,
        *,
        body: schema.SettingsFieldUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Update a field in an Organization.

        PATCH /fields/{id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="field",
            operation="update_field",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    def update_field_choice(
        self,
        field_id: str,
        choice_id: str,
        *,
        body: schema.SettingsFieldChoiceUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Update a choice of a field.

        PATCH /fields/{field_id}/choices/{choice_id}

        Experimental: this endpoint is experimental upstream.
        """
        return self._client._dispatch(
            namespace="field",
            operation="update_field_choice",
            params={
                "field_id": field_id,
                "choice_id": choice_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )
