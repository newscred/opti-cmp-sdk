"""Token and OAuth authentication."""

from __future__ import annotations

from typing import Any, cast

from ..._types import EndpointOptions, Response
from ._oauth import (
    async_lock,
    authorization_code_params,
    build_authorization_url,
    client_credentials_params,
    grant_params,
    needs_refresh,
    request_token,
    request_token_async,
    request_user_info,
    request_user_info_async,
    store_token,
    token_url,
)
from ._types import (
    AuthOptions,
    AuthState,
    AuthToken,
    OAuth,
    OAuthState,
    TokenAuth,
)

__all__ = [
    "AsyncOAuthMethods",
    "AuthOptions",
    "AuthToken",
    "OAuth",
    "OAuthMethods",
    "TokenAuth",
    "auth_plugin",
]

_NOT_CONFIGURED = (
    "OAuth is not configured. Pass auth=OAuth(...) when creating the client, "
    "or call client.authenticate(OAuth(...))."
)


def authenticate(state: AuthState, options: AuthOptions | None = None) -> None:
    """Replace the client's credentials. `None` clears them."""
    if options is None:
        state.token = None
        state.oauth = None
    elif isinstance(options, OAuth):
        state.oauth = OAuthState(options=options, token=options.token)
        state.token = None
    else:
        # TokenAuth validates the token in __post_init__.
        state.token = options.token
        state.oauth = None


class _OAuthBase:
    def __init__(self, client: Any, state: AuthState) -> None:
        self._client = client
        self._state = state

    def _require(self) -> OAuthState:
        if self._state.oauth is None:
            raise ValueError(_NOT_CONFIGURED)
        return self._state.oauth

    def get_authorization_url(
        self,
        *,
        redirect_uri: str,
        authorization_url: str | None = None,
        scope: str | None = None,
        state: str | None = None,
    ) -> str:
        oauth = self._require()
        return build_authorization_url(
            client_id=oauth.options.client_id,
            redirect_uri=redirect_uri,
            authorization_url=authorization_url,
            scope=scope,
            state=state,
        )


class OAuthMethods(_OAuthBase):
    """Synchronous OAuth operations, exposed as ``client.oauth``."""

    def exchange_code(self, *, code: str, redirect_uri: str) -> AuthToken:
        oauth = self._require()
        token = request_token(
            self._client._httpx,
            token_url(oauth.options),
            authorization_code_params(
                oauth.options, code=code, redirect_uri=redirect_uri
            ),
        )
        # Held for the store, not the fetch: a concurrent `ensure_token`
        # refresh must not interleave with this write.
        with oauth.lock:
            return self._store(oauth, token)

    def get_client_credentials_token(self) -> AuthToken:
        oauth = self._require()
        token = request_token(
            self._client._httpx,
            token_url(oauth.options),
            client_credentials_params(oauth.options),
        )
        with oauth.lock:
            return self._store(oauth, token)

    def refresh_token(self) -> AuthToken:
        """Force a refresh now.

        Unlike :meth:`ensure_token`, concurrent callers are serialised rather
        than coalesced: asking explicitly for a refresh gets you one.
        """
        oauth = self._require()
        with oauth.lock:
            return self._refresh(oauth)

    def get_user_info(self, *, user_info_url: str | None = None) -> Response[Any]:
        token = self.ensure_token()
        return request_user_info(self._client._httpx, token.access_token, user_info_url)

    def ensure_token(self) -> AuthToken:
        oauth = self._require()
        if not needs_refresh(oauth):
            return cast(AuthToken, oauth.token)
        with oauth.lock:
            # Another thread may have refreshed while this one waited.
            if not needs_refresh(oauth):
                return cast(AuthToken, oauth.token)
            return self._refresh(oauth)

    def _refresh(self, oauth: OAuthState) -> AuthToken:
        token = request_token(
            self._client._httpx, token_url(oauth.options), grant_params(oauth)
        )
        return self._store(oauth, token)

    def _store(self, oauth: OAuthState, token: AuthToken) -> AuthToken:
        stored = store_token(oauth, token)
        callback = oauth.options.on_token_refresh
        if callback is not None:
            result = callback(stored)
            if hasattr(result, "__await__"):
                # Nothing here can await it, and discarding it would persist no
                # token while only emitting a "never awaited" RuntimeWarning.
                result.close()
                raise TypeError(
                    "on_token_refresh must be a regular function on a "
                    "synchronous client; use AsyncOptiCMP for a coroutine"
                )
        return stored


