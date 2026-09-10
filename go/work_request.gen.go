// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type WorkRequestService struct{ client *Client }

type AddAttachmentToWorkRequestParams struct {
	ID   string
	Body schema.AttachmentRequest
}

// AddAttachmentToWorkRequest POST /work-requests/{id}/attachments
func (s *WorkRequestService) AddAttachmentToWorkRequest(ctx context.Context, params AddAttachmentToWorkRequestParams) (*Response[schema.AttachmentResponse], error) {
	r := &request{method: "POST", path: "/work-requests/{id}/attachments"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.AttachmentResponse](ctx, s.client, r)
}

type AddCommentToWorkRequestParams struct {
	ID   string
	Body schema.CommentWithReplyCreateRequest
}

// AddCommentToWorkRequest POST /work-requests/{id}/comments
func (s *WorkRequestService) AddCommentToWorkRequest(ctx context.Context, params AddCommentToWorkRequestParams) (*Response[schema.WorkRequestCommentResponse], error) {
	r := &request{method: "POST", path: "/work-requests/{id}/comments"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.WorkRequestCommentResponse](ctx, s.client, r)
}

type CreateCampaignFromWorkRequestParams struct {
	ID   string
	Body schema.WorkRequestCampaignRequest
}

// CreateCampaignFromWorkRequest POST /work-requests/{id}/campaigns
func (s *WorkRequestService) CreateCampaignFromWorkRequest(ctx context.Context, params CreateCampaignFromWorkRequestParams) (*Response[schema.WorkRequestCampaignResponse], error) {
	r := &request{method: "POST", path: "/work-requests/{id}/campaigns"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.WorkRequestCampaignResponse](ctx, s.client, r)
}

type CreateTaskFromWorkRequestParams struct {
	ID   string
	Body schema.WorkRequestTaskRequest
}

// CreateTaskFromWorkRequest POST /work-requests/{id}/tasks
func (s *WorkRequestService) CreateTaskFromWorkRequest(ctx context.Context, params CreateTaskFromWorkRequestParams) (*Response[schema.WorkRequestTaskResponse], error) {
	r := &request{method: "POST", path: "/work-requests/{id}/tasks"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.WorkRequestTaskResponse](ctx, s.client, r)
}

type CreateWorkRequestParams struct {
	Body schema.WorkRequestCreateRequest
}

// CreateWorkRequest POST /work-requests
func (s *WorkRequestService) CreateWorkRequest(ctx context.Context, params CreateWorkRequestParams) (*Response[schema.WorkRequestResponse], error) {
	r := &request{method: "POST", path: "/work-requests"}
	r.body = params.Body
	return do[schema.WorkRequestResponse](ctx, s.client, r)
}

type CreateWorkRequestCreativeAssetParams struct {
	ID   string
	Body schema.CreativeAssetRequest
}

// CreateWorkRequestCreativeAsset POST /work-requests/{id}/creative-assets
func (s *WorkRequestService) CreateWorkRequestCreativeAsset(ctx context.Context, params CreateWorkRequestCreativeAssetParams) (*Response[schema.CreativeAssetResponse], error) {
	r := &request{method: "POST", path: "/work-requests/{id}/creative-assets"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.CreativeAssetResponse](ctx, s.client, r)
}

type DeleteWorkRequestAttachmentParams struct {
	WorkRequestID string
	AttachmentID  string
}

// DeleteWorkRequestAttachment DELETE /work-requests/{work_request_id}/attachments/{attachment_id}
func (s *WorkRequestService) DeleteWorkRequestAttachment(ctx context.Context, params DeleteWorkRequestAttachmentParams) (*Response[any], error) {
	r := &request{method: "DELETE", path: "/work-requests/{work_request_id}/attachments/{attachment_id}"}
	r.pathParams = map[string]string{
		"work_request_id": fmt.Sprint(params.WorkRequestID),
		"attachment_id":   fmt.Sprint(params.AttachmentID),
	}
	return do[any](ctx, s.client, r)
}

type DeleteWorkRequestCreativeAssetParams struct {
	WorkRequestID   string
	CreativeAssetID string
}

// DeleteWorkRequestCreativeAsset DELETE /work-requests/{work_request_id}/creative-assets/{creative_asset_id}
func (s *WorkRequestService) DeleteWorkRequestCreativeAsset(ctx context.Context, params DeleteWorkRequestCreativeAssetParams) (*Response[any], error) {
	r := &request{method: "DELETE", path: "/work-requests/{work_request_id}/creative-assets/{creative_asset_id}"}
	r.pathParams = map[string]string{
		"work_request_id":   fmt.Sprint(params.WorkRequestID),
		"creative_asset_id": fmt.Sprint(params.CreativeAssetID),
	}
	return do[any](ctx, s.client, r)
}

type GetWorkRequestParams struct {
	ID string
}

