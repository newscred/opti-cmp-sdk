// Command gen generates the typed endpoint surface and schema types for the
// opticmp package from the shared OpenAPI specification.
//
// It runs oapi-codegen for the schema types, then emits one file per namespace
// with typed methods, plus namespaces.gen.go wiring them onto the client.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/codegen"
)

const marker = "// Auto-generated - DO NOT EDIT"

var httpMethods = []string{"get", "post", "put", "patch", "delete"}

func main() {
	specDir := flag.String("spec", "../../specification", "directory holding the specification JSON files")
	outDir := flag.String("out", ".", "output directory (the opticmp package)")
	flag.Parse()

	absSpec, err := filepath.Abs(*specDir)
	must(err)
	absOut, err := filepath.Abs(*outDir)
	must(err)

	spec := loadJSON(filepath.Join(absSpec, "openapi.json"))
	names := loadEndpointNames(filepath.Join(absSpec, "endpoint-names.json"))

	// Prime oapi-codegen's initialism normalizer up front, so the name helpers
	// (codegen.ToCamelCaseWithInitialisms) work for the passes below that run
	// before the schema Generate call.
	primeNormalizer()

	// Hoist inline request/response bodies into components.schemas and record
	// per-operation metadata for the endpoint model.
	meta := hoistAndCollect(spec)

	// Name nested inline object schemas so oapi-codegen emits named types
	// instead of anonymous nested structs.
	tagNestedObjects(spec)

	known := componentNames(spec)

	// 1. Schema types via oapi-codegen (fed the hoisted spec).
	generateSchemaTypes(spec, absOut)

	// 2. Endpoint surface.
	nss := buildNamespaces(spec, names, meta, known)
	writeNamespaces(nss, absOut)

	// Remove any previously generated files this run did not write (a renamed
	// or dropped namespace, or the older *_gen.go convention). Done after
	// writing, not before, so `go generate`'s file scan never sees a file
	// vanish mid-run — existing files are overwritten in place.
	written := map[string]bool{
		filepath.Join(absOut, "namespaces.gen.go"):       true,
		filepath.Join(absOut, "schema", "schema.gen.go"): true,
	}
	for _, ns := range nss {
		written[filepath.Join(absOut, ns.file+".gen.go")] = true
	}
	removeOrphans(absOut, written)

	fmt.Printf("Generated %d namespaces into %s\n", len(nss), absOut)
}

// ---------------------------------------------------------------------------
// Spec loading
// ---------------------------------------------------------------------------

func loadJSON(path string) map[string]any {
	data, err := os.ReadFile(path)
	must(err)
	var out map[string]any
	must(json.Unmarshal(data, &out))
	return out
}

// endpointName is one entry of endpoint-names.json.
type endpointName struct {
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	Deprecated bool   `json:"deprecated"`
}

func loadEndpointNames(path string) map[string]map[string]endpointName {
	data, err := os.ReadFile(path)
	must(err)
	var out map[string]map[string]endpointName
	must(json.Unmarshal(data, &out))
	return out
}

func componentNames(spec map[string]any) map[string]bool {
	out := map[string]bool{}
	for name := range schemasMap(spec) {
		out[name] = true
	}
	return out
}

func schemasMap(spec map[string]any) map[string]any {
	comps, _ := spec["components"].(map[string]any)
	if comps == nil {
		comps = map[string]any{}
		spec["components"] = comps
	}
	schemas, _ := comps["schemas"].(map[string]any)
	if schemas == nil {
		schemas = map[string]any{}
		comps["schemas"] = schemas
	}
	return schemas
}

// ---------------------------------------------------------------------------
// Hoisting inline bodies + operation metadata
// ---------------------------------------------------------------------------

type opMeta struct {
	responseType    string // "" => untyped (any)
	requestType     string // "" => no body
	requestRequired bool
	paginated       bool
}

// metaKey identifies an operation by path and method.
func metaKey(path, method string) string { return method + " " + path }

