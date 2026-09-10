// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type TaskService struct{ client *Client }

type AddAssetToTaskParams struct {
	ID   string
	Body schema.TaskAssetRequest
}

// AddAssetToTask POST /tasks/{id}/assets
func (s *TaskService) AddAssetToTask(ctx context.Context, params AddAssetToTaskParams) (*Response[schema.TaskAssetResponse], error) {
	r := &request{method: "POST", path: "/tasks/{id}/assets"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.TaskAssetResponse](ctx, s.client, r)
}

type AddCommentToTaskParams struct {
	TaskID string
	Body   schema.CommentCreateRequest
}

// AddCommentToTask POST /tasks/{task_id}/comments
func (s *TaskService) AddCommentToTask(ctx context.Context, params AddCommentToTaskParams) (*Response[schema.TaskCommentResponse], error) {
	r := &request{method: "POST", path: "/tasks/{task_id}/comments"}
	r.pathParams = map[string]string{
		"task_id": fmt.Sprint(params.TaskID),
	}
	r.body = params.Body
	return do[schema.TaskCommentResponse](ctx, s.client, r)
}

type AddCommentToTaskAssetParams struct {
	TaskID  string
	AssetID string
	Body    schema.CommentCreateRequest
}

// AddCommentToTaskAsset POST /tasks/{task_id}/assets/{asset_id}/comments
func (s *TaskService) AddCommentToTaskAsset(ctx context.Context, params AddCommentToTaskAssetParams) (*Response[schema.TaskAssetCommentResponse], error) {
	r := &request{method: "POST", path: "/tasks/{task_id}/assets/{asset_id}/comments"}
	r.pathParams = map[string]string{
		"task_id":  fmt.Sprint(params.TaskID),
		"asset_id": fmt.Sprint(params.AssetID),
	}
	r.body = params.Body
	return do[schema.TaskAssetCommentResponse](ctx, s.client, r)
}

type AddCommentToTaskSubStepParams struct {
	TaskID    string
	StepID    string
	SubStepID string
	Body      schema.CommentWithReplyCreateRequest
}

// AddCommentToTaskSubStep POST /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments
func (s *TaskService) AddCommentToTaskSubStep(ctx context.Context, params AddCommentToTaskSubStepParams) (*Response[schema.TaskSubStepCommentResponse], error) {
	r := &request{method: "POST", path: "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments"}
	r.pathParams = map[string]string{
		"task_id":     fmt.Sprint(params.TaskID),
		"step_id":     fmt.Sprint(params.StepID),
		"sub_step_id": fmt.Sprint(params.SubStepID),
	}
	r.body = params.Body
	return do[schema.TaskSubStepCommentResponse](ctx, s.client, r)
}

type AddDraftToTaskAssetParams struct {
	TaskID  string
	AssetID string
	Body    schema.TaskAssetRequest
}

// AddDraftToTaskAsset POST /tasks/{task_id}/assets/{asset_id}/drafts
func (s *TaskService) AddDraftToTaskAsset(ctx context.Context, params AddDraftToTaskAssetParams) (*Response[schema.TaskAssetDraftResponse], error) {
	r := &request{method: "POST", path: "/tasks/{task_id}/assets/{asset_id}/drafts"}
	r.pathParams = map[string]string{
		"task_id":  fmt.Sprint(params.TaskID),
		"asset_id": fmt.Sprint(params.AssetID),
	}
	r.body = params.Body
	return do[schema.TaskAssetDraftResponse](ctx, s.client, r)
}

type AddFieldToTaskParams struct {
	TaskID string
	Body   schema.ObjectFieldCreateRequest
}

// AddFieldToTask POST tasks/{task_id}/fields
func (s *TaskService) AddFieldToTask(ctx context.Context, params AddFieldToTaskParams) (*Response[schema.ObjectFieldCreateResponse], error) {
	r := &request{method: "POST", path: "/tasks/{task_id}/fields"}
	r.pathParams = map[string]string{
		"task_id": fmt.Sprint(params.TaskID),
	}
	r.body = params.Body
	return do[schema.ObjectFieldCreateResponse](ctx, s.client, r)
}

