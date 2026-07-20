// Auto-generated - DO NOT EDIT

import type * as Schema from "./schema.js";

export interface GetAssetUrlParams {
  asset_id: string;
}

export interface GetTaskAssetDraftBrandComplianceParams {
  task_id: string;
  asset_id: string;
  draft_id: string;
}

export interface ListBrandComplianceCategoriesParams {
  offset?: number;
  page_size?: number;
}

export interface UpdateTaskAssetDraftBrandComplianceParams {
  task_id: string;
  asset_id: string;
  draft_id: string;
  body: Schema.TaskAssetDraftBrandComplianceRequest;
}

export interface AddAttachmentToCampaignParams {
  id: string;
  body: Schema.AttachmentRequest;
}

export interface AddCommentToCampaignParams {
  id: string;
  body: Schema.CommentCreateRequest;
}

export interface AddFieldToCampaignParams {
  id: string;
  body: Schema.ObjectFieldCreateRequest;
}

export interface CreateCampaignParams {
  body: Schema.CampaignCreateRequest;
}

export interface GetCampaignParams {
  id: string;
}

export interface GetCampaignBriefParams {
  id: string;
}

export interface ListCampaignFieldsParams {
  id: string;
  offset?: number;
  page_size?: number;
}

export interface ListCampaignsParams {
  owner?: string;
  start_date?: string;
  end_date?: string;
  offset?: number;
  page_size?: number;
}

export interface UpdateCampaignParams {
  id: string;
  body: Schema.CampaignUpdateRequest;
}

export interface UpdateCampaignFieldParams {
  campaign_id: string;
  field_id: string;
  body: Schema.CampaignFieldUpdateRequest;
}

export interface CreateEventParams {
  body: Schema.EventCreateRequest;
}

export interface GetEventParams {
  id: string;
}

export interface ListEventFieldsParams {
  id: string;
  offset?: number;
  page_size?: number;
}

export interface ListEventsParams {
  start_date?: string;
  end_date?: string;
  campaign_id?: string;
  offset?: number;
  page_size?: number;
}

export interface UpdateEventParams {
  id: string;
  body: Schema.EventUpdateRequest;
}

export interface UpdateEventFieldsParams {
  id: string;
  body: Schema.EventFieldsUpdateRequest;
}

export interface CreateFieldParams {
  body: Schema.SettingsFieldCreateRequest;
}

export interface CreateFieldChoicesParams {
  id: string;
  body: Schema.SettingsFieldChoiceCreateRequest;
}

export interface DeleteFieldChoiceParams {
  field_id: string;
  choice_id: string;
}

export interface ListFieldsParams {
  ids?: string;
  offset?: number;
  page_size?: number;
}

export interface UpdateFieldParams {
  id: string;
  body: Schema.SettingsFieldUpdateRequest;
}

export interface UpdateFieldChoiceParams {
  field_id: string;
  choice_id: string;
  body: Schema.SettingsFieldChoiceUpdateRequest;
}

export interface ListLabelGroupsParams {
  /** Source organization type to filter by */
  source_org_type?: 'current' | 'related';
  offset?: number;
  page_size?: number;
}

export interface AddAssetPermissionsParams {
  asset_id: string;
  body: Schema.AssetPermissionBulkCreateRequest;
}

export interface AddFolderPermissionsParams {
  id: string;
  body: Schema.FolderPermissionBulkCreateRequest;
}

export interface CreateAssetParams {
  body: Schema.LibraryAssetCreateRequest;
}

export interface CreateAssetLineageParams {
  asset_id: string;
  body: Schema.AssetLineageCreateRequest;
}

export interface CreateAssetVersionParams {
  asset_id: string;
  body: Schema.AssetVersionCreateRequest;
}

export interface CreateFileUrlsParams {
  body: Schema.FileUrlBulkCreateRequest;
}

export interface CreateFolderParams {
  body: Schema.FolderCreateRequest;
}

