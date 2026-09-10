// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type BrandComplianceService struct{ client *Client }

type GetTaskAssetDraftBrandComplianceParams struct {
	TaskID  string
	AssetID string
	DraftID string
}

// GetTaskAssetDraftBrandCompliance GET /tasks/{task_id}/assets/{asset_id}/drafts/{draft_id}/brand-compliance
func (s *BrandComplianceService) GetTaskAssetDraftBrandCompliance(ctx context.Context, params GetTaskAssetDraftBrandComplianceParams) (*Response[schema.TaskAssetDraftBrandComplianceResponse], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/assets/{asset_id}/drafts/{draft_id}/brand-compliance"}
	r.pathParams = map[string]string{
		"task_id":  fmt.Sprint(params.TaskID),
		"asset_id": fmt.Sprint(params.AssetID),
		"draft_id": fmt.Sprint(params.DraftID),
	}
	return do[schema.TaskAssetDraftBrandComplianceResponse](ctx, s.client, r)
}

type ListBrandComplianceCategoriesParams struct {
	Offset   *int
	PageSize *int
}

// ListBrandComplianceCategories GET /brand-compliance/categories
func (s *BrandComplianceService) ListBrandComplianceCategories(ctx context.Context, params ListBrandComplianceCategoriesParams) (*Response[schema.ListBrandComplianceCategoriesResponse], error) {
	r := &request{method: "GET", path: "/brand-compliance/categories"}
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
	return do[schema.ListBrandComplianceCategoriesResponse](ctx, s.client, r)
}

type UpdateTaskAssetDraftBrandComplianceParams struct {
	TaskID  string
	AssetID string
	DraftID string
	Body    schema.TaskAssetDraftBrandComplianceRequest
}

// UpdateTaskAssetDraftBrandCompliance PUT /tasks/{task_id}/assets/{asset_id}/drafts/{draft_id}/brand-compliance
func (s *BrandComplianceService) UpdateTaskAssetDraftBrandCompliance(ctx context.Context, params UpdateTaskAssetDraftBrandComplianceParams) (*Response[schema.TaskAssetDraftBrandComplianceResponse], error) {
	r := &request{method: "PUT", path: "/tasks/{task_id}/assets/{asset_id}/drafts/{draft_id}/brand-compliance"}
	r.pathParams = map[string]string{
		"task_id":  fmt.Sprint(params.TaskID),
		"asset_id": fmt.Sprint(params.AssetID),
		"draft_id": fmt.Sprint(params.DraftID),
	}
	r.body = params.Body
	return do[schema.TaskAssetDraftBrandComplianceResponse](ctx, s.client, r)
}
