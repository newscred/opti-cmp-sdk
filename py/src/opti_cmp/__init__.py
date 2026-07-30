"""Python SDK for the Optimizely CMP API."""

from typing import TYPE_CHECKING

from ._client import AsyncOptiCMP, BaseClient, DynamicNamespace, OptiCMP, Plugin
from ._endpoint import Endpoint, endpoint
from ._errors import (
    APIConnectionError,
    APIError,
    APITimeoutError,
    InvalidRequestError,
    OptiCMPError,
)
from ._hook import (
    AfterHook,
    AsyncSingularHook,
    BeforeHook,
    ErrorHook,
    SingularHook,
    WrapHook,
)
from ._object import APIObject
from ._request import AsyncRequest, Request
from ._types import (
    EndpointOptions,
    EndpointParams,
    Headers,
    HeaderTypes,
    RequestConfig,
    RequestOptions,
    Response,
    Route,
    Routes,
)
from .plugins.auth import (
    AuthOptions,
    AuthToken,
    OAuth,
    TokenAuth,
)
from .version import VERSION

if TYPE_CHECKING:
    from ._generated import objects, schema
else:
    # Nothing at runtime instantiates these: every generated endpoint module
    # imports them under `TYPE_CHECKING`, and `objects` only declares attributes
    # for the checker. They are ~8k lines of class statements, so importing them
    # eagerly costs about 14% of `import opti_cmp` for callers that never write
    # an annotation. PEP 562 defers that to first access.
    _LAZY = ("objects", "schema")

    def __getattr__(name: str) -> object:
        if name in _LAZY:
            import importlib

            module = importlib.import_module(f"._generated.{name}", __name__)
            globals()[name] = module
            return module
        raise AttributeError(f"module {__name__!r} has no attribute {name!r}")

    def __dir__() -> list[str]:
        return sorted(__all__)


__all__ = [
    "VERSION",
    "APIConnectionError",
    "APIError",
    "APIObject",
    "APITimeoutError",
    "AfterHook",
    "AsyncOptiCMP",
    "AsyncRequest",
    "AsyncSingularHook",
    "AuthOptions",
    "AuthToken",
    "BaseClient",
    "BeforeHook",
    "DynamicNamespace",
    "Endpoint",
    "EndpointOptions",
    "EndpointParams",
    "ErrorHook",
    "HeaderTypes",
    "Headers",
    "InvalidRequestError",
    "OAuth",
    "OptiCMP",
    "OptiCMPError",
    "Plugin",
    "Request",
    "RequestConfig",
    "RequestOptions",
    "Response",
    "Route",
    "Routes",
    "SingularHook",
    "TokenAuth",
    "WrapHook",
    "endpoint",
    "objects",
    "schema",
]
