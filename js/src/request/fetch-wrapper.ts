import type { Headers, RequestOptions } from "../endpoint/types.js";
import type { Response } from "./types.js";

import { HTTPError } from "../error/index.js";

export async function fetchWrapper<T>(
  requestOptions: RequestOptions,
): Promise<Response<T>> {
  const { body, headers, method, request, url } = requestOptions;

  const options: RequestInit = {
    body: body as RequestInit["body"],
    headers,
    method,
    signal: AbortSignal.any(
      [
        request?.signal,
        request?.timeout ? AbortSignal.timeout(request.timeout) : null,
      ].filter((signal): signal is AbortSignal => Boolean(signal)),
    ),
  };

  let responseStatus = 500;

  const responseHeaders: Headers = {};

  const fetchFn = request?.fetch ?? fetch;

  try {
    const response = await fetchFn(url, options);

    responseStatus = response.status;

    for (const [field, value] of response.headers) {
      responseHeaders[field] = value;
    }
    const data = await getResponseData(response);

    if (!response.ok) {
      throw new HTTPError(response.statusText, responseStatus, {
        data,
        headers: responseHeaders,
        request: requestOptions,
      });
    }

    return {
      data: data as T,
      headers: responseHeaders,
      status: responseStatus,
      url: response.url || url,
    } as Response<T>;
  } catch (error) {
    if (error instanceof HTTPError) {
      throw error;
    }

    throw new HTTPError(
      error instanceof Error ? error.message : String(error),
      responseStatus,
      {
        headers: responseHeaders,
        request: requestOptions,
      },
    );
  }
}

async function getResponseData(
  response: globalThis.Response,
): Promise<unknown> {
  const contentType = response.headers.get("content-type");

  if (response.status === 204) {
    return undefined;
  }

  if (contentType?.includes("application/json")) {
    return response.json();
  }

  if (
    contentType?.startsWith("text/") ||
    contentType?.includes("charset=utf-8")
  ) {
    return response.text();
  }

  return response.arrayBuffer();
}
