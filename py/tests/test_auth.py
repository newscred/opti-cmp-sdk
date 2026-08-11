"""Token and OAuth authentication."""

from __future__ import annotations

import asyncio
import threading
import time
from collections.abc import Callable
from concurrent.futures import ThreadPoolExecutor
from dataclasses import replace
from typing import Any
from urllib.parse import parse_qs, urlsplit

import httpx
import pytest

from opti_cmp import APIError, AsyncOptiCMP, AuthToken, OAuth, OptiCMP, TokenAuth

TOKEN_URL = "https://accounts.cmp.optimizely.com/o/oauth2/v1/token"
USERINFO_URL = "https://accounts.cmp.optimizely.com/o/oauth2/v1/userinfo"

CLIENT_CREDENTIALS = OAuth(
    grant_type="client_credentials",
    client_id="id-1",
    client_secret="secret-1",
)

AUTHORIZATION_CODE = OAuth(
    grant_type="authorization_code",
    client_id="id-1",
    client_secret="secret-1",
)


class Server:
    """Answers token/userinfo calls and records every request."""

    def __init__(
        self,
        *,
        token: dict[str, Any] | Callable[[int], dict[str, Any]] | None = None,
        token_status: int = 200,
        userinfo: dict[str, Any] | None = None,
        userinfo_status: int = 200,
    ) -> None:
        self.requests: list[httpx.Request] = []
        self.token_payload = token or {
            "access_token": "access-1",
            "expires_in": 3600,
            "token_type": "Bearer",
        }
        self.token_status = token_status
        self.userinfo_payload = userinfo or {"sub": "user-1"}
        self.userinfo_status = userinfo_status
        self.token_calls = 0

    def handler(self, request: httpx.Request) -> httpx.Response:
        self.requests.append(request)
        path = urlsplit(str(request.url)).path

        if path.endswith("/oauth2/v1/token"):
            self.token_calls += 1
            payload = self.token_payload
            if callable(payload):
                payload = payload(self.token_calls)
            return httpx.Response(self.token_status, json=payload)

        if path.endswith("/oauth2/v1/userinfo"):
            return httpx.Response(self.userinfo_status, json=self.userinfo_payload)

        return httpx.Response(200, json={"ok": True})

    @property
    def transport(self) -> httpx.MockTransport:
        return httpx.MockTransport(self.handler)

    def token_form(self, index: int = 0) -> dict[str, str]:
        forms = [
            parse_qs(r.read().decode())
            for r in self.requests
            if urlsplit(str(r.url)).path.endswith("/token")
        ]
        return {key: value[0] for key, value in forms[index].items()}

    def api_requests(self) -> list[httpx.Request]:
        return [
            r for r in self.requests if "accounts.cmp.optimizely.com" not in str(r.url)
        ]


def expiring_token(access: str, expires_in: int) -> dict[str, Any]:
    return {"access_token": access, "expires_in": expires_in, "token_type": "Bearer"}


class TestNoAuth:
    def test_leaves_requests_unauthenticated(self) -> None:
        server = Server()
        with OptiCMP(transport=server.transport) as client:
            client.request("/campaigns")

        assert "authorization" not in server.api_requests()[0].headers


class TestTokenAuth:
    def test_sets_a_bearer_header(self) -> None:
        server = Server()
        with OptiCMP(auth=TokenAuth("tok-1"), transport=server.transport) as client:
            client.request("/campaigns")

        assert server.api_requests()[0].headers["authorization"] == "Bearer tok-1"

    def test_can_be_configured_later(self) -> None:
        server = Server()
        with OptiCMP(transport=server.transport) as client:
            client.authenticate(TokenAuth("later"))
            client.request("/campaigns")

        assert server.api_requests()[0].headers["authorization"] == "Bearer later"

    def test_can_be_cleared(self) -> None:
        server = Server()
        with OptiCMP(auth=TokenAuth("tok-1"), transport=server.transport) as client:
            client.authenticate(None)
            client.request("/campaigns")

        assert "authorization" not in server.api_requests()[0].headers

    def test_rejects_an_empty_token(self) -> None:
        with pytest.raises(ValueError, match="requires a token"):
            TokenAuth("")


class TestCredentialsStayOutOfRepr:
    """A repr lands in tracebacks, logs and debugger panes."""

    def test_oauth_hides_the_client_secret(self) -> None:
        secret = "super-secret"
        auth = OAuth(
            client_id="id-1", client_secret=secret, grant_type="client_credentials"
        )
        assert secret not in repr(auth)
        assert "id-1" in repr(auth)
        assert auth.client_secret == secret

    def test_token_auth_hides_the_token(self) -> None:
        auth = TokenAuth("super-secret")
        assert "super-secret" not in repr(auth)
        assert auth.token == "super-secret"

    def test_auth_token_hides_both_tokens(self) -> None:
        token = AuthToken(
            access_token="access-secret",
            expires_at=1.0,
            refresh_token="refresh-secret",
        )
        assert "secret" not in repr(token)
        assert token.access_token == "access-secret"