func hoistAndCollect(spec map[string]any) map[string]opMeta {
	meta := map[string]opMeta{}
	schemas := schemasMap(spec)
	paths, _ := spec["paths"].(map[string]any)

	for path, pv := range paths {
		item, _ := pv.(map[string]any)
		for _, method := range httpMethods {
			op, _ := item[method].(map[string]any)
			if op == nil {
				continue
			}
			opID, _ := op["operationId"].(string)
			base := pascal(opID)

			m := opMeta{}
			if schema, _ := responseSchema(op); schema != nil {
				name := refName(schema)
				if name == "" { // inline: hoist it
					name = base + "Response"
					schemas[name] = schema
					m.paginated = isPaginated(schema)
				}
				m.responseType = name
			}
			if schema, required := requestSchema(op); schema != nil {
				name := refName(schema)
				if name == "" {
					name = base + "Request"
					schemas[name] = schema
				}
				m.requestType = name
				m.requestRequired = required
			}
			meta[metaKey(path, method)] = m
		}
	}
	return meta
}

// responseSchema returns the first 2xx application/json schema.
func responseSchema(op map[string]any) (map[string]any, bool) {
	responses, _ := op["responses"].(map[string]any)
	for _, code := range []string{"200", "201", "202"} {
		resp, _ := responses[code].(map[string]any)
		if resp == nil {
			continue
		}
		if schema := jsonSchema(resp["content"]); schema != nil {
			return schema, true
		}
	}
	return nil, false
}

func requestSchema(op map[string]any) (map[string]any, bool) {
	rb, _ := op["requestBody"].(map[string]any)
	if rb == nil {
		return nil, false
	}
	content, _ := rb["content"].(map[string]any)
	schema := pickSchema(content, "application/json", "multipart/form-data")
	if schema == nil {
		return nil, false
	}
	required, _ := rb["required"].(bool)
	return schema, required
}

func jsonSchema(content any) map[string]any {
	return pickSchema(content, "application/json")
}

func pickSchema(content any, mediaTypes ...string) map[string]any {
	c, _ := content.(map[string]any)
	if c == nil {
		return nil
	}
	for _, mt := range mediaTypes {
		media, _ := c[mt].(map[string]any)
		if media == nil {
			continue
		}
		if schema, _ := media["schema"].(map[string]any); schema != nil {
			return schema
		}
	}
	return nil
}

func isPaginated(schema map[string]any) bool {
	props, _ := schema["properties"].(map[string]any)
	if props == nil {
		return false
	}
	data, _ := props["data"].(map[string]any)
	_, hasPagination := props["pagination"]
	return data != nil && data["type"] == "array" && hasPagination
}

// tagNestedObjects gives every nested inline object schema a name via the
// x-go-type-name extension, so oapi-codegen emits a named type (schema.FooBar)
// instead of an anonymous nested struct. Names follow oapi's own nested
// convention: the enclosing type name concatenated with the Pascal-cased
// property path (matching how it already names nested enums), run through the
// same initialism normalizer so the casing matches the surrounding types.
func tagNestedObjects(spec map[string]any) {
	schemas := schemasMap(spec)
	// Reserve the emitted type names of the components so a nested name never
	// collides with a top-level type.
	used := map[string]bool{}
	for name := range schemas {
		used[codegen.ToCamelCaseWithInitialisms(name)] = true
	}
	for _, name := range sortedKeys(schemas) {
		if sc, ok := schemas[name].(map[string]any); ok {
			tagDescendants(sc, codegen.ToCamelCaseWithInitialisms(name), used)
		}
	}
}