export interface CreateStructuredContentParams {
  body: Schema.LibraryStructuredContentCreateRequest;
}

export interface DeleteAssetLineageParams {
  asset_id: string;
  lineage_id: string;
}

export interface DeleteFolderParams {
  id: string;
}

export interface DeleteImageParams {
  id: string;
}

export interface DeleteRawFileParams {
  id: string;
}

export interface DeleteVideoParams {
  id: string;
}

export interface GetArticleParams {
  id: string;
}

export interface GetFolderParams {
  id: string;
}

export interface GetImageParams {
  id: string;
}

export interface GetRawFileParams {
  id: string;
}

export interface GetRenditionParams {
  id: string;
}

export interface GetRenditionConfigParams {
  id: string;
}

export interface GetStructuredContentParams {
  id: string;
}

export interface GetVideoParams {
  id: string;
}

export interface ListAssetFieldsParams {
  asset_id: string;
  offset?: number;
  page_size?: number;
}

export interface ListAssetLineagesParams {
  asset_id?: string;
  used_in?: 'external';
  /** Date and time as the lower limit to filter asset lineages by `created_at`, in ISO 8601 UTC format */
  created_at__from?: string;
  /** Date and time as the upper limit to filter asset lineages by `created_at`, in ISO 8601 UTC format */
  created_at__to?: string;
  offset?: number;
  page_size?: number;
}

export interface ListAssetPermissionsParams {
  asset_id: string;
  access?: 'view' | 'edit' | 'comment' | 'delete';
  max_access?: 'view' | 'edit' | 'comment' | 'delete';
  min_access?: 'view' | 'edit' | 'comment' | 'delete';
}

export interface ListAssetRenditionsParams {
  asset_id: string;
  offset?: number;
  page_size?: number;
}

export interface ListAssetsParams {
  /** Asset type to filter by. Example: `type=image&type=video` */
  type?: 'article' | 'image' | 'video' | 'raw_file' | 'structured_content'[];
  /** Label to filter by. **Must be passed as urlencoded string**. Labels that do not exist are ignored. If none of the provided labels exist, filtering is not applied. Example – `label=%7B%22group%22%3A%22ee63e3ee43925bb5cc8cd17b817d93ee%22%2C%22values%22%3A%5B%226706efc7828cd6aaedbc0434139cd3e1%22%2C%221f32651216cf2aefcaa08be1ea7dedf1%22%5D%7D` */
  label?: Record<string, unknown>[];
  /** List of fields to filter by. **Must be passed as base64 encoded string**. Fields that do not exist are ignored. If none of the provided fields exist, filtering is not applied. Example – `fields=Wwp7CiJpZCI6ICI2N2E4NDZhMWM3NzU1YTFwNThpNjh5MzVhIiwKInZhbHVlcyI6IFsiNjdhODQ2YTFjNzc1YWU1YWExYTE0YTA1Il0KfQpd=` */
  fields?: Record<string, unknown>[];
  /** Date and time as the lower limit to filter assets by `created_at`, in ISO 8601 UTC format */
  created_at__from?: string;
  /** Date and time as the upper limit to filter assets by `created_at`, in ISO 8601 UTC format */
  created_at__to?: string;
  /** Date and time as the lower limit to filter assets by `modified_at`, in ISO 8601 UTC format */
  modified_at__from?: string;
  /** Date and time as the upper limit to filter assets by `modified_at`, in ISO 8601 UTC format */
  modified_at__to?: string;
  /** ID of the library folder to include assets from */
  folder_id?: string;
  /** Indicates whether assets from subfolders need to be included */
  include_subfolder_assets?: boolean;
  /** Search assets by title or content description */
  search_text?: string;
  /** ID of the campaign to include assets from */
  campaign_id?: string;
  offset?: number;
  page_size?: number;
}

export interface ListFolderPermissionsParams {
  id: string;
  access?: 'view' | 'edit' | 'comment' | 'delete';
  max_access?: 'view' | 'edit' | 'comment' | 'delete';
  min_access?: 'view' | 'edit' | 'comment' | 'delete';
}