type AddStructuredContentToTaskParams struct {
	TaskID string
	Body   schema.TaskStructuredContentCreateRequest
}

// AddStructuredContentToTask POST /tasks/{task_id}/structured-contents
func (s *TaskService) AddStructuredContentToTask(ctx context.Context, params AddStructuredContentToTaskParams) (*Response[schema.ContentDetailsModel], error) {
	r := &request{method: "POST", path: "/tasks/{task_id}/structured-contents"}
	r.pathParams = map[string]string{
		"task_id": fmt.Sprint(params.TaskID),
	}
	r.body = params.Body
	return do[schema.ContentDetailsModel](ctx, s.client, r)
}

type AddURLToTaskParams struct {
	TaskID string
	Body   schema.AddURLToTaskRequest
}

// AddURLToTask POST /tasks/{task_id}/urls
func (s *TaskService) AddURLToTask(ctx context.Context, params AddURLToTaskParams) (*Response[schema.TaskURLResponse], error) {
	r := &request{method: "POST", path: "/tasks/{task_id}/urls"}
	r.pathParams = map[string]string{
		"task_id": fmt.Sprint(params.TaskID),
	}
	r.body = params.Body
	return do[schema.TaskURLResponse](ctx, s.client, r)
}

type CreateTaskParams struct {
	Body schema.TaskCreateRequest
}

// CreateTask POST /tasks
func (s *TaskService) CreateTask(ctx context.Context, params CreateTaskParams) (*Response[schema.TaskResponse], error) {
	r := &request{method: "POST", path: "/tasks"}
	r.body = params.Body
	return do[schema.TaskResponse](ctx, s.client, r)
}

type CreateTaskPublishingIntentParams struct {
	TaskID string
	Body   schema.TaskPublishingIntentCreateRequest
}

// CreateTaskPublishingIntent POST /tasks/{task_id}/publishing-intents
func (s *TaskService) CreateTaskPublishingIntent(ctx context.Context, params CreateTaskPublishingIntentParams) (*Response[schema.TaskPublishingIntentResponse], error) {
	r := &request{method: "POST", path: "/tasks/{task_id}/publishing-intents"}
	r.pathParams = map[string]string{
		"task_id": fmt.Sprint(params.TaskID),
	}
	r.body = params.Body
	return do[schema.TaskPublishingIntentResponse](ctx, s.client, r)
}

type CreateTaskStructuredContentDraftParams struct {
	TaskID    string
	ContentID string
	Body      schema.TaskStructuredContentDraftRequest
}

// CreateTaskStructuredContentDraft POST /tasks/{task_id}/structured-contents/{content_id}/drafts
func (s *TaskService) CreateTaskStructuredContentDraft(ctx context.Context, params CreateTaskStructuredContentDraftParams) (*Response[schema.CreateTaskStructuredContentDraftResponse], error) {
	r := &request{method: "POST", path: "/tasks/{task_id}/structured-contents/{content_id}/drafts"}
	r.pathParams = map[string]string{
		"task_id":    fmt.Sprint(params.TaskID),
		"content_id": fmt.Sprint(params.ContentID),
	}
	r.body = params.Body
	return do[schema.CreateTaskStructuredContentDraftResponse](ctx, s.client, r)
}

type DeleteTaskStructuredContentParams struct {
	TaskID    string
	ContentID string
}

// DeleteTaskStructuredContent DELETE /tasks/{task_id}/structured-contents/{content_id}
func (s *TaskService) DeleteTaskStructuredContent(ctx context.Context, params DeleteTaskStructuredContentParams) (*Response[any], error) {
	r := &request{method: "DELETE", path: "/tasks/{task_id}/structured-contents/{content_id}"}
	r.pathParams = map[string]string{
		"task_id":    fmt.Sprint(params.TaskID),
		"content_id": fmt.Sprint(params.ContentID),
	}
	return do[any](ctx, s.client, r)
}

type DeleteTaskSubStepCommentParams struct {
	TaskID    string
	StepID    string
	SubStepID string
	CommentID string
}

