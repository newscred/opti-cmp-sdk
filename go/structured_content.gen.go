// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type StructuredContentService struct{ client *Client }

type AcknowledgeSCContentPreviewParams struct {
	ContentID string
	VersionID string
	PreviewID string
	Body      schema.SCContentPreviewAcknowledgeRequest
}

// AcknowledgeSCContentPreview POST /structured-content/contents/{content_id}/versions/{version_id}/previews/{preview_id}/acknowledge
func (s *StructuredContentService) AcknowledgeSCContentPreview(ctx context.Context, params AcknowledgeSCContentPreviewParams) (*Response[schema.AcknowledgeSCContentPreviewResponse], error) {
	r := &request{method: "POST", path: "/structured-content/contents/{content_id}/versions/{version_id}/previews/{preview_id}/acknowledge"}
	r.pathParams = map[string]string{
		"content_id": fmt.Sprint(params.ContentID),
		"version_id": fmt.Sprint(params.VersionID),
		"preview_id": fmt.Sprint(params.PreviewID),
	}
	r.body = params.Body
	return do[schema.AcknowledgeSCContentPreviewResponse](ctx, s.client, r)
}

type CompleteSCContentPreviewParams struct {
	ContentID string
	VersionID string
	PreviewID string
	Body      schema.SCContentPreviewCompleteRequest
}

// CompleteSCContentPreview POST /structured-content/contents/{content_id}/versions/{version_id}/previews/{preview_id}/complete
func (s *StructuredContentService) CompleteSCContentPreview(ctx context.Context, params CompleteSCContentPreviewParams) (*Response[schema.CompleteSCContentPreviewResponse], error) {
	r := &request{method: "POST", path: "/structured-content/contents/{content_id}/versions/{version_id}/previews/{preview_id}/complete"}
	r.pathParams = map[string]string{
		"content_id": fmt.Sprint(params.ContentID),
		"version_id": fmt.Sprint(params.VersionID),
		"preview_id": fmt.Sprint(params.PreviewID),
	}
	r.body = params.Body
	return do[schema.CompleteSCContentPreviewResponse](ctx, s.client, r)
}

type CreateSCContentTypeParams struct {
	Body schema.SCContentTypeCreateRequest
}

// CreateSCContentType POST /structured-content/content-types
func (s *StructuredContentService) CreateSCContentType(ctx context.Context, params CreateSCContentTypeParams) (*Response[schema.SCContentTypeCreateResponse], error) {
	r := &request{method: "POST", path: "/structured-content/content-types"}
	r.body = params.Body
	return do[schema.SCContentTypeCreateResponse](ctx, s.client, r)
}

type CreateSCContentTypeManagedMigrationParams struct {
	ContentTypeID string
	Body          schema.SCContentTypeManagedMigrationCreateRequest
}

// CreateSCContentTypeManagedMigration POST /structured-content/content-types/{content_type_id}/managed-migrations
func (s *StructuredContentService) CreateSCContentTypeManagedMigration(ctx context.Context, params CreateSCContentTypeManagedMigrationParams) (*Response[schema.CreateSCContentTypeManagedMigrationResponse], error) {
	r := &request{method: "POST", path: "/structured-content/content-types/{content_type_id}/managed-migrations"}
	r.pathParams = map[string]string{
		"content_type_id": fmt.Sprint(params.ContentTypeID),
	}
	r.body = params.Body
	return do[schema.CreateSCContentTypeManagedMigrationResponse](ctx, s.client, r)
}

type CreateSCContentTypeVersionParams struct {
	ContentTypeID string
	Body          schema.SCContentTypeVersionCreateRequest
}

// CreateSCContentTypeVersion POST /structured-content/content-types/{content_type_id}/versions
func (s *StructuredContentService) CreateSCContentTypeVersion(ctx context.Context, params CreateSCContentTypeVersionParams) (*Response[schema.SCContentTypeCreateResponse], error) {
	r := &request{method: "POST", path: "/structured-content/content-types/{content_type_id}/versions"}
	r.pathParams = map[string]string{
		"content_type_id": fmt.Sprint(params.ContentTypeID),
	}
	r.body = params.Body
	return do[schema.SCContentTypeCreateResponse](ctx, s.client, r)
}

type DeleteSCContentTypeManagedMigrationParams struct {
	ContentTypeID string
	JobID         string
}

// DeleteSCContentTypeManagedMigration DELETE /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}
func (s *StructuredContentService) DeleteSCContentTypeManagedMigration(ctx context.Context, params DeleteSCContentTypeManagedMigrationParams) (*Response[any], error) {
	r := &request{method: "DELETE", path: "/structured-content/content-types/{content_type_id}/managed-migrations/{job_id}"}
	r.pathParams = map[string]string{
		"content_type_id": fmt.Sprint(params.ContentTypeID),
		"job_id":          fmt.Sprint(params.JobID),
	}
	return do[any](ctx, s.client, r)
}

