import { parseTemplate as parseUrlTemplate } from "url-template";

import type {
  EndpointDefaults,
  EndpointParams,
  RequestOptions,
} from "./types.js";

import { extractUrlVariableNames } from "../utils/extract-url-variable-names.js";

function addQueryParameters(
  url: string,
  params: Record<string, unknown> = {},
): string {
  const names = Object.keys(params).filter(
    (name) => typeof params[name] !== "undefined",
  );
  if (names.length === 0) return url;

  const separator = url.includes("?") ? "&" : "?";
  const queryString = names
    .flatMap((name) => {
      if (Array.isArray(params[name])) {
        return params[name]
          .map(
            (value: string) => `${name}=${encodeURIComponent(String(value))}`,
          )
          .join("&");
      }
      return `${name}=${encodeURIComponent(String(params[name]))}`;
    })
    .join("&");

  return `${url}${separator}${queryString}`;
}

const contentType = {
  formData: "multipart/form-data",
  json: "application/json; charset=utf-8",
  urlEncoded: "application/x-www-form-urlencoded",
};

export function parse(options: EndpointDefaults): RequestOptions {
  const {
    baseUrl,
    body: _body,
    headers,
    method,
    request,
    url: _url,
    ...params
  } = options;

  const urlVariableNames = extractUrlVariableNames(_url);

  let url = parseUrlTemplate(_url).expand(params);
  if (!/^http/.test(url)) {
    url = `${baseUrl}${url}`;
  }

  const remainingParams: EndpointParams = {};
  for (const [key, value] of Object.entries(params)) {
    if (!urlVariableNames.includes(key)) {
      remainingParams[key] = value;
    }
  }

  let body: unknown;

  if (["DELETE", "GET"].includes(method)) {
    url = addQueryParameters(url, remainingParams);
  } else if (typeof _body !== "undefined") {
    body = _body;
  } else if (Object.keys(remainingParams).length > 0) {
    body = remainingParams;
  }

  if (typeof body !== "undefined") {
    headers["content-type"] = contentType.json;
    body = JSON.stringify(body);
  }

  return {
    body,
    headers,
    method,
    request,
    url,
  };
}