// DeleteTaskSubStepComment DELETE /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}
func (s *TaskService) DeleteTaskSubStepComment(ctx context.Context, params DeleteTaskSubStepCommentParams) (*Response[any], error) {
	r := &request{method: "DELETE", path: "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}"}
	r.pathParams = map[string]string{
		"task_id":     fmt.Sprint(params.TaskID),
		"step_id":     fmt.Sprint(params.StepID),
		"sub_step_id": fmt.Sprint(params.SubStepID),
		"comment_id":  fmt.Sprint(params.CommentID),
	}
	return do[any](ctx, s.client, r)
}

type GetTaskParams struct {
	ID string
}

// GetTask GET /tasks/{id}
func (s *TaskService) GetTask(ctx context.Context, params GetTaskParams) (*Response[schema.TaskResponse], error) {
	r := &request{method: "GET", path: "/tasks/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.TaskResponse](ctx, s.client, r)
}

type GetTaskArticleParams struct {
	TaskID    string
	ArticleID string
}

// GetTaskArticle GET /tasks/{task_id}/articles/{article_id}
func (s *TaskService) GetTaskArticle(ctx context.Context, params GetTaskArticleParams) (*Response[schema.TaskArticle], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/articles/{article_id}"}
	r.pathParams = map[string]string{
		"task_id":    fmt.Sprint(params.TaskID),
		"article_id": fmt.Sprint(params.ArticleID),
	}
	return do[schema.TaskArticle](ctx, s.client, r)
}

type GetTaskBriefParams struct {
	ID string
}

// GetTaskBrief GET /tasks/{id}/brief
func (s *TaskService) GetTaskBrief(ctx context.Context, params GetTaskBriefParams) (*Response[schema.TaskBriefResponse], error) {
	r := &request{method: "GET", path: "/tasks/{id}/brief"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.TaskBriefResponse](ctx, s.client, r)
}

type GetTaskCustomFieldParams struct {
	TaskID        string
	CustomFieldID string
}

// GetTaskCustomField GET /tasks/{task_id}/custom-fields/{custom_field_id}
func (s *TaskService) GetTaskCustomField(ctx context.Context, params GetTaskCustomFieldParams) (*Response[schema.TaskCustomField], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/custom-fields/{custom_field_id}"}
	r.pathParams = map[string]string{
		"task_id":         fmt.Sprint(params.TaskID),
		"custom_field_id": fmt.Sprint(params.CustomFieldID),
	}
	return do[schema.TaskCustomField](ctx, s.client, r)
}

type GetTaskImageParams struct {
	TaskID  string
	ImageID string
}

// GetTaskImage GET /tasks/{task_id}/images/{image_id}
func (s *TaskService) GetTaskImage(ctx context.Context, params GetTaskImageParams) (*Response[schema.TaskImage], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/images/{image_id}"}
	r.pathParams = map[string]string{
		"task_id":  fmt.Sprint(params.TaskID),
		"image_id": fmt.Sprint(params.ImageID),
	}
	return do[schema.TaskImage](ctx, s.client, r)
}

type GetTaskRawFileParams struct {
	TaskID    string
	RawFileID string
}

// GetTaskRawFile GET /tasks/{task_id}/raw-files/{raw_file_id}
func (s *TaskService) GetTaskRawFile(ctx context.Context, params GetTaskRawFileParams) (*Response[schema.TaskRawFile], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/raw-files/{raw_file_id}"}
	r.pathParams = map[string]string{
		"task_id":     fmt.Sprint(params.TaskID),
		"raw_file_id": fmt.Sprint(params.RawFileID),
	}
	return do[schema.TaskRawFile](ctx, s.client, r)
}

type GetTaskStructuredContentParams struct {
	TaskID    string
	ContentID string
}

// GetTaskStructuredContent GET /tasks/{task_id}/structured-contents/{content_id}
func (s *TaskService) GetTaskStructuredContent(ctx context.Context, params GetTaskStructuredContentParams) (*Response[schema.ContentDetailsModel], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/structured-contents/{content_id}"}
	r.pathParams = map[string]string{
		"task_id":    fmt.Sprint(params.TaskID),
		"content_id": fmt.Sprint(params.ContentID),
	}
	return do[schema.ContentDetailsModel](ctx, s.client, r)
}

type GetTaskSubStepParams struct {
	TaskID    string
	StepID    string
	SubStepID string
}

// GetTaskSubStep GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}
func (s *TaskService) GetTaskSubStep(ctx context.Context, params GetTaskSubStepParams) (*Response[schema.TaskSubStep], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}"}
	r.pathParams = map[string]string{
		"task_id":     fmt.Sprint(params.TaskID),
		"step_id":     fmt.Sprint(params.StepID),
		"sub_step_id": fmt.Sprint(params.SubStepID),
	}
	return do[schema.TaskSubStep](ctx, s.client, r)
}

