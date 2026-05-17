import { describe, expect, it } from "vitest";

import { extractUrlVariableNames } from "./extract-url-variable-names.js";

describe("extractUrlVariableNames", () => {
  it("should return empty array for URL with no variables", () => {
    expect(extractUrlVariableNames("/users")).toEqual([]);
    expect(extractUrlVariableNames("/api/v1/items")).toEqual([]);
    expect(extractUrlVariableNames("")).toEqual([]);
  });

  it("should extract single variable", () => {
    expect(extractUrlVariableNames("/users/{id}")).toEqual(["id"]);
    expect(extractUrlVariableNames("/{resource}")).toEqual(["resource"]);
  });

  it("should extract multiple variables", () => {
    expect(extractUrlVariableNames("/users/{userId}/posts/{postId}")).toEqual([
      "userId",
      "postId",
    ]);
  });

  it("should handle variables with operators", () => {
    expect(extractUrlVariableNames("/search{?q,limit}")).toEqual([
      "q",
      "limit",
    ]);
    expect(extractUrlVariableNames("/files{/path*}")).toEqual(["path"]);
  });

  it("should handle adjacent variables", () => {
    expect(extractUrlVariableNames("/{a}{b}")).toEqual(["a", "b"]);
  });

  it("should handle variables at different positions", () => {
    expect(extractUrlVariableNames("{base}/users/{id}")).toEqual([
      "base",
      "id",
    ]);
  });
});
