// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
)

type TaskStepService struct{ client *Client }

type UpdateTaskStepParams struct {
	TaskID string
	StepID string
	Body   schema.TaskStepRequest
}

// UpdateTaskStep PATCH /tasks/{task_id}/steps/{step_id}
func (s *TaskStepService) UpdateTaskStep(ctx context.Context, params UpdateTaskStepParams) (*Response[schema.TaskStep], error) {
	r := &request{method: "PATCH", path: "/tasks/{task_id}/steps/{step_id}"}
	r.pathParams = map[string]string{
		"task_id": fmt.Sprint(params.TaskID),
		"step_id": fmt.Sprint(params.StepID),
	}
	r.body = params.Body
	return do[schema.TaskStep](ctx, s.client, r)
}
