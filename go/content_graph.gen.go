// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type ContentGraphService struct{ client *Client }

type ListContentGraphsParams struct {
	// Type of the content graph instances to return. Use `dam` for instances holding library assets and `sc` for instances holding structured contents.
	Type string
	// Number of results to return per page
	PageSize *int
}

// ListContentGraphs GET /content-graphs
func (s *ContentGraphService) ListContentGraphs(ctx context.Context, params ListContentGraphsParams) (*Response[schema.ContentGraphListResponse], error) {
	r := &request{method: "GET", path: "/content-graphs"}
	q := url.Values{}
	q.Set("type", fmt.Sprint(params.Type))
	if params.PageSize != nil {
		q.Set("page_size", fmt.Sprint(*params.PageSize))
	}
	if len(q) > 0 {
		r.query = q
	}
	return do[schema.ContentGraphListResponse](ctx, s.client, r)
}

type QueryContentGraphParams struct {
	CgInstanceID string
	Body         schema.ContentGraphQueryRequest
}

// QueryContentGraph POST /content-graphs/{cg_instance_id}/query
func (s *ContentGraphService) QueryContentGraph(ctx context.Context, params QueryContentGraphParams) (*Response[schema.ContentGraphQueryResponse], error) {
	r := &request{method: "POST", path: "/content-graphs/{cg_instance_id}/query"}
	r.pathParams = map[string]string{
		"cg_instance_id": fmt.Sprint(params.CgInstanceID),
	}
	r.body = params.Body
	return do[schema.ContentGraphQueryResponse](ctx, s.client, r)
}