export interface ListFoldersParams {
  /** ID of the parent folder to filter by */
  parent_folder_id?: string;
  offset?: number;
  page_size?: number;
}

export interface ListRelatedAssetsParams {
  asset_id: string;
}

export interface RemoveAssetPermissionParams {
  asset_id: string;
  accessor_id: string;
}

export interface RemoveFolderPermissionParams {
  id: string;
  accessor_id: string;
}

export interface ReplaceRelatedAssetsParams {
  asset_id: string;
  body: Schema.ReplaceRelatedAssetsRequest;
}

export interface UpdateAssetFieldParams {
  asset_id: string;
  field_id: string;
  body: Schema.AssetFieldUpdateRequest;
}

export interface UpdateAssetFieldsParams {
  asset_id: string;
  body: Schema.AssetFieldsUpdateRequest;
}

export interface UpdateAssetPermissionParams {
  asset_id: string;
  accessor_id: string;
  body: Schema.AssetPermissionUpdateRequest;
}

export interface UpdateFolderParams {
  id: string;
  body: Schema.FolderUpdateRequest;
}

export interface UpdateFolderPermissionParams {
  id: string;
  accessor_id: string;
  body: Schema.FolderPermissionUpdateRequest;
}

export interface UpdateImageParams {
  id: string;
  body: Schema.LibraryImageUpdateRequest;
}

export interface UpdateRawFileParams {
  id: string;
  body: Schema.LibraryRawFileUpdateRequest;
}

export interface UpdateStructuredContentParams {
  id: string;
  body: Schema.LibraryStructuredContentUpdateRequest;
}

export interface UpdateVideoParams {
  id: string;
  body: Schema.LibraryVideoUpdateRequest;
}

export interface CreateMilestoneParams {
  body: Schema.MilestoneCreateRequest;
}

export interface GetMilestoneParams {
  id: string;
}

export interface ListMilestonesParams {
  campaign_id?: string;
  due_date__from?: string;
  due_date__to?: string;
  offset?: number;
  page_size?: number;
}

export interface UpdateMilestoneParams {
  id: string;
  body: Schema.MilestoneUpdateRequest;
}

export interface BulkCreatePublishingEventMetadataParams {
  /** Unique identifier of the publishing event */
  publishing_event_id: string;
  body: Schema.PublishingEventMetadataBulkCreateRequest;
}

export interface GetPublishingEventParams {
  /** Unique identifier of the publishing event */
  publishing_event_id: string;
}

export interface GetPublishingEventAssetMetadataParams {
  /** Unique identifier of the publishing event */
  publishing_event_id: string;
  /** Unique identifier of the asset */
  asset_id: string;
  /** Unique identifier of the publishing metadata */
  publishing_metadata_id: string;
}

export interface ListPublishingChannelsParams {
  /** Number of results to return per page */
  page_size?: number;
}

export interface ListPublishingEventMetadataParams {
  /** Unique identifier of the publishing event */
  publishing_event_id: string;
  /** Publishing status of the asset */
  status?: 'published' | 'unpublished' | 'synced' | 'failed';
  /** Type of asset */
  asset_type?: 'article' | 'image' | 'video' | 'raw_file' | 'structured_content';
  /** Unique identifier of the asset. */
  asset_id?: string;
  /** The locale to which the asset is being published to. */
  locale?: string;
}

export type GetSettingsParams = Record<string, never>;

export interface UpdateSettingsParams {
  /** If `execute=true` the settings are created or updated. Otherwise, the endpoint returns only the changeset. */
  execute?: boolean;
  /** If `overwrite_workflows=true` the existing workflows are overwritten. Otherwise, a new workflow is created where a prefix `Copy of` is added to the workflow's name. */
  overwrite_workflows?: boolean;
  body: Schema.SettingsResources;
}

