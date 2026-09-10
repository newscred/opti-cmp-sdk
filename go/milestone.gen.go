// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type MilestoneService struct{ client *Client }

type CreateMilestoneParams struct {
	Body schema.MilestoneCreateRequest
}

// CreateMilestone POST /milestones
func (s *MilestoneService) CreateMilestone(ctx context.Context, params CreateMilestoneParams) (*Response[schema.MilestoneResponse], error) {
	r := &request{method: "POST", path: "/milestones"}
	r.body = params.Body
	return do[schema.MilestoneResponse](ctx, s.client, r)
}

type GetMilestoneParams struct {
	ID string
}

// GetMilestone GET /milestones/{id}
func (s *MilestoneService) GetMilestone(ctx context.Context, params GetMilestoneParams) (*Response[schema.MilestoneResponse], error) {
	r := &request{method: "GET", path: "/milestones/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.MilestoneResponse](ctx, s.client, r)
}

type ListMilestonesParams struct {
	CampaignID  *string
	DueDateFrom *string
	DueDateTo   *string
	Offset      *int
	PageSize    *int
}

// ListMilestones GET /milestones
func (s *MilestoneService) ListMilestones(ctx context.Context, params ListMilestonesParams) (*Response[schema.ListMilestonesResponse], error) {
	r := &request{method: "GET", path: "/milestones"}
	q := url.Values{}
	if params.CampaignID != nil {
		q.Set("campaign_id", fmt.Sprint(*params.CampaignID))
	}
	if params.DueDateFrom != nil {
		q.Set("due_date__from", fmt.Sprint(*params.DueDateFrom))
	}
	if params.DueDateTo != nil {
		q.Set("due_date__to", fmt.Sprint(*params.DueDateTo))
	}
	if params.Offset != nil {
		q.Set("offset", fmt.Sprint(*params.Offset))
	}
	if params.PageSize != nil {
		q.Set("page_size", fmt.Sprint(*params.PageSize))
	}
	if len(q) > 0 {
		r.query = q
	}
	return do[schema.ListMilestonesResponse](ctx, s.client, r)
}

type UpdateMilestoneParams struct {
	ID   string
	Body schema.MilestoneUpdateRequest
}

// UpdateMilestone PATCH /milestones/{id}
func (s *MilestoneService) UpdateMilestone(ctx context.Context, params UpdateMilestoneParams) (*Response[schema.MilestoneResponse], error) {
	r := &request{method: "PATCH", path: "/milestones/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.MilestoneResponse](ctx, s.client, r)
}
