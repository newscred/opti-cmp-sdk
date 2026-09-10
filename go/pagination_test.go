package opticmp

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPaginationHelpers(t *testing.T) {
	var page2Hit bool
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Query().Get("offset") == "10" {
			page2Hit = true
			_, _ = w.Write([]byte(`{"data":[],"pagination":{"next":null,"previous":null}}`))
			return
		}
		// Page 1 points "next" at this same server with offset=10.
		_, _ = w.Write([]byte(`{"data":[],"pagination":{"next":"` + srvURL(r) + `?offset=10","previous":null}}`))
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))
	type list struct{}
	resp, err := do[list](context.Background(), c, &request{method: "GET", path: "/assets"})
	if err != nil {
		t.Fatal(err)
	}

	if !HasNextPage(resp) {
		t.Fatal("HasNextPage = false, want true")
	}
	if HasPreviousPage(resp) {
		t.Fatal("HasPreviousPage = true, want false")
	}

	next, err := GetNextPage(context.Background(), resp)
	if err != nil {
		t.Fatal(err)
	}
	if !page2Hit {
		t.Fatal("next page URL was not followed")
	}
	if HasNextPage(next) {
		t.Fatal("last page reported a next page")
	}

	// No previous page available -> error.
	if _, err := GetPreviousPage(context.Background(), resp); err == nil {
		t.Fatal("want error for missing previous page")
	}
}

func srvURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + r.Host + r.URL.Path
}
