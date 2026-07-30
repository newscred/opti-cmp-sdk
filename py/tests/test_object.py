"""Payload access: `response.data.title` as well as `response.data["title"]`."""

from __future__ import annotations

import json
from collections.abc import Mapping
from typing import Any, cast

import httpx
import pytest

from opti_cmp import APIError, APIObject, AsyncOptiCMP, OptiCMP

from .conftest import json_transport

CAMPAIGN = {
    "id": "camp-1",
    "title": "Q3",
    "budget": {"budgeted_amount": "100.00", "currency_code": "USD"},
    "labels": [
        {"group": {"name": "priority"}, "values": [{"name": "urgent"}]},
        {"group": {"name": "region"}, "values": [{"name": "eu"}]},
    ],
}


class TestAPIObject:
    def test_reads_by_attribute_and_by_key(self) -> None:
        obj: Any = APIObject({"title": "Q3"})
        assert obj.title == "Q3"
        assert obj["title"] == "Q3"

    def test_missing_key_raises_attribute_error(self) -> None:
        obj: Any = APIObject({"title": "Q3"})
        with pytest.raises(AttributeError, match="titel"):
            _ = obj.titel

    def test_get_is_the_safe_form_for_optional_fields(self) -> None:
        obj = APIObject({"title": "Q3"})
        assert obj.get("description") is None
        assert obj.get("description", "none") == "none"

    def test_is_still_a_mapping(self) -> None:
        obj = APIObject({"title": "Q3"})
        assert isinstance(obj, Mapping)
        assert dict(obj) == {"title": "Q3"}
        assert json.loads(json.dumps(obj)) == {"title": "Q3"}
        assert {**obj} == {"title": "Q3"}

    def test_dir_lists_the_keys_for_repl_completion(self) -> None:
        assert "title" in dir(APIObject({"title": "Q3"}))


class TestResponsePayloads:
    def test_attribute_access_on_a_response(self) -> None:
        with OptiCMP(transport=json_transport(CAMPAIGN)) as client:
            campaign = client.campaign.get_campaign("camp-1").data

        assert campaign.title == "Q3"
        assert campaign["title"] == "Q3"

    def test_nested_objects_are_wrapped(self) -> None:
        with OptiCMP(transport=json_transport(CAMPAIGN)) as client:
            campaign = client.campaign.get_campaign("camp-1").data

        budget = campaign.budget
        assert budget is not None
        assert budget.budgeted_amount == "100.00"
        assert budget["currency_code"] == "USD"

    def test_objects_inside_lists_are_wrapped(self) -> None:
        with OptiCMP(transport=json_transport(CAMPAIGN)) as client:
            campaign = client.campaign.get_campaign("camp-1").data

        assert [label.group.name for label in campaign.labels] == [
            "priority",
            "region",
        ]

    def test_a_field_the_api_omitted_raises(self) -> None:
        with OptiCMP(transport=json_transport({"id": "camp-1"})) as client:
            campaign = client.campaign.get_campaign("camp-1").data

        with pytest.raises(AttributeError):
            _ = campaign.description
        assert campaign.get("description") is None

    def test_fields_the_spec_does_not_know_still_arrive(self) -> None:
        # Forward compatibility: nothing is parsed into a rigid model, so a new
        # upstream field is reachable rather than dropped or fatal.
        with OptiCMP(
            transport=json_transport({**CAMPAIGN, "brand_new": "value"})
        ) as client:
            campaign = client.campaign.get_campaign("camp-1").data

        assert campaign.get("brand_new") == "value"
        # Reachable as an attribute too; the checker cannot know it yet.
        assert cast(Any, campaign).brand_new == "value"

    def test_paginated_items_support_attributes(self) -> None:
        page = {"data": [{"title": "a"}], "pagination": {"next": None}}
        with OptiCMP(transport=json_transport(page)) as client:
            response = client.campaign.list_campaigns()
            assert [item.title for item in response.data["data"]] == ["a"]

    def test_error_bodies_are_wrapped_too(self) -> None:
        with OptiCMP(
            transport=json_transport({"message": "Not Found"}, status=404)
        ) as client:
            with pytest.raises(APIError) as info:
                client.campaign.get_campaign("missing")

        assert info.value.data.message == "Not Found"

    def test_non_json_payloads_are_left_alone(self) -> None:
        transport = httpx.MockTransport(
            lambda _r: httpx.Response(
                200, text="hello", headers={"content-type": "text/plain"}
            )
        )
        with OptiCMP(transport=transport) as client:
            assert client.request("/whatever").data == "hello"

    async def test_async_payloads_are_wrapped(self) -> None:
        async with AsyncOptiCMP(transport=json_transport(CAMPAIGN)) as client:
            campaign = (await client.campaign.get_campaign("camp-1")).data

        budget = campaign.budget
        assert budget is not None
        assert budget.budgeted_amount == "100.00"
