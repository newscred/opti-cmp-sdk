import { readFileSync, writeFileSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = dirname(fileURLToPath(import.meta.url));
const SPEC_DIR = join(__dirname, "../..", "specification");
const TYPES_DIR = join(__dirname, "..", "src", "types");

interface Operation {
  operationId?: string;
  responses?: Record<string, OperationResponse>;
}

interface OperationResponse {
  content?: {
    "application/json"?: {
      schema?: Record<string, unknown> & { $ref?: string };
    };
  };
}

type PathMethods = Record<string, Operation>;

function isValidIdentifier(name: string): boolean {
  return /^[a-zA-Z_$][a-zA-Z0-9_$]*$/.test(name);
}

function main(): void {
  console.log("Generating schema exports...");

  const schemasPath = join(SPEC_DIR, "schemas.json");
  const schemas: Record<string, unknown> = JSON.parse(
    readFileSync(schemasPath, "utf-8"),
  );

  const schemaNames = Object.keys(schemas).sort();
  const validNames = schemaNames.filter(isValidIdentifier);
  const skippedNames = schemaNames.filter((name) => !isValidIdentifier(name));

  // Find inline response schemas from paths.json
  const pathsPath = join(SPEC_DIR, "paths.json");
  const paths: Record<string, PathMethods> = JSON.parse(
    readFileSync(pathsPath, "utf-8"),
  );
  const inlineResponseTypes: Array<{
    name: string;
    operationId: string;
    status: string;
  }> = [];

  for (const methods of Object.values(paths)) {
    for (const operation of Object.values(methods)) {
      if (!operation.operationId) continue;

      for (const [status, response] of Object.entries(
        operation.responses || {},
      )) {
        if (!status.startsWith("2")) continue;

        const schema = response?.content?.["application/json"]?.schema;
        if (schema && !("$ref" in schema)) {
          inlineResponseTypes.push({
            name: `${toPascalCase(operation.operationId)}Response`,
            operationId: operation.operationId,
            status,
          });
          break;
        }
      }
    }
  }

  inlineResponseTypes.sort((a, b) => a.name.localeCompare(b.name));

  const lines: string[] = [
    "// Auto-generated - DO NOT EDIT",
    "",
    `import type { components, operations } from "./openapi.js";`,
    "",
  ];

  for (const name of validNames) {
    lines.push(`export type ${name} = components["schemas"]["${name}"];`);
  }

  if (inlineResponseTypes.length > 0) {
    lines.push("");
    lines.push("// Inline response schemas");
    for (const { name, operationId, status } of inlineResponseTypes) {
      lines.push(
        `export type ${name} = operations["${operationId}"]["responses"]["${status}"]["content"]["application/json"];`,
      );
    }
  }

  lines.push("");

  const outputPath = join(TYPES_DIR, "schema.ts");
  writeFileSync(outputPath, lines.join("\n"));
  console.log(`  Written: ${outputPath}`);
  console.log(`  Exported ${validNames.length} schema types`);
  console.log(`  Exported ${inlineResponseTypes.length} inline response types`);
  if (skippedNames.length > 0) {
    console.log(
      `  Skipped ${skippedNames.length} invalid identifiers: ${skippedNames.join(", ")}`,
    );
  }

  // Generate schema-public.ts with only public types
  const publicSuffixes = ["Response", "Request", "Payload", "ListResponseItem"];
  const publicTypes = validNames.filter((name) =>
    publicSuffixes.some((suffix) => name.endsWith(suffix)),
  );
  const inlinePublicTypes = inlineResponseTypes.map((t) => t.name);
  const allPublicTypes = [...publicTypes, ...inlinePublicTypes].sort();

  const publicLines: string[] = ["// Auto-generated - DO NOT EDIT", ""];

  publicLines.push(`export type {`);
  for (const name of allPublicTypes) {
    publicLines.push(`  ${name},`);
  }
  publicLines.push(`} from "./schema.js";`);

  publicLines.push("");

  const publicOutputPath = join(TYPES_DIR, "schema-public.ts");
  writeFileSync(publicOutputPath, publicLines.join("\n"));
  console.log(`  Written: ${publicOutputPath}`);
  console.log(
    `  Exported ${allPublicTypes.length} public schema types (${publicTypes.length} named + ${inlinePublicTypes.length} inline)`,
  );

  console.log("Done!");
}

function toPascalCase(str: string): string {
  return str.charAt(0).toUpperCase() + str.slice(1);
}

main();
