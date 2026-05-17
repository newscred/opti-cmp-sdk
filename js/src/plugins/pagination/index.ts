import type { APIClient } from "../../client/types.js";
import type { Response } from "../../request/types.js";
import type { Pagination } from "../../types/schema.js";

import { getPage } from "./get-page.js";

export default function paginationPlugin(client: APIClient): void {
  client.hasNextPage = <T extends { pagination: Pagination }>(
    response: Response<T>,
  ): boolean => Boolean(response.data.pagination?.next);

  client.hasPreviousPage = <T extends { pagination: Pagination }>(
    response: Response<T>,
  ): boolean => Boolean(response.data.pagination?.previous);

  client.getNextPage = <T extends { pagination: Pagination }>(
    response: Response<T>,
  ) => getPage(client, "next", response);

  client.getPreviousPage = <T extends { pagination: Pagination }>(
    response: Response<T>,
  ) => getPage(client, "previous", response);
}