// tagDescendants walks the inline object schemas reachable from sc through its
// properties, array items, additionalProperties, and allOf members, tagging
// each. prefix is the enclosing type name.
func tagDescendants(sc map[string]any, prefix string, used map[string]bool) {
	if props, ok := sc["properties"].(map[string]any); ok {
		for _, prop := range sortedKeys(props) {
			tagSlot(props, prop, prefix+pascal(prop), used)
		}
	}
	if sc["type"] == "array" {
		tagSlot(sc, "items", prefix+"Item", used)
	}
	if _, ok := sc["additionalProperties"].(map[string]any); ok {
		tagSlot(sc, "additionalProperties", prefix+"Value", used)
	}
	if members, ok := sc["allOf"].([]any); ok {
		for _, m := range members {
			if mm, ok := m.(map[string]any); ok {
				// allOf members merge into the enclosing type, so tag their
				// descendants under the same prefix.
				tagDescendants(mm, prefix, used)
			}
		}
	}
}

// tagSlot processes the schema at container[key]. An inline object is tagged
// with a unique x-go-type-name derived from candidate; an array recurses into
// its items (which take the property-derived name). $refs are left untouched.
func tagSlot(container map[string]any, key, candidate string, used map[string]bool) {
	s, ok := container[key].(map[string]any)
	if !ok {
		return
	}
	if _, isRef := s["$ref"]; isRef {
		return
	}
	switch {
	case isInlineObject(s):
		name := uniqueName(candidate, used)
		s["x-go-type-name"] = name
		tagDescendants(s, name, used)
	case s["type"] == "array":
		tagSlot(s, "items", candidate, used)
	}
}

// isInlineObject reports whether s is an inline object schema (not a $ref) with
// at least one property, worth naming.
func isInlineObject(s map[string]any) bool {
	if _, ok := s["$ref"]; ok {
		return false
	}
	props, ok := s["properties"].(map[string]any)
	return ok && len(props) > 0
}

// uniqueName normalizes rawBase to a Go type name and makes it unique against
// used, reserving the result.
func uniqueName(rawBase string, used map[string]bool) string {
	base := codegen.ToCamelCaseWithInitialisms(rawBase)
	name := base
	for i := 2; used[name]; i++ {
		name = base + strconv.Itoa(i)
	}
	used[name] = true
	return name
}

func sortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// stripKey recursively deletes every occurrence of key from the spec.
func stripKey(node any, key string) {
	switch v := node.(type) {
	case map[string]any:
		delete(v, key)
		for _, child := range v {
			stripKey(child, key)
		}
	case []any:
		for _, child := range v {
			stripKey(child, key)
		}
	}
}

func refName(schema map[string]any) string {
	ref, _ := schema["$ref"].(string)
	if ref == "" {
		return ""
	}
	parts := strings.Split(ref, "/")
	return parts[len(parts)-1]
}

// ---------------------------------------------------------------------------
// oapi-codegen
// ---------------------------------------------------------------------------

// schemaPkg is the import path of the generated schema subpackage.
const schemaPkg = "github.com/newscred/opti-cmp-sdk/go/schema"

// primeNormalizer runs a throwaway Generate so oapi-codegen initializes its
// package-global initialism map. After this, codegen.ToCamelCaseWithInitialisms
// applies initialisms; without it the map is nil and it silently skips them.
func primeNormalizer() {
	doc, err := openapi3.NewLoader().LoadFromData(
		[]byte(`{"openapi":"3.0.0","info":{"title":"x","version":"0"},"paths":{}}`))
	must(err)
	_, err = codegen.Generate(doc, codegen.Configuration{
		PackageName:   "schema",
		Generate:      codegen.GenerateOptions{Models: true},
		OutputOptions: codegen.OutputOptions{NameNormalizer: string(codegen.NameNormalizerFunctionToCamelCaseWithInitialisms)},
	})
	must(err)
	if codegen.ToCamelCaseWithInitialisms("id") != "ID" {
		log.Fatal("oapi-codegen initialisms not initialized")
	}
}

