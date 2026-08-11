"""Exercise the SDK against the real Optimizely CMP API.

Read-only: this lists and fetches, and never creates, updates or deletes
anything in your organisation.

    export OPTI_CMP_TOKEN=<token>
    uv run python examples/smoke.py

or, for the client-credentials OAuth grant:

    export OPTI_CMP_CLIENT_ID=<id> OPTI_CMP_CLIENT_SECRET=<secret>
    uv run python examples/smoke.py

Point it somewhere other than production with OPTI_CMP_BASE_URL.
"""

from __future__ import annotations

import asyncio
import os
import sys
from typing import Any

from opti_cmp import (
    APIConnectionError,
    APIError,
    AsyncOptiCMP,
    AuthOptions,
    EndpointOptions,
    OAuth,
    OptiCMP,
    TokenAuth,
)


def credentials() -> AuthOptions:
    if token := os.environ.get("OPTI_CMP_TOKEN"):
        return TokenAuth(token)

    client_id = os.environ.get("OPTI_CMP_CLIENT_ID")
    client_secret = os.environ.get("OPTI_CMP_CLIENT_SECRET")
    if client_id and client_secret:
        return OAuth(
            client_id=client_id,
            client_secret=client_secret,
            grant_type="client_credentials",
        )

    sys.exit(
        "Set OPTI_CMP_TOKEN, or both OPTI_CMP_CLIENT_ID and "
        "OPTI_CMP_CLIENT_SECRET. See the docstring at the top of this file."
    )


def options() -> dict[str, Any]:
    opts: dict[str, Any] = {"auth": credentials(), "request": {"timeout": 30.0}}
    if base_url := os.environ.get("OPTI_CMP_BASE_URL"):
        opts["base_url"] = base_url
    return opts


def heading(text: str) -> None:
    print(f"\n{text}\n{'-' * len(text)}")


def main() -> None:
    client = OptiCMP(**options())

    # Log every call. Generated endpoints carry the operation's identity, which
    # is what makes a usable metric label — the concrete URL would be one label
    # per campaign. A raw `client.request(...)`, including the one behind
    # `get_next_page`, has no operation, so fall back to the route.
    @client.request_hook.before
    def log(options: EndpointOptions) -> None:
        print(f"  -> {options.operation or f'{options.method} {options.url}'}")

    with client:
        heading("List campaigns")
        page = client.campaign.list_campaigns(page_size=5)
        campaigns = page.data["data"]
        print(f"  {len(campaigns)} campaign(s) on the first page")
        for campaign in campaigns:
            # Attribute access, checked against the specification.
            print(f"    {campaign.id}  {campaign.title}")

        heading("Pagination")
        if client.has_next_page(page):
            following = client.get_next_page(page)
            print(f"  next page: {len(following.data['data'])} more")
        else:
            print("  only one page")

        heading("Fetch one campaign")
        if campaigns:
            detail = client.campaign.get_campaign(campaigns[0].id).data
            print(f"  title      : {detail.title}")
            print(f"  status     : {detail.status}")
            # Nested objects are wrapped, so this chains — but `budget` is
            # nullable upstream, and the type checker knows it, so narrow first.
            budget = detail.budget
            print(f"  currency   : {budget.currency_code if budget else '(none)'}")
            # `description` is nullable too.
            print(f"  description: {detail.description or '(none)'}")
        else:
            print("  no campaigns to fetch")

        heading("Error handling")
        try:
            client.campaign.get_campaign("00000000-0000-0000-0000-000000000000")
        except APIError as error:
            print(f"  APIError {error.status}: {error.data}")
        else:
            print("  unexpectedly found that campaign")

        heading("Response envelope")
        print(f"  status : {page.status}")
        print(f"  request id: {page.headers.get('x-request-id', '(none)')}")


async def main_async() -> None:
    heading("Async client")
    async with AsyncOptiCMP(**options()) as client:
        page = await client.campaign.list_campaigns(page_size=1)
        print(f"  status {page.status}, {len(page.data['data'])} campaign(s)")


if __name__ == "__main__":
    try:
        main()
        asyncio.run(main_async())
    except APIConnectionError as error:
        # No response arrived, so there is no status to report.
        sys.exit(f"\nCould not reach the API: {error}")
    except APIError as error:
        sys.exit(f"\nRequest failed with {error.status}: {error.data}")
    print("\nDone.")
