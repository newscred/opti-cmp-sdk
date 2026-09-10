package main

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// ---------------------------------------------------------------------------
// Name helpers
// ---------------------------------------------------------------------------

// pascal converts a camelCase, snake_case or kebab-case name to PascalCase,
// preserving existing internal capitals (listAssets -> ListAssets).
func pascal(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == ' ' || r == '.'
	})
	var b strings.Builder
	for _, p := range parts {
		r := []rune(p)
		if len(r) == 0 {
			continue
		}
		b.WriteString(strings.ToUpper(string(r[0])))
		b.WriteString(string(r[1:]))
	}
	return b.String()
}

// snake converts a camelCase name to snake_case (for file names).
func snake(s string) string {
	var b strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				b.WriteRune('_')
			}
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// ---------------------------------------------------------------------------
// Emission
// ---------------------------------------------------------------------------

func writeNamespaces(nss []namespace, outDir string) {
	for _, ns := range nss {
		src := emitNamespaceFile(ns)
		writeGo(filepath.Join(outDir, ns.file+".gen.go"), src)
	}
	writeGo(filepath.Join(outDir, "namespaces.gen.go"), emitNamespacesFile(nss))
}

func emitNamespacesFile(nss []namespace) string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s\n\npackage opticmp\n\n", marker)
	b.WriteString("// namespaces holds the endpoint namespace fields, embedded into Client.\n")
	b.WriteString("type namespaces struct {\n")
	for _, ns := range nss {
		fmt.Fprintf(&b, "\t%s *%s\n", ns.field, ns.service)
	}
	b.WriteString("}\n\n")
	b.WriteString("func (c *Client) initNamespaces() {\n")
	for _, ns := range nss {
		fmt.Fprintf(&b, "\tc.%s = &%s{client: c}\n", ns.field, ns.service)
	}
	b.WriteString("}\n")
	return b.String()
}

func emitNamespaceFile(ns namespace) string {
	imports := neededImports(ns)

	var b strings.Builder
	fmt.Fprintf(&b, "%s\n\npackage opticmp\n\n", marker)
	if len(imports) > 0 {
		b.WriteString("import (\n")
		for _, imp := range imports {
			fmt.Fprintf(&b, "\t%q\n", imp)
		}
		b.WriteString(")\n\n")
	}

	fmt.Fprintf(&b, "type %s struct{ client *Client }\n\n", ns.service)

	for _, ep := range ns.endpoints {
		emitParamsType(&b, ep)
		emitMethod(&b, ns, ep)
	}
	return b.String()
}

func neededImports(ns namespace) []string {
	var fmtNeeded, urlNeeded, httpNeeded, schemaNeeded bool
	for _, ep := range ns.endpoints {
		if strings.HasPrefix(ep.returnType, "schema.") || strings.HasPrefix(ep.bodyType, "schema.") {
			schemaNeeded = true
		}
		for _, p := range ep.params {
			fmtNeeded = true
			if strings.HasPrefix(p.goType, "schema.") {
				schemaNeeded = true
			}
			switch p.in {
			case "query":
				urlNeeded = true
			case "header":
				httpNeeded = true
			}
		}
	}
	imports := []string{"context"}
	if fmtNeeded {
		imports = append(imports, "fmt")
	}
	if httpNeeded {
		imports = append(imports, "net/http")
	}
	if urlNeeded {
		imports = append(imports, "net/url")
	}
	if schemaNeeded {
		imports = append(imports, schemaPkg)
	}
	return imports
}

func emitParamsType(b *strings.Builder, ep endpoint) {
	if !ep.hasParamsArg() {
		return
	}
	fmt.Fprintf(b, "type %s struct {\n", ep.paramsType())
	for _, p := range ep.params {
		if p.doc != "" {
			fmt.Fprintf(b, "\t// %s\n", oneLine(p.doc))
		}
		fmt.Fprintf(b, "\t%s %s\n", p.field, fieldType(p))
	}
	if ep.hasBody {
		if ep.bodyRequired {
			fmt.Fprintf(b, "\tBody %s\n", ep.bodyType)
		} else {
			fmt.Fprintf(b, "\tBody *%s\n", ep.bodyType)
		}
	}
	b.WriteString("}\n\n")
}

// fieldType returns the Go type of a param field: arrays as-is, optional scalars
// as pointers, required scalars as values.
func fieldType(p param) string {
	if p.isArray || p.required {
		return p.goType
	}
	return "*" + p.goType
}