export interface AcknowledgeSCContentPreviewParams {
  content_id: string;
  version_id: string;
  preview_id: string;
  body: Schema.SCContentPreviewAcknowledgeRequest;
}

export interface CompleteSCContentPreviewParams {
  content_id: string;
  version_id: string;
  preview_id: string;
  body: Schema.SCContentPreviewCompleteRequest;
}

export interface CreateSCContentTypeParams {
  body: Schema.SCContentTypeCreateRequest;
}

export interface CreateSCContentTypeManagedMigrationParams {
  content_type_id: string;
  body: Schema.SCContentTypeManagedMigrationCreateRequest;
}

export interface CreateSCContentTypeVersionParams {
  content_type_id: string;
  body: Schema.SCContentTypeVersionCreateRequest;
}

export interface DeleteSCContentTypeManagedMigrationParams {
  content_type_id: string;
  job_id: string;
}

export interface GetSCContentTypeParams {
  content_type_id: string;
}

export interface GetSCContentTypeManagedMigrationParams {
  content_type_id: string;
  job_id: string;
}

export interface GetSCContentTypeVersionParams {
  content_type_id: string;
  version_id: string;
}

export interface ListSCContentTypeManagedMigrationsParams {
  content_type_id: string;
  /** Whether include a summary of content migration status (total, not started, succeeded, errored). */
  content_migration_summary?: boolean;
  /** Pagination offset (number of jobs to skip). */
  offset?: number;
  /** Pagination limit (number of jobs to return). */
  limit?: number;
}

export interface ListSCContentTypesParams {
  source?: string;
  disabled?: boolean;
  list?: Schema.ContentTypeListingOption;
}

export interface ListSCContentTypeVersionsParams {
  content_type_id: string;
}

export interface MigrateSCContentParams {
  content_id: string;
  body: Schema.SCContentMigrationCreateRequest;
}

export interface StartSCContentTypeManagedMigrationParams {
  content_type_id: string;
  job_id: string;
}

export interface UpdateSCContentTypeParams {
  content_type_id: string;
  body: Schema.SCContentTypeUpdateRequest;
}

export interface UpdateSCContentTypeManagedMigrationParams {
  content_type_id: string;
  job_id: string;
  body: Record<string, unknown>;
}

export interface ValidateSCContentTypeManagedMigrationParams {
  content_type_id: string;
  body: Schema.SCContentTypeManagedMigrationValidateRequest;
}

export interface AddAssetToTaskParams {
  id: string;
  body: Schema.TaskAssetRequest;
}

export interface AddCommentToTaskParams {
  task_id: string;
  body: Schema.CommentCreateRequest;
}

export interface AddCommentToTaskAssetParams {
  task_id: string;
  asset_id: string;
  body: Schema.CommentCreateRequest;
}

export interface AddCommentToTaskSubStepParams {
  task_id: string;
  step_id: string;
  sub_step_id: string;
  body: Schema.CommentWithReplyCreateRequest;
}

export interface AddDraftToTaskAssetParams {
  task_id: string;
  asset_id: string;
  body: Schema.TaskAssetRequest;
}

export interface AddFieldToTaskParams {
  task_id: string;
  body: Schema.ObjectFieldCreateRequest;
}

export interface AddStructuredContentToTaskParams {
  task_id: string;
  body: Schema.TaskStructuredContentCreateRequest;
}

export interface AddUrlToTaskParams {
  task_id: string;
  body: Schema.AddUrlToTaskRequest;
}

export interface CreateTaskParams {
  body: Schema.TaskCreateRequest;
}

export interface CreateTaskPublishingIntentParams {
  task_id: string;
  body: Schema.TaskPublishingIntentCreateRequest;
}

export interface CreateTaskStructuredContentDraftParams {
  task_id: string;
  content_id: string;
  body: Schema.TaskStructuredContentDraftRequest;
}

export interface DeleteTaskStructuredContentParams {
  task_id: string;
  content_id: string;
}