type GetSCContentTypeParams struct {
	ContentTypeID string
}

// GetSCContentType GET /structured-content/content-types/{content_type_id}
func (s *StructuredContentService) GetSCContentType(ctx context.Context, params GetSCContentTypeParams) (*Response[schema.SCContentType], error) {
	r := &request{method: "GET", path: "/structured-content/content-types/{content_type_id}"}
	r.pathParams = map[string]string{
		"content_type_id": fmt.Sprint(params.ContentTypeID),
	}
	return do[schema.SCContentType](ctx, s.client, r)
}

type GetSCContentTypeManagedMigrationParams struct {
	ContentTypeID string
	JobID         string
}

// GetSCContentTypeManagedMigration GET /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}
func (s *StructuredContentService) GetSCContentTypeManagedMigration(ctx context.Context, params GetSCContentTypeManagedMigrationParams) (*Response[schema.SCContentTypeManagedMigrationResponse], error) {
	r := &request{method: "GET", path: "/structured-content/content-types/{content_type_id}/managed-migrations/{job_id}"}
	r.pathParams = map[string]string{
		"content_type_id": fmt.Sprint(params.ContentTypeID),
		"job_id":          fmt.Sprint(params.JobID),
	}
	return do[schema.SCContentTypeManagedMigrationResponse](ctx, s.client, r)
}

type GetSCContentTypeVersionParams struct {
	ContentTypeID string
	VersionID     string
}

// GetSCContentTypeVersion GET /structured-content/content-types/{content_type_id}/versions/{version_id}
func (s *StructuredContentService) GetSCContentTypeVersion(ctx context.Context, params GetSCContentTypeVersionParams) (*Response[schema.SCContentTypeVersion], error) {
	r := &request{method: "GET", path: "/structured-content/content-types/{content_type_id}/versions/{version_id}"}
	r.pathParams = map[string]string{
		"content_type_id": fmt.Sprint(params.ContentTypeID),
		"version_id":      fmt.Sprint(params.VersionID),
	}
	return do[schema.SCContentTypeVersion](ctx, s.client, r)
}

type ListSCContentTypeManagedMigrationsParams struct {
	ContentTypeID string
	// Whether include a summary of content migration status (total, not started, succeeded, errored).
	ContentMigrationSummary *bool
	// Pagination offset (number of jobs to skip).
	Offset *int
	// Pagination limit (number of jobs to return).
	Limit *int
}

// ListSCContentTypeManagedMigrations GET /structured-content/content-types/{content_type_id}/managed-migrations
func (s *StructuredContentService) ListSCContentTypeManagedMigrations(ctx context.Context, params ListSCContentTypeManagedMigrationsParams) (*Response[schema.ListSCContentTypeManagedMigrationsResponse], error) {
	r := &request{method: "GET", path: "/structured-content/content-types/{content_type_id}/managed-migrations"}
	r.pathParams = map[string]string{
		"content_type_id": fmt.Sprint(params.ContentTypeID),
	}
	q := url.Values{}
	if params.ContentMigrationSummary != nil {
		q.Set("content_migration_summary", fmt.Sprint(*params.ContentMigrationSummary))
	}
	if params.Offset != nil {
		q.Set("offset", fmt.Sprint(*params.Offset))
	}
	if params.Limit != nil {
		q.Set("limit", fmt.Sprint(*params.Limit))
	}
	if len(q) > 0 {
		r.query = q
	}
	return do[schema.ListSCContentTypeManagedMigrationsResponse](ctx, s.client, r)
}

type ListSCContentTypeVersionsParams struct {
	ContentTypeID string
}

// ListSCContentTypeVersions GET /structured-content/content-types/{content_type_id}/versions
func (s *StructuredContentService) ListSCContentTypeVersions(ctx context.Context, params ListSCContentTypeVersionsParams) (*Response[schema.ListSCContentTypeVersionsResponse], error) {
	r := &request{method: "GET", path: "/structured-content/content-types/{content_type_id}/versions"}
	r.pathParams = map[string]string{
		"content_type_id": fmt.Sprint(params.ContentTypeID),
	}
	return do[schema.ListSCContentTypeVersionsResponse](ctx, s.client, r)
}

type ListSCContentTypesParams struct {
	Source   *string
	Disabled *bool
	List     *schema.ContentTypeListingOption
}