class AsyncOAuthMethods(_OAuthBase):
    """Asynchronous OAuth operations, exposed as ``client.oauth``."""

    async def exchange_code(self, *, code: str, redirect_uri: str) -> AuthToken:
        oauth = self._require()
        token = await request_token_async(
            self._client._httpx,
            token_url(oauth.options),
            authorization_code_params(
                oauth.options, code=code, redirect_uri=redirect_uri
            ),
        )
        # Held for the store, not the fetch: a concurrent `ensure_token`
        # refresh must not interleave with this write.
        async with async_lock(oauth):
            return await self._store(oauth, token)

    async def get_client_credentials_token(self) -> AuthToken:
        oauth = self._require()
        token = await request_token_async(
            self._client._httpx,
            token_url(oauth.options),
            client_credentials_params(oauth.options),
        )
        async with async_lock(oauth):
            return await self._store(oauth, token)

    async def refresh_token(self) -> AuthToken:
        oauth = self._require()
        async with async_lock(oauth):
            return await self._refresh(oauth)

    async def get_user_info(self, *, user_info_url: str | None = None) -> Response[Any]:
        token = await self.ensure_token()
        return await request_user_info_async(
            self._client._httpx, token.access_token, user_info_url
        )

    async def ensure_token(self) -> AuthToken:
        oauth = self._require()
        if not needs_refresh(oauth):
            return cast(AuthToken, oauth.token)
        async with async_lock(oauth):
            if not needs_refresh(oauth):
                return cast(AuthToken, oauth.token)
            return await self._refresh(oauth)

    async def _refresh(self, oauth: OAuthState) -> AuthToken:
        token = await request_token_async(
            self._client._httpx, token_url(oauth.options), grant_params(oauth)
        )
        return await self._store(oauth, token)

    async def _store(self, oauth: OAuthState, token: AuthToken) -> AuthToken:
        stored = store_token(oauth, token)
        callback = oauth.options.on_token_refresh
        if callback is not None:
            result = callback(stored)
            if hasattr(result, "__await__"):
                await result
        return stored


def _authorize(options: EndpointOptions, access_token: str | None) -> None:
    if access_token:
        options.headers["authorization"] = f"Bearer {access_token}"


def _make_before(state: AuthState, methods: OAuthMethods) -> Any:
    def before(options: EndpointOptions) -> None:
        token = methods.ensure_token().access_token if state.oauth else state.token
        _authorize(options, token)

    return before


def _make_before_async(state: AuthState, methods: AsyncOAuthMethods) -> Any:
    async def before(options: EndpointOptions) -> None:
        token = (
            (await methods.ensure_token()).access_token if state.oauth else state.token
        )
        _authorize(options, token)

    return before


def auth_plugin(client: Any, options: dict[str, Any]) -> None:
    state = AuthState()

    if client._is_async:
        async_methods = AsyncOAuthMethods(client, state)
        client.oauth = async_methods
        before = _make_before_async(state, async_methods)
    else:
        sync_methods = OAuthMethods(client, state)
        client.oauth = sync_methods
        before = _make_before(state, sync_methods)

    client.authenticate = lambda auth=None: authenticate(state, auth)

    if options.get("auth") is not None:
        authenticate(state, options["auth"])

    client.request_hook.before(before)