export interface DeleteTaskSubStepCommentParams {
  task_id: string;
  step_id: string;
  sub_step_id: string;
  comment_id: string;
}

export interface GetTaskParams {
  id: string;
}

export interface GetTaskArticleParams {
  task_id: string;
  article_id: string;
}

export interface GetTaskBriefParams {
  id: string;
}

export interface GetTaskCustomFieldParams {
  task_id: string;
  custom_field_id: string;
}

export interface GetTaskImageParams {
  task_id: string;
  image_id: string;
}

export interface GetTaskRawFileParams {
  task_id: string;
  raw_file_id: string;
}

export interface GetTaskStructuredContentParams {
  task_id: string;
  content_id: string;
}

export interface GetTaskSubStepParams {
  task_id: string;
  step_id: string;
  sub_step_id: string;
}

export interface GetTaskSubStepCommentParams {
  task_id: string;
  step_id: string;
  sub_step_id: string;
  comment_id: string;
}

export interface GetTaskSubStepExternalWorkParams {
  task_id: string;
  step_id: string;
  sub_step_id: string;
}

export interface GetTaskVideoParams {
  task_id: string;
  video_id: string;
}

export interface ListTaskAssetCommentsParams {
  task_id: string;
  asset_id: string;
  offset?: number;
  page_size?: number;
}

export interface ListTaskAssetDraftsParams {
  task_id: string;
  asset_id: string;
  offset?: number;
  page_size?: number;
}

export interface ListTaskAssetFieldsParams {
  task_id: string;
  asset_id: string;
  offset?: number;
  page_size?: number;
}

export interface ListTaskAssetsParams {
  id: string;
  offset?: number;
  page_size?: number;
}

export interface ListTaskAttachmentsParams {
  id: string;
  offset?: number;
  page_size?: number;
}

export interface ListTaskCustomFieldChoicesParams {
  task_id: string;
  custom_field_id: string;
  offset?: number;
  page_size?: number;
}

export interface ListTaskCustomFieldsParams {
  id: string;
  offset?: number;
  page_size?: number;
}

export interface ListTaskFieldsParams {
  task_id: string;
  offset?: number;
  page_size?: number;
}

export interface ListTasksParams {
  search_key?: string;
  campaign?: string;
  workflow?: string;
  milestone?: string;
  start_date?: string;
  due_date?: string;
  status?: 'Archived' | 'Completed' | 'Overdue' | 'Not Started' | 'In Progress' | 'On Hold'[];
  offset?: number;
  page_size?: number;
}

export interface ListTaskSubStepCommentsParams {
  task_id: string;
  step_id: string;
  sub_step_id: string;
  offset?: number;
  page_size?: number;
}

export interface ListTaskSubStepFieldsParams {
  task_id: string;
  step_id: string;
  sub_step_id: string;
  offset?: number;
  page_size?: number;
}

export interface RemoveTaskFieldParams {
  task_id: string;
  field_id: string;
}

export interface UpdateTaskParams {
  id: string;
  body: Schema.TaskUpdateRequest;
}

export interface UpdateTaskAssetFieldsParams {
  task_id: string;
  asset_id: string;
  body: Schema.TaskAssetFieldsUpdateRequest;
}

export interface UpdateTaskCustomFieldParams {
  task_id: string;
  custom_field_id: string;
  body: Schema.TaskCustomFieldUpdateRequest;
}

export interface UpdateTaskFieldParams {
  task_id: string;
  field_id: string;
  body: Schema.TaskFieldUpdateRequest;
}

export interface UpdateTaskStructuredContentParams {
  task_id: string;
  content_id: string;
  body: Schema.TaskStructuredContentUpdateRequest;
}

export interface UpdateTaskSubStepParams {
  task_id: string;
  step_id: string;
  sub_step_id: string;
  body: Schema.TaskSubStepRequest;
}

export interface UpdateTaskSubStepCommentParams {
  task_id: string;
  step_id: string;
  sub_step_id: string;
  comment_id: string;
  body: Schema.TaskSubStepCommentUpdateRequest;
}