type GetTaskSubStepCommentParams struct {
	TaskID    string
	StepID    string
	SubStepID string
	CommentID string
}

// GetTaskSubStepComment GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}
func (s *TaskService) GetTaskSubStepComment(ctx context.Context, params GetTaskSubStepCommentParams) (*Response[schema.TaskSubStepCommentResponse], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}"}
	r.pathParams = map[string]string{
		"task_id":     fmt.Sprint(params.TaskID),
		"step_id":     fmt.Sprint(params.StepID),
		"sub_step_id": fmt.Sprint(params.SubStepID),
		"comment_id":  fmt.Sprint(params.CommentID),
	}
	return do[schema.TaskSubStepCommentResponse](ctx, s.client, r)
}

type GetTaskSubStepExternalWorkParams struct {
	TaskID    string
	StepID    string
	SubStepID string
}

// GetTaskSubStepExternalWork GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/external-work
func (s *TaskService) GetTaskSubStepExternalWork(ctx context.Context, params GetTaskSubStepExternalWorkParams) (*Response[schema.TaskExternalWorkResponse], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/external-work"}
	r.pathParams = map[string]string{
		"task_id":     fmt.Sprint(params.TaskID),
		"step_id":     fmt.Sprint(params.StepID),
		"sub_step_id": fmt.Sprint(params.SubStepID),
	}
	return do[schema.TaskExternalWorkResponse](ctx, s.client, r)
}

type GetTaskVideoParams struct {
	TaskID  string
	VideoID string
}

// GetTaskVideo GET /tasks/{task_id}/videos/{video_id}
func (s *TaskService) GetTaskVideo(ctx context.Context, params GetTaskVideoParams) (*Response[schema.TaskVideo], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/videos/{video_id}"}
	r.pathParams = map[string]string{
		"task_id":  fmt.Sprint(params.TaskID),
		"video_id": fmt.Sprint(params.VideoID),
	}
	return do[schema.TaskVideo](ctx, s.client, r)
}

type ListTaskAssetCommentsParams struct {
	TaskID   string
	AssetID  string
	Offset   *int
	PageSize *int
}