class TestClientCredentials:
    def test_fetches_a_token_then_authorizes_the_request(self) -> None:
        server = Server()
        with OptiCMP(auth=CLIENT_CREDENTIALS, transport=server.transport) as client:
            client.request("/campaigns")

        assert server.token_calls == 1
        assert server.token_form() == {
            "client_id": "id-1",
            "client_secret": "secret-1",
            "grant_type": "client_credentials",
        }
        assert server.api_requests()[0].headers["authorization"] == "Bearer access-1"

    def test_includes_scope_and_honours_a_custom_token_url(self) -> None:
        server = Server()
        with OptiCMP(
            auth=replace(
                CLIENT_CREDENTIALS,
                scope="openid profile",
                token_url=TOKEN_URL,
            ),
            transport=server.transport,
        ) as client:
            client.request("/campaigns")

        assert server.token_form()["scope"] == "openid profile"

    def test_reuses_a_valid_token(self) -> None:
        server = Server()
        with OptiCMP(auth=CLIENT_CREDENTIALS, transport=server.transport) as client:
            client.request("/campaigns")
            client.request("/campaigns")

        assert server.token_calls == 1

    def test_refreshes_a_token_inside_the_expiry_buffer(self) -> None:
        # 30s of life left, and the buffer is 60s, so it must refetch.
        server = Server(token=lambda n: expiring_token(f"access-{n}", 30))
        with OptiCMP(auth=CLIENT_CREDENTIALS, transport=server.transport) as client:
            client.request("/campaigns")
            client.request("/campaigns")

        assert server.token_calls == 2
        assert server.api_requests()[1].headers["authorization"] == "Bearer access-2"

    def test_uses_a_seeded_token_without_fetching(self) -> None:
        server = Server()
        seeded = AuthToken(access_token="seeded", expires_at=time.time() + 3600)
        with OptiCMP(
            auth=replace(CLIENT_CREDENTIALS, token=seeded), transport=server.transport
        ) as client:
            client.request("/campaigns")

        assert server.token_calls == 0
        assert server.api_requests()[0].headers["authorization"] == "Bearer seeded"

    def test_refetches_when_the_seeded_token_expired(self) -> None:
        server = Server()
        stale = AuthToken(access_token="stale", expires_at=time.time() - 10)
        with OptiCMP(
            auth=replace(CLIENT_CREDENTIALS, token=stale), transport=server.transport
        ) as client:
            client.request("/campaigns")

        assert server.token_calls == 1
        assert server.api_requests()[0].headers["authorization"] == "Bearer access-1"

    def test_invokes_on_token_refresh(self) -> None:
        issued: list[AuthToken] = []
        server = Server()
        with OptiCMP(
            auth=replace(CLIENT_CREDENTIALS, on_token_refresh=issued.append),
            transport=server.transport,
        ) as client:
            client.request("/campaigns")

        assert [token.access_token for token in issued] == ["access-1"]

    def test_raises_http_error_when_the_token_endpoint_fails(self) -> None:
        server = Server(token_status=401, token={"error": "invalid_client"})
        with OptiCMP(auth=CLIENT_CREDENTIALS, transport=server.transport) as client:
            with pytest.raises(APIError) as info:
                client.request("/campaigns")

        assert info.value.status == 401
        assert info.value.data == {"error": "invalid_client"}

    def test_token_errors_do_not_leak_the_client_secret(self) -> None:
        server = Server(token_status=401, token={"error": "invalid_client"})
        with OptiCMP(auth=CLIENT_CREDENTIALS, transport=server.transport) as client:
            with pytest.raises(APIError) as info:
                client.request("/campaigns")

        assert info.value.request is not None
        assert info.value.request.body is None

    def test_retries_after_a_failed_refresh(self) -> None:
        state = {"fail": True}

        def handler(request: httpx.Request) -> httpx.Response:
            if urlsplit(str(request.url)).path.endswith("/token"):
                if state["fail"]:
                    state["fail"] = False
                    return httpx.Response(500, json={"error": "server_error"})
                return httpx.Response(200, json=expiring_token("recovered", 3600))
            return httpx.Response(200, json={"ok": True})

        with OptiCMP(
            auth=CLIENT_CREDENTIALS, transport=httpx.MockTransport(handler)
        ) as client:
            with pytest.raises(APIError):
                client.request("/campaigns")
            # The failure must not leave a poisoned in-flight refresh behind.
            client.request("/campaigns")