export interface UpdateTaskSubStepExternalWorkParams {
  task_id: string;
  step_id: string;
  sub_step_id: string;
  body: Schema.TaskExternalWorkRequest;
}

export interface UpdateTaskStepParams {
  task_id: string;
  step_id: string;
  body: Schema.TaskStepRequest;
}

export interface GetTeamParams {
  id: string;
}

export interface ListTeamsParams {
  offset?: number;
  page_size?: number;
}

export interface GetTemplateParams {
  template_id: string;
}

export interface ListTemplatesParams {
  search?: string;
  applicable_to?: 'work_request' | 'task_brief' | 'campaign_brief';
  include_inactive?: boolean;
}

export interface CompleteMultipartUploadParams {
  /** ID of the multipart upload to complete */
  id: string;
}

export interface CreateMultipartUploadParams {
  body: Record<string, unknown>;
}

export interface GetMultipartUploadStatusParams {
  /** ID of the multipart upload to check status for */
  id: string;
}

export type GetUploadUrlParams = Record<string, never>;

export interface FindUserByEmailParams {
  email: string;
}

export interface GetUserParams {
  id: string;
}

export type ListUsersParams = Record<string, never>;

export interface GetWorkflowParams {
  workflow_id: string;
}

export interface ListWorkflowsParams {
  offset?: number;
  page_size?: number;
}

export interface AddAttachmentToWorkRequestParams {
  id: string;
  body: Schema.AttachmentRequest;
}

export interface AddCommentToWorkRequestParams {
  id: string;
  body: Schema.CommentWithReplyCreateRequest;
}

export interface CreateCampaignFromWorkRequestParams {
  id: string;
  body: Schema.WorkRequestCampaignRequest;
}

export interface CreateTaskFromWorkRequestParams {
  id: string;
  body: Schema.WorkRequestTaskRequest;
}

export interface CreateWorkRequestParams {
  body: Schema.WorkRequestCreateRequest;
}

export interface CreateWorkRequestCreativeAssetParams {
  id: string;
  body: Schema.CreativeAssetRequest;
}

export interface DeleteWorkRequestAttachmentParams {
  work_request_id: string;
  attachment_id: string;
}

export interface DeleteWorkRequestCreativeAssetParams {
  work_request_id: string;
  creative_asset_id: string;
}

export interface GetWorkRequestParams {
  id: string;
}

export interface GetWorkRequestCommentParams {
  work_request_id: string;
  comment_id: string;
}

export interface ListWorkRequestApprovedAssetsParams {
  id: string;
  offset?: number;
  page_size?: number;
}

export interface ListWorkRequestCommentsParams {
  id: string;
  offset?: number;
  page_size?: number;
}

export interface ListWorkRequestRelatedResourcesParams {
  id: string;
  offset?: number;
  page_size?: number;
}

export interface ListWorkRequestsParams {
  created_by?: string;
  status?: 'Submitted' | 'Accepted' | 'Completed' | 'Declined';
  order_by?: 'priority' | 'status' | 'created_at' | 'due_date';
  order_as?: 'asc' | 'desc';
  /** Unique identifiers of the templates. You can append multiple times, for example `template_ids=template1&template_ids=template2`. */
  template_ids?: string[];
  /** Date and time as the lower limit to filter work requests by `created_at`, in ISO 8601 UTC format */
  created_at__from?: string;
  /** Date and time as the upper limit to filter work requests by `created_at`, in ISO 8601 UTC format */
  created_at__to?: string;
  offset?: number;
  page_size?: number;
}

export interface UpdateWorkRequestParams {
  id: string;
  body: Schema.WorkRequestUpdateRequest;
}

export interface UpdateWorkRequestFormFieldParams {
  work_request_id: string;
  form_field_identifier: string;
  body: Schema.WorkRequestFormFieldUpdateRequest;
}

