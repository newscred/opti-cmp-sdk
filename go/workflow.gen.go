// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type WorkflowService struct{ client *Client }

type GetWorkflowParams struct {
	WorkflowID string
}

// GetWorkflow GET /workflows/{workflow_id}
func (s *WorkflowService) GetWorkflow(ctx context.Context, params GetWorkflowParams) (*Response[schema.WorkflowResponse], error) {
	r := &request{method: "GET", path: "/workflows/{workflow_id}"}
	r.pathParams = map[string]string{
		"workflow_id": fmt.Sprint(params.WorkflowID),
	}
	return do[schema.WorkflowResponse](ctx, s.client, r)
}

type ListWorkflowsParams struct {
	Offset   *int
	PageSize *int
}

// ListWorkflows GET /workflows
func (s *WorkflowService) ListWorkflows(ctx context.Context, params ListWorkflowsParams) (*Response[schema.ListWorkflowsResponse], error) {
	r := &request{method: "GET", path: "/workflows"}
	q := url.Values{}
	if params.Offset != nil {
		q.Set("offset", fmt.Sprint(*params.Offset))
	}
	if params.PageSize != nil {
		q.Set("page_size", fmt.Sprint(*params.PageSize))
	}
	if len(q) > 0 {
		r.query = q
	}
	return do[schema.ListWorkflowsResponse](ctx, s.client, r)
}
