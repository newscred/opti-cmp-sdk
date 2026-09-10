// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type CampaignService struct{ client *Client }

type AddAttachmentToCampaignParams struct {
	ID   string
	Body schema.AttachmentRequest
}

// AddAttachmentToCampaign POST /campaigns/{id}/attachments
func (s *CampaignService) AddAttachmentToCampaign(ctx context.Context, params AddAttachmentToCampaignParams) (*Response[schema.AttachmentResponse], error) {
	r := &request{method: "POST", path: "/campaigns/{id}/attachments"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.AttachmentResponse](ctx, s.client, r)
}

type AddCommentToCampaignParams struct {
	ID   string
	Body schema.CommentCreateRequest
}

// AddCommentToCampaign POST /campaigns/{id}/comments
func (s *CampaignService) AddCommentToCampaign(ctx context.Context, params AddCommentToCampaignParams) (*Response[schema.CampaignCommentResponse], error) {
	r := &request{method: "POST", path: "/campaigns/{id}/comments"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.CampaignCommentResponse](ctx, s.client, r)
}

type AddFieldToCampaignParams struct {
	ID   string
	Body schema.ObjectFieldCreateRequest
}

// AddFieldToCampaign POST campaigns/{id}/fields
func (s *CampaignService) AddFieldToCampaign(ctx context.Context, params AddFieldToCampaignParams) (*Response[schema.ObjectFieldCreateResponse], error) {
	r := &request{method: "POST", path: "/campaigns/{id}/fields"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.ObjectFieldCreateResponse](ctx, s.client, r)
}

type CreateCampaignParams struct {
	Body schema.CampaignCreateRequest
}

// CreateCampaign POST /campaigns
func (s *CampaignService) CreateCampaign(ctx context.Context, params CreateCampaignParams) (*Response[schema.CampaignResponse], error) {
	r := &request{method: "POST", path: "/campaigns"}
	r.body = params.Body
	return do[schema.CampaignResponse](ctx, s.client, r)
}

type GetCampaignParams struct {
	ID string
}

// GetCampaign GET /campaigns/{id}
func (s *CampaignService) GetCampaign(ctx context.Context, params GetCampaignParams) (*Response[schema.CampaignResponse], error) {
	r := &request{method: "GET", path: "/campaigns/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.CampaignResponse](ctx, s.client, r)
}

type GetCampaignBriefParams struct {
	ID string
}

// GetCampaignBrief GET /campaigns/{id}/brief
func (s *CampaignService) GetCampaignBrief(ctx context.Context, params GetCampaignBriefParams) (*Response[schema.CampaignBriefResponse], error) {
	r := &request{method: "GET", path: "/campaigns/{id}/brief"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.CampaignBriefResponse](ctx, s.client, r)
}

type ListCampaignFieldsParams struct {
	ID       string
	Offset   *int
	PageSize *int
}

// ListCampaignFields GET campaigns/{id}/fields
func (s *CampaignService) ListCampaignFields(ctx context.Context, params ListCampaignFieldsParams) (*Response[schema.ListCampaignFieldsResponse], error) {
	r := &request{method: "GET", path: "/campaigns/{id}/fields"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
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
	return do[schema.ListCampaignFieldsResponse](ctx, s.client, r)
}

type ListCampaignsParams struct {
	Owner     *string
	StartDate *string
	EndDate   *string
	Offset    *int
	PageSize  *int
}

// ListCampaigns GET /campaigns
func (s *CampaignService) ListCampaigns(ctx context.Context, params ListCampaignsParams) (*Response[schema.ListCampaignsResponse], error) {
	r := &request{method: "GET", path: "/campaigns"}
	q := url.Values{}
	if params.Owner != nil {
		q.Set("owner", fmt.Sprint(*params.Owner))
	}
	if params.StartDate != nil {
		q.Set("start_date", fmt.Sprint(*params.StartDate))
	}
	if params.EndDate != nil {
		q.Set("end_date", fmt.Sprint(*params.EndDate))
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
	return do[schema.ListCampaignsResponse](ctx, s.client, r)
}

type UpdateCampaignParams struct {
	ID   string
	Body schema.CampaignUpdateRequest
}

// UpdateCampaign PATCH /campaigns/{id}
func (s *CampaignService) UpdateCampaign(ctx context.Context, params UpdateCampaignParams) (*Response[any], error) {
	r := &request{method: "PATCH", path: "/campaigns/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[any](ctx, s.client, r)
}

type UpdateCampaignFieldParams struct {
	CampaignID string
	FieldID    string
	Body       schema.CampaignFieldUpdateRequest
}

// UpdateCampaignField PUT /campaigns/{campaign_id}/fields/{field_id}
func (s *CampaignService) UpdateCampaignField(ctx context.Context, params UpdateCampaignFieldParams) (*Response[schema.CampaignFieldUpdateResponse], error) {
	r := &request{method: "PUT", path: "/campaigns/{campaign_id}/fields/{field_id}"}
	r.pathParams = map[string]string{
		"campaign_id": fmt.Sprint(params.CampaignID),
		"field_id":    fmt.Sprint(params.FieldID),
	}
	r.body = params.Body
	return do[schema.CampaignFieldUpdateResponse](ctx, s.client, r)
}
