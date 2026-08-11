"""OAuth grant flows and token refresh."""

from __future__ import annotations

import time
from typing import TYPE_CHECKING, Any

import httpx

from ..._request import asend, handle_response, send, transport_error
from ..._types import Headers, RequestOptions, Response
from ._types import (
    DEFAULT_AUTHORIZATION_URL,
    DEFAULT_TOKEN_URL,
    DEFAULT_USERINFO_URL,
    EXPIRY_BUFFER_SECONDS,
    AuthToken,
    OAuth,
    OAuthState,
)

if TYPE_CHECKING:
    import asyncio

_FORM_HEADERS = {
    "accept": "application/json",
    "content-type": "application/x-www-form-urlencoded",
}


def build_authorization_url(
    *,
    client_id: str,
    redirect_uri: str,
    authorization_url: str | None = None,
    scope: str | None = None,
    state: str | None = None,
) -> str:
    query: dict[str, str] = {
        "client_id": client_id,
        "redirect_uri": redirect_uri,
        "response_type": "code",
    }
    if scope:
        query["scope"] = scope
    if state:
        query["state"] = state

    base = httpx.URL(authorization_url or DEFAULT_AUTHORIZATION_URL)
    return str(base.copy_merge_params(query))


def authorization_code_params(
    options: OAuth, *, code: str, redirect_uri: str
) -> dict[str, str]:
    return {
        "client_id": options.client_id,
        "client_secret": options.client_secret,
        "code": code,
        "grant_type": "authorization_code",
        "redirect_uri": redirect_uri,
    }


def client_credentials_params(options: OAuth) -> dict[str, str]:
    params = {
        "client_id": options.client_id,
        "client_secret": options.client_secret,
        "grant_type": "client_credentials",
    }
    scope = options.scope
    if scope:
        params["scope"] = scope
    return params


def refresh_params(options: OAuth, refresh_token: str) -> dict[str, str]:
    return {
        "client_id": options.client_id,
        "client_secret": options.client_secret,
        "grant_type": "refresh_token",
        "refresh_token": refresh_token,
    }


def grant_params(state: OAuthState) -> dict[str, str]:
    """Parameters for the grant that renews the current token."""
    options = state.options
    if options.grant_type == "authorization_code":
        refresh_token = state.token.refresh_token if state.token else None
        if not refresh_token:
            raise ValueError("Cannot refresh the access token without a refresh token")
        return refresh_params(options, refresh_token)
    return client_credentials_params(options)


def token_url(options: OAuth) -> str:
    return options.token_url or DEFAULT_TOKEN_URL


def needs_refresh(state: OAuthState) -> bool:
    return (
        state.token is None
        or time.time() >= state.token.expires_at - EXPIRY_BUFFER_SECONDS
    )


def store_token(state: OAuthState, token: AuthToken) -> AuthToken:
    """Keep the previous refresh token when the grant response omits one."""
    if token.refresh_token is None and state.token is not None:
        token.refresh_token = state.token.refresh_token
    state.token = token
    return token


def request_token(client: httpx.Client, url: str, params: dict[str, str]) -> AuthToken:
    options = _token_request(url)
    try:
        response = client.post(url, data=params, headers=_FORM_HEADERS)
    except httpx.HTTPError as error:
        raise transport_error(error, options) from error
    return _token_from(response, options)


async def request_token_async(
    client: httpx.AsyncClient, url: str, params: dict[str, str]
) -> AuthToken:
    options = _token_request(url)
    try:
        response = await client.post(url, data=params, headers=_FORM_HEADERS)
    except httpx.HTTPError as error:
        raise transport_error(error, options) from error
    return _token_from(response, options)


def request_user_info(
    client: httpx.Client, access_token: str, url: str | None = None
) -> Response[dict[str, Any]]:
    return send(client, _user_info_request(access_token, url))


async def request_user_info_async(
    client: httpx.AsyncClient, access_token: str, url: str | None = None
) -> Response[dict[str, Any]]:
    return await asend(client, _user_info_request(access_token, url))


def async_lock(state: OAuthState) -> asyncio.Lock:
    """The refresh lock for this state, created on first use.

    An `asyncio.Lock` binds to the loop that first contends it, so an
    `AsyncOptiCMP` belongs to one event loop. Reusing one across two
    `asyncio.run()` calls, or across threads each running their own loop, is
    not supported — build a client per loop.
    """
    # Imported here, not at module scope: this module is reached unconditionally
    # from `opti_cmp/__init__`, and asyncio is ~13ms of the package's import
    # time. A sync-only process never calls this.
    import asyncio

    if state.async_lock is None:
        state.async_lock = asyncio.Lock()
    lock: asyncio.Lock = state.async_lock
    return lock


def _user_info_request(access_token: str, url: str | None) -> RequestOptions:
    return RequestOptions(
        method="GET",
        url=url or DEFAULT_USERINFO_URL,
        headers=Headers(
            {"accept": "application/json", "authorization": f"Bearer {access_token}"}
        ),
    )


def _token_request(url: str) -> RequestOptions:
    """The request as errors should describe it.

    `body` is left unset on purpose: the real one carries client_secret,
    refresh_token and code, none of which belong in a traceback.
    """
    return RequestOptions(method="POST", url=url)


def _token_from(response: httpx.Response, options: RequestOptions) -> AuthToken:
    payload = handle_response(response, options).data
    return AuthToken(
        access_token=payload["access_token"],
        expires_at=time.time() + payload["expires_in"],
        refresh_token=payload.get("refresh_token"),
    )
