import type { RequestMethod } from "../../endpoint/types.js";

export type Route = {
  method: RequestMethod;
  params?: Record<string, string>;
  url: string;
};

export type Routes = Record<string, Record<string, Route>>;
