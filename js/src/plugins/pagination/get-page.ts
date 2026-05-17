import type { APIClient } from "../../client/types.js";
import type { Response } from "../../request/types.js";
import type { Pagination } from "../../types/schema.js";
import type { Direction } from "./types.js";

import { HTTPError } from "../../error/index.js";

export async function getPage<T extends { pagination: Pagination }>(
  client: APIClient,
  direction: Direction,
  response: Response<T>,
): Promise<Response<T>> {
  const url = response.data.pagination?.[direction];

  if (!url) {
    throw new HTTPError(`No ${direction} page available`, 404, {});
  }

  return client.request<T>(url);
}