func generateSchemaTypes(spec map[string]any, outDir string) {
	// Feed oapi-codegen the components only, without paths. It then emits just
	// the schema types and none of the per-operation Params types, which this
	// SDK does not use (params are generated in package opticmp). The hoisted
	// inline body schemas live in components, so they are still emitted, and
	// skip-prune keeps every component.
	tmpSpec := make(map[string]any, len(spec))
	for key, value := range spec {
		if key == "paths" {
			continue
		}
		tmpSpec[key] = value
	}
	// oapi-codegen emits oneOf discriminator helpers and switch cases in
	// Go-map order, which is not stable across runs. Dropping the discriminator
	// makes it emit the union helpers in the ordered oneOf member sequence, so
	// the output is deterministic. Union types keep their AsX/FromX helpers.
	stripKey(tmpSpec, "discriminator")

	data, err := json.Marshal(tmpSpec)
	must(err)
	doc, err := openapi3.NewLoader().LoadFromData(data)
	must(err)

	// name-normalizer upper-cases Go initialisms (Id -> ID, Url -> URL). Calling
	// Generate in-process also primes oapi-codegen's initialism state, so
	// toCamelInit (codegen.ToCamelCaseWithInitialisms) matches these names.
	code, err := codegen.Generate(doc, codegen.Configuration{
		PackageName: "schema",
		Generate:    codegen.GenerateOptions{Models: true},
		OutputOptions: codegen.OutputOptions{
			SkipPrune:      true, // keep the hoisted inline body types
			NameNormalizer: string(codegen.NameNormalizerFunctionToCamelCaseWithInitialisms),
		},
	})
	must(err)

	schemaDir := filepath.Join(outDir, "schema")
	must(os.MkdirAll(schemaDir, 0o755))
	out := filepath.Join(schemaDir, "schema.gen.go")
	must(os.WriteFile(out, []byte(code), 0o644))
	normalizeHeader(out)
}

// normalizeHeader replaces oapi-codegen's package doc comment with the standard
// marker so there is a single package comment.
func normalizeHeader(path string) {
	data, err := os.ReadFile(path)
	must(err)
	text := string(data)
	idx := strings.Index(text, "\npackage schema")
	if idx == -1 {
		return
	}
	rest := text[idx+1:]
	must(os.WriteFile(path, []byte(marker+"\n\n"+rest), 0o644))
}

// removeOrphans deletes generated-pattern files that this run did not write.
// Hand-written files (client.go, *_test.go, ...) never match the patterns, so
// they are left alone. In steady state nothing is removed.
func removeOrphans(outDir string, written map[string]bool) {
	patterns := []string{
		filepath.Join(outDir, "*.gen.go"),
		filepath.Join(outDir, "*_gen.go"), // older convention
		filepath.Join(outDir, "schema", "*.gen.go"),
		filepath.Join(outDir, "schema", "*_gen.go"),
	}
	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		must(err)
		for _, path := range matches {
			if !written[path] {
				must(os.Remove(path))
			}
		}
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ---------------------------------------------------------------------------
// Model
// ---------------------------------------------------------------------------

type param struct {
	wire     string
	field    string
	goType   string
	isArray  bool
	required bool
	in       string // path | query | header
	doc      string
}

type endpoint struct {
	goName       string
	httpMethod   string // GET, POST, ...
	path         string
	params       []param
	hasBody      bool
	bodyType     string
	bodyRequired bool
	returnType   string // "" => any
	deprecated   bool
	summary      string
}

func (e endpoint) paramsType() string { return e.goName + "Params" }

func (e endpoint) hasParamsArg() bool { return len(e.params) > 0 || e.hasBody }

type namespace struct {
	field     string
	service   string
	file      string
	endpoints []endpoint
}

func buildNamespaces(spec map[string]any, names map[string]map[string]endpointName, meta map[string]opMeta, known map[string]bool) []namespace {
	paths, _ := spec["paths"].(map[string]any)
	byField := map[string]*namespace{}

	for path, methods := range names {
		item, _ := paths[path].(map[string]any)
		if item == nil {
			continue
		}
		pathParams := paramList(item["parameters"])

		for _, method := range httpMethods {
			mapping, ok := methods[method]
			if !ok {
				continue
			}
			op, _ := item[method].(map[string]any)
			if op == nil {
				continue
			}

			ep := endpoint{
				goName:     codegen.ToCamelCaseWithInitialisms(mapping.Name),
				httpMethod: strings.ToUpper(method),
				path:       path,
				deprecated: mapping.Deprecated,
			}
			ep.summary, _ = op["summary"].(string)

			m := meta[metaKey(path, method)]
			if m.responseType != "" && known[m.responseType] {
				ep.returnType = "schema." + codegen.ToCamelCaseWithInitialisms(m.responseType)
			}
			if m.requestType != "" {
				ep.hasBody = true
				ep.bodyType = goRefType(m.requestType, known)
				ep.bodyRequired = m.requestRequired
			}

			ep.params = mergeParams(pathParams, paramList(op["parameters"]))
			if m.paginated {
				ep.params = withPaginationParams(ep.params)
			}

			field := codegen.ToCamelCaseWithInitialisms(mapping.Namespace)
			ns := byField[field]
			if ns == nil {
				ns = &namespace{
					field:   field,
					service: field + "Service",
					file:    snake(mapping.Namespace),
				}
				byField[field] = ns
			}
			ns.endpoints = append(ns.endpoints, ep)
		}
	}

	out := make([]namespace, 0, len(byField))
	for _, ns := range byField {
		sort.Slice(ns.endpoints, func(i, j int) bool { return ns.endpoints[i].goName < ns.endpoints[j].goName })
		out = append(out, *ns)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].field < out[j].field })
	return out
}