// ListTaskAssetComments GET /tasks/{task_id}/assets/{asset_id}/comments
func (s *TaskService) ListTaskAssetComments(ctx context.Context, params ListTaskAssetCommentsParams) (*Response[schema.ListTaskAssetCommentsResponse], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/assets/{asset_id}/comments"}
	r.pathParams = map[string]string{
		"task_id":  fmt.Sprint(params.TaskID),
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
	return do[schema.ListTaskAssetCommentsResponse](ctx, s.client, r)
}

type ListTaskAssetDraftsParams struct {
	TaskID   string
	AssetID  string
	Offset   *int
	PageSize *int
}

// ListTaskAssetDrafts GET /tasks/{task_id}/assets/{asset_id}/drafts
func (s *TaskService) ListTaskAssetDrafts(ctx context.Context, params ListTaskAssetDraftsParams) (*Response[schema.ListTaskAssetDraftsResponse], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/assets/{asset_id}/drafts"}
	r.pathParams = map[string]string{
		"task_id":  fmt.Sprint(params.TaskID),
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
	return do[schema.ListTaskAssetDraftsResponse](ctx, s.client, r)
}

type ListTaskAssetFieldsParams struct {
	TaskID   string
	AssetID  string
	Offset   *int
	PageSize *int
}

// ListTaskAssetFields GET /tasks/{task_id}/assets/{asset_id}/fields
func (s *TaskService) ListTaskAssetFields(ctx context.Context, params ListTaskAssetFieldsParams) (*Response[schema.ListTaskAssetFieldsResponse], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/assets/{asset_id}/fields"}
	r.pathParams = map[string]string{
		"task_id":  fmt.Sprint(params.TaskID),
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
	return do[schema.ListTaskAssetFieldsResponse](ctx, s.client, r)
}

type ListTaskAssetsParams struct {
	ID       string
	Offset   *int
	PageSize *int
}

// ListTaskAssets GET /tasks/{id}/assets
func (s *TaskService) ListTaskAssets(ctx context.Context, params ListTaskAssetsParams) (*Response[schema.ListTaskAssetsResponse], error) {
	r := &request{method: "GET", path: "/tasks/{id}/assets"}
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
	return do[schema.ListTaskAssetsResponse](ctx, s.client, r)
}

type ListTaskAttachmentsParams struct {
	ID       string
	Offset   *int
	PageSize *int
}

// ListTaskAttachments GET /tasks/{id}/attachments
func (s *TaskService) ListTaskAttachments(ctx context.Context, params ListTaskAttachmentsParams) (*Response[schema.ListTaskAttachmentsResponse], error) {
	r := &request{method: "GET", path: "/tasks/{id}/attachments"}
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
	return do[schema.ListTaskAttachmentsResponse](ctx, s.client, r)
}

type ListTaskCustomFieldChoicesParams struct {
	TaskID        string
	CustomFieldID string
	Offset        *int
	PageSize      *int
}

// ListTaskCustomFieldChoices GET /tasks/{task_id}/custom-fields/{custom_field_id}/choices
func (s *TaskService) ListTaskCustomFieldChoices(ctx context.Context, params ListTaskCustomFieldChoicesParams) (*Response[schema.ListTaskCustomFieldChoicesResponse], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/custom-fields/{custom_field_id}/choices"}
	r.pathParams = map[string]string{
		"task_id":         fmt.Sprint(params.TaskID),
		"custom_field_id": fmt.Sprint(params.CustomFieldID),
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
	return do[schema.ListTaskCustomFieldChoicesResponse](ctx, s.client, r)
}

type ListTaskCustomFieldsParams struct {
	ID       string
	Offset   *int
	PageSize *int
}

// ListTaskCustomFields GET tasks/{id}/custom-fields
func (s *TaskService) ListTaskCustomFields(ctx context.Context, params ListTaskCustomFieldsParams) (*Response[schema.ListTaskCustomFieldsResponse], error) {
	r := &request{method: "GET", path: "/tasks/{id}/custom-fields"}
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
	return do[schema.ListTaskCustomFieldsResponse](ctx, s.client, r)
}

type ListTaskFieldsParams struct {
	TaskID   string
	Offset   *int
	PageSize *int
}

// ListTaskFields GET tasks/{task_id}/fields
func (s *TaskService) ListTaskFields(ctx context.Context, params ListTaskFieldsParams) (*Response[schema.ListTaskFieldsResponse], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/fields"}
	r.pathParams = map[string]string{
		"task_id": fmt.Sprint(params.TaskID),
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
	return do[schema.ListTaskFieldsResponse](ctx, s.client, r)
}

type ListTaskSubStepCommentsParams struct {
	TaskID    string
	StepID    string
	SubStepID string
	Offset    *int
	PageSize  *int
}

// ListTaskSubStepComments GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments
func (s *TaskService) ListTaskSubStepComments(ctx context.Context, params ListTaskSubStepCommentsParams) (*Response[schema.ListTaskSubStepCommentsResponse], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments"}
	r.pathParams = map[string]string{
		"task_id":     fmt.Sprint(params.TaskID),
		"step_id":     fmt.Sprint(params.StepID),
		"sub_step_id": fmt.Sprint(params.SubStepID),
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
	return do[schema.ListTaskSubStepCommentsResponse](ctx, s.client, r)
}

type ListTaskSubStepFieldsParams struct {
	TaskID    string
	StepID    string
	SubStepID string
	Offset    *int
	PageSize  *int
}

// ListTaskSubStepFields GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/fields
func (s *TaskService) ListTaskSubStepFields(ctx context.Context, params ListTaskSubStepFieldsParams) (*Response[schema.ListTaskSubStepFieldsResponse], error) {
	r := &request{method: "GET", path: "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/fields"}
	r.pathParams = map[string]string{
		"task_id":     fmt.Sprint(params.TaskID),
		"step_id":     fmt.Sprint(params.StepID),
		"sub_step_id": fmt.Sprint(params.SubStepID),
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
	return do[schema.ListTaskSubStepFieldsResponse](ctx, s.client, r)
}

type ListTasksParams struct {
	SearchKey *string
	Campaign  *string
	Workflow  *string
	Milestone *string
	StartDate *string
	DueDate   *string
	Status    []string
	Offset    *int
	PageSize  *int
}

// ListTasks GET /tasks
func (s *TaskService) ListTasks(ctx context.Context, params ListTasksParams) (*Response[schema.ListTasksResponse], error) {
	r := &request{method: "GET", path: "/tasks"}
	q := url.Values{}
	if params.SearchKey != nil {
		q.Set("search_key", fmt.Sprint(*params.SearchKey))
	}
	if params.Campaign != nil {
		q.Set("campaign", fmt.Sprint(*params.Campaign))
	}
	if params.Workflow != nil {
		q.Set("workflow", fmt.Sprint(*params.Workflow))
	}
	if params.Milestone != nil {
		q.Set("milestone", fmt.Sprint(*params.Milestone))
	}
	if params.StartDate != nil {
		q.Set("start_date", fmt.Sprint(*params.StartDate))
	}
	if params.DueDate != nil {
		q.Set("due_date", fmt.Sprint(*params.DueDate))
	}
	for _, v := range params.Status {
		q.Add("status", fmt.Sprint(v))
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
	return do[schema.ListTasksResponse](ctx, s.client, r)
}

type RemoveTaskFieldParams struct {
	TaskID  string
	FieldID string
}

// RemoveTaskField DELETE tasks/{task_id}/fields/{field_id}
func (s *TaskService) RemoveTaskField(ctx context.Context, params RemoveTaskFieldParams) (*Response[any], error) {
	r := &request{method: "DELETE", path: "/tasks/{task_id}/fields/{field_id}"}
	r.pathParams = map[string]string{
		"task_id":  fmt.Sprint(params.TaskID),
		"field_id": fmt.Sprint(params.FieldID),
	}
	return do[any](ctx, s.client, r)
}

type UpdateTaskParams struct {
	ID   string
	Body schema.TaskUpdateRequest
}

// UpdateTask PATCH /tasks/{id}
func (s *TaskService) UpdateTask(ctx context.Context, params UpdateTaskParams) (*Response[schema.TaskResponse], error) {
	r := &request{method: "PATCH", path: "/tasks/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.TaskResponse](ctx, s.client, r)
}

type UpdateTaskAssetFieldsParams struct {
	TaskID  string
	AssetID string
	Body    schema.TaskAssetFieldsUpdateRequest
}

// UpdateTaskAssetFields PUT tasks/{task_id}/assets/{asset_id}/fields
func (s *TaskService) UpdateTaskAssetFields(ctx context.Context, params UpdateTaskAssetFieldsParams) (*Response[any], error) {
	r := &request{method: "PUT", path: "/tasks/{task_id}/assets/{asset_id}/fields"}
	r.pathParams = map[string]string{
		"task_id":  fmt.Sprint(params.TaskID),
		"asset_id": fmt.Sprint(params.AssetID),
	}
	r.body = params.Body
	return do[any](ctx, s.client, r)
}

type UpdateTaskCustomFieldParams struct {
	TaskID        string
	CustomFieldID string
	Body          schema.TaskCustomFieldUpdateRequest
}

// UpdateTaskCustomField PATCH /tasks/{task_id}/custom-fields/{custom_field_id}
func (s *TaskService) UpdateTaskCustomField(ctx context.Context, params UpdateTaskCustomFieldParams) (*Response[schema.TaskCustomField], error) {
	r := &request{method: "PATCH", path: "/tasks/{task_id}/custom-fields/{custom_field_id}"}
	r.pathParams = map[string]string{
		"task_id":         fmt.Sprint(params.TaskID),
		"custom_field_id": fmt.Sprint(params.CustomFieldID),
	}
	r.body = params.Body
	return do[schema.TaskCustomField](ctx, s.client, r)
}

type UpdateTaskFieldParams struct {
	TaskID  string
	FieldID string
	Body    schema.TaskFieldUpdateRequest
}

// UpdateTaskField PUT tasks/{taks_id}/fields/{field_id}
func (s *TaskService) UpdateTaskField(ctx context.Context, params UpdateTaskFieldParams) (*Response[any], error) {
	r := &request{method: "PUT", path: "/tasks/{task_id}/fields/{field_id}"}
	r.pathParams = map[string]string{
		"task_id":  fmt.Sprint(params.TaskID),
		"field_id": fmt.Sprint(params.FieldID),
	}
	r.body = params.Body
	return do[any](ctx, s.client, r)
}

type UpdateTaskStructuredContentParams struct {
	TaskID    string
	ContentID string
	Body      schema.TaskStructuredContentUpdateRequest
}

// UpdateTaskStructuredContent PATCH /tasks/{task_id}/structured-contents/{content_id}
func (s *TaskService) UpdateTaskStructuredContent(ctx context.Context, params UpdateTaskStructuredContentParams) (*Response[schema.ContentDetailsModel], error) {
	r := &request{method: "PATCH", path: "/tasks/{task_id}/structured-contents/{content_id}"}
	r.pathParams = map[string]string{
		"task_id":    fmt.Sprint(params.TaskID),
		"content_id": fmt.Sprint(params.ContentID),
	}
	r.body = params.Body
	return do[schema.ContentDetailsModel](ctx, s.client, r)
}

type UpdateTaskSubStepParams struct {
	TaskID    string
	StepID    string
	SubStepID string
	Body      schema.TaskSubStepRequest
}

// UpdateTaskSubStep PATCH /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}
func (s *TaskService) UpdateTaskSubStep(ctx context.Context, params UpdateTaskSubStepParams) (*Response[schema.TaskSubStep], error) {
	r := &request{method: "PATCH", path: "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}"}
	r.pathParams = map[string]string{
		"task_id":     fmt.Sprint(params.TaskID),
		"step_id":     fmt.Sprint(params.StepID),
		"sub_step_id": fmt.Sprint(params.SubStepID),
	}
	r.body = params.Body
	return do[schema.TaskSubStep](ctx, s.client, r)
}

type UpdateTaskSubStepCommentParams struct {
	TaskID    string
	StepID    string
	SubStepID string
	CommentID string
	Body      schema.TaskSubStepCommentUpdateRequest
}

// UpdateTaskSubStepComment PATCH /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}
func (s *TaskService) UpdateTaskSubStepComment(ctx context.Context, params UpdateTaskSubStepCommentParams) (*Response[schema.TaskSubStepCommentResponse], error) {
	r := &request{method: "PATCH", path: "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}"}
	r.pathParams = map[string]string{
		"task_id":     fmt.Sprint(params.TaskID),
		"step_id":     fmt.Sprint(params.StepID),
		"sub_step_id": fmt.Sprint(params.SubStepID),
		"comment_id":  fmt.Sprint(params.CommentID),
	}
	r.body = params.Body
	return do[schema.TaskSubStepCommentResponse](ctx, s.client, r)
}

type UpdateTaskSubStepExternalWorkParams struct {
	TaskID    string
	StepID    string
	SubStepID string
	Body      schema.TaskExternalWorkRequest
}

// UpdateTaskSubStepExternalWork PATCH /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/external-work
func (s *TaskService) UpdateTaskSubStepExternalWork(ctx context.Context, params UpdateTaskSubStepExternalWorkParams) (*Response[schema.TaskExternalWorkResponse], error) {
	r := &request{method: "PATCH", path: "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/external-work"}
	r.pathParams = map[string]string{
		"task_id":     fmt.Sprint(params.TaskID),
		"step_id":     fmt.Sprint(params.StepID),
		"sub_step_id": fmt.Sprint(params.SubStepID),
	}
	r.body = params.Body
	return do[schema.TaskExternalWorkResponse](ctx, s.client, r)
}
