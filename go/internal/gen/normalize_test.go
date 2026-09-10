package main

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/codegen"
)

// TestInitialismNormalizer guards the reuse of oapi-codegen's normalizer. Its
// ToCamelCaseWithInitialisms only applies initialisms after Generate has primed
// the package global, so this primes it the same way generateSchemaTypes does,
// then checks the casing this generator relies on for its type references.
func TestInitialismNormalizer(t *testing.T) {
	const spec = `{
		"openapi": "3.0.0",
		"info": {"title": "t", "version": "1"},
		"paths": {},
		"components": {"schemas": {"X": {"type": "object", "properties": {"task_id": {"type": "string"}}}}}
	}`
	doc, err := openapi3.NewLoader().LoadFromData([]byte(spec))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := codegen.Generate(doc, codegen.Configuration{
		PackageName:   "schema",
		Generate:      codegen.GenerateOptions{Models: true},
		OutputOptions: codegen.OutputOptions{NameNormalizer: string(codegen.NameNormalizerFunctionToCamelCaseWithInitialisms)},
	}); err != nil {
		t.Fatal(err)
	}

	cases := map[string]string{
		"task_id":       "TaskID",
		"get_asset_url": "GetAssetURL",
		"html_body":     "HTMLBody",
		"api_url":       "APIURL",
		"page_size":     "PageSize",
	}
	for in, want := range cases {
		if got := codegen.ToCamelCaseWithInitialisms(in); got != want {
			t.Errorf("ToCamelCaseWithInitialisms(%q) = %q, want %q", in, got, want)
		}
	}
}
