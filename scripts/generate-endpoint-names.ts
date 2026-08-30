import { existsSync, readFileSync, writeFileSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = dirname(fileURLToPath(import.meta.url));
const SPEC_DIR = join(__dirname, "..", "specification");

interface EndpointNameMapping {
  [path: string]: {
    [method: string]: {
      deprecated?: boolean;
      name: string;
      namespace: string;
    };
  };
}

interface OperationObject {
  deprecated?: boolean;
  operationId?: string;
  summary?: string;
  tags?: string[];
}

interface PathItem {
  delete?: OperationObject;
  get?: OperationObject;
  head?: OperationObject;
  options?: OperationObject;
  patch?: OperationObject;
  post?: OperationObject;
  put?: OperationObject;
}

const HTTP_METHODS = [
  "get",
  "post",
  "put",
  "patch",
  "delete",
  "options",
  "head",
] as const;

type HttpMethod = (typeof HTTP_METHODS)[number];

const TAG_TO_NAMESPACE: Record<string, string> = {
  Assets: "asset",
  "Brand Compliance": "brandCompliance",
  Campaigns: "campaign",
  "Content Graph": "contentGraph",
  Events: "event",
  Fields: "field",
  Labels: "label",
  Library: "library",
  Milestones: "milestone",
  Publishing: "publishing",
  Settings: "settings",
  "Structured Contents": "structuredContent",
  Tasks: "task",
  "Task Step": "taskStep",
  Teams: "team",
  Templates: "template",
  Uploader: "uploader",
  Users: "user",
  Workflows: "workflow",
  "Work Requests": "workRequest",
};

function deriveMethodName(
  method: HttpMethod,
  path: string,
  operation: OperationObject,
): string {
  if (operation.operationId) {
    const parts = operation.operationId.split("_");
    if (parts.length > 1) {
      return toCamelCase(parts.slice(1).join("_"));
    }
    return toCamelCase(operation.operationId);
  }

  const segments = path
    .split("/")
    .filter((s) => s && !s.startsWith("{"))
    .map((s) => s.replace(/-/g, "_"));

  const lastSegment = segments[segments.length - 1] || "resource";
  const hasIdParam =
    path.includes(`{${lastSegment.replace(/s$/, "")}_id}`) ||
    path.endsWith("}");

  const prefixMap: Record<HttpMethod, string> = {
    delete: "delete",
    get: hasIdParam ? "get" : "list",
    head: "head",
    options: "options",
    patch: "update",
    post: "create",
    put: "update",
  };

  const prefix = prefixMap[method];
  const resourceName = toCamelCase(lastSegment);

  if (prefix === "list") {
    return resourceName;
  }
  return `${prefix}${resourceName.charAt(0).toUpperCase()}${resourceName.slice(1)}`;
}

function main(): void {
  console.log("Generating endpoint names...");

  const pathsFile = join(SPEC_DIR, "paths.json");
  if (!existsSync(pathsFile)) {
    throw new Error("paths.json not found. Run generate:spec first.");
  }

  const paths: Record<string, PathItem> = JSON.parse(
    readFileSync(pathsFile, "utf-8"),
  );
  const endpointNames: EndpointNameMapping = {};

  let totalEndpoints = 0;

  for (const [path, pathItem] of Object.entries(paths)) {
    endpointNames[path] = {};

    for (const method of HTTP_METHODS) {
      const operation = pathItem[method];
      if (!operation) continue;

      totalEndpoints++;

      const tag = operation.tags?.[0] || "default";
      const namespace = toNamespace(tag);
      const name = deriveMethodName(method, path, operation);

      endpointNames[path][method] = {
        deprecated: operation.deprecated,
        name,
        namespace,
      };
    }
  }

  const outputPath = join(SPEC_DIR, "endpoint-names.json");
  writeFileSync(outputPath, JSON.stringify(endpointNames, null, 2) + "\n");

  console.log(`  Total endpoints: ${totalEndpoints}`);
  console.log(`  Written: ${outputPath}`);
  console.log("Done!");
}

function toCamelCase(str: string): string {
  return str
    .replace(/[-_\s]+(.)?/g, (_, c) => (c ? c.toUpperCase() : ""))
    .replace(/^./, (c) => c.toLowerCase());
}

function toNamespace(tag: string): string {
  if (TAG_TO_NAMESPACE[tag]) {
    return TAG_TO_NAMESPACE[tag];
  }
  throw new Error("Namespace not found for tag: " + tag);
}

main();
