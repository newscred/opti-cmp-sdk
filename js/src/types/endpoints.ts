// Auto-generated - DO NOT EDIT

import type { Response } from "../request/types.js";
import type * as Params from "./params.js";
import type * as Schema from "./schema.js";

export interface APIEndpoints {
  asset: {
    /**
     * Find URL of an asset by ID
     *
     * `GET /asset-urls/{asset_id}`
     */
    getAssetUrl: (
      params: Params.GetAssetUrlParams,
    ) => Promise<void>;
  };
  brandCompliance: {
    /**
     * **Experimental:** Get the brand compliance details of a draft of an asset of a task.
     *
     * `GET /tasks/{task_id}/assets/{asset_id}/drafts/{draft_id}/brand-compliance`
     */
    getTaskAssetDraftBrandCompliance: (
      params: Params.GetTaskAssetDraftBrandComplianceParams,
    ) => Promise<Response<Schema.TaskAssetDraftBrandComplianceResponse>>;
    /**
     * **Experimental:** Get a list of brand compliance categories.
     *
     * `GET /brand-compliance/categories`
     */
    listBrandComplianceCategories: (
      params: Params.ListBrandComplianceCategoriesParams,
    ) => Promise<Response<Schema.ListBrandComplianceCategoriesResponse>>;
    /**
     * **Experimental:** Update the brand compliance details of a draft of an asset of a task.
     *
     * `PUT /tasks/{task_id}/assets/{asset_id}/drafts/{draft_id}/brand-compliance`
     */
    updateTaskAssetDraftBrandCompliance: (
      params: Params.UpdateTaskAssetDraftBrandComplianceParams,
    ) => Promise<Response<Schema.TaskAssetDraftBrandComplianceResponse>>;
  };
  campaign: {
    /**
     * **Experimental:** Add an attachment to the campaign.
     *
     * `POST /campaigns/{id}/attachments`
     */
    addAttachmentToCampaign: (
      params: Params.AddAttachmentToCampaignParams,
    ) => Promise<Response<Schema.AttachmentResponse>>;
    /**
     * **Experimental:** Post a comment on a campaign.
     *
     * `POST /campaigns/{id}/comments`
     */
    addCommentToCampaign: (
      params: Params.AddCommentToCampaignParams,
    ) => Promise<Response<Schema.CampaignCommentResponse>>;
    /**
     * **Experimental:** Add a field to a campaign.
     *
     * `POST /campaigns/{id}/fields`
     */
    addFieldToCampaign: (
      params: Params.AddFieldToCampaignParams,
    ) => Promise<Response<Schema.ObjectFieldCreateResponse>>;
    /**
     * **Experimental:** Create a campaign
     *
     * `POST /campaigns`
     */
    createCampaign: (
      params: Params.CreateCampaignParams,
    ) => Promise<Response<Schema.CampaignResponse>>;
    /**
     * Get a campaign
     *
     * `GET /campaigns/{id}`
     */
    getCampaign: (
      params: Params.GetCampaignParams,
    ) => Promise<Response<Schema.CampaignResponse>>;
    /**
     * Get brief of the campaign
     *
     * `GET /campaigns/{id}/brief`
     */
    getCampaignBrief: (
      params: Params.GetCampaignBriefParams,
    ) => Promise<Response<Schema.CampaignBriefResponse>>;
    /**
     * **Experimental:** Get the list of fields of a campaign.
     *
     * `GET /campaigns/{id}/fields`
     */
    listCampaignFields: (
      params: Params.ListCampaignFieldsParams,
    ) => Promise<Response<Schema.ListCampaignFieldsResponse>>;
    /**
     * **Experimental:** Get a list of campaigns.
     *
     * `GET /campaigns`
     */
    listCampaigns: (
      params: Params.ListCampaignsParams,
    ) => Promise<Response<Schema.ListCampaignsResponse>>;
    /**
     * **Experimental:** Update a campaign
     *
     * `PATCH /campaigns/{id}`
     */
    updateCampaign: (
      params: Params.UpdateCampaignParams,
    ) => Promise<void>;
    /**
     * **Experimental:** Update the field value of a campaign.
     *
     * `PUT /campaigns/{campaign_id}/fields/{field_id}`
     */
    updateCampaignField: (
      params: Params.UpdateCampaignFieldParams,
    ) => Promise<Response<Schema.CampaignFieldUpdateResponse>>;
  };
  compose: {
    /**
     * **Experimental:** Get a content schema. Returns 403 if Compose is not enabled for the organization.
     *
     * `GET /compose/content-schemas/{schema_id}`
     *
     * @experimental
     */
    getComposeContentSchema: (
      params: Params.GetComposeContentSchemaParams,
    ) => Promise<Response<Schema.ComposeContentSchemaResponse>>;
    /**
     * **Experimental:** Get all content schemas of the organization. The list is not paginated, so `pagination.next` and `pagination.previous` are always `null`. Returns 403 if Compose is not enabled for the organization.
     *
     * `GET /compose/content-schemas`
     *
     * @experimental
     */
    listComposeContentSchemas: (
      params: Params.ListComposeContentSchemasParams,
    ) => Promise<Response<Schema.ListComposeContentSchemasResponse>>;
  };
  contentGraph: {
    /**
     * **Experimental:** Get the list of content graph instances provisioned for the organization. Use this to discover the `instance_id` of the instance to query, before running a query against it.
     *
     * An organization with no content graph instance of the requested type gets an empty `data` list rather than an error.
     *
     * `GET /content-graphs`
     */
    listContentGraphs: (
      params: Params.ListContentGraphsParams,
    ) => Promise<Response<Schema.ContentGraphListResponse>>;
    /**
     * **Experimental:** Run a GraphQL query against a content graph instance. Get the `instance_id` from `GET /content-graphs`.
     *
     * Use this to check whether a given asset has synced to the graph, by querying for its identifier and seeing whether the graph returns it.
     *
     * **Reading the response.** A `200` carries the graph's JSON payload (`data`, and `extensions` when present) passed through unchanged. If the graph rejects a well-formed query, the endpoint returns the graph's own `4xx` status with its GraphQL `errors` body (for example `429` when the graph rate-limits). A `502` means the content graph is unavailable or returned a server error, and a `504` means it could not be reached in time — both are transient, so retry. A `500` is a server-side failure on our end and, unlike `502`/`504`, will not be resolved by retrying.
     *
     * `POST /content-graphs/{instance_id}/query`
     */
    queryContentGraph: (
      params: Params.QueryContentGraphParams,
    ) => Promise<Response<Schema.ContentGraphQueryResponse>>;
  };
  event: {
    /**
     * **Experimental:** Create an event.
     *
     * `POST /events`
     */
    createEvent: (
      params: Params.CreateEventParams,
    ) => Promise<Response<Schema.EventResponse>>;
    /**
     * **Experimental:** Get an event.
     *
     * `GET /events/{id}`
     */
    getEvent: (
      params: Params.GetEventParams,
    ) => Promise<Response<Schema.EventResponse>>;
    /**
     * **Experimental:** Get the list of fields of an event.
     *
     * `GET /events/{id}/fields`
     */
    listEventFields: (
      params: Params.ListEventFieldsParams,
    ) => Promise<Response<Schema.ListEventFieldsResponse>>;
    /**
     * **Experimental:** Get a list of events.
     *
     * `GET /events`
     */
    listEvents: (
      params: Params.ListEventsParams,
    ) => Promise<Response<Schema.ListEventsResponse>>;
    /**
     * **Experimental:** Update an event.
     *
     * `PATCH /events/{id}`
     */
    updateEvent: (
      params: Params.UpdateEventParams,
    ) => Promise<Response<Schema.EventResponse>>;
    /**
     * **Experimental:** Replace existing fields of an event.
     *
     * `PUT /events/{id}/fields`
     */
    updateEventFields: (
      params: Params.UpdateEventFieldsParams,
    ) => Promise<void>;
  };
  field: {
    /**
     * **Experimental:** Add a new field to an Organization.
     *
     * `POST /fields`
     */
    createField: (
      params: Params.CreateFieldParams,
    ) => Promise<Response<Schema.CreateFieldResponse>>;
    /**
     * **Experimental:** Create choices in a field.
     *
     * `POST /fields/{id}/choices`
     */
    createFieldChoices: (
      params: Params.CreateFieldChoicesParams,
    ) => Promise<Response<Schema.SettingsFieldChoiceCreateResponse>>;
    /**
     * **Experimental:** Delete a choice from a field. A request that deletes the last active choice of a field fails with a `400` error.
     *
     * `DELETE /fields/{field_id}/choices/{choice_id}`
     */
    deleteFieldChoice: (
      params: Params.DeleteFieldChoiceParams,
    ) => Promise<void>;
    /**
     * **Experimental:** Get the list of fields of an Organization.
     *
     * `GET /fields`
     */
    listFields: (
      params: Params.ListFieldsParams,
    ) => Promise<Response<Schema.ListFieldsResponse>>;
    /**
     * **Experimental:** Update a field in an Organization.
     *
     * `PATCH /fields/{id}`
     */
    updateField: (
      params: Params.UpdateFieldParams,
    ) => Promise<void>;
    /**
     * **Experimental:** Update a choice of a field.
     *
     * `PATCH /fields/{field_id}/choices/{choice_id}`
     */
    updateFieldChoice: (
      params: Params.UpdateFieldChoiceParams,
    ) => Promise<void>;
  };
  label: {
    /**
     * Get the list of label groups. Label groups are sorted by `name` in ascending order.
     *
     * `GET /label-groups`
     */
    listLabelGroups: (
      params: Params.ListLabelGroupsParams,
    ) => Promise<Response<Schema.ListLabelGroupsResponse>>;
  };
  library: {
    /**
     * **Experimental:** Grant asset access to users or teams
     *
     * `POST /assets/{asset_id}/permissions`
     */
    addAssetPermissions: (
      params: Params.AddAssetPermissionsParams,
    ) => Promise<void>;
    /**
     * **Experimental:** Grant folder access to users or teams
     *
     * `POST /folders/{id}/permissions`
     */
    addFolderPermissions: (
      params: Params.AddFolderPermissionsParams,
    ) => Promise<void>;
    /**
     * Create a new asset. Supports only `images`, `videos`, and `raw files`. See [Upload assets](https://docs.developers.optimizely.com/content-marketing-platform/docs/upload-assets) to upload an asset.
     *
     * `POST /assets`
     */
    createAsset: (
      params: Params.CreateAssetParams,
    ) => Promise<Response<Schema.AssetResponse>>;
    /**
     * **Experimental:** Add a new external asset lineage.
     *
     * `POST /assets/{asset_id}/lineages`
     */
    createAssetLineage: (
      params: Params.CreateAssetLineageParams,
    ) => Promise<Response<Schema.AssetLineageResponse>>;
    /**
     * Add a new version to a library asset. Supports adding versions to only `images`, `videos`, and `raw files` type assets. See [Upload assets](https://docs.developers.optimizely.com/content-marketing-platform/docs/upload-assets) to upload a version to a library asset.
     *
     * `POST /assets/{asset_id}/versions`
     */
    createAssetVersion: (
      params: Params.CreateAssetVersionParams,
    ) => Promise<Response<Schema.LibraryAssetVersionResponse>>;
    /**
     * Generates download URLs of files given file guid.
     *
     * `POST /file-urls`
     */
    createFileUrls: (
      params: Params.CreateFileUrlsParams,
    ) => Promise<Response<Schema.BatchFileUrlResponse>>;
    /**
     * Create a folder. Use the `parent_folder_id` field to create a nested folder.
     *
     * `POST /folders`
     */
    createFolder: (
      params: Params.CreateFolderParams,
    ) => Promise<Response<Schema.FolderResponse>>;
    /**
     * **Experimental:** Creates a structured content.
     *
     * `POST /structured-contents`
     */
    createStructuredContent: (
      params: Params.CreateStructuredContentParams,
    ) => Promise<Response<Schema.LibraryStructuredContent>>;
    /**
     * **Experimental:** Delete an asset lineage.
     *
     * `DELETE /assets/{asset_id}/lineages/{lineage_id}`
     */
    deleteAssetLineage: (
      params: Params.DeleteAssetLineageParams,
    ) => Promise<void>;
    /**
     * **Experimental:** Delete a folder.
     *
     * `DELETE /folders/{id}`
     */
    deleteFolder: (
      params: Params.DeleteFolderParams,
    ) => Promise<void>;
    /**
     * Delete an image
     *
     * `DELETE /images/{id}`
     */
    deleteImage: (
      params: Params.DeleteImageParams,
    ) => Promise<void>;
    /**
     * Delete a raw file
     *
     * `DELETE /raw-files/{id}`
     */
    deleteRawFile: (
      params: Params.DeleteRawFileParams,
    ) => Promise<void>;
    /**
     * Delete a video
     *
     * `DELETE /videos/{id}`
     */
    deleteVideo: (
      params: Params.DeleteVideoParams,
    ) => Promise<void>;
    /**
     * Get an article
     *
     * `GET /articles/{id}`
     */
    getArticle: (
      params: Params.GetArticleParams,
    ) => Promise<Response<Schema.LibraryArticle>>;
    /**
     * **Experimental:** Get a compose content from the library, with the field values of its published version. Returns 404 if the content was never published.
     *
     * `GET /compose-contents/{content_id}`
     *
     * @experimental
     */
    getComposeContent: (
      params: Params.GetComposeContentParams,
    ) => Promise<Response<Schema.LibraryComposeContentResponse>>;
    /**
     * Get a folder
     *
     * `GET /folders/{id}`
     */
    getFolder: (
      params: Params.GetFolderParams,
    ) => Promise<Response<Schema.FolderResponse>>;
    /**
     * Get an image
     *
     * `GET /images/{id}`
     */
    getImage: (
      params: Params.GetImageParams,
    ) => Promise<Response<Schema.LibraryImage>>;
    /**
     * Get a raw file
     *
     * `GET /raw-files/{id}`
     */
    getRawFile: (
      params: Params.GetRawFileParams,
    ) => Promise<Response<Schema.LibraryRawFile>>;
    /**
     * **Experimental:** Get a rendition given its `id`.
     *
     * `GET /renditions/{id}`
     */
    getRendition: (
      params: Params.GetRenditionParams,
    ) => Promise<Response<Schema.DetailedAssetRenditionResponse>>;
    /**
     * **Experimental:** Get a rendition configuration of an organization given its `id`.
     *
     * `GET /rendition-configs/{id}`
     */
    getRenditionConfig: (
      params: Params.GetRenditionConfigParams,
    ) => Promise<Response<Schema.RenditionConfigResponse>>;
    /**
     * **Experimental:** Get the structured content.
     *
     * `GET /structured-contents/{id}`
     */
    getStructuredContent: (
      params: Params.GetStructuredContentParams,
    ) => Promise<Response<Schema.LibraryStructuredContent>>;
    /**
     * Get a video
     *
     * `GET /videos/{id}`
     */
    getVideo: (
      params: Params.GetVideoParams,
    ) => Promise<Response<Schema.LibraryVideo>>;
    /**
     * **Experimental:** Get the list of fields of an asset.
     *
     * `GET /assets/{asset_id}/fields`
     */
    listAssetFields: (
      params: Params.ListAssetFieldsParams,
    ) => Promise<Response<Schema.ListAssetFieldsResponse>>;
    /**
     * **Experimental:** Get the list of asset lineage.
     *
     * `GET /asset-lineages`
     */
    listAssetLineages: (
      params: Params.ListAssetLineagesParams,
    ) => Promise<Response<Schema.ListAssetLineagesResponse>>;
    /**
     * **Experimental:** List of entities that have permission to access the asset
     *
     * `GET /assets/{asset_id}/permissions`
     */
    listAssetPermissions: (
      params: Params.ListAssetPermissionsParams,
    ) => Promise<Response<Schema.AssetPermissionListResponseItem>>;
    /**
     * **Experimental:** Get the renditions of an asset given its `id`.
     *
     * `GET /assets/{asset_id}/renditions`
     */
    listAssetRenditions: (
      params: Params.ListAssetRenditionsParams,
    ) => Promise<Response<Schema.ListAssetRenditionsResponse>>;
    /**
     * Get the list of assets. Assets are sorted by `modified_at` in descending order.
     * To get assets that are not inside any folder, pass `include_subfolder_assets=false` in the query param.
     *
     * `GET /assets`
     */
    listAssets: (
      params: Params.ListAssetsParams,
    ) => Promise<Response<Schema.ListAssetsResponse>>;
    /**
     * **Experimental:** List of entities that have permission to access the folder
     *
     * `GET /folders/{id}/permissions`
     */
    listFolderPermissions: (
      params: Params.ListFolderPermissionsParams,
    ) => Promise<Response<Schema.FolderPermissionListResponseItem>>;
    /**
     * Get the list of folders sorted by `modified_at` in descending order.
     *
     * `GET /folders`
     */
    listFolders: (
      params: Params.ListFoldersParams,
    ) => Promise<Response<Schema.ListFoldersResponse>>;
    /**
     * Returns a paginated list of related assets for the specified asset. Only supported for asset types: article, image, video, raw_file, and structured_content.
     *
     * `GET /assets/{asset_id}/related-assets`
     */
    listRelatedAssets: (
      params: Params.ListRelatedAssetsParams,
    ) => Promise<Response<Schema.RelatedAssetsListResponse>>;
    /**
     * **Experimental:** Remove accessor's access from an asset
     *
     * `DELETE /assets/{asset_id}/permissions/{accessor_id}`
     */
    removeAssetPermission: (
      params: Params.RemoveAssetPermissionParams,
    ) => Promise<void>;
    /**
     * **Experimental:** Remove accessor's access from a folder
     *
     * `DELETE /folders/{id}/permissions/{accessor_id}`
     */
    removeFolderPermission: (
      params: Params.RemoveFolderPermissionParams,
    ) => Promise<void>;
    /**
     * Replaces all related assets for the specified asset. Only supported for asset types: article, image, video, raw_file, and structured_content.
     *
     * `PUT /assets/{asset_id}/related-assets`
     */
    replaceRelatedAssets: (
      params: Params.ReplaceRelatedAssetsParams,
    ) => Promise<Response<Schema.ReplaceRelatedAssetsResponse>>;
    /**
     * **Experimental:** Update the field value of an asset.
     *
     * `PUT /assets/{asset_id}/fields/{field_id}`
     */
    updateAssetField: (
      params: Params.UpdateAssetFieldParams,
    ) => Promise<Response<Schema.AssetFieldUpdateResponse>>;
    /**
     * **Experimental:** Replace existing fields of an asset.
     *
     * `PUT /assets/{asset_id}/fields`
     */
    updateAssetFields: (
      params: Params.UpdateAssetFieldsParams,
    ) => Promise<Response<Schema.UpdateAssetFieldsResponse>>;
    /**
     * **Experimental:** Update accessor's access level and ownership of asset
     *
     * `PATCH /assets/{asset_id}/permissions/{accessor_id}`
     */
    updateAssetPermission: (
      params: Params.UpdateAssetPermissionParams,
    ) => Promise<void>;
    /**
     * **Experimental:** Update a folder's name and/or parent.
     *
     * `PATCH /folders/{id}`
     */
    updateFolder: (
      params: Params.UpdateFolderParams,
    ) => Promise<Response<Schema.FolderResponse>>;
    /**
     * **Experimental:** Update accessor's access level and ownership of folder
     *
     * `PATCH /folders/{id}/permissions/{accessor_id}`
     */
    updateFolderPermission: (
      params: Params.UpdateFolderPermissionParams,
    ) => Promise<void>;
    /**
     * Updates an image
     *
     * `PATCH /images/{id}`
     */
    updateImage: (
      params: Params.UpdateImageParams,
    ) => Promise<Response<Schema.LibraryImage>>;
    /**
     * Updates a raw file
     *
     * `PATCH /raw-files/{id}`
     */
    updateRawFile: (
      params: Params.UpdateRawFileParams,
    ) => Promise<Response<Schema.LibraryRawFile>>;
    /**
     * **Experimental:** Updates a structured content.
     *
     * `PATCH /structured-contents/{id}`
     */
    updateStructuredContent: (
      params: Params.UpdateStructuredContentParams,
    ) => Promise<Response<Schema.LibraryStructuredContent>>;
    /**
     * Updates a video
     *
     * `PATCH /videos/{id}`
     */
    updateVideo: (
      params: Params.UpdateVideoParams,
    ) => Promise<Response<Schema.LibraryVideo>>;
  };
  milestone: {
    /**
     * **Experimental:** Creates a milestone. Defaults to the organization's campaign if campaign_id is omitted. Tasks are validated against the selected campaign and must belong to it.
     *
     * `POST /milestones`
     */
    createMilestone: (
      params: Params.CreateMilestoneParams,
    ) => Promise<Response<Schema.MilestoneResponse>>;
    /**
     * **Experimental:** Get a milestone by ID.
     *
     * `GET /milestones/{id}`
     */
    getMilestone: (
      params: Params.GetMilestoneParams,
    ) => Promise<Response<Schema.MilestoneResponse>>;
    /**
     * **Experimental:** Get a list of milestones.
     *
     * `GET /milestones`
     */
    listMilestones: (
      params: Params.ListMilestonesParams,
    ) => Promise<Response<Schema.ListMilestonesResponse>>;
    /**
     * **Experimental:** Updates a milestone. All fields are optional. Only provided fields will be updated. Note: If tasks array is empty, it will remove all task associations for that milestone.
     *
     * `PATCH /milestones/{id}`
     */
    updateMilestone: (
      params: Params.UpdateMilestoneParams,
    ) => Promise<Response<Schema.MilestoneResponse>>;
  };
  publishing: {
    /**
     * Create asset metadata in bulk.
     *
     * `POST /publishing-events/{publishing_event_id}/publishing-metadata`
     */
    bulkCreatePublishingEventMetadata: (
      params: Params.BulkCreatePublishingEventMetadataParams,
    ) => Promise<Response<Schema.PublishingEventMetadataBulkCreateResponse>>;
    /**
     * Post a progress update describing what is happening while a publishing event is being processed, so that CMP can show it to the user.
     *
     * `POST /publishing-events/{publishing_event_id}/progress-updates`
     */
    createPublishingEventProgressUpdate: (
      params: Params.CreatePublishingEventProgressUpdateParams,
    ) => Promise<void>;
    /**
     * Get the publishing event by ID
     *
     * `GET /publishing-events/{publishing_event_id}`
     */
    getPublishingEvent: (
      params: Params.GetPublishingEventParams,
    ) => Promise<Response<Schema.PublishingEventResponse>>;
    /**
     * Get publishing metadata.
     *
     * `GET /publishing-events/{publishing_event_id}/assets/{asset_id}/publishing-metadata/{publishing_metadata_id}`
     */
    getPublishingEventAssetMetadata: (
      params: Params.GetPublishingEventAssetMetadataParams,
    ) => Promise<Response<Schema.PublishingEventMetadataResponse>>;
    /**
     * **Experimental:** Get the list of publishing channels available in the organization. Use this to populate a channel picker before creating a publishing intent for a task.
     *
     * `GET /publishing-channels`
     */
    listPublishingChannels: (
      params: Params.ListPublishingChannelsParams,
    ) => Promise<Response<Schema.PublishingChannelListResponse>>;
    /**
     * Get the list of publishing metadata for each asset of a publishing event.
     *
     * `GET /publishing-events/{publishing_event_id}/publishing-metadata`
     */
    listPublishingEventMetadata: (
      params: Params.ListPublishingEventMetadataParams,
    ) => Promise<Response<Schema.PublishingEventMetadataListResponse>>;
    /**
     * Get the progress updates posted for a publishing event, newest first.
     *
     * `GET /publishing-events/{publishing_event_id}/progress-updates`
     */
    listPublishingEventProgressUpdates: (
      params: Params.ListPublishingEventProgressUpdatesParams,
    ) => Promise<Response<Schema.PublishingEventProgressUpdateListResponse>>;
  };
  settings: {
    /**
     * **Experimental:** Get settings.
     *
     * `GET /settings`
     */
    getSettings: (
      params: Params.GetSettingsParams,
    ) => Promise<Response<Schema.Settings>>;
    /**
     * **Experimental:** Create or update settings.
     * ## Import behavior:
     * - Label:
     *     - If a label-group with the same name exists, the label-group is updated.
     *       - Labels are merged.
     *       - If a label with the same name exists, it is overwritten.
     *     - If no label-group with the same name exists, a new label-group is created.
     * - Custom Fields:
     *     - If a custom field with the same name exists, the custom field is updated and its choices are merged.
     *     - If no custom field with the same name exists, a new custom field will be created.
     * - Workflows:
     *     - If the query param overwrite_workflows=true the existing workflows are overwritten.
     *       Otherwise, a new workflow will be created where the prefix `Copy of` is added to the name of the workflow.
     * - Webhooks:
     *     - If a webhook with the same name exists, the webhook is updated and its events are merged.
     *     - If multiple webhooks with the same name exist, a new webhook is created.
     *     - If no webhook with the same name exists, a new webhook is created.
     * - Apps:
     *     - If an app with the same name exists, the app is updated and its authorization callback URLs are merged.
     *     - If multiple apps with the same name exist, a new app is created.
     *     - If no app with the same name exists, a new app is created.
     * - Templates:
     *     - If a template with the same name exists, the template is updated.
     *     - If multiple templates with the same name exist, a new template is created.
     *     - If no template with the same name exists, a new template is created.
     * - Routing Rules:
     *     - If one single routing rule with the same name exists, the routing rule is updated.
     *     - If multiple routing rules with the same name exist, a new routing rule is created.
     *     - If no routing rule with the same name exists, a new routing rule is created.
     *
     * `POST /settings`
     */
    updateSettings: (
      params: Params.UpdateSettingsParams,
    ) => Promise<Response<Schema.SettingsUpdateResponse>>;
  };
  structuredContent: {
    /**
     * **Experimental:** Acknowledge content preview. Content preview can be acknowledged only once. So make sure you are acknowledging only the content previews targeted for your integration. Otherwise it will stall acknowledgment from other integrations.
     *
     * `POST /structured-content/contents/{content_id}/versions/{version_id}/previews/{preview_id}/acknowledge`
     */
    acknowledgeSCContentPreview: (
      params: Params.AcknowledgeSCContentPreviewParams,
    ) => Promise<Response<unknown>>;
    /**
     * **Experimental:** Complete content preview.
     *
     * `POST /structured-content/contents/{content_id}/versions/{version_id}/previews/{preview_id}/complete`
     */
    completeSCContentPreview: (
      params: Params.CompleteSCContentPreviewParams,
    ) => Promise<Response<unknown>>;
    /**
     * **Experimental:** Create content type.
     *
     * `POST /structured-content/content-types`
     */
    createSCContentType: (
      params: Params.CreateSCContentTypeParams,
    ) => Promise<Response<Schema.SCContentTypeCreateResponse>>;
    /**
     * **Experimental:** Create a new managed migration job.
     *
     * `POST /structured-content/content-types/{content_type_id}/managed-migrations`
     */
    createSCContentTypeManagedMigration: (
      params: Params.CreateSCContentTypeManagedMigrationParams,
    ) => Promise<Response<Schema.CreateSCContentTypeManagedMigrationResponse>>;
    /**
     * **Experimental:** Add Content Type Version.
     *
     * `POST /structured-content/content-types/{content_type_id}/versions`
     */
    createSCContentTypeVersion: (
      params: Params.CreateSCContentTypeVersionParams,
    ) => Promise<Response<Schema.SCContentTypeCreateResponse>>;
    /**
     * **Experimental:** Delete a managed migration job with not_started status.
     *
     * `DELETE /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}`
     */
    deleteSCContentTypeManagedMigration: (
      params: Params.DeleteSCContentTypeManagedMigrationParams,
    ) => Promise<void>;
    /**
     * **Experimental:** Get content type.
     *
     * `GET /structured-content/content-types/{content_type_id}`
     */
    getSCContentType: (
      params: Params.GetSCContentTypeParams,
    ) => Promise<Response<Schema.SCContentType>>;
    /**
     * **Experimental:** Get details of a specific managed migration job.
     *
     * `GET /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}`
     */
    getSCContentTypeManagedMigration: (
      params: Params.GetSCContentTypeManagedMigrationParams,
    ) => Promise<Response<Schema.SCContentTypeManagedMigrationResponse>>;
    /**
     * **Experimental:** Get content type version.
     *
     * `GET /structured-content/content-types/{content_type_id}/versions/{version_id}`
     */
    getSCContentTypeVersion: (
      params: Params.GetSCContentTypeVersionParams,
    ) => Promise<Response<Schema.SCContentTypeVersion>>;
    /**
     * **Experimental:** Retrieves a list of managed migration jobs for a specific content type.
     *
     * `GET /structured-content/content-types/{content_type_id}/managed-migrations`
     */
    listSCContentTypeManagedMigrations: (
      params: Params.ListSCContentTypeManagedMigrationsParams,
    ) => Promise<Response<Schema.SCContentTypeManagedMigrationResponse[]>>;
    /**
     * **Experimental:** Get content types.
     *
     * `GET /structured-content/content-types`
     */
    listSCContentTypes: (
      params: Params.ListSCContentTypesParams,
    ) => Promise<Response<Schema.BaseContentTypeModel[]>>;
    /**
     * **Experimental:** Get content type versions.
     *
     * `GET /structured-content/content-types/{content_type_id}/versions`
     */
    listSCContentTypeVersions: (
      params: Params.ListSCContentTypeVersionsParams,
    ) => Promise<Response<Schema.BaseContentTypeVersionModel[]>>;
    /**
     * Migrate content to a specific content type version. **Experimental:**
     *
     * `POST /structured-content/contents/{content_id}/migration`
     */
    migrateSCContent: (
      params: Params.MigrateSCContentParams,
    ) => Promise<Response<Schema.MigrateSCContentResponse>>;
    /**
     * **Experimental:** Start managed migration job.
     *
     * `POST /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}/start`
     */
    startSCContentTypeManagedMigration: (
      params: Params.StartSCContentTypeManagedMigrationParams,
    ) => Promise<Response<Schema.SCContentTypeManagedMigrationStartResponse>>;
    /**
     * **Experimental:** Update content type.
     *
     * `POST /structured-content/content-types/{content_type_id}`
     */
    updateSCContentType: (
      params: Params.UpdateSCContentTypeParams,
    ) => Promise<Response<Schema.SCContentTypeUpdateResponse>>;
    /**
     * **Experimental:** Update a managed migration job.
     *
     * `PATCH /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}`
     */
    updateSCContentTypeManagedMigration: (
      params: Params.UpdateSCContentTypeManagedMigrationParams,
    ) => Promise<Response<Schema.UpdateSCContentTypeManagedMigrationResponse>>;
    /**
     * **Experimental:** Check managed migration possible or not.
     *
     * `POST /structured-content/content-types/{content_type_id}/managed-migrations/validate`
     */
    validateSCContentTypeManagedMigration: (
      params: Params.ValidateSCContentTypeManagedMigrationParams,
    ) => Promise<Response<Schema.ValidateSCContentTypeManagedMigrationResponse>>;
  };
  task: {
    /**
     * This API allows you to add an asset to a task. There currently two mechanisms to add an asset - either by uploading an asset first and then calling this API (for only `images`, `videos`, and `raw files`), or by adding an asset from the library. The 'type' parameter in the request body will determine the mechanism used to add the asset. See [Upload assets](https://docs.developers.optimizely.com/content-marketing-platform/docs/upload-assets) to upload an asset.
     *
     * `POST /tasks/{id}/assets`
     */
    addAssetToTask: (
      params: Params.AddAssetToTaskParams,
    ) => Promise<Response<Schema.TaskAssetResponse>>;
    /**
     * **Experimental:** Post a comment on a task.
     *
     * `POST /tasks/{task_id}/comments`
     */
    addCommentToTask: (
      params: Params.AddCommentToTaskParams,
    ) => Promise<Response<Schema.TaskCommentResponse>>;
    /**
     * Add a comment to the latest version of a task asset. Supports adding comments to only **images**, **videos** and **raw files** type assets.
     *
     * `POST /tasks/{task_id}/assets/{asset_id}/comments`
     */
    addCommentToTaskAsset: (
      params: Params.AddCommentToTaskAssetParams,
    ) => Promise<Response<Schema.TaskAssetCommentResponse>>;
    /**
     * Create a task substep comment
     *
     * `POST /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments`
     */
    addCommentToTaskSubStep: (
      params: Params.AddCommentToTaskSubStepParams,
    ) => Promise<Response<Schema.TaskSubStepCommentResponse>>;
    /**
     * Add a new draft to a task asset. You can add a draft to only `images`, `videos`, and `raw files` type task assets. See [Upload assets](https://docs.developers.optimizely.com/content-marketing-platform/docs/upload-assets) to upload an asset draft.
     *
     * `POST /tasks/{task_id}/assets/{asset_id}/drafts`
     */
    addDraftToTaskAsset: (
      params: Params.AddDraftToTaskAssetParams,
    ) => Promise<Response<Schema.TaskAssetDraftResponse>>;
    /**
     * **Experimental:** Add a field to a task.
     *
     * `POST /tasks/{task_id}/fields`
     */
    addFieldToTask: (
      params: Params.AddFieldToTaskParams,
    ) => Promise<Response<Schema.ObjectFieldCreateResponse>>;
    /**
     * **Experimental:** Add a new structured content to task
     *
     * `POST /tasks/{task_id}/structured-contents`
     */
    addStructuredContentToTask: (
      params: Params.AddStructuredContentToTaskParams,
    ) => Promise<Response<Schema.ContentDetailsModel>>;
    /**
     * **Experimental:** Adds a URL to a task.
     *
     * `POST /tasks/{task_id}/urls`
     */
    addUrlToTask: (
      params: Params.AddUrlToTaskParams,
    ) => Promise<Response<Schema.TaskUrlResponse>>;
    /**
     * **Experimental:** Create a task
     *
     * `POST /tasks`
     */
    createTask: (
      params: Params.CreateTaskParams,
    ) => Promise<Response<Schema.TaskResponse>>;
    /**
     * **Experimental:** Create a compose content in a task. The new content has no field values and the status `draft`. Returns 403 if Compose is not enabled for the organization, and 422 if the task is completed or archived.
     *
     * `POST /tasks/{task_id}/compose-contents`
     *
     * @experimental
     */
    createTaskComposeContent: (
      params: Params.CreateTaskComposeContentParams,
    ) => Promise<Response<Schema.TaskComposeContentResponse>>;
    /**
     * **Experimental:** Create a publishing intent for a task on the specified publishing channel.
     *
     * `POST /tasks/{task_id}/publishing-intents`
     */
    createTaskPublishingIntent: (
      params: Params.CreateTaskPublishingIntentParams,
    ) => Promise<Response<Schema.TaskPublishingIntentResponse>>;
    /**
     * **Experimental:** Create structured content draft using a task.
     *
     * `POST /tasks/{task_id}/structured-contents/{content_id}/drafts`
     */
    createTaskStructuredContentDraft: (
      params: Params.CreateTaskStructuredContentDraftParams,
    ) => Promise<Response<Record<string, unknown>>>;
    /**
     * **Experimental:** Add a preview URL to the latest version of a structured content in a task. The task must be neither archived nor completed.
     *
     * `POST /tasks/{task_id}/structured-contents/{content_id}/previews`
     */
    createTaskStructuredContentPreview: (
      params: Params.CreateTaskStructuredContentPreviewParams,
    ) => Promise<void>;
    /**
     * **Experimental:** Remove a compose content from a task. The content is not deleted from Compose, and references to it from other content are not checked. Returns 422 if the task is completed or archived.
     *
     * `DELETE /tasks/{task_id}/compose-contents/{content_id}`
     *
     * @experimental
     */
    deleteTaskComposeContent: (
      params: Params.DeleteTaskComposeContentParams,
    ) => Promise<void>;
    /**
     * **Experimental:** Delete a structured content from task
     *
     * `DELETE /tasks/{task_id}/structured-contents/{content_id}`
     */
    deleteTaskStructuredContent: (
      params: Params.DeleteTaskStructuredContentParams,
    ) => Promise<void>;
    /**
     * Delete a task substep comment
     *
     * `DELETE /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}`
     */
    deleteTaskSubStepComment: (
      params: Params.DeleteTaskSubStepCommentParams,
    ) => Promise<void>;
    /**
     * Get a task
     *
     * `GET /tasks/{id}`
     */
    getTask: (
      params: Params.GetTaskParams,
    ) => Promise<Response<Schema.TaskResponse>>;
    /**
     * Get an article associated with the task identified by `task_id`.
     *
     * `GET /tasks/{task_id}/articles/{article_id}`
     */
    getTaskArticle: (
      params: Params.GetTaskArticleParams,
    ) => Promise<Response<Schema.TaskArticle>>;
    /**
     * Get brief of the task
     *
     * `GET /tasks/{id}/brief`
     */
    getTaskBrief: (
      params: Params.GetTaskBriefParams,
    ) => Promise<Response<Schema.TaskBriefResponse>>;
    /**
     * **Experimental:** Get a compose content of a task, with its current (draft) field values. Reference and embedded schema fields are returned as links only. An embedded child content can also be fetched by its own id when its parent content belongs to the task.
     *
     * `GET /tasks/{task_id}/compose-contents/{content_id}`
     *
     * @experimental
     */
    getTaskComposeContent: (
      params: Params.GetTaskComposeContentParams,
    ) => Promise<Response<Schema.TaskComposeContentResponse>>;
    /**
     * Get a custom field of a task
     *
     * `GET /tasks/{task_id}/custom-fields/{custom_field_id}`
     */
    getTaskCustomField: (
      params: Params.GetTaskCustomFieldParams,
    ) => Promise<Response<Schema.TaskCustomField>>;
    /**
     * Get an image associated with the task identified by `task_id`.
     *
     * `GET /tasks/{task_id}/images/{image_id}`
     */
    getTaskImage: (
      params: Params.GetTaskImageParams,
    ) => Promise<Response<Schema.TaskImage>>;
    /**
     * Get a raw file associated with the task identified by `task_id`.
     *
     * `GET /tasks/{task_id}/raw-files/{raw_file_id}`
     */
    getTaskRawFile: (
      params: Params.GetTaskRawFileParams,
    ) => Promise<Response<Schema.TaskRawFile>>;
    /**
     * **Experimental:** Get a task structured content by its guid.
     *
     * `GET /tasks/{task_id}/structured-contents/{content_id}`
     */
    getTaskStructuredContent: (
      params: Params.GetTaskStructuredContentParams,
    ) => Promise<Response<Schema.ContentDetailsModel>>;
    /**
     * Get the substep of a task
     *
     * `GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}`
     */
    getTaskSubStep: (
      params: Params.GetTaskSubStepParams,
    ) => Promise<Response<Schema.TaskSubStep>>;
    /**
     * Get a task substep comment
     *
     * `GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}`
     */
    getTaskSubStepComment: (
      params: Params.GetTaskSubStepCommentParams,
    ) => Promise<Response<Schema.TaskSubStepCommentResponse>>;
    /**
     * Get the external work information of a task external substep
     *
     * `GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/external-work`
     */
    getTaskSubStepExternalWork: (
      params: Params.GetTaskSubStepExternalWorkParams,
    ) => Promise<Response<Schema.TaskExternalWorkResponse>>;
    /**
     * Get a video associated with the task identified by task_id
     *
     * `GET /tasks/{task_id}/videos/{video_id}`
     */
    getTaskVideo: (
      params: Params.GetTaskVideoParams,
    ) => Promise<Response<Schema.TaskVideo>>;
    /**
     * List all the comments for the latest version of a task asset.
     *
     * `GET /tasks/{task_id}/assets/{asset_id}/comments`
     */
    listTaskAssetComments: (
      params: Params.ListTaskAssetCommentsParams,
    ) => Promise<Response<Schema.ListTaskAssetCommentsResponse>>;
    /**
     * **Experimental:** Get the list of drafts on an asset of a task.
     *
     * `GET /tasks/{task_id}/assets/{asset_id}/drafts`
     */
    listTaskAssetDrafts: (
      params: Params.ListTaskAssetDraftsParams,
    ) => Promise<Response<Schema.ListTaskAssetDraftsResponse>>;
    /**
     * **Experimental:** List all the fields of a task asset.
     *
     * `GET /tasks/{task_id}/assets/{asset_id}/fields`
     */
    listTaskAssetFields: (
      params: Params.ListTaskAssetFieldsParams,
    ) => Promise<Response<Schema.ListTaskAssetFieldsResponse>>;
    /**
     * Get the list of the assets or contents of a task
     *
     * `GET /tasks/{id}/assets`
     */
    listTaskAssets: (
      params: Params.ListTaskAssetsParams,
    ) => Promise<Response<Schema.ListTaskAssetsResponse>>;
    /**
     * Get the list of the attachments of a task
     *
     * `GET /tasks/{id}/attachments`
     */
    listTaskAttachments: (
      params: Params.ListTaskAttachmentsParams,
    ) => Promise<Response<Schema.ListTaskAttachmentsResponse>>;
    /**
     * Get the list of the choices of a custom field in a task.
     *
     * `GET /tasks/{task_id}/custom-fields/{custom_field_id}/choices`
     */
    listTaskCustomFieldChoices: (
      params: Params.ListTaskCustomFieldChoicesParams,
    ) => Promise<Response<Schema.ListTaskCustomFieldChoicesResponse>>;
    /**
     * Get the list of custom fields added to a task.
     *
     * `GET /tasks/{id}/custom-fields`
     */
    listTaskCustomFields: (
      params: Params.ListTaskCustomFieldsParams,
    ) => Promise<Response<Schema.ListTaskCustomFieldsResponse>>;
    /**
     * **Experimental:** Get the list of fields of a task.
     *
     * `GET /tasks/{task_id}/fields`
     */
    listTaskFields: (
      params: Params.ListTaskFieldsParams,
    ) => Promise<Response<Schema.ListTaskFieldsResponse>>;
    /**
     * **Experimental:** Get a list of tasks.
     *
     * Pick one style of date filter per request. Either use `start_date` and `due_date`, or use the `__from` and `__to` filters. A request that mixes the two styles returns a `400` error.
     *
     * `GET /tasks`
     */
    listTasks: (
      params: Params.ListTasksParams,
    ) => Promise<Response<Schema.ListTasksResponse>>;
    /**
     * **Experimental:** Get the list of versions of a task structured content, newest first. These are the versions that CMP lists for the content.
     *
     * `GET /tasks/{task_id}/structured-contents/{content_id}/versions`
     */
    listTaskStructuredContentVersions: (
      params: Params.ListTaskStructuredContentVersionsParams,
    ) => Promise<Response<Schema.TaskStructuredContentVersionListResponse>>;
    /**
     * Get the list of the comments of a task substep
     *
     * `GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments`
     */
    listTaskSubStepComments: (
      params: Params.ListTaskSubStepCommentsParams,
    ) => Promise<Response<Schema.ListTaskSubStepCommentsResponse>>;
    /**
     * **Experimental:** Get the list of fields of a substep of a task.
     *
     * `GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/fields`
     */
    listTaskSubStepFields: (
      params: Params.ListTaskSubStepFieldsParams,
    ) => Promise<Response<Schema.ListTaskSubStepFieldsResponse>>;
    /**
     * **Experimental:** Remove a field of a task.
     *
     * `DELETE /tasks/{task_id}/fields/{field_id}`
     */
    removeTaskField: (
      params: Params.RemoveTaskFieldParams,
    ) => Promise<void>;
    /**
     * Update a task
     *
     * `PATCH /tasks/{id}`
     */
    updateTask: (
      params: Params.UpdateTaskParams,
    ) => Promise<Response<Schema.TaskResponse>>;
    /**
     * **Experimental:** Replace fields of an asset in a task.
     *
     * `PUT /tasks/{task_id}/assets/{asset_id}/fields`
     */
    updateTaskAssetFields: (
      params: Params.UpdateTaskAssetFieldsParams,
    ) => Promise<void>;
    /**
     * **Experimental:** Update the field values of a compose content of a task. The given fields are merged into the existing fields. Field values are checked against the content schema structure; value rules (`validations`) are not checked on update. Returns 422 if the task is completed or archived.
     *
     * `PATCH /tasks/{task_id}/compose-contents/{content_id}`
     *
     * @experimental
     */
    updateTaskComposeContent: (
      params: Params.UpdateTaskComposeContentParams,
    ) => Promise<Response<Schema.TaskComposeContentResponse>>;
    /**
     * Update a custom field of a task.
     * You can update the following types:  `text_field`, `multi_line_text_field`, `checkboxes`, `dropdown`, `multi_select_dropdown`, `multiple_choice`, `date_field`, `rich_text_field`
     *
     * `PATCH /tasks/{task_id}/custom-fields/{custom_field_id}`
     */
    updateTaskCustomField: (
      params: Params.UpdateTaskCustomFieldParams,
    ) => Promise<Response<Schema.TaskCustomField>>;
    /**
     * **Experimental:** Update a field of a task.
     *
     * `PUT /tasks/{task_id}/fields/{field_id}`
     */
    updateTaskField: (
      params: Params.UpdateTaskFieldParams,
    ) => Promise<void>;
    /**
     * **Experimental:** Update structured content using a task.
     *
     * `PATCH /tasks/{task_id}/structured-contents/{content_id}`
     */
    updateTaskStructuredContent: (
      params: Params.UpdateTaskStructuredContentParams,
    ) => Promise<Response<Schema.ContentDetailsModel>>;
    /**
     * Update a substep of a step in a task
     *
     * `PATCH /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}`
     */
    updateTaskSubStep: (
      params: Params.UpdateTaskSubStepParams,
    ) => Promise<Response<Schema.TaskSubStep>>;
    /**
     * Update a task substep comment
     *
     * `PATCH /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}`
     */
    updateTaskSubStepComment: (
      params: Params.UpdateTaskSubStepCommentParams,
    ) => Promise<Response<Schema.TaskSubStepCommentResponse>>;
    /**
     * Update the external work information of a task external substep
     *
     * `PATCH /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/external-work`
     */
    updateTaskSubStepExternalWork: (
      params: Params.UpdateTaskSubStepExternalWorkParams,
    ) => Promise<Response<Schema.TaskExternalWorkResponse>>;
  };
  taskStep: {
    /**
     * **Experimental:** Update a step in a task
     *
     * `PATCH /tasks/{task_id}/steps/{step_id}`
     */
    updateTaskStep: (
      params: Params.UpdateTaskStepParams,
    ) => Promise<Response<Schema.TaskStep>>;
  };
  team: {
    /**
     * **Experimental:** Get a team.
     *
     * `GET /teams/{id}`
     */
    getTeam: (
      params: Params.GetTeamParams,
    ) => Promise<Response<Schema.TeamWithUsers>>;
    /**
     * **Experimental:** Get a list of teams.
     *
     * `GET /teams`
     */
    listTeams: (
      params: Params.ListTeamsParams,
    ) => Promise<Response<Schema.ListTeamsResponse>>;
  };
  template: {
    /**
     * **Experimental:** Get a template by ID.
     *
     * `GET /templates/{template_id}`
     */
    getTemplate: (
      params: Params.GetTemplateParams,
    ) => Promise<Response<Schema.TemplateResponse>>;
    /**
     * **Experimental:** Get a list of templates.
     *
     * `GET /templates`
     */
    listTemplates: (
      params: Params.ListTemplatesParams,
    ) => Promise<Response<Schema.TemplateListResponse>>;
  };
  uploader: {
    /**
     * **Experimental:** Initiate completion of a multipart upload after all parts have been uploaded
     *
     * `POST /multipart-uploads/{id}/complete`
     */
    completeMultipartUpload: (
      params: Params.CompleteMultipartUploadParams,
    ) => Promise<Response<Schema.CompleteMultipartUploadResponse>>;
    /**
     * **Experimental:** Create pre-signed URLs for multipart upload of large files
     *
     * `POST /multipart-uploads`
     */
    createMultipartUpload: (
      params: Params.CreateMultipartUploadParams,
    ) => Promise<Response<Schema.CreateMultipartUploadResponse>>;
    /**
     * **Experimental:** Retrieve the current status of a multipart upload
     *
     * `GET /multipart-uploads/{id}/status`
     */
    getMultipartUploadStatus: (
      params: Params.GetMultipartUploadStatusParams,
    ) => Promise<Response<Schema.GetMultipartUploadStatusResponse>>;
    /**
     * Get a presigned URL and related meta fields to upload a file. **The validity time for this URL is 60 minutes**.
     *
     * `GET /upload-url`
     */
    getUploadUrl: (
      params: Params.GetUploadUrlParams,
    ) => Promise<Response<Schema.GetUploadUrlResponse>>;
  };
  user: {
    /**
     * Find a user by email address
     *
     * `GET /users`
     */
    findUserByEmail: (
      params: Params.FindUserByEmailParams,
    ) => Promise<void>;
    /**
     * Get a user
     *
     * `GET /users/{id}`
     */
    getUser: (
      params: Params.GetUserParams,
    ) => Promise<Response<Schema.UserResponse>>;
    /**
     * Get list of users.
     *
     * `GET /userlist`
     */
    listUsers: (
      params: Params.ListUsersParams,
    ) => Promise<Response<Schema.UserListResponse>>;
  };
  workflow: {
    /**
     * **Experimental:** Get a workflow by ID.
     *
     * `GET /workflows/{workflow_id}`
     */
    getWorkflow: (
      params: Params.GetWorkflowParams,
    ) => Promise<Response<Schema.WorkflowResponse>>;
    /**
     * **Experimental:** Get a list of workflows.
     *
     * `GET /workflows`
     */
    listWorkflows: (
      params: Params.ListWorkflowsParams,
    ) => Promise<Response<Schema.ListWorkflowsResponse>>;
  };
  workRequest: {
    /**
     * **Experimental:** Create attachments for a work request.
     *
     * `POST /work-requests/{id}/attachments`
     */
    addAttachmentToWorkRequest: (
      params: Params.AddAttachmentToWorkRequestParams,
    ) => Promise<Response<Schema.AttachmentResponse>>;
    /**
     * **Experimental:** Post a comment on a work request.
     *
     * `POST /work-requests/{id}/comments`
     */
    addCommentToWorkRequest: (
      params: Params.AddCommentToWorkRequestParams,
    ) => Promise<Response<Schema.WorkRequestCommentResponse>>;
    /**
     * **Experimental:** Create a new campaign from a work request.
     *
     * `POST /work-requests/{id}/campaigns`
     */
    createCampaignFromWorkRequest: (
      params: Params.CreateCampaignFromWorkRequestParams,
    ) => Promise<Response<Schema.WorkRequestCampaignResponse>>;
    /**
     * **Experimental:** Create a new task from a work request.
     *
     * `POST /work-requests/{id}/tasks`
     */
    createTaskFromWorkRequest: (
      params: Params.CreateTaskFromWorkRequestParams,
    ) => Promise<Response<Schema.WorkRequestTaskResponse>>;
    /**
     * **Experimental:** Create a work request.
     *
     * `POST /work-requests`
     */
    createWorkRequest: (
      params: Params.CreateWorkRequestParams,
    ) => Promise<Response<Schema.WorkRequestResponse>>;
    /**
     * **Experimental:** Create creative assets for a work request.
     *
     * `POST /work-requests/{id}/creative-assets`
     */
    createWorkRequestCreativeAsset: (
      params: Params.CreateWorkRequestCreativeAssetParams,
    ) => Promise<Response<Schema.CreativeAssetResponse>>;
    /**
     * **Experimental:** Delete a work request attachment.
     *
     * `DELETE /work-requests/{work_request_id}/attachments/{attachment_id}`
     */
    deleteWorkRequestAttachment: (
      params: Params.DeleteWorkRequestAttachmentParams,
    ) => Promise<void>;
    /**
     * **Experimental:** Delete a work request creative asset.
     *
     * `DELETE /work-requests/{work_request_id}/creative-assets/{creative_asset_id}`
     */
    deleteWorkRequestCreativeAsset: (
      params: Params.DeleteWorkRequestCreativeAssetParams,
    ) => Promise<void>;
    /**
     * **Experimental:** Get a work request by ID.
     *
     * `GET /work-requests/{id}`
     */
    getWorkRequest: (
      params: Params.GetWorkRequestParams,
    ) => Promise<Response<Schema.WorkRequestResponse>>;
    /**
     * **Experimental:** Get a work request comment by ID.
     *
     * `GET /work-requests/{work_request_id}/comments/{comment_id}`
     */
    getWorkRequestComment: (
      params: Params.GetWorkRequestCommentParams,
    ) => Promise<Response<Schema.WorkRequestCommentResponse>>;
    /**
     * **Experimental:** Get a list of approved assets of a work request.
     *
     * `GET /work-requests/{id}/approved-assets`
     */
    listWorkRequestApprovedAssets: (
      params: Params.ListWorkRequestApprovedAssetsParams,
    ) => Promise<Response<Schema.ListWorkRequestApprovedAssetsResponse>>;
    /**
     * **Experimental:** Get list of comments for a work request.
     *
     * `GET /work-requests/{id}/comments`
     */
    listWorkRequestComments: (
      params: Params.ListWorkRequestCommentsParams,
    ) => Promise<Response<Schema.ListWorkRequestCommentsResponse>>;
    /**
     * **Experimental:** Get a list of work request related resources.
     *
     * `GET /work-requests/{id}/related-resources`
     */
    listWorkRequestRelatedResources: (
      params: Params.ListWorkRequestRelatedResourcesParams,
    ) => Promise<Response<Schema.ListWorkRequestRelatedResourcesResponse>>;
    /**
     * **Experimental:** Get a list of work requests.
     *
     * `GET /work-requests`
     */
    listWorkRequests: (
      params: Params.ListWorkRequestsParams,
    ) => Promise<Response<Schema.ListWorkRequestsResponse>>;
    /**
     * **Experimental:** Updates a work request.
     *
     * `PATCH /work-requests/{id}`
     */
    updateWorkRequest: (
      params: Params.UpdateWorkRequestParams,
    ) => Promise<void>;
    /**
     * **Experimental:** Updates a work request form field.
     *
     * `PUT /work-requests/{work_request_id}/form-fields/{form_field_identifier}`
     */
    updateWorkRequestFormField: (
      params: Params.UpdateWorkRequestFormFieldParams,
    ) => Promise<Response<Schema.WorkRequestRequestFormFieldUpdateResponse>>;
  };
}
