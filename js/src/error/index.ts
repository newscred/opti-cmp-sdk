export class HTTPError extends Error {
  data?: unknown;
  headers?: Record<string, string>;
  request?: {
    body?: unknown;
    headers: Record<string, string>;
    method: string;
    url: string;
  };
  status: number;

  constructor(
    message: string,
    status: number,
    options: {
      data?: unknown;
      headers?: Record<string, string>;
      request?: {
        body?: unknown;
        headers: Record<string, string>;
        method: string;
        url: string;
      };
    },
  ) {
    super(message);

    // Maintains proper stack trace (only available on V8)
    if (Error.captureStackTrace)
      Error.captureStackTrace(this, this.constructor);

    this.name = "HTTPError";
    this.data = options.data;
    this.headers = options.headers;
    this.request = options.request;
    this.status = status;
  }
}
