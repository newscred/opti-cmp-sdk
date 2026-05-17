import { mkdirSync, readFileSync, writeFileSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import openapiTS, { astToString } from "openapi-typescript";

const __dirname = dirname(fileURLToPath(import.meta.url));
const SPEC_DIR = join(__dirname, "../..", "specification");
const TYPES_DIR = join(__dirname, "..", "src", "types");

async function main(): Promise<void> {
  console.log("Generating OpenAPI types...");

  const specPath = join(SPEC_DIR, "openapi.json");
  const spec = JSON.parse(readFileSync(specPath, "utf-8"));

  const ast = await openapiTS(spec, {
    alphabetize: false,
    exportType: true,
  });

  const output = astToString(ast);

  mkdirSync(TYPES_DIR, { recursive: true });
  const outputPath = join(TYPES_DIR, "openapi.d.ts");
  writeFileSync(outputPath, `// Auto-generated - DO NOT EDIT\n\n${output}`);
  console.log(`  Written: ${outputPath}`);

  console.log("Done!");
}

main().catch((err) => {
  console.error("Error:", err);
  process.exit(1);
});
