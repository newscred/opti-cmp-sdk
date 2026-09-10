package opticmp

// Ptr returns a pointer to v, for setting optional fields where a nil pointer
// means "unset" and is distinct from the zero value.
//
// As of Go 1.26 the built-in new accepts a value expression, so new(v) does the
// same thing; prefer it once the module's Go version allows it.
func Ptr[T any](v T) *T { return &v }
