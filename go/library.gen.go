// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type LibraryService struct{ client *Client }

type AddAssetPermissionsParams struct {
	AssetID string
	Body    schema.AssetPermissionBulkCreateRequest
}

// AddAssetPermissions POST /assets/{asset_id}/permissions
func (s *LibraryService) AddAssetPermissions(ctx context.Context, params AddAssetPermissionsParams) (*Response[any], error) {
	r := &request{method: "POST", path: "/assets/{asset_id}/permissions"}
	r.pathParams = map[string]string{
		"asset_id": fmt.Sprint(params.AssetID),
	}
	r.body = params.Body
	return do[any](ctx, s.client, r)
}

type AddFolderPermissionsParams struct {
	ID   string
	Body schema.FolderPermissionBulkCreateRequest
}

// AddFolderPermissions POST /folders/{id}/permissions
func (s *LibraryService) AddFolderPermissions(ctx context.Context, params AddFolderPermissionsParams) (*Response[any], error) {
	r := &request{method: "POST", path: "/folders/{id}/permissions"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[any](ctx, s.client, r)
}

type CreateAssetParams struct {
	Body schema.LibraryAssetCreateRequest
}

// CreateAsset POST /assets
func (s *LibraryService) CreateAsset(ctx context.Context, params CreateAssetParams) (*Response[schema.AssetResponse], error) {
	r := &request{method: "POST", path: "/assets"}
	r.body = params.Body
	return do[schema.AssetResponse](ctx, s.client, r)
}

type CreateAssetLineageParams struct {
	AssetID string
	Body    schema.AssetLineageCreateRequest
}

// CreateAssetLineage POST /assets/{asset_id}/lineages
func (s *LibraryService) CreateAssetLineage(ctx context.Context, params CreateAssetLineageParams) (*Response[schema.AssetLineageResponse], error) {
	r := &request{method: "POST", path: "/assets/{asset_id}/lineages"}
	r.pathParams = map[string]string{
		"asset_id": fmt.Sprint(params.AssetID),
	}
	r.body = params.Body
	return do[schema.AssetLineageResponse](ctx, s.client, r)
}

type CreateAssetVersionParams struct {
	AssetID string
	Body    schema.AssetVersionCreateRequest
}

// CreateAssetVersion POST /assets/{asset_id}/versions
func (s *LibraryService) CreateAssetVersion(ctx context.Context, params CreateAssetVersionParams) (*Response[schema.LibraryAssetVersionResponse], error) {
	r := &request{method: "POST", path: "/assets/{asset_id}/versions"}
	r.pathParams = map[string]string{
		"asset_id": fmt.Sprint(params.AssetID),
	}
	r.body = params.Body
	return do[schema.LibraryAssetVersionResponse](ctx, s.client, r)
}

type CreateFileUrlsParams struct {
	Body schema.FileURLBulkCreateRequest
}

// CreateFileUrls POST /file-urls
func (s *LibraryService) CreateFileUrls(ctx context.Context, params CreateFileUrlsParams) (*Response[schema.BatchFileURLResponse], error) {
	r := &request{method: "POST", path: "/file-urls"}
	r.body = params.Body
	return do[schema.BatchFileURLResponse](ctx, s.client, r)
}

type CreateFolderParams struct {
	Body schema.FolderCreateRequest
}

// CreateFolder POST /folders
func (s *LibraryService) CreateFolder(ctx context.Context, params CreateFolderParams) (*Response[schema.FolderResponse], error) {
	r := &request{method: "POST", path: "/folders"}
	r.body = params.Body
	return do[schema.FolderResponse](ctx, s.client, r)
}

type CreateStructuredContentParams struct {
	Body schema.LibraryStructuredContentCreateRequest
}

// CreateStructuredContent POST /structured-contents
func (s *LibraryService) CreateStructuredContent(ctx context.Context, params CreateStructuredContentParams) (*Response[schema.LibraryStructuredContent], error) {
	r := &request{method: "POST", path: "/structured-contents"}
	r.body = params.Body
	return do[schema.LibraryStructuredContent](ctx, s.client, r)
}

type DeleteAssetLineageParams struct {
	AssetID   string
	LineageID string
}

// DeleteAssetLineage DELETE /assets/{asset_id}/lineages/{lineage_id}
func (s *LibraryService) DeleteAssetLineage(ctx context.Context, params DeleteAssetLineageParams) (*Response[any], error) {
	r := &request{method: "DELETE", path: "/assets/{asset_id}/lineages/{lineage_id}"}
	r.pathParams = map[string]string{
		"asset_id":   fmt.Sprint(params.AssetID),
		"lineage_id": fmt.Sprint(params.LineageID),
	}
	return do[any](ctx, s.client, r)
}

type DeleteFolderParams struct {
	ID string
}

// DeleteFolder DELETE /folders/{id}
func (s *LibraryService) DeleteFolder(ctx context.Context, params DeleteFolderParams) (*Response[any], error) {
	r := &request{method: "DELETE", path: "/folders/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[any](ctx, s.client, r)
}

type DeleteImageParams struct {
	ID string
}

// DeleteImage DELETE /images/{id}
func (s *LibraryService) DeleteImage(ctx context.Context, params DeleteImageParams) (*Response[any], error) {
	r := &request{method: "DELETE", path: "/images/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[any](ctx, s.client, r)
}

type DeleteRawFileParams struct {
	ID string
}

// DeleteRawFile DELETE /raw-files/{id}
func (s *LibraryService) DeleteRawFile(ctx context.Context, params DeleteRawFileParams) (*Response[any], error) {
	r := &request{method: "DELETE", path: "/raw-files/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[any](ctx, s.client, r)
}

type DeleteVideoParams struct {
	ID string
}

// DeleteVideo DELETE /videos/{id}
func (s *LibraryService) DeleteVideo(ctx context.Context, params DeleteVideoParams) (*Response[any], error) {
	r := &request{method: "DELETE", path: "/videos/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[any](ctx, s.client, r)
}

type GetArticleParams struct {
	ID string
}

// GetArticle GET /articles/{id}
func (s *LibraryService) GetArticle(ctx context.Context, params GetArticleParams) (*Response[schema.LibraryArticle], error) {
	r := &request{method: "GET", path: "/articles/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.LibraryArticle](ctx, s.client, r)
}

type GetFolderParams struct {
	ID string
}

// GetFolder GET /folders/{id}
func (s *LibraryService) GetFolder(ctx context.Context, params GetFolderParams) (*Response[schema.FolderResponse], error) {
	r := &request{method: "GET", path: "/folders/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.FolderResponse](ctx, s.client, r)
}

type GetImageParams struct {
	ID string
}

// GetImage GET /images/{id}
func (s *LibraryService) GetImage(ctx context.Context, params GetImageParams) (*Response[schema.LibraryImage], error) {
	r := &request{method: "GET", path: "/images/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.LibraryImage](ctx, s.client, r)
}

type GetRawFileParams struct {
	ID string
}

// GetRawFile GET /raw-files/{id}
func (s *LibraryService) GetRawFile(ctx context.Context, params GetRawFileParams) (*Response[schema.LibraryRawFile], error) {
	r := &request{method: "GET", path: "/raw-files/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.LibraryRawFile](ctx, s.client, r)
}

type GetRenditionParams struct {
	ID string
}

// GetRendition GET /renditions/{id}
func (s *LibraryService) GetRendition(ctx context.Context, params GetRenditionParams) (*Response[schema.DetailedAssetRenditionResponse], error) {
	r := &request{method: "GET", path: "/renditions/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.DetailedAssetRenditionResponse](ctx, s.client, r)
}

type GetRenditionConfigParams struct {
	ID string
}

// GetRenditionConfig GET /rendition-configs/{id}
func (s *LibraryService) GetRenditionConfig(ctx context.Context, params GetRenditionConfigParams) (*Response[schema.RenditionConfigResponse], error) {
	r := &request{method: "GET", path: "/rendition-configs/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.RenditionConfigResponse](ctx, s.client, r)
}

type GetStructuredContentParams struct {
	ID string
}

// GetStructuredContent GET /structured-contents/{id}
func (s *LibraryService) GetStructuredContent(ctx context.Context, params GetStructuredContentParams) (*Response[schema.LibraryStructuredContent], error) {
	r := &request{method: "GET", path: "/structured-contents/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.LibraryStructuredContent](ctx, s.client, r)
}

type GetVideoParams struct {
	ID string
}

// GetVideo GET /videos/{id}
func (s *LibraryService) GetVideo(ctx context.Context, params GetVideoParams) (*Response[schema.LibraryVideo], error) {
	r := &request{method: "GET", path: "/videos/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.LibraryVideo](ctx, s.client, r)
}

type ListAssetFieldsParams struct {
	AssetID  string
	Offset   *int
	PageSize *int
}

// ListAssetFields GET assets/{asset_id}/fields
func (s *LibraryService) ListAssetFields(ctx context.Context, params ListAssetFieldsParams) (*Response[schema.ListAssetFieldsResponse], error) {
	r := &request{method: "GET", path: "/assets/{asset_id}/fields"}
	r.pathParams = map[string]string{
		"asset_id": fmt.Sprint(params.AssetID),
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
	return do[schema.ListAssetFieldsResponse](ctx, s.client, r)
}

type ListAssetLineagesParams struct {
	AssetID *string
	UsedIn  *string
	// Date and time as the lower limit to filter asset lineages by `created_at`, in ISO 8601 UTC format
	CreatedAtFrom *string
	// Date and time as the upper limit to filter asset lineages by `created_at`, in ISO 8601 UTC format
	CreatedAtTo *string
	Offset      *int
	PageSize    *int
}

// ListAssetLineages GET /asset-lineages
func (s *LibraryService) ListAssetLineages(ctx context.Context, params ListAssetLineagesParams) (*Response[schema.ListAssetLineagesResponse], error) {
	r := &request{method: "GET", path: "/asset-lineages"}
	q := url.Values{}
	if params.AssetID != nil {
		q.Set("asset_id", fmt.Sprint(*params.AssetID))
	}
	if params.UsedIn != nil {
		q.Set("used_in", fmt.Sprint(*params.UsedIn))
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
	return do[schema.ListAssetLineagesResponse](ctx, s.client, r)
}

type ListAssetPermissionsParams struct {
	AssetID   string
	Access    *string
	MaxAccess *string
	MinAccess *string
}

// ListAssetPermissions GET /assets/{asset_id}/permissions
func (s *LibraryService) ListAssetPermissions(ctx context.Context, params ListAssetPermissionsParams) (*Response[schema.AssetPermissionListResponseItem], error) {
	r := &request{method: "GET", path: "/assets/{asset_id}/permissions"}
	r.pathParams = map[string]string{
		"asset_id": fmt.Sprint(params.AssetID),
	}
	q := url.Values{}
	if params.Access != nil {
		q.Set("access", fmt.Sprint(*params.Access))
	}
	if params.MaxAccess != nil {
		q.Set("max_access", fmt.Sprint(*params.MaxAccess))
	}
	if params.MinAccess != nil {
		q.Set("min_access", fmt.Sprint(*params.MinAccess))
	}
	if len(q) > 0 {
		r.query = q
	}
	return do[schema.AssetPermissionListResponseItem](ctx, s.client, r)
}

type ListAssetRenditionsParams struct {
	AssetID  string
	Offset   *int
	PageSize *int
}

// ListAssetRenditions GET /assets/{asset_id}/renditions
func (s *LibraryService) ListAssetRenditions(ctx context.Context, params ListAssetRenditionsParams) (*Response[schema.ListAssetRenditionsResponse], error) {
	r := &request{method: "GET", path: "/assets/{asset_id}/renditions"}
	r.pathParams = map[string]string{
		"asset_id": fmt.Sprint(params.AssetID),
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
	return do[schema.ListAssetRenditionsResponse](ctx, s.client, r)
}

type ListAssetsParams struct {
	// Asset type to filter by. Example: `type=image&type=video`
	Type []string
	// Label to filter by. **Must be passed as urlencoded string**. Labels that do not exist are ignored. If none of the provided labels exist, filtering is not applied. Example – `label=%7B%22group%22%3A%22ee63e3ee43925bb5cc8cd17b817d93ee%22%2C%22values%22%3A%5B%226706efc7828cd6aaedbc0434139cd3e1%22%2C%221f32651216cf2aefcaa08be1ea7dedf1%22%5D%7D`
	Label []string
	// List of fields to filter by. **Must be passed as base64 encoded string**. Fields that do not exist are ignored. If none of the provided fields exist, filtering is not applied. Example – `fields=Wwp7CiJpZCI6ICI2N2E4NDZhMWM3NzU1YTFwNThpNjh5MzVhIiwKInZhbHVlcyI6IFsiNjdhODQ2YTFjNzc1YWU1YWExYTE0YTA1Il0KfQpd=`
	Fields []string
	// Date and time as the lower limit to filter assets by `created_at`, in ISO 8601 UTC format
	CreatedAtFrom *string
	// Date and time as the upper limit to filter assets by `created_at`, in ISO 8601 UTC format
	CreatedAtTo *string
	// Date and time as the lower limit to filter assets by `modified_at`, in ISO 8601 UTC format
	ModifiedAtFrom *string
	// Date and time as the upper limit to filter assets by `modified_at`, in ISO 8601 UTC format
	ModifiedAtTo *string
	// ID of the library folder to include assets from
	FolderID *string
	// Indicates whether assets from subfolders need to be included
	IncludeSubfolderAssets *bool
	// Search assets by title or content description
	SearchText *string
	// ID of the campaign to include assets from
	CampaignID *string
	Offset     *int
	PageSize   *int
}

// ListAssets GET /assets
func (s *LibraryService) ListAssets(ctx context.Context, params ListAssetsParams) (*Response[schema.ListAssetsResponse], error) {
	r := &request{method: "GET", path: "/assets"}
	q := url.Values{}
	for _, v := range params.Type {
		q.Add("type", fmt.Sprint(v))
	}
	for _, v := range params.Label {
		q.Add("label", fmt.Sprint(v))
	}
	for _, v := range params.Fields {
		q.Add("fields", fmt.Sprint(v))
	}
	if params.CreatedAtFrom != nil {
		q.Set("created_at__from", fmt.Sprint(*params.CreatedAtFrom))
	}
	if params.CreatedAtTo != nil {
		q.Set("created_at__to", fmt.Sprint(*params.CreatedAtTo))
	}
	if params.ModifiedAtFrom != nil {
		q.Set("modified_at__from", fmt.Sprint(*params.ModifiedAtFrom))
	}
	if params.ModifiedAtTo != nil {
		q.Set("modified_at__to", fmt.Sprint(*params.ModifiedAtTo))
	}
	if params.FolderID != nil {
		q.Set("folder_id", fmt.Sprint(*params.FolderID))
	}
	if params.IncludeSubfolderAssets != nil {
		q.Set("include_subfolder_assets", fmt.Sprint(*params.IncludeSubfolderAssets))
	}
	if params.SearchText != nil {
		q.Set("search_text", fmt.Sprint(*params.SearchText))
	}
	if params.CampaignID != nil {
		q.Set("campaign_id", fmt.Sprint(*params.CampaignID))
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
	return do[schema.ListAssetsResponse](ctx, s.client, r)
}

type ListFolderPermissionsParams struct {
	ID        string
	Access    *string
	MaxAccess *string
	MinAccess *string
}

// ListFolderPermissions GET /folders/{id}/permissions
func (s *LibraryService) ListFolderPermissions(ctx context.Context, params ListFolderPermissionsParams) (*Response[schema.FolderPermissionListResponseItem], error) {
	r := &request{method: "GET", path: "/folders/{id}/permissions"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	q := url.Values{}
	if params.Access != nil {
		q.Set("access", fmt.Sprint(*params.Access))
	}
	if params.MaxAccess != nil {
		q.Set("max_access", fmt.Sprint(*params.MaxAccess))
	}
	if params.MinAccess != nil {
		q.Set("min_access", fmt.Sprint(*params.MinAccess))
	}
	if len(q) > 0 {
		r.query = q
	}
	return do[schema.FolderPermissionListResponseItem](ctx, s.client, r)
}

type ListFoldersParams struct {
	// ID of the parent folder to filter by
	ParentFolderID *string
	Offset         *int
	PageSize       *int
}

// ListFolders GET /folders
func (s *LibraryService) ListFolders(ctx context.Context, params ListFoldersParams) (*Response[schema.ListFoldersResponse], error) {
	r := &request{method: "GET", path: "/folders"}
	q := url.Values{}
	if params.ParentFolderID != nil {
		q.Set("parent_folder_id", fmt.Sprint(*params.ParentFolderID))
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
	return do[schema.ListFoldersResponse](ctx, s.client, r)
}

type ListRelatedAssetsParams struct {
	AssetID string
}

// ListRelatedAssets GET /assets/{asset_id}/related-assets
func (s *LibraryService) ListRelatedAssets(ctx context.Context, params ListRelatedAssetsParams) (*Response[schema.RelatedAssetsListResponse], error) {
	r := &request{method: "GET", path: "/assets/{asset_id}/related-assets"}
	r.pathParams = map[string]string{
		"asset_id": fmt.Sprint(params.AssetID),
	}
	return do[schema.RelatedAssetsListResponse](ctx, s.client, r)
}

type RemoveAssetPermissionParams struct {
	AssetID    string
	AccessorID string
}

// RemoveAssetPermission DELETE /asstes/{asset_id}/permissions/{accessor_id}
func (s *LibraryService) RemoveAssetPermission(ctx context.Context, params RemoveAssetPermissionParams) (*Response[any], error) {
	r := &request{method: "DELETE", path: "/assets/{asset_id}/permissions/{accessor_id}"}
	r.pathParams = map[string]string{
		"asset_id":    fmt.Sprint(params.AssetID),
		"accessor_id": fmt.Sprint(params.AccessorID),
	}
	return do[any](ctx, s.client, r)
}

type RemoveFolderPermissionParams struct {
	ID         string
	AccessorID string
}

// RemoveFolderPermission DELETE /folders/{id}/permissions/{accessor_id}
func (s *LibraryService) RemoveFolderPermission(ctx context.Context, params RemoveFolderPermissionParams) (*Response[any], error) {
	r := &request{method: "DELETE", path: "/folders/{id}/permissions/{accessor_id}"}
	r.pathParams = map[string]string{
		"id":          fmt.Sprint(params.ID),
		"accessor_id": fmt.Sprint(params.AccessorID),
	}
	return do[any](ctx, s.client, r)
}

type ReplaceRelatedAssetsParams struct {
	AssetID string
	Body    schema.ReplaceRelatedAssetsRequest
}

// ReplaceRelatedAssets PUT /assets/{asset_id}/related-assets
func (s *LibraryService) ReplaceRelatedAssets(ctx context.Context, params ReplaceRelatedAssetsParams) (*Response[schema.ReplaceRelatedAssetsResponse], error) {
	r := &request{method: "PUT", path: "/assets/{asset_id}/related-assets"}
	r.pathParams = map[string]string{
		"asset_id": fmt.Sprint(params.AssetID),
	}
	r.body = params.Body
	return do[schema.ReplaceRelatedAssetsResponse](ctx, s.client, r)
}

type UpdateAssetFieldParams struct {
	AssetID string
	FieldID string
	Body    schema.AssetFieldUpdateRequest
}

// UpdateAssetField PUT /assets/{asset_id}/fields/{field_id}
func (s *LibraryService) UpdateAssetField(ctx context.Context, params UpdateAssetFieldParams) (*Response[schema.AssetFieldUpdateResponse], error) {
	r := &request{method: "PUT", path: "/assets/{asset_id}/fields/{field_id}"}
	r.pathParams = map[string]string{
		"asset_id": fmt.Sprint(params.AssetID),
		"field_id": fmt.Sprint(params.FieldID),
	}
	r.body = params.Body
	return do[schema.AssetFieldUpdateResponse](ctx, s.client, r)
}

type UpdateAssetFieldsParams struct {
	AssetID string
	Body    schema.AssetFieldsUpdateRequest
}

// UpdateAssetFields PUT /assets/{asset_id}/fields
func (s *LibraryService) UpdateAssetFields(ctx context.Context, params UpdateAssetFieldsParams) (*Response[schema.UpdateAssetFieldsResponse], error) {
	r := &request{method: "PUT", path: "/assets/{asset_id}/fields"}
	r.pathParams = map[string]string{
		"asset_id": fmt.Sprint(params.AssetID),
	}
	r.body = params.Body
	return do[schema.UpdateAssetFieldsResponse](ctx, s.client, r)
}

type UpdateAssetPermissionParams struct {
	AssetID    string
	AccessorID string
	Body       schema.AssetPermissionUpdateRequest
}

// UpdateAssetPermission PATCH /assets/{asset_id}/permissions/{accessor_id}
func (s *LibraryService) UpdateAssetPermission(ctx context.Context, params UpdateAssetPermissionParams) (*Response[any], error) {
	r := &request{method: "PATCH", path: "/assets/{asset_id}/permissions/{accessor_id}"}
	r.pathParams = map[string]string{
		"asset_id":    fmt.Sprint(params.AssetID),
		"accessor_id": fmt.Sprint(params.AccessorID),
	}
	r.body = params.Body
	return do[any](ctx, s.client, r)
}

type UpdateFolderParams struct {
	ID   string
	Body schema.FolderUpdateRequest
}

// UpdateFolder PATCH /folders/{id}
func (s *LibraryService) UpdateFolder(ctx context.Context, params UpdateFolderParams) (*Response[schema.FolderResponse], error) {
	r := &request{method: "PATCH", path: "/folders/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.FolderResponse](ctx, s.client, r)
}

type UpdateFolderPermissionParams struct {
	ID         string
	AccessorID string
	Body       schema.FolderPermissionUpdateRequest
}

// UpdateFolderPermission PATCH /folders/{id}/permissions/{accessor_id}
func (s *LibraryService) UpdateFolderPermission(ctx context.Context, params UpdateFolderPermissionParams) (*Response[any], error) {
	r := &request{method: "PATCH", path: "/folders/{id}/permissions/{accessor_id}"}
	r.pathParams = map[string]string{
		"id":          fmt.Sprint(params.ID),
		"accessor_id": fmt.Sprint(params.AccessorID),
	}
	r.body = params.Body
	return do[any](ctx, s.client, r)
}

type UpdateImageParams struct {
	ID   string
	Body schema.LibraryImageUpdateRequest
}

// UpdateImage PATCH /images/{id}
func (s *LibraryService) UpdateImage(ctx context.Context, params UpdateImageParams) (*Response[schema.LibraryImage], error) {
	r := &request{method: "PATCH", path: "/images/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.LibraryImage](ctx, s.client, r)
}

type UpdateRawFileParams struct {
	ID   string
	Body schema.LibraryRawFileUpdateRequest
}

// UpdateRawFile PATCH /raw-files/{id}
func (s *LibraryService) UpdateRawFile(ctx context.Context, params UpdateRawFileParams) (*Response[schema.LibraryRawFile], error) {
	r := &request{method: "PATCH", path: "/raw-files/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.LibraryRawFile](ctx, s.client, r)
}

type UpdateStructuredContentParams struct {
	ID   string
	Body schema.LibraryStructuredContentUpdateRequest
}

// UpdateStructuredContent PATCH /structured-contents/{id}
func (s *LibraryService) UpdateStructuredContent(ctx context.Context, params UpdateStructuredContentParams) (*Response[schema.LibraryStructuredContent], error) {
	r := &request{method: "PATCH", path: "/structured-contents/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.LibraryStructuredContent](ctx, s.client, r)
}

type UpdateVideoParams struct {
	ID   string
	Body schema.LibraryVideoUpdateRequest
}

// UpdateVideo PATCH /videos/{id}
func (s *LibraryService) UpdateVideo(ctx context.Context, params UpdateVideoParams) (*Response[schema.LibraryVideo], error) {
	r := &request{method: "PATCH", path: "/videos/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.LibraryVideo](ctx, s.client, r)
}