func emitMethod(b *strings.Builder, ns namespace, ep endpoint) {
	ret := ep.returnType
	if ret == "" {
		ret = "any"
	}

	if ep.summary != "" {
		fmt.Fprintf(b, "// %s %s\n", ep.goName, oneLine(ep.summary))
	}
	if ep.deprecated {
		if ep.summary != "" {
			b.WriteString("//\n")
		} else {
			fmt.Fprintf(b, "// %s is deprecated.\n//\n", ep.goName)
		}
		b.WriteString("// Deprecated: this endpoint is deprecated.\n")
	}

	if ep.hasParamsArg() {
		fmt.Fprintf(b, "func (s *%s) %s(ctx context.Context, params %s) (*Response[%s], error) {\n",
			ns.service, ep.goName, ep.paramsType(), ret)
	} else {
		fmt.Fprintf(b, "func (s *%s) %s(ctx context.Context) (*Response[%s], error) {\n",
			ns.service, ep.goName, ret)
	}

	fmt.Fprintf(b, "\tr := &request{method: %q, path: %q}\n", ep.httpMethod, ep.path)

	emitPathParams(b, ep)
	emitQueryParams(b, ep)
	emitHeaderParams(b, ep)
	if ep.hasBody {
		if ep.bodyRequired {
			b.WriteString("\tr.body = params.Body\n")
		} else {
			b.WriteString("\tif params.Body != nil {\n\t\tr.body = params.Body\n\t}\n")
		}
	}

	fmt.Fprintf(b, "\treturn do[%s](ctx, s.client, r)\n}\n\n", ret)
}

func emitPathParams(b *strings.Builder, ep endpoint) {
	var ps []param
	for _, p := range ep.params {
		if p.in == "path" {
			ps = append(ps, p)
		}
	}
	if len(ps) == 0 {
		return
	}
	b.WriteString("\tr.pathParams = map[string]string{\n")
	for _, p := range ps {
		fmt.Fprintf(b, "\t\t%q: fmt.Sprint(params.%s),\n", p.wire, p.field)
	}
	b.WriteString("\t}\n")
}

func emitQueryParams(b *strings.Builder, ep endpoint) {
	var ps []param
	for _, p := range ep.params {
		if p.in == "query" {
			ps = append(ps, p)
		}
	}
	if len(ps) == 0 {
		return
	}
	b.WriteString("\tq := url.Values{}\n")
	for _, p := range ps {
		emitValueSet(b, p, "q", "Add")
	}
	b.WriteString("\tif len(q) > 0 {\n\t\tr.query = q\n\t}\n")
}

func emitHeaderParams(b *strings.Builder, ep endpoint) {
	var ps []param
	for _, p := range ep.params {
		if p.in == "header" {
			ps = append(ps, p)
		}
	}
	if len(ps) == 0 {
		return
	}
	b.WriteString("\th := http.Header{}\n")
	for _, p := range ps {
		emitValueSet(b, p, "h", "Add")
	}
	b.WriteString("\tif len(h) > 0 {\n\t\tr.header = h\n\t}\n")
}

// emitValueSet writes the code that puts one param into a url.Values / Header.
func emitValueSet(b *strings.Builder, p param, target, addMethod string) {
	switch {
	case p.isArray:
		fmt.Fprintf(b, "\tfor _, v := range params.%s {\n\t\t%s.%s(%q, fmt.Sprint(v))\n\t}\n",
			p.field, target, addMethod, p.wire)
	case p.required:
		fmt.Fprintf(b, "\t%s.Set(%q, fmt.Sprint(params.%s))\n", target, p.wire, p.field)
	default:
		fmt.Fprintf(b, "\tif params.%s != nil {\n\t\t%s.Set(%q, fmt.Sprint(*params.%s))\n\t}\n",
			p.field, target, p.wire, p.field)
	}
}

func oneLine(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", " ")
	return strings.Join(strings.Fields(s), " ")
}

func writeGo(path, src string) {
	formatted, err := format.Source([]byte(src))
	if err != nil {
		// Write the unformatted source to aid debugging, then fail.
		_ = os.WriteFile(path, []byte(src), 0o644)
		must(fmt.Errorf("format %s: %w", path, err))
	}
	must(os.WriteFile(path, formatted, 0o644))
}