// GetWorkRequest GET /work-requests/{id}
func (s *WorkRequestService) GetWorkRequest(ctx context.Context, params GetWorkRequestParams) (*Response[schema.WorkRequestResponse], error) {
	r := &request{method: "GET", path: "/work-requests/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.WorkRequestResponse](ctx, s.client, r)
}

type GetWorkRequestCommentParams struct {
	WorkRequestID string
	CommentID     string
}

// GetWorkRequestComment GET /work-requests/{work_request_id}/comments/{comment_id}
func (s *WorkRequestService) GetWorkRequestComment(ctx context.Context, params GetWorkRequestCommentParams) (*Response[schema.WorkRequestCommentResponse], error) {
	r := &request{method: "GET", path: "/work-requests/{work_request_id}/comments/{comment_id}"}
	r.pathParams = map[string]string{
		"work_request_id": fmt.Sprint(params.WorkRequestID),
		"comment_id":      fmt.Sprint(params.CommentID),
	}
	return do[schema.WorkRequestCommentResponse](ctx, s.client, r)
}

type ListWorkRequestApprovedAssetsParams struct {
	ID       string
	Offset   *int
	PageSize *int
}

// ListWorkRequestApprovedAssets GET /work-requests/{id}/approved-assets
func (s *WorkRequestService) ListWorkRequestApprovedAssets(ctx context.Context, params ListWorkRequestApprovedAssetsParams) (*Response[schema.ListWorkRequestApprovedAssetsResponse], error) {
	r := &request{method: "GET", path: "/work-requests/{id}/approved-assets"}
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
	return do[schema.ListWorkRequestApprovedAssetsResponse](ctx, s.client, r)
}

type ListWorkRequestCommentsParams struct {
	ID       string
	Offset   *int
	PageSize *int
}

// ListWorkRequestComments GET /work-requests/{id}/comments
func (s *WorkRequestService) ListWorkRequestComments(ctx context.Context, params ListWorkRequestCommentsParams) (*Response[schema.ListWorkRequestCommentsResponse], error) {
	r := &request{method: "GET", path: "/work-requests/{id}/comments"}
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
	return do[schema.ListWorkRequestCommentsResponse](ctx, s.client, r)
}

type ListWorkRequestRelatedResourcesParams struct {
	ID       string
	Offset   *int
	PageSize *int
}

// ListWorkRequestRelatedResources GET /work-requests/{id}/related-resources
func (s *WorkRequestService) ListWorkRequestRelatedResources(ctx context.Context, params ListWorkRequestRelatedResourcesParams) (*Response[schema.ListWorkRequestRelatedResourcesResponse], error) {
	r := &request{method: "GET", path: "/work-requests/{id}/related-resources"}
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
	return do[schema.ListWorkRequestRelatedResourcesResponse](ctx, s.client, r)
}

type ListWorkRequestsParams struct {
	CreatedBy *string
	Status    *string
	OrderBy   *string
	OrderAs   *string
	// Unique identifiers of the templates. You can append multiple times, for example `template_ids=template1&template_ids=template2`.
	TemplateIds []string
	// Date and time as the lower limit to filter work requests by `created_at`, in ISO 8601 UTC format
	CreatedAtFrom *string
	// Date and time as the upper limit to filter work requests by `created_at`, in ISO 8601 UTC format
	CreatedAtTo *string
	Offset      *int
	PageSize    *int
}

// ListWorkRequests GET /work-requests
func (s *WorkRequestService) ListWorkRequests(ctx context.Context, params ListWorkRequestsParams) (*Response[schema.ListWorkRequestsResponse], error) {
	r := &request{method: "GET", path: "/work-requests"}
	q := url.Values{}
	if params.CreatedBy != nil {
		q.Set("created_by", fmt.Sprint(*params.CreatedBy))
	}
	if params.Status != nil {
		q.Set("status", fmt.Sprint(*params.Status))
	}
	if params.OrderBy != nil {
		q.Set("order_by", fmt.Sprint(*params.OrderBy))
	}
	if params.OrderAs != nil {
		q.Set("order_as", fmt.Sprint(*params.OrderAs))
	}
	for _, v := range params.TemplateIds {
		q.Add("template_ids", fmt.Sprint(v))
	}
	if params.CreatedAtFrom != nil {
		q.Set("created_at__from", fmt.Sprint(*params.CreatedAtFrom))
	}
	if params.CreatedAtTo != nil {
		q.Set("created_at__to", fmt.Sprint(*params.CreatedAtTo))
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
	return do[schema.ListWorkRequestsResponse](ctx, s.client, r)
}

type UpdateWorkRequestParams struct {
	ID   string
	Body schema.WorkRequestUpdateRequest
}

// UpdateWorkRequest PATCH /work-requests/{id}
func (s *WorkRequestService) UpdateWorkRequest(ctx context.Context, params UpdateWorkRequestParams) (*Response[any], error) {
	r := &request{method: "PATCH", path: "/work-requests/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[any](ctx, s.client, r)
}

type UpdateWorkRequestFormFieldParams struct {
	WorkRequestID       string
	FormFieldIdentifier string
	Body                schema.WorkRequestFormFieldUpdateRequest
}

// UpdateWorkRequestFormField PUT /work-requests/{work_request_id}/form-fields/{form_field_identifier}
func (s *WorkRequestService) UpdateWorkRequestFormField(ctx context.Context, params UpdateWorkRequestFormFieldParams) (*Response[schema.WorkRequestRequestFormFieldUpdateResponse], error) {
	r := &request{method: "PUT", path: "/work-requests/{work_request_id}/form-fields/{form_field_identifier}"}
	r.pathParams = map[string]string{
		"work_request_id":       fmt.Sprint(params.WorkRequestID),
		"form_field_identifier": fmt.Sprint(params.FormFieldIdentifier),
	}
	r.body = params.Body
	return do[schema.WorkRequestRequestFormFieldUpdateResponse](ctx, s.client, r)
}
