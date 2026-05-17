import { existsSync, mkdirSync, readFileSync, writeFileSync } from "node:fs";
import path, { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import { parse as parseYaml } from "yaml";

const __dirname = dirname(fileURLToPath(import.meta.url));
const SPEC_DIR = join(__dirname, "..", "specification");
const OPENAPI_SPEC_URI =
  process.env.OPENAPI_SPEC_URI ||
  "https://docs.developers.optimizely.com/content-marketing-platform/openapi/optimizely-cmp-open-api-documentation.json";

interface OpenAPISpec {
  components: {
    schemas: Record<string, unknown>;
    securitySchemes?: Record<string, unknown>;
  };
  info: { title: string; version: string };
  openapi: string;
  paths: Record<string, unknown>;
  security?: Array<Record<string, string[]>>;
  servers: Array<{ description?: string; url: string }>;
  tags?: Array<{ description?: string; name: string }>;
}

async function loadSpec(): Promise<OpenAPISpec> {
  let content: string;

  if (OPENAPI_SPEC_URI.startsWith("http")) {
    console.log("Fetching OpenAPI specification from remote...");
    console.log(`  URL: ${OPENAPI_SPEC_URI}`);

    const response = await fetch(OPENAPI_SPEC_URI);
    if (!response.ok) {
      throw new Error(
        `Failed to fetch spec: ${response.status} ${response.statusText}`,
      );
    }

    content = await response.text();
  } else if (existsSync(OPENAPI_SPEC_URI)) {
    console.log("Loading OpenAPI specification from local file...");
    console.log(`  Path: ${OPENAPI_SPEC_URI}`);

    content = readFileSync(OPENAPI_SPEC_URI, "utf-8");
  } else {
    throw new Error(
      "Invalid OPENAPI_SPEC_URI. Must be a valid URL or existing file path.",
    );
  }

  switch (path.extname(OPENAPI_SPEC_URI.toLowerCase())) {
    case ".json":
      return JSON.parse(content) as OpenAPISpec;
    case ".yaml":
    case ".yml":
      return parseYaml(content) as OpenAPISpec;
    default:
      throw new Error("unsupported format");
  }
}

async function main(): Promise<void> {
  const spec = await loadSpec();

  console.log(`\nSpec info:`);
  console.log(`  Title: ${spec.info.title}`);
  console.log(`  Version: ${spec.info.version}`);
  console.log(`  OpenAPI: ${spec.openapi}`);
  console.log(`  Paths: ${Object.keys(spec.paths).length}`);
  console.log(
    `  Schemas: ${Object.keys(spec.components?.schemas || {}).length}`,
  );

  console.log("\nWriting specification files...");

  // @ts-expect-error ...
  writeJson(join(SPEC_DIR, "openapi.json"), sortObjectKeys(spec));
  writeJson(join(SPEC_DIR, "paths.json"), sortObjectKeys(spec.paths));
  writeJson(
    join(SPEC_DIR, "schemas.json"),
    sortObjectKeys(spec.components?.schemas || {}),
  );
  writeJson(join(SPEC_DIR, "tags.json"), spec.tags || []);
  writeJson(join(SPEC_DIR, "security.json"), {
    security: spec.security || [],
    securitySchemes: spec.components?.securitySchemes || {},
  });
  writeJson(join(SPEC_DIR, "servers.json"), spec.servers || []);

  console.log("\nDone!");
}

function sortObjectKeys<T extends Record<string, unknown>>(obj: T): T {
  if (typeof obj !== "object" || obj === null || Array.isArray(obj)) {
    return obj;
  }
  const sorted = {} as T;
  for (const key of Object.keys(obj).sort()) {
    const value = obj[key];
    sorted[key as keyof T] = (
      typeof value === "object" && value !== null
        ? Array.isArray(value)
          ? value.map((v) => (typeof v === "object" ? sortObjectKeys(v) : v))
          : sortObjectKeys(value as Record<string, unknown>)
        : value
    ) as T[keyof T];
  }
  return sorted;
}

function writeJson(filePath: string, data: unknown): void {
  mkdirSync(dirname(filePath), { recursive: true });
  writeFileSync(filePath, JSON.stringify(data, null, 2) + "\n");
  console.log(`  Written: ${filePath}`);
}

main().catch((err) => {
  console.error("Error:", err);
  process.exit(1);
});
