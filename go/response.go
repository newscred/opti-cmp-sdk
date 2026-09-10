package opticmp

import "net/http"

// Response wraps a decoded API response. Data holds the decoded body typed to
// the endpoint's return type.
type Response[T any] struct {
	Data   T
	Header http.Header
	Status int
	URL    string

	// client and raw support the pagination helpers, which read the pagination
	// URLs from the untyped body and follow them.
	client *Client
	raw    []byte
}
