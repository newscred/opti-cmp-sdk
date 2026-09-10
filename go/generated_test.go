package opticmp

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGeneratedSurface exercises generated namespace methods end to end: path
// parameter expansion, query building, and typed decoding.
func TestGeneratedSurface(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/tasks/task-1":
			_, _ = w.Write([]byte(`{"id":"task-1","title":"Write docs"}`))
		case "/assets":
			if r.URL.Query().Get("page_size") != "2" {
				t.Errorf("page_size query = %q", r.URL.Query().Get("page_size"))
			}
			_, _ = w.Write([]byte(`{"data":[{"id":"a1"}],"pagination":{"next":null,"previous":null},"total_count":1}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL), WithToken("tok"))
	ctx := context.Background()

	task, err := c.Task.GetTask(ctx, GetTaskParams{ID: "task-1"})
	if err != nil {
		t.Fatal(err)
	}
	if task.Data.ID != "task-1" {
		t.Fatalf("task id = %v", task.Data.ID)
	}

	assets, err := c.Library.ListAssets(ctx, ListAssetsParams{PageSize: Ptr(2)})
	if err != nil {
		t.Fatal(err)
	}
	if len(assets.Data.Data) != 1 {
		t.Fatalf("assets = %d, want 1", len(assets.Data.Data))
	}
}