class TestAuthorizationCode:
    def test_refreshes_with_the_refresh_token_grant(self) -> None:
        server = Server(
            token={
                "access_token": "access-2",
                "refresh_token": "refresh-2",
                "expires_in": 3600,
                "token_type": "Bearer",
            }
        )
        seeded = AuthToken(
            access_token="old", expires_at=time.time() - 1, refresh_token="refresh-1"
        )
        with OptiCMP(
            auth=replace(AUTHORIZATION_CODE, token=seeded), transport=server.transport
        ) as client:
            client.request("/campaigns")

        form = server.token_form()
        assert form["grant_type"] == "refresh_token"
        assert form["refresh_token"] == "refresh-1"
        assert client.oauth.ensure_token().refresh_token == "refresh-2"

    def test_keeps_the_previous_refresh_token_when_none_is_returned(self) -> None:
        server = Server()  # response has no refresh_token
        seeded = AuthToken(
            access_token="old", expires_at=time.time() - 1, refresh_token="refresh-1"
        )
        with OptiCMP(
            auth=replace(AUTHORIZATION_CODE, token=seeded), transport=server.transport
        ) as client:
            client.request("/campaigns")
            assert client.oauth.ensure_token().refresh_token == "refresh-1"

    def test_refusing_to_refresh_without_a_refresh_token(self) -> None:
        with OptiCMP(auth=AUTHORIZATION_CODE, transport=Server().transport) as client:
            with pytest.raises(ValueError, match="without a refresh token"):
                client.request("/campaigns")


class TestOAuthMethods:
    def test_methods_raise_until_oauth_is_configured(self) -> None:
        with OptiCMP(transport=Server().transport) as client:
            with pytest.raises(ValueError, match="OAuth is not configured"):
                client.oauth.get_client_credentials_token()

    def test_oauth_can_be_configured_later(self) -> None:
        server = Server()
        with OptiCMP(transport=server.transport) as client:
            client.authenticate(CLIENT_CREDENTIALS)
            client.request("/campaigns")

        assert server.api_requests()[0].headers["authorization"] == "Bearer access-1"

    def test_get_client_credentials_token_returns_and_stores(self) -> None:
        server = Server()
        with OptiCMP(auth=CLIENT_CREDENTIALS, transport=server.transport) as client:
            token = client.oauth.get_client_credentials_token()
            assert token.access_token == "access-1"

            client.request("/campaigns")

        assert server.token_calls == 1

    def test_refresh_token_forces_a_new_token(self) -> None:
        server = Server(token=lambda n: expiring_token(f"access-{n}", 3600))
        with OptiCMP(auth=CLIENT_CREDENTIALS, transport=server.transport) as client:
            client.request("/campaigns")
            assert client.oauth.refresh_token().access_token == "access-2"

        assert server.token_calls == 2

    def test_get_authorization_url(self) -> None:
        with OptiCMP(auth=AUTHORIZATION_CODE, transport=Server().transport) as client:
            url = client.oauth.get_authorization_url(
                redirect_uri="https://app.example.com/cb",
                scope="openid",
                state="xyz",
            )

        parts = urlsplit(url)
        query = {key: value[0] for key, value in parse_qs(parts.query).items()}
        assert parts.netloc == "accounts.cmp.optimizely.com"
        assert parts.path == "/o/oauth2/v1/auth"
        assert query == {
            "client_id": "id-1",
            "redirect_uri": "https://app.example.com/cb",
            "response_type": "code",
            "scope": "openid",
            "state": "xyz",
        }

    def test_get_authorization_url_keeps_the_base_url_s_own_query(self) -> None:
        with OptiCMP(auth=AUTHORIZATION_CODE, transport=Server().transport) as client:
            url = client.oauth.get_authorization_url(
                redirect_uri="https://app.example.com/cb",
                authorization_url="https://idp.example.com/auth?tenant=acme",
            )

        query = {key: value[0] for key, value in parse_qs(urlsplit(url).query).items()}
        assert query == {
            "tenant": "acme",
            "client_id": "id-1",
            "redirect_uri": "https://app.example.com/cb",
            "response_type": "code",
        }

    def test_exchange_code_stores_the_token(self) -> None:
        server = Server()
        with OptiCMP(auth=AUTHORIZATION_CODE, transport=server.transport) as client:
            token = client.oauth.exchange_code(
                code="auth-code", redirect_uri="https://app.example.com/cb"
            )
            assert token.access_token == "access-1"

            client.request("/campaigns")

        form = server.token_form()
        assert form["grant_type"] == "authorization_code"
        assert form["code"] == "auth-code"
        assert server.api_requests()[0].headers["authorization"] == "Bearer access-1"

    def test_get_user_info_returns_claims(self) -> None:
        server = Server(userinfo={"sub": "user-1", "email": "a@b.c"})
        with OptiCMP(auth=CLIENT_CREDENTIALS, transport=server.transport) as client:
            response = client.oauth.get_user_info()

        assert response.status == 200
        assert response.data == {"sub": "user-1", "email": "a@b.c"}

    def test_get_user_info_raises_on_failure(self) -> None:
        server = Server(userinfo_status=401, userinfo={"error": "invalid_token"})
        with OptiCMP(auth=CLIENT_CREDENTIALS, transport=server.transport) as client:
            with pytest.raises(APIError) as info:
                client.oauth.get_user_info()

        assert info.value.status == 401


