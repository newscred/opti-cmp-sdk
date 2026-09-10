package opticmp

import "testing"

func TestHTTPErrorFormat(t *testing.T) {
	err := &HTTPError{Status: 404, Message: "not found"}
	if got, want := err.Error(), "opticmp: HTTP 404: not found"; got != want {
		t.Fatalf("Error() = %q, want %q", got, want)
	}
}
