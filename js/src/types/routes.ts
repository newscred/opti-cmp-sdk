import type { paths } from "./openapi.js";

export type EndpointPath = Prettify<keyof paths>;

export type EndpointRoute = Prettify<
  {
    [P in keyof paths]: ExtractRoutes<P>;
  }[keyof paths]
>;

type ExtractRoutes<P extends keyof paths> = {
  [M in HttpMethod]: paths[P][M] extends { responses: unknown }
    ? `${UppercaseMethod<M>} ${P & string}`
    : never;
}[HttpMethod];

type HttpMethod = "delete" | "get" | "patch" | "post" | "put";

type Prettify<T> = {
  [K in keyof T]: T[K];
} & {};

type UppercaseMethod<M extends string> = M extends "get"
  ? "GET"
  : M extends "put"
    ? "PUT"
    : M extends "post"
      ? "POST"
      : M extends "delete"
        ? "DELETE"
        : M extends "patch"
          ? "PATCH"
          : never;