class TestAsyncAuth:
    async def test_token_auth(self) -> None:
        server = Server()
        async with AsyncOptiCMP(
            auth=TokenAuth("tok-1"), transport=server.transport
        ) as client:
            await client.request("/campaigns")

        assert server.api_requests()[0].headers["authorization"] == "Bearer tok-1"

    async def test_client_credentials(self) -> None:
        server = Server()
        async with AsyncOptiCMP(
            auth=CLIENT_CREDENTIALS, transport=server.transport
        ) as client:
            await client.request("/campaigns")

        assert server.token_calls == 1
        assert server.api_requests()[0].headers["authorization"] == "Bearer access-1"

    async def test_concurrent_requests_share_one_token_fetch(self) -> None:
        server = Server()
        async with AsyncOptiCMP(
            auth=CLIENT_CREDENTIALS, transport=server.transport
        ) as client:
            await asyncio.gather(*(client.request("/campaigns") for _ in range(5)))

        assert server.token_calls == 1

    async def test_awaits_an_async_on_token_refresh(self) -> None:
        issued: list[str] = []

        async def on_refresh(token: AuthToken) -> None:
            await asyncio.sleep(0)
            issued.append(token.access_token)

        server = Server()
        async with AsyncOptiCMP(
            auth=replace(CLIENT_CREDENTIALS, on_token_refresh=on_refresh),
            transport=server.transport,
        ) as client:
            await client.request("/campaigns")

        assert issued == ["access-1"]

    async def test_get_user_info(self) -> None:
        server = Server(userinfo={"sub": "user-9"})
        async with AsyncOptiCMP(
            auth=CLIENT_CREDENTIALS, transport=server.transport
        ) as client:
            response = await client.oauth.get_user_info()

        assert response.data == {"sub": "user-9"}


class TestSyncRefreshLock:
    def test_concurrent_threads_share_one_token_fetch(self) -> None:
        """The sync path coalesces on a threading.Lock, the async one on an
        asyncio.Lock. Both need proving; only the async one was covered.
        """
        server = Server()
        # A barrier rather than a sleep: every thread is guaranteed to be inside
        # the request before any of them reaches the token fetch.
        start = threading.Barrier(5)

        with OptiCMP(auth=CLIENT_CREDENTIALS, transport=server.transport) as client:

            def call(_: int) -> None:
                start.wait(timeout=5)
                client.request("/campaigns")

            with ThreadPoolExecutor(max_workers=5) as pool:
                list(pool.map(call, range(5)))

        assert server.token_calls == 1
        assert len(server.api_requests()) == 5

    def test_a_failed_refresh_does_not_hold_the_lock(self) -> None:
        server = Server(token_status=500)
        with OptiCMP(auth=CLIENT_CREDENTIALS, transport=server.transport) as client:
            with pytest.raises(APIError):
                client.request("/campaigns")
            # A lock left held here would deadlock rather than raise again.
            with pytest.raises(APIError):
                client.request("/campaigns")

        assert server.token_calls == 2

    def test_rejects_a_coroutine_on_token_refresh(self) -> None:
        async def on_refresh(token: AuthToken) -> None: ...

        server = Server()
        with OptiCMP(
            auth=replace(CLIENT_CREDENTIALS, on_token_refresh=on_refresh),
            transport=server.transport,
        ) as client:
            with pytest.raises(TypeError, match="synchronous client"):
                client.request("/campaigns")


class TestTokenUnits:
    def test_rejects_a_millisecond_expires_at(self) -> None:
        # The TypeScript SDK stores milliseconds; carrying one over would make
        # the token look valid for 50,000 years.
        with pytest.raises(ValueError, match="milliseconds"):
            AuthToken(access_token="a", expires_at=time.time() * 1000)

    def test_accepts_a_second_expires_at(self) -> None:
        token = AuthToken(access_token="a", expires_at=time.time() + 3600)
        assert token.access_token == "a"
