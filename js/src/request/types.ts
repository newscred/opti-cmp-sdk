import type {
  Endpoint,
  EndpointOptions,
  EndpointParams,
  Headers,
} from "../endpoint/types.js";
import type { EndpointRoute } from "../types/routes.js";

export interface Request {
  <T>(
    route: EndpointRoute | (string & {}),
    params?: EndpointParams,
  ): Promise<Response<T>>;
  <T>(options: EndpointOptions): Promise<Response<T>>;
  defaults(params: EndpointParams): Request;
  endpoint: Endpoint;
}

export type Response<T> = {
  data: T;
  headers: Headers;
  status: number;
  url: string;
};
