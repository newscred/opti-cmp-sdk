import Mustache from "mustache";
import { mkdirSync, readFileSync, writeFileSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = dirname(fileURLToPath(import.meta.url));
const SPEC_DIR = join(__dirname, "../..", "specification");
const SRC_DIR = join(__dirname, "..", "src");
const TYPES_DIR = join(SRC_DIR, "types");
const TEMPLATES_DIR = join(__dirname, "..", "templates");

interface EndpointMapping {
  deprecated?: boolean;
  name: string;
  namespace: string;
}

interface OperationObject {
  deprecated?: boolean;
  parameters?: ParameterObject[];
  requestBody?: RequestBodyObject;
  responses?: Record<string, ResponseObject>;
  summary?: string;
  tags?: string[];
}

interface ParameterObject {
  description?: string;
  in: "cookie" | "header" | "path" | "query";
  name: string;
  required?: boolean;
  schema?: SchemaObject;
}

interface PathItem {
  delete?: OperationObject;
  get?: OperationObject;
  parameters?: ParameterObject[];
  patch?: OperationObject;
  post?: OperationObject;
  put?: OperationObject;
}

interface RequestBodyObject {
  content?: {
    [mediaType: string]: {
      schema?: SchemaObject;
    };
  };
  required?: boolean;
}

interface ResponseObject {
  content?: {
    [mediaType: string]: {
      schema?: SchemaObject;
    };
  };
}

interface SchemaObject {
  $ref?: string;
  allOf?: SchemaObject[];
  anyOf?: SchemaObject[];
  enum?: string[];
  format?: string;
  items?: SchemaObject;
  nullable?: boolean;
  oneOf?: SchemaObject[];
  properties?: Record<string, SchemaObject>;
  required?: string[];
  type?: string;
}

interface SlimRoute {
  method: string;
  params: Record<string, string>;
  url: string;
}

interface TemplateData {
  namespaces: TemplateNamespace[];
}

interface TemplateEndpoint {
  bodyRequired?: boolean;
  bodyType?: string;
  deprecated?: boolean;
  hasBody: boolean;
  isEmpty: boolean;
  name: string;
  params: TemplateParam[];
  paramsType: string;
  returnType: string;
}

interface TemplateNamespace {
  endpoints: TemplateEndpoint[];
  namespace: string;
}

interface TemplateParam {
  description?: string;
  in: "header" | "path" | "query";
  name: string;
  required: boolean;
  type: string;
}

function toPascalCase(str: string): string {
  return str
    .replace(/[-_\s]+(.)?/g, (_, c) => (c ? c.toUpperCase() : ""))
    .replace(/^./, (c) => c.toUpperCase());
}

const HTTP_METHODS = ["get", "post", "put", "patch", "delete"] as const;

function getRequestBodyType(
  requestBody: RequestBodyObject | undefined,
  schemaNames: Set<string>,
  endpointName: string,
): undefined | { bodyRequired: boolean; bodyType: string } {
  if (!requestBody?.content) return undefined;

  const jsonContent = requestBody.content["application/json"];
  const formContent = requestBody.content["multipart/form-data"];
  const content = jsonContent || formContent;

  if (!content?.schema) return undefined;

  const schema = content.schema;

  // Check for inline object schema (will be named {EndpointName}Request)
  if (schema.type === "object" && schema.properties && !schema.$ref) {
    const inlineName = `${toPascalCase(endpointName)}Request`;
    if (schemaNames.has(inlineName)) {
      return {
        bodyRequired: requestBody.required ?? false,
        bodyType: `Schema.${inlineName}`,
      };
    }
  }

  return {
    bodyRequired: requestBody.required ?? false,
    bodyType: schemaToType(schema, schemaNames),
  };
}

function getReturnType(
  responses: Record<string, ResponseObject> | undefined,
  schemaNames: Set<string>,
  endpointName: string,
): { hasPagination?: boolean; returnType: string } {
  if (!responses) return { returnType: "void" };

  for (const code of ["200", "201", "202", "204"]) {
    const response = responses[code];
    if (!response?.content) continue;

    const jsonContent = response.content["application/json"];
    if (!jsonContent?.schema) continue;

    const schema = jsonContent.schema;
    let hasPagination = false;

    if (
      schema.properties?.data?.type === "array" &&
      schema.properties?.pagination
    ) {
      hasPagination = true;
    }

    let inlineSchemaName = "";
    if (
      schema.type === "object" &&
      schema.properties &&
      !schema.$ref &&
      endpointName
    ) {
      const inlineName = `${toPascalCase(endpointName)}Response`;
      inlineSchemaName = `Schema.${inlineName}`;
    }
    return {
      hasPagination,
      returnType: `Response<${schemaToType(schema, schemaNames, inlineSchemaName)}>`,
    };
  }

  return { returnType: "void" };
}

function main(): void {
  console.log("Generating endpoints...");

  const pathsFile = join(SPEC_DIR, "paths.json");
  const endpointNamesFile = join(SPEC_DIR, "endpoint-names.json");
  const schemasFile = join(SPEC_DIR, "schemas.json");

  const paths: Record<string, PathItem> = JSON.parse(
    readFileSync(pathsFile, "utf-8"),
  );
  const endpointNames: Record<
    string,
    Record<string, EndpointMapping>
  > = JSON.parse(readFileSync(endpointNamesFile, "utf-8"));
  const schemas: Record<string, unknown> = JSON.parse(
    readFileSync(schemasFile, "utf-8"),
  );
  const schemaNames = new Set(Object.keys(schemas));

  const namespaceMap = new Map<string, TemplateEndpoint[]>();
  const slimRoutes: Record<string, Record<string, SlimRoute>> = {};
  let totalRoutes = 0;

  for (const [path, pathItem] of Object.entries(paths)) {
    const pathEndpoints = endpointNames[path];
    if (!pathEndpoints) continue;

    for (const method of HTTP_METHODS) {
      const operation = pathItem[method];
      const mapping = pathEndpoints[method];
      if (!operation || !mapping) continue;

      const { name, namespace } = mapping;

      // Process for templates
      const params = processParameters(
        pathItem.parameters,
        operation.parameters,
        schemaNames,
      );
      const bodyInfo = getRequestBodyType(
        operation.requestBody,
        schemaNames,
        name,
      );
      const { hasPagination, returnType } = getReturnType(
        operation.responses,
        schemaNames,
        name,
      );
      if (hasPagination) {
        if (!params.find((p) => p.name === "offset")) {
          params.push({
            in: "query",
            name: "offset",
            required: false,
            type: "number",
          });
        }
        if (!params.find((p) => p.name === "page_size")) {
          params.push({
            in: "query",
            name: "page_size",
            required: false,
            type: "number",
          });
        }
      }

      const hasBody = !!bodyInfo;
      const isEmpty = params.length === 0 && !hasBody;

      const endpoint: TemplateEndpoint = {
        bodyRequired: bodyInfo?.bodyRequired,
        bodyType: bodyInfo?.bodyType,
        deprecated: operation.deprecated,
        hasBody,
        isEmpty,
        name,
        params,
        paramsType: `${toPascalCase(name)}Params`,
        returnType,
      };

      if (!namespaceMap.has(namespace)) {
        namespaceMap.set(namespace, []);
      }
      namespaceMap.get(namespace)!.push(endpoint);

      // Build slim routes
      if (!slimRoutes[namespace]) {
        slimRoutes[namespace] = {};
      }
      slimRoutes[namespace][name] = {
        method: method.toUpperCase(),
        params: Object.fromEntries(
          params.filter((p) => p.in !== "path").map((p) => [p.name, p.in]),
        ),
        url: path,
      };

      totalRoutes++;
    }
  }

  // Sort namespaces and endpoints
  const sortedNamespaces = Array.from(namespaceMap.keys()).sort((a, b) => {
    return a.toLowerCase() < b.toLowerCase() ? -1 : 1;
  });
  const templateData: TemplateData = {
    namespaces: sortedNamespaces.map((ns) => ({
      endpoints: namespaceMap
        .get(ns)!
        .sort((a, b) => a.name.localeCompare(b.name)),
      namespace: ns,
    })),
  };

  mkdirSync(TYPES_DIR, { recursive: true });

  // Generate params.ts
  console.log("  Generating params types...");
  const paramsTemplate = readFileSync(
    join(TEMPLATES_DIR, "params.mustache"),
    "utf-8",
  );
  const paramsTypes = Mustache.render(paramsTemplate, templateData);
  writeFileSync(join(TYPES_DIR, "params.ts"), paramsTypes);
  console.log(`    Written: ${join(TYPES_DIR, "params.ts")}`);

  // Generate endpoints.ts
  console.log("  Generating endpoint types...");
  const endpointsTemplate = readFileSync(
    join(TEMPLATES_DIR, "endpoints.mustache"),
    "utf-8",
  );
  const endpointTypes = Mustache.render(endpointsTemplate, templateData);
  writeFileSync(join(TYPES_DIR, "endpoints.ts"), endpointTypes);
  console.log(`    Written: ${join(TYPES_DIR, "endpoints.ts")}`);

  // Generate slim routes.json
  console.log("  Generating routes.json...");
  const sortedSlimRoutes: Record<string, Record<string, SlimRoute>> = {};
  for (const ns of sortedNamespaces) {
    sortedSlimRoutes[ns] = {};
    for (const name of Object.keys(slimRoutes[ns] || {}).sort()) {
      sortedSlimRoutes[ns][name] = slimRoutes[ns][name];
    }
  }
  writeFileSync(
    join(SRC_DIR, "plugins/register-api-endpoints/routes.json"),
    JSON.stringify(sortedSlimRoutes) + "\n",
  );
  console.log(`    Written: ${join(SRC_DIR, "routes.json")}`);

  console.log(`  Namespaces: ${sortedNamespaces.length}`);
  console.log(`  Total routes: ${totalRoutes}`);
  console.log("Done!");
}

function processParameters(
  pathParams: ParameterObject[] | undefined,
  opParams: ParameterObject[] | undefined,
  schemaNames: Set<string>,
): TemplateParam[] {
  const params: TemplateParam[] = [];
  const allParams = [...(pathParams || []), ...(opParams || [])];
  const seen = new Set<string>();

  for (const param of allParams) {
    if (param.in === "cookie") continue;
    if (!param.name || seen.has(param.name)) continue;
    seen.add(param.name);

    let type: string;
    if (param.schema?.enum) {
      type = param.schema.enum.map((e) => `'${e}'`).join(" | ");
    } else {
      type = schemaToType(param.schema, schemaNames);
    }

    params.push({
      description: param.description,
      in: param.in as "header" | "path" | "query",
      name: param.name,
      required: param.required ?? param.in === "path",
      type,
    });
  }

  return params;
}

function resolveRef(ref: string): string {
  const parts = ref.split("/");
  return parts[parts.length - 1];
}

function schemaToType(
  schema: SchemaObject | undefined,
  schemaNames: Set<string>,
  inlineSchemaName?: string,
): string {
  if (!schema) return "unknown";

  if (schema.$ref) {
    const refName = resolveRef(schema.$ref);
    return schemaNames.has(refName) ? `Schema.${refName}` : refName;
  }

  if (schema.allOf || schema.oneOf || schema.anyOf) {
    const schemas = schema.allOf || schema.oneOf || schema.anyOf || [];
    const types = schemas.map((s) => schemaToType(s, schemaNames));
    return types.join(" | ") || "unknown";
  }

  switch (schema.type) {
    case "array":
      return `${schemaToType(schema.items, schemaNames)}[]`;
    case "boolean":
      return "boolean";
    case "integer":
    case "number":
      return "number";
    case "object":
      return inlineSchemaName || "Record<string, unknown>";
    case "string":
      if (schema.enum) return schema.enum.map((e) => `'${e}'`).join(" | ");
      return "string";
    default:
      return "unknown";
  }
}

main();
