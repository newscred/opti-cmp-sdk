from __future__ import annotations

import threading
from collections.abc import Callable
from dataclasses import dataclass, field
from typing import Any, Literal

DEFAULT_AUTHORIZATION_URL = "https://accounts.cmp.optimizely.com/o/oauth2/v1/auth"
DEFAULT_TOKEN_URL = "https://accounts.cmp.optimizely.com/o/oauth2/v1/token"
DEFAULT_USERINFO_URL = "https://accounts.cmp.optimizely.com/o/oauth2/v1/userinfo"

#: Refresh this many seconds before the access token actually expires.
EXPIRY_BUFFER_SECONDS = 60.0

GrantType = Literal["authorization_code", "client_credentials"]


# Credentials are kept out of __repr__ so they do not surface in tracebacks,
# logs, or a debugger's variable pane.


@dataclass
class AuthToken:
    access_token: str = field(repr=False)
    #: Unix epoch **seconds**, as `time.time()` returns. The TypeScript SDK
    #: stores milliseconds, so a token persisted by it must be divided by 1000
    #: before being seeded here — otherwise the expiry never arrives and a dead
    #: token is sent until the API rejects it.
    expires_at: float
    refresh_token: str | None = field(default=None, repr=False)

    def __post_init__(self) -> None:
        # Cheap guard against the millisecond mix-up: year 10000 in seconds.
        if self.expires_at > 253_402_300_800:
            raise ValueError(
                "expires_at looks like milliseconds; it must be Unix epoch seconds"
            )


@dataclass(frozen=True)
class TokenAuth:
    """A static access token."""

    token: str = field(repr=False)

    def __post_init__(self) -> None:
        if not self.token:
            raise ValueError("Token authentication requires a token")


@dataclass(frozen=True)
class OAuth:
    """OAuth credentials. Tokens are fetched on demand and refreshed as needed."""

    client_id: str
    client_secret: str = field(repr=False)
    grant_type: GrantType
    #: Client-credentials only; ignored for the authorization-code grant, which
    #: takes its scope from the authorization request.
    scope: str | None = None
    #: Seed an already-issued token to skip the first round trip.
    token: AuthToken | None = None
    token_url: str | None = None
    on_token_refresh: Callable[[AuthToken], Any] | None = None


#: The class is the discriminator, so no `type` field is needed.
AuthOptions = TokenAuth | OAuth


@dataclass
class OAuthState:
    options: OAuth
    token: AuthToken | None = None
    lock: threading.Lock = field(default_factory=threading.Lock)
    #: Created lazily so a sync client never touches asyncio.
    async_lock: Any = None


@dataclass
class AuthState:
    token: str | None = None
    oauth: OAuthState | None = None
