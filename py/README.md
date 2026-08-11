# opti-cmp

Python SDK for the Optimizely CMP API.

## Install

```bash
pip install opti-cmp
# or
uv add opti-cmp
```

## Usage

```python
from opti_cmp import OptiCMP, TokenAuth

with OptiCMP(auth=TokenAuth("<auth-token>")) as client:
    campaign = client.campaign.get_campaign("campaign-id").data

    print(campaign.title)

    # `budget` is optional in the specification, so narrow before reaching in.
    if campaign.budget:
        print(campaign.budget.currency_code)
```

Every call returns a `Response` with `data`, `headers`, `status` and `url`.
`headers` is an `httpx.Headers`, so lookups ignore case and repeated headers
such as `Set-Cookie` are preserved:

```python
response.headers["X-Request-Id"]  # any case
response.headers.get_list("set-cookie")  # every value
```

### Reading payloads

`response.data` reads either by attribute or by key, and nesting works both ways:

```python
campaign.title  # checked against the specification
campaign["title"]  # same value, untyped
campaign.budget.currency_code  # nested objects are wrapped too
[label.group.name for label in campaign.labels]
```

Attribute access is the checked path: a type checker knows every field the
specification declares, so `campaign.titel` is caught before you run it. The
subscript form always works but is untyped.

Three cases worth knowing:

- **A field the API omitted** raises `AttributeError`. Use
  `campaign.get("description")` for fields that are genuinely optional.
- **A field newer than the specification** still arrives and is still readable —
  nothing is parsed into a rigid model, so new upstream fields are never
  dropped. A type checker will not know about it yet, so reach it with
  `campaign.get("new_field")`.
- **Request bodies** stay plain dicts, checked key by key against a `TypedDict`:

  ```python
  client.campaign.update_campaign("campaign-id", body={"title": "Q3"})
  ```

The generated types are exported if you want to annotate your own code:
`opti_cmp.schema` holds the request-body `TypedDict`s, `opti_cmp.objects` the
response classes.

```python
from opti_cmp import schema


def rename(title: str) -> schema.CampaignUpdateRequest:
    return {"title": title}
```

### Async

```python
from opti_cmp import AsyncOptiCMP, TokenAuth

async with AsyncOptiCMP(auth=TokenAuth("<auth-token>")) as client:
    campaign = (await client.campaign.get_campaign("campaign-id")).data
```

### OAuth

```python
from opti_cmp import OAuth, OptiCMP

client = OptiCMP(
    auth=OAuth(
        grant_type="client_credentials",
        client_id="<client-id>",
        client_secret="<client-secret>",
    )
)
```

Access tokens are fetched on demand and refreshed a minute before they expire.
For the authorization code grant, use `client.oauth.get_authorization_url(...)`
and `client.oauth.exchange_code(...)`, and pass `on_token_refresh` to persist
new tokens.

Seeding a token you stored earlier skips the first round trip:

```python
from opti_cmp import AuthToken

OAuth(..., token=AuthToken(access_token="...", expires_at=1767225600.0))
```

> `expires_at` is Unix epoch **seconds**, as `time.time()` returns. The
> TypeScript SDK stores milliseconds — divide by 1000 when carrying a token
> across. A millisecond value is rejected rather than silently never expiring.

### Pagination

Paginated endpoints return one page at a time. Step through them with
`has_next_page` and `get_next_page`:

```python
response = client.campaign.list_campaigns(page_size=50)

while True:
    for campaign in response.data["data"]:
        print(campaign.title)
    if not client.has_next_page(response):
        break
    response = client.get_next_page(response)
```

`has_previous_page` and `get_previous_page` walk the other way. On
`AsyncOptiCMP`, `get_next_page` and `get_previous_page` are awaitable;
`has_next_page` and `has_previous_page` are not.

`get_next_page` raises `ValueError` when there is no next page, so check
`has_next_page` first.

### Errors

```python
from opti_cmp import (
    APIConnectionError,
    APIError,
    APITimeoutError,
    InvalidRequestError,
    OptiCMPError,
)

try:
    client.campaign.get_campaign("missing")
except APIError as error:  # the server answered
    print(error.status, error.data)
except APITimeoutError:  # no response before the deadline
    ...
except APIConnectionError:  # never reached the server; no status exists
    ...
except InvalidRequestError:  # bad URL or header; nothing was sent
    ...
except OptiCMPError:  # any of the above
    ...
```

