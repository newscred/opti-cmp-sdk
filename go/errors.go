package opticmp

import (
	"fmt"
	"net/http"
)

// RequestSpec describes the request that produced an [HTTPError]. Body is nil
// for requests that carry secrets (for example token requests), so it is safe
// to log.
type RequestSpec struct {
	Method string
	URL    string
	Header http.Header
	Body   any
}

// HTTPError is returned when a request fails to send or the API responds with a
// non-2xx status. Status is 0 when the request never reached the server.
type HTTPError struct {
	Status  int
	Message string
	Data    any
	Header  http.Header
	Request *RequestSpec
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("opticmp: HTTP %d: %s", e.Status, e.Message)
}
