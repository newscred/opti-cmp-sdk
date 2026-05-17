// Auto-generated - DO NOT EDIT

import type { Response } from "../request/types.js";
import type * as Params from "./params.js";
import type * as Schema from "./schema.js";

export interface APIEndpoints {
  asset: {
    getAssetUrl: (
      params: Params.GetAssetUrlParams,
    ) => Promise<void>;
  };
  brandCompliance: {
    getTaskAssetDraftBrandCompliance: (
      params: Params.GetTaskAssetDraftBrandComplianceParams,
    ) => Promise<Response<Schema.TaskAssetDraftBrandComplianceResponse>>;
    listBrandComplianceCategories: (
      params: Params.ListBrandComplianceCategoriesParams,
    ) => Promise<Response<Schema.ListBrandComplianceCategoriesResponse>>;
    updateTaskAssetDraftBrandCompliance: (
      params: Params.UpdateTaskAssetDraftBrandComplianceParams,
    ) => Promise<Response<Schema.TaskAssetDraftBrandComplianceResponse>>;
  };
  campaign: {
    addAttachmentToCampaign: (
      params: Params.AddAttachmentToCampaignParams,
    ) => Promise<Response<Schema.AttachmentResponse>>;
    addCommentToCampaign: (
      params: Params.AddCommentToCampaignParams,
    ) => Promise<Response<Schema.CampaignCommentResponse>>;
    addFieldToCampaign: (
      params: Params.AddFieldToCampaignParams,
    ) => Promise<Response<Schema.ObjectFieldCreateResponse>>;
    createCampaign: (
      params: Params.CreateCampaignParams,
    ) => Promise<Response<Schema.CampaignResponse>>;
    getCampaign: (
      params: Params.GetCampaignParams,
    ) => Promise<Response<Schema.CampaignResponse>>;
    getCampaignBrief: (
      params: Params.GetCampaignBriefParams,
    ) => Promise<Response<Schema.CampaignBriefResponse>>;
    listCampaignFields: (
      params: Params.ListCampaignFieldsParams,
    ) => Promise<Response<Schema.ListCampaignFieldsResponse>>;
    listCampaigns: (
      params: Params.ListCampaignsParams,
    ) => Promise<Response<Schema.ListCampaignsResponse>>;
    updateCampaign: (
      params: Params.UpdateCampaignParams,
    ) => Promise<void>;
    updateCampaignField: (
      params: Params.UpdateCampaignFieldParams,
    ) => Promise<Response<Schema.CampaignFieldUpdateResponse>>;
  };
  event: {
    createEvent: (
      params: Params.CreateEventParams,
    ) => Promise<Response<Schema.EventResponse>>;
    getEvent: (
      params: Params.GetEventParams,
    ) => Promise<Response<Schema.EventResponse>>;
    listEventFields: (
      params: Params.ListEventFieldsParams,
    ) => Promise<Response<Schema.ListEventFieldsResponse>>;
    listEvents: (
      params: Params.ListEventsParams,
    ) => Promise<Response<Schema.ListEventsResponse>>;
    updateEvent: (
      params: Params.UpdateEventParams,
    ) => Promise<Response<Schema.EventResponse>>;
    updateEventFields: (
      params: Params.UpdateEventFieldsParams,
    ) => Promise<void>;
  };
  field: {
    createField: (
      params: Params.CreateFieldParams,
    ) => Promise<Response<Schema.CreateFieldResponse>>;
    createFieldChoices: (
      params: Params.CreateFieldChoicesParams,
    ) => Promise<Response<Schema.SettingsFieldChoiceCreateResponse>>;
    deleteFieldChoice: (
      params: Params.DeleteFieldChoiceParams,
    ) => Promise<void>;
    listFields: (
      params: Params.ListFieldsParams,
    ) => Promise<Response<Schema.ListFieldsResponse>>;
    updateField: (
      params: Params.UpdateFieldParams,
    ) => Promise<void>;
    updateFieldChoice: (
      params: Params.UpdateFieldChoiceParams,
    ) => Promise<void>;
  };
  label: {
    listLabelGroups: (
      params: Params.ListLabelGroupsParams,
    ) => Promise<Response<Schema.ListLabelGroupsResponse>>;
  };
  library: {
    addAssetPermissions: (
      params: Params.AddAssetPermissionsParams,
    ) => Promise<void>;
    addFolderPermissions: (
      params: Params.AddFolderPermissionsParams,
    ) => Promise<void>;
    createAsset: (
      params: Params.CreateAssetParams,
    ) => Promise<Response<Schema.AssetResponse>>;
    createAssetLineage: (
      params: Params.CreateAssetLineageParams,
    ) => Promise<Response<Schema.AssetLineageResponse>>;
    createAssetVersion: (
      params: Params.CreateAssetVersionParams,
    ) => Promise<Response<Schema.LibraryAssetVersionResponse>>;
    createFileUrls: (
      params: Params.CreateFileUrlsParams,
    ) => Promise<Response<Schema.BatchFileUrlResponse>>;
    createFolder: (
      params: Params.CreateFolderParams,
    ) => Promise<Response<Schema.FolderResponse>>;
    createStructuredContent: (
      params: Params.CreateStructuredContentParams,
    ) => Promise<Response<Schema.LibraryStructuredContent>>;
    deleteAssetLineage: (
      params: Params.DeleteAssetLineageParams,
    ) => Promise<void>;
    deleteFolder: (
      params: Params.DeleteFolderParams,
    ) => Promise<void>;
    deleteImage: (
      params: Params.DeleteImageParams,
    ) => Promise<void>;
    deleteRawFile: (
      params: Params.DeleteRawFileParams,
    ) => Promise<void>;
    deleteVideo: (
      params: Params.DeleteVideoParams,
    ) => Promise<void>;
    getArticle: (
      params: Params.GetArticleParams,
    ) => Promise<Response<Schema.LibraryArticle>>;
    getFolder: (
      params: Params.GetFolderParams,
    ) => Promise<Response<Schema.FolderResponse>>;
    getImage: (
      params: Params.GetImageParams,
    ) => Promise<Response<Schema.LibraryImage>>;
    getRawFile: (
      params: Params.GetRawFileParams,
    ) => Promise<Response<Schema.LibraryRawFile>>;
    getRendition: (
      params: Params.GetRenditionParams,
    ) => Promise<Response<Schema.DetailedAssetRenditionResponse>>;
    getRenditionConfig: (
      params: Params.GetRenditionConfigParams,
    ) => Promise<Response<Schema.RenditionConfigResponse>>;
    getStructuredContent: (
      params: Params.GetStructuredContentParams,
    ) => Promise<Response<Schema.LibraryStructuredContent>>;
    getVideo: (
      params: Params.GetVideoParams,
    ) => Promise<Response<Schema.LibraryVideo>>;
    listAssetFields: (
      params: Params.ListAssetFieldsParams,
    ) => Promise<Response<Schema.ListAssetFieldsResponse>>;
    listAssetLineages: (
      params: Params.ListAssetLineagesParams,
    ) => Promise<Response<Schema.ListAssetLineagesResponse>>;
    listAssetPermissions: (
      params: Params.ListAssetPermissionsParams,
    ) => Promise<Response<Schema.AssetPermissionListResponseItem>>;
    listAssetRenditions: (
      params: Params.ListAssetRenditionsParams,
    ) => Promise<Response<Schema.ListAssetRenditionsResponse>>;
    listAssets: (
      params: Params.ListAssetsParams,
    ) => Promise<Response<Schema.ListAssetsResponse>>;
    listFolderPermissions: (
      params: Params.ListFolderPermissionsParams,
    ) => Promise<Response<Schema.FolderPermissionListResponseItem>>;
    listFolders: (
      params: Params.ListFoldersParams,
    ) => Promise<Response<Schema.ListFoldersResponse>>;
    removeAssetPermission: (
      params: Params.RemoveAssetPermissionParams,
    ) => Promise<void>;
    removeFolderPermission: (
      params: Params.RemoveFolderPermissionParams,
    ) => Promise<void>;
    updateAssetField: (
      params: Params.UpdateAssetFieldParams,
    ) => Promise<Response<Schema.AssetFieldUpdateResponse>>;
    updateAssetFields: (
      params: Params.UpdateAssetFieldsParams,
    ) => Promise<Response<Schema.UpdateAssetFieldsResponse>>;
    updateAssetPermission: (
      params: Params.UpdateAssetPermissionParams,
    ) => Promise<void>;
    updateFolder: (
      params: Params.UpdateFolderParams,
    ) => Promise<Response<Schema.FolderResponse>>;
    updateFolderPermission: (
      params: Params.UpdateFolderPermissionParams,
    ) => Promise<void>;
    updateImage: (
      params: Params.UpdateImageParams,
    ) => Promise<Response<Schema.LibraryImage>>;
    updateRawFile: (
      params: Params.UpdateRawFileParams,
    ) => Promise<Response<Schema.LibraryRawFile>>;
    updateStructuredContent: (
      params: Params.UpdateStructuredContentParams,
    ) => Promise<Response<Schema.LibraryStructuredContent>>;
    updateVideo: (
      params: Params.UpdateVideoParams,
    ) => Promise<Response<Schema.LibraryVideo>>;
  };
  milestone: {
    createMilestone: (
      params: Params.CreateMilestoneParams,
    ) => Promise<Response<Schema.MilestoneResponse>>;
    getMilestone: (
      params: Params.GetMilestoneParams,
    ) => Promise<Response<Schema.MilestoneResponse>>;
    listMilestones: (
      params: Params.ListMilestonesParams,
    ) => Promise<Response<Schema.ListMilestonesResponse>>;
    updateMilestone: (
      params: Params.UpdateMilestoneParams,
    ) => Promise<Response<Schema.MilestoneResponse>>;
  };
  publishing: {
    bulkCreatePublishingEventMetadata: (
      params: Params.BulkCreatePublishingEventMetadataParams,
    ) => Promise<Response<Schema.PublishingEventMetadataBulkCreateResponse>>;
    getPublishingEvent: (
      params: Params.GetPublishingEventParams,
    ) => Promise<Response<Schema.PublishingEventResponse>>;
    getPublishingEventAssetMetadata: (
      params: Params.GetPublishingEventAssetMetadataParams,
    ) => Promise<Response<Schema.PublishingEventMetadataResponse>>;
    listPublishingEventMetadata: (
      params: Params.ListPublishingEventMetadataParams,
    ) => Promise<Response<Schema.PublishingEventMetadataListResponse>>;
  };
  settings: {
    getSettings: (
      params: Params.GetSettingsParams,
    ) => Promise<Response<Schema.Settings>>;
    updateSettings: (
      params: Params.UpdateSettingsParams,
    ) => Promise<Response<Schema.SettingsUpdateResponse>>;
  };
  structuredContent: {
    acknowledgeSCContentPreview: (
      params: Params.AcknowledgeSCContentPreviewParams,
    ) => Promise<Response<unknown>>;
    completeSCContentPreview: (
      params: Params.CompleteSCContentPreviewParams,
    ) => Promise<Response<unknown>>;
    createSCContentType: (
      params: Params.CreateSCContentTypeParams,
    ) => Promise<Response<Schema.SCContentTypeCreateResponse>>;
    createSCContentTypeManagedMigration: (
      params: Params.CreateSCContentTypeManagedMigrationParams,
    ) => Promise<Response<Schema.CreateSCContentTypeManagedMigrationResponse>>;
    createSCContentTypeVersion: (
      params: Params.CreateSCContentTypeVersionParams,
    ) => Promise<Response<Schema.SCContentTypeCreateResponse>>;
    deleteSCContentTypeManagedMigration: (
      params: Params.DeleteSCContentTypeManagedMigrationParams,
    ) => Promise<void>;
    getSCContentType: (
      params: Params.GetSCContentTypeParams,
    ) => Promise<Response<Schema.SCContentType>>;
    getSCContentTypeManagedMigration: (
      params: Params.GetSCContentTypeManagedMigrationParams,
    ) => Promise<Response<Schema.SCContentTypeManagedMigrationResponse>>;
    getSCContentTypeVersion: (
      params: Params.GetSCContentTypeVersionParams,
    ) => Promise<Response<Schema.SCContentTypeVersion>>;
    listSCContentTypeManagedMigrations: (
      params: Params.ListSCContentTypeManagedMigrationsParams,
    ) => Promise<Response<Schema.SCContentTypeManagedMigrationResponse[]>>;
    listSCContentTypes: (
      params: Params.ListSCContentTypesParams,
    ) => Promise<Response<Schema.BaseContentTypeModel[]>>;
    listSCContentTypeVersions: (
      params: Params.ListSCContentTypeVersionsParams,
    ) => Promise<Response<Schema.BaseContentTypeVersionModel[]>>;
    migrateSCContent: (
      params: Params.MigrateSCContentParams,
    ) => Promise<Response<Schema.MigrateSCContentResponse>>;
    startSCContentTypeManagedMigration: (
      params: Params.StartSCContentTypeManagedMigrationParams,
    ) => Promise<Response<Schema.SCContentTypeManagedMigrationStartResponse>>;
    updateSCContentType: (
      params: Params.UpdateSCContentTypeParams,
    ) => Promise<Response<Schema.SCContentTypeUpdateResponse>>;
    updateSCContentTypeManagedMigration: (
      params: Params.UpdateSCContentTypeManagedMigrationParams,
    ) => Promise<Response<Schema.UpdateSCContentTypeManagedMigrationResponse>>;
    validateSCContentTypeManagedMigration: (
      params: Params.ValidateSCContentTypeManagedMigrationParams,
    ) => Promise<Response<Schema.ValidateSCContentTypeManagedMigrationResponse>>;
  };
  task: {
    addAssetToTask: (
      params: Params.AddAssetToTaskParams,
    ) => Promise<Response<Schema.TaskAssetResponse>>;
    addCommentToTask: (
      params: Params.AddCommentToTaskParams,
    ) => Promise<Response<Schema.TaskCommentResponse>>;
    addCommentToTaskAsset: (
      params: Params.AddCommentToTaskAssetParams,
    ) => Promise<Response<Schema.TaskAssetCommentResponse>>;
    addCommentToTaskSubStep: (
      params: Params.AddCommentToTaskSubStepParams,
    ) => Promise<Response<Schema.TaskSubStepCommentResponse>>;
    addDraftToTaskAsset: (
      params: Params.AddDraftToTaskAssetParams,
    ) => Promise<Response<Schema.TaskAssetDraftResponse>>;
    addFieldToTask: (
      params: Params.AddFieldToTaskParams,
    ) => Promise<Response<Schema.ObjectFieldCreateResponse>>;
    addStructuredContentToTask: (
      params: Params.AddStructuredContentToTaskParams,
    ) => Promise<Response<Schema.ContentDetailsModel>>;
    addUrlToTask: (
      params: Params.AddUrlToTaskParams,
    ) => Promise<Response<Schema.TaskUrlResponse>>;
    createTask: (
      params: Params.CreateTaskParams,
    ) => Promise<Response<Schema.TaskResponse>>;
    createTaskStructuredContentDraft: (
      params: Params.CreateTaskStructuredContentDraftParams,
    ) => Promise<Response<Record<string, unknown>>>;
    deleteTaskStructuredContent: (
      params: Params.DeleteTaskStructuredContentParams,
    ) => Promise<void>;
    deleteTaskSubStepComment: (
      params: Params.DeleteTaskSubStepCommentParams,
    ) => Promise<void>;
    getTask: (
      params: Params.GetTaskParams,
    ) => Promise<Response<Schema.TaskResponse>>;
    getTaskArticle: (
      params: Params.GetTaskArticleParams,
    ) => Promise<Response<Schema.TaskArticle>>;
    getTaskBrief: (
      params: Params.GetTaskBriefParams,
    ) => Promise<Response<Schema.TaskBriefResponse>>;
    getTaskCustomField: (
      params: Params.GetTaskCustomFieldParams,
    ) => Promise<Response<Schema.TaskCustomField>>;
    getTaskImage: (
      params: Params.GetTaskImageParams,
    ) => Promise<Response<Schema.TaskImage>>;
    getTaskRawFile: (
      params: Params.GetTaskRawFileParams,
    ) => Promise<Response<Schema.TaskRawFile>>;
    getTaskStructuredContent: (
      params: Params.GetTaskStructuredContentParams,
    ) => Promise<Response<Schema.ContentDetailsModel>>;
    getTaskSubStep: (
      params: Params.GetTaskSubStepParams,
    ) => Promise<Response<Schema.TaskSubStep>>;
    getTaskSubStepComment: (
      params: Params.GetTaskSubStepCommentParams,
    ) => Promise<Response<Schema.TaskSubStepCommentResponse>>;
    getTaskSubStepExternalWork: (
      params: Params.GetTaskSubStepExternalWorkParams,
    ) => Promise<Response<Schema.TaskExternalWorkResponse>>;
    getTaskVideo: (
      params: Params.GetTaskVideoParams,
    ) => Promise<Response<Schema.TaskVideo>>;
    listTaskAssetComments: (
      params: Params.ListTaskAssetCommentsParams,
    ) => Promise<Response<Schema.ListTaskAssetCommentsResponse>>;
    listTaskAssetDrafts: (
      params: Params.ListTaskAssetDraftsParams,
    ) => Promise<Response<Schema.ListTaskAssetDraftsResponse>>;
    listTaskAssetFields: (
      params: Params.ListTaskAssetFieldsParams,
    ) => Promise<Response<Schema.ListTaskAssetFieldsResponse>>;
    listTaskAssets: (
      params: Params.ListTaskAssetsParams,
    ) => Promise<Response<Schema.ListTaskAssetsResponse>>;
    listTaskAttachments: (
      params: Params.ListTaskAttachmentsParams,
    ) => Promise<Response<Schema.ListTaskAttachmentsResponse>>;
    listTaskCustomFieldChoices: (
      params: Params.ListTaskCustomFieldChoicesParams,
    ) => Promise<Response<Schema.ListTaskCustomFieldChoicesResponse>>;
    listTaskCustomFields: (
      params: Params.ListTaskCustomFieldsParams,
    ) => Promise<Response<Schema.ListTaskCustomFieldsResponse>>;
    listTaskFields: (
      params: Params.ListTaskFieldsParams,
    ) => Promise<Response<Schema.ListTaskFieldsResponse>>;
    listTasks: (
      params: Params.ListTasksParams,
    ) => Promise<Response<Schema.ListTasksResponse>>;
    listTaskSubStepComments: (
      params: Params.ListTaskSubStepCommentsParams,
    ) => Promise<Response<Schema.ListTaskSubStepCommentsResponse>>;
    listTaskSubStepFields: (
      params: Params.ListTaskSubStepFieldsParams,
    ) => Promise<Response<Schema.ListTaskSubStepFieldsResponse>>;
    removeTaskField: (
      params: Params.RemoveTaskFieldParams,
    ) => Promise<void>;
    updateTask: (
      params: Params.UpdateTaskParams,
    ) => Promise<Response<Schema.TaskResponse>>;
    updateTaskAssetFields: (
      params: Params.UpdateTaskAssetFieldsParams,
    ) => Promise<void>;
    updateTaskCustomField: (
      params: Params.UpdateTaskCustomFieldParams,
    ) => Promise<Response<Schema.TaskCustomField>>;
    updateTaskField: (
      params: Params.UpdateTaskFieldParams,
    ) => Promise<void>;
    updateTaskStructuredContent: (
      params: Params.UpdateTaskStructuredContentParams,
    ) => Promise<Response<Schema.ContentDetailsModel>>;
    updateTaskSubStep: (
      params: Params.UpdateTaskSubStepParams,
    ) => Promise<Response<Schema.TaskSubStep>>;
    updateTaskSubStepComment: (
      params: Params.UpdateTaskSubStepCommentParams,
    ) => Promise<Response<Schema.TaskSubStepCommentResponse>>;
    updateTaskSubStepExternalWork: (
      params: Params.UpdateTaskSubStepExternalWorkParams,
    ) => Promise<Response<Schema.TaskExternalWorkResponse>>;
  };
  taskStep: {
    updateTaskStep: (
      params: Params.UpdateTaskStepParams,
    ) => Promise<Response<Schema.TaskStep>>;
  };
  team: {
    getTeam: (
      params: Params.GetTeamParams,
    ) => Promise<Response<Schema.TeamWithUsers>>;
    listTeams: (
      params: Params.ListTeamsParams,
    ) => Promise<Response<Schema.ListTeamsResponse>>;
  };
  template: {
    getTemplate: (
      params: Params.GetTemplateParams,
    ) => Promise<Response<Schema.TemplateResponse>>;
    listTemplates: (
      params: Params.ListTemplatesParams,
    ) => Promise<Response<Schema.TemplateListResponse>>;
  };
  uploader: {
    completeMultipartUpload: (
      params: Params.CompleteMultipartUploadParams,
    ) => Promise<Response<Schema.CompleteMultipartUploadResponse>>;
    createMultipartUpload: (
      params: Params.CreateMultipartUploadParams,
    ) => Promise<Response<Schema.CreateMultipartUploadResponse>>;
    getMultipartUploadStatus: (
      params: Params.GetMultipartUploadStatusParams,
    ) => Promise<Response<Schema.GetMultipartUploadStatusResponse>>;
    getUploadUrl: (
      params: Params.GetUploadUrlParams,
    ) => Promise<Response<Schema.GetUploadUrlResponse>>;
  };
  user: {
    findUserByEmail: (
      params: Params.FindUserByEmailParams,
    ) => Promise<void>;
    getUser: (
      params: Params.GetUserParams,
    ) => Promise<Response<Schema.UserResponse>>;
    listUsers: (
      params: Params.ListUsersParams,
    ) => Promise<Response<Schema.UserListResponse>>;
  };
  workflow: {
    getWorkflow: (
      params: Params.GetWorkflowParams,
    ) => Promise<Response<Schema.WorkflowResponse>>;
    listWorkflows: (
      params: Params.ListWorkflowsParams,
    ) => Promise<Response<Schema.ListWorkflowsResponse>>;
  };
  workRequest: {
    addAttachmentToWorkRequest: (
      params: Params.AddAttachmentToWorkRequestParams,
    ) => Promise<Response<Schema.AttachmentResponse>>;
    addCommentToWorkRequest: (
      params: Params.AddCommentToWorkRequestParams,
    ) => Promise<Response<Schema.WorkRequestCommentResponse>>;
    createCampaignFromWorkRequest: (
      params: Params.CreateCampaignFromWorkRequestParams,
    ) => Promise<Response<Schema.WorkRequestCampaignResponse>>;
    createTaskFromWorkRequest: (
      params: Params.CreateTaskFromWorkRequestParams,
    ) => Promise<Response<Schema.WorkRequestTaskResponse>>;
    createWorkRequest: (
      params: Params.CreateWorkRequestParams,
    ) => Promise<Response<Schema.WorkRequestResponse>>;
    createWorkRequestCreativeAsset: (
      params: Params.CreateWorkRequestCreativeAssetParams,
    ) => Promise<Response<Schema.CreativeAssetResponse>>;
    deleteWorkRequestAttachment: (
      params: Params.DeleteWorkRequestAttachmentParams,
    ) => Promise<void>;
    deleteWorkRequestCreativeAsset: (
      params: Params.DeleteWorkRequestCreativeAssetParams,
    ) => Promise<void>;
    getWorkRequest: (
      params: Params.GetWorkRequestParams,
    ) => Promise<Response<Schema.WorkRequestResponse>>;
    getWorkRequestComment: (
      params: Params.GetWorkRequestCommentParams,
    ) => Promise<Response<Schema.WorkRequestCommentResponse>>;
    listWorkRequestApprovedAssets: (
      params: Params.ListWorkRequestApprovedAssetsParams,
    ) => Promise<Response<Schema.ListWorkRequestApprovedAssetsResponse>>;
    listWorkRequestComments: (
      params: Params.ListWorkRequestCommentsParams,
    ) => Promise<Response<Schema.ListWorkRequestCommentsResponse>>;
    listWorkRequestRelatedResources: (
      params: Params.ListWorkRequestRelatedResourcesParams,
    ) => Promise<Response<Schema.ListWorkRequestRelatedResourcesResponse>>;
    listWorkRequests: (
      params: Params.ListWorkRequestsParams,
    ) => Promise<Response<Schema.ListWorkRequestsResponse>>;
    updateWorkRequest: (
      params: Params.UpdateWorkRequestParams,
    ) => Promise<void>;
    updateWorkRequestFormField: (
      params: Params.UpdateWorkRequestFormFieldParams,
    ) => Promise<Response<Schema.WorkRequestRequestFormFieldUpdateResponse>>;
  };
}