func paramList(v any) []map[string]any {
	arr, _ := v.([]any)
	out := make([]map[string]any, 0, len(arr))
	for _, it := range arr {
		if m, ok := it.(map[string]any); ok {
			out = append(out, m)
		}
	}
	return out
}

func mergeParams(pathParams, opParams []map[string]any) []param {
	var out []param
	seen := map[string]bool{}
	for _, raw := range append(append([]map[string]any{}, pathParams...), opParams...) {
		in, _ := raw["in"].(string)
		if in == "cookie" {
			continue
		}
		name, _ := raw["name"].(string)
		if name == "" || seen[name] {
			continue
		}
		seen[name] = true

		schema, _ := raw["schema"].(map[string]any)
		goType, isArray := paramGoType(schema)
		required, _ := raw["required"].(bool)
		if in == "path" {
			required = true
		}
		doc, _ := raw["description"].(string)
		out = append(out, param{
			wire:     name,
			field:    codegen.ToCamelCaseWithInitialisms(name),
			goType:   goType,
			isArray:  isArray,
			required: required,
			in:       in,
			doc:      doc,
		})
	}
	return out
}

func withPaginationParams(params []param) []param {
	has := map[string]bool{}
	for _, p := range params {
		has[p.wire] = true
	}
	for _, name := range []string{"offset", "page_size"} {
		if !has[name] {
			params = append(params, param{
				wire:   name,
				field:  codegen.ToCamelCaseWithInitialisms(name),
				goType: "int",
				in:     "query",
			})
		}
	}
	return params
}

func paramGoType(schema map[string]any) (string, bool) {
	if schema == nil {
		return "string", false
	}
	if name := refName(schema); name != "" {
		return "schema." + codegen.ToCamelCaseWithInitialisms(name), false
	}
	switch t, _ := schema["type"].(string); t {
	case "array":
		items, _ := schema["items"].(map[string]any)
		et, _ := paramGoType(items)
		return "[]" + et, true
	case "integer":
		return "int", false
	case "number":
		return "float64", false
	case "boolean":
		return "bool", false
	default:
		return "string", false
	}
}

func goRefType(name string, known map[string]bool) string {
	if known[name] {
		return "schema." + codegen.ToCamelCaseWithInitialisms(name)
	}
	return "any"
}