// ListSCContentTypes GET /structured-content/content-types
func (s *StructuredContentService) ListSCContentTypes(ctx context.Context, params ListSCContentTypesParams) (*Response[schema.ListSCContentTypesResponse], error) {
	r := &request{method: "GET", path: "/structured-content/content-types"}
	q := url.Values{}
	if params.Source != nil {
		q.Set("source", fmt.Sprint(*params.Source))
	}
	if params.Disabled != nil {
		q.Set("disabled", fmt.Sprint(*params.Disabled))
	}
	if params.List != nil {
		q.Set("list", fmt.Sprint(*params.List))
	}
	if len(q) > 0 {
		r.query = q
	}
	return do[schema.ListSCContentTypesResponse](ctx, s.client, r)
}

type MigrateSCContentParams struct {
	ContentID string
	Body      schema.SCContentMigrationCreateRequest
}

// MigrateSCContent POST /structured-content/contents/{content_id}/migration
func (s *StructuredContentService) MigrateSCContent(ctx context.Context, params MigrateSCContentParams) (*Response[schema.MigrateSCContentResponse], error) {
	r := &request{method: "POST", path: "/structured-content/contents/{content_id}/migration"}
	r.pathParams = map[string]string{
		"content_id": fmt.Sprint(params.ContentID),
	}
	r.body = params.Body
	return do[schema.MigrateSCContentResponse](ctx, s.client, r)
}

type StartSCContentTypeManagedMigrationParams struct {
	ContentTypeID string
	JobID         string
}

// StartSCContentTypeManagedMigration POST /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}/start
func (s *StructuredContentService) StartSCContentTypeManagedMigration(ctx context.Context, params StartSCContentTypeManagedMigrationParams) (*Response[schema.SCContentTypeManagedMigrationStartResponse], error) {
	r := &request{method: "POST", path: "/structured-content/content-types/{content_type_id}/managed-migrations/{job_id}/start"}
	r.pathParams = map[string]string{
		"content_type_id": fmt.Sprint(params.ContentTypeID),
		"job_id":          fmt.Sprint(params.JobID),
	}
	return do[schema.SCContentTypeManagedMigrationStartResponse](ctx, s.client, r)
}

type UpdateSCContentTypeParams struct {
	ContentTypeID string
	Body          schema.SCContentTypeUpdateRequest
}

// UpdateSCContentType POST /structured-content/content-types/{content_type_id}
func (s *StructuredContentService) UpdateSCContentType(ctx context.Context, params UpdateSCContentTypeParams) (*Response[schema.SCContentTypeUpdateResponse], error) {
	r := &request{method: "POST", path: "/structured-content/content-types/{content_type_id}"}
	r.pathParams = map[string]string{
		"content_type_id": fmt.Sprint(params.ContentTypeID),
	}
	r.body = params.Body
	return do[schema.SCContentTypeUpdateResponse](ctx, s.client, r)
}

type UpdateSCContentTypeManagedMigrationParams struct {
	ContentTypeID string
	JobID         string
	Body          schema.UpdateSCContentTypeManagedMigrationRequest
}

// UpdateSCContentTypeManagedMigration PATCH /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}
func (s *StructuredContentService) UpdateSCContentTypeManagedMigration(ctx context.Context, params UpdateSCContentTypeManagedMigrationParams) (*Response[schema.UpdateSCContentTypeManagedMigrationResponse], error) {
	r := &request{method: "PATCH", path: "/structured-content/content-types/{content_type_id}/managed-migrations/{job_id}"}
	r.pathParams = map[string]string{
		"content_type_id": fmt.Sprint(params.ContentTypeID),
		"job_id":          fmt.Sprint(params.JobID),
	}
	r.body = params.Body
	return do[schema.UpdateSCContentTypeManagedMigrationResponse](ctx, s.client, r)
}

type ValidateSCContentTypeManagedMigrationParams struct {
	ContentTypeID string
	Body          schema.SCContentTypeManagedMigrationValidateRequest
}

// ValidateSCContentTypeManagedMigration POST /structured-content/content-types/{content_type_id}/managed-migrations/validate
func (s *StructuredContentService) ValidateSCContentTypeManagedMigration(ctx context.Context, params ValidateSCContentTypeManagedMigrationParams) (*Response[schema.ValidateSCContentTypeManagedMigrationResponse], error) {
	r := &request{method: "POST", path: "/structured-content/content-types/{content_type_id}/managed-migrations/validate"}
	r.pathParams = map[string]string{
		"content_type_id": fmt.Sprint(params.ContentTypeID),
	}
	r.body = params.Body
	return do[schema.ValidateSCContentTypeManagedMigrationResponse](ctx, s.client, r)
}
