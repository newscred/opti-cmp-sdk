package opticmp

import (
	"context"
	"encoding/json"
	"net/http"
)

// pageURL reads a pagination URL ("next" or "previous") from the untyped
// response body. It returns "" when the direction has no page.
func pageURL[T any](r *Response[T], direction string) string {
	if r == nil || len(r.raw) == 0 {
		return ""
	}
	var body struct {
		Pagination struct {
			Next     *string `json:"next"`
			Previous *string `json:"previous"`
		} `json:"pagination"`
	}
	if json.Unmarshal(r.raw, &body) != nil {
		return ""
	}
	value := body.Pagination.Next
	if direction == "previous" {
		value = body.Pagination.Previous
	}
	if value == nil {
		return ""
	}
	return *value
}

// HasNextPage reports whether the response has a next page.
func HasNextPage[T any](r *Response[T]) bool { return pageURL(r, "next") != "" }

// HasPreviousPage reports whether the response has a previous page.
func HasPreviousPage[T any](r *Response[T]) bool { return pageURL(r, "previous") != "" }

// GetNextPage fetches the next page of a paginated response.
func GetNextPage[T any](ctx context.Context, r *Response[T]) (*Response[T], error) {
	return getPage(ctx, r, "next")
}

// GetPreviousPage fetches the previous page of a paginated response.
func GetPreviousPage[T any](ctx context.Context, r *Response[T]) (*Response[T], error) {
	return getPage(ctx, r, "previous")
}

func getPage[T any](ctx context.Context, r *Response[T], direction string) (*Response[T], error) {
	url := pageURL(r, direction)
	if url == "" {
		return nil, &HTTPError{Status: http.StatusNotFound, Message: "no " + direction + " page available"}
	}
	return do[T](ctx, r.client, &request{method: http.MethodGet, path: url})
}