A failure that never produced a response is an `APIConnectionError`, not an
`APIError` with an invented status, so a real server 500 and an unreachable
host are distinguishable.

### Hooks

`client.request_hook` wraps every request. Register callbacks by call or as
decorators — use it for logging, metrics, tracing headers, retries, or custom
error handling.

```python
import time

from opti_cmp import APIError

client = OptiCMP(auth=TokenAuth("<auth-token>"))


@client.request_hook.before
def add_tenant(options):
    """Runs before the request. Mutate `options` in place."""
    options.headers["x-tenant"] = "acme"


@client.request_hook.after
def log_response(response, options):
    """Runs on success. Return non-None to replace the response."""
    log.info("%s %s -> %s", options.method, options.url, response.status)


@client.request_hook.error
def count_failures(error, options):
    """Runs on failure. Return None to let the error propagate."""
    # Only APIError carries a status — a connection failure has none.
    status = error.status if isinstance(error, APIError) else None
    metrics.increment("cmp.error", status=status)


@client.request_hook.wrap
def measure(inner, options):
    """Wraps the whole call. You are responsible for calling `inner`."""
    started = time.perf_counter()
    try:
        return inner(options)
    finally:
        metrics.timing("cmp.duration", time.perf_counter() - started)
```

`wrap` encloses the before hooks too, so a retry wrap re-runs them — which is
what you want, because OAuth fetches its token from a before hook. The
last-registered wrap sits closest to the request.

`options` is the merged endpoint options. The parts the SDK understands read
and write as attributes — `method`, `url`, `headers`, `request`, `namespace`,
`operation` — and the call's own parameters are reachable by key:

```python
options.headers["x-tenant"] = "acme"
options.url  # "/campaigns/{id}", still templated
options["page_size"]  # the call's parameters
```

`url` stays templated, which is what you want as a metric label; the concrete
URL would be one label per campaign.

`namespace` and `operation` are set for generated endpoint calls only. A raw
`client.request(...)` has neither — including the one behind `get_next_page` —
so they read as `None` rather than raising:

```python
label = options.operation or f"{options.method} {options.url}"
```

It is still a `dict`, so `options["headers"]` works too.

Every registered hook of a kind runs, in registration order. For `after` and
`error`, returning `None` means "observed, no change":

- an `error` hook that only logs will **not** swallow the exception,
- returning a value from an `error` hook recovers, and the first hook to do so
  wins,
- returning a value from an `after` hook replaces the response.

Use `client.request_hook.remove(callback)` to unregister.

On `AsyncOptiCMP`, hooks may be either plain functions or coroutine functions.

> **Note for users of the JavaScript SDK:** its error hooks recover unless they
> re-throw. Here an error hook must return a value to recover, so the common
> log-only hook is safe by default.

#### Why there is only one hook

`httpx`'s own `event_hooks` are deliberately not exposed. They can observe a
request and a response, but their return values are ignored, so they cannot
replace a response, and httpx has no error event at all — it accepts
`event_hooks={"error": [...]}` without complaint and never fires it. Everything
they can do, `request_hook` does, with the operation's identity and the parsed
response instead of raw bytes.

Passing `event_hooks=` is accepted and ignored: unrecognised keywords are
collected for plugins rather than rejected, so nothing validates them.

For control below the SDK — retries, proxies, connection limits, or an existing
httpx instrumentation stack — pass a custom `transport=`, which httpx builds on
directly.

### Timeouts

Requests time out after 30 seconds, with a 10-second connect deadline. Set
`request={"timeout": seconds}` on the client to change it, or per call to
override it once. `None` disables the deadline — use it for large uploads.

```python
client = OptiCMP(auth=TokenAuth("<auth-token>"), request={"timeout": 60.0})

client.campaign.list_campaigns(request={"timeout": None})
```

Timeouts are **seconds**, the Python convention; the TypeScript SDK's
`AbortSignal.timeout` takes milliseconds.

## Example

`examples/smoke.py` runs the SDK against the real API — read-only, so it never
creates or changes anything:

```bash
export OPTI_CMP_TOKEN=<token>
uv run python examples/smoke.py
```

It covers auth, listing, pagination, fetching one record, attribute access on
nested payloads, error handling and the async client.

## Development

The endpoint surface is generated from the shared `specification/` directory at
the repository root.

```bash
uv sync
uv run python scripts/generate_schema_types.py
uv run python scripts/generate_endpoints.py
uv run ruff check . && uv run ruff format --check .
uv run mypy
uv run pytest
```
