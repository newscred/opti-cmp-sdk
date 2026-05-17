// Auto-generated - DO NOT EDIT

export type paths = {
    "/articles/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /articles/{id}
         * @description Get an article
         */
        get: operations["getArticle"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/asset-lineages": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /asset-lineages
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get the list of asset lineage.
         */
        get: operations["listAssetLineages"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/asset-urls/{asset_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /asset-urls/{asset_id}
         * @description Find URL of an asset by ID
         */
        get: operations["getAssetUrl"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/assets": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /assets
         * @description Get the list of assets. Assets are sorted by `modified_at` in descending order.
         *     To get assets that are not inside any folder, pass `include_subfolder_assets=false` in the query param.
         */
        get: operations["listAssets"];
        put?: never;
        /**
         * POST /assets
         * @description Create a new asset. Supports only `images`, `videos`, and `raw files`. See [Upload assets](https://docs.developers.optimizely.com/content-marketing-platform/docs/upload-assets) to upload an asset.
         */
        post: operations["createAsset"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/assets/{asset_id}/fields": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET assets/{asset_id}/fields
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get the list of fields of an asset.
         */
        get: operations["listAssetFields"];
        /**
         * PUT /assets/{asset_id}/fields
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Replace existing fields of an asset.
         */
        put: operations["updateAssetFields"];
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/assets/{asset_id}/fields/{field_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        /**
         * PUT /assets/{asset_id}/fields/{field_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Update the field value of an asset.
         */
        put: operations["updateAssetField"];
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/assets/{asset_id}/lineages": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /assets/{asset_id}/lineages
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Add a new external asset lineage.
         */
        post: operations["createAssetLineage"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/assets/{asset_id}/lineages/{lineage_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        /**
         * DELETE /assets/{asset_id}/lineages/{lineage_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Delete an asset lineage.
         */
        delete: operations["deleteAssetLineage"];
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/assets/{asset_id}/permissions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /assets/{asset_id}/permissions
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ List of entities that have permission to access the asset
         */
        get: operations["listAssetPermissions"];
        put?: never;
        /**
         * POST /assets/{asset_id}/permissions
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Grant asset access to users or teams
         */
        post: operations["addAssetPermissions"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/assets/{asset_id}/permissions/{accessor_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        /**
         * DELETE /asstes/{asset_id}/permissions/{accessor_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Remove accessor's access from an asset
         */
        delete: operations["removeAssetPermission"];
        options?: never;
        head?: never;
        /**
         * PATCH /assets/{asset_id}/permissions/{accessor_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Update accessor's access level and ownership of asset
         */
        patch: operations["updateAssetPermission"];
        trace?: never;
    };
    "/assets/{asset_id}/renditions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /assets/{asset_id}/renditions
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get the renditions of an asset given its `id`.
         */
        get: operations["listAssetRenditions"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/assets/{asset_id}/versions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /assets/{asset_id}/versions
         * @description Add a new version to a library asset. Supports adding versions to only `images`, `videos`, and `raw files` type assets. See [Upload assets](https://docs.developers.optimizely.com/content-marketing-platform/docs/upload-assets) to upload a version to a library asset.
         */
        post: operations["createAssetVersion"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/brand-compliance/categories": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /brand-compliance/categories
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a list of brand compliance categories.
         */
        get: operations["listBrandComplianceCategories"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/campaigns": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /campaigns
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a list of campaigns.
         */
        get: operations["listCampaigns"];
        put?: never;
        /**
         * POST /campaigns
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Create a campaign
         */
        post: operations["createCampaign"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/campaigns/{campaign_id}/fields/{field_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        /**
         * PUT /campaigns/{campaign_id}/fields/{field_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Update the field value of a campaign.
         */
        put: operations["updateCampaignField"];
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/campaigns/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /campaigns/{id}
         * @description Get a campaign
         */
        get: operations["getCampaign"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        /**
         * PATCH /campaigns/{id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Update a campaign
         */
        patch: operations["updateCampaign"];
        trace?: never;
    };
    "/campaigns/{id}/attachments": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /campaigns/{id}/attachments
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Add an attachment to the campaign.
         */
        post: operations["addAttachmentToCampaign"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/campaigns/{id}/brief": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /campaigns/{id}/brief
         * @description Get brief of the campaign
         */
        get: operations["getCampaignBrief"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/campaigns/{id}/comments": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /campaigns/{id}/comments
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Post a comment on a campaign.
         */
        post: operations["addCommentToCampaign"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/campaigns/{id}/fields": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET campaigns/{id}/fields
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get the list of fields of a campaign.
         */
        get: operations["listCampaignFields"];
        put?: never;
        /**
         * POST campaigns/{id}/fields
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Add a field to a campaign.
         */
        post: operations["addFieldToCampaign"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/events": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /events
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a list of events.
         */
        get: operations["listEvents"];
        put?: never;
        /**
         * POST /events
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Create an event.
         */
        post: operations["createEvent"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/events/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /events/{id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get an event.
         */
        get: operations["getEvent"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        /**
         * PATCH /events/{id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Update an event.
         */
        patch: operations["updateEvent"];
        trace?: never;
    };
    "/events/{id}/fields": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET events/{id}/fields
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get the list of fields of an event.
         */
        get: operations["listEventFields"];
        /**
         * PUT /events/{id}/fields
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Replace existing fields of an event.
         */
        put: operations["updateEventFields"];
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/fields": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /fields
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get the list of fields of an Organization.
         */
        get: operations["listFields"];
        put?: never;
        /**
         * POST /fields
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Add a new field to an Organization.
         */
        post: operations["createField"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/fields/{field_id}/choices/{choice_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        /**
         * DELETE /fields/{field_id}/choices/{choice_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Delete a choice from a field.
         */
        delete: operations["deleteFieldChoice"];
        options?: never;
        head?: never;
        /**
         * PATCH /fields/{field_id}/choices/{choice_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Update a choice of a field.
         */
        patch: operations["updateFieldChoice"];
        trace?: never;
    };
    "/fields/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        /**
         * PATCH /fields/{id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Update a field in an Organization.
         */
        patch: operations["updateField"];
        trace?: never;
    };
    "/fields/{id}/choices": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /fields/{id}/choices
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Create choices in a field.
         */
        post: operations["createFieldChoices"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/file-urls": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /file-urls
         * @description Generates download URLs of files given file guid.
         */
        post: operations["createFileUrls"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/folders": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /folders
         * @description Get the list of folders sorted by `modified_at` in descending order.
         */
        get: operations["listFolders"];
        put?: never;
        /**
         * POST /folders
         * @description Create a folder. Use the `parent_folder_id` field to create a nested folder.
         */
        post: operations["createFolder"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/folders/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /folders/{id}
         * @description Get a folder
         */
        get: operations["getFolder"];
        put?: never;
        post?: never;
        /**
         * DELETE /folders/{id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Delete a folder.
         */
        delete: operations["deleteFolder"];
        options?: never;
        head?: never;
        /**
         * PATCH /folders/{id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Update a folder's name and/or parent.
         */
        patch: operations["updateFolder"];
        trace?: never;
    };
    "/folders/{id}/permissions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /folders/{id}/permissions
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ List of entities that have permission to access the folder
         */
        get: operations["listFolderPermissions"];
        put?: never;
        /**
         * POST /folders/{id}/permissions
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Grant folder access to users or teams
         */
        post: operations["addFolderPermissions"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/folders/{id}/permissions/{accessor_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        /**
         * DELETE /folders/{id}/permissions/{accessor_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Remove accessor's access from a folder
         */
        delete: operations["removeFolderPermission"];
        options?: never;
        head?: never;
        /**
         * PATCH /folders/{id}/permissions/{accessor_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Update accessor's access level and ownership of folder
         */
        patch: operations["updateFolderPermission"];
        trace?: never;
    };
    "/images/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /images/{id}
         * @description Get an image
         */
        get: operations["getImage"];
        put?: never;
        post?: never;
        /**
         * DELETE /images/{id}
         * @description Delete an image
         */
        delete: operations["deleteImage"];
        options?: never;
        head?: never;
        /**
         * PATCH /images/{id}
         * @description Updates an image
         */
        patch: operations["updateImage"];
        trace?: never;
    };
    "/label-groups": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /label-groups
         * @description Get the list of label groups. Label groups are sorted by `name` in ascending order.
         */
        get: operations["listLabelGroups"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/milestones": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /milestones
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a list of milestones.
         */
        get: operations["listMilestones"];
        put?: never;
        /**
         * POST /milestones
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Creates a milestone. Defaults to the organization's campaign if campaign_id is omitted. Tasks are validated against the selected campaign and must belong to it.
         */
        post: operations["createMilestone"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/milestones/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /milestones/{id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a milestone by ID.
         */
        get: operations["getMilestone"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        /**
         * PATCH /milestones/{id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Updates a milestone. All fields are optional. Only provided fields will be updated. Note: If tasks array is empty, it will remove all task associations for that milestone.
         */
        patch: operations["updateMilestone"];
        trace?: never;
    };
    "/multipart-uploads": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /v3/multipart-uploads
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Create pre-signed URLs for multipart upload of large files
         */
        post: operations["createMultipartUpload"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/multipart-uploads/{id}/complete": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /v3/multipart-uploads/{id}/complete
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Initiate completion of a multipart upload after all parts have been uploaded
         */
        post: operations["completeMultipartUpload"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/multipart-uploads/{id}/status": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /v3/multipart-uploads/{id}/status
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Retrieve the current status of a multipart upload
         */
        get: operations["getMultipartUploadStatus"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/publishing-events/{publishing_event_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /v3/publishing-events/{publishing_event_id}
         * @description Get the publishing event by ID
         */
        get: operations["getPublishingEvent"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/publishing-events/{publishing_event_id}/assets/{asset_id}/publishing-metadata/{publishing_metadata_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET v3/publishing-events/{publishing_event_id}/assets/{asset_id}/publishing-metadata/{publishing_metadata_id}
         * @description Get publishing metadata.
         */
        get: operations["getPublishingEventAssetMetadata"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/publishing-events/{publishing_event_id}/publishing-metadata": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /v3/publishing-events/{publishing_event_id}/publishing-metadata
         * @description Get the list of publishing metadata for each asset of a publishing event.
         */
        get: operations["listPublishingEventMetadata"];
        put?: never;
        /**
         * POST /v3/publishing-events/{publishing_event_id}/publishing-metadata
         * @description Create asset metadata in bulk.
         */
        post: operations["bulkCreatePublishingEventMetadata"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/raw-files/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /raw-files/{id}
         * @description Get a raw file
         */
        get: operations["getRawFile"];
        put?: never;
        post?: never;
        /**
         * DELETE /raw-files/{id}
         * @description Delete a raw file
         */
        delete: operations["deleteRawFile"];
        options?: never;
        head?: never;
        /**
         * PATCH /raw-files/{id}
         * @description Updates a raw file
         */
        patch: operations["updateRawFile"];
        trace?: never;
    };
    "/rendition-configs/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /rendition-configs/{id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a rendition configuration of an organization given its `id`.
         */
        get: operations["getRenditionConfig"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/renditions/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /renditions/{id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a rendition given its `id`.
         */
        get: operations["getRendition"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/settings": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /settings
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get settings.
         */
        get: operations["getSettings"];
        put?: never;
        /**
         * POST /settings
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Create or update settings.
         *     ## Import behavior:
         *     - Label:
         *         - If a label-group with the same name exists, the label-group is updated.
         *           - Labels are merged.
         *           - If a label with the same name exists, it is overwritten.
         *         - If no label-group with the same name exists, a new label-group is created.
         *     - Custom Fields:
         *         - If a custom field with the same name exists, the custom field is updated and its choices are merged.
         *         - If no custom field with the same name exists, a new custom field will be created.
         *     - Workflows:
         *         - If the query param overwrite_workflows=true the existing workflows are overwritten.
         *           Otherwise, a new workflow will be created where the prefix `Copy of` is added to the name of the workflow.
         *     - Webhooks:
         *         - If a webhook with the same name exists, the webhook is updated and its events are merged.
         *         - If multiple webhooks with the same name exist, a new webhook is created.
         *         - If no webhook with the same name exists, a new webhook is created.
         *     - Apps:
         *         - If an app with the same name exists, the app is updated and its authorization callback URLs are merged.
         *         - If multiple apps with the same name exist, a new app is created.
         *         - If no app with the same name exists, a new app is created.
         *     - Templates:
         *         - If a template with the same name exists, the template is updated.
         *         - If multiple templates with the same name exist, a new template is created.
         *         - If no template with the same name exists, a new template is created.
         *     - Routing Rules:
         *         - If one single routing rule with the same name exists, the routing rule is updated.
         *         - If multiple routing rules with the same name exist, a new routing rule is created.
         *         - If no routing rule with the same name exists, a new routing rule is created.
         */
        post: operations["updateSettings"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/structured-content/content-types": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /structured-content/content-types
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get content types.
         */
        get: operations["listSCContentTypes"];
        put?: never;
        /**
         * POST /structured-content/content-types
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Create content type.
         */
        post: operations["createSCContentType"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/structured-content/content-types/{content_type_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /structured-content/content-types/{content_type_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get content type.
         */
        get: operations["getSCContentType"];
        put?: never;
        /**
         * POST /structured-content/content-types/{content_type_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Update content type.
         */
        post: operations["updateSCContentType"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/structured-content/content-types/{content_type_id}/managed-migrations": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /structured-content/content-types/{content_type_id}/managed-migrations
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Retrieves a list of managed migration jobs for a specific content type.
         */
        get: operations["listSCContentTypeManagedMigrations"];
        put?: never;
        /**
         * POST /structured-content/content-types/{content_type_id}/managed-migrations
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Create a new managed migration job.
         */
        post: operations["createSCContentTypeManagedMigration"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/structured-content/content-types/{content_type_id}/managed-migrations/validate": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /structured-content/content-types/{content_type_id}/managed-migrations/validate
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Check managed migration possible or not.
         */
        post: operations["validateSCContentTypeManagedMigration"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/structured-content/content-types/{content_type_id}/managed-migrations/{job_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get details of a specific managed migration job.
         */
        get: operations["getSCContentTypeManagedMigration"];
        put?: never;
        post?: never;
        /**
         * DELETE /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Delete a managed migration job with not_started status.
         */
        delete: operations["deleteSCContentTypeManagedMigration"];
        options?: never;
        head?: never;
        /**
         * PATCH /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Update a managed migration job.
         */
        patch: operations["updateSCContentTypeManagedMigration"];
        trace?: never;
    };
    "/structured-content/content-types/{content_type_id}/managed-migrations/{job_id}/start": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /structured-content/content-types/{content_type_id}/managed-migrations/{job_id}/start
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Start managed migration job.
         */
        post: operations["startSCContentTypeManagedMigration"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/structured-content/content-types/{content_type_id}/versions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /structured-content/content-types/{content_type_id}/versions
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get content type versions.
         */
        get: operations["listSCContentTypeVersions"];
        put?: never;
        /**
         * POST /structured-content/content-types/{content_type_id}/versions
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Add Content Type Version.
         */
        post: operations["createSCContentTypeVersion"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/structured-content/content-types/{content_type_id}/versions/{version_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /structured-content/content-types/{content_type_id}/versions/{version_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get content type version.
         */
        get: operations["getSCContentTypeVersion"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/structured-content/contents/{content_id}/migration": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /structured-content/contents/{content_id}/migration
         * @description Migrate content to a specific content type version. @EXPERIMENTAL@BADGE@PLACEHOLDER@
         */
        post: operations["migrateSCContent"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/structured-content/contents/{content_id}/versions/{version_id}/previews/{preview_id}/acknowledge": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /structured-content/contents/{content_id}/versions/{version_id}/previews/{preview_id}/acknowledge
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Acknowledge content preview. Content preview can be acknowledged only once. So make sure you are acknowledging only the content previews targeted for your integration. Otherwise it will stall acknowledgment from other integrations.
         */
        post: operations["acknowledgeSCContentPreview"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/structured-content/contents/{content_id}/versions/{version_id}/previews/{preview_id}/complete": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /structured-content/contents/{content_id}/versions/{version_id}/previews/{preview_id}/complete
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Complete content preview.
         */
        post: operations["completeSCContentPreview"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/structured-contents": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /structured-contents
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Creates a structured content.
         */
        post: operations["createStructuredContent"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/structured-contents/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /structured-contents/{id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get the structured content.
         */
        get: operations["getStructuredContent"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        /**
         * PATCH /structured-contents/{id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Updates a structured content.
         */
        patch: operations["updateStructuredContent"];
        trace?: never;
    };
    "/tasks": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a list of tasks.
         */
        get: operations["listTasks"];
        put?: never;
        /**
         * POST /tasks
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Create a task
         */
        post: operations["createTask"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{id}
         * @description Get a task
         */
        get: operations["getTask"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        /**
         * PATCH /tasks/{id}
         * @description Update a task
         */
        patch: operations["updateTask"];
        trace?: never;
    };
    "/tasks/{id}/assets": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{id}/assets
         * @description Get the list of the assets or contents of a task
         */
        get: operations["listTaskAssets"];
        put?: never;
        /**
         * POST /tasks/{id}/assets
         * @description This API allows you to add an asset to a task. There currently two mechanisms to add an asset - either by uploading an asset first and then calling this API (for only `images`, `videos`, and `raw files`), or by adding an asset from the library. The 'type' parameter in the request body will determine the mechanism used to add the asset. See [Upload assets](https://docs.developers.optimizely.com/content-marketing-platform/docs/upload-assets) to upload an asset.
         */
        post: operations["addAssetToTask"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{id}/attachments": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{id}/attachments
         * @description Get the list of the attachments of a task
         */
        get: operations["listTaskAttachments"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{id}/brief": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{id}/brief
         * @description Get brief of the task
         */
        get: operations["getTaskBrief"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{id}/custom-fields": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET tasks/{id}/custom-fields
         * @description Get the list of custom fields added to a task.
         */
        get: operations["listTaskCustomFields"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/articles/{article_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/articles/{article_id}
         * @description Get an article associated with the task identified by `task_id`.
         */
        get: operations["getTaskArticle"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/assets/{asset_id}/comments": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/assets/{asset_id}/comments
         * @description List all the comments for the latest version of a task asset.
         */
        get: operations["listTaskAssetComments"];
        put?: never;
        /**
         * POST /tasks/{task_id}/assets/{asset_id}/comments
         * @description Add a comment to the latest version of a task asset. Supports adding comments to only **images**, **videos** and **raw files** type assets.
         */
        post: operations["addCommentToTaskAsset"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/assets/{asset_id}/drafts": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/assets/{asset_id}/drafts
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get the list of drafts on an asset of a task.
         */
        get: operations["listTaskAssetDrafts"];
        put?: never;
        /**
         * POST /tasks/{task_id}/assets/{asset_id}/drafts
         * @description Add a new draft to a task asset. You can add a draft to only `images`, `videos`, and `raw files` type task assets. See [Upload assets](https://docs.developers.optimizely.com/content-marketing-platform/docs/upload-assets) to upload an asset draft.
         */
        post: operations["addDraftToTaskAsset"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/assets/{asset_id}/drafts/{draft_id}/brand-compliance": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/assets/{asset_id}/drafts/{draft_id}/brand-compliance
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get the brand compliance details of a draft of an asset of a task.
         */
        get: operations["getTaskAssetDraftBrandCompliance"];
        /**
         * PUT /tasks/{task_id}/assets/{asset_id}/drafts/{draft_id}/brand-compliance
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Update the brand compliance details of a draft of an asset of a task.
         */
        put: operations["updateTaskAssetDraftBrandCompliance"];
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/assets/{asset_id}/fields": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/assets/{asset_id}/fields
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ List all the fields of a task asset.
         */
        get: operations["listTaskAssetFields"];
        /**
         * PUT tasks/{task_id}/assets/{asset_id}/fields
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Replace fields of an asset in a task.
         */
        put: operations["updateTaskAssetFields"];
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/comments": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /tasks/{task_id}/comments
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Post a comment on a task.
         */
        post: operations["addCommentToTask"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/custom-fields/{custom_field_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/custom-fields/{custom_field_id}
         * @description Get a custom field of a task
         */
        get: operations["getTaskCustomField"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        /**
         * PATCH /tasks/{task_id}/custom-fields/{custom_field_id}
         * @description Update a custom field of a task.
         *     You can update the following types:  `text_field`, `multi_line_text_field`, `checkboxes`, `dropdown`, `multi_select_dropdown`, `multiple_choice`, `date_field`, `rich_text_field`
         */
        patch: operations["updateTaskCustomField"];
        trace?: never;
    };
    "/tasks/{task_id}/custom-fields/{custom_field_id}/choices": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/custom-fields/{custom_field_id}/choices
         * @description Get the list of the choices of a custom field in a task.
         */
        get: operations["listTaskCustomFieldChoices"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/fields": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET tasks/{task_id}/fields
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get the list of fields of a task.
         */
        get: operations["listTaskFields"];
        put?: never;
        /**
         * POST tasks/{task_id}/fields
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Add a field to a task.
         */
        post: operations["addFieldToTask"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/fields/{field_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        /**
         * PUT tasks/{taks_id}/fields/{field_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Update a field of a task.
         */
        put: operations["updateTaskField"];
        post?: never;
        /**
         * DELETE tasks/{task_id}/fields/{field_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Remove a field of a task.
         */
        delete: operations["removeTaskField"];
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/images/{image_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/images/{image_id}
         * @description Get an image associated with the task identified by `task_id`.
         */
        get: operations["getTaskImage"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/raw-files/{raw_file_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/raw-files/{raw_file_id}
         * @description Get a raw file associated with the task identified by `task_id`.
         */
        get: operations["getTaskRawFile"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/steps/{step_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        /**
         * PATCH /tasks/{task_id}/steps/{step_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Update a step in a task
         */
        patch: operations["updateTaskStep"];
        trace?: never;
    };
    "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}
         * @description Get the substep of a task
         */
        get: operations["getTaskSubStep"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        /**
         * PATCH /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}
         * @description Update a substep of a step in a task
         */
        patch: operations["updateTaskSubStep"];
        trace?: never;
    };
    "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments
         * @description Get the list of the comments of a task substep
         */
        get: operations["listTaskSubStepComments"];
        put?: never;
        /**
         * POST /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments
         * @description Create a task substep comment
         */
        post: operations["addCommentToTaskSubStep"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}
         * @description Get a task substep comment
         */
        get: operations["getTaskSubStepComment"];
        put?: never;
        post?: never;
        /**
         * DELETE /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}
         * @description Delete a task substep comment
         */
        delete: operations["deleteTaskSubStepComment"];
        options?: never;
        head?: never;
        /**
         * PATCH /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}
         * @description Update a task substep comment
         */
        patch: operations["updateTaskSubStepComment"];
        trace?: never;
    };
    "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/external-work": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/external-work
         * @description Get the external work information of a task external substep
         */
        get: operations["getTaskSubStepExternalWork"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        /**
         * PATCH /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/external-work
         * @description Update the external work information of a task external substep
         */
        patch: operations["updateTaskSubStepExternalWork"];
        trace?: never;
    };
    "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/fields": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/fields
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get the list of fields of a substep of a task.
         */
        get: operations["listTaskSubStepFields"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/structured-contents": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /tasks/{task_id}/structured-contents
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Add a new structured content to task
         */
        post: operations["addStructuredContentToTask"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/structured-contents/{content_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/structured-contents/{content_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a task structured content by its guid.
         */
        get: operations["getTaskStructuredContent"];
        put?: never;
        post?: never;
        /**
         * DELETE /tasks/{task_id}/structured-contents/{content_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Delete a structured content from task
         */
        delete: operations["deleteTaskStructuredContent"];
        options?: never;
        head?: never;
        /**
         * PATCH /tasks/{task_id}/structured-contents/{content_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Update structured content using a task.
         */
        patch: operations["updateTaskStructuredContent"];
        trace?: never;
    };
    "/tasks/{task_id}/structured-contents/{content_id}/drafts": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /tasks/{task_id}/structured-contents/{content_id}/drafts
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Create structured content draft using a task.
         */
        post: operations["createTaskStructuredContentDraft"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/urls": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /tasks/{task_id}/urls
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Adds a URL to a task.
         */
        post: operations["addUrlToTask"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tasks/{task_id}/videos/{video_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /tasks/{task_id}/videos/{video_id}
         * @description Get a video associated with the task identified by task_id
         */
        get: operations["getTaskVideo"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/teams": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /teams
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a list of teams.
         */
        get: operations["listTeams"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/teams/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /teams/{id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a team.
         */
        get: operations["getTeam"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/templates": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /templates
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a list of templates.
         */
        get: operations["listTemplates"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/templates/{template_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /templates/{template_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a template by ID.
         */
        get: operations["getTemplate"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/upload-url": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /upload-url
         * @description Get a presigned URL and related meta fields to upload a file. **The validity time for this URL is 60 minutes**.
         */
        get: operations["getUploadUrl"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/userlist": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /userlist
         * @description Get list of users.
         */
        get: operations["listUsers"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/users": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /users
         * @description Find a user by email address
         */
        get: operations["findUserByEmail"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/users/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /users/{id}
         * @description Get a user
         */
        get: operations["getUser"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/videos/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /videos/{id}
         * @description Get a video
         */
        get: operations["getVideo"];
        put?: never;
        post?: never;
        /**
         * DELETE /videos/{id}
         * @description Delete a video
         */
        delete: operations["deleteVideo"];
        options?: never;
        head?: never;
        /**
         * PATCH /videos/{id}
         * @description Updates a video
         */
        patch: operations["updateVideo"];
        trace?: never;
    };
    "/work-requests": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /work-requests
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a list of work requests.
         */
        get: operations["listWorkRequests"];
        put?: never;
        /**
         * POST /work-requests
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Create a work request.
         */
        post: operations["createWorkRequest"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/work-requests/{id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /work-requests/{id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a work request by ID.
         */
        get: operations["getWorkRequest"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        /**
         * PATCH /work-requests/{id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Updates a work request.
         */
        patch: operations["updateWorkRequest"];
        trace?: never;
    };
    "/work-requests/{id}/approved-assets": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /work-requests/{id}/approved-assets
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a list of approved assets of a work request.
         */
        get: operations["listWorkRequestApprovedAssets"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/work-requests/{id}/attachments": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /work-requests/{id}/attachments
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Create attachments for a work request.
         */
        post: operations["addAttachmentToWorkRequest"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/work-requests/{id}/campaigns": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /work-requests/{id}/campaigns
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Create a new campaign from a work request.
         */
        post: operations["createCampaignFromWorkRequest"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/work-requests/{id}/comments": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /work-requests/{id}/comments
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get list of comments for a work request.
         */
        get: operations["listWorkRequestComments"];
        put?: never;
        /**
         * POST /work-requests/{id}/comments
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Post a comment on a work request.
         */
        post: operations["addCommentToWorkRequest"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/work-requests/{id}/creative-assets": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /work-requests/{id}/creative-assets
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Create creative assets for a work request.
         */
        post: operations["createWorkRequestCreativeAsset"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/work-requests/{id}/related-resources": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /work-requests/{id}/related-resources
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a list of work request related resources.
         */
        get: operations["listWorkRequestRelatedResources"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/work-requests/{id}/tasks": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * POST /work-requests/{id}/tasks
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Create a new task from a work request.
         */
        post: operations["createTaskFromWorkRequest"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/work-requests/{work_request_id}/attachments/{attachment_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        /**
         * DELETE /work-requests/{work_request_id}/attachments/{attachment_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Delete a work request attachment.
         */
        delete: operations["deleteWorkRequestAttachment"];
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/work-requests/{work_request_id}/comments/{comment_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /work-requests/{work_request_id}/comments/{comment_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a work request comment by ID.
         */
        get: operations["getWorkRequestComment"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/work-requests/{work_request_id}/creative-assets/{creative_asset_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        /**
         * DELETE /work-requests/{work_request_id}/creative-assets/{creative_asset_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Delete a work request creative asset.
         */
        delete: operations["deleteWorkRequestCreativeAsset"];
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/work-requests/{work_request_id}/form-fields/{form_field_identifier}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        /**
         * PUT /work-requests/{work_request_id}/form-fields/{form_field_identifier}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Updates a work request form field.
         */
        put: operations["updateWorkRequestFormField"];
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/workflows": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /workflows
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a list of workflows.
         */
        get: operations["listWorkflows"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/workflows/{workflow_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * GET /workflows/{workflow_id}
         * @description @EXPERIMENTAL@BADGE@PLACEHOLDER@ Get a workflow by ID.
         */
        get: operations["getWorkflow"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
};
export type webhooks = Record<string, never>;
export type components = {
    schemas: {
        AddUrlToTaskRequest: {
            /**
             * @description Title of the URL.
             * @example Google
             */
            title?: string;
            /**
             * @description The URL to be added to the task.
             * @example https://www.google.com
             */
            url: string;
        };
        AllowedContentTypeItem: {
            /** @description Name of the item */
            name: string;
            /** @description URL of the item */
            url: string;
        };
        /** @description Author of the article */
        ArticleAuthorResponse: {
            /**
             * @description Name of Author
             * @example John Doe
             */
            name: string | null;
        };
        /** @description Content of the asset */
        AssetContent: {
            /**
             * @description Type of the content.
             *      - article  – `html_body`
             *      - image, raw_file, video  – `url`
             *      - structured_content  – `api_url`
             * @example url
             * @enum {string}
             */
            type: "url" | "api_url" | "html_body";
            /**
             * @description Content of the asset.  - article – the html body - image, raw_file, video – the download URL - structured_content – api url
             * @example http://images.cmp.optimizely.com/Zz0xODQ3NDU3Y2Y2YmYzOTlmNjkyOTgyZDY3Y2I3YWM2OA==S
             */
            value: string;
        };
        AssetFieldListResponseItem: components["schemas"]["AssetFieldTypeCommon"] | components["schemas"]["AssetFieldTypeLabel"] | components["schemas"]["AssetFieldTypeDropdown"] | components["schemas"]["AssetFieldTypeRadio"] | components["schemas"]["AssetFieldTypeCheckbox"] | components["schemas"]["AssetFieldTypeSimpleNumber"] | components["schemas"]["AssetFieldTypePercentageNumber"] | components["schemas"]["AssetFieldTypeCurrencyNumber"];
        AssetFieldTypeCheckbox: components["schemas"]["BaseAssetField"] & {
            /** @description Choices of the checkbox */
            choices: {
                /**
                 * @description Identifier of the choice
                 * @example 6ceee2f4fa3411ecb37802420ac8001b
                 */
                id: string;
                /**
                 * @description Name of the choice
                 * @example Checkbox 1 Choice 1
                 */
                name: string;
            }[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "checkbox";
        };
        AssetFieldTypeCommon: components["schemas"]["BaseAssetField"] & Record<string, never> & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "date" | "image" | "rich_text" | "text" | "text_area" | "video";
        };
        AssetFieldTypeCurrencyNumber: components["schemas"]["BaseAssetField"] & {
            /**
             * @description Currency code of the numerical field
             * @example USD
             */
            currency_code: string;
            /**
             * @description Decimal place of the numerical field
             * @example 2
             */
            decimal_places: number | null;
            /**
             * @description Whether the numerical field has a thousand separator
             * @example true
             */
            has_thousand_separator: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "currency_number";
        };
        AssetFieldTypeDropdown: components["schemas"]["BaseAssetField"] & {
            /** @description Choices of the dropdown */
            choices: {
                /**
                 * @description Identifier of the choice
                 * @example 6ceee2f4fa3411ecb37802420ac8001b
                 */
                id: string;
                /**
                 * @description Name of the choice
                 * @example Dropdown 1 Choice 1
                 */
                name: string;
            }[];
            /** @description Allow users to select multiple values from the dropdown */
            is_multi_select: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "dropdown";
        };
        AssetFieldTypeLabel: components["schemas"]["BaseAssetField"] & {
            /** @description Choices of the label */
            choices: {
                /**
                 * @description Identifier of the choice
                 * @example 6ceee2f4fa3411ecb37802420ac8001b
                 */
                id: string;
                /**
                 * @description Name of the choice
                 * @example Label 1 Choice 1
                 */
                name: string;
            }[];
            /** @description Select multiple values from the label */
            is_multi_select: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "label";
        };
        AssetFieldTypePercentageNumber: components["schemas"]["BaseAssetField"] & {
            /**
             * @description Decimal place of the numerical field
             * @example 2
             */
            decimal_places: number | null;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "percentage_number";
        };
        AssetFieldTypeRadio: components["schemas"]["BaseAssetField"] & {
            /** @description Choices of the radio button */
            choices: {
                /**
                 * @description Identifier of the choice
                 * @example 6ceee2f4fa3411ecb37802420ac8001b
                 */
                id: string;
                /**
                 * @description Name of the choice
                 * @example Radio 1 Choice 1
                 */
                name: string;
            }[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "radio_button";
        };
        AssetFieldTypeSimpleNumber: components["schemas"]["BaseAssetField"] & {
            /**
             * @description Decimal place of the numerical field
             * @example 2
             */
            decimal_places: number | null;
            /**
             * @description Whether the numerical field has a thousand separator
             * @example true
             */
            has_thousand_separator: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "simple_number";
        };
        /**
         * @description Type of the field
         * @example image
         * @enum {string}
         */
        AssetFieldTypes: "checkbox" | "currency_number" | "date" | "dropdown" | "label" | "radio_button" | "text_area" | "percentage_number" | "rich_text" | "simple_number" | "text";
        AssetFieldUpdateRequest: components["schemas"]["StringTypeObjectFieldUpdatePayload"] | components["schemas"]["MultiChoiceTypeObjectFieldUpdatePayload"] | components["schemas"]["DropdownTypeObjectFieldUpdatePayload"] | components["schemas"]["RadioButtonTypeObjectFieldUpdatePayload"] | components["schemas"]["NumberTypeObjectFieldUpdatePayload"] | components["schemas"]["DateTypeObjectFieldUpdatePayload"] | components["schemas"]["AssetTypeObjectFieldUpdatePayload"];
        AssetFieldUpdateResponse: {
            /**
             * @description Unique identifier for the field
             * @example 5a7f910511b0a72230ce6631
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the asset
                 * @example https://api.cmp.optimizely.com/v3/assets/5f857f30e1c4a2038d6179e9
                 */
                asset: string;
                /**
                 * @description URL of the asset fields
                 * @example https://api.cmp.optimizely.com/v3/assets/5f857f30e1c4a2038d6179e9/fields
                 */
                asset_fields: string;
            };
            /**
             * @description Name of the field
             * @example Image Field
             */
            name: string;
            type: components["schemas"]["AssetFieldTypes"];
            /** @description List of values selected or from user input */
            values: string[];
        };
        AssetFieldsUpdateRequest: (components["schemas"]["StringTypeObjectField"] | components["schemas"]["MultiChoiceTypeObjectField"] | components["schemas"]["DropdownTypeObjectField"] | components["schemas"]["RadioButtonTypeObjectField"] | components["schemas"]["NumberTypeObjectField"] | components["schemas"]["DateTypeObjectField"] | components["schemas"]["AssetTypeObjectField"])[];
        AssetLineageCreateRequest: {
            /**
             * @description Website favicon url
             * @example https://domain.xyz/path/to/favicon
             */
            icon_url?: string;
            /**
             * @description Name of the source
             * @example Some website name
             */
            name: string;
            /**
             * @description Unique identifier of a rendition of an existing version of the asset
             * @example 1d7491a891b00a721e0419630a7f8a4b
             */
            rendition_id?: string;
            /**
             * @description URI of the resource where the asset is used. If this is an external url, make sure the url starts with http:// or https://.
             * @example https://domain.xyz/path/to/resource
             */
            uri: string;
        };
        AssetLineageResponse: {
            /**
             * @description Unique identifier of the asset
             * @example 5d7f910551b00a722e0418830cee6632
             */
            asset_id: string;
            /**
             * Format: date-time
             * @description Date and time on which the asset lineage was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Icon url of the asset lineage
             * @example https://somewebsite.com/abc.png
             */
            icon_url: string | null;
            /**
             * @description Unique identifier of the asset lineage
             * @example 5d7f910551b00a722e0418830cee6631
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the asset
                 * @example https://api.cmp.optimizely.com/v3/articles/5d7f910551bw0a722e0418830cee6632
                 */
                asset: string;
                /**
                 * @description URL of the asset lineage
                 * @example https://api.cmp.optimizely.com/v3/assets/5d7f910551bw0a722e0418830cee6632/lineages/5d7f910551bw0a722e0418830cee6631
                 */
                self: string;
            };
            /**
             * @description Name of the source
             * @example Sample Name
             */
            name: string;
            /**
             * @description Unique identifier of a rendition of the asset
             * @example 1d7491a891b00a721e0419630a7f8a4b
             */
            rendition_id: string | null;
            /**
             * @description Unique identifier of the location where the asset is being used
             * @example https://somewebsite.com
             */
            uri: string;
            /**
             * @description Indicates where the asset is being used
             * @example external
             * @enum {string}
             */
            used_in: "external";
            /**
             * @description Unique identifier of the asset version
             * @example 5d7f910551b00a722e0418830cee6633
             */
            version_id: string;
        };
        AssetPermissionBulkCreateRequest: {
            /** @description List of permissions to be granted to asset */
            permissions: {
                /**
                 * @description level of access the accessor has to the asset
                 * @example view
                 * @enum {string}
                 */
                access_type: "view" | "edit";
                /**
                 * @description Unique identifier of accessor
                 * @example 5d7f910551b00a722e0418830cee6632
                 */
                id: string;
                /**
                 * @description Indicates if the accessor is the owner of the asset. A team cannot be an owner.
                 * @default false
                 * @example false
                 */
                is_owner: boolean;
            }[];
            /**
             * @description Type of the accessor
             * @example user
             * @enum {string}
             */
            type: "user" | "team";
        };
        AssetPermissionListResponseItem: {
            /** @description List of Permissions */
            data: components["schemas"]["BasePermissionsResponseSchema"][];
            pagination: components["schemas"]["Pagination"] & {
                /** @example https://api.cmp.optimizely.com/v3/assets/c00bfb60703f11ef805602420ac8001b/permissions?offset=10&page_size=10 */
                next?: string | null;
            };
        };
        AssetPermissionUpdateRequest: {
            /**
             * @description Set the level of access the accessor has to the asset
             * @example view
             * @enum {string}
             */
            access_type?: "view" | "edit";
            /**
             * @description Used to set an accessor as owner. This field can only be set to true as an owner cannot be removed without another accessor being set as the new owner. A team cannot be an owner.
             * @example true
             */
            is_owner?: boolean;
        };
        AssetResponse: {
            content: components["schemas"]["AssetContent"];
            /**
             * Format: date-time
             * @description Date and time on which the asset was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description File extension of the asset
             * @example png
             */
            file_extension: string | null;
            /**
             * @description Location of the folder containing the asset in the library
             * @example /all assets/important assets
             */
            file_location: string;
            /**
             * @description ID of the folder containing the asset in the library
             * @example 6bb8db20a5b611ebae319b7c541b1a5a
             */
            folder_id: string | null;
            /**
             * @description Unique identifier for the asset
             * @example 5d7f910551b00a722e0418830cee6631
             */
            id: string;
            /**
             * @description Whether the asset is archived or not
             * @example false
             */
            is_archived: boolean;
            /** @description Labels associated with the asset */
            labels: components["schemas"]["ResourceLabelResponse"][];
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the asset (only GET is supported)
                 * @example https://api.cmp.optimizely.com/v3/images/5d7f910551bw0a722e0418830cee6631
                 */
                self: string;
            };
            /**
             * @description MIME type of the asset
             * @example image/png
             */
            mime_type: string;
            /**
             * Format: date-time
             * @description Date and time of the most recent modification of the asset, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            modified_at: string;
            /**
             * @description ID of the asset's owner organization
             * @example 5108c3a9becac35915111191
             */
            owner_organization_id: string;
            /**
             * @description URL of the asset's thumbnail
             * @example https://images1.cmp.optimizely.com/Pz1kYmI9Z2FkYTJtZWI5VGI1WZq4MTZkNTdmGHjM5OGRmYq==
             */
            thumbnail_url: string | null;
            /**
             * @description Title of the asset
             * @example sample_image.png
             */
            title: string;
            /**
             * @description Type of the asset
             * @example image
             * @enum {string}
             */
            type: "article" | "image" | "video" | "raw_file" | "structured_content";
        };
        AssetTypeObjectField: components["schemas"]["BaseObjectField"] & components["schemas"]["AssetTypeObjectFieldUpdatePayload"] & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "image" | "video";
        };
        AssetTypeObjectFieldUpdatePayload: components["schemas"]["BaseObjectFieldUpdatePayload"] & {
            /** @description Payload for fields with attachments */
            values?: {
                /**
                 * @description Unique identifier of the file upload session. This is the `upload_meta_fields.key` field retrieved from the `/v3/upload-url` endpoint.
                 * @example ce8995aea58b11ea8cd90242ac120005
                 */
                key: string;
                /**
                 * @description Title of the image or video
                 * @example Sample.jpeg
                 */
                title: string;
            }[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "image" | "video";
        };
        AssetVersionCreateRequest: {
            /** @description Unique identifier of the file upload session. This is the `upload_meta_fields.key` field retrieved from the `/v3/upload-url` endpoint. */
            key: string;
            /**
             * @description Title of the asset
             * @example sample_image.png
             */
            title: string;
        };
        AttachmentRequest: {
            /**
             * @description Unique identifier of the file upload session. This is the `upload_meta_fields.key` field retrieved from the `/v3/upload-url` endpoint.
             * @example ce8995aea58b11ea8cd90242ac120005
             */
            key: string;
            /**
             * @description Filename of the attachment
             * @example sample_image.png
             */
            name: string;
        };
        AttachmentResponse: {
            /**
             * @description Unique identifier for the attachment
             * @example 5a7f910511b0a72230ce6631
             */
            id: string;
            /**
             * @description Filename of the attachment
             * @example sample_image.png
             */
            name: string;
            /**
             * @description Download URL of the attachment
             * @example https://files.cmp.optimizely.com/download/2115bfe4450c11ebaae8000c291b51d4
             */
            url: string;
        };
        BaseAssetField: {
            /**
             * @description Identifier of the field
             * @example 64be245ec0d79e79fdf1ad84
             */
            id: string;
            /**
             * @description Name of the field
             * @example Simple Text
             */
            name: string;
            /**
             * @description Type of the field
             * @example text
             */
            type: string;
            /** @description List of selected values or from user input */
            values: string[];
        };
        BaseAssetRenditionResponse: {
            /**
             * Format: date-time
             * @description Date and time on which the rendition was created, in ISO 8601 UTC format
             * @example 2022-09-27T08:18:41Z
             */
            created_at: string;
            /**
             * @description Unique identifier of the rendition
             * @example 5d7f910551b00a722e0418830cee6631
             */
            id: string;
            /**
             * @description MIME type of the rendition
             * @example image/png
             */
            mime_type: string;
            /**
             * @description Name of the rendition
             * @example Facebook
             */
            name: string;
            /**
             * @description URL of the rendition
             * @example https://files.cmp.optimizely.com/download/6ceee2f4fa3411ecb37802420ac8001b
             */
            url: string;
        };
        BaseContentTypeModel: {
            /** @description Status determining whether the content type is a component */
            component: boolean;
            /** @description Unique identifier of the content type */
            content_type_guid: string;
            /**
             * Format: date-time
             * @description Date and time on which the content type was created, in ISO 8601 UTC format
             */
            created_at: string;
            /** @description Unique identifier of the user who created the content type */
            created_by: string;
            /** @description Description of the content type */
            description?: string;
            /** @description Disabled status of the content type */
            disabled?: boolean;
            /** @description Meta Links */
            links: {
                /**
                 * @description URL of the content type
                 * @example https://api.cmp.optimizely.com/v3/structured-content/content-types/9fda53cf66a14fd487251480ca695c7b
                 */
                self?: string;
                /**
                 * @description URL of the content type versions
                 * @example https://api.cmp.optimizely.com/v3/structured-content/content-types/9fda53cf66a14fd487251480ca695c7b/versions
                 */
                versions?: string;
            };
            /** @description Name of the content type */
            name: string;
            /** @description Source of the content type */
            source?: string;
            /** @description Source of the content type */
            source_id?: string;
            /** @description Source metadata of the content type */
            source_metadata?: string;
            /** @description Thumbnail GUID of the content type */
            thumbnail_guid?: string;
            /**
             * Format: date-time
             * @description Date and time on which the content type was last updated, in ISO 8601 UTC format
             */
            updated_at: string;
            /** @description Unique identifier of the user who last updated the content type */
            updated_by: string;
        };
        BaseContentTypeVersionModel: {
            /**
             * Format: date-time
             * @description Date and time on which the version was created, in ISO 8601 UTC format
             */
            created_at: string;
            /** @description Unique identifier of the user who created the version */
            created_by: string;
            /** @description Expected locales of the version */
            expected_locales: string[];
            /** @description Indicates whether the version is the latest */
            latest: boolean;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the content type
                 * @example https://api.cmp.optimizely.com/v3/structured-content/content-types/555876bf63e94ea4b3ec4da318736414
                 */
                content_type?: string;
                /**
                 * @description URL of the content type version
                 * @example https://api.cmp.optimizely.com/v3/structured-content/content-types/555876bf63e94ea4b3ec4da318736414/versions/31a6d15b09c54a04935203eb958ae058
                 */
                self?: string;
            };
            /** @description Unique identifier of the version */
            version_guid: string;
        };
        BaseFieldDefinition: {
            base_type: components["schemas"]["BaseFieldDefinitionType"];
            core: components["schemas"]["CoreFieldDef"];
            default_values?: (components["schemas"]["LocationDefaultValue"] | boolean | string | Record<string, never> | unknown[])[];
        };
        /**
         * @description An enumeration.
         * @enum {string}
         */
        BaseFieldDefinitionType: "boolean" | "json" | "url" | "location";
        BaseFormFieldRequest: {
            /**
             * @description Identifier of the form field collected from the template form field
             * @example brief
             */
            identifier: string;
            /**
             * @description Type of the form field
             * @example brief
             */
            type: string;
            /** @description Values for the form field to populate in the work request */
            values: unknown[];
        };
        BaseFormFieldResponse: {
            /**
             * @description Identifier of the form field collected from the template form field
             * @example brief
             */
            identifier: string;
            /**
             * @description Type of the form field
             * @example brief
             */
            type: string;
            /** @description Values of the form field */
            values: unknown[];
        };
        BaseObjectField: components["schemas"]["BaseObjectFieldUpdatePayload"] & {
            /**
             * @description Unique identifier of the field
             * @example 41aac313bda1e40b6cde7116fc24a
             */
            id: string;
        };
        BaseObjectFieldUpdatePayload: {
            /**
             * @description Type of the field
             * @example checkbox
             */
            type: string;
            /** @description Values for the campaign field */
            values: unknown[];
        };
        BasePermissionsResponseSchema: {
            /**
             * @description level of access the accessor has to the resource
             * @example view
             * @enum {string}
             */
            access_type: "view" | "edit" | "comment" | "delete";
            /**
             * @description Unique identifier of the accessor
             * @example 5d7f910551b00a722e0418830cee6631
             */
            id: string;
            /**
             * @description Indicates if the accessor is the owner of the object
             * @example false
             */
            is_owner: boolean;
            links: {
                /**
                 * @description links for user and team type objects
                 * @example https://api.cmp.optimizely.com/v3/users/5f857f30e1c4a2038d6179e9
                 */
                accessor: string | null;
            };
            /**
             * @description Name of the accessor
             * @example John Doe
             */
            name: string;
            /**
             * @description Type of the accessor
             * @example user
             * @enum {string}
             */
            type: "user" | "team" | "organization";
        };
        BaseSettingsFieldCreatePayload: {
            /**
             * @description Used to help understand the functionality of the field.
             * @example This field is used to check options
             */
            helper_text?: string;
            /**
             * @description Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
             * @example true
             */
            is_active: boolean;
            /**
             * @description Name of the field
             * @example Sample Checkbox
             */
            name: string;
            /**
             * @description Type of the field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "date" | "image" | "rich_text" | "text" | "text_area" | "video";
        };
        BaseSettingsFieldUpdatePayload: {
            /**
             * @description Used to help understand the functionality of the field.
             * @example This is a helper text
             */
            helper_text?: string;
            /**
             * @description Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
             * @example true
             */
            is_active?: boolean;
            /**
             * @description Name of the field
             * @example Sample Field
             */
            name?: string;
            /**
             * @description Type of the field. This field is required but not updateable. Type of a field will always stay the same. Update request body must contain atleast one other updateable field. (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "checkbox" | "date" | "image" | "radio_button" | "rich_text" | "text" | "text_area" | "video";
        };
        BaseSettingsFieldsResponse: {
            /**
             * @description Used to help understand functionality of field
             * @example This field is used to check options
             */
            helper_text: string;
            /**
             * @description Identifier of the field
             * @example 6ceee2f4fa3411ecb37802420ac8001b
             */
            id: string;
            /**
             * @description Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
             * @example true
             */
            is_active: boolean;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the list of fields filtered only with the field
                 * @example https://api.cmp.optimizely.com/v3/fields?ids=8p7f910551i00722y0418830aee6612
                 */
                source: string;
            };
            /**
             * @description Name of the field
             * @example Sample Checkbox
             */
            name: string;
            /**
             * @description Type of the field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "date" | "image" | "rich_text" | "text" | "text_area" | "video";
        };
        BaseTemplateFormField: {
            /**
             * @description Helper text of the field
             * @example Sample helper text
             */
            helper_text: string | null;
            /**
             * @description Identifier of the field
             * @example attachment
             */
            identifier: string;
            /**
             * @description Specify if the field can be used to create any resource
             * @example false
             */
            is_readonly: boolean;
            /**
             * @description Required status of the field
             * @example false
             */
            is_required: boolean;
            /**
             * @description Label of the field
             * @example Attach a file...
             */
            label: string | null;
            /**
             * @description Sort order of the field
             * @example 0
             */
            sort_order: number;
            /**
             * @description Type of the field
             * @example label
             * @enum {string}
             */
            type: "brief" | "checkbox" | "currency_number" | "date" | "dropdown" | "file" | "instruction" | "label" | "percentage_number" | "radio_button" | "richtext" | "section" | "simple_number" | "text" | "text_area";
        };
        BaseTemplateResponse: {
            /** @description List of resources the template is applicable to */
            applicable_to: ("work_request" | "task_brief" | "campaign_brief")[];
            /**
             * @description Description of the template
             * @example Sample Desciption
             */
            description: string;
            /**
             * @description Unique identifier of the template
             * @example 9119a313057e401189407116fcd3aa24
             */
            id: string;
            /**
             * @description Active status of the template
             * @example true
             */
            is_active: boolean;
            /**
             * @description Title of the template
             * @example Sample Template Title
             */
            title: string;
        };
        BatchFileUrlResponse: {
            /**
             * @description List of urls for Files.
             * @example {
             *       "000": null,
             *       "791569bab45811eebb2802420ac8001c": "http://images.cmp.optimizely.com/Zz0xODQ3NDU3Y2Y2YmYzOTlmNjkyOTgyZDY3Y2I3YWM2OA==",
             *       "947769bab45811eebb2802420ac8072n": "http://images.cmp.optimizely.com/Zz0xODQ3NDU3Y2Y2YmYzOTlmNjkyOTgyZDY3Y2I3YWM2OA==?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOlsiNDVhY2MyYmVhMGExMTF"
             *     }
             */
            urls: {
                [key: string]: string | null;
            };
        };
        BooleanFieldValueModel: {
            /** @description Value of the field value */
            bool_value: boolean;
            /** @description Order index of the field value */
            order_index?: number;
        };
        BrandComplianceCategoriesResponse: {
            criteria: {
                /**
                 * @description Details about the criteria
                 * @example Description-1A
                 */
                description?: string;
                /**
                 * @description Unique identifier of the criteria
                 * @example 66cebe190bc97e0151eecddb
                 */
                id: string;
                /**
                 * @description Name of the criteria
                 * @example Sub-category-1A
                 */
                name?: string;
            }[];
            /**
             * @description Unique identifier of the category
             * @example 66cebe190bc97e0151eecdda
             */
            id: string;
            /**
             * @description Name of the category
             * @example Category-1
             */
            name?: string;
        };
        BriefTypeFormFieldRequest: components["schemas"]["BaseFormFieldRequest"] & {
            /** @description Array of brief values */
            values?: (components["schemas"]["WorkRequestTextBriefRequest"] | components["schemas"]["WorkRequestAttachmentBriefRequest"])[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "brief";
        };
        BriefTypeFormFieldResponse: components["schemas"]["BaseFormFieldResponse"] & {
            /** @description Array of brief values */
            values?: (components["schemas"]["WorkRequestTextBriefResponse"] | components["schemas"]["WorkRequestAttachmentBriefResponse"])[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "brief";
        };
        /** @description Budget of the campaign */
        BudgetResponse: {
            /**
             * @description Planned amount for the campaign
             * @example 1000.50
             */
            budgeted_amount: string;
            /**
             * @description Currency code of the budget
             * @example USD
             */
            currency_code: string;
        };
        CampaignBriefResponse: {
            /** @description List of fields of the brief */
            fields: {
                /**
                 * @description Name of the brief field
                 * @example My Dropdown
                 */
                name: string;
                /** @description List of values set for the brief field */
                value: string[];
            }[];
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the campaign
                 * @example https://api.cmp.optimizely.com/v3/campaigns/5f857f30e1c4a2038d6179e9
                 */
                campaign: string;
                /**
                 * @description URL of the brief
                 * @example https://api.cmp.optimizely.com/v3/campaigns/5f857f30e1c4a2038d6179e9/brief
                 */
                self: string;
            };
            template?: (components["schemas"]["CampaignBriefTemplateValue"] | components["schemas"]["NullValue"]) | null;
            /**
             * @description Title of the campaign brief
             * @example Awesome Campaign Brief
             */
            title: string;
            /**
             * @description Type of the campaign brief
             * @example template
             * @enum {string}
             */
            type: "template" | "attachment" | "text";
        };
        /** @description Template info of the brief if the brief is of template type */
        CampaignBriefTemplateValue: {
            /**
             * @description Unique identifier of the template
             * @example 9nu8ue9wf8u9nusd9f
             */
            id: string;
            /**
             * @description Name of the template
             * @example My Template
             */
            name: string;
        };
        CampaignCommentResponse: {
            /**
             * @description List of attachments of the comment
             * @example [
             *       {
             *         "id": "a113667245d111eb8945000c",
             *         "name": "sample.png",
             *         "url": "https://files.cmp.optimizely.com/download/96c314a645d111eb8945000c291b51d4?token="
             *       }
             *     ]
             */
            attachments: components["schemas"]["AttachmentResponse"][];
            /**
             * @description Unique identifier of the user who posted the comment
             * @example 62a8455e29fed3045d9c7a8b
             */
            comment_by: string;
            /**
             * Format: date-time
             * @description Creation date and time of the campaign comment in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2020-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Unique identifier of the campaign comment
             * @example 5fe38c39574b52a62a089239
             */
            id: string;
            /**
             * @description Determines if the comment is resolved in Optimizely CMP
             * @example false
             */
            is_resolved: boolean;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the campaign
                 * @example https://api.cmp.optimizely.com/v3/campaigns/5fe38c39574b52a62a089239
                 */
                campaign: string;
                /**
                 * @description URL of the user who posted the comment
                 * @example https://api.cmp.optimizely.com/v3/users/62a8455e29fed3045d9c7a8b
                 */
                comment_by: string;
            };
            /**
             * Format: date-time
             * @description Last modification date and time of the campaign comment in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2020-10-06T14:15:30Z
             */
            modified_at: string;
            /**
             * @description Unique identifier of the parent comment
             * @example 5fe38c39574b52a62a089238
             */
            parent_comment_id: string | null;
            /**
             * @description Content of the comment
             * @example Well done!
             */
            value: string;
        };
        CampaignCreateRequest: {
            /**
             * @description Color of the campaign
             * @example #ec2e3b
             * @enum {string}
             */
            color?: "#333" | "#ec2e3b" | "#fc594b" | "#e2575b" | "#f45097" | "#de55a9" | "#bd417f" | "#a959cc" | "#6456b7" | "#4db2fc" | "#3a97be" | "#36c5a3" | "#9bc94e" | "#fdad3c" | "#ef803c" | "#fd8876" | "#ea868a" | "#f781b5" | "#e785c1" | "#cf69a3" | "#c187db" | "#9086cc" | "#79c8fd" | "#5ab5d0" | "#4fd5bd" | "#b8d877" | "#fec454" | "#f3a35a" | "#d3998f" | "#c79697" | "#d096af" | "#c897b4" | "#ae7c98" | "#b094be" | "#908bae" | "#97bed8" | "#7aa8b6" | "#7abeb2" | "#bece9d" | "#e3c68d" | "#d3ac85" | "#b9a29e" | "#b29f9f" | "#b9a1ac" | "#b5a1ae" | "#9a8692" | "#a89bac" | "#908e9b" | "#a8b9c3" | "#8ea1a7" | "#96b1ac" | "#c1c8b5" | "#d4c8b0" | "#c0b0a1" | "#212121" | "#CE515B" | "#F16281" | "#EB7777" | "#A25555" | "#B05F8C" | "#AF8EC3" | "#21B7EC" | "#3AB0C9" | "#669BA8" | "#7ABEB2" | "#5F9DCE" | "#717CB9" | "#4C557E" | "#449C6C" | "#63B05F" | "#4D9B93" | "#ACC050" | "#A29F55" | "#8AA872" | "#E4844C" | "#EE9808" | "#DE7009" | "#FCD404" | "#CBB039" | "#99896A" | "#2e2a26" | "#ccccdc";
            /**
             * @description Description of the campaign
             * @example Sample Description
             */
            description?: string | null;
            /**
             * Format: date
             * @description End date of the campaign, in ISO 8601 UTC format (2026-02-10).
             * @example 2026-10-06
             */
            end_date?: string;
            /**
             * @description ID of the owner of the campaign
             * @example 665d5717c8ba0201647a0991
             */
            owner_id?: string;
            /**
             * @description ID of the parent campaign
             * @example 665d5717c8ba0201647a0992
             */
            parent_campaign?: string | null;
            /**
             * Format: date
             * @description Start date of the campaign, in ISO 8601 UTC format (2025-02-10).
             * @example 2025-10-01
             */
            start_date?: string;
            /**
             * @description Title of the campaign
             * @example Sample Title
             */
            title: string;
        };
        /**
         * @description Type of the field
         * @example image
         * @enum {string}
         */
        CampaignFieldTypes: "checkbox" | "currency_number" | "date" | "dropdown" | "label" | "radio_button" | "text_area" | "percentage_number" | "rich_text" | "simple_number" | "text" | "image" | "video";
        CampaignFieldUpdateRequest: components["schemas"]["StringTypeObjectFieldUpdatePayload"] | components["schemas"]["MultiChoiceTypeObjectFieldUpdatePayload"] | components["schemas"]["DropdownTypeObjectFieldUpdatePayload"] | components["schemas"]["RadioButtonTypeObjectFieldUpdatePayload"] | components["schemas"]["NumberTypeObjectFieldUpdatePayload"] | components["schemas"]["DateTypeObjectFieldUpdatePayload"] | components["schemas"]["AssetTypeObjectFieldUpdatePayload"];
        CampaignFieldUpdateResponse: {
            /**
             * @description Unique identifier for the field
             * @example 5a7f910511b0a72230ce6631
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the campaign
                 * @example https://api.cmp.optimizely.com/v3/campaigns/5f857f30e1c4a2038d6179e9
                 */
                campaign: string;
                /**
                 * @description URL of the campaign fields
                 * @example https://api.cmp.optimizely.com/v3/campaigns/5f857f30e1c4a2038d6179e9/fields
                 */
                campaign_fields: string;
            };
            /**
             * @description Name of the field
             * @example Image Field
             */
            name: string;
            type: components["schemas"]["CampaignFieldTypes"];
            /** @description List of values selected or from user input */
            values: string[];
        };
        CampaignListResponseItem: {
            /**
             * @description Description of the campaign
             * @example The awesome campaign's description
             */
            description: string | null;
            /**
             * @description End date of the campaign in ISO 8601 format: `YYYY-MM-DD`
             * @example 2019-10-16
             */
            end_date: string | null;
            /**
             * @description Unique identifier of the campaign
             * @example 8q7f910551b00a722e0418830cee6612
             */
            id: string;
            /**
             * @description Whether the campaign is hidden or not
             * @example false
             */
            is_hidden: boolean;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the campaign
                 * @example https://api.cmp.optimizely.com/v3/campaigns/8q7f910551b00a722e0418830cee6612
                 */
                self: string;
            };
            /** @description Owner information of the campaign */
            owner: {
                /**
                 * @description Unique identifier of the owner of the campaign
                 * @example 8q7f910551b00a722e0418830cee8211
                 */
                id?: string;
            };
            /**
             * @description Reference identifier of the campaign
             * @example CPN-1589
             */
            reference_id: string;
            /**
             * @description Start date of the campaign in ISO 8601 format: `YYYY-MM-DD`
             * @example 2019-10-06
             */
            start_date: string | null;
            /**
             * @description Title of the campaign
             * @example The awesome campaign
             */
            title: string;
        };
        CampaignResponse: {
            budget: (components["schemas"]["BudgetResponse"] | components["schemas"]["NullValue"]) | null;
            /**
             * Format: date-time
             * @description Creation date and time of the campaign,  in ISO 8601 UTC format (`2035-02-10T10:40:45Z`)
             * @example 2035-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Description of the campaign
             * @example The awesome campaign's description
             */
            description: string | null;
            /**
             * @description End date of the campaign in ISO 8601 format: `YYYY-MM-DD`
             * @example 2019-10-16
             */
            end_date: string | null;
            /**
             * @description Unique identifier of the campaign
             * @example 8q7f910551b00a722e0418830cee6612
             */
            id: string;
            /**
             * @description Whether the campaign is hidden or not
             * @example false
             */
            is_hidden: boolean;
            /** @description Labels associated to the campaign */
            labels: components["schemas"]["ResourceLabelResponse"][];
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the campaign brief
                 * @example https://api.cmp.optimizely.com/v3/campaigns/8q7f910551b00a722e0418830cee6612/brief
                 */
                brief: string | null;
                /** @description URLs of the child campaigns */
                child_campaigns: string[];
                /**
                 * @description URL of the campaign owner
                 * @example https://api.cmp.optimizely.com/v3/users/6ea1bf781398c3147393arr1
                 */
                owner: string | null;
                /**
                 * @description URL of the parent campaign
                 * @example https://api.cmp.optimizely.com/v3/campaigns/7u7f910551b00a722e0418830cee6643
                 */
                parent_campaign: string | null;
                /**
                 * @description URL of the campaign
                 * @example https://api.cmp.optimizely.com/v3/campaigns/8q7f910551b00a722e0418830cee6612
                 */
                self: string;
            };
            /**
             * @description Reference identifier of the campaign
             * @example CPN-1589
             */
            reference_id: string;
            /**
             * @description Start date of the campaign in ISO 8601 format: `YYYY-MM-DD`
             * @example 2019-10-06
             */
            start_date: string | null;
            /**
             * @description Status of the campaign
             * @example On Track
             * @enum {string}
             */
            status: "Not Started" | "Off Track" | "On Track" | "Complete" | "At Risk";
            /**
             * @description Title of the campaign
             * @example The awesome campaign
             */
            title: string;
        };
        CampaignUpdateRequest: {
            /**
             * @description End date of the campaign in ISO 8601 format: `YYYY-MM-DD`
             * @example 2022-09-06
             */
            end_date?: string | null;
            /**
             * @description Unique identifier of the campaign owner
             * @example 8q7f910551b00a722e0418830cee6612
             */
            owner_id?: string;
            /**
             * @description Start date of the campaign in ISO 8601 format: `YYYY-MM-DD`
             * @example 2022-09-06
             */
            start_date?: string | null;
            /**
             * @description Title of the campaign
             * @example The awesome campaign
             */
            title?: string;
        };
        CheckboxAndRadioTypeSettingsFieldCreatePayload: components["schemas"]["BaseSettingsFieldCreatePayload"] & {
            /** @description Choices of the field */
            choices: {
                /**
                 * @description Name of the choice
                 * @example Choice 1
                 */
                name: string;
            }[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "checkbox" | "radio_button";
        };
        CheckboxAndRadioTypeSettingsFieldResponse: components["schemas"]["BaseSettingsFieldsResponse"] & {
            /** @description Choices of the field */
            choices: {
                /**
                 * @description Identifier of the choice
                 * @example 6ceee2f4fa3411ecb37802420ac8001b
                 */
                id: string;
                /**
                 * @description Name of the choice
                 * @example Choice 1
                 */
                name: string;
            }[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "checkbox" | "radio_button";
        };
        /**
         * @description An enumeration.
         * @enum {string}
         */
        ChoiceDisplayOption: "radio" | "dropdown" | "tag" | "checkbox";
        ChoiceFieldDefinition: {
            /** @description Choices of the field */
            choices: {
                [key: string]: string;
            };
            core: components["schemas"]["CoreFieldDef"];
            /** @description Default values of the field */
            default_values?: string[];
            display_option?: components["schemas"]["ChoiceDisplayOption"];
        };
        ChoiceFieldValueModel: {
            /** @description Choice key of the field value */
            choice_key: string;
            /** @description Order index of the field value */
            order_index?: number;
        };
        CommentCreateRequest: {
            /** @description List of comment attachments. */
            attachments?: components["schemas"]["AttachmentRequest"][];
            /**
             * @description Content of the comment. Markdown is supported. To mention a user belonging to the organization, use the `@[name](openapi-user-link)` format.
             * @example Comment mentioning @[Organization User](https://api.cmp.optimizely.com/v3/users/5fe38aeb574b52a62a089238)
             */
            value: string;
        };
        CommentWithReplyCreateRequest: {
            /** @description List of comment attachments. */
            attachments?: components["schemas"]["AttachmentRequest"][];
            /**
             * @description Parent comment id to reply.
             * @example da80cfd7bbf84959a8a981acbad996b3
             */
            parent_comment_id?: string;
            /**
             * @description Content of the comment. Markdown is supported. To mention a user belonging to the organization, use the `@[name](openapi-user-link)` format.
             * @example Comment mentioning @[Organization User](https://api.cmp.optimizely.com/v3/users/5fe38aeb574b52a62a089238)
             */
            value: string;
        };
        ContentDetailsModel: {
            /** @description Unique identifier of the content */
            content_guid: string;
            content_type?: components["schemas"]["VersionedContentTypeModel"];
            /** @description Unique identifier of the content type */
            content_type_guid: string;
            /** @description Name of the content type */
            content_type_name: string;
            /**
             * Format: date-time
             * @description Date and time on which the content was created, in ISO 8601 UTC format
             */
            created_at: string;
            /** @description Unique identifier of the user who created the content */
            created_by: string;
            /** @description Expired status of the content */
            expired?: boolean;
            /**
             * Format: date-time
             * @description Date and time on which the content will expire, in ISO 8601 UTC format
             */
            expiry_datetime?: string;
            latest_fields_version: components["schemas"]["ContentFieldsVersionDetails"];
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the content definition
                 * @example https://api.cmp.optimizely.com/v3/structured-content/content-types/5291e00b990a49f1857adc024fd26620/versions/a8b48d529701452aa56a4e752e1b0ffc
                 */
                definition?: string;
                /**
                 * @description URL of the content
                 * @example https://api.cmp.optimizely.com/v3/structured-content/contents/a24b834428a043ab8caed1ded4606f7d
                 */
                self?: string;
            };
            /**
             * @description The primary locale of the content
             * @example en_US
             */
            primary_locale?: string;
            /** @description This is true when the content is not embedded in another content through a content type reference field */
            root_content: boolean;
            /** @description Source of the content */
            source?: string;
            /** @description Source ID of the content */
            source_id?: string;
            /** @description Source metadata of the content */
            source_metadata?: string;
            /** @description Template GUID of the content */
            template_guid?: string;
            /** @description Title of the content */
            title: string;
            /**
             * Format: date-time
             * @description Date and time on which the content was last updated, in ISO 8601 UTC format
             */
            updated_at: string;
            /** @description Unique identifier of the user who last updated the content */
            updated_by: string;
        };
        ContentFieldValueExpandedModel: {
            content_details: components["schemas"]["ContentDetailsModel"];
            /** @description Content GUID of the field value */
            content_guid: string;
            /** @description Content URL of the field value */
            content_url?: string;
            /** @description Indicates whether the field value is embedded */
            embedded: boolean;
            /** @description Order index of the field value */
            order_index?: number;
        };
        ContentFieldValueModel: {
            /** @description Content GUID of the field value */
            content_guid: string;
            /** @description Content URL of the field value */
            content_url?: string;
            /** @description Indicates whether the field value is embedded */
            embedded: boolean;
            /** @description Order index of the field value */
            order_index?: number;
        };
        ContentFieldValueWithEmbeddedModel: {
            /**
             * Contentdetails
             * @description The embedded content object
             */
            content_details: components["schemas"]["RecursiveStructuredContent"];
        };
        ContentFieldWithEmbeddedPatchValueModel: {
            content_id: string;
            patch_details: components["schemas"]["RecursivePatchStructuredContentFields"];
        };
        ContentFieldsVersionDetails: {
            /** @description Content hash of the version */
            content_hash: string;
            /**
             * Format: date-time
             * @description Date and time on which the version was created, in ISO 8601 UTC format
             */
            created_at: string;
            /** @description Unique identifier of the user who created the version */
            created_by: string;
            /** @description List of fields */
            fields: {
                [key: string]: components["schemas"]["LocalizedFieldValues"][];
            };
            /** @description Source ID of the version */
            source_id?: string;
            /** @description Source metadata of the version */
            source_metadata?: string;
            validation?: components["schemas"]["ContentFieldsVersionValidation"];
            /** @description Unique identifier of the version */
            version_guid: string;
        };
        ContentFieldsVersionValidation: {
            /** @description Validation for the fields */
            fields?: {
                [key: string]: {
                    [key: string]: components["schemas"]["error_reason"];
                };
            };
        };
        ContentMigrationSummary: {
            /** @description The number of content items that failed to migrate. */
            errored?: number;
            /** @description The number of content items that haven't started migrating. */
            not_started?: number;
            /** @description The number of content items that are skipped due to content type version ID mismatch. */
            skipped?: number;
            /** @description The number of content items that migrated successfully. */
            succeeded?: number;
            /** @description The total number of content items. */
            total?: number;
        };
        ContentTypeFieldDefinition: {
            /** @description Whether to allow ref editing */
            allow_ref_edit?: boolean;
            /** @description List of allowed content types */
            allowed_content_types: string[];
            /** @description Links related to the content type */
            content_type_links?: {
                [key: string]: components["schemas"]["AllowedContentTypeItem"];
            };
            core: components["schemas"]["CoreFieldDef"];
            /** @description Default value for the field */
            default_value?: string;
            ref_type: components["schemas"]["ContentTypeFieldEmbedMixConfig"];
        };
        /**
         * @description Ref Type:
         *       * `1` Refer only
         *       * `2` Refer and create embed
         *       * `3` Create only
         * @enum {integer}
         */
        ContentTypeFieldEmbedMixConfig: 1 | 2 | 3;
        /**
         * @description List Type:
         *       * `1` - Component only
         *       * `2` - Content only
         *       * `3` - Component and Content only
         * @enum {string}
         */
        ContentTypeListingOption: "1" | "2" | "3";
        CoreContentType: {
            /** @description Indicates whether the content type is a component */
            component: boolean;
            /** @description Description of the content type */
            description?: string;
            /** @description Disabled status of the content type */
            disabled?: boolean;
            /** @description Name of the content type */
            name: string;
            /** @description Thumbnail GUID of the content type */
            thumbnail_guid?: string;
        };
        CoreFieldDef: {
            /** @description Editor metadata of the field */
            editor_metadata?: Record<string, never> | unknown[];
            field_type?: components["schemas"]["FieldType"];
            /** @description Helptext of the field */
            help_text?: string;
            /** @description Indicates whether the field is a list */
            is_list: boolean;
            /** @description Required status of the field */
            is_required: boolean;
            /** @description Key of the field */
            key: string;
            /** @description Maximum length of the list */
            max_list_length?: number;
            /** @description Minimum length of the list */
            min_list_length?: number;
            /** @description Name of the field */
            name: string;
            /** @description Indicates whether the field needs internationalization */
            need_internationalization: boolean;
            /** @description Order index of the field */
            order_index?: number;
            /** @description Source ID of the field */
            source_id?: string;
            /** @description Source metadata of the field */
            source_metadata?: string;
        };
        CreativeAssetRequest: {
            /**
             * @description Unique identifier of the file upload session. This is the `upload_meta_fields.key` field retrieved from the `/v3/upload-url` endpoint.
             * @example ce8995aea58b11ea8cd90242ac120005
             */
            key: string;
            /**
             * @description Filename of the creative asset
             * @example sample_image.png
             */
            name: string;
        };
        CreativeAssetResponse: {
            /**
             * @description Unique identifier for the creative asset
             * @example 5a7f910511b0a72230ce6631
             */
            id: string;
            /**
             * @description Filename of the creative asset
             * @example sample_image.png
             */
            name: string;
            /**
             * @description Download URL of the creative asset
             * @example https://files.cmp.optimizely.com/download/2115bfe4450c11ebaae8000c291b51d4
             */
            url: string;
        };
        CurrencyCustomField: {
            /**
             * @description Curency code of custom field
             * @default USD
             * @example USD
             */
            currency_code: string;
            /**
             * @description Value of decimal places
             * @default null
             * @example 1
             */
            decimal_places: number | null;
            /**
             * @description Indicator of thousand separator
             * @default false
             * @example false
             */
            has_thousand_separator: boolean;
            /**
             * @description Name of the custom field
             * @example A Currency Custom Field
             */
            label: string;
        };
        CurrencyNumberTypeSettingsFieldCreatePayload: components["schemas"]["GenericNumberTypeSettingsFieldCreatePayload"] & {
            /**
             * @description Currency code of the numerical field. This must be a valid currency code. If this field is empty, `USD` is set as the default.
             * @default USD
             * @example USD
             */
            currency_code: string;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "currency_number";
        };
        CurrencyNumberTypeSettingsFieldResponse: components["schemas"]["GenericNumberTypeSettingsFieldResponse"] & {
            /**
             * @description Currency code of the numerical field
             * @example USD
             */
            currency_code: string;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "currency_number";
        };
        CurrencyNumberTypeSettingsFieldUpdatePayload: components["schemas"]["GenericNumberTypeSettingsFieldUpdatePayload"] & {
            /**
             * @description Currency code of the numerical field. This must be a valid currency code. If this field is empty, `USD` is set as the default.
             * @default USD
             * @example USD
             */
            currency_code: string;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "currency_number";
        };
        DateTypeFormFieldRequest: components["schemas"]["BaseFormFieldRequest"] & {
            /** @description Array of date-time values, in ISO 8601 UTC format (2020-02-10T10:40:45Z) */
            values?: string[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "date";
        };
        DateTypeFormFieldResponse: components["schemas"]["BaseFormFieldResponse"] & {
            /** @description Array of date-time values, in ISO 8601 UTC format (2020-02-10T10:40:45Z) */
            values?: string[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "date";
        };
        DateTypeObjectField: components["schemas"]["BaseObjectField"] & components["schemas"]["DateTypeObjectFieldUpdatePayload"] & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "date";
        };
        DateTypeObjectFieldUpdatePayload: components["schemas"]["BaseObjectFieldUpdatePayload"] & {
            /** @description Array of date-time values, in ISO 8601 UTC format (2020-02-10T10:40:45Z) */
            values?: string[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "date";
        };
        DatetimeFieldDefinition: {
            core: components["schemas"]["CoreFieldDef"];
            /** @description Default values of the field */
            default_values?: string[];
            /**
             * Format: date-time
             * @description Maximum date of the field
             */
            max_date: string;
            /**
             * Format: date-time
             * @description Minimum date of the field
             */
            min_date: string;
        };
        DatetimeFieldValueModel: {
            /**
             * Format: date-time
             * @description Value of the field value
             */
            datetime_value: string;
            /** @description Order index of the field value */
            order_index?: number;
        };
        DeleteFieldValueModel: {
            /** @description Should be true if the field value needs to be deleted */
            delete: boolean;
            /** @description Order index of the field value */
            order_index?: number;
        };
        DeleteLocaleFieldValueModel: {
            /** @description Should be true if the field value with locale needs to be deleted */
            delete: boolean;
            locale: string;
        };
        DetailedAssetRenditionResponse: components["schemas"]["BaseAssetRenditionResponse"] & {
            /**
             * @description Alternative text for rendition
             * @example image title
             */
            alt_text: string | null;
            /**
             * @description Type of the original asset the rendition was generated from
             * @example image
             */
            asset_type: string;
            /**
             * @description The height (in pixels) of the rendition
             * @example 400
             */
            height: number | null;
            /**
             * Format: date-time
             * @description Date and time on which the rendition was last modified, in ISO 8601 UTC format
             * @example 2022-09-27T09:14:30Z
             */
            modified_at: string;
            /**
             * @description ID of the original asset the rendition was generated from
             * @example 5d7f910551b00a722e0418830cee6631
             */
            original_asset_id: string;
            /**
             * @description ID of rendition config used to generate the rendition
             * @example f910551b00ae0418830cee6631452651
             */
            rendition_config_id?: string;
            /**
             * @description Status of the rendition's generation
             * @example Done
             * @enum {string}
             */
            status: "Done" | "InProgress" | "Error";
            /**
             * @description The width (in pixels) of the rendition
             * @example 600
             */
            width: number | null;
        };
        DropdownTypeObjectField: components["schemas"]["BaseObjectField"] & components["schemas"]["DropdownTypeObjectFieldUpdatePayload"] & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "dropdown";
        };
        DropdownTypeObjectFieldUpdatePayload: components["schemas"]["BaseObjectFieldUpdatePayload"] & {
            /** @description Array of choice ID. Multiple choice values are not acceptable for `is_multi_select=false`. */
            values?: string[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "dropdown";
        };
        DropdownTypeSettingsFieldCreatePayload: components["schemas"]["BaseSettingsFieldCreatePayload"] & {
            /** @description Choices of the field */
            choices: {
                /**
                 * @description Name of the choice
                 * @example Choice 1
                 */
                name: string;
            }[];
            /** @description Allow users to select multiple values */
            is_multi_select: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "dropdown";
        };
        DropdownTypeSettingsFieldResponse: components["schemas"]["BaseSettingsFieldsResponse"] & {
            /** @description Choices of the field */
            choices: {
                /**
                 * @description Identifier of the choice
                 * @example 6ceee2f4fa3411ecb37802420ac8001b
                 */
                id: string;
                /**
                 * @description Name of the choice
                 * @example Choice 1
                 */
                name: string;
            }[];
            /** @description Allow users to select multiple values */
            is_multi_select: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "dropdown";
        };
        /** @description Error payload */
        Error: {
            /** @description Additional information */
            errors?: {
                [key: string]: unknown;
            };
            /**
             * @description Message describing the error
             * @example Not found
             */
            message: string;
        } & {
            [key: string]: unknown;
        };
        EventCreateRequest: {
            /**
             * @description ID of the campaign under which the event is to be created
             * @example 8p7f91i55722y04188a30cee6612
             */
            campaign_id?: string;
            /**
             * @description Description of the event
             * @example This awesome event is created to celebrate our clients
             */
            description?: string;
            /**
             * Format: date-time
             * @description End date and time of the event in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`
             * @example 2024-12-24T12:18:00Z
             */
            end_date: string;
            /** @description List of fields to be associated with the event */
            fields?: (components["schemas"]["StringTypeObjectField"] | components["schemas"]["MultiChoiceTypeObjectField"] | components["schemas"]["DropdownTypeObjectField"] | components["schemas"]["RadioButtonTypeObjectField"] | components["schemas"]["NumberTypeObjectField"] | components["schemas"]["DateTypeObjectField"] | components["schemas"]["AssetTypeObjectField"])[];
            /**
             * @description Indicate if the event is all day long
             * @example false
             */
            is_all_day: boolean;
            /**
             * Format: date-time
             * @description Start date and time of the event in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`
             * @example 1996-09-09T12:18:00Z
             */
            start_date: string;
            /**
             * @description Title of the event
             * @example The awesome event
             */
            title: string;
        };
        EventFieldsUpdateRequest: (components["schemas"]["MultiChoiceTypeObjectField"] | components["schemas"]["StringTypeObjectField"] | components["schemas"]["DropdownTypeObjectField"] | components["schemas"]["RadioButtonTypeObjectField"] | components["schemas"]["NumberTypeObjectField"] | components["schemas"]["DateTypeObjectField"] | components["schemas"]["AssetTypeObjectField"])[];
        EventResponse: {
            /**
             * @description Unique identifier of the campaign associated with the event
             * @example 66d953340019d7b86833ac6dg
             */
            campaign_id: string | null;
            /**
             * @description Unique identifier of the user who created the event
             * @example 66d837c62373533177b59db3
             */
            created_by: string;
            /**
             * @description Description of the event
             * @example The awesome event's description
             */
            description: string | null;
            /**
             * @description End date of the event in ISO 8601 UTC format
             * @example 2024-11-26T11:30:34Z
             */
            end_date: string;
            /**
             * @description Unique identifier of the event
             * @example 674ea40a7a9cac80ff78c95c
             */
            id: string;
            /**
             * @description Whether the event is a day long event or not
             * @example false
             */
            is_all_day: boolean;
            /**
             * @description Whether the event is archived or not
             * @example false
             */
            is_archived: boolean;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the campaign associated with the event
                 * @example https://api.cmp.optimizely.com/v3/campaigns/66d953340019d7b86833ac6d
                 */
                campaign: string | null;
                /**
                 * @description URL of the event
                 * @example https://api.cmp.optimizely.com/v3/events/674ea40a7a9cac80ff78c95c
                 */
                self: string;
            };
            /**
             * @description Reference ID of the event
             * @example EVT-3
             */
            reference_id: string;
            /**
             * @description Start date of the event in ISO 8601 UTC format
             * @example 2024-11-26T11:00:34Z
             */
            start_date: string;
            /**
             * @description Title of the event
             * @example The awesome event
             */
            title: string;
        };
        EventUpdateRequest: {
            /**
             * @description Description of the event
             * @example This awesome event is created to celebrate our clients
             */
            description?: string;
            /**
             * Format: date-time
             * @description End date and time of the event in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`
             * @example 2024-12-24T12:18:00Z
             */
            end_date?: string;
            /**
             * @description Indicate if the event is all day long
             * @example false
             */
            is_all_day?: boolean;
            /**
             * Format: date-time
             * @description Start date and time of the event in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`
             * @example 1996-09-09T12:18:00Z
             */
            start_date?: string;
            /**
             * @description Title of the event
             * @example The awesome event
             */
            title?: string;
        };
        FeaturedImageResponse: {
            /**
             * @description Attribution text of the featured image
             * @example This is a sample attribution text
             */
            attribution_text: string | null;
            /**
             * @description Caption of the featured image
             * @example This is a sample caption
             */
            caption: string;
            /**
             * @description Description of the featured image
             * @example This is a sample description
             */
            description: string | null;
            /**
             * @description Height of the featured image in pixels
             * @example 400
             */
            height: number;
            /**
             * @description MIME type of the featured image
             * @example image/jpeg
             */
            mime_type: string;
            /** @description Source of the featured image */
            source: {
                /**
                 * @description Organization or vendor name of the featured image
                 * @example Reuters
                 */
                name: string | null;
            };
            /**
             * @description URL of the featured image thumbnail
             * @example https://images-cdn.cmp.optimizely.com/Zz01Y2FkOWFjZWQ5OTMxMWViYjY3OTEzNDMzOWM3ZDNhNA==?width=75&height=75
             */
            thumbnail: string | null;
            /**
             * @description URL of the featured image
             * @example https://images-cdn.cmp.optimizely.com/sample.jpeg
             */
            url: string;
            /**
             * @description Width of the featured image in pixels
             * @example 700
             */
            width: number;
        };
        FieldBaseType: {
            /**
             * @description Identifier of the field
             * @example 64be245ec0d79e79fdf1ad84
             */
            id: string;
            /**
             * @description Name of the field
             * @example Dropdown 101
             */
            name: string;
            /**
             * @description Type of the field
             * @example text
             */
            type: string;
            /** @description List of selected values or from user input */
            values: string[];
        };
        FieldListResponseItem: components["schemas"]["FieldTypeCommon"] | components["schemas"]["FieldTypeLabel"] | components["schemas"]["FieldTypeDropdown"] | components["schemas"]["FieldTypeRadio"] | components["schemas"]["FieldTypeCheckbox"] | components["schemas"]["FieldTypeSimpleNumber"] | components["schemas"]["FieldTypePercentageNumber"] | components["schemas"]["FieldTypeCurrencyNumber"];
        /**
         * @description An enumeration.
         * @enum {string}
         */
        FieldType: "boolean" | "number" | "text-field" | "rich-text" | "datetime" | "library-asset" | "content-type" | "json" | "url" | "choice" | "location";
        FieldTypeCheckbox: components["schemas"]["FieldBaseType"] & {
            /** @description Choices of the checkbox */
            choices: {
                /**
                 * @description Identifier of the choice
                 * @example 6ceee2f4fa3411ecb37802420ac8001b
                 */
                id: string;
                /**
                 * @description Name of the choice
                 * @example Checkbox 1 Choice 1
                 */
                name: string;
            }[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "checkbox";
        };
        FieldTypeCommon: components["schemas"]["FieldBaseType"] & Record<string, never> & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "date" | "image" | "rich_text" | "text" | "text_area" | "video";
        };
        FieldTypeCurrencyNumber: components["schemas"]["FieldBaseType"] & {
            /**
             * @description Currency code of the numerical field
             * @example USD
             */
            currency_code: string;
            /**
             * @description Decimal place of the numerical field
             * @example 2
             */
            decimal_places: number | null;
            /**
             * @description Whether the numerical field has a thousand separator
             * @example true
             */
            has_thousand_separator: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "currency_number";
        };
        FieldTypeDropdown: components["schemas"]["FieldBaseType"] & {
            /** @description Choices of the dropdown */
            choices: {
                /**
                 * @description Identifier of the choice
                 * @example 6ceee2f4fa3411ecb37802420ac8001b
                 */
                id: string;
                /**
                 * @description Name of the choice
                 * @example Dropdown 1 Choice 1
                 */
                name: string;
            }[];
            /** @description Select multiple values from the dropdown */
            is_multi_select: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "dropdown";
        };
        FieldTypeLabel: components["schemas"]["FieldBaseType"] & {
            /** @description Choices of the label */
            choices: {
                /**
                 * @description Identifier of the choice
                 * @example 6ceee2f4fa3411ecb37802420ac8001b
                 */
                id: string;
                /**
                 * @description Name of the choice
                 * @example Label 1 Choice 1
                 */
                name: string;
            }[];
            /** @description Select multiple values from the label */
            is_multi_select: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "label";
        };
        FieldTypePercentageNumber: components["schemas"]["FieldBaseType"] & {
            /**
             * @description Decimal place of the numerical field
             * @example 2
             */
            decimal_places: number | null;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "percentage_number";
        };
        FieldTypeRadio: components["schemas"]["FieldBaseType"] & {
            /** @description Choices of the radio button */
            choices: {
                /**
                 * @description Identifier of the choice
                 * @example 6ceee2f4fa3411ecb37802420ac8001b
                 */
                id: string;
                /**
                 * @description Name of the choice
                 * @example Radio 1 Choice 1
                 */
                name: string;
            }[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "radio_button";
        };
        FieldTypeSimpleNumber: components["schemas"]["FieldBaseType"] & {
            /**
             * @description Decimal place of the numerical field
             * @example 2
             */
            decimal_places: number | null;
            /**
             * @description Whether the numerical field has a thousand separator
             * @example true
             */
            has_thousand_separator: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "simple_number";
        };
        FileTypeFormFieldRequest: components["schemas"]["BaseFormFieldRequest"] & {
            /** @description Array of file inputs to add a file */
            values?: components["schemas"]["FileTypeFormFieldValueRequest"][];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "file";
        };
        FileTypeFormFieldResponse: components["schemas"]["BaseFormFieldResponse"] & {
            /** @description Array of files */
            values?: components["schemas"]["FileTypeFormFieldValueResponse"][];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "file";
        };
        FileTypeFormFieldValueRequest: {
            /**
             * @description Unique identifier of the file upload session. This is the `upload_meta_fields.key` field retrieved from the `/v3/upload-url` endpoint.
             * @example ce8995aea58b11ea8cd90242ac120005
             */
            key: string;
            /** @description Name of the file */
            name: string;
        };
        FileTypeFormFieldValueResponse: {
            /**
             * @description Unique identifier of the file
             * @example 6dzf4ds4fds4f4564dzsfd
             */
            id: string;
            /**
             * @description Name of the file
             * @example Sample Name
             */
            name: string;
            /**
             * @description URL of the file
             * @example https://files.cmp.optimizely.com/download/8nv8svsdnivjusdijuvisdv?token=token
             */
            url: string;
        };
        FileUrlBulkCreateRequest: {
            /** @description List of guids of assets representing binary files to generate URLs for the corresponding files */
            guids: string[];
        };
        FolderCreateRequest: {
            /**
             * @description Name of the folder
             * @example Optimizely Digital Asset Management
             */
            name: string;
            /**
             * @description ID of the parent folder
             * @example 5e6f910551b00a722e0418830cee6630
             */
            parent_folder_id?: string | null;
        };
        FolderPermissionBulkCreateRequest: {
            /** @description List of permissions to be granted to folder */
            permissions: {
                /**
                 * @description level of access the accessor has to the folder
                 * @example view
                 * @enum {string}
                 */
                access_type: "view" | "edit";
                /**
                 * @description Unique identifier of accessor
                 * @example 5d7f910551b00a722e0418830cee6632
                 */
                id: string;
                /**
                 * @description Indicates if the accessor is the owner of the folder. A team cannot be an owner.
                 * @default false
                 * @example false
                 */
                is_owner: boolean;
            }[];
            /**
             * @description Type of the accessor
             * @example user
             * @enum {string}
             */
            type: "user" | "team";
        };
        FolderPermissionListResponseItem: {
            /** @description List of Permissions */
            data: components["schemas"]["BasePermissionsResponseSchema"][];
            pagination: components["schemas"]["Pagination"] & {
                /** @example https://api.cmp.optimizely.com/v3/folders/c00bfb60703f11ef805602420ac8001b/permissions?offset=10&page_size=10 */
                next?: string | null;
            };
        };
        FolderPermissionUpdateRequest: {
            /**
             * @description Set the level of access the accessor has to the folder
             * @example view
             * @enum {string}
             */
            access_type?: "view" | "edit";
            /**
             * @description Used to set an accessor as owner. This field can only be set to true as an owner cannot be removed without another accessor being set as the new owner. A team cannot be an owner.
             * @example true
             */
            is_owner?: boolean;
        };
        FolderResponse: {
            /**
             * Format: date-time
             * @description Date and time on which the folder was created,  in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Unique identifier of the folder
             * @example 5d7f910551b00a722e0418830cee6632
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the assets inside the folder
                 * @example https://api.cmp.optimizely.com/v3/assets?folder_id=5d7f910551b00a722e0418830cee6632&include_subfolder_assets=false
                 */
                assets: string;
                /**
                 * @description URL of the children list of the folder
                 * @example https://api.cmp.optimizely.com/v3/folders?parent_folder_id=5d7f910551b00a722e0418830cee6632
                 */
                child_folders: string;
                /**
                 * @description URL of the parent of the folder
                 * @example https://api.cmp.optimizely.com/v3/folders/1d9d8aeca10811ebbc640242ac12001b
                 */
                parent_folder: string | null;
                /**
                 * @description URL of the folder
                 * @example https://api.cmp.optimizely.com/v3/folders/5d7f910551b00a722e0418830cee6632
                 */
                self: string;
            };
            /**
             * Format: date-time
             * @description Date and time of the most recent modification of the folder, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            modified_at: string;
            /**
             * @description Name of the folder
             * @example icons
             */
            name: string;
            /**
             * @description ID of the parent folder
             * @example 1d9d8aeca10811ebbc640242ac12001b
             */
            parent_folder_id: string | null;
            /**
             * @description Folder location
             * @example /images/icons
             */
            path: string;
        };
        FolderUpdateRequest: {
            /**
             * @description Name of the folder
             * @example Optimizely Digital Asset Management
             */
            name?: string;
            /**
             * @description ID of the parent folder. If you pass null, the folder will be moved to the root.
             * @example 5e6f910551b00a722e0418830cee6630
             */
            parent_folder_id?: string | null;
        } | unknown | unknown;
        FormField: {
            /**
             * @description Type of the custom field when the form field is a custom field
             * @example multichoice
             */
            custom_field_type: string | null;
            /**
             * @description Type of template form field
             * @example custom_field
             */
            field_type: string;
            /**
             * @description Helper text for form field
             * @example This is a helper text
             */
            help: string;
            /**
             * @description Identifier of the form field
             * @example SingleChoice1
             */
            identifier: string;
            /**
             * @description Indicates whether the form field is required
             * @example false
             */
            is_required: boolean;
            /** @description List of the template logic rules */
            logic_rules: components["schemas"]["LogicRule"][];
            /**
             * @description Sort order of the form field
             * @example 0
             */
            sort_order: number;
        };
        GenericNumberTypeSettingsFieldCreatePayload: components["schemas"]["BaseSettingsFieldCreatePayload"] & {
            /**
             * @description Decimal place of the numerical field. This must be greather than 0. If this field is empty, 2 is set as the default.
             * @default 2
             * @example 2
             */
            decimal_places: number;
            /**
             * @description Whether the numerical field has a thousand separator
             * @default false
             * @example true
             */
            has_thousand_separator: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "percentage_number" | "simple_number";
        };
        GenericNumberTypeSettingsFieldResponse: components["schemas"]["BaseSettingsFieldsResponse"] & {
            /**
             * @description Decimal place of the numerical field
             * @example 2
             */
            decimal_places: number;
            /**
             * @description Whether the numerical field has a thousand separator
             * @example true
             */
            has_thousand_separator: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "percentage_number" | "simple_number";
        };
        GenericNumberTypeSettingsFieldUpdatePayload: components["schemas"]["BaseSettingsFieldUpdatePayload"] & {
            /**
             * @description Decimal place of the numerical field. This must be greather than 0. If this field is empty, 2 is set as the default.
             * @default 2
             * @example 2
             */
            decimal_places: number;
            /**
             * @description Whether the numerical field has a thousand separator
             * @default false
             * @example true
             */
            has_thousand_separator: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "percentage_number" | "simple_number";
        };
        /** @description Choice for field types other than `label` */
        GenericSettingsFieldChoiceCreatePayload: {
            /**
             * @description Name of the choice
             * @example Choice 1
             */
            name: string;
        };
        /** @description Choice for field types other than `label` */
        GenericSettingsFieldChoiceUpdatePayload: {
            /**
             * @description Name of the choice
             * @example Choice 1
             */
            name: string;
        };
        HTTPException: {
            /** @description Error message details. */
            detail?: string;
        };
        HTTPValidationError: {
            /** @description Details of the error */
            detail?: components["schemas"]["ValidationError"][];
        };
        Instruction: {
            /**
             * @description Description of the template instruction
             * @example Test instruction description
             */
            description: string;
            /**
             * @description Id of the template instruction
             * @example 9ec506f480dbc2d5f6919168ae448c40
             */
            id: string;
        };
        JSONFieldValueModel: {
            /** @description Value of the field value */
            json_value: Record<string, never> | unknown[];
            /** @description Order index of the field value */
            order_index?: number;
        };
        KeyedPreviewCompletedModel: {
            /**
             * Format: uri
             * @description URL for completed preview.
             */
            completed: string;
            /**
             * @description If not specified, the `locale` will be infered from the `Content-Language` header from the `completed` URL.
             *      In case of multiple locales, only the first one will be considered and it will be converted to snake case.
             *      For example, if the `Content-Language` header from the `completed` URL returns `en-US, de, bn`, it will be inferred as `en_US`.
             *      If the `locale` is not specified and there is no `Content-Language` header, this will default to `en_US`.
             * @default en_US
             */
            locale: string;
            /**
             * @description If not specified, the `mimeType` will be inferred from the `Content-Type` header from the `completed` URL.
             *      If the `mimeType` is not specified and there is no `Content-Type` header, this will default to `text/html`.
             * @default text/html
             */
            mimeType: string;
            /**
             * @description Human readable preview name to distinguish between generated preview outcomes.
             * @default
             */
            name: string;
            /**
             * @description Priority of a preview. For multiple previews, this will be used to sort the previews.
             * @default 0
             */
            orderIndex: number;
        };
        KeyedPreviewErrorModel: {
            /**
             * Format: uri
             * @description URL for error preview.
             */
            error: string;
            /**
             * @description If not specified, the `locale` will be infered from the `Content-Language` header from the `error` URL.
             *      In case of multiple locales, only the first one will be considered and it will be converted to snake case.
             *      For example, if the `Content-Language` header from the `completed` URL returns `en-US, de, bn`, it will be inferred as `en_US`.
             *      If the `locale` is not specified and there is no `Content-Language` header, this will default to `en_US`.
             * @default en_US
             */
            locale: string;
            /**
             * @description If not specified, the `mimeType` will be inferred from the `Content-Type` header from the `error` URL.
             *      If the `mimeType` is not specified and there is no `Content-Type` header, this will default to `text/html`.
             * @default text/html
             */
            mimeType: string;
            /**
             * @description Human readable preview name to distinguish between generated preview outcomes.
             * @default
             */
            name: string;
            /**
             * @description Priority of a preview. For multiple previews, this will be used to sort the previews.
             * @default 0
             */
            orderIndex: number;
        };
        LabelAndDropdownTypeSettingsFieldUpdatePayload: components["schemas"]["BaseSettingsFieldUpdatePayload"] & {
            /** @description Allow users to select multiple values */
            is_multi_select?: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "dropdown" | "label";
        };
        LabelGroup: {
            /**
             * @description Unique identifier for the label group
             * @example 693c3b28b10a22800e1936fa53324ac1
             */
            id: string;
            /**
             * @description Name of the label group
             * @example Content Pillar
             */
            name: string;
            /**
             * @description Source organization type where the label group is listed.
             *     `current` means the label group is listed in the organization with which the Open API App is associated.
             *     `related` means the label group is listed in one of the organizations related to the `current` organization.
             * @enum {string}
             */
            source_org_type: "current" | "related";
            /** @description Values of the label group */
            values: components["schemas"]["LabelGroupValue"][];
        };
        /** @description Values of a label group */
        LabelGroupValue: {
            /**
             * @description Unique identifier for the label group value
             * @example 1693bd22eac001bae8f33eab29ee5bc2
             */
            id: string;
            /**
             * @description Name of the label group value
             * @example Pillar 1
             */
            name: string;
        };
        LabelOnlyCustomField: {
            /**
             * @description Name of the custom field
             * @example A Custom Field
             */
            label: string;
        };
        LabelTypeSettingsFieldChoiceCreatePayload: components["schemas"]["GenericSettingsFieldChoiceCreatePayload"] & {
            /**
             * @description Code of choice color. Colors are limited to the following:
             *     - Electric Blue (#4ECFD5)
             *     - Golden Rod (#FFC700)
             *     - Salmon Pink (#FF98A7)
             *     - Blue Violet (#702BD5)
             *     - Spring Green (#5FEEAD)
             *     - Lavender Blue (#D6C4F2)
             *     - Sky (#5F9FF6)
             *     - Persian blue (#CB5DEB)
             *     - Cool Gray (#9694B3)
             * @example #FFC700
             */
            color: string;
        };
        /** @description Choice for field type of `label` */
        LabelTypeSettingsFieldChoiceUpdatePayload: {
            /**
             * @description Code of choice color. Colors are limited to the following:
             *     - Electric Blue (#4ECFD5)
             *     - Golden Rod (#FFC700)
             *     - Salmon Pink (#FF98A7)
             *     - Blue Violet (#702BD5)
             *     - Spring Green (#5FEEAD)
             *     - Lavender Blue (#D6C4F2)
             *     - Sky (#5F9FF6)
             *     - Persian blue (#CB5DEB)
             *     - Cool Gray (#9694B3)
             * @example #FFC700
             */
            color?: string;
            /**
             * @description Name of the choice
             * @example Choice 1
             */
            name?: string;
        };
        LabelTypeSettingsFieldCreatePayload: components["schemas"]["BaseSettingsFieldCreatePayload"] & {
            /** @description Choices of the field */
            choices: {
                /**
                 * @description Code of choice color. Colors are limited to the following:
                 *     - Electric Blue (#4ECFD5)
                 *     - Golden Rod (#FFC700)
                 *     - Salmon Pink (#FF98A7)
                 *     - Blue Violet (#702BD5)
                 *     - Spring Green (#5FEEAD)
                 *     - Lavender Blue (#D6C4F2)
                 *     - Sky (#5F9FF6)
                 *     - Persian blue (#CB5DEB)
                 *     - Cool Gray (#9694B3)
                 * @example #FFC700
                 */
                color: string;
                /**
                 * @description Name of the choice
                 * @example Choice 1
                 */
                name: string;
            }[];
            /** @description Allow users to select multiple values */
            is_multi_select: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "label";
        };
        LabelTypeSettingsFieldResponse: components["schemas"]["BaseSettingsFieldsResponse"] & {
            /** @description Choices of the field */
            choices: {
                /**
                 * @description Code of choice color. Colors are limited to the following:
                 *     - Electric Blue (#4ECFD5)
                 *     - Golden Rod (#FFC700)
                 *     - Salmon Pink (#FF98A7)
                 *     - Blue Violet (#702BD5)
                 *     - Spring Green (#5FEEAD)
                 *     - Lavender Blue (#D6C4F2)
                 *     - Sky (#5F9FF6)
                 *     - Persian blue (#CB5DEB)
                 *     - Cool Gray (#9694B3)
                 * @example #4ECFD5
                 */
                color: string | null;
                /**
                 * @description Identifier of the choice
                 * @example 6ceee2f4fa3411ecb37802420ac8001b
                 */
                id: string;
                /**
                 * @description Name of the choice
                 * @example Choice 1
                 */
                name: string;
            }[];
            /** @description Allow users to select multiple values */
            is_multi_select: boolean;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "label";
        };
        LibraryArticle: {
            /** @description List of authors of the article */
            authors: components["schemas"]["ArticleAuthorResponse"][];
            /**
             * Format: date-time
             * @description Date and time on which the article was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * Format: date-time
             * @description Date and time for the expiration of the article, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            expires_at: string | null;
            /**
             * @description Location of the folder containing the article in the library
             * @example /all assets/
             */
            file_location: string;
            /**
             * @description ID of the folder containing the article in the library
             * @example 6bb8db20a5b611ebae319b7c541b1a5a
             */
            folder_id: string | null;
            /**
             * @description Unique identifier of the source article in case of copied or translated articles
             * @example 7949c82ce3ae11eba378dbacdf18f20c
             */
            group_id: string | null;
            /**
             * @description Content of the article
             * @example <p>Article text</p>
             */
            html_body: string;
            /**
             * @description Unique identifier of the article
             * @example 5d7f910551b00a722e0418830cee6631
             */
            id: string;
            /** @description Featured images of the article */
            images: components["schemas"]["FeaturedImageResponse"][];
            /**
             * @description Whether the article is archived or not
             * @example false
             */
            is_archived: boolean;
            /** @description Labels associated with the article */
            labels: components["schemas"]["ResourceLabelResponse"][];
            /**
             * @description Language code depicting the language of the article
             * @example eng
             */
            lang_code: string | null;
            /**
             * @description Meta description of the article
             * @example Article text
             */
            meta_description: string | null;
            /**
             * @description Meta Keywords of the article
             * @example [
             *       "key",
             *       "word",
             *       "keyword"
             *     ]
             */
            meta_keywords: string[];
            /**
             * @description Meta title of the article
             * @example 3 Ways Influencer Marketing Will Further Mature in 2020
             */
            meta_title: string | null;
            /**
             * @description Meta URL of the article
             * @example /
             */
            meta_url: string | null;
            /**
             * Format: date-time
             * @description Date and time of the most recent modification of the article, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            modified_at: string;
            /**
             * @description ID of the article owner
             * @example 5108c3a9becac35915111191
             */
            owner_id: string;
            /**
             * @description ID of owner organization of the article
             * @example 5108c3a9becac35915111191
             */
            owner_organization_id: string;
            /**
             * @description Unique tracking pixel of the article
             * @example https://pixel.cmp.optimizely.com/px.gif?key=YXJ0aWNsZT0zZmJjN2Y5NmUzYzcxMWViOGVmNTAyNDJhYzEyMDAxOA==
             */
            pixel_key: string;
            /**
             * @description Link to the vendor article for Marketplace article or Task article ID for created article
             * @example https://source-vendor.com/sample-article
             */
            source_article: string | null;
            /**
             * @description Name of the article's source
             * @example Optimizely CMP
             */
            source_name: string | null;
            /** @description Tags for the article */
            tags: components["schemas"]["Tag"][];
            /**
             * @description URL of the thumbnail of the article
             * @example http://images.cmp.optimizely.com/Zz0xODQ3NDU3Y2Y2YmYzOTlmNjkyOTgyZDY3Y2I3YWM2OA==
             */
            thumbnail_url: string | null;
            /**
             * @description Title of the article
             * @example 3 Ways Influencer Marketing Will Further Mature in 2020
             */
            title: string;
            /**
             * @description Link to the published article
             * @example https://example.com/sample-article
             */
            url: string | null;
            /**
             * @description Version id of article
             * @example 5d7f910551b00a722e0418830cee6631
             */
            version_id: string;
            /**
             * @description Version number of article
             * @example 2
             */
            version_number: number;
        };
        LibraryAssetCreateRequest: {
            /**
             * @description ID of the folder where the asset should be added
             * @example 5d7f910551b00a722e0418830cee6631
             */
            folder_id?: string | null;
            /** @description Unique identifier of the file upload session. This is the `upload_meta_fields.key` field retrieved from the `/v3/upload-url` endpoint. */
            key: string;
            /**
             * @description Title of the asset
             * @example sample_image.png
             */
            title: string;
        };
        LibraryAssetDefaultValue: {
            /** @description Unique identifier of the asset */
            asset_guid: string;
            asset_type: components["schemas"]["LibraryAssetType"];
        };
        LibraryAssetFieldDefinition: {
            allowed_types: components["schemas"]["LibraryAssetType"][];
            core: components["schemas"]["CoreFieldDef"];
            /** @description Default values for the field */
            default_values?: components["schemas"]["LibraryAssetDefaultValue"][];
        };
        LibraryAssetFieldValueModel: {
            /** @description Asset GUID of the field value */
            asset_guid: string;
            asset_type: components["schemas"]["LibraryAssetType"];
            /** @description Meta links */
            links?: {
                [key: string]: string;
            };
            /** @description Order index of the field value */
            order_index?: number;
        };
        /**
         * @description An enumeration.
         * @enum {string}
         */
        LibraryAssetType: "article" | "image" | "video" | "raw_file" | "structured_content";
        LibraryAssetVersionResponse: {
            /**
             * @description Unique identifier of the asset
             * @example 5d7f910551b00a722e0418830cee6631
             */
            asset_id: string;
            /** @description Content of the version */
            content: {
                /**
                 * @description Type of the content
                 * @example url
                 * @enum {string}
                 */
                type: "url";
                /**
                 * @description Content of the version. It is the download URL for image, raw file, and video.
                 * @example http://images.cmp.optimizely.com/Zz0xODQ3NDU3Y2Y2YmYzOTlmNjkyOTgyZDY3Y2I3YWM2OA==S
                 */
                value: string;
            };
            /**
             * Format: date-time
             * @description Date and time on which the version was created,  in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the asset that the version is associated with
                 * @example https://api.cmp.optimizely.com/v3/images/5d7f910551b00a722e0418830cee6631
                 */
                asset: string;
            };
            /**
             * @description MIME type of the version
             * @example image/png
             */
            mime_type: string;
            /**
             * @description Title of the asset
             * @example sample_image.png
             */
            title: string;
            /**
             * @description Type of the version
             * @example image
             * @enum {string}
             */
            type: "image" | "video" | "raw_file";
            /**
             * @description The serial number of the version
             * @example 2
             */
            version_number: number;
        };
        LibraryImage: {
            /**
             * @description Alternative text for image
             * @example image title
             */
            alt_text: string | null;
            /**
             * @description Attribution of Image
             * @example Nature
             */
            attribution_text: string | null;
            /**
             * Format: date-time
             * @description Date and time on which the image was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Description of the image
             * @example Very important image
             */
            description: string | null;
            /**
             * Format: date-time
             * @description Date and time for the expiration of the image, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            expires_at: string | null;
            /**
             * @description File extension of the image file
             * @example jpg
             */
            file_extension: string | null;
            /**
             * @description Location of the folder containing the image in the library
             * @example /images/
             */
            file_location: string;
            /**
             * @description Size of the image in bytes
             * @example 11244
             */
            file_size: number;
            /** @description Focal point of the image */
            focal_point: {
                /**
                 * @description X coordinate of the focal point
                 * @example 35
                 */
                x: number;
                /**
                 * @description Y coordinate of the focal point
                 * @example 75
                 */
                y: number;
            } | null;
            /**
             * @description ID of the folder containing the image in the library
             * @example 6bb8db20a5b611ebae319b7c541b1a5a
             */
            folder_id: string | null;
            /**
             * @description Unique identifier of the image
             * @example 5d7f910551b00a722e0418830cee6632
             */
            id: string;
            /** @description Width and height of the image in pixels */
            image_resolution: {
                /**
                 * @description Height of the image in pixels
                 * @example 800
                 */
                height: number;
                /**
                 * @description Width of the image in pixels
                 * @example 700
                 */
                width: number;
            };
            /**
             * @description Whether the image is archived or not
             * @example true
             */
            is_archived: boolean;
            /**
             * @description Whether the image URL is public or not
             * @example true
             */
            is_public: boolean;
            /** @description Labels associated with the image */
            labels: components["schemas"]["ResourceLabelResponse"][];
            /**
             * @description MIME type of the image
             * @example image/jpeg
             */
            mime_type: string;
            /**
             * Format: date-time
             * @description Date and time of the most recent modification of the image, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            modified_at: string;
            /**
             * @description ID of the image owner
             * @example 5108c3a9becac35915111191
             */
            owner_id: string;
            /**
             * @description ID of owner organization of the image
             * @example 5108c3a9becac35915111191
             */
            owner_organization_id: string;
            /** @description Tags for the image */
            tags: components["schemas"]["Tag"][];
            /**
             * @description URL of the thumbnail of the image
             * @example http://images.cmp.optimizely.com/Zz0xODQ3NDU3Y2Y2YmYzOTlmNjkyOTgyZDY3Y2I3YWM2OA==
             */
            thumbnail_url: string | null;
            /**
             * @description Title of the image
             * @example cat.jpeg
             */
            title: string;
            /**
             * @description Download the URL of the image
             * @example http://images.cmp.optimizely.com/Zz0xODQ3NDU3Y2Y2YmYzOTlmNjkyOTgyZDY3Y2I3YWM2OA==
             */
            url: string;
            /**
             * @description Version id of image
             * @example 5d7f910551b00a722e0418830cee6631
             */
            version_id: string;
            /**
             * @description Version number of image
             * @example 2
             */
            version_number: number;
        };
        LibraryImageUpdateRequest: {
            /**
             * @description Alt text of the image. Set an empty string to clear the alt text.
             * @example Sample alt text
             */
            alt_text?: string;
            /**
             * @description Attribution text of the image. Set an empty string to clear the attribution text.
             * @example Sample attribution text
             */
            attribution_text?: string;
            /**
             * @description Description of the image. Set an empty string to clear the description.
             * @example Image of a cute cat
             */
            description?: string;
            /**
             * Format: date-time
             * @description Date and time when the image expires, in ISO 8601 UTC format
             * @example 2019-10-07T13:15:30Z
             */
            expires_at?: string | null;
            /**
             * @description ID of the folder containing the image in the library
             * @example 6bb8db20a5b611ebae319b7c541b1a5a
             */
            folder_id?: string | null;
            /**
             * @description Whether the image should be archived
             * @example true
             */
            is_archived?: boolean;
            /**
             * @description Whether the image URL should be public
             * @example true
             */
            is_public?: boolean;
            /**
             * @deprecated
             * @description Please use [PUT /assets/{asset_id}/fields](https://docs.developers.optimizely.com/content-marketing-platform/reference/put_assets-asset-id-fields) or  [PUT /assets/{asset_id}/fields/{field_id}](https://docs.developers.optimizely.com/content-marketing-platform/reference/put_assets-asset-id-fields-field-id)  API to update label type asset fields
             */
            labels?: components["schemas"]["ResourceLabelRequest"][];
            /**
             * @description Tags of the image. Provided tags will replace the existing ones. Set an empty array to clear the existing tags.
             * @example [
             *       "dhaka",
             *       "new york",
             *       "london"
             *     ]
             */
            tags?: string[];
            /**
             * @description Title of the image
             * @example cat.jpeg
             */
            title?: string;
        };
        LibraryRawFile: {
            /**
             * @description Attribution of raw file
             * @example Text
             */
            attribution_text: string | null;
            /**
             * Format: date-time
             * @description Date and time on which the raw file was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Description of the raw file
             * @example Very important file
             */
            description: string | null;
            /**
             * Format: date-time
             * @description Date and time for the expiration of the raw file, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            expires_at: string | null;
            /**
             * @description File extension of the raw file
             * @example zip
             */
            file_extension: string | null;
            /**
             * @description Location of the folder containing the raw file in the library
             * @example /all files/
             */
            file_location: string;
            /**
             * @description Size of the raw file in bytes
             * @example 11244
             */
            file_size: number;
            /**
             * @description ID of the folder containing the raw file in the library
             * @example 6bb8db20a5b611ebae319b7c541b1a5a
             */
            folder_id: string | null;
            /**
             * @description Unique identifier of the raw file
             * @example 5d7f910551b00a722e0418830cee6634
             */
            id: string;
            /**
             * @description Whether the raw file is archived or not
             * @example false
             */
            is_archived: boolean;
            /**
             * @description Whether the raw file URL is public or not
             * @example true
             */
            is_public: boolean;
            /** @description Labels associated to the raw file */
            labels: components["schemas"]["ResourceLabelResponse"][];
            /**
             * @description MIME type of the raw file
             * @example application/zip
             */
            mime_type: string;
            /**
             * Format: date-time
             * @description Date and time of the most recent modification of the raw file, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            modified_at: string;
            /**
             * @description ID of the raw file owner
             * @example 5108c3a9becac35915111191
             */
            owner_id: string;
            /**
             * @description ID of owner organization of the raw file
             * @example 5108c3a9becac35915111191
             */
            owner_organization_id: string;
            /** @description Tags for the raw file */
            tags: components["schemas"]["Tag"][];
            /**
             * @description URL of the thumbnail of the raw file
             * @example http://images.cmp.optimizely.com/Zz0xODQ3NDU3Y2Y2YmYzOTlmNjkyOTgyZDY3Y2I3YWM2OA==
             */
            thumbnail_url: string | null;
            /**
             * @description Title of the raw file
             * @example documents.zip
             */
            title: string;
            /**
             * @description Download the URL of the raw file
             * @example https://files.cmp.optimizely.com/171451644651701b96f1122009f026bc
             */
            url: string;
            /**
             * @description Version id of raw file
             * @example 5d7f910551b00a722e0418830cee6631
             */
            version_id: string;
            /**
             * @description Version number of raw file
             * @example 2
             */
            version_number: number;
        };
        LibraryRawFileUpdateRequest: {
            /**
             * @description Attribution text of the raw file. Set an empty string to clear the attribution text.
             * @example Sample attribution text
             */
            attribution_text?: string;
            /**
             * @description Description of the raw file. Set an empty string to clear the description.
             * @example A zip file that contains some documents
             */
            description?: string;
            /**
             * Format: date-time
             * @description Date and time when the raw file expires, in ISO 8601 UTC format
             * @example 2019-10-07T13:15:30Z
             */
            expires_at?: string | null;
            /**
             * @description ID of the folder containing the raw file in the library
             * @example 6bb8db20a5b611ebae319b7c541b1a5a
             */
            folder_id?: string | null;
            /**
             * @description Whether the raw file should be archived
             * @example true
             */
            is_archived?: boolean;
            /**
             * @description Whether the raw file URL should be public
             * @example true
             */
            is_public?: boolean;
            /**
             * @deprecated
             * @description Please use [PUT /assets/{asset_id}/fields](https://docs.developers.optimizely.com/content-marketing-platform/reference/put_assets-asset-id-fields) or  [PUT /assets/{asset_id}/fields/{field_id}](https://docs.developers.optimizely.com/content-marketing-platform/reference/put_assets-asset-id-fields-field-id)  API to update label type asset fields
             */
            labels?: components["schemas"]["ResourceLabelRequest"][];
            /**
             * @description Tags of the raw file. Provided tags will replace the existing ones. Set an empty array to clear the existing tags.
             * @example [
             *       "dhaka",
             *       "new york",
             *       "london"
             *     ]
             */
            tags?: string[];
            /**
             * @description Title of the raw file
             * @example documents.zip
             */
            title?: string;
        };
        LibraryStructuredContent: {
            content_body?: components["schemas"]["ContentDetailsModel"];
            /**
             * Format: date-time
             * @description Date and time on which the structured content was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Location of the folder containing the structured content in the library
             * @example /all files/
             */
            file_location: string;
            /**
             * @description ID of the folder containing the structured content in the library
             * @example 6bb8db20a5b611ebae319b7c541b1a5a
             */
            folder_id: string | null;
            /**
             * @description Unique identifier of the structured content
             * @example a24b834427a043ab8caed1ded4606f7d
             */
            id: string;
            /**
             * @description Whether the structured content is archived or not
             * @example true
             */
            is_archived: boolean;
            /** @description Labels associated with the structured content */
            labels: components["schemas"]["ResourceLabelResponse"][];
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the structured content
                 * @example https://api.cmp.optimizely.com/v3/structured-contents/a24b834427a043ab8caed1ded4606f7d
                 */
                self: string;
            };
            /**
             * Format: date-time
             * @description Date and time of the most recent modification of the structured content, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            modified_at: string;
            /**
             * @description ID of owner organization of the structured content
             * @example 5108c3a9becac35915111191
             */
            owner_organization_id: string;
            /**
             * @description Title of the structured content
             * @example Some webpage
             */
            title: string;
        };
        LibraryStructuredContentCreateRequest: {
            /** @description The structured content body */
            content_body: {
                /**
                 * Contenttypeguid
                 * @description The guid of the content's content type
                 */
                content_type_guid: string;
                /**
                 * Expired
                 * @description Expired status of the content
                 */
                expired?: boolean;
                /**
                 * Expirydatetime
                 * Format: date-time
                 * @description Date and time on which the content will expire, in ISO 8601 UTC format
                 */
                expiry_datetime?: string | null;
                /** fields */
                fields?: components["schemas"]["StructuredContentBody"];
                /**
                 * Rootcontent
                 * @description This is true when the content is not embedded in another content through a content type reference field
                 */
                root_content: boolean;
                /**
                 * Source
                 * @description Source of the content
                 */
                source?: string | null;
                /**
                 * Sourceid
                 * @description Source id of the content
                 */
                source_id?: string | null;
                /**
                 * Sourcemetadata
                 * @description Source metadata of the content
                 */
                source_metadata?: string | null;
                /**
                 * @description Title of the structured content
                 * @example Some webpage
                 */
                title?: string;
            };
            /**
             * @description ID of the folder containing the structured content in the library
             * @example 6bb8db20a5b611ebae319b7c541b1a5a
             */
            folder_id?: string | null;
            /** @description Title of the structured content asset */
            title?: string;
        };
        LibraryStructuredContentUpdateRequest: {
            /** @description The structured content body */
            content_body?: {
                /**
                 * Expired
                 * @description Expired status of the content
                 */
                expired?: boolean | null;
                /**
                 * Expirydatetime
                 * Format: date-time
                 * @description Date and time on which the content will expire, in ISO 8601 UTC format
                 */
                expiry_datetime?: string | null;
                /** Contentbody */
                fields?: components["schemas"]["StructuredContentFields"];
                /**
                 * Source
                 * @description Source of the content
                 */
                source?: string | null;
                /**
                 * Sourceid
                 * @description Source id of the content
                 */
                source_id?: string | null;
                /**
                 * Sourcemetadata
                 * @description Source metadata of the content
                 */
                source_metadata?: string | null;
                /**
                 * @description Title of the structured content
                 * @example Some webpage
                 */
                title?: string | null;
            };
            /**
             * @description Whether the structured content should be archived
             * @example true
             */
            is_archived?: boolean;
            /** @description Title of the structured content asset */
            title?: string | null;
        };
        LibraryVideo: {
            /**
             * @description Alternative text for video
             * @example video title
             */
            alt_text: string | null;
            /**
             * @description Attribution of Video
             * @example Sci-fi
             */
            attribution_text: string | null;
            /**
             * Format: date-time
             * @description Date and time on which the video was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Description of the video
             * @example Very important video
             */
            description: string | null;
            /**
             * Format: date-time
             * @description Date and time for the expiration of the video, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            expires_at: string | null;
            /**
             * @description File extension of the video file
             * @example mp4
             */
            file_extension: string | null;
            /**
             * @description Location of the folder containing the video in the library
             * @example /example/
             */
            file_location: string;
            /**
             * @description Size of the video in bytes
             * @example 11244
             */
            file_size: number;
            /**
             * @description ID of the folder containing the video in the library
             * @example 6bb8db20a5b611ebae319b7c541b1a5a
             */
            folder_id: string | null;
            /**
             * @description Unique identifier for the video
             * @example 5d7f910551b00a722e0418830cee6633
             */
            id: string;
            /**
             * @description Whether the video is archived or not
             * @example true
             */
            is_archived: boolean;
            /**
             * @description Whether the video URL is public
             * @example true
             */
            is_public: boolean;
            /** @description Labels associated with the video */
            labels: components["schemas"]["ResourceLabelResponse"][];
            /**
             * @description MIME type of the video
             * @example video/mp4
             */
            mime_type: string;
            /**
             * Format: date-time
             * @description Date and time of the most recent modification of the video, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            modified_at: string;
            /**
             * @description ID of the video owner
             * @example 5108c3a9becac35915111191
             */
            owner_id: string;
            /**
             * @description ID of owner organization of the video
             * @example 5108c3a9becac35915111191
             */
            owner_organization_id: string;
            /** @description Tags for the video */
            tags: components["schemas"]["Tag"][];
            /**
             * @description URL of the thumbnail of the video
             * @example http://images.cmp.optimizely.com/Zz0xODQ3NDU3Y2Y2YmYzOTlmNjkyOTgyZDY3Y2I3YWM2OA==
             */
            thumbnail_url: string | null;
            /**
             * @description Title of the video
             * @example product.mp4
             */
            title: string;
            /**
             * @description Download the URL of the video
             * @example https://videos.cmp.optimizely.com/03a747babe81ceb55763fa085bqa20dc
             */
            url: string;
            /**
             * @description Version id of video
             * @example 5d7f910551b00a722e0418830cee6631
             */
            version_id: string;
            /**
             * @description Version number of video
             * @example 2
             */
            version_number: number;
        };
        LibraryVideoUpdateRequest: {
            /**
             * @description Alt text of the video. Set an empty string to clear the alt text.
             * @example Sample alt text
             */
            alt_text?: string;
            /**
             * @description Attribution text of the video. Set an empty string to clear the attribution text.
             * @example Sample attribution text
             */
            attribution_text?: string;
            /**
             * @description Description of the video. Set an empty string to clear the description.
             * @example A short video on product demo
             */
            description?: string;
            /**
             * Format: date-time
             * @description Date and time when the video expires, in ISO 8601 UTC format
             * @example 2019-10-07T13:15:30Z
             */
            expires_at?: string | null;
            /**
             * @description ID of the folder containing the video in the library
             * @example 6bb8db20a5b611ebae319b7c541b1a5a
             */
            folder_id?: string | null;
            /**
             * @description Whether the video should be archived
             * @example true
             */
            is_archived?: boolean;
            /**
             * @description Whether the video URL should be public
             * @example true
             */
            is_public?: boolean;
            /**
             * @deprecated
             * @description Please use [PUT /assets/{asset_id}/fields](https://docs.developers.optimizely.com/content-marketing-platform/reference/put_assets-asset-id-fields) or  [PUT /assets/{asset_id}/fields/{field_id}](https://docs.developers.optimizely.com/content-marketing-platform/reference/put_assets-asset-id-fields-field-id)  API to update label type asset fields
             */
            labels?: components["schemas"]["ResourceLabelRequest"][];
            /**
             * @description Tags of the video. Provided tags will replace the existing ones. Set an empty array to clear the existing tags.
             * @example [
             *       "dhaka",
             *       "new york",
             *       "london"
             *     ]
             */
            tags?: string[];
            /**
             * @description Title of the video
             * @example product.mp4
             */
            title?: string;
        };
        LocalizedFieldValues: {
            field_values: (components["schemas"]["NumberFieldValueModel"] | components["schemas"]["BooleanFieldValueModel"] | components["schemas"]["DatetimeFieldValueModel"] | components["schemas"]["TextFieldValueModel"] | components["schemas"]["RichTextFieldValueModel"] | components["schemas"]["LibraryAssetFieldValueModel"] | components["schemas"]["ContentFieldValueExpandedModel"] | components["schemas"]["ContentFieldValueModel"] | components["schemas"]["URLFieldValueModel"] | components["schemas"]["JSONFieldValueModel"] | components["schemas"]["ChoiceFieldValueModel"] | components["schemas"]["LocationFieldValueModel"])[];
            locale: string;
        };
        LocalizedFieldValuesWithEmbeddedContent: {
            field_values: (components["schemas"]["NumberFieldValueModel"] | components["schemas"]["BooleanFieldValueModel"] | components["schemas"]["DatetimeFieldValueModel"] | components["schemas"]["TextFieldValueModel"] | components["schemas"]["RichTextFieldValueModel"] | components["schemas"]["LibraryAssetFieldValueModel"] | components["schemas"]["ContentFieldValueWithEmbeddedModel"] | components["schemas"]["ContentFieldValueModel"] | components["schemas"]["URLFieldValueModel"] | components["schemas"]["JSONFieldValueModel"] | components["schemas"]["ChoiceFieldValueModel"] | components["schemas"]["LocationFieldValueModel"])[];
            locale: string;
        };
        LocationDefaultValue: {
            /** @description Latitude of the location */
            latitude: number;
            /** @description Longitude of the location */
            longitude: number;
        };
        LocationFieldValueModel: {
            /** @description Longitude of the field value */
            latitude: number;
            /** @description Longitude of the field value */
            longitude: number;
            /** @description Order index of the field value */
            order_index?: number;
        };
        LogicRule: {
            /** @description Template logic rule action */
            action: components["schemas"]["LogicRuleAction"];
            /** @description Logic rule action associated condition */
            condition: components["schemas"]["LogicRuleCondition"];
        };
        /** @description Template logic rule action */
        LogicRuleAction: {
            target_field: components["schemas"]["LogicRuleTargetField"];
            /**
             * @description Type of logic rule action
             * @example jump_to
             */
            type: string;
            /**
             * @description Values to consider on logic rule action
             * @example []
             */
            values: string[];
        };
        /** @description Logic rule action associated condition */
        LogicRuleCondition: {
            /**
             * @description The operator for logic rule action
             * @example equal
             */
            operator: string;
            /**
             * @description List of labels or custom field choices
             * @example [
             *       "choice1"
             *     ]
             */
            values: string[];
        };
        /** @description Template logic rule action's target field */
        LogicRuleTargetField: {
            /**
             * @description Type of custom field if the target field is a custom field
             * @default null
             * @example null
             */
            custom_field_type: string | null;
            /**
             * @description Name of the target field identifier
             * @example end_of_form
             */
            identifier: string;
            /**
             * @description Type of the target field identifier
             * @example end_of_form
             */
            type: string;
        };
        MilestoneCreateRequest: {
            /**
             * @description Id of the campaign to be associated with the milestone
             * @example 63f1c2b675be4132854d2741
             */
            campaign_id?: string;
            /**
             * @description Description of the milestone. If provided, this should be between 1 and 250 characters.
             * @example Complete all tasks for the Q1 product launch
             */
            description?: string | null;
            /**
             * Format: date-time
             * @description Date and time on which the milestone will expire, in ISO 8601 UTC format
             * @example 2023-10-07T13:15:30Z
             */
            due_date: string;
            /**
             * @description Hex color code for the milestone label
             * @example #4ECFD5
             */
            hex_color: string;
            /**
             * @description List of tasks to be associated with the milestone
             * @example [
             *       {
             *         "id": "63f1c2b675be4132854d2741"
             *       },
             *       {
             *         "id": "63f1c2b675be4132854d2742"
             *       }
             *     ]
             */
            tasks?: {
                /** @description Task ID */
                id: string;
            }[];
            /**
             * @description Title of the milestone. This should be between 1 and 80 characters.
             * @example Q1 Product Launch
             */
            title: string;
        };
        MilestoneResponse: {
            /** @description Campaign associated with the milestone */
            campaign: {
                /**
                 * @description Unique identifier of the campaign
                 * @example 63f1c2b675be4132854d2742
                 */
                id: string;
            } | null;
            /**
             * @description Color of the milestone label
             * @example #4ECFD5
             */
            color: string;
            /**
             * @description Description of the milestone
             * @example This is a milestone
             */
            description: string | null;
            /**
             * Format: date-time
             * @description Date and time on which the milestone will expire, in ISO 8601 UTC format
             * @example 2023-10-07T13:15:30Z
             */
            due_date: string;
            /**
             * @description Unique identifier of the milestone
             * @example 63f1c2b675be4132854d2741
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the milestone
                 * @example https://api.cmp.optimizely.com/v3/milestones/63f1c2b675be4132854d2741
                 */
                self: string;
            };
            /**
             * @description Title of the milestone
             * @example Sample Milestone
             */
            title: string;
        };
        MilestoneUpdateRequest: {
            /**
             * @description Id of the campaign to be associated with the milestone
             * @example 63f1c2b675be4132854d2741
             */
            campaign_id?: string;
            /**
             * @description Description of the milestone. If provided, this should be between 1 and 250 characters.
             * @example Complete all tasks for the Q1 product launch
             */
            description?: string | null;
            /**
             * Format: date-time
             * @description Date and time on which the milestone will expire, in ISO 8601 UTC format
             * @example 2023-10-07T13:15:30Z
             */
            due_date?: string;
            /**
             * @description Hex color code for the milestone label
             * @example #4ECFD5
             */
            hex_color?: string;
            /**
             * @description List of tasks to be associated with the milestone. This can be an empty array, but in that case it will remove all task associations for that milestone.
             * @example [
             *       {
             *         "id": "63f1c2b675be4132854d2741"
             *       },
             *       {
             *         "id": "63f1c2b675be4132854d2742"
             *       }
             *     ]
             */
            tasks?: {
                /** @description Task ID */
                id: string;
            }[];
            /**
             * @description Title of the milestone. This should be between 1 and 80 characters.
             * @example Q1 Product Launch
             */
            title?: string;
        };
        MultiChoiceTypeFormFieldRequest: components["schemas"]["BaseFormFieldRequest"] & {
            /** @description Array of choice IDs */
            values?: string[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "checkbox" | "dropdown" | "label";
        };
        MultiChoiceTypeFormFieldResponse: components["schemas"]["BaseFormFieldResponse"] & {
            /** @description Array of choices */
            values?: {
                /**
                 * @description Unique identifier of the choice
                 * @example 9119a313057e401189407116fcd3
                 */
                id: string;
                /**
                 * @description Name of the choice
                 * @example Choice 1
                 */
                name: string;
            }[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "checkbox" | "dropdown" | "label";
        };
        MultiChoiceTypeObjectField: components["schemas"]["BaseObjectField"] & components["schemas"]["MultiChoiceTypeObjectFieldUpdatePayload"] & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "checkbox" | "label";
        };
        MultiChoiceTypeObjectFieldUpdatePayload: components["schemas"]["BaseObjectFieldUpdatePayload"] & {
            /** @description Array of choice IDs */
            values?: string[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "checkbox" | "label";
        };
        MultipleOptionCustomField: {
            /**
             * @description Custom field choices
             * @example [
             *       "choice1",
             *       "choice2"
             *     ]
             */
            choices: string[];
            /**
             * @description Name of the custom field
             * @example A custom field with choice
             */
            label: string;
        };
        NullValue: unknown;
        NumberCustomField: {
            /**
             * @description Value of decimal places
             * @example 1
             */
            decimal_places: number | null;
            /**
             * @description Indicator of thousand separator
             * @default false
             * @example false
             */
            has_thousand_separator: boolean;
            /**
             * @description Name of the custom field
             * @example A Number Custom Field
             */
            label: string;
        };
        NumberFieldDefinition: {
            core: components["schemas"]["CoreFieldDef"];
            /** @description Default values of the field */
            default_values?: number[];
            /** @description Maximum value of the field */
            max_value?: number;
            /** @description Minimum value of the field */
            min_value?: number;
        };
        NumberFieldValueModel: {
            /** @description Value of the field value */
            num_value: number;
            /** @description Order index of the field value */
            order_index?: number;
        };
        NumberTypeFormFieldRequest: components["schemas"]["BaseFormFieldRequest"] & {
            /** @description Array of numbers */
            values?: number[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "currency_number" | "percentage_number" | "simple_number";
        };
        NumberTypeFormFieldResponse: components["schemas"]["BaseFormFieldResponse"] & {
            /** @description Array of numeric strings */
            values?: string[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "currency_number" | "percentage_number" | "simple_number";
        };
        NumberTypeObjectField: components["schemas"]["BaseObjectField"] & components["schemas"]["NumberTypeObjectFieldUpdatePayload"] & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "currency_number" | "percentage_number" | "simple_number";
        };
        NumberTypeObjectFieldUpdatePayload: components["schemas"]["BaseObjectFieldUpdatePayload"] & {
            /** @description Accepts a single numeric value */
            values?: number[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "currency_number" | "percentage_number" | "simple_number";
        };
        ObjectFieldCreateRequest: components["schemas"]["StringTypeObjectField"] | components["schemas"]["MultiChoiceTypeObjectField"] | components["schemas"]["DropdownTypeObjectField"] | components["schemas"]["RadioButtonTypeObjectField"] | components["schemas"]["NumberTypeObjectField"] | components["schemas"]["DateTypeObjectField"] | components["schemas"]["AssetTypeObjectField"];
        ObjectFieldCreateResponse: components["schemas"]["FieldTypeCommon"] | components["schemas"]["FieldTypeLabel"] | components["schemas"]["FieldTypeDropdown"] | components["schemas"]["FieldTypeRadio"] | components["schemas"]["FieldTypeCheckbox"] | components["schemas"]["FieldTypeSimpleNumber"] | components["schemas"]["FieldTypePercentageNumber"] | components["schemas"]["FieldTypeCurrencyNumber"];
        /** @description Pagination related information */
        Pagination: {
            /**
             * @description URL to the next page
             * @example https://api.cmp.optimizely.com/<some-path-to-next-page>?offset=10&page_size=10
             */
            next: string | null;
            /**
             * @description URL to the previous page
             * @example null
             */
            previous: string | null;
        };
        PatchLocalizedFieldValuesWithEmbeddedContent: {
            field_values: (components["schemas"]["NumberFieldValueModel"] | components["schemas"]["BooleanFieldValueModel"] | components["schemas"]["DatetimeFieldValueModel"] | components["schemas"]["TextFieldValueModel"] | components["schemas"]["RichTextFieldValueModel"] | components["schemas"]["LibraryAssetFieldValueModel"] | components["schemas"]["ContentFieldValueWithEmbeddedModel"] | components["schemas"]["ContentFieldValueModel"] | components["schemas"]["URLFieldValueModel"] | components["schemas"]["JSONFieldValueModel"] | components["schemas"]["ChoiceFieldValueModel"] | components["schemas"]["LocationFieldValueModel"] | components["schemas"]["DeleteFieldValueModel"] | components["schemas"]["ContentFieldWithEmbeddedPatchValueModel"])[];
            /** @description The default value should be en_US */
            locale: string;
        };
        PercentageCustomField: {
            /**
             * @description Value of decimal places
             * @default null
             * @example 1
             */
            decimal_places: number | null;
            /**
             * @description Name of the custom field
             * @example A Percentage Custom Field
             */
            label: string;
        };
        PublishingEventAssets: {
            /**
             * @description Unique identifier of the asset
             * @example 5d7f910551b00a722e0418830cee2212
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of asset
                 * @example https://api.cmp.optimizely.com/v3/articles/5d7f910551b00a722e0418830cee2212
                 */
                self: string;
            };
            /** @description List of information related to publishing metadata for the asset */
            publishing_metadata: {
                /**
                 * @description Unique identifier of the publishing metadata
                 * @example 5e46745616s564s4564964
                 */
                id: string;
                /** @description Meta links */
                links: {
                    /**
                     * @description URL of the publishing metadata of the asset
                     * @example https://api.cmp.optimizely.com/v3/publishing-events/5d7f910551b00a722e0418830cee5534/assets/5d7f910551b00a722e0418830cee2212/publishing-metadata/5e46745616s564s4564964
                     */
                    self: string;
                };
            }[];
            /**
             * @description Type of the asset: `article`, `image`, `video`, `raw_file`, `structured_content`
             * @example article
             */
            type: string;
        };
        PublishingEventMetadataBulkCreateRequest: {
            /** @description List publishing metadata to be posted */
            data: components["schemas"]["PublishingEventMetadataCreateRequest"][];
        };
        PublishingEventMetadataBulkCreateResponse: {
            /** @description List of successfully posted publishing metadata */
            data: components["schemas"]["PublishingEventMetadataResponse"][];
            /** @description List of errors of assets for which publishing metadata failed to post */
            errors: {
                /**
                 * @description Unique identifier of the asset that failed to update its publishing metadata
                 * @example 4567m474974987479856456457
                 */
                asset_id: string;
                /**
                 * @description Custom error code
                 * @example canonical-link-error
                 * @enum {string}
                 */
                error_code: "canonical-link-error" | "unknown-asset" | "metadata-exists" | "duplicate-metadata" | "invalid-status" | "missing-public-url" | "public-url-not-allowed" | "domain-not-whitelisted" | "missing-publishing-destination-updated-at" | "publishing-destination-updated-at-not-allowed" | "missing-locale" | "locale-not-allowed";
                /**
                 * @description locale of the asset
                 * @example en
                 */
                locale: string | null;
                /**
                 * @description Description of the error
                 * @example A canonical URL already exists for the task article '5e46456144645674564456'
                 */
                message: string;
            }[];
        };
        PublishingEventMetadataCreateRequest: {
            /**
             * @description Unique identifier of the asset
             * @example 4567m474974987479856456457
             */
            asset_id: string;
            /**
             * @description The locale to which the asset is being published
             * @example en
             */
            locale?: string;
            /**
             * @description public url of asset
             * @example https://example.com/test
             */
            public_url?: string;
            /**
             * Format: date-time
             * @description Timestamp of when the publishing destination of the asset was updated
             * @example 2019-10-06T13:15:30Z
             */
            publishing_destination_updated_at?: string;
            /**
             * @description Publishing status of the asset
             * @example published
             * @enum {string}
             */
            status: "published" | "unpublished" | "synced" | "failed";
            /**
             * @description Any message patched by the integrator about the asset's status
             * @example This asset is in review
             */
            status_message?: string;
        };
        PublishingEventMetadataListResponse: {
            /** @description List of successfully posted publishing metadata */
            data?: components["schemas"]["PublishingEventMetadataResponse"][];
        };
        PublishingEventMetadataResponse: {
            /**
             * @description Unique identifier of asset
             * @example 4567m74974987479856456456
             */
            asset_id: string;
            /**
             * @description Type of the asset
             * @example article
             * @enum {string}
             */
            asset_type: "article" | "image" | "video" | "raw_file" | "structured_content";
            /**
             * @description Unique identifier of publishing metadata
             * @example 5ebcd5644967474414564
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the asset
                 * @example https://api.cmp.optimizely.com/v3/articles/4567m474974987479856456456
                 */
                asset: string;
                /**
                 * @description URL of the publishing event
                 * @example https://api.cmp.optimizely.com/v3/publishing-events/1d9d8aeca10811ebbc640242ac12001b
                 */
                publishing_event: string;
                /**
                 * @description URL of the publishing metadata
                 * @example https://api.cmp.optimizely.com/v3/publishing-events/1d9d8aeca10811ebbc640242ac12001b/assets/1d9d8aeca10811ebbc640242ac12003c/publishing-metadata/5ebcd5644967474414564
                 */
                self: string;
            };
            /**
             * @description The locale to which the asset is being published.
             * @example en
             */
            locale: string | null;
            /**
             * @description Public URL of the asset
             * @example https://example.com/test
             */
            public_url: string | null;
            /**
             * Format: date-time
             * @description Timestamp of when the publishing destination of the asset was updated
             * @example 2019-10-06T13:15:30Z
             */
            publishing_destination_updated_at: string | null;
            /**
             * @description Publishing status of the asset
             * @example published
             * @enum {string|null}
             */
            status: "published" | "unpublished" | "synced" | "failed" | null;
            /**
             * @description Any message patched by the integrator about the asset's status
             * @example This asset is in review
             */
            status_message: string | null;
        };
        PublishingEventResponse: {
            /** @description List of assets associated with the publishing event. */
            assets: components["schemas"]["PublishingEventAssets"][];
            /**
             * @description Unique identifier of the publishing event
             * @example 5d7f910551b00a722e0418830cee5534
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of list of publishing metadata
                 * @example https://api.cmp.optimizely.com/v3/publishing-events/5d7f910551b00a722e0418830cee5534/publishing-metadata
                 */
                publishing_metadata: string;
                /**
                 * @description URL of the publishing event
                 * @example https://api.cmp.optimizely.com/v3/publishing-events/5d7f910551b00a722e0418830cee2212
                 */
                self: string;
            };
        };
        RadioButtonTypeFormFieldRequest: components["schemas"]["BaseFormFieldRequest"] & {
            /** @description Array of choice IDs */
            values?: string[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "radio_button";
        };
        RadioButtonTypeFormFieldResponse: components["schemas"]["MultiChoiceTypeFormFieldResponse"] & {
            /** @description Array of choices */
            values?: {
                /**
                 * @description Unique identifier of the choice
                 * @example 9119a313057e401189407116fcd3
                 */
                id: string;
                /**
                 * @description Name of the choice
                 * @example Choice 1
                 */
                name: string;
            }[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "radio_button";
        };
        RadioButtonTypeObjectField: components["schemas"]["BaseObjectField"] & components["schemas"]["RadioButtonTypeObjectFieldUpdatePayload"] & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "radio_button";
        };
        RadioButtonTypeObjectFieldUpdatePayload: components["schemas"]["BaseObjectFieldUpdatePayload"] & {
            /** @description Accepts a single choice Id. */
            values?: string[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "radio_button";
        };
        RecursivePatchStructuredContentFields: {
            /** Contentbody */
            content_body?: components["schemas"]["StructuredContentFields"];
            /**
             * Expired
             * @description Expired status of the content
             */
            expired?: boolean;
            /**
             * Expirydatetime
             * Format: date-time
             * @description Date and time on which the content will expire, in ISO 8601 UTC format
             */
            expiry_datetime?: string | null;
            /**
             * Rootcontent
             * @description This is true when the content is not embedded in another content through a content type reference field
             */
            root_content?: boolean;
            /**
             * Source
             * @description Source of the content
             */
            source?: string | null;
            /**
             * Sourceid
             * @description Source id of the content
             */
            source_id?: string | null;
            /** Sourcemetadata */
            source_metadata?: string;
            /**
             * @description Title of the structured content
             * @example Some webpage
             */
            title?: string;
        };
        RecursiveStructuredContent: {
            /**
             * Contenttypeguid
             * @description The guid of the content's content type
             */
            content_type_guid: string;
            /**
             * Expired
             * @description Expired status of the content
             */
            expired?: boolean;
            /**
             * Expirydatetime
             * Format: date-time
             * @description Date and time on which the content will expire, in ISO 8601 UTC format
             */
            expiry_datetime?: string | null;
            /**
             * Rootcontent
             * @description This is true when the content is not embedded in another content through a content type reference field
             */
            root_content: boolean;
            /**
             * Source
             * @description Source of the content
             */
            source?: string | null;
            /**
             * Sourceid
             * @description Source id of the content
             */
            source_id?: string | null;
            /**
             * Sourcemetadata
             * @description Source metadata of the content
             */
            source_metadata?: string | null;
            /**
             * @description Title of the structured content
             * @example Some webpage
             */
            title?: string;
        } & {
            [key: string]: components["schemas"]["LocalizedFieldValuesWithEmbeddedContent"][];
        };
        RenditionConfigResponse: {
            /**
             * @description The asset type this config applies to
             * @example image
             */
            asset_type: string;
            /**
             * @description The expected height (in pixels) of the renditions generated using this config
             * @example 400
             */
            height: number | null;
            /**
             * @description Unique identifier of the rendition config
             * @example 5d7f910551b00a722e0418830cee6631
             */
            id: string;
            /**
             * @description The file format of the target asset
             * @example JPG
             */
            input_format: string | null;
            /**
             * @description Name of the rendition config
             * @example Facebook
             */
            name: string;
            /**
             * @description The file format of the generated rendition
             * @example PNG
             */
            output_format: string | null;
            /**
             * @description The quality (percent) of the renditions generated using this config
             * @example 80
             */
            quality: number | null;
            /**
             * @description The expected width (in pixels) of the renditions generated using this config
             * @example 600
             */
            width: number | null;
        };
        ResourceChoiceResponse: {
            /**
             * @description Unique identifier of the choice
             * @example 9119a313057e401189407116fcd3
             */
            id: string;
            /**
             * @description Name of the choice
             * @example Some choice
             */
            name: string;
        };
        /** @description Payload to associate labels to a resource */
        ResourceLabelRequest: {
            /**
             * @description Label group
             * @example 2467e583a60e23fda2b89db81a453cd2
             */
            group: string;
            /** @description List of values of the label */
            values: string[];
        };
        /** @description Label associated to a fetched resource */
        ResourceLabelResponse: {
            /** @description Label group */
            group: {
                /**
                 * @description Identifier of the label group
                 * @example 2467e583a60e23fda2b89db81a453cd2
                 */
                id: string;
                /**
                 * @description Name of the label group
                 * @example Content Format
                 */
                name: string;
            };
            /** @description List of values of the label */
            values: {
                /**
                 * @description Identifier of the label value
                 * @example 71c378f3fee3d822759d1bdc2aab628c
                 */
                id: string;
                /**
                 * @description Name of the label value
                 * @example Photos
                 */
                name: string;
            }[];
        };
        RichTextFieldDefinition: {
            core: components["schemas"]["CoreFieldDef"];
            /** @description Default values of the field */
            default_values?: string[];
            /** @description Maximum length of the field */
            max_visual_text_length?: number;
            /** @description Minimum length of the field */
            min_visual_text_length: number;
        };
        RichTextFieldValueModel: {
            /** @description Order index of the field value */
            order_index?: number;
            /** @description Value of the field value */
            rich_text_value: string;
        };
        SCContentMigrationCreateRequest: {
            /** @description Unique identifier of the user who migrated the content */
            created_by: string;
            /** @description List of fields */
            fields?: {
                [key: string]: components["schemas"]["LocalizedFieldValues"][];
            };
            /** @description Unique identifier of the content type version to migrate the content to */
            new_content_type_version_id?: string;
            /**
             * Source
             * @description Source of the content
             */
            source?: string | null;
            /**
             * Sourceid
             * @description Source id of the content
             */
            source_id?: string | null;
        };
        SCContentPreviewAcknowledgeRequest: {
            /** @description Unique identifier of the user who acknowledged the preview request */
            acknowledged_by: string;
            /** @description Content hash of the preview request */
            content_hash: string;
        };
        SCContentPreviewCompleteRequest: {
            keyed_previews: {
                [key: string]: string | components["schemas"]["KeyedPreviewCompletedModel"] | components["schemas"]["KeyedPreviewErrorModel"];
            };
        };
        SCContentType: components["schemas"]["BaseContentTypeModel"] & {
            latest_version: components["schemas"]["SCContentTypeVersion"];
        };
        SCContentTypeCreateRequest: {
            /** @description Unique identifier of the user who is creating the content type */
            created_by: string;
            details: components["schemas"]["CoreContentType"];
            /** @description Expected locales for the content type */
            expected_locales?: string[];
            field_definitions: (components["schemas"]["ContentTypeFieldDefinition"] | components["schemas"]["LibraryAssetFieldDefinition"] | components["schemas"]["TextFieldDefinition"] | components["schemas"]["BaseFieldDefinition"] | components["schemas"]["DatetimeFieldDefinition"] | components["schemas"]["RichTextFieldDefinition"] | components["schemas"]["ChoiceFieldDefinition"] | components["schemas"]["NumberFieldDefinition"])[];
            /** @description Source for the content type */
            source?: string;
            /** @description Source ID for the content type */
            source_id?: string;
            /** @description Source metadata for the content type */
            source_metadata?: string;
        };
        SCContentTypeCreateResponse: {
            /** @description Unique identifier of the content type */
            content_type_guid?: string;
            /** @description Unique identifier of the content type version */
            content_type_version_guid?: string;
            /** @description Created status of the requested content type */
            created: boolean;
        };
        SCContentTypeManagedMigrationCreateRequest: {
            /** @description Unique identifier of the user who created the job. */
            created_by: string;
            /** @description Default values for the migration job, if any. */
            default_values?: components["schemas"]["LocalizedFieldValues"];
            /** @description The ID of the source content type version to migrate from. */
            source_content_type_version_id: string;
        };
        SCContentTypeManagedMigrationResponse: {
            content_migration_summary?: components["schemas"]["ContentMigrationSummary"];
            /** @description The ID of the content type. */
            content_type_id?: string;
            /**
             * Format: date-time
             * @description Date and time on which the managed migration job was created,  in ISO 8601 UTC format (2020-02-10T10:40:45Z).
             * @example 2019-10-06T13:15:30Z
             */
            created_at?: string;
            /** @description The default values for the content type. */
            default_values?: Record<string, never>;
            /** @description The ID of the managed migration job. */
            id?: string;
            /** @description The ID of the instance. */
            instance_id?: string;
            /** @description The ID of the source content type version. */
            source_content_type_version_id?: string;
            /**
             * @description The status of the managed migration job.
             * @enum {string}
             */
            status?: "not_started" | "in_progress" | "success" | "error";
            /** @description The ID of the target content type version. */
            target_content_type_version_id?: string;
            /**
             * @description Date and time on which the managed migration job was last updated,  in ISO 8601 UTC format (2020-02-10T10:40:45Z).
             * @example 2019-10-06T13:15:30Z
             */
            updated_at?: string;
        };
        SCContentTypeManagedMigrationStartResponse: {
            /** @description The unique identifier of job. */
            job_id?: string;
            /** @description True if job started, otherwise false. */
            started?: boolean;
        };
        SCContentTypeManagedMigrationValidateRequest: {
            /** @description Default values for the migration, if any. */
            default_values?: components["schemas"]["LocalizedFieldValues"];
            /** @description The ID of the source content type version to migrate from. */
            source_content_type_version_id: string;
        };
        SCContentTypeUpdateRequest: {
            details: components["schemas"]["CoreContentType"];
            /** @description Source metadata of the content type */
            source_metadata?: string;
            /** @description Unique identifier of the user who updated the content type */
            updated_by: string;
        };
        SCContentTypeUpdateResponse: {
            /** @description Updated status of the request */
            updated: boolean;
        };
        SCContentTypeVersion: components["schemas"]["BaseContentTypeVersionModel"] & {
            field_definitions: (components["schemas"]["ContentTypeFieldDefinition"] | components["schemas"]["LibraryAssetFieldDefinition"] | components["schemas"]["TextFieldDefinition"] | components["schemas"]["BaseFieldDefinition"] | components["schemas"]["DatetimeFieldDefinition"] | components["schemas"]["RichTextFieldDefinition"] | components["schemas"]["ChoiceFieldDefinition"] | components["schemas"]["NumberFieldDefinition"])[];
        };
        SCContentTypeVersionCreateRequest: {
            /** @description Unique identifier of the user who created the version */
            created_by: string;
            /** @description Expected locales of the version */
            expected_locales?: string[];
            field_definitions: (components["schemas"]["ContentTypeFieldDefinition"] | components["schemas"]["LibraryAssetFieldDefinition"] | components["schemas"]["TextFieldDefinition"] | components["schemas"]["BaseFieldDefinition"] | components["schemas"]["DatetimeFieldDefinition"] | components["schemas"]["RichTextFieldDefinition"] | components["schemas"]["ChoiceFieldDefinition"] | components["schemas"]["NumberFieldDefinition"])[];
        };
        Section: {
            /**
             * @description Id of the template section
             * @example 9ec506f480dbc2d5f6919168ae448c41
             */
            id: string;
            /**
             * @description Name of the template section
             * @example Test section
             */
            label: string;
        };
        Settings: {
            /** @description Fetched settings */
            resources: components["schemas"]["SettingsResources"];
        };
        SettingsApp: {
            /**
             * @description List of the authorization callback URLs
             * @example [
             *       "https://example1.com",
             *       "https://example2.com"
             *     ]
             */
            authorization_callback_urls: string[];
            /**
             * @description Description of the app
             * @example Sample description
             */
            description: string;
            /**
             * @description Environment of the app
             * @example production
             */
            env: string;
            /**
             * @description Indicates whether an email should be exposed in an API response
             * @example true
             */
            expose_email: boolean;
            /**
             * @description Home page URL for app
             * @example https://example1.com
             */
            homepage_url: string;
            /**
             * @description Name of the app
             * @example App 1
             */
            name: string;
        };
        SettingsChangeset: {
            /** @description Created settings */
            create: components["schemas"]["SettingsChangesetBase"];
            /** @description Updated settings */
            update: components["schemas"]["SettingsChangesetBase"];
        };
        SettingsChangesetBase: {
            /**
             * @description List of the app names
             * @example [
             *       "A App"
             *     ]
             */
            apps: string[];
            /**
             * @description List of the custom field names
             * @example [
             *       "A Custom Field"
             *     ]
             */
            custom_fields: string[];
            /**
             * @description List of the label names
             * @example [
             *       "A LabelGroup"
             *     ]
             */
            labels: string[];
            /**
             * @description List of the routing rule names
             * @example [
             *       "A routing Rule"
             *     ]
             */
            routing_rules: string[];
            /**
             * @description List of the template names
             * @example [
             *       "A Template"
             *     ]
             */
            templates: string[];
            /**
             * @description List of the webhook names
             * @example [
             *       "A Webhook"
             *     ]
             */
            webhooks: string[];
            /**
             * @description List of the workflow names
             * @example [
             *       "A WorkFlow"
             *     ]
             */
            workflows: string[];
        };
        /** @description Settings for custom fields */
        SettingsCustomField: {
            /** @description List of checkbox custom fields */
            checkbox: components["schemas"]["MultipleOptionCustomField"][];
            /** @description List of currency custom fields */
            currency: components["schemas"]["CurrencyCustomField"][];
            /** @description List of date custom fields */
            date: components["schemas"]["LabelOnlyCustomField"][];
            /** @description List of dropdown custom fields */
            dropdown: components["schemas"]["MultipleOptionCustomField"][];
            /** @description List of image custom fields */
            image: components["schemas"]["LabelOnlyCustomField"][];
            /** @description List of multi-select dropdown custom fields */
            multi_select_dropdown: components["schemas"]["MultipleOptionCustomField"][];
            /** @description List of multichoice custom fields */
            multichoice: components["schemas"]["MultipleOptionCustomField"][];
            /** @description List of number custom fields */
            number: components["schemas"]["NumberCustomField"][];
            /** @description List of percentage custom fields */
            percentage: components["schemas"]["PercentageCustomField"][];
            /** @description List of rich text custom fields */
            richtext: components["schemas"]["LabelOnlyCustomField"][];
            /** @description List of string custom fields */
            string: components["schemas"]["LabelOnlyCustomField"][];
            /** @description List of textarea custom fields */
            textarea: components["schemas"]["LabelOnlyCustomField"][];
            /** @description List of video custom fields */
            video: components["schemas"]["LabelOnlyCustomField"][];
        };
        /** @description Choices of the field */
        SettingsFieldChoiceCreateRequest: (components["schemas"]["GenericSettingsFieldChoiceCreatePayload"] | components["schemas"]["LabelTypeSettingsFieldChoiceCreatePayload"])[];
        /** @description Choices of the field */
        SettingsFieldChoiceCreateResponse: {
            /**
             * @description Identifier of the field choice
             * @example 675ca12f4fa3411ecb37802420ac82ab7
             */
            id: string;
            /**
             * @description Name of the choice
             * @example Choice 1
             */
            name: string;
        }[];
        /** @description Payload for updating choice */
        SettingsFieldChoiceUpdateRequest: components["schemas"]["GenericSettingsFieldChoiceUpdatePayload"] | components["schemas"]["LabelTypeSettingsFieldChoiceUpdatePayload"];
        SettingsFieldCreateRequest: components["schemas"]["BaseSettingsFieldCreatePayload"] | components["schemas"]["LabelTypeSettingsFieldCreatePayload"] | components["schemas"]["DropdownTypeSettingsFieldCreatePayload"] | components["schemas"]["CheckboxAndRadioTypeSettingsFieldCreatePayload"] | components["schemas"]["GenericNumberTypeSettingsFieldCreatePayload"] | components["schemas"]["CurrencyNumberTypeSettingsFieldCreatePayload"];
        SettingsFieldUpdateRequest: components["schemas"]["BaseSettingsFieldUpdatePayload"] | components["schemas"]["LabelAndDropdownTypeSettingsFieldUpdatePayload"] | components["schemas"]["GenericNumberTypeSettingsFieldUpdatePayload"] | components["schemas"]["CurrencyNumberTypeSettingsFieldUpdatePayload"];
        SettingsLabelGroup: {
            /**
             * @description Single or multiple value support indicator for the label group
             * @example true
             */
            has_single_value: boolean;
            /**
             * @description Name of the label group
             * @example Content Pillar
             */
            label_type: string;
            /** @description List of labels of the label group */
            labels: components["schemas"]["SettingsLabelGroupLabel"][];
        };
        SettingsLabelGroupLabel: {
            /**
             * @description Color of the label
             * @default null
             * @example #4ECFD5
             */
            color: string | null;
            /**
             * @description Name of the label
             * @example Pillar1
             */
            name: string;
        };
        SettingsResources: {
            /** @description Settings apps */
            apps: components["schemas"]["SettingsApp"][];
            custom_fields: components["schemas"]["SettingsCustomField"];
            /** @description Settings labels */
            labels: components["schemas"]["SettingsLabelGroup"][];
            /** @description Settings routing rules */
            routing_rules: components["schemas"]["SettingsRoutingRule"][];
            /** @description Settings templates */
            templates: components["schemas"]["SettingsTemplate"][];
            /** @description Settings webhooks */
            webhooks: components["schemas"]["SettingsWebhook"][];
            /** @description Settings workflows */
            workflows: components["schemas"]["SettingsWorkflow"][];
        };
        SettingsRoutingRule: {
            /**
             * @description Description of the routing rule
             * @example Sample routing rule description
             */
            description: string;
            /**
             * @description Name of the routing rule
             * @example example Routing Rule
             */
            name: string;
            /** @description List of the associated rules */
            rules: {
                /**
                 * @description Associated custom field type
                 * @example multichoice
                 */
                custom_field_type: string | null;
                /**
                 * @description Name of the rule field type
                 * @example customField
                 */
                field_type: string;
                /**
                 * @description Name of the rule identifier
                 * @example exampleCustomField1
                 */
                identifier: string | null;
                /**
                 * @description Name of the rule operator
                 * @example equal
                 */
                operator: string;
                /**
                 * @description Unit of rule's field. Not null when the rule's `field_type` value is `dueDate`
                 * @example days
                 * @enum {string|null}
                 */
                unit: "days" | "weeks" | "months" | "years" | null;
                /**
                 * @description List of the associated label, custom field, or due date choice
                 * @example [
                 *       "value1",
                 *       "value2"
                 *     ]
                 */
                value: string[];
            }[];
            /**
             * @description Name of the associated template
             * @example Sample template name
             */
            template_name: string | null;
            /**
             * @description List of the routing rule watchers email
             * @example [
             *       "examplewatcher@gmail.com"
             *     ]
             */
            watchers: string[];
        };
        SettingsTemplate: {
            /**
             * @description Description of the template
             * @example Template description
             */
            description: string;
            /** @description List of the template form_fields */
            form_fields: components["schemas"]["FormField"][];
            /** @description List of the template instructions */
            instructions: components["schemas"]["Instruction"][];
            /**
             * @description Name of the template
             * @example Template 1
             */
            name: string;
            /**
             * @description Public description of the template
             * @example Template public description
             */
            public_description: string;
            /** @description List of the template sections */
            sections: components["schemas"]["Section"][];
            /**
             * @description List of areas where you can use the template
             * @example [
             *       "campaign_brief",
             *       "task_brief",
             *       "work_request"
             *     ]
             */
            types: string[];
        };
        SettingsUpdateResponse: {
            /** @description Changeset of created or updated settings */
            changeset: components["schemas"]["SettingsChangeset"];
            /** @description Created or updated settings */
            resources: components["schemas"]["SettingsResources"];
        };
        SettingsWebhook: {
            /**
             * @description Callback URL of the webhook
             * @example https://exampledomain.com
             */
            callback_url: string;
            /**
             * @description Description of the webhook
             * @example Sample description
             */
            description: string;
            /**
             * @description List of the subscribed event names for the webhook
             * @example [
             *       "asset_added",
             *       "asset_removed"
             *     ]
             */
            event_names: string[];
            /**
             * @description Name of the webhook
             * @example Webhook 1
             */
            name: string;
            /**
             * @description Webhook callback URL secret
             * @example secret121
             */
            secret: string;
        };
        SettingsWorkflow: {
            /** @description List of the disallowed channels */
            blacklisted_channels: components["schemas"]["WorkflowChannel"][];
            /** @description List of the associated custom field */
            custom_fields: components["schemas"]["WorkflowCustomfield"][];
            /** @description List of the associated channels */
            default_channels: components["schemas"]["WorkflowChannel"][];
            /**
             * @description Description of the workflow
             * @default null
             * @example Sample description
             */
            description: string | null;
            /**
             * @description Indicates whether the workflow is flexible
             * @example false
             */
            is_flexible_workflow: boolean;
            /**
             * @description Indicates whether the contents of the task created from this workflow are not pushed to the library by default
             * @example false
             */
            is_not_pushed_to_library_by_default: boolean;
            /**
             * @description Indicates whether the contents of the task created from this workflow are released as an asset
             * @example false
             */
            is_released_as_asset: boolean;
            /** @description List of the associated labels */
            labels: components["schemas"]["WorkflowLabel"][];
            /**
             * @description Name of the workflow
             * @example A Workflow
             */
            name: string;
            /** @description List of the workflow steps */
            steps: components["schemas"]["WorkflowStep"][];
        };
        StringTypeObjectField: components["schemas"]["BaseObjectField"] & components["schemas"]["StringTypeObjectFieldUpdatePayload"] & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "rich_text" | "text" | "text_area";
        };
        StringTypeObjectFieldUpdatePayload: components["schemas"]["BaseObjectFieldUpdatePayload"] & {
            /** @description Any single string */
            values?: string[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "rich_text" | "text" | "text_area";
        };
        /** @description The fields of structured content */
        StructuredContentBody: {
            /** @description Whether the structured content has embedded content reference field values */
            has_embedded?: boolean | null;
        } & {
            [key: string]: components["schemas"]["LocalizedFieldValuesWithEmbeddedContent"][];
        };
        /** @description The fields of structured content */
        StructuredContentFields: {
            [key: string]: (components["schemas"]["PatchLocalizedFieldValuesWithEmbeddedContent"] | components["schemas"]["DeleteLocaleFieldValueModel"])[];
        };
        Tag: {
            /**
             * @description Unique identifier of the tag
             * @example 5d7f910551b00a722e0418830cee6633
             */
            guid: string;
            /**
             * @description Name of the tag
             * @example Sample Tag
             */
            name: string;
        };
        TaskArticle: {
            /**
             * Format: date-time
             * @description Date and time on which the article was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Content of the article
             * @example <p>Article text</p>
             */
            html_body: string;
            /**
             * @description Unique identifier of the article
             * @example 5d7f910551b00a722e0418830cee6631
             */
            id: string;
            /** @description Labels associated with the article */
            labels: components["schemas"]["ResourceLabelResponse"][];
            /** @description Meta links */
            links: {
                /**
                 * @description URL to POST draft to a task asset – not supported for `article` so the value will be `null`
                 * @example null
                 */
                drafts: string | null;
                /**
                 * @description URL of the article
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/articles/5d7f910551b00a722e0418830cee6631
                 */
                self: string;
                /**
                 * @description URL of the task that the article is associated with
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9
                 */
                task: string;
                /** @description Web URLs */
                web_urls: {
                    /**
                     * @description Web URL of the drafts of the article – not supported for `article` so the value will be `null`
                     * @example null
                     */
                    drafts: string | null;
                    /**
                     * @description Web URL of the article
                     * @example https://cmp.optimizely.com/cloud/taskv3/5f857f30e1c4a2038d6179e9?contentTabGuid=5d7f910551b00a722e0418830cee6631
                     */
                    self: string;
                    /**
                     * @description Web URL of the task that the article is associated with
                     * @example https://cmp.optimizely.com/cloud/taskv3/5f857f30e1c4a2038d6179e9
                     */
                    task: string;
                };
            };
            /**
             * Format: date-time
             * @description Date and time of the most recent modification of the article, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            modified_at: string;
            /**
             * @description Title of the article
             * @example 3 Ways Influencer Marketing Will Further Mature in 2020
             */
            title: string;
            /**
             * @description Public URL of the article
             * @example https://cmp.optimizely.com/new-article-published.html
             */
            url: string | null;
        };
        TaskAssetCommentResponse: {
            /**
             * @description List of attachments of the comment
             * @example [
             *       {
             *         "id": "a113667245d111eb8945000c",
             *         "name": "sample.png",
             *         "url": "https://files.cmp.optimizely.com/download/96c314a645d111eb8945000c291b51d4?token="
             *       }
             *     ]
             */
            attachments: components["schemas"]["AttachmentResponse"][];
            /**
             * Format: date-time
             * @description Creation date and time of the comment, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2020-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Unique identifier of the comment
             * @example 5fe38c39574b52a62a089239
             */
            id: string;
            /**
             * @description Determines if the comment is resolved in Optimizely CMP
             * @example false
             */
            is_resolved: boolean;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the asset
                 * @example https://api.cmp.optimizely.com/v3/tasks/5d7f910551b00a722e0418830cee5534/images/6d7f910551b00a722e0418830cee5564
                 */
                asset: string;
                /**
                 * @description URL of the task
                 * @example https://api.cmp.optimizely.com/v3/tasks/5d7f910551b00a722e0418830cee5534
                 */
                task: string;
            };
            /**
             * Format: date-time
             * @description Last modification date and time of the comment, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2020-10-06T14:15:30Z
             */
            modified_at: string;
            /**
             * @description Content of the comment
             * @example Well done!
             */
            value: string;
        };
        TaskAssetDraftBrandComplianceRequest: {
            /** @description List of compliance categories */
            categories: {
                /** @description List of criteria */
                criteria: {
                    /**
                     * @description Description of the criterion
                     * @example For all CMYK, RGB and HTML breakdowns, please refer to the full Colour Policy guideline
                     */
                    description: string | null;
                    /**
                     * @description Unique identifier of the criterion
                     * @example 66dd74f41a335b59cd98915c
                     */
                    id: string;
                    /**
                     * @description Whether the criterion is selected or not
                     * @example true
                     */
                    selected: boolean | null;
                }[];
                /**
                 * @description Unique identifier of the category
                 * @example 66dd74f41a335b59cd98915b
                 */
                id: string;
                /**
                 * @description Notes left by the reviewer for the category
                 * @example Look's good
                 */
                notes: string | null;
                /**
                 * @description Brand compliance status of the category
                 * @example compliant
                 * @enum {string|null}
                 */
                status: "compliant" | "not compliant" | "not applicable" | null;
            }[];
            /**
             * @description Brand compliance status of the draft
             * @example approved
             * @enum {string}
             */
            status: "approved" | "declined" | "not_reviewed";
        };
        TaskAssetDraftBrandComplianceResponse: {
            /** @description List of compliance categories */
            categories: {
                /** @description List of criteria */
                criteria: {
                    /**
                     * @description Description of the criterion
                     * @example For all CMYK, RGB and HTML breakdowns, please refer to the full Colour Policy guideline
                     */
                    description: string | null;
                    /**
                     * @description Unique identifier of the criterion
                     * @example 66dd74f41a335b59cd98915c
                     */
                    id: string;
                    /**
                     * @description Name of the criterion
                     * @example Sub-category-1A
                     */
                    name: string;
                    /**
                     * @description Whether the criterion is selected or not
                     * @example true
                     */
                    selected: boolean | null;
                }[];
                /**
                 * @description Unique identifier of the category
                 * @example 66dd74f41a335b59cd98915b
                 */
                id: string;
                /**
                 * @description Name of the category
                 * @example Category-1
                 */
                name: string;
                /**
                 * @description Notes left by the reviewer for the category
                 * @example Look's good
                 */
                notes: string | null;
                /**
                 * @description Brand compliance status of the category
                 * @example compliant
                 * @enum {string|null}
                 */
                status: "compliant" | "not compliant" | "not applicable" | null;
            }[];
            /**
             * Format: date-time
             * @description Date and time of when the draft was reviewed, in ISO 8601 UTC format
             * @example 2020-10-06T13:15:30Z
             */
            reviewed_at: string | null;
            /**
             * @description Unique identifier of the user who reviewed the draft
             * @example 66d837c62373533177b59db3
             */
            reviewed_by: string | null;
            /**
             * @description Brand compliance status of the draft
             * @example approved
             * @enum {string}
             */
            status: "approved" | "declined" | "not_reviewed";
        };
        TaskAssetDraftListResponseItem: {
            /** @description Content of the draft */
            content: {
                /**
                 * @description Type of the content
                 * @example url
                 * @enum {string}
                 */
                type: "url";
                /**
                 * @description Content of the draft; download URL for `image`, `raw file`, and `video`
                 * @example http://images.cmp.optimizely.com/Zz0xODQ3NDU3Y2Y2YmYzOTlmNjkyOTgyZDY3Y2I3YWM2OA==S
                 */
                value: string;
            };
            /**
             * Format: date-time
             * @description Date and time on which the draft was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description The serial number of the draft
             * @example 2
             */
            draft_number: number;
            /**
             * @description Unique identifier of the draft
             * @example 5d7f910551b00a722e0418830cee6631
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the asset that the draft is associated with
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/images/5d7f910551b00a722e0418830cee6631
                 */
                asset: string;
                /**
                 * @description URL of the task that the asset is associated with
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9
                 */
                task: string;
            };
            /**
             * @description MIME type of the draft
             * @example image/png
             */
            mime_type: string;
            /**
             * @description Title of the draft
             * @example sample_image.png
             */
            title: string;
            /**
             * @description Type of the draft
             * @example image
             * @enum {string}
             */
            type: "image" | "video" | "raw_file";
        };
        TaskAssetDraftResponse: {
            /**
             * @description Unique identifier of the asset
             * @example 5d7f910551b00a722e0418830cee6631
             */
            asset_id: string;
            /** @description Content of the draft */
            content: {
                /**
                 * @description Type of the content
                 * @example url
                 * @enum {string}
                 */
                type: "url";
                /**
                 * @description Content of the draft; download URL for `image`, `raw file`, and `video`
                 * @example http://images.cmp.optimizely.com/Zz0xODQ3NDU3Y2Y2YmYzOTlmNjkyOTgyZDY3Y2I3YWM2OA==S
                 */
                value: string;
            };
            /**
             * Format: date-time
             * @description Date and time on which the draft was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description The serial number of the draft
             * @example 2
             */
            draft_number: number;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the asset that the draft is associated with
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/images/5d7f910551b00a722e0418830cee6631
                 */
                asset: string;
                /**
                 * @description URL of the task that the asset is associated with
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9
                 */
                task: string;
            };
            /**
             * @description MIME type of the draft
             * @example image/png
             */
            mime_type: string;
            /**
             * @description Title of the draft
             * @example sample_image.png
             */
            title: string;
            /**
             * @description Type of the draft
             * @example image
             * @enum {string}
             */
            type: "image" | "video" | "raw_file";
        };
        TaskAssetFieldsUpdateRequest: (components["schemas"]["StringTypeObjectField"] | components["schemas"]["MultiChoiceTypeObjectField"] | components["schemas"]["DropdownTypeObjectField"] | components["schemas"]["RadioButtonTypeObjectField"] | components["schemas"]["NumberTypeObjectField"] | components["schemas"]["DateTypeObjectField"] | components["schemas"]["AssetTypeObjectField"])[];
        TaskAssetRequest: components["schemas"]["TaskAssetRequestForDirectUpload"] | components["schemas"]["TaskAssetRequestForLibrary"];
        TaskAssetRequestForDirectUpload: {
            /** @description Unique identifier of the file upload session. This is the `upload_meta_fields.key` field retrieved from the `/v3/upload-url` endpoint. */
            key: string;
            /**
             * @description Title of the asset
             * @example sample_image.png
             */
            title: string;
            /**
             * @description The mechanism for adding the asset to the task (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "direct_upload";
        };
        TaskAssetRequestForLibrary: {
            /**
             * Format: uuid
             * @description Unique identifier of the asset
             * @example d1efd3c4625111f0baa002420ac8001d
             */
            content_id: string;
            /**
             * @description The type of asset
             * @example image
             * @enum {string}
             */
            content_type: "article" | "image" | "raw_file" | "structured_content" | "video";
            /**
             * @description The mechanism for adding the asset to the task (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "fork_from_library";
        };
        TaskAssetResponse: {
            content: components["schemas"]["AssetContent"];
            /**
             * Format: date-time
             * @description Date and time on which the asset was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Unique identifier for the asset
             * @example 5d7f910551b00a722e0418830cee6631
             */
            id: string;
            /** @description Labels associated with the asset */
            labels: components["schemas"]["ResourceLabelResponse"][];
            /** @description Meta links */
            links: {
                /**
                 * @description URL to POST draft to the asset - null for `article` - non-null for `image`, `video`, and `raw file`.
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/assets/5d7f910551b00a722e0418830cee6631/drafts
                 */
                drafts: string | null;
                /**
                 * @description URL of the asset
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/image/5d7f910551b00a722e0418830cee6631
                 */
                self: string;
                /**
                 * @description URL of the task that the asset is associated with
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9
                 */
                task: string;
                /** @description Web URLs */
                web_urls: {
                    /**
                     * @description Web URL of the drafts of the asset - null for `article` - non-null for `image`, `video`, and `raw file`.
                     * @example https://cmp.optimizely.com/cloud/task/5f857f30e1c4a2038d6179e9/image/5d7f910551b00a722e0418830cee6631
                     */
                    drafts: string | null;
                    /**
                     * @description Web URL of the asset
                     * @example https://cmp.optimizely.com/cloud/taskv3/5f857f30e1c4a2038d6179e9?contentTabGuid=5d7f910551b00a722e0418830cee6631
                     */
                    self: string;
                    /**
                     * @description Web URL of the task that the asset is associated with
                     * @example https://cmp.optimizely.com/cloud/taskv3/5f857f30e1c4a2038d6179e9
                     */
                    task: string;
                };
            };
            /**
             * @description MIME type of the asset
             * @example image/png
             */
            mime_type: string;
            /**
             * Format: date-time
             * @description Date and time of the most recent modification of the asset, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            modified_at: string;
            /**
             * @description Title of the asset
             * @example sample_image.png
             */
            title: string;
            /**
             * @description Type of the asset
             * @example image
             * @enum {string}
             */
            type: "article" | "image" | "video" | "raw_file" | "structured_content";
        };
        TaskBriefResponse: {
            /** @description List of fields of the brief */
            fields: {
                /**
                 * @description Name of the brief field
                 * @example My Dropdown
                 */
                name: string;
                /**
                 * @description Unique identifier of the settings field. If available, can be used to look up the field details using settings API.
                 * @example 67c5684de8b943290648a693
                 */
                settings_field_id?: string | null;
                /** @description List of values set for the brief field */
                value: string[];
            }[];
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the brief
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/brief
                 */
                self: string;
                /**
                 * @description URL of the task
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9
                 */
                task: string;
            };
            template?: (components["schemas"]["TaskBriefTemplateValue"] | components["schemas"]["NullValue"]) | null;
            /**
             * @description Title of the task brief
             * @example Awesome Task Brief
             */
            title: string;
            /**
             * @description Type of the task brief
             * @example template
             * @enum {string}
             */
            type: "template" | "attachment" | "text";
        };
        /** @description Template info of the brief if the brief is of template type */
        TaskBriefTemplateValue: {
            /**
             * @description Unique identifier of the template
             * @example 9nu8ue9wf8u9nusd9f
             */
            id: string;
            /**
             * @description Name of the template
             * @example My Template
             */
            name: string;
        };
        TaskCommentResponse: {
            /**
             * @description List of attachments of the comment
             * @example [
             *       {
             *         "id": "a113667245d111eb8945000c",
             *         "name": "sample.png",
             *         "url": "https://files.cmp.optimizely.com/download/96c314a645d111eb8945000c291b51d4?token="
             *       }
             *     ]
             */
            attachments: components["schemas"]["AttachmentResponse"][];
            /**
             * @description Unique identifier of the user who posted the comment
             * @example 62a8455e29fed3045d9c7a8b
             */
            comment_by: string;
            /**
             * Format: date-time
             * @description Creation date and time of the task comment, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2020-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Unique identifier of the task comment
             * @example 5fe38c39574b52a62a089239
             */
            id: string;
            /**
             * @description Determines if the comment is resolved in Optimizely CMP
             * @example false
             */
            is_resolved: boolean;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the user who posted the comment
                 * @example https://api.cmp.optimizely.com/v3/users/62a8455e29fed3045d9c7a8b
                 */
                comment_by: string;
                /**
                 * @description URL of the task
                 * @example https://api.cmp.optimizely.com/v3/tasks/5fe38c39574b52a62a089239
                 */
                task: string;
            };
            /**
             * Format: date-time
             * @description Last modification date and time of the task comment, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2020-10-06T14:15:30Z
             */
            modified_at: string;
            /**
             * @description Unique identifier of the parent comment
             * @example 5fe38c39574b52a62a089238
             */
            parent_comment_id: string | null;
            /**
             * @description Content of the comment
             * @example Well done!
             */
            value: string;
        };
        TaskCreateRequest: {
            /**
             * @description Id of the campaign to be associated with the task
             * @example 6684f31be89a526d030cfc0f
             */
            campaign_id?: string;
            /**
             * Format: date-time
             * @description Due date and time of the task, in ISO 8601 UTC format (2020-02-10T10:40:45Z). The `due_at` field is required if `workflow_id` is present and the workflow does not have step duration. If `workflow_id` is present and the workflow has step duration, then either `start_at` or `due_at` is required.
             * @example 2020-10-06T13:15:30Z
             */
            due_at?: string;
            /**
             * @description Id of the owner of the task
             * @example 665d5717c8ba0201647a0991
             */
            owner_id?: string;
            /**
             * Format: date-time
             * @description Start date and time of the task, in ISO 8601 UTC format (2020-02-10T10:40:45Z). If `workflow_id` is present and the workflow has step duration, then either `start_at` or `due_at` is required. If both `start_at` and `due_at` is provided with a `workflow_id` that has step duration, then the `start_at` field will be ignored.
             * @example 2020-10-01T06:00:00Z
             */
            start_at?: string;
            /**
             * @description Title of the task
             * @example Awesome Task
             */
            title: string;
            /**
             * @description Id of the workflow to be used in the task
             * @example 665d56034854d100ad7cb8fb
             */
            workflow_id?: string;
        };
        TaskCustomField: {
            /**
             * @description Unique identifier for the custom field
             * @example 9nu8ue9wf8u9nusd9q
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the available custom field choices
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/custom-fields/9nu8ue9wf8u9nusd9q/choices
                 */
                choices: string | null;
                /**
                 * @description URL of the custom field
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/custom-fields/9nu8ue9wf8u9nusd9q
                 */
                self: string;
            };
            /**
             * @description Name of the custom field
             * @example My Dropdown
             */
            name: string;
            /**
             * @description Type of the custom field
             * @example dropdown
             * @enum {string}
             */
            type: "text_field" | "multi_line_text_field" | "checkboxes" | "dropdown" | "multi_select_dropdown" | "multiple_choice" | "date_field" | "image" | "video" | "rich_text_field";
            /** @description List of values of the custom field */
            values: {
                /**
                 * @description Identifier of the custom field value
                 * @example 1nu8ue9wf8u9nusd9u
                 */
                id: string | null;
                /**
                 * @description Name of the custom field value
                 * @example Some text
                 */
                name: string;
            }[];
        };
        TaskCustomFieldChoiceListResponseItem: {
            /**
             * @description Unique identifier of the custom field choice
             * @example j90uv0sd9i0si09sdip
             */
            id: string;
            /**
             * @description Name of the custom field choice
             * @example Some option
             */
            name: string;
        };
        TaskCustomFieldUpdateRequest: {
            /**
             * @description An array of choice IDs, if the custom field has choices. Otherwise, text values to update the custom field values with. Must contain at least one id/value in the array. For custom fields of type `checkboxes`, you can provide an array of multiple-choice IDs as values.
             *     Value for the custom field of type `date_field` must follow this format: `YYYY-MM-DD`.
             */
            values: string[];
        };
        TaskExternalWorkRequest: {
            /**
             * @description Identifier of the external work
             * @example MY-PROJ-123
             */
            identifier?: string | null;
            /**
             * @description Status of the external work
             * @example In Progress
             */
            status?: string | null;
            /**
             * @description Title of the external work
             * @example A very important ticket
             */
            title?: string | null;
            /**
             * @description Link to the external work (must be a valid URL with HTTPS scheme)
             * @example https://example.com/some-project/MY-PROJ-123
             */
            url?: string | null;
        };
        TaskExternalWorkResponse: {
            /**
             * @description Name of the external system
             * @example jira
             */
            external_system: string;
            /**
             * @description Identifier of the external work
             * @example MY-PROJ-123
             */
            identifier: string | null;
            /** @description Meta links */
            links: {
                /**
                 * @description URL to get or update the external work information
                 * @example https://api.cmp.optimizely.com/v3/tasks/5d7f910551b00a722e0418830cee6631/steps/32982hf94j98cnr48034nv035/sub-steps/9n390809r384nv503459075034nv5/external-work
                 */
                self: string;
            };
            /**
             * @description Status of the external work
             * @example In Progress
             */
            status: string | null;
            /**
             * @description Title of the external work
             * @example A very important ticket
             */
            title: string | null;
            /**
             * @description Link to the external work
             * @example https://example.com/some-project/MY-PROJ-123
             */
            url: string | null;
        };
        TaskFieldUpdateRequest: components["schemas"]["StringTypeObjectFieldUpdatePayload"] | components["schemas"]["MultiChoiceTypeObjectFieldUpdatePayload"] | components["schemas"]["DropdownTypeObjectFieldUpdatePayload"] | components["schemas"]["RadioButtonTypeObjectFieldUpdatePayload"] | components["schemas"]["NumberTypeObjectFieldUpdatePayload"] | components["schemas"]["DateTypeObjectFieldUpdatePayload"] | components["schemas"]["AssetTypeObjectFieldUpdatePayload"];
        TaskImage: {
            /**
             * Format: date-time
             * @description Date and time on which the image was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Size of the image in bytes
             * @example 11244
             */
            file_size: number;
            /**
             * @description Unique identifier of the image
             * @example 5d7f910551b00a722e0418830cee6632
             */
            id: string;
            /** @description Width and height of the image in pixels */
            image_resolution: {
                /**
                 * @description Height of the image in pixels
                 * @example 800
                 */
                height: number;
                /**
                 * @description Width of the image in pixels
                 * @example 700
                 */
                width: number;
            };
            /** @description Labels associated with the image */
            labels: components["schemas"]["ResourceLabelResponse"][];
            /** @description Meta links */
            links: {
                /**
                 * @description URL to POST draft to the image
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/assets/5d7f910551b00a722e0418830cee6632/drafts
                 */
                drafts: string;
                /**
                 * @description URL of the image
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/images/5d7f910551b00a722e0418830cee6632
                 */
                self: string;
                /**
                 * @description URL of the task that the image is associated with
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9
                 */
                task: string;
                /** @description Web URLs */
                web_urls: {
                    /**
                     * @description Web URL of the drafts of the image
                     * @example https://cmp.optimizely.com/cloud/task/5f857f30e1c4a2038d6179e9/image/5d7f910551b00a722e0418830cee6632
                     */
                    drafts: string;
                    /**
                     * @description Web URL of the image
                     * @example https://cmp.optimizely.com/cloud/taskv3/5f857f30e1c4a2038d6179e9?contentTabGuid=5d7f910551b00a722e0418830cee6632
                     */
                    self: string;
                    /**
                     * @description Web URL of the task that the image is associated with
                     * @example https://cmp.optimizely.com/cloud/taskv3/5f857f30e1c4a2038d6179e9
                     */
                    task: string;
                };
            };
            /**
             * @description MIME type of the image
             * @example image/jpeg
             */
            mime_type: string;
            /**
             * Format: date-time
             * @description Date and time of the most recent modification of the image, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            modified_at: string;
            /**
             * @description Title of the image
             * @example cat.jpeg
             */
            title: string;
            /**
             * @description Download the URL of the image
             * @example http://images.cmp.optimizely.com/Zz0xODQ3NDU3Y2Y2YmYzOTlmNjkyOTgyZDY3Y2I3YWM2OA==
             */
            url: string;
        };
        TaskListResponseItem: {
            /**
             * @description Unique identifier of the campaign
             * @example 8n7f910o51b00b722o0418n30cie8211
             */
            campaign_id: Record<string, never>;
            /**
             * @description Due date of the task in ISO 8601 UTC format
             * @example 2024-06-01T06:00:00Z
             */
            due_at: string | null;
            /**
             * @description Unique identifier of the task
             * @example 837f910551b00a722e0418830cee6612
             */
            id: string;
            /**
             * @description Whether the task is archived or not
             * @example false
             */
            is_archived: boolean;
            /**
             * @description Whether the task is completed or not
             * @example false
             */
            is_completed: boolean;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the campaign that the task is associated with
                 * @example https://api.cmp.optimizely.com/v3/campaigns/2e7f9107351b00a722e0412230cee5534
                 */
                campaign: string | null;
                /**
                 * @description URL of the milestone that the task is associated with
                 * @example https://api.cmp.optimizely.com/v3/milestones/63f1c2b675be4132854d2741
                 */
                milestone: string | null;
                /**
                 * @description URL of the task
                 * @example https://api.cmp.optimizely.com/v3/tasks/5d7f910551b00a722e0418830cee5534
                 */
                self: string;
                /** @description Web URLs */
                web_urls: {
                    /**
                     * @description Web URL of the task
                     * @example https://cmp.optimizely.com/cloud/taskv3/5c52cc153db27d0e5141146f
                     */
                    self: string;
                };
                /**
                 * @description URL of the workflow the task
                 * @example https://api.cmp.optimizely.com/v3/workflows/63f1c2b675be4132854d2741
                 */
                workflow: string | null;
            };
            /**
             * @description Unique identifier of the milestone
             * @example 8n7f910o51b00b722o0418n30cie8211
             */
            milestone_id: Record<string, never>;
            /**
             * @description Modified date of the task in ISO 8601 UTC format
             * @example 2024-06-01T06:00:00Z
             */
            modified_at: string | null;
            /**
             * @description Reference identifier of the task
             * @example TSK-1589
             */
            reference_id: string;
            /**
             * @description Start date of the task in ISO 8601 UTC format
             * @example 2024-06-01T06:00:00Z
             */
            start_at: string | null;
            /**
             * @description Current status of the task
             * @example Not Started
             * @enum {string}
             */
            status: "Archived" | "Completed" | "Overdue" | "Not Started" | "In Progress" | "On Hold";
            /**
             * @description Title of the task
             * @example The awesome task
             */
            title: string;
            /**
             * @description Unique identifier of the workflow
             * @example 8n7f910o51b00b722o0418n30cie8211
             */
            workflow_id: Record<string, never>;
        };
        TaskRawFile: {
            /**
             * Format: date-time
             * @description Date and time on which the raw file was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Size of the raw file in bytes
             * @example 11244
             */
            file_size: number;
            /**
             * @description Unique identifier of the raw file
             * @example 5d7f910551b00a722e0418830cee6634
             */
            id: string;
            /** @description Labels associated to the raw file */
            labels: components["schemas"]["ResourceLabelResponse"][];
            /** @description Meta links */
            links: {
                /**
                 * @description URL to POST draft to the raw file
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/assets/5d7f910551b00a722e0418830cee6634/drafts
                 */
                drafts: string;
                /**
                 * @description URL of the raw file
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/raw-files/5d7f910551b00a722e0418830cee6634
                 */
                self: string;
                /**
                 * @description URL of the task that the raw file is associated with
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9
                 */
                task: string;
                /** @description Web URLs */
                web_urls: {
                    /**
                     * @description Web URL of the drafts of the raw file
                     * @example https://cmp.optimizely.com/cloud/task/5f857f30e1c4a2038d6179e9/raw_file/5d7f910551b00a722e0418830cee6634
                     */
                    drafts: string;
                    /**
                     * @description Web URL of the raw file
                     * @example https://cmp.optimizely.com/cloud/taskv3/5f857f30e1c4a2038d6179e9?contentTabGuid=5d7f910551b00a722e0418830cee6634
                     */
                    self: string;
                    /**
                     * @description Web URL of the task that the raw file is associated with
                     * @example https://cmp.optimizely.com/cloud/taskv3/5f857f30e1c4a2038d6179e9
                     */
                    task: string;
                };
            };
            /**
             * @description MIME type of the raw file
             * @example application/zip
             */
            mime_type: string;
            /**
             * Format: date-time
             * @description Date and time of the most recent modification of the raw file,  in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            modified_at: string;
            /**
             * @description Title of the raw file
             * @example documents.zip
             */
            title: string;
            /**
             * @description Download the URL of the raw file
             * @example https://files.cmp.optimizely.com/171451644651701b96f1122009f026bc
             */
            url: string;
        };
        TaskResponse: {
            /**
             * Format: date-time
             * @description Due date and time of the task, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2020-10-06T13:15:30Z
             */
            due_at: string | null;
            /**
             * @description Unique identifier of the task
             * @example 5d7f910551b00a722e0418830cee5534
             */
            id: string;
            /**
             * @description Determines whether the task is archived
             * @example false
             */
            is_archived: boolean;
            /**
             * @description Determines whether the task is completed
             * @example false
             */
            is_completed: boolean;
            /** @description Labels associated with the task */
            labels: components["schemas"]["ResourceLabelResponse"][];
            /** @description Meta links */
            links: {
                /**
                 * @description URL for getting the list of the task assets or adding a new asset to the task
                 * @example https://api.cmp.optimizely.com/v3/tasks/5d7f910551b00a722e0418830cee5534/assets
                 */
                assets: string;
                /**
                 * @description URL for getting the list of the task attachments
                 * @example https://api.cmp.optimizely.com/v3/tasks/5d7f910551b00a722e0418830cee5534/attachments
                 */
                attachments: string;
                /**
                 * @description URL of the task brief
                 * @example https://api.cmp.optimizely.com/v3/tasks/5d7f910551b00a722e0418830cee5534/brief
                 */
                brief: string | null;
                /**
                 * @description URL of the campaign that the task is associated with
                 * @example https://api.cmp.optimizely.com/v3/campaigns/2e7f9107351b00a722e0412230cee5534
                 */
                campaign: string;
                /**
                 * @description URL of the list of custom fields added to the task
                 * @example https://api.cmp.optimizely.com/v3/tasks/5d7f910551b00a722e0418830cee5534/custom-fields
                 */
                custom_fields: string | null;
                /**
                 * @description URL of the milestone that the task is associated with
                 * @example https://api.cmp.optimizely.com/v3/milestones/63f1c2b675be4132854d2741
                 */
                milestone: string | null;
                /**
                 * @description URL of the task
                 * @example https://api.cmp.optimizely.com/v3/tasks/5d7f910551b00a722e0418830cee5534
                 */
                self: string;
                /** @description Web URLs */
                web_urls: {
                    /**
                     * @description Web URL of the task brief
                     * @example https://cmp.optimizely.com/cloud/taskv3/5c52cc153db27d0e5141146f?activeTab=brief
                     */
                    brief: string;
                    /**
                     * @description Web URL of the task
                     * @example https://cmp.optimizely.com/cloud/taskv3/5c52cc153db27d0e5141146f
                     */
                    self: string;
                };
            };
            /**
             * @description Reference ID of the task
             * @example TSK-123
             */
            reference_id: string;
            /**
             * Format: date-time
             * @description Start date and time of the task, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            start_at: string | null;
            /**
             * @description Current status of the task
             * @example Not Started
             * @enum {string}
             */
            status: "Archived" | "Completed" | "Overdue" | "Not Started" | "In Progress" | "On Hold";
            /** @description Steps of the task */
            steps: components["schemas"]["TaskStep"][];
            /**
             * @description Title of the task
             * @example Awesome Task
             */
            title: string;
            /**
             * @description Unique identifier of workflow
             * @example 8q7f910551b00a722e0418830cee6612
             */
            workflow_id?: string | null;
        };
        TaskStep: {
            /**
             * @description Description of the step
             * @example Description of the step
             */
            description: string | null;
            /**
             * Format: date-time
             * @description Due date and time of the step containing the substep,  in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2020-10-06T13:15:30Z
             */
            due_at: string | null;
            /**
             * @description Unique identifier of the step
             * @example 6d7f910551b00a722e0418830cee5564
             */
            id: string;
            /**
             * @description Determines whether the step is completed in Optimizely CMP
             * @example false
             */
            is_completed: boolean;
            /** @description Substeps of the step */
            sub_steps: components["schemas"]["TaskSubStep"][];
            /**
             * @description Title of the step
             * @example Sample step
             */
            title: string;
        };
        TaskStepRequest: {
            /**
             * Format: date-time
             * @description Due date and time of the task step, in ISO 8601 UTC format (2020-02-10T10:40:45Z). Point to note: in the CMP web interface, a popup is shown while updating a step's due date to inform the user of temporal or sequential constraints based on the due date of other steps. However, this API will update the step's due date while superceding any validation for the workflow.
             * @example 2025-10-06T13:00:00Z
             */
            due_at?: string;
        };
        TaskStructuredContentCreateRequest: {
            /**
             * @description Unique identifier of the content type
             * @example 5d7f910551b00a722e0418830cee6633
             */
            content_type_guid: string;
            /**
             * Expired
             * @description Expired status of the content
             */
            expired?: boolean;
            /**
             * Expirydatetime
             * Format: date-time
             * @description Date and time on which the content will expire, in ISO 8601 UTC format
             * @example 2019-10-07T13:15:30Z
             */
            expiry_datetime?: string | null;
            /** fields */
            fields?: components["schemas"]["StructuredContentBody"];
            /**
             * @description The primary locale of the content
             * @example en_US
             */
            primary_locale?: string;
            /**
             * Rootcontent
             * @description This is true when the content is not embedded in another content through a content type reference field
             */
            root_content?: boolean;
            /**
             * Source
             * @description Source of the content
             */
            source?: string | null;
            /**
             * Sourceid
             * @description Source id of the content
             */
            source_id?: string | null;
            /**
             * Sourcemetadata
             * @description Source metadata of the content
             */
            source_metadata?: string | null;
            /**
             * @description Unique identifier of the template
             * @example 5d7f910551b00a722e0418830cee6633
             */
            template_guid?: string;
            /**
             * @description Title of the structured content
             * @example Some webpage
             */
            title?: string;
        };
        TaskStructuredContentDraftRequest: {
            /**
             * contentbody
             * @description The fields of structured content
             */
            content_body: components["schemas"]["LocalizedFieldValues"];
        };
        TaskStructuredContentUpdateRequest: {
            /**
             * Expired
             * @description Expired status of the content
             */
            expired?: boolean;
            /**
             * Expirydatetime
             * Format: date-time
             * @description Date and time on which the content will expire, in ISO 8601 UTC format
             * @example 2019-10-07T13:15:30Z
             */
            expiry_datetime?: string | null;
            /**
             * fields
             * @description The fields of structured content
             */
            fields?: components["schemas"]["StructuredContentFields"];
            /**
             * @description The primary locale of the content
             * @example en_US
             */
            primary_locale?: string;
            /**
             * Source
             * @description Source of the content
             */
            source?: string | null;
            /**
             * Sourceid
             * @description Source id of the content
             */
            source_id?: string | null;
            /**
             * Sourcemetadata
             * @description Source metadata of the content
             */
            source_metadata?: string | null;
            /**
             * @description Title of the structured content
             * @example Some webpage
             */
            title?: string;
        };
        TaskSubStep: {
            /**
             * @description ID of the assignee of the step containing the substep
             * @example 5fdf31d57f0d0e362e2b5908
             */
            assignee_id: string | null;
            /**
             * @description Type of the assignee of the step containing the substep.
             * @example user
             * @enum {string}
             */
            assignee_type: "user" | "team";
            /**
             * @description Unique identifier of the substep
             * @example 9c7f910551b00a722e0418830cee2564
             */
            id: string;
            /**
             * @description Determines whether the substep is completed in Optimizely CMP
             * @example false
             */
            is_completed: boolean;
            /**
             * @description Determines whether the substep is external in Optimizely CMP
             * @example true
             */
            is_external: boolean;
            /**
             * @description Determines whether the substep is In Progress in Optimizely CMP
             * @example false
             */
            is_in_progress: boolean;
            /**
             * @description Determines whether the substep is skipped in Optimizely CMP
             * @example false
             */
            is_skipped: boolean;
            /** @description Meta links */
            links: {
                /**
                 * @description Assignee of the step containing the substep
                 * @example https://api.cmp.optimizely.com/v3/users/5fdf31d57f0d0e362e2b5908
                 */
                assignee?: string | null;
                /**
                 * @description URL of the external work associated with substep if substep is external in _Optimizely CMP_
                 * @example https://api.cmp.optimizely.com/v3/tasks/5d7f910551b00a722e0418830cee5534/steps/6d7f910551b00a722e0418830cee5564/sub-steps/9c7f910551b00a722e0418830cee2564/external-work
                 */
                external_work: string | null;
                /**
                 * @description URL of the substep
                 * @example https://api.cmp.optimizely.com/v3/tasks/5d7f910551b00a722e0418830cee5534/steps/6d7f910551b00a722e0418830cee5564/sub-steps/9c7f910551b00a722e0418830cee2564
                 */
                self: string;
                /**
                 * @description URL of the task
                 * @example https://api.cmp.optimizely.com/v3/tasks/5d7f910551b00a722e0418830cee5534
                 */
                task: string;
            };
            /**
             * @description Title of the substep
             * @example Sample step
             */
            title: string;
        };
        TaskSubStepCommentResponse: {
            /**
             * @description List of attachments of the comment
             * @example [
             *       {
             *         "id": "a113667245d111eb8945000c",
             *         "name": "sample.png",
             *         "url": "https://files.cmp.optimizely.com/download/96c314a645d111eb8945000c291b51d4?token="
             *       }
             *     ]
             */
            attachments: components["schemas"]["AttachmentResponse"][];
            /**
             * Format: date-time
             * @description Creation date and time of the task substep comment,  in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2020-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Unique identifier of the substep comment
             * @example 5fe38c39574b52a62a089239
             */
            id: string;
            /**
             * @description Determines if the comment is resolved in Optimizely CMP
             * @example false
             */
            is_resolved: boolean;
            /** @description Meta links */
            links?: {
                /**
                 * @description Creator of the substep comment
                 * @example https://api.cmp.optimizely.com/v3/users/5fe38eea574b52a62a08923a
                 */
                comment_by: string;
                /**
                 * @description URL of the substep comment
                 * @example https://api.cmp.optimizely.com/v3/tasks/5d7f910551b00a722e0418830cee5534/steps/6d7f910551b00a722e0418830cee5564/sub-steps/9c7f910551b00a722e0418830cee2564
                 */
                self: string;
                /**
                 * @description URL of the substep
                 * @example https://api.cmp.optimizely.com/v3/tasks/5d7f910551b00a722e0418830cee5534/steps/6d7f910551b00a722e0418830cee5564/sub-steps/9c7f910551b00a722e0418830cee2564
                 */
                sub_step: string;
                /**
                 * @description URL of the task
                 * @example https://api.cmp.optimizely.com/v3/tasks/5d7f910551b00a722e0418830cee5534
                 */
                task: string;
            };
            /**
             * Format: date-time
             * @description Last modification date and time of the task substep comment,  in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2020-10-06T14:15:30Z
             */
            modified_at: string;
            /**
             * @description Content of the comment
             * @example Well done!
             */
            value: string;
        };
        TaskSubStepCommentUpdateRequest: {
            /**
             * @description Updated comment value. Markdown is supported. To mention a user belonging to the organization, use `@[name](openapi-user-link)` format.
             * @example Updated comment mentioning @[Organization User](https://api.cmp.optimizely.com/v3/users/5fe38aeb574b52a62a089238)
             */
            value: string;
        };
        TaskSubStepRequest: {
            /**
             * @description ID of the assignee for the substep; cannot be passed together with `is_completed`, `is_in_progress` or `is_skipped`
             * @example 5fdf31d57f0d0e362e2b5908
             */
            assignee_id?: string | null;
            /**
             * @description Type of the assignee for the substep. If not provided with `assignee_id`, defaults to `user`.
             * @example user
             * @enum {string}
             */
            assignee_type?: "user" | "team";
            /**
             * @description Whether the substep should be completed; cannot be passed together with `is_in_progress`, `is_skipped` or `assignee_id`
             * @example true
             * @enum {boolean}
             */
            is_completed?: true;
            /**
             * @description Whether the substep should be progressed; cannot be passed together with `is_completed`, `is_skipped` or `assignee_id`
             * @example true
             * @enum {boolean}
             */
            is_in_progress?: true;
            /**
             * @description Whether the substep should be skipped; cannot be passed together with `is_in_progress`, `is_completed` or `assignee_id`
             * @enum {boolean}
             */
            is_skipped?: true;
        };
        TaskUpdateRequest: {
            /**
             * @description ID of the campaign to be associated with the task
             * @example 6684f31be89a526d030cfc0f
             */
            campaign_id?: string;
            /**
             * Format: date-time
             * @description Due date and time of the task
             * @example 2025-10-06T13:00:00Z
             */
            due_at?: string;
            /** @description Labels to associate */
            labels?: components["schemas"]["ResourceLabelRequest"][];
            /**
             * @description ID of the owner for the task
             * @example 6853b3e37ed5e42cff067434
             */
            owner_id?: string;
            /**
             * Format: date-time
             * @description Start date of the task
             * @example 2025-10-06T13:00:00Z
             */
            start_at?: string;
            /** @description Title of the task */
            title?: string;
            /**
             * @description ID of the workflow to be used in the task
             * @example 665d56034854d100ad7cb8fb
             */
            workflow_id?: string;
        };
        TaskUrlResponse: {
            /**
             * Format: date-time
             * @description Timestamp when the URL was added to the task
             * @example 2024-06-01T06:00:00Z
             */
            created_at: string;
            /**
             * @description Unique identifier of the URL added to the task
             * @example 692851fbbd8c9217e28f7ad1
             */
            id: string;
        };
        TaskVideo: {
            /**
             * Format: date-time
             * @description Date and time when the video was created,  in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Size of the video in bytes
             * @example 11244
             */
            file_size: number;
            /**
             * @description Unique identifier for the video
             * @example 5d7f910551b00a722e0418830cee6633
             */
            id: string;
            /** @description Labels associated with the video */
            labels: components["schemas"]["ResourceLabelResponse"][];
            /** @description Meta links */
            links: {
                /**
                 * @description URL to POST draft to the video
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/assets/5d7f910551b00a722e0418830cee6633/drafts
                 */
                drafts: string;
                /**
                 * @description URL of the video
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/videos/5d7f910551b00a722e0418830cee6633
                 */
                self: string;
                /**
                 * @description URL of the task that the video is associated with
                 * @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9
                 */
                task: string;
                /** @description Web URLs */
                web_urls: {
                    /**
                     * @description Web URL of the drafts of the video
                     * @example https://cmp.optimizely.com/cloud/task/5f857f30e1c4a2038d6179e9/video/5d7f910551b00a722e0418830cee6633
                     */
                    drafts: string;
                    /**
                     * @description Web URL of the video
                     * @example https://cmp.optimizely.com/cloud/taskv3/5f857f30e1c4a2038d6179e9?contentTabGuid=5d7f910551b00a722e0418830cee6633
                     */
                    self: string;
                    /**
                     * @description Web URL of the task that the video is associated with
                     * @example https://cmp.optimizely.com/cloud/taskv3/5f857f30e1c4a2038d6179e9
                     */
                    task: string;
                };
            };
            /**
             * @description MIME type of the video
             * @example video/mp4
             */
            mime_type: string;
            /**
             * Format: date-time
             * @description Date and time of the most recent modification of the video, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2019-10-07T13:15:30Z
             */
            modified_at: string;
            /**
             * @description Title of the video
             * @example product.mp4
             */
            title: string;
            /**
             * @description Download URL of the video
             * @example https://videos.cmp.optimizely.com/03a747babe81ceb55763fa085bqa20dc
             */
            url: string;
        };
        Team: {
            /**
             * @description Unique identifier of the team
             * @example 5d7f910551b00a722e0418830cee6633
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the team
                 * @example https://api.cmp.optimizely.com/v3/teams/5d7f910551b00a722e0418830cee6633
                 */
                self: string;
            };
            /**
             * @description Name of the team
             * @example Sample Team
             */
            name: string;
        };
        TeamWithUsers: {
            /**
             * @description Unique identifier of the team
             * @example 5d7f910551b00a722e0418830cee6633
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the team
                 * @example https://api.cmp.optimizely.com/v3/teams/5d7f910551b00a722e0418830cee6633
                 */
                self: string;
            };
            /**
             * @description Name of the team
             * @example Sample Team
             */
            name: string;
            /** @description List of users in the team */
            users: {
                /**
                 * @description Email address of the user
                 * @example john.doe@example.com
                 */
                email: string;
                /**
                 * @description Unique identifier of the user
                 * @example 5d7f910551b00a722e0418830cee6633
                 */
                id: string;
                /**
                 * @description Name of the user
                 * @example John Doe
                 */
                name: string;
            }[];
        };
        TemplateChoiceFormField: components["schemas"]["BaseTemplateFormField"] & {
            logic_rules: components["schemas"]["TemplateLogicRule"];
            /** @description Type-specific meta-information */
            type_specific_meta: {
                /** @description Choices for the field */
                choices: components["schemas"]["ResourceChoiceResponse"][];
                /**
                 * @description Whether multiple selection is allowed
                 * @example true
                 */
                is_multi_select: boolean;
            };
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "checkbox" | "dropdown" | "label" | "radio_button";
        };
        TemplateCurrencyNumberFormField: components["schemas"]["BaseTemplateFormField"] & {
            logic_rules: components["schemas"]["TemplateLogicRule"];
            /** @description Type-specific meta-information */
            type_specific_meta: {
                /**
                 * @description Currency code of the field
                 * @example USD
                 */
                currency_code: string;
                /**
                 * @description Decimal places allowed for the field
                 * @example 2
                 */
                decimal_places: number | null;
                /**
                 * @description Whether the field has thousand separators
                 * @example true
                 */
                has_thousand_separator: boolean;
            };
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "currency_number";
        };
        TemplateDefaultFormField: components["schemas"]["BaseTemplateFormField"] & {
            logic_rules: components["schemas"]["TemplateLogicRule"];
            /**
             * @description Type-specific meta-information
             * @example null
             * @enum {object|null}
             */
            type_specific_meta: never | null;
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "brief" | "date" | "file" | "richtext" | "section" | "text" | "text_area";
        };
        TemplateInstructionFormField: components["schemas"]["BaseTemplateFormField"] & {
            logic_rules: components["schemas"]["TemplateLogicRule"];
            /** @description Type-specific meta-information */
            type_specific_meta: {
                /**
                 * @description Description of the field
                 * @example just a message to inform something
                 */
                description: string;
            };
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "instruction";
        };
        TemplateListResponse: {
            /** @description List of templates */
            data: (components["schemas"]["BaseTemplateResponse"] & {
                /** @description Meta links */
                links: {
                    /**
                     * @description URL of the template
                     * @example https://api.cmp.optimizely.com/v3/templates/9119a313057e401189407116fcd3aa24
                     */
                    self: string;
                };
            })[];
            pagination: components["schemas"]["Pagination"] & {
                /** @example https://api.cmp.optimizely.com/v3/templates?offset=10&page_size=10 */
                next?: string | null;
            };
        };
        /** @description List of logic rules */
        TemplateLogicRule: {
            action: components["schemas"]["TemplateLogicRuleAction"];
            condition: components["schemas"]["TemplateLogicRuleCondition"];
        }[];
        /** @description Action to perform for the rule */
        TemplateLogicRuleAction: {
            /** @description Target field for the action */
            target_field: {
                /**
                 * @description Identifier of the target field
                 * @example 70ca0e8a954911ecb4be0242ac160014
                 */
                identifier: string;
            };
            /**
             * @description Type of the action
             * @example show_values
             * @enum {string}
             */
            type: "jump_to" | "show_values";
            /** @description Values for the action */
            values: string[];
        };
        /** @description Condition of the rule */
        TemplateLogicRuleCondition: {
            /**
             * @description Operator of the condition
             * @example any_of
             * @enum {string}
             */
            operator: "any_of" | "equal" | "not_equal";
            /** @description Values of the condition */
            values: string[];
        };
        TemplatePercentageNumberFormField: components["schemas"]["BaseTemplateFormField"] & {
            logic_rules: components["schemas"]["TemplateLogicRule"];
            /** @description Type-specific meta-information */
            type_specific_meta: {
                /**
                 * @description Decimal places allowed for the field
                 * @example 2
                 */
                decimal_places: number | null;
            };
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "percentage_number";
        };
        TemplateResponse: components["schemas"]["BaseTemplateResponse"] & {
            /** @description List of form fields */
            form_fields: (components["schemas"]["TemplateDefaultFormField"] | components["schemas"]["TemplateCurrencyNumberFormField"] | components["schemas"]["TemplateSimpleNumberFormField"] | components["schemas"]["TemplatePercentageNumberFormField"] | components["schemas"]["TemplateChoiceFormField"] | components["schemas"]["TemplateInstructionFormField"])[];
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the template
                 * @example https://api.cmp.optimizely.com/v3/templates/9119a313057e401189407116fcd3aa24
                 */
                self: string;
            };
        };
        TemplateSimpleNumberFormField: components["schemas"]["BaseTemplateFormField"] & {
            logic_rules: components["schemas"]["TemplateLogicRule"];
            /** @description Type-specific meta-information */
            type_specific_meta: {
                /**
                 * @description Decimal places allowed for the field
                 * @example 2
                 */
                decimal_places: number | null;
                /**
                 * @description Whether the field has thousand separators
                 * @example true
                 */
                has_thousand_separator: boolean;
            };
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "simple_number";
        };
        TextFieldDefinition: {
            core: components["schemas"]["CoreFieldDef"];
            /** @description Default values of the field */
            default_values?: string[];
            /** @description Maximum length of the field */
            max_length?: number;
            /** @description Minimum length of the field */
            min_length: number;
            /** @description Validation pattern of the field */
            validation_pattern: string;
        };
        TextFieldValueModel: {
            /** @description Order index of the field value */
            order_index?: number;
            /** @description Value of the field value */
            text_value: string;
        };
        TextTypeFormFieldRequest: components["schemas"]["BaseFormFieldRequest"] & {
            values?: string[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "richtext" | "text" | "text_area";
        };
        TextTypeFormFieldResponse: components["schemas"]["BaseFormFieldResponse"] & {
            values?: string[];
        } & {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "richtext" | "text" | "text_area";
        };
        URLFieldValueModel: {
            /** @description Order index of the field value */
            order_index?: number;
            /**
             * Format: uri
             * @description URL of the field value
             */
            url: string;
        };
        UserListResponse: {
            /** @description List of users */
            data: (components["schemas"]["UserListResponseItem"] & {
                /** @description Meta links */
                links: {
                    /**
                     * @description URL of the user
                     * @example https://api.cmp.optimizely.com/v3/users/9119a313057e401189407116fcd3aa24
                     */
                    self: string;
                };
            })[];
            pagination: components["schemas"]["Pagination"] & {
                /** @example https://api.cmp.optimizely.com/v3/users?offset=10&page_size=10 */
                next?: string | null;
            };
        };
        UserListResponseItem: {
            /**
             * @description First name of the user
             * @example John
             */
            first_name: string;
            /**
             * @description Full name of the user
             * @example John Doe
             */
            full_name: string;
            /**
             * @description Unique identifier of the user
             * @example 5fe4925ef8a9378056bf4e37
             */
            id: string;
            /**
             * @description URL of the profile picture of the user - null if the user has not provided an image
             * @example https://images.cmp.optimizely.com/6a485f72ba424f9397c71ddf0cfb8f59
             */
            image_url: string | null;
            /**
             * @description Last name of the user
             * @example Doe
             */
            last_name: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the user
                 * @example https://api.cmp.optimizely.com/v3/users/5fe4925ef8a9378056bf4e37
                 */
                self: string;
            };
            /** @description List of roles assigned to the user */
            roles: {
                /**
                 * @description Name of the role
                 * @example Creator
                 */
                name: string;
            }[];
        };
        UserResponse: {
            /**
             * @description Email of the user – null unless the application is configured to expose the user email
             * @example john.doe@example.com
             */
            email: string | null;
            /**
             * @description First name of the user
             * @example John
             */
            first_name: string;
            /**
             * @description Full name of the user
             * @example John Doe
             */
            full_name: string;
            /**
             * @description Unique identifier of the user
             * @example 5fe4925ef8a9378056bf4e37
             */
            id: string;
            /**
             * @description URL of the profile picture of the user – null if the user has not provided an image
             * @example https://images.cmp.optimizely.com/6a485f72ba424f9397c71ddf0cfb8f59
             */
            image_url: string | null;
            /**
             * @description Last name of the user
             * @example Doe
             */
            last_name: string;
            /** @description Meta links */
            links?: {
                /**
                 * @description URL of the user
                 * @example https://api.cmp.optimizely.com/v3/users/5fe4925ef8a9378056bf4e37
                 */
                self?: string;
            };
        };
        ValidationError: {
            /** @description Locations where the error occurred */
            loc: string[];
            /** @description Message describing the error */
            msg: string;
            /**
             * Error Type
             * @description Type of the error
             */
            type: string;
        };
        VersionedContentTypeModel: {
            /** @description Indicates whether the content type is a component */
            component: boolean;
            /** @description Unique identifier of the content type */
            content_type_guid: string;
            /**
             * Format: date-time
             * @description Date and time on which the content type was created, in ISO 8601 UTC format
             */
            created_at: string;
            /** @description Unique identifier of the user who created the content type */
            created_by: string;
            /** @description Description of the content type */
            description?: string;
            /** @description Disabled status of the content type */
            disabled?: boolean;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the content type
                 * @example https://api.cmp.optimizely.com/v3/structured-content/content-types/9fda53cf66a14fd487251480ca695c7b
                 */
                self?: string;
                /**
                 * @description URL of the content type versions
                 * @example https://api.cmp.optimizely.com/v3/structured-content/content-types/9fda53cf66a14fd487251480ca695c7b/versions
                 */
                versions?: string;
            };
            /** @description Name of the content type */
            name: string;
            /** @description Source of the content type */
            source?: string;
            /** @description Source of the content type */
            source_id?: string;
            /** @description Source metadata of the content type */
            source_metadata?: string;
            /** @description Thumbnail GUID of the content type */
            thumbnail_guid?: string;
            /**
             * Format: date-time
             * @description Date and time on which the content type was last updated, in ISO 8601 UTC format
             */
            updated_at: string;
            /** @description Unique identifier of the user who last updated the content type */
            updated_by: string;
            version: components["schemas"]["SCContentTypeVersion"];
        };
        WorkRequestApprovedAssetResponse: {
            /** @description Content of the approved asset */
            content_type: {
                /**
                 * @description Type of the content. 'html_body for 'article' and 'url' for `image`, `raw file`, and `video`
                 * @example url
                 * @enum {string}
                 */
                type: "html_body" | "url";
                /**
                 * @description Content of the asset. 'html_body for 'article' and 'url' for `image`, `raw file`, and `video`
                 * @example http://images.cmp.optimizely.com/Zz0xODQ3NDU3Y2Y2YmYzOTlmNjkyOTgyZDY3Y2I3YWM2OA==S
                 */
                value: string;
            };
            /**
             * @description Unique identifier of the work request's approved asset
             * @example 74f12a064b1611eda2fe02420ac8001d
             */
            id: string;
            /**
             * @description Mime type of the work request's approved asset
             * @example image/png
             */
            mime_type: string;
            /**
             * @description Title of the work request's approved asset
             * @example Nature.png
             */
            title: string;
            /**
             * @description Type of the work request's approved asset
             * @example image
             * @enum {string}
             */
            type: "article" | "image" | "video" | "raw_file";
        };
        /** @description Brief from uploaded attachment */
        WorkRequestAttachmentBriefRequest: {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "attachment_brief";
            value: {
                /**
                 * @description Unique identifier of the file upload session. This should be the value of `upload_meta_fields.key` retrieved using the `/v3/upload-url` endpoint
                 * @example 6dzf4ds4fds4f4564dzsfd
                 */
                key: string;
                /**
                 * @description Name of the attachment brief
                 * @example Sample-name
                 */
                name: string;
            };
        };
        /** @description Brief from uploaded attachment */
        WorkRequestAttachmentBriefResponse: {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "attachment_brief";
            value: {
                /**
                 * @description Unique identifier of the attachment brief
                 * @example 6dzf4ds4fds4f4564dzsfd
                 */
                id: string;
                /**
                 * @description Name of the attachment brief
                 * @example Sample-name
                 */
                name: string;
                /**
                 * @description URL of the attachment brief
                 * @example https://files.cmp.optimizely.com/download/8nv8svsdnivjusdijuvisdv?token=token
                 */
                url: string;
            };
        };
        WorkRequestCampaignRequest: {
            /**
             * @description Description of the campaign
             * @example The awesome campaign's description
             */
            description?: string | null;
            /**
             * Format: date-time
             * @description End date of the campaign, in ISO 8601 UTC format (2022-08-31T06:00:00Z)
             * @example 2022-08-31T06:00:00Z
             */
            end_date?: string | null;
            /**
             * @description Unique identifier of the campaign owner
             * @example 8q7f910551b00a722e0418830cee6612
             */
            owner_id?: string;
            /**
             * @description Unique identifier of a campaign
             * @example 6679686ade80bdf1dfda5ec3ab32ddf
             */
            parent_campaign_id?: string;
            /**
             * Format: date-time
             * @description Start date of the campaign, in ISO 8601 UTC format (2022-08-31T06:00:00Z)
             * @example 2022-08-31T06:00:00Z
             */
            start_date?: string | null;
            /**
             * @description Title of the campaign
             * @example The awesome campaign
             */
            title: string;
        };
        WorkRequestCampaignResponse: {
            /**
             * @description Description of the campaign
             * @example The awesome campaign's description
             */
            description: string | null;
            /**
             * Format: date-time
             * @description End date of the campaign, in ISO 8601 UTC format (2022-08-31T06:00:00Z)
             * @example 2022-08-31T06:00:00Z
             */
            end_date: string | null;
            /**
             * @description Unique identifier of the campaign
             * @example 8q7f910551b00a722e0418830cee6612
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the campaign brief
                 * @example https://api.cmp.optimizely.com/v3/campaigns/8q7f910551b00a722e0418830cee6612/brief
                 */
                brief: string | null;
                /**
                 * @description URL of the parent campaign
                 * @example https://api.cmp.optimizely.com/v3/campaigns/6679686ade80bdf1dfda5ec3ab32ddf
                 */
                parent_campaign: string | null;
                /**
                 * @description URL of the campaign
                 * @example https://api.cmp.optimizely.com/v3/campaigns/8q7f910551b00a722e0418830cee6612
                 */
                self: string;
            };
            /**
             * @description Unique identifier of the campaign owner
             * @example 8q7f910551b00a722e0418830cee6612
             */
            owner_id: string;
            /**
             * @description Unique identifier of the parent campaign
             * @example 6679686ade80bdf1dfda5ec3ab32ddf
             */
            parent_campaign_id: string;
            /**
             * @description Reference identifier of the campaign
             * @example CPN-1589
             */
            reference_id: string;
            /**
             * Format: date-time
             * @description Start date of the campaign, in ISO 8601 UTC format (2022-08-31T06:00:00Z)
             * @example 2022-08-31T06:00:00Z
             */
            start_date: string | null;
            /**
             * @description Title of the campaign
             * @example The awesome campaign
             */
            title: string;
        };
        WorkRequestCommentResponse: {
            /**
             * @description List of attachments of the comment
             * @example [
             *       {
             *         "id": "a113667245d111eb8945000c",
             *         "name": "sample.png",
             *         "url": "https://files.cmp.optimizely.com/download/96c314a645d111eb8945000c291b51d4?token="
             *       }
             *     ]
             */
            attachments: components["schemas"]["AttachmentResponse"][];
            /**
             * @description Unique identifier of the user who posted the comment
             * @example 62a8455e29fed3045d9c7a8b
             */
            comment_by: string;
            /**
             * Format: date-time
             * @description Creation date and time of the work request comment, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2020-10-06T13:15:30Z
             */
            created_at: string;
            /**
             * @description Unique identifier of the work request comment
             * @example 5fe38c39574b52a62a089239
             */
            id: string;
            /**
             * @description Determines if the comment is resolved
             * @example false
             */
            is_resolved: boolean;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the user who posted the comment
                 * @example https://api.cmp.optimizely.com/v3/users/62a8455e29fed3045d9c7a8b
                 */
                comment_by: string;
                /**
                 * @description URL of the work request
                 * @example https://api.cmp.optimizely.com/v3/work-requests/da80cfd7bbf84959a8a981acbad996b3
                 */
                work_request: string;
            };
            /**
             * Format: date-time
             * @description Last modification date and time of the work request comment, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
             * @example 2020-10-06T14:15:30Z
             */
            modified_at: string;
            /**
             * @description Unique identifier of the parent comment
             * @example 5fe38c39574b52a62a089238
             */
            parent_comment_id: string | null;
            /**
             * @description Content of the comment
             * @example Well done!
             */
            value: string;
        };
        WorkRequestCreateRequest: {
            /** @description List of assigned users to the work request */
            assignees?: string[];
            /** @description Form fields from the work request template to populate in the work request. **NOTE** You must provide values for all the required form fields. */
            form_fields: (components["schemas"]["BriefTypeFormFieldRequest"] | components["schemas"]["DateTypeFormFieldRequest"] | components["schemas"]["FileTypeFormFieldRequest"] | components["schemas"]["MultiChoiceTypeFormFieldRequest"] | components["schemas"]["NumberTypeFormFieldRequest"] | components["schemas"]["RadioButtonTypeFormFieldRequest"] | components["schemas"]["TextTypeFormFieldRequest"])[];
            /**
             * @description The template on which the work request should be created
             * @example 1265ca474ba987cc925164aa7
             */
            template_id: string;
        };
        WorkRequestFormFieldUpdateRequest: components["schemas"]["WorkRequestRequestFormFieldBriefTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldCheckboxTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldCurrencyNumberTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldDateTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldDropdownTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldFileTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldLabelTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldPercentageNumberTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldRadioButtonTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldRichtextTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldSimpleNumberTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldTextAreaTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldTextTypePayload"];
        WorkRequestRelatedResourceResponse: {
            /**
             * Format: date-time
             * @description Date and time on which the work request related resource was created, in ISO 8601 UTC format (2022-04-30T06:00:00.000Z)
             * @example 2022-04-30T06:00:00.000Z
             */
            created_at: string;
            /**
             * @description Unique identifier of the related resource
             * @example 5fe4925ef8a9378056bf4e37
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the related resource
                 * @example https://api.cmp.optimizely.com/v3/tasks/5fe4925ef8a9378056bf4e37
                 */
                self: string | null;
            };
            /**
             * @description Related resource relation type
             * @example started
             * @enum {string}
             */
            relation_type: "started" | "linked";
            /**
             * @description Related resource type. task/campaign/event etc
             * @example task
             */
            resource_type: string;
        };
        WorkRequestRequestFormFieldBriefTypePayload: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "brief";
            /** @description Brief value payload */
            values: (components["schemas"]["WorkRequestTextBriefRequest"] | components["schemas"]["WorkRequestAttachmentBriefRequest"])[];
        };
        WorkRequestRequestFormFieldBriefTypeResponse: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "brief";
            /** @description Brief value payload */
            values: (components["schemas"]["WorkRequestTextBriefResponse"] | components["schemas"]["WorkRequestAttachmentBriefResponse"])[];
        };
        WorkRequestRequestFormFieldCheckboxTypePayload: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "checkbox";
            /**
             * @description Unique identifiers of the choices
             * @example [
             *       "63342902e84c7544973c07ce",
             *       "633429b2e84c7544973c07e7"
             *     ]
             */
            values: string[];
        };
        WorkRequestRequestFormFieldCheckboxTypeResponse: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "checkbox";
            /**
             * @description Values for the form field
             * @example [
             *       {
             *         "id": "63342902e84c7544973c07ce",
             *         "name": "Option 1"
             *       },
             *       {
             *         "id": "633429b2e84c7544973c07e7",
             *         "name": "Option 2"
             *       }
             *     ]
             */
            values: {
                /** @description Unique identifier of the choice */
                id: string;
                /** @description Name of the choice */
                name: string;
            }[];
        };
        WorkRequestRequestFormFieldCurrencyNumberTypePayload: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "currency_number";
            /**
             * @description Currency amount for the form field
             * @example [
             *       "123.45"
             *     ]
             */
            values: string[];
        };
        WorkRequestRequestFormFieldDateTypePayload: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "date";
            /**
             * Format: date-time
             * @description Date and time, in ISO 8601 UTC format
             * @example [
             *       "2022-04-30T06:00:00.00Z"
             *     ]
             */
            values: string[];
        };
        WorkRequestRequestFormFieldDropdownTypePayload: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "dropdown";
            /**
             * @description Unique identifier of the choice for the form field
             * @example [
             *       "63342902e84c7544973c07ce"
             *     ]
             */
            values: string[];
        };
        WorkRequestRequestFormFieldDropdownTypeResponse: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "dropdown";
            /**
             * @description Value for the form field
             * @example [
             *       {
             *         "id": "63342902e84c7544973c07ce",
             *         "name": "Option 1"
             *       }
             *     ]
             */
            values: {
                /** @description Unique identifier of the choice */
                id: string;
                /** @description Name of the choice */
                name: string;
            }[];
        };
        WorkRequestRequestFormFieldFileTypePayload: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "file";
            /**
             * @description uploaded file values
             * @example [
             *       {
             *         "key": "123456/789654",
             *         "name": "My File"
             *       }
             *     ]
             */
            values: {
                /** @description Unique identifier of the file upload session. This is the `upload_meta_fields.key` field was retrieved from the `/v3/upload-url` endpoint. */
                key: string;
                /** @description Name for the file */
                name: string;
            }[];
        };
        WorkRequestRequestFormFieldFileTypeResponse: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "file";
            /**
             * @description Values for the form field
             * @example [
             *       {
             *         "id": "a113667245d111eb8945000c",
             *         "name": "My File",
             *         "url": "https://files.cmp.optimizely.com/download/96c314a645d111eb8945000c291b51d4?token="
             *       }
             *     ]
             */
            values: {
                /** @description Unique identifier of the file */
                id: string;
                /** @description Name for the file */
                name: string;
                /** @description Download URL of the file */
                url: string;
            }[];
        };
        WorkRequestRequestFormFieldLabelTypePayload: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "label";
            /**
             * @description Unique identifiers of label values for the form field
             * @example [
             *       "63342902e84c7544973c07ce"
             *     ]
             */
            values: string[];
        };
        WorkRequestRequestFormFieldLabelTypeResponse: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "label";
            /**
             * @description Values for the form field
             * @example [
             *       {
             *         "id": "63342902e84c7544973c07ce",
             *         "name": "Label 1"
             *       }
             *     ]
             */
            values: {
                /** @description Unique identifiers of label value */
                id: string;
                /** @description Name of the label value */
                name: string;
            }[];
        };
        WorkRequestRequestFormFieldPercentageNumberTypePayload: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "percentage_number";
            /**
             * @description Value for the form field
             * @example [
             *       "123.456"
             *     ]
             */
            values: string[];
        };
        WorkRequestRequestFormFieldRadioButtonTypePayload: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "radio_button";
            /**
             * @description Unique identifier of the choice for the form field
             * @example [
             *       "63342902e84c7544973c07ce"
             *     ]
             */
            values: string[];
        };
        WorkRequestRequestFormFieldRadioButtonTypeResponse: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "radio_button";
            /**
             * @description Value for the form field
             * @example [
             *       {
             *         "id": "63342902e84c7544973c07ce",
             *         "name": "Choice 1"
             *       }
             *     ]
             */
            values: {
                /** @description Unique identifier of the choice */
                id?: string;
                /** @description Name of the choice value */
                name?: string;
            }[];
        };
        WorkRequestRequestFormFieldRichtextTypePayload: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "richtext";
            /**
             * @description Value of the form field
             * @example [
             *       "<h1>hello world</h1>"
             *     ]
             */
            values: string[];
        };
        WorkRequestRequestFormFieldSimpleNumberTypePayload: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "simple_number";
            /**
             * @description Value of the form field
             * @example [
             *       "123.456"
             *     ]
             */
            values: string[];
        };
        WorkRequestRequestFormFieldTextAreaTypePayload: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "text_area";
            /**
             * @description Value of the form field
             * @example [
             *       "Hello world"
             *     ]
             */
            values: string[];
        };
        WorkRequestRequestFormFieldTextTypePayload: {
            /**
             * @description Type of the form field (enum property replaced by openapi-typescript)
             * @enum {string}
             */
            type: "text";
            /**
             * @description Value of the form field
             * @example [
             *       "Hello world"
             *     ]
             */
            values: string[];
        };
        WorkRequestRequestFormFieldUpdateResponse: components["schemas"]["WorkRequestRequestFormFieldBriefTypeResponse"] | components["schemas"]["WorkRequestRequestFormFieldCheckboxTypeResponse"] | components["schemas"]["WorkRequestRequestFormFieldCurrencyNumberTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldDateTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldDropdownTypeResponse"] | components["schemas"]["WorkRequestRequestFormFieldFileTypeResponse"] | components["schemas"]["WorkRequestRequestFormFieldLabelTypeResponse"] | components["schemas"]["WorkRequestRequestFormFieldPercentageNumberTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldRadioButtonTypeResponse"] | components["schemas"]["WorkRequestRequestFormFieldRichtextTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldSimpleNumberTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldTextAreaTypePayload"] | components["schemas"]["WorkRequestRequestFormFieldTextTypePayload"];
        WorkRequestResponse: {
            /** @description List of assigned users to the work request */
            assignees: {
                /**
                 * @description Full name of the assignee
                 * @example John Doe
                 */
                full_name: string;
                /**
                 * @description Unique identifier of the assignee
                 * @example 6asdcnfnfdsfdsfd54sgfd54sf
                 */
                id: string;
                /**
                 * @description URL of the profile picture of the assignee. `null` if the assignee has not provided an image
                 * @example https://images.cmp.optimizely.com/6a485f72ba424f9397c71ddf0cfb8f59
                 */
                image_url: string | null;
                /** @description Meta links */
                links: {
                    /**
                     * @description URL of the assigned user
                     * @example https://api.cmp.optimizely.com/v3/users/6asdcnfnfdsfdsfd54sgfd54sf
                     */
                    self: string;
                };
                /**
                 * @description Type of the assignee
                 * @example user
                 * @enum {string}
                 */
                type: "user" | "team";
            }[];
            /**
             * Format: date-time
             * @description Date and time on which the work request was created, in ISO 8601 UTC format (2022-04-30T06:00:00.000Z)
             * @example 2022-04-30T06:00:00.000Z
             */
            created_at: string;
            /** @description Creator of the work request */
            created_by: {
                /**
                 * @description Unique identifier of the creator of the work request
                 * @example 6asdcnfnfdsfdsfd54sgfd54sf
                 */
                id: string;
                /** @description Meta links */
                links: {
                    /**
                     * @description URL of the creator
                     * @example https://api.cmp.optimizely.com/v3/users/6asdcnfnfdsfdsfd54sgfd54sf
                     */
                    self: string;
                };
            };
            /**
             * @description Populated form fields following the template
             * @example [
             *       {
             *         "identifier": "title",
             *         "type": "text",
             *         "values": [
             *           "Some title"
             *         ]
             *       },
             *       {
             *         "identifier": "brief",
             *         "type": "brief",
             *         "values": [
             *           {
             *             "type": "attachment_brief",
             *             "value": {
             *               "id": "attachment-1",
             *               "name": "Attachment",
             *               "url": "http://images.cmp.optimizely.com/attachment-1?token=jwt.auth.token"
             *             }
             *           }
             *         ]
             *       },
             *       {
             *         "identifier": "630c486a90546926f034ea28",
             *         "type": "richtext",
             *         "values": [
             *           "<b>rich</b> text"
             *         ]
             *       },
             *       {
             *         "identifier": "630c479690546926f034ea13",
             *         "type": "label",
             *         "values": [
             *           {
             *             "id": "630c479690546926f034ea15",
             *             "name": "value-1"
             *           }
             *         ]
             *       },
             *       {
             *         "identifier": "630c47be90546926f034ea17",
             *         "type": "dropdown",
             *         "values": [
             *           {
             *             "id": "630c47be90546926f034ea18",
             *             "name": "option 1"
             *           },
             *           {
             *             "id": "630c47be90546926f034ea20",
             *             "name": "option 3"
             *           }
             *         ]
             *       },
             *       {
             *         "identifier": "630c482b90546926f034ea22",
             *         "type": "simple_number",
             *         "values": [
             *           "1234.345678"
             *         ]
             *       },
             *       {
             *         "identifier": "630c484b90546926f034ea23",
             *         "type": "percentage_number",
             *         "values": [
             *           "23"
             *         ]
             *       },
             *       {
             *         "identifier": "630c481b90546926f034ea21",
             *         "type": "currency_number",
             *         "values": [
             *           "123.56"
             *         ]
             *       },
             *       {
             *         "identifier": "630c485e90546926f034ea24",
             *         "type": "checkbox",
             *         "values": [
             *           {
             *             "id": "630c485e90546926f034ea25",
             *             "name": "choice 1"
             *           },
             *           {
             *             "id": "630c485e90546926f034ea26",
             *             "name": "choice 6"
             *           }
             *         ]
             *       },
             *       {
             *         "identifier": "630c47d290546926f034ea1b",
             *         "type": "radio_button",
             *         "values": [
             *           {
             *             "id": "630c47d290546926f034ea1c",
             *             "name": "choice 2"
             *           }
             *         ]
             *       },
             *       {
             *         "identifier": "creative_assets",
             *         "type": "file",
             *         "values": [
             *           {
             *             "id": "creative-asset-1",
             *             "name": "Creative Asset 1",
             *             "url": "http://images.cmp.optimizely.com/creative-asset-1?token=jwt.auth.token"
             *           },
             *           {
             *             "id": "creative-asset-2",
             *             "name": "Creative Asset 2",
             *             "url": "http://images.cmp.optimizely.com/creative-asset-2?token=jwt.auth.token"
             *           }
             *         ]
             *       },
             *       {
             *         "identifier": "due_date",
             *         "type": "date",
             *         "values": [
             *           "2022-02-20T12:02:03Z"
             *         ]
             *       },
             *       {
             *         "identifier": "media_links",
             *         "type": "text_area",
             *         "values": [
             *           "some text\nmore text"
             *         ]
             *       }
             *     ]
             */
            form_fields: (components["schemas"]["BriefTypeFormFieldResponse"] | components["schemas"]["DateTypeFormFieldResponse"] | components["schemas"]["FileTypeFormFieldResponse"] | components["schemas"]["MultiChoiceTypeFormFieldResponse"] | components["schemas"]["NumberTypeFormFieldResponse"] | components["schemas"]["RadioButtonTypeFormFieldResponse"] | components["schemas"]["TextTypeFormFieldResponse"])[];
            /**
             * @description Unique identifier of the work request
             * @example 5fe4925ef8a9378056bf4e37
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL to add attachments to work request
                 * @example https://api.cmp.optimizely.com/v3/work-requests/5fe4925ef8a9378056bf4e37/attachments
                 */
                attachments: string;
                /**
                 * @description URL to create a campaign from the work request
                 * @example https://api.cmp.optimizely.com/v3/work-requests/5fe4925ef8a9378056bf4e37/campaigns
                 */
                campaigns: string;
                /**
                 * @description URL to create or get work request comments
                 * @example https://api.cmp.optimizely.com/v3/work-requests/5fe4925ef8a9378056bf4e37/comments
                 */
                comments: string;
                /**
                 * @description URL to add creative assets to the work request
                 * @example https://api.cmp.optimizely.com/v3/work-requests/5fe4925ef8a9378056bf4e37/creative-assets
                 */
                creative_assets: string;
                /**
                 * @description URL to get related resources
                 * @example https://api.cmp.optimizely.com/v3/work-requests/5fe4925ef8a9378056bf4e37/related-resources
                 */
                related_resources: string;
                /**
                 * @description URL of the work request
                 * @example https://api.cmp.optimizely.com/v3/work-requests/5fe4925ef8a9378056bf4e37
                 */
                self: string;
                /**
                 * @description URL to create a task from the work request
                 * @example https://api.cmp.optimizely.com/v3/work-requests/5fe4925ef8a9378056bf4e37/tasks
                 */
                tasks: string;
            };
            /**
             * Format: date-time
             * @description Date and time on which the work request was last modified, in ISO 8601 UTC format (2022-04-30T08:00:00.000Z)
             * @example 2022-04-30T08:00:00.000Z
             */
            modified_at: string;
            /**
             * @description Priority of the work request
             * @example Medium
             * @enum {string}
             */
            priority: "Low" | "Medium" | "High";
            /**
             * @description Reference identifier of the work request
             * @example WRQ-23
             */
            reference_id: string;
            /**
             * @description Current status of the work request
             * @example Accepted
             * @enum {string}
             */
            status: "Accepted" | "Declined" | "Submitted" | "Completed";
            /** @description The template, based on the work request that was created */
            template: {
                /**
                 * @description Unique identifier of the template
                 * @example 1265ca474ba987cc925164aa7
                 */
                id: string;
                /** @description Meta links */
                links: {
                    /**
                     * @description URL of the template
                     * @example https://api.cmp.optimizely.com/v3/templates/9119a313057e401189407116fcd3aa24
                     */
                    self: string;
                };
                /**
                 * @description Title of the template
                 * @example Sample template title
                 */
                title: string;
            };
        };
        WorkRequestTaskRequest: {
            /**
             * @description List of asset ids which will not inherit any fields.
             * @example [
             *       "74f12a064b1611eda2fe02420ac8001d"
             *     ]
             */
            asset_ids_without_field_inheritance?: string[];
            /**
             * @description Unique identifier of parent campaign
             * @example 8q7f910551b00a722e0418830cee6612
             */
            campaign_id?: string | null;
            /**
             * Format: date-time
             * @description Due date and time of the task, in ISO 8601 UTC format (2020-02-10T10:40:45Z). The `due_at` field is required if `workflow_id` is present and the workflow does not have step duration. If `workflow_id` is present and the workflow has step duration, then either `start_at` or `due_at` is required.
             * @example 2022-12-31T06:00:00Z
             */
            due_at?: string | null;
            /**
             * @description Specifies the source from which to inherit fields.
             *
             *     Supported values:
             *     - `work_request` (default) : Fields are created from the work request
             *     - `workflow` : Fields are inherited from the workflow specified by `workflow_id`
             * @default work_request
             * @example work_request
             */
            inherit_fields_from: string;
            /**
             * @description Unique identifier of the task owner. If passed as null, then the owner will be the token user.
             * @example 8q7f910551b00a722e0418830cee6612
             */
            owner_id?: string | null;
            /**
             * Format: date-time
             * @description Start date and time of the task, in ISO 8601 UTC format (2020-02-10T10:40:45Z). If `workflow_id` is present and the workflow has step duration, then either `start_at` or `due_at` is required. If both `start_at` and `due_at` is provided with a `workflow_id` that has step duration, then the `start_at` field will be ignored.
             * @example 2022-12-01T06:00:00Z
             */
            start_at?: string | null;
            /**
             * @description Title of the task. If not provided. it will be populated by the work request title.
             * @example The awesome task
             */
            title?: string;
            /**
             * @description Unique identifier of workflow
             * @example 8q7f910551b00a722e0418830cee6612
             */
            workflow_id?: string | null;
        };
        WorkRequestTaskResponse: {
            /**
             * Format: date-time
             * @description Due date of the task, in ISO 8601 UTC format (2022-12-31T06:00:00Z)
             * @example 2022-12-31T06:00:00Z
             */
            due_at: string | null;
            /**
             * @description Unique identifier of the task
             * @example 8q7f910551b00a722e0418830cee6612
             */
            id: string;
            /** @description Meta links */
            links: {
                /**
                 * @description URL of the task owner
                 * @example https://api.cmp.optimizely.com/v3/users/8q7f910551b00a722e0418830cee6612
                 */
                owner?: string | null;
                /**
                 * @description URL of the task
                 * @example https://api.cmp.optimizely.com/v3/tasks/8q7f910551b00a722e0418830cee6612
                 */
                self: string;
            };
            /**
             * @description Unique identifier of the owner
             * @example 8q7f910551b00a722e0418830cee6612
             */
            owner_id: string;
            /**
             * @description Reference Id of the task
             * @example TSK-1111
             */
            reference_id?: string;
            /**
             * Format: date-time
             * @description Start date of the task, in ISO 8601 UTC format (2022-12-01T06:00:00Z)
             * @example 2022-12-01T06:00:00Z
             */
            start_at: string | null;
            /**
             * @description Title of the task
             * @example The awesome task
             */
            title: string;
        };
        /** @description Written brief */
        WorkRequestTextBriefRequest: {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "text_brief";
            /** @example <p>Hello</p> */
            value: string;
        };
        /** @description Written brief */
        WorkRequestTextBriefResponse: {
            /**
             * @description discriminator enum property added by openapi-typescript
             * @enum {string}
             */
            type: "text_brief";
            /** @example <p>Hello</p> */
            value: string;
        };
        WorkRequestUpdateRequest: {
            /** @description List of users assigned to the work request */
            assignees?: string[];
            /**
             * @description Update the priority of the work request
             * @example Low
             * @enum {string}
             */
            priority?: "Low" | "Medium" | "High";
            /**
             * @description Update the status of the work request
             * @example Submitted
             * @enum {string}
             */
            status?: "Accepted" | "Declined" | "Submitted" | "Completed";
        };
        WorkflowChannel: {
            /**
             * @description Name of the channel
             * @example exampleChannel
             */
            display_name: string;
            /**
             * @description Type of channel
             * @example twitter
             */
            type: string;
        };
        WorkflowCustomfield: {
            /**
             * @description List of the custom field choices
             * @default []
             * @example [
             *       "choice1",
             *       "choice2"
             *     ]
             */
            default_value: string[];
            /**
             * @description Name of the custom field
             * @example exampleCustomField
             */
            field: string;
            /**
             * @description Indicates whether a custom field is required
             * @example false
             */
            is_required: boolean;
            /**
             * @description Type of custom field
             * @example equal
             */
            type: string;
        };
        WorkflowLabel: {
            /**
             * @description List of label names
             * @example [
             *       "label1",
             *       "label2"
             *     ]
             */
            default_value: string[];
            /**
             * @description Indicates whether a label is required
             * @example true
             */
            is_required: boolean;
            /**
             * @description Indicates whether a label is required at task creation
             * @example false
             */
            is_required_at_task_creation: boolean;
            /**
             * @description Name of the Label Group
             * @example Content Pillar
             */
            label_type: string;
        };
        WorkflowListResponseItem: {
            /**
             * @description Description of the workflow
             * @example Nice description
             */
            description: string | null;
            /**
             * @description Unique identifier of the workflow
             * @example 6e7f910551b00a722e04188
             */
            id: string;
            /**
             * @description Indicates whether the workflow is flexible
             * @example true
             */
            is_flexible: boolean;
            /**
             * @description Name of the workflow
             * @example The awesome workflow
             */
            name: string;
        };
        WorkflowResponse: {
            /**
             * @description Description of the workflow
             * @example Nice description
             */
            description: string | null;
            /** @description List of fields */
            fields: components["schemas"]["ObjectFieldCreateResponse"][];
            /**
             * @description Unique identifier of the workflow
             * @example 6e7f910551b00a722e04188
             */
            id: string;
            /**
             * @description Indicates whether the workflow is active
             * @example true
             */
            is_active: boolean;
            /**
             * @description Indicates whether the asset approval of a workflow is enabled
             * @example true
             */
            is_asset_approval_enabled: boolean;
            /**
             * @description Indicates whether the workflow is flexible
             * @example true
             */
            is_flexible: boolean;
            /**
             * @description Indicates whether the resource management of a workflow is enabled
             * @example true
             */
            is_resource_management_enabled: boolean;
            /**
             * @description Indicates whether the smart duration of a workflow is enabled
             * @example true
             */
            is_smart_duration_enabled: boolean;
            /**
             * @description Name of the workflow
             * @example The awesome workflow
             */
            name: string;
            /** @description List of steps */
            steps: {
                /**
                 * @description Identifier of the step
                 * @example 6e7f910551b00a722e04181
                 */
                id: string;
                /**
                 * @description Name of the step
                 * @example Step 1
                 */
                name: string;
                /** @description List of sub-steps */
                sub_steps: {
                    /**
                     * @description List of actions of a sub-step
                     * @example [
                     *       {
                     *         "name": "Approve"
                     *       },
                     *       {
                     *         "name": "Find Content"
                     *       },
                     *       {
                     *         "name": "Write/Edit Content"
                     *       },
                     *       {
                     *         "name": "Publish"
                     *       },
                     *       {
                     *         "name": "Set Publish Destination"
                     *       },
                     *       {
                     *         "name": "Edit Fields"
                     *       },
                     *       {
                     *         "name": "Set Publish Date"
                     *       },
                     *       {
                     *         "name": "Write/Edit Share"
                     *       },
                     *       {
                     *         "name": "Select Image"
                     *       }
                     *     ]
                     */
                    actions: {
                        /**
                         * @description Name of the action
                         * @example Approve
                         */
                        name: string;
                    }[];
                    /** @description List of assignees of a sub-step */
                    assignees: {
                        /**
                         * @description Identifier of the sub-step assignee
                         * @example 6e7f910551b00a722e04182
                         */
                        id: string;
                        /** @description Meta links */
                        links: {
                            /**
                             * @description URL of the assignee
                             * @example https://api.cmp.optimizely.com/v3/users/6e7f910551b00a722e04182
                             */
                            self: string;
                        };
                        /**
                         * @description Name of the assignee
                         * @example Ned Stark
                         */
                        name: string;
                        /**
                         * @description Type of the assignee
                         * @example user
                         * @enum {string}
                         */
                        type: "user" | "team";
                    }[];
                    /**
                     * @description Identifier of the sub-step
                     * @example 6e7f910551b00a722e04182
                     */
                    id: string;
                    /**
                     * @description Name of the sub-step
                     * @example Step 1 Substep 1
                     */
                    name: string;
                    /**
                     * @description Type of the sub-step
                     * @example default
                     * @enum {string}
                     */
                    type: "default" | "external";
                }[];
            }[];
        };
        WorkflowStep: {
            /**
             * @description Description of the workflow step
             * @default null
             * @example This is a dummy example
             */
            description: string | null;
            /**
             * @description Duration of the workflow step
             * @default null
             * @example 1
             */
            duration: number | null;
            /**
             * @description Name of the workflow step
             * @example Step 1
             */
            label: string;
            /** @description List of the substeps */
            substeps: components["schemas"]["WorkflowSubStep"][];
        };
        WorkflowSubStep: {
            /**
             * @description List of actions
             * @example [
             *       "findContent",
             *       "writeEditSocialPost"
             *     ]
             */
            actions: string[];
            /**
             * @description Description of the substep
             * @default null
             * @example This is a substep description
             */
            description: string | null;
            /** @description External substep configuration */
            external_sub_step_config?: {
                /**
                 * @description Indicates whether the user interaction is allowed
                 * @default false
                 * @example false
                 */
                is_user_interaction_allowed: boolean;
            };
            /**
             * @description Name of the external system if the substep is external
             * @default null
             * @example null
             */
            external_system: string | null;
            /**
             * @description Indicates whether a substep is external
             * @default false
             * @example false
             */
            is_external: boolean | null;
            /**
             * @description Name of the substep
             * @default null
             * @example Substep 1
             */
            label: string | null;
        };
        /**
         * error_reason
         * @description An enumeration.
         * @enum {string}
         */
        error_reason: "REQUIRED_FIELD_ABSENT" | "LIST_EMPTY" | "INVALID_STATE" | "INVALID_SELF_REF" | "MIN_NOT_MET" | "MAX_NOT_MET" | "PATTERN_ERROR";
    };
    responses: {
        /** @description Client error */
        ClientError: {
            headers: {
                [name: string]: unknown;
            };
            content: {
                /**
                 * @example {
                 *       "message": "Unsupported arguments: a,b,c"
                 *     }
                 */
                "application/json": components["schemas"]["Error"];
            };
        };
        /** @description Permission error */
        Forbidden: {
            headers: {
                [name: string]: unknown;
            };
            content: {
                /**
                 * @example {
                 *       "message": "You do not have the permission to perform this operation"
                 *     }
                 */
                "application/json": components["schemas"]["Error"];
            };
        };
        /** @description Not found error */
        NotFound: {
            headers: {
                [name: string]: unknown;
            };
            content: {
                /**
                 * @example {
                 *       "message": "Resource not found"
                 *     }
                 */
                "application/json": components["schemas"]["Error"];
            };
        };
        /** @description Authorization error */
        Unauthorized: {
            headers: {
                [name: string]: unknown;
            };
            content: {
                /**
                 * @example {
                 *       "message": "Unauthorized"
                 *     }
                 */
                "application/json": components["schemas"]["Error"];
            };
        };
        /** @description Unprocessable entity error */
        UnprocessableEntity: {
            headers: {
                [name: string]: unknown;
            };
            content: {
                /**
                 * @example {
                 *       "message": "Resource cannot be processed"
                 *     }
                 */
                "application/json": components["schemas"]["Error"];
            };
        };
    };
    parameters: {
        /**
         * @description Starting index of results (zero indexed)
         * @example 5
         */
        offset: number;
        /**
         * @description Number of results to return per page
         * @example 15
         */
        page_size: number;
    };
    requestBodies: never;
    headers: never;
    pathItems: never;
};
export type $defs = Record<string, never>;
export interface operations {
    getArticle: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched article */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LibraryArticle"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    listAssetLineages: {
        parameters: {
            query?: {
                /** @example 1265ca474ba987cc925164aa7 */
                asset_id?: string;
                /** @example external */
                used_in?: "external";
                /**
                 * @description Date and time as the lower limit to filter asset lineages by `created_at`, in ISO 8601 UTC format
                 * @example 2018-11-30T13:32:44Z
                 */
                created_at__from?: string;
                /**
                 * @description Date and time as the upper limit to filter asset lineages by `created_at`, in ISO 8601 UTC format
                 * @example 2018-11-30T13:32:44Z
                 */
                created_at__to?: string;
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched asset lineage */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of asset lineage */
                        data: components["schemas"]["AssetLineageResponse"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/asset-lineages?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    getAssetUrl: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                asset_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Redirect to the discovered asset URL */
            302: {
                headers: {
                    /**
                     * @description URL of the discovered asset
                     * @example https://api.cmp.optimizely.com/v3/raw-files/5fe4925ef8a9378056bf4e37
                     */
                    Location?: string;
                    /**
                     * @description Location of the asset
                     * @example library
                     */
                    "X-Resource-Location"?: "library" | "task";
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    listAssets: {
        parameters: {
            query?: {
                /** @description Asset type to filter by. Example: `type=image&type=video` */
                type?: ("article" | "image" | "video" | "raw_file" | "structured_content")[];
                /** @description Label to filter by. **Must be passed as urlencoded string**. Labels that do not exist are ignored. If none of the provided labels exist, filtering is not applied. Example – `label=%7B%22group%22%3A%22ee63e3ee43925bb5cc8cd17b817d93ee%22%2C%22values%22%3A%5B%226706efc7828cd6aaedbc0434139cd3e1%22%2C%221f32651216cf2aefcaa08be1ea7dedf1%22%5D%7D` */
                label?: {
                    /**
                     * @description Identifier of the label group
                     * @example 2467e583a60e23fda2b89db81a453cd2
                     */
                    group: string;
                    /** @description List of identifiers of the label values */
                    values: string[];
                }[];
                /** @description List of fields to filter by. **Must be passed as base64 encoded string**. Fields that do not exist are ignored. If none of the provided fields exist, filtering is not applied. Example – `fields=Wwp7CiJpZCI6ICI2N2E4NDZhMWM3NzU1YTFwNThpNjh5MzVhIiwKInZhbHVlcyI6IFsiNjdhODQ2YTFjNzc1YWU1YWExYTE0YTA1Il0KfQpd=` */
                fields?: {
                    /**
                     * @description Identifier of the field
                     * @example 2467e583a60e23fda2b89db81a453cd2
                     */
                    id: string;
                    /** @description List of identifiers of the field values */
                    values: string[];
                }[];
                /**
                 * @description Date and time as the lower limit to filter assets by `created_at`, in ISO 8601 UTC format
                 * @example 2018-11-30T13:32:44Z
                 */
                created_at__from?: string;
                /**
                 * @description Date and time as the upper limit to filter assets by `created_at`, in ISO 8601 UTC format
                 * @example 2018-11-30T13:32:44Z
                 */
                created_at__to?: string;
                /**
                 * @description Date and time as the lower limit to filter assets by `modified_at`, in ISO 8601 UTC format
                 * @example 2018-11-30T13:32:44Z
                 */
                modified_at__from?: string;
                /**
                 * @description Date and time as the upper limit to filter assets by `modified_at`, in ISO 8601 UTC format
                 * @example 2018-11-30T13:32:44Z
                 */
                modified_at__to?: string;
                /**
                 * @description ID of the library folder to include assets from
                 * @example 6bb8db20a5b611ebae319b7c541b1a7f
                 */
                folder_id?: string;
                /**
                 * @description Indicates whether assets from subfolders need to be included
                 * @example false
                 */
                include_subfolder_assets?: boolean;
                /**
                 * @description Search assets by title or content description
                 * @example Cute Cat
                 */
                search_text?: string;
                /**
                 * @description ID of the campaign to include assets from
                 * @example 673218bee23887779e9e9a6f
                 */
                campaign_id?: string;
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched assets */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of assets */
                        data: components["schemas"]["AssetResponse"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/assets?offset=10&page_size=10 */
                            next?: string | null;
                        };
                        /** @description Total number of assets found based on the request */
                        total_count: number;
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    createAsset: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** @description Payload to create an asset */
        requestBody: {
            content: {
                "application/json": components["schemas"]["LibraryAssetCreateRequest"];
            };
        };
        responses: {
            /** @description Created asset */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AssetResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    listAssetFields: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                asset_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched fields */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of fields */
                        data: components["schemas"]["AssetFieldListResponseItem"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/assets/5f857f30e1c4a2038d6179e9/fields?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    updateAssetFields: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                asset_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to replace the fields */
        requestBody: {
            content: {
                "application/json": components["schemas"]["AssetFieldsUpdateRequest"];
            };
        };
        responses: {
            /** @description List of fields after replacement */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of fields */
                        data: components["schemas"]["AssetFieldListResponseItem"][];
                        /** @description Meta links */
                        links: {
                            /**
                             * @description URL to get the asset fields
                             * @example https://api.cmp.optimizely.com/v3/assets/5f857f30e1c4a2038d6179e9/fields
                             */
                            asset_fields: string;
                            /**
                             * @description URL of the asset
                             * @example https://api.cmp.optimizely.com/v3/asset-urls/5f857f30e1c4a2038d6179e9
                             */
                            asset_url: string;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    updateAssetField: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                asset_id: string;
                /** @example 4f857f30e1c4a2038d6179e7 */
                field_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the field value */
        requestBody: {
            content: {
                "application/json": components["schemas"]["AssetFieldUpdateRequest"];
            };
        };
        responses: {
            /** @description Updated the asset field */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AssetFieldUpdateResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    createAssetLineage: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                asset_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to add an external asset lineage for a library asset */
        requestBody: {
            content: {
                "application/json": components["schemas"]["AssetLineageCreateRequest"];
            };
        };
        responses: {
            /** @description Created external asset lineage */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AssetLineageResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    deleteAssetLineage: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                asset_id: string;
                /** @example ff3f8e024f0611ed993202420ac8001b */
                lineage_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Deleted the asset lineage */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listAssetPermissions: {
        parameters: {
            query?: {
                /** @example view */
                access?: "view" | "edit" | "comment" | "delete";
                /** @example view */
                max_access?: "view" | "edit" | "comment" | "delete";
                /** @example view */
                min_access?: "view" | "edit" | "comment" | "delete";
            };
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                asset_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of permissions */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AssetPermissionListResponseItem"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    addAssetPermissions: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                asset_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to add list of accessors to asset permissions */
        requestBody: {
            content: {
                "application/json": components["schemas"]["AssetPermissionBulkCreateRequest"];
            };
        };
        responses: {
            /** @description Permission access granted */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    removeAssetPermission: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                asset_id: string;
                /** @example 5d7f910551b00a722e0418830cee6631 */
                accessor_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Accessor's access removed */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    updateAssetPermission: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                asset_id: string;
                /** @example 5d7f910551b00a722e0418830cee6631 */
                accessor_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update permission of accessors */
        requestBody: {
            content: {
                "application/json": components["schemas"]["AssetPermissionUpdateRequest"];
            };
        };
        responses: {
            /** @description Permission access granted */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listAssetRenditions: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                asset_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched renditions */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of renditions */
                        data: components["schemas"]["BaseAssetRenditionResponse"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/assets/5d7f910551b00a722e0418830cee6631/renditions?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    createAssetVersion: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                asset_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to add a version to a library asset */
        requestBody: {
            content: {
                "application/json": components["schemas"]["AssetVersionCreateRequest"];
            };
        };
        responses: {
            /** @description Created version */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LibraryAssetVersionResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listBrandComplianceCategories: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched brand compliance categories */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of brand compliance categories */
                        data: components["schemas"]["BrandComplianceCategoriesResponse"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/brand-compliance/categories?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    listCampaigns: {
        parameters: {
            query?: {
                /** @example 127a32121da23c3a543 */
                owner?: string;
                /** @example 2022-08-24T00:00:00.000Z */
                start_date?: string;
                /** @example 2022-08-24T00:00:00.000Z */
                end_date?: string;
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched campaigns */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of campaigns */
                        data: components["schemas"]["CampaignListResponseItem"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/campaigns?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    createCampaign: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** @description Payload to create a campaign */
        requestBody: {
            content: {
                "application/json": components["schemas"]["CampaignCreateRequest"];
            };
        };
        responses: {
            /** @description Created campaign */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["CampaignResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    updateCampaignField: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                campaign_id: string;
                /** @example 4f857f30e1c4a2038d6179e7 */
                field_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the field value */
        requestBody: {
            content: {
                "application/json": components["schemas"]["CampaignFieldUpdateRequest"];
            };
        };
        responses: {
            /** @description Updated the campaign field */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["CampaignFieldUpdateResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getCampaign: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 8q7f910551b00a722e0418830cee6612 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched campaign */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["CampaignResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    updateCampaign: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 8q7f910551b00a722e0418830cee6612 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the campaign */
        requestBody: {
            content: {
                "application/json": components["schemas"]["CampaignUpdateRequest"];
            };
        };
        responses: {
            /** @description Campaign updated */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    addAttachmentToCampaign: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to add an attachment */
        requestBody: {
            content: {
                "application/json": components["schemas"]["AttachmentRequest"];
            };
        };
        responses: {
            /** @description Attachment added to the campaign */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AttachmentResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getCampaignBrief: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched brief of the campaign */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["CampaignBriefResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    addCommentToCampaign: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example da80cfd7bbf84959a8a981acbad996b3 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to add a comment */
        requestBody: {
            content: {
                "application/json": components["schemas"]["CommentCreateRequest"];
            };
        };
        responses: {
            /** @description Created a comment for the campaign */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["CampaignCommentResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    listCampaignFields: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched fields */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of fields */
                        data: components["schemas"]["FieldListResponseItem"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/campaigns/5f857f30e1c4a2038d6179e9/fields?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    addFieldToCampaign: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to add a field to a campaign */
        requestBody: {
            content: {
                "application/json": components["schemas"]["ObjectFieldCreateRequest"];
            };
        };
        responses: {
            /** @description Added field to a campaign. */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ObjectFieldCreateResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listEvents: {
        parameters: {
            query?: {
                /** @example 2024-11-24T00:00:00.000Z */
                start_date?: string;
                /** @example 2024-11-26T00:00:00.000Z */
                end_date?: string;
                /** @example 66d953340019d7b86833ac6d */
                campaign_id?: string;
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched events */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of events */
                        data: components["schemas"]["EventResponse"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/events?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    createEvent: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** @description Payload to create a new event */
        requestBody: {
            content: {
                "application/json": components["schemas"]["EventCreateRequest"];
            };
        };
        responses: {
            /** @description Created event details */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["EventResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    getEvent: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched event details */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["EventResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    updateEvent: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update an event */
        requestBody: {
            content: {
                "application/json": components["schemas"]["EventUpdateRequest"];
            };
        };
        responses: {
            /** @description Updated event details */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["EventResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listEventFields: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched fields */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of fields */
                        data: components["schemas"]["FieldListResponseItem"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/events/5f857f30e1c4a2038d6179e9/fields?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    updateEventFields: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to replace the fields */
        requestBody: {
            content: {
                "application/json": components["schemas"]["EventFieldsUpdateRequest"];
            };
        };
        responses: {
            /** @description Event fields updated */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listFields: {
        parameters: {
            query?: {
                /** @example 12p5ca474ba9i7cc92y164aa7 */
                ids?: string;
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched fields */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of fields */
                        data: (components["schemas"]["BaseSettingsFieldsResponse"] | components["schemas"]["LabelTypeSettingsFieldResponse"] | components["schemas"]["DropdownTypeSettingsFieldResponse"] | components["schemas"]["CheckboxAndRadioTypeSettingsFieldResponse"] | components["schemas"]["GenericNumberTypeSettingsFieldResponse"] | components["schemas"]["CurrencyNumberTypeSettingsFieldResponse"])[];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/fields?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    createField: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** @description Payload to create a new field in an Organization */
        requestBody: {
            content: {
                "application/json": components["schemas"]["SettingsFieldCreateRequest"];
            };
        };
        responses: {
            /** @description Created new field in an Organization */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /**
                         * @description Identifier of the field that has been created.
                         * @example 6ceee2f4fa3411ecb37802420ac8001b
                         */
                        id: string;
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    deleteFieldChoice: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 8a7f91p551b00i722e04y8830cae6612 */
                field_id: string;
                /** @example 8a7b910551b00a722e0418860cee6612 */
                choice_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Successful response */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    updateFieldChoice: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 8a7f91p551b00i722e04y8830cae6612 */
                field_id: string;
                /** @example 8a7b91p51b00i722e041y860aee6612 */
                choice_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update a choice in the field */
        requestBody: {
            content: {
                "application/json": components["schemas"]["SettingsFieldChoiceUpdateRequest"];
            };
        };
        responses: {
            /** @description Successful response */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    updateField: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 8q7f910551b00a722e0418830cee6612 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update in an Organization. */
        requestBody: {
            content: {
                "application/json": components["schemas"]["SettingsFieldUpdateRequest"];
            };
        };
        responses: {
            /** @description Updated field in an Organization */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    createFieldChoices: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 8a7f910551b00a722e0418830cee6612 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to create new choices in the field */
        requestBody: {
            content: {
                "application/json": components["schemas"]["SettingsFieldChoiceCreateRequest"];
            };
        };
        responses: {
            /** @description Field with all the choices including the newly created ones */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["SettingsFieldChoiceCreateResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    createFileUrls: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** @description Payload to generate download URLs for files. */
        requestBody: {
            content: {
                "application/json": components["schemas"]["FileUrlBulkCreateRequest"];
            };
        };
        responses: {
            /** @description Generated URLs */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["BatchFileUrlResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    listFolders: {
        parameters: {
            query?: {
                /**
                 * @description ID of the parent folder to filter by
                 * @example 1d9d8aeca10811ebbc640242ac12001b
                 */
                parent_folder_id?: string;
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched folders */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of folders */
                        data: components["schemas"]["FolderResponse"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/folders?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    createFolder: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** @description Payload to create a folder */
        requestBody: {
            content: {
                "application/json": components["schemas"]["FolderCreateRequest"];
            };
        };
        responses: {
            /** @description Created folder */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["FolderResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getFolder: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6632 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched folder */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["FolderResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    deleteFolder: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6632 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Folder deleted successfully */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    updateFolder: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6632 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update a folder */
        requestBody: {
            content: {
                "application/json": components["schemas"]["FolderUpdateRequest"];
            };
        };
        responses: {
            /** @description Updated Folder */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["FolderResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    listFolderPermissions: {
        parameters: {
            query?: {
                /** @example view */
                access?: "view" | "edit" | "comment" | "delete";
                /** @example view */
                max_access?: "view" | "edit" | "comment" | "delete";
                /** @example view */
                min_access?: "view" | "edit" | "comment" | "delete";
            };
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of permissions */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["FolderPermissionListResponseItem"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    addFolderPermissions: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to add list of accessors to folder permissions */
        requestBody: {
            content: {
                "application/json": components["schemas"]["FolderPermissionBulkCreateRequest"];
            };
        };
        responses: {
            /** @description Permission access granted */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    removeFolderPermission: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                id: string;
                /** @example 5d7f910551b00a722e0418830cee6631 */
                accessor_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Accessor's access removed */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    updateFolderPermission: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                id: string;
                /** @example 5d7f910551b00a722e0418830cee6631 */
                accessor_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update permission of accessors */
        requestBody: {
            content: {
                "application/json": components["schemas"]["FolderPermissionUpdateRequest"];
            };
        };
        responses: {
            /** @description Permission access granted */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    getImage: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6632 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched image */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LibraryImage"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    deleteImage: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6632 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Deleted the image */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    updateImage: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the image */
        requestBody: {
            content: {
                "application/json": components["schemas"]["LibraryImageUpdateRequest"];
            };
        };
        responses: {
            /** @description Updated Image */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LibraryImage"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listLabelGroups: {
        parameters: {
            query?: {
                /**
                 * @description Source organization type to filter by
                 * @example current
                 */
                source_org_type?: "current" | "related";
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched label groups */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of label groups */
                        data: components["schemas"]["LabelGroup"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/label-groups?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    listMilestones: {
        parameters: {
            query?: {
                /** @example 63f1c2b675be4132854d2742 */
                campaign_id?: string | null;
                /** @example 2023-10-01T00:00:00Z */
                due_date__from?: string;
                /** @example 2023-10-31T23:59:59Z */
                due_date__to?: string;
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched milestones */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of milestones */
                        data: components["schemas"]["MilestoneResponse"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/milestones?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    createMilestone: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** @description Payload to create a milestone */
        requestBody: {
            content: {
                "application/json": components["schemas"]["MilestoneCreateRequest"];
            };
        };
        responses: {
            /** @description Created milestone */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["MilestoneResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getMilestone: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 63f1c2b675be4132854d2741 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Details of the milestone */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["MilestoneResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    updateMilestone: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 63f1c2b675be4132854d2741 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update a milestone */
        requestBody: {
            content: {
                "application/json": components["schemas"]["MilestoneUpdateRequest"];
            };
        };
        responses: {
            /** @description Updated milestone */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["MilestoneResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    createMultipartUpload: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": {
                    /**
                     * @description Size of the file in bytes
                     * @example 104857600
                     */
                    file_size: number;
                    /**
                     * @description The size of each part of the file to be uploaded. The size of the last part can be smaller and hence can be ignored. `upload_part_count` is calculated based on `file_size` and `part_size`. **The maximum allowed `upload_part_count` is 10000.** So use larger `part_size` for larger files to keep `upload_part_count` under 10000.
                     * @default 5242880
                     * @example 6000000
                     */
                    part_size?: number;
                };
            };
        };
        responses: {
            /** @description Successful operation */
            201: {
                headers: {
                    /** @example gzip */
                    "Content-Encoding"?: "gzip";
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /**
                         * Format: date-time
                         * @description Expiration time of the upload URLs
                         * @example 2025-03-03T12:00:00Z
                         */
                        expires_at: string;
                        /**
                         * @description ID of the multipart upload
                         * @example abcdef1234567890
                         */
                        id: string;
                        links: {
                            /**
                             * Format: uri
                             * @description URL to complete the upload
                             * @example https://api.cmp.optimizely.com/v3/multipart-uploads/abcdef1234567890/complete
                             */
                            complete: string;
                            /**
                             * Format: uri
                             * @description URL to check the upload status
                             * @example https://api.cmp.optimizely.com/v3/multipart-uploads/abcdef1234567890/status
                             */
                            status: string;
                        };
                        /**
                         * @description Number of parts for the upload
                         * @example 10
                         */
                        upload_part_count: number;
                        /**
                         * @description Array of pre-signed URLs for each part. **Each file part must be uploaded to the URL using PUT requests following the order of the URLs**
                         * @example [
                         *       "https://s3.amazonaws.com/xyz/abcdef1234567890/0",
                         *       "https://s3.amazonaws.com/xyz/abcdef1234567890/1"
                         *     ]
                         */
                        upload_part_urls: string[];
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    completeMultipartUpload: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /**
                 * @description ID of the multipart upload to complete
                 * @example abcdef1234567890
                 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Upload completion initiated successfully */
            202: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /**
                         * @description Unique identifier of the file combining all parts. **This key will be needed for various file / attachment addition to different CMP resources using endpoints e.g. `POST /v3/assets`, `POST /v3/campaigns/{id}/attachments` etc.**
                         * @example ce8995aea58b11ea8cd90242ac120005/abcdef1234567890
                         */
                        key: string;
                        links: {
                            /**
                             * Format: uri
                             * @description URL to check upload status
                             * @example https://api.cmp.optimizely.com/v3/multipart-uploads/abcdef1234567890/status
                             */
                            status: string;
                        };
                    };
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    getMultipartUploadStatus: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /**
                 * @description ID of the multipart upload to check status for
                 * @example abcdef1234567890
                 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Successfully retrieved upload status */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /**
                         * Format: date-time
                         * @description Expiration time of the upload
                         * @example 2025-03-03T12:00:00Z
                         */
                        expires_at: string;
                        /**
                         * @description ID of the multipart upload
                         * @example abcdef1234567890
                         */
                        id: string;
                        /**
                         * @description Unique identifier of the file combining all parts. **This key will be needed for various file / attachment addition to different CMP resources using endpoints e.g. `POST /v3/assets`, `POST /v3/campaigns/{id}/attachments` etc.**
                         * @example ce8995aea58b11ea8cd90242ac120005/abcdef1234567890
                         */
                        key: string;
                        links: {
                            /**
                             * Format: uri
                             * @description URL to this status endpoint
                             * @example https://api.cmp.optimizely.com/v3/multipart-uploads/abcdef1234567890/status
                             */
                            self: string;
                        };
                        /**
                         * @description Current status of the upload
                         * @example UPLOAD_COMPLETION_IN_PROGRESS
                         * @enum {string}
                         */
                        status: "UPLOAD_COMPLETION_NOT_STARTED" | "UPLOAD_COMPLETION_IN_PROGRESS" | "UPLOAD_COMPLETION_SUCCEEDED" | "UPLOAD_COMPLETION_FAILED";
                        /**
                         * @description Detailed status message
                         * @example Some error message related to the status
                         */
                        status_message: string | null;
                    };
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getPublishingEvent: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /**
                 * @description Unique identifier of the publishing event
                 * @example 1d9d8aeca10811ebbc640242ac12001b
                 */
                publishing_event_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched publishing event */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["PublishingEventResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getPublishingEventAssetMetadata: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /**
                 * @description Unique identifier of the publishing event
                 * @example 1d9d8aeca10811ebbc640242ac12001b
                 */
                publishing_event_id: string;
                /**
                 * @description Unique identifier of the asset
                 * @example 1d9d8aeca10811ebbc640242ac12003c
                 */
                asset_id: string;
                /**
                 * @description Unique identifier of the publishing metadata
                 * @example 5ebcd5644967474414564
                 */
                publishing_metadata_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched publishing metadata */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["PublishingEventMetadataResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    listPublishingEventMetadata: {
        parameters: {
            query?: {
                /**
                 * @description Publishing status of the asset
                 * @example published
                 */
                status?: "published" | "unpublished" | "synced" | "failed";
                /**
                 * @description Type of asset
                 * @example article
                 */
                asset_type?: "article" | "image" | "video" | "raw_file" | "structured_content";
                /**
                 * @description Unique identifier of the asset.
                 * @example 1d9d8aeca10811ebbc640242ac12001b
                 */
                asset_id?: string;
                /**
                 * @description The locale to which the asset is being published to.
                 * @example en
                 */
                locale?: string;
            };
            header?: never;
            path: {
                /**
                 * @description Unique identifier of the publishing event
                 * @example 1d9d8aeca10811ebbc640242ac12001b
                 */
                publishing_event_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of Fetched publishing metadata */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["PublishingEventMetadataListResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    bulkCreatePublishingEventMetadata: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /**
                 * @description Unique identifier of the publishing event
                 * @example 1d9d8aeca10811ebbc640242ac12001b
                 */
                publishing_event_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to create publishing metadata for all assets in a publishing event */
        requestBody: {
            content: {
                "application/json": components["schemas"]["PublishingEventMetadataBulkCreateRequest"];
            };
        };
        responses: {
            /** @description Fetched all publishing metadata associated with the publishing event */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["PublishingEventMetadataBulkCreateResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getRawFile: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6634 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched raw file */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LibraryRawFile"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    deleteRawFile: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6634 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Deleted the raw file */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    updateRawFile: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the raw file */
        requestBody: {
            content: {
                "application/json": components["schemas"]["LibraryRawFileUpdateRequest"];
            };
        };
        responses: {
            /** @description Raw file updated */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LibraryRawFile"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    getRenditionConfig: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Data of a rendition configuration */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["RenditionConfigResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getRendition: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example c2603a6e0b4811eebea95edf08e23f48 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Details of the rendition */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["DetailedAssetRenditionResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getSettings: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched settings */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["Settings"];
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    updateSettings: {
        parameters: {
            query?: {
                /**
                 * @description If `execute=true` the settings are created or updated. Otherwise, the endpoint returns only the changeset.
                 * @example true
                 */
                execute?: boolean;
                /**
                 * @description If `overwrite_workflows=true` the existing workflows are overwritten. Otherwise, a new workflow is created where a prefix `Copy of` is added to the workflow's name.
                 * @example true
                 */
                overwrite_workflows?: boolean;
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** @description Payload to create or update settings */
        requestBody: {
            content: {
                "application/json": components["schemas"]["SettingsResources"];
            };
        };
        responses: {
            /** @description Created or updated settings and changeset */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["SettingsUpdateResponse"];
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listSCContentTypes: {
        parameters: {
            query?: {
                source?: string;
                disabled?: boolean;
                list?: components["schemas"]["ContentTypeListingOption"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Successful response */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["BaseContentTypeModel"][];
                };
            };
            /** @description Validation error */
            422: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["HTTPValidationError"];
                };
            };
        };
    };
    createSCContentType: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["SCContentTypeCreateRequest"];
            };
        };
        responses: {
            /** @description Successful response */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["SCContentTypeCreateResponse"];
                };
            };
            /** @description Created */
            201: {
                headers: {
                    /** @description URL to get the created resource */
                    Location?: string;
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["SCContentTypeCreateResponse"];
                };
            };
            /** @description Validation error */
            422: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["HTTPValidationError"];
                };
            };
        };
    };
    getSCContentType: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                content_type_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Successful response */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["SCContentType"];
                };
            };
            /** @description Validation error */
            422: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["HTTPValidationError"];
                };
            };
        };
    };
    updateSCContentType: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                content_type_id: string;
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["SCContentTypeUpdateRequest"];
            };
        };
        responses: {
            /** @description Successful response */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["SCContentTypeUpdateResponse"];
                };
            };
            /** @description Validation error */
            422: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["HTTPValidationError"];
                };
            };
        };
    };
    listSCContentTypeManagedMigrations: {
        parameters: {
            query?: {
                /**
                 * @description Whether include a summary of content migration status (total, not started, succeeded, errored).
                 * @example true
                 */
                content_migration_summary?: boolean;
                /**
                 * @description Pagination offset (number of jobs to skip).
                 * @example 10
                 */
                offset?: number;
                /**
                 * @description Pagination limit (number of jobs to return).
                 * @example 25
                 */
                limit?: number;
            };
            header?: never;
            path: {
                /** @example 645cb61c966d0c591320f636 */
                content_type_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of managed migration jobs for the content type. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["SCContentTypeManagedMigrationResponse"][];
                };
            };
        };
    };
    createSCContentTypeManagedMigration: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 645cb61c966d0c591320f636 */
                content_type_id: string;
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["SCContentTypeManagedMigrationCreateRequest"];
            };
        };
        responses: {
            /** @description Successful response. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description False if the managed migration job is not created. */
                        created: boolean;
                    };
                };
            };
            /** @description Created. */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description True if the managed migration job is created. */
                        created: boolean;
                        /** @description The ID of the created managed migration job. */
                        job_id?: string;
                    };
                };
            };
            400: components["responses"]["ClientError"];
            404: components["responses"]["NotFound"];
        };
    };
    validateSCContentTypeManagedMigration: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 645cb61c966d0c591320f636 */
                content_type_id: string;
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["SCContentTypeManagedMigrationValidateRequest"];
            };
        };
        responses: {
            /** @description Successful response. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description False if the managed migration not possible */
                        is_managed_migration_possible: boolean;
                    };
                };
            };
            /** @description Error response when no content type found by content_type_id and instance_id */
            404: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    /**
                     * @example {
                     *       "detail": "Could not find the content type"
                     *     }
                     */
                    "application/json": {
                        /** @description error message details */
                        detail: string;
                    };
                };
            };
        };
    };
    getSCContentTypeManagedMigration: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 645cb61c966d0c591320f636 */
                content_type_id: string;
                job_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Successful response. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["SCContentTypeManagedMigrationResponse"];
                };
            };
            /** @description Not Found. */
            404: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description The error message. */
                        detail?: string;
                    };
                };
            };
        };
    };
    deleteSCContentTypeManagedMigration: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 645cb61c966d0c591320f636 */
                content_type_id: string;
                job_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Successful response */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            404: components["responses"]["NotFound"];
        };
    };
    updateSCContentTypeManagedMigration: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 645cb61c966d0c591320f636 */
                content_type_id: string;
                job_id: string;
            };
            cookie?: never;
        };
        /** @description Payload for default field values. */
        requestBody: {
            content: {
                "application/json": {
                    default_values: components["schemas"]["LocalizedFieldValues"];
                };
            };
        };
        responses: {
            /** @description Successful response */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description True if the managed migration job is updated. */
                        updated?: boolean;
                    };
                };
            };
            /** @description Validation error */
            422: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["HTTPValidationError"];
                };
            };
        };
    };
    startSCContentTypeManagedMigration: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                content_type_id: string;
                job_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Successful response with started=true */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    /**
                     * @example {
                     *       "job_id": "sample_job_id",
                     *       "started": true
                     *     }
                     */
                    "application/json": components["schemas"]["SCContentTypeManagedMigrationStartResponse"];
                };
            };
            /** @description Error response when job found by job_id but job status is invalid (status not in [error or not_started]) */
            400: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    /**
                     * @example {
                     *       "detail": "No Job found for job_id: {job_id} and instance_id: {instance_id} with status not_started or error"
                     *     }
                     */
                    "application/json": components["schemas"]["HTTPException"];
                };
            };
            /** @description Error response when no job found by job_id with status not_started, error */
            404: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    /**
                     * @example {
                     *       "message": "No Job found for job_id: {job_id} and instance_id: {instance_id}."
                     *     }
                     */
                    "application/json": components["schemas"]["HTTPException"];
                };
            };
        };
    };
    listSCContentTypeVersions: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                content_type_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Successful response */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["BaseContentTypeVersionModel"][];
                };
            };
            /** @description Validation error */
            422: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["HTTPValidationError"];
                };
            };
        };
    };
    createSCContentTypeVersion: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                content_type_id: string;
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["SCContentTypeVersionCreateRequest"];
            };
        };
        responses: {
            /** @description Successful response */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["SCContentTypeCreateResponse"];
                };
            };
            /** @description Created */
            201: {
                headers: {
                    /** @description URL to get the created resource */
                    Location?: string;
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["SCContentTypeCreateResponse"];
                };
            };
            /** @description Validation error */
            422: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["HTTPValidationError"];
                };
            };
        };
    };
    getSCContentTypeVersion: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                content_type_id: string;
                version_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Successful response */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["SCContentTypeVersion"];
                };
            };
            /** @description Validation error */
            422: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["HTTPValidationError"];
                };
            };
        };
    };
    migrateSCContent: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                content_id: string;
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["SCContentMigrationCreateRequest"];
            };
        };
        responses: {
            /** @description Successful response */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description The identifier of the content */
                        content_guid?: string;
                        /** @description The hashed fingerprint of the content version */
                        content_hash?: string;
                        /** @description Whether the version was created or not */
                        created?: boolean;
                        /** @description The identifier of the content version */
                        version_guid?: string;
                    };
                };
            };
            400: components["responses"]["ClientError"];
            404: components["responses"]["NotFound"];
            409: components["responses"]["ClientError"];
            /** @description Validation error */
            422: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["HTTPValidationError"];
                };
            };
        };
    };
    acknowledgeSCContentPreview: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                content_id: string;
                version_id: string;
                preview_id: string;
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["SCContentPreviewAcknowledgeRequest"];
            };
        };
        responses: {
            /** @description Successful response */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": unknown;
                };
            };
            /** @description Validation error */
            422: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["HTTPValidationError"];
                };
            };
        };
    };
    completeSCContentPreview: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                content_id: string;
                version_id: string;
                preview_id: string;
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["SCContentPreviewCompleteRequest"];
            };
        };
        responses: {
            /** @description Successful response */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": unknown;
                };
            };
            /** @description Validation error */
            422: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["HTTPValidationError"];
                };
            };
        };
    };
    createStructuredContent: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** @description Payload to create the structured content */
        requestBody: {
            content: {
                "application/json": components["schemas"]["LibraryStructuredContentCreateRequest"];
            };
        };
        responses: {
            /** @description Structured content created */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LibraryStructuredContent"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    getStructuredContent: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example a24b834427a043ab8caed1ded4606f7d */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched structured content */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LibraryStructuredContent"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    updateStructuredContent: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the structured content */
        requestBody: {
            content: {
                "application/json": components["schemas"]["LibraryStructuredContentUpdateRequest"];
            };
        };
        responses: {
            /** @description Structured content updated */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LibraryStructuredContent"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listTasks: {
        parameters: {
            query?: {
                /** @example task 1 */
                search_key?: string;
                /** @example 8n7f910o51b00b722o0418n30cie8211 */
                campaign?: string;
                /** @example 8n7f910o51b00b722o0418n30cie8211 */
                workflow?: string;
                /** @example 8n7f910o51b00b722o0418n30cie8211 */
                milestone?: string;
                /** @example 2022-08-24T00:00:00.000Z */
                start_date?: string;
                /** @example 2022-08-24T00:00:00.000Z */
                due_date?: string;
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched tasks */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of tasks */
                        data: components["schemas"]["TaskListResponseItem"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/tasks?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    createTask: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** @description Payload to create a task */
        requestBody: {
            content: {
                "application/json": components["schemas"]["TaskCreateRequest"];
            };
        };
        responses: {
            /** @description Created task */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getTask: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched task */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    updateTask: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the task */
        requestBody: {
            content: {
                "application/json": components["schemas"]["TaskUpdateRequest"];
            };
        };
        responses: {
            /** @description Updated task */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    listTaskAssets: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched assets */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of assets */
                        data: components["schemas"]["TaskAssetResponse"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/assets?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    addAssetToTask: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                id: string;
            };
            cookie?: never;
        };
        /**
         * @description Payload to add an asset. There are two mechanisms for adding assets to a task:
         *     - `direct_upload` - An asset that has been uploaded using the [Upload assets](https://docs.developers.optimizely.com/content-marketing-platform/docs/upload-assets) endpoint.
         *     - `fork_from_library` - This will fork an existing asset from the library to the task.
         *     The `type` field is optional, and will default to `direct_upload` if not provided.
         *     Note: Renditions is currently not supported for library assets.
         */
        requestBody: {
            content: {
                "application/json": components["schemas"]["TaskAssetRequest"];
            };
        };
        responses: {
            /** @description Created asset */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskAssetResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listTaskAttachments: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched attachments */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of attachments */
                        data: components["schemas"]["AttachmentResponse"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/tasks/attachments?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getTaskBrief: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched brief of the task */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskBriefResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    listTaskCustomFields: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched custom fields */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of custom fields */
                        data: components["schemas"]["TaskCustomField"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/custom-fields?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getTaskArticle: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
                /** @example 5d7f910551b00a722e0418830cee6631 */
                article_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched article */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskArticle"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    listTaskAssetComments: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                task_id: string;
                /** @example 897f910551b00a722e0418830cee6612 */
                asset_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of comments for the asset */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of comments */
                        data?: components["schemas"]["TaskAssetCommentResponse"][];
                        pagination?: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/tasks/7d7f910551b00a722e0418830cee6612/assets/897f910551b00a722e0418830cee6612/comments?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    addCommentToTaskAsset: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                task_id: string;
                /** @example 897f910551b00a722e0418830cee6612 */
                asset_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to add comment */
        requestBody: {
            content: {
                "application/json": components["schemas"]["CommentCreateRequest"];
            };
        };
        responses: {
            /** @description Added comment */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskAssetCommentResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listTaskAssetDrafts: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
                /** @example 5d7f910551b00a722e0418830cee6631 */
                asset_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched drafts */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of drafts */
                        data: components["schemas"]["TaskAssetDraftListResponseItem"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/assets/5d7f910551b00a722e0418830cee6631/drafts?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    addDraftToTaskAsset: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
                /** @example 5d7f910551b00a722e0418830cee6631 */
                asset_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to add a draft to an asset */
        requestBody: {
            content: {
                "application/json": components["schemas"]["TaskAssetRequest"];
            };
        };
        responses: {
            /** @description Created draft */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskAssetDraftResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    getTaskAssetDraftBrandCompliance: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
                /** @example 5d7f910551b00a722e0418830cee6631 */
                asset_id: string;
                draft_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Brand compliance details of the asset draft */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskAssetDraftBrandComplianceResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            404: components["responses"]["NotFound"];
        };
    };
    updateTaskAssetDraftBrandCompliance: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
                /** @example 5d7f910551b00a722e0418830cee6631 */
                asset_id: string;
                draft_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the brand compliance */
        requestBody: {
            content: {
                "application/json": components["schemas"]["TaskAssetDraftBrandComplianceRequest"];
            };
        };
        responses: {
            /** @description Brand compliance details of the asset draft */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskAssetDraftBrandComplianceResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            404: components["responses"]["NotFound"];
        };
    };
    listTaskAssetFields: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                task_id: string;
                /** @example 897f910551b00a722e0418830cee6612 */
                asset_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fields for the asset */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of fields */
                        data?: components["schemas"]["AssetFieldListResponseItem"][];
                        pagination?: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/tasks/7d7f910551b00a722e0418830cee6612/assets/897f910551b00a722e0418830cee6612/fields?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    updateTaskAssetFields: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
                /** @example 5f857f30e1c4a2038d6179e9 */
                asset_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to replace the fields */
        requestBody: {
            content: {
                "application/json": components["schemas"]["TaskAssetFieldsUpdateRequest"];
            };
        };
        responses: {
            /** @description Fields successfulled updated. */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    addCommentToTask: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example da80cfd7bbf84959a8a981acbad996b3 */
                task_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to add a comment */
        requestBody: {
            content: {
                "application/json": components["schemas"]["CommentCreateRequest"];
            };
        };
        responses: {
            /** @description Created a comment for the task */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskCommentResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getTaskCustomField: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
                /** @example 9nu8ue9wf8u9nusd9q */
                custom_field_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched custom field */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskCustomField"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    updateTaskCustomField: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
                /** @example 9nu8ue9wf8u9nusd9q */
                custom_field_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the custom field. */
        requestBody: {
            content: {
                "application/json": components["schemas"]["TaskCustomFieldUpdateRequest"];
            };
        };
        responses: {
            /** @description Updated custom field */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskCustomField"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listTaskCustomFieldChoices: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
                /** @example 9nu8ue9wf8u9nusd9q */
                custom_field_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched custom field choices */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of custom field choices */
                        data: components["schemas"]["TaskCustomFieldChoiceListResponseItem"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/custom-fields/j90uv0sd9i0si09sdip/choices?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    listTaskFields: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched fields */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of fields */
                        data: components["schemas"]["FieldListResponseItem"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/tasks/5f857f30e1c4a2038d6179e9/fields?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    addFieldToTask: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to add a field to a task */
        requestBody: {
            content: {
                "application/json": components["schemas"]["ObjectFieldCreateRequest"];
            };
        };
        responses: {
            /** @description Added field to a task. */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ObjectFieldCreateResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    updateTaskField: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 124p857f3i220y38ad6179e9 */
                task_id: string;
                /** @example 124p857f3i220y38ad6179e9 */
                field_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update a field of a task */
        requestBody: {
            content: {
                "application/json": components["schemas"]["TaskFieldUpdateRequest"];
            };
        };
        responses: {
            /** @description The field of the task was updated. */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    removeTaskField: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
                /** @example fa63cece6aa811efb70202420ac80016 */
                field_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description The field is removed from the task */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getTaskImage: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
                /** @example 5d7f910551b00a722e0418830cee6632 */
                image_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched image */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskImage"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getTaskRawFile: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
                /** @example 5d7f910551b00a722e0418830cee6634 */
                raw_file_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched raw file */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskRawFile"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    updateTaskStep: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                task_id: string;
                /** @example 897f910551b00a722e0418830cee6612 */
                step_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the step */
        requestBody: {
            content: {
                "application/json": components["schemas"]["TaskStepRequest"];
            };
        };
        responses: {
            /** @description The updated step object */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskStep"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    getTaskSubStep: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                task_id: string;
                /** @example 897f910551b00a722e0418830cee6612 */
                step_id: string;
                /** @example 700f910551b00a722e0418830cee6612 */
                sub_step_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched task substep information */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskSubStep"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    updateTaskSubStep: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                task_id: string;
                /** @example 897f910551b00a722e0418830cee6612 */
                step_id: string;
                /** @example 700f910551b00a722e0418830cee6612 */
                sub_step_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the substep */
        requestBody: {
            content: {
                "application/json": components["schemas"]["TaskSubStepRequest"];
            };
        };
        responses: {
            /** @description Fetched task substep information */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskSubStep"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listTaskSubStepComments: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                task_id: string;
                /** @example 897f910551b00a722e0418830cee6612 */
                step_id: string;
                /** @example 700f910551b00a722e0418830cee6612 */
                sub_step_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched task substep comments */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of substep comments */
                        data: components["schemas"]["TaskSubStepCommentResponse"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/tasks/7d7f910551b00a722e0418830cee6612/steps/897f910551b00a722e0418830cee6612/sub-steps/700f910551b00a722e0418830cee6612/comments?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    addCommentToTaskSubStep: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                task_id: string;
                /** @example 897f910551b00a722e0418830cee6612 */
                step_id: string;
                /** @example 700f910551b00a722e0418830cee6612 */
                sub_step_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to create substep comment */
        requestBody: {
            content: {
                "application/json": components["schemas"]["CommentWithReplyCreateRequest"];
            };
        };
        responses: {
            /** @description Created task substep comment */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskSubStepCommentResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    getTaskSubStepComment: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                task_id: string;
                /** @example 897f910551b00a722e0418830cee6612 */
                step_id: string;
                /** @example 700f910551b00a722e0418830cee6612 */
                sub_step_id: string;
                /** @example 5fe3886c574b52a62a089237 */
                comment_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched task substep comment */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskSubStepCommentResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    deleteTaskSubStepComment: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                task_id: string;
                /** @example 897f910551b00a722e0418830cee6612 */
                step_id: string;
                /** @example 700f910551b00a722e0418830cee6612 */
                sub_step_id: string;
                /** @example 5fe3886c574b52a62a089237 */
                comment_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Deleted task substep comment */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    updateTaskSubStepComment: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                task_id: string;
                /** @example 897f910551b00a722e0418830cee6612 */
                step_id: string;
                /** @example 700f910551b00a722e0418830cee6612 */
                sub_step_id: string;
                /** @example 5fe3886c574b52a62a089237 */
                comment_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the substep comment */
        requestBody: {
            content: {
                "application/json": components["schemas"]["TaskSubStepCommentUpdateRequest"];
            };
        };
        responses: {
            /** @description Updated task substep comment */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskSubStepCommentResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    getTaskSubStepExternalWork: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                task_id: string;
                /** @example 897f910551b00a722e0418830cee6612 */
                step_id: string;
                /** @example 700f910551b00a722e0418830cee6612 */
                sub_step_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched external work information */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskExternalWorkResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    updateTaskSubStepExternalWork: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                task_id: string;
                /** @example 897f910551b00a722e0418830cee6612 */
                step_id: string;
                /** @example 700f910551b00a722e0418830cee6612 */
                sub_step_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the external work information */
        requestBody: {
            content: {
                "application/json": components["schemas"]["TaskExternalWorkRequest"];
            };
        };
        responses: {
            /** @description Updated external work information */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskExternalWorkResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listTaskSubStepFields: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path: {
                /** @example 7d7f910551b00a722e0418830cee6612 */
                task_id: string;
                /** @example 897f910551b00a722e0418830cee6612 */
                step_id: string;
                /** @example 700f910551b00a722e0418830cee6612 */
                sub_step_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched fields */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of fields */
                        data: (components["schemas"]["FieldTypeLabel"] | components["schemas"]["FieldTypeDropdown"] | components["schemas"]["FieldTypeRadio"] | components["schemas"]["FieldTypeCheckbox"])[];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/tasks/7d7f910551b00a722e0418830cee6612/steps/897f910551b00a722e0418830cee6612/sub-steps/700f910551b00a722e0418830cee6612/fields?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    addStructuredContentToTask: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to add a structured content */
        requestBody: {
            content: {
                "application/json": components["schemas"]["TaskStructuredContentCreateRequest"];
            };
        };
        responses: {
            /** @description Created content */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ContentDetailsModel"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    getTaskStructuredContent: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 645cb61c966d0c591320f636 */
                task_id: string;
                /** @example 5fe38c39574b52a62a089239 */
                content_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Structured content by Id for the task */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ContentDetailsModel"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    deleteTaskStructuredContent: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                task_id: string;
                /** @example ff3f8e024f0611ed993202420ac8001b */
                content_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Deleted the content */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    updateTaskStructuredContent: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example da80cfd7bbf84959a8a981acbad996b3 */
                task_id: string;
                /** @example da80cfd7bbf84959a8a981acbad996b3 */
                content_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update structured content */
        requestBody: {
            content: {
                "application/json": components["schemas"]["TaskStructuredContentUpdateRequest"];
            };
        };
        responses: {
            /** @description Updated structured content as response */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ContentDetailsModel"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    createTaskStructuredContentDraft: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example da80cfd7bbf84959a8a981acbad996b3 */
                task_id: string;
                /** @example da80cfd7bbf84959a8a981acbad996b3 */
                content_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to create structured content drafts */
        requestBody: {
            content: {
                "application/json": components["schemas"]["TaskStructuredContentDraftRequest"];
            };
        };
        responses: {
            /** @description Created structured content draft as response */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    /** @example null */
                    "application/json": Record<string, never>;
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    addUrlToTask: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 68fdf93d09caf1463f145dc6 */
                task_id: string;
            };
            cookie?: never;
        };
        /** @description Payload to add a URL to a task */
        requestBody: {
            content: {
                "application/json": components["schemas"]["AddUrlToTaskRequest"];
            };
        };
        responses: {
            /** @description Created task URL */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskUrlResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    getTaskVideo: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5f857f30e1c4a2038d6179e9 */
                task_id: string;
                /** @example 5d7f910551b00a722e0418830cee6633 */
                video_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched video */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TaskVideo"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    listTeams: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched teams */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of teams */
                        data: components["schemas"]["Team"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/teams?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    getTeam: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6633 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched the team */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TeamWithUsers"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    listTemplates: {
        parameters: {
            query?: {
                /** @example Kangaroo */
                search?: string;
                /** @example work_request */
                applicable_to?: "work_request" | "task_brief" | "campaign_brief";
                /** @example true */
                include_inactive?: boolean;
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched templates */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["TemplateListResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    getTemplate: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 9119a313057e401189407116fcd3aa24 */
                template_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched template */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    /**
                     * @example {
                     *       "applicable_to": [
                     *         "campaign_brief",
                     *         "task_brief",
                     *         "work_request"
                     *       ],
                     *       "description": "Sample Desciption",
                     *       "form_fields": [
                     *         {
                     *           "helper_text": "Sample helper text",
                     *           "identifier": "attachment",
                     *           "is_readonly": false,
                     *           "is_required": false,
                     *           "label": "Attach a file...",
                     *           "logic_rules": [
                     *             {
                     *               "action": {
                     *                 "target_field": {
                     *                   "identifier": "70ca0e8a954911ecb4be0242ac160014"
                     *                 },
                     *                 "type": "show_values",
                     *                 "values": [
                     *                   "73e05aa2b5d311eca17f0242ac160019"
                     *                 ]
                     *               },
                     *               "condition": {
                     *                 "operator": "any_of",
                     *                 "values": [
                     *                   "4a8b901e9fd911ecbefa0242ac160019"
                     *                 ]
                     *               }
                     *             }
                     *           ],
                     *           "sort_order": 0,
                     *           "type": "file",
                     *           "type_specific_meta": null
                     *         },
                     *         {
                     *           "helper_text": "Sample helper text",
                     *           "identifier": "73e05aa2b54nda3c320242ac1d5291",
                     *           "is_readonly": false,
                     *           "is_required": true,
                     *           "label": "Select publish destination",
                     *           "logic_rules": [],
                     *           "sort_order": 1,
                     *           "type": "dropdown",
                     *           "type_specific_meta": {
                     *             "choices": [
                     *               {
                     *                 "id": "62553abbb7def00b4e60ded5",
                     *                 "name": "Sample Dropdown Field Choice"
                     *               }
                     *             ],
                     *             "is_multi_select": false
                     *           }
                     *         },
                     *         {
                     *           "helper_text": "Sample helper text",
                     *           "identifier": "73e05aa2b54nda3c320242ac1d5291",
                     *           "is_readonly": false,
                     *           "is_required": false,
                     *           "label": "Input value",
                     *           "logic_rules": [],
                     *           "sort_order": 2,
                     *           "type": "simple_number",
                     *           "type_specific_meta": {
                     *             "decimal_places": 3,
                     *             "has_thousand_separator": false
                     *           }
                     *         },
                     *         {
                     *           "helper_text": "Sample helper text",
                     *           "identifier": "73e05aa2b54nda3c320242ac1d5291",
                     *           "is_readonly": false,
                     *           "is_required": false,
                     *           "label": "Input value",
                     *           "logic_rules": [],
                     *           "sort_order": 3,
                     *           "type": "percentage_number",
                     *           "type_specific_meta": {
                     *             "decimal_places": 3
                     *           }
                     *         },
                     *         {
                     *           "helper_text": "Sample helper text",
                     *           "identifier": "73e05aa2b54nda3c320242ac1d5291",
                     *           "is_readonly": false,
                     *           "is_required": false,
                     *           "label": "Input value",
                     *           "logic_rules": [],
                     *           "sort_order": 4,
                     *           "type": "currency_number",
                     *           "type_specific_meta": {
                     *             "currency_code": "USD",
                     *             "decimal_places": 2,
                     *             "has_thousand_separator": true
                     *           }
                     *         },
                     *         {
                     *           "helper_text": "Sample helper text",
                     *           "identifier": "73e05aa2b54nda3c320242ac1d5291",
                     *           "is_readonly": false,
                     *           "is_required": false,
                     *           "label": "Input value",
                     *           "logic_rules": [],
                     *           "sort_order": 5,
                     *           "type": "richtext",
                     *           "type_specific_meta": null
                     *         },
                     *         {
                     *           "helper_text": "Sample helper text",
                     *           "identifier": "73e05aa2b54nda3c320242ac1d5291",
                     *           "is_readonly": true,
                     *           "is_required": false,
                     *           "label": "Another section",
                     *           "logic_rules": [],
                     *           "sort_order": 5,
                     *           "type": "section",
                     *           "type_specific_meta": null
                     *         },
                     *         {
                     *           "helper_text": "Sample helper text",
                     *           "identifier": "73e05aa2b54nda3c320242ac1d5291",
                     *           "is_readonly": false,
                     *           "is_required": false,
                     *           "label": "select value",
                     *           "logic_rules": [],
                     *           "sort_order": 7,
                     *           "type": "radio_button",
                     *           "type_specific_meta": {
                     *             "choices": [
                     *               {
                     *                 "id": "62553abbb7def00b4e60ded5",
                     *                 "name": "Sample MultiChoice Field Choice"
                     *               }
                     *             ],
                     *             "is_multi_select": false
                     *           }
                     *         },
                     *         {
                     *           "helper_text": "Sample helper text",
                     *           "identifier": "73e05aa2b54nda3c320242ac1d5291",
                     *           "is_readonly": false,
                     *           "is_required": false,
                     *           "label": "select value",
                     *           "logic_rules": [],
                     *           "sort_order": 8,
                     *           "type": "label",
                     *           "type_specific_meta": {
                     *             "choices": [
                     *               {
                     *                 "id": "9119a313057e401189407116fcd3",
                     *                 "name": "Sample Label"
                     *               }
                     *             ],
                     *             "is_multi_select": false
                     *           }
                     *         }
                     *       ],
                     *       "id": "9119a313057e401189407116fcd3aa24",
                     *       "is_active": true,
                     *       "title": "Sample Template Title"
                     *     }
                     */
                    "application/json": components["schemas"]["TemplateResponse"];
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getUploadUrl: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched presigned upload URL and related meta fields */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description Related meta fields to upload a file. **Items within this field must be sent along as paylod of the POST request, along with the file, to the presigned upload URL** */
                        upload_meta_fields: {
                            /**
                             * @description Unique identifier of the file. **This key will be needed for various file / attachment addition to different CMP resources using endpoints e.g. `POST /v3/assets`, `POST /v3/campaigns/{id}/attachments` etc.**
                             * @example ce8995aea58b11ea8cd90242ac120005
                             */
                            key: string;
                            /** @example eyJleHBpcmF0aW9uIjogIjIwMjAtMDYtMDNUMTY6MTc6NTbbbbbbImNvbmRpdGlvbnMiOiBbeyJidWNrZXQiOiAibmV3c2NyZWQtY21wLWZpbGVzIn0sIHsia2V5IjogIjY3Y2NhMzcxYTVhZDExZWFiNDNhMGExMjlmMDhlZWQ5In0sIHsieC1hbXotYWxnb3JpdGhtIjogIkFXUzQtSE1BQy1TSEEyNTYifSwgeyJ4LWFtei1jcmVkZW50aWFsIjogIkFTSUFVTlVMQkwzV082RDdSWkZILzIwMjAwNjAzL3VzLWVhc3QtMS9zMy9hd3M0X3JlcXVlc3QifSwgeyJ4LWFtei1kYXRlIjogIjIwMjAwNjAzVDE1MTc1N1oifSwgeyJ4LWFtei1zZWN1cml0eS10b2tlbiI6ICJJUW9KYjNKcFoybHVYMlZqRUE4YUNYVnpMV1ZoYzNRdE1TSkhNRVVDSUQrc2FrelZIK05zN0Z4UmtsaW9CZFdldjdZRXlIcHBoVjEyNnRsL0ZadjNBaUVBODFOV1l5TUt1L0UvZ1FyaUlhaVhMaXpxSFFuNnlxTWlWSzQzeU9xYTUaaaaaaU1JZHhBQUdnd3pNRFF4TmpBMU16QXhOVFlpRERxMlN3dVJTR2ZGL2hGRXpDcVJBMlo4NmpkRzNQczMzY3c5bFpFeVBpbmdWSnV0WUQ0RXZuU291NFR1NG9WaGhEZUwyS0QxdFREZmMrNHJmVWRMQ2Y1SWpzb3VRYldMekIzczRXZ2xMQUJFOFJ0VUM5K0o1bVpJUTlYNzJVa0xOUjA0UURpVlliSjdxNS95S25VNUFSZWUzQXdqVHp3UWZPS3dsTnRFODYyVWo4MXlWbDhMdS81d3kwaU5hZFpvbkZPcVBONldJWURpdi9MVzlLKzhWNjRNS2E5bDhtcG1kVll3RWJ2dTRiMENJVFIxQ0tjUHlzVkFNODBKdnZ0aklvQTZXc3RDcWxiM0ZHZEFGVmlnaDdId3JKL3greFNZVU5YVzFvaFNlUExDNXFPSCtoZ2ZqRFVOSDMvbXhEV2V4dWhQbDFGSjF6VlpZQzBMY2U0Y1F6dEtPbTgzZFBqcWI4a09pQVBUWVB3MGFmYnl0QUQyVE8vTmhZdHlvU1BzL0xWS1FYWXhyYk5MZ3N6azRMUjdLeUEydlA5NkJRMll6aVhnWDhTTm13bklDcWRiY2dzZitvNnBWOXRlTlJIczdIS3N2Qm9HYVFESTVhUDQrcmQ1K0h0aWM3RUZqSDlDUkk1bnBpL3cvTVlodkc0dXljWDR4cEJSeE5iSURWbFpuYTBEaUkvVkxmV3JjOWtkUFBhdm10QVhtVHR1U2c4aFFVMnNzMSt0dEJqaE1NRGkzdllGT3VzQnJEMDMyc1BlMWx0QUNRSzlVQUxEUmFYT2VIWjNXbThTVjZlMTErWVhnd3JKZkloQlp2a053OGxHdkRuN2FSK2hXUkpPazljY3pHOC9DTzRqZk8xVlpkZ3N2YU5PaEhpc01sQXVOMVJXallXUGJWRkRoajV3ekZ0OTQrcmJwdFd3dk9vamtXVElUM0o5SWtSNTF6SWJhc1VqTW4rOHRiSjFJbHJnWVFOaDVaaUVPTGZ1bFVHZWl1L2xXdERFbFhKUy9sS0pUTnV1YU1DRnU0NnBocTlvVE9EMWtmSVYrUEFxZEZYV2xQbVZKK3lXeGRLellSV1creDdERjJ6dnZmZm1nc0JBWU01eDh6aEJ3YWIrSXkvTGR6eWdQVUFnaTFTUE9oSktFRjNBejU4aVArdGFDbUxhdmo5cGFRPT0ifV19 */
                            policy: string;
                            /** @example AWS4-HMAC-SHA256 */
                            "x-amz-algorithm": string;
                            /** @example ASIAAAAAAA3WO6D7RZFH/20343603/us-east-1/s3/aws4_request */
                            "x-amz-credential": string;
                            /** @example 20200603T151757Z */
                            "x-amz-date": string;
                            /** @example IQoJb3JpZ2luX2VjEA8ccccccWVhc3QtMSJHMEUCID+sakzVH+Ns7FxRklioBdWev7YEyHpphV126tl/FZv3AiEA81NWYyMKu/E/gQriIaiXLizqHQn6yqMiVK43yOqa52AqtAMIdxAAGgwzMDQxNjA1MzAxNTYiDDq2SwuRSGfFddddddqRA2Z86jdG3Ps33cw9lZEyPingVJutYD4EvnSou4Tu4oVhhDeL2KD1tTDfc+4rfUdLCf5IjsouQbWLzB3s4WglLABE8RtUC9+J5mZIQ9X72UkLNR04QDiVYbJ7q5/yKnU5ARee3AwjTzwQfOKwlNtE862Uj81yVl8Lu/5wy0iNadZonFOqPN6WIYDiv/LW9K+8V64MKa9l8mpmdVYwEbvu4b0CITR1CKcPysVAM80JvvtjIoA6WstCqlb3FGdAFVigh7HwrJ/x+xSYUNXW1ohSePLC5qOH+hgfjDUNH3/mxDWexuhPl1FJ1zVZYC0Lce4cQztKOm83dPjqb8kOiAPTYPw0afbytAD2TO/NhYtyoSPs/LVKQXYxrbNLgszk4LR7KyA2vP96BQ2YziXgX8SNmwnICqdbcgsf+o6pV9teNRHs7HKsvBoGaQDI5aP4+rd5+Htic7EFjH9CRI5npi/w/MYhvG4uycX4xpBRxNbIDVlZna0DiI/VLfWrc9kdPPavmtAXmTtuSg8hQU2ss1+ttBjhMMDi3vYFOusBrD032sPe1ltACQK9UALDRaXOeHZ3Wm8SV6e11+YXgwrJfIhBZvkNw8lGvDn7aR+hWRJOk9cczG8/CO4jfO1VZdgsvaNOhHisMlAuN1RWjYWPbVFDhj5wzFt94+rbptWwvOojkWTIT3J9IkR51zIbasUjMn+8tbJ1IlrgYQNh5ZiEOLfulUGeiu/lWtDElXJS/lKJTNuuaMCFu46phq9oTOD1kfIV+PAqdFXWlPmVJ+yWxdKzYRWW+x7DF2zvvffmgsBAYM5x8zhBwab+Iy/LdzygPUAgi1SPOhJKEF3Az58iP+taCmLavj9paQ== */
                            "x-amz-security-token"?: string;
                            /** @example 0e67ede744d4aed62aaaaaa8e87648b3e98a3e611ccccccc89963bcf70bd02f9 */
                            "x-amz-signature": string;
                        };
                        /**
                         * @description Presigned URL to use for uploading a file with an **HTTP POST** request as `multipart/form-data`
                         * @example https://cmp-files.s3-accelerate.amazonaws.com/
                         */
                        url: string;
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    listUsers: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of users */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["UserListResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    findUserByEmail: {
        parameters: {
            query: {
                /** @example john.doe@example.com */
                email: string;
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Redirect to the discovered user */
            302: {
                headers: {
                    /**
                     * @description URL of the discovered user
                     * @example https://api.cmp.optimizely.com/v3/users/5fe4925ef8a9378056bf4e37
                     */
                    Location?: string;
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getUser: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5fe4925ef8a9378056bf4e37 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched user */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["UserResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    getVideo: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6633 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched video */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LibraryVideo"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    deleteVideo: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6633 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Deleted the video */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    updateVideo: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5d7f910551b00a722e0418830cee6631 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the video */
        requestBody: {
            content: {
                "application/json": components["schemas"]["LibraryVideoUpdateRequest"];
            };
        };
        responses: {
            /** @description Video Updated */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["LibraryVideo"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listWorkRequests: {
        parameters: {
            query?: {
                /**
                 * @example [
                 *       "1265ca474ba987cc925164aa7",
                 *       "606995f714ace404593cce4c"
                 *     ]
                 */
                created_by?: string;
                /**
                 * @example [
                 *       "Accepted",
                 *       "Submitted"
                 *     ]
                 */
                status?: "Submitted" | "Accepted" | "Completed" | "Declined";
                /** @example due_date */
                order_by?: "priority" | "status" | "created_at" | "due_date";
                /** @example desc */
                order_as?: "asc" | "desc";
                /**
                 * @description Unique identifiers of the templates. You can append multiple times, for example `template_ids=template1&template_ids=template2`.
                 * @example [
                 *       "1265ca474ba987cc925164aa7",
                 *       "606995f714ace404593cce4c"
                 *     ]
                 */
                template_ids?: string[];
                /**
                 * @description Date and time as the lower limit to filter work requests by `created_at`, in ISO 8601 UTC format
                 * @example 2018-11-30T13:32:44Z
                 */
                created_at__from?: string;
                /**
                 * @description Date and time as the upper limit to filter work requests by `created_at`, in ISO 8601 UTC format
                 * @example 2018-12-31T23:59:59Z
                 */
                created_at__to?: string;
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched work requests */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of work requests */
                        data: components["schemas"]["WorkRequestResponse"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/work-requests?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    createWorkRequest: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** @description Payload to create a work request */
        requestBody: {
            content: {
                /**
                 * @example {
                 *       "assignees": [
                 *         "74ba987cc9251689abc123"
                 *       ],
                 *       "form_fields": [
                 *         {
                 *           "identifier": "title",
                 *           "type": "text",
                 *           "values": [
                 *             "Some title"
                 *           ]
                 *         },
                 *         {
                 *           "identifier": "brief",
                 *           "type": "brief",
                 *           "values": [
                 *             {
                 *               "type": "attachment_brief",
                 *               "value": {
                 *                 "key": "10d5c89aebe611eca11302420ac80002",
                 *                 "name": "Attachment"
                 *               }
                 *             }
                 *           ]
                 *         },
                 *         {
                 *           "identifier": "630c486a90546926f034ea28",
                 *           "type": "richtext",
                 *           "values": [
                 *             "<b>rich</b> text"
                 *           ]
                 *         },
                 *         {
                 *           "identifier": "630c479690546926f034ea13",
                 *           "type": "label",
                 *           "values": [
                 *             "630c479690546926f034ea15"
                 *           ]
                 *         },
                 *         {
                 *           "identifier": "630c47be90546926f034ea17",
                 *           "type": "dropdown",
                 *           "values": [
                 *             "630c47be90546926f034ea18"
                 *           ]
                 *         },
                 *         {
                 *           "identifier": "630c482b90546926f034ea22",
                 *           "type": "simple_number",
                 *           "values": [
                 *             1234.345678
                 *           ]
                 *         },
                 *         {
                 *           "identifier": "630c484b90546926f034ea23",
                 *           "type": "percentage_number",
                 *           "values": [
                 *             23
                 *           ]
                 *         },
                 *         {
                 *           "identifier": "630c481b90546926f034ea21",
                 *           "type": "currency_number",
                 *           "values": [
                 *             123.56
                 *           ]
                 *         },
                 *         {
                 *           "identifier": "630c485e90546926f034ea24",
                 *           "type": "checkbox",
                 *           "values": [
                 *             "630c485e90546926f034ea25"
                 *           ]
                 *         },
                 *         {
                 *           "identifier": "630c47d290546926f034ea1b",
                 *           "type": "radio_button",
                 *           "values": [
                 *             "630c47d290546926f034ea1c"
                 *           ]
                 *         },
                 *         {
                 *           "identifier": "creative_assets",
                 *           "type": "file",
                 *           "values": [
                 *             {
                 *               "key": "creative-asset-key-1",
                 *               "name": "Creative Asset 1"
                 *             },
                 *             {
                 *               "key": "creative-asset-key-2",
                 *               "name": "Creative Asset 2"
                 *             }
                 *           ]
                 *         },
                 *         {
                 *           "identifier": "due_date",
                 *           "type": "date",
                 *           "values": [
                 *             "2022-02-20T12:02:03.000000Z"
                 *           ]
                 *         },
                 *         {
                 *           "identifier": "media_links",
                 *           "type": "text_area",
                 *           "values": [
                 *             "some text\nmore text"
                 *           ]
                 *         }
                 *       ],
                 *       "template_id": "5c89aebe611eca11302420af"
                 *     }
                 */
                "application/json": components["schemas"]["WorkRequestCreateRequest"];
            };
        };
        responses: {
            /** @description Created work request */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["WorkRequestResponse"];
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    getWorkRequest: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 9119a313057e401189407116fcd3aa24 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched work request */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["WorkRequestResponse"];
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    updateWorkRequest: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 9119a313057e401189407116fcd3aa24 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the work request */
        requestBody: {
            content: {
                "application/json": components["schemas"]["WorkRequestUpdateRequest"];
            };
        };
        responses: {
            /** @description Work request updated */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listWorkRequestApprovedAssets: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path: {
                /** @example da80cfd7bbf84959a8a981acbad996b3 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of approved assets for the work request */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of work request approved assets */
                        data: components["schemas"]["WorkRequestApprovedAssetResponse"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/work-requests/da80cfd7bbf84959a8a981acbad996b3/approved-assets?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    addAttachmentToWorkRequest: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example da80cfd7bbf84959a8a981acbad996b3 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to create an attachment */
        requestBody: {
            content: {
                "application/json": components["schemas"]["AttachmentRequest"];
            };
        };
        responses: {
            /** @description Created attachment for the work request */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AttachmentResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    createCampaignFromWorkRequest: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example da80cfd7bbf84959a8a981acbad996b3 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to create a campaign */
        requestBody: {
            content: {
                "application/json": components["schemas"]["WorkRequestCampaignRequest"];
            };
        };
        responses: {
            /** @description Created a campaign from a work request */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["WorkRequestCampaignResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    listWorkRequestComments: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path: {
                /** @example da80cfd7bbf84959a8a981acbad996b3 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of comments for the work request */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of work request comments */
                        data: components["schemas"]["WorkRequestCommentResponse"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/work-requests/da80cfd7bbf84959a8a981acbad996b3/comments?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    addCommentToWorkRequest: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example da80cfd7bbf84959a8a981acbad996b3 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to add comment */
        requestBody: {
            content: {
                "application/json": components["schemas"]["CommentWithReplyCreateRequest"];
            };
        };
        responses: {
            /** @description Created comment for the work request */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["WorkRequestCommentResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    createWorkRequestCreativeAsset: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example da80cfd7bbf84959a8a981acbad996b3 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to create creative assets */
        requestBody: {
            content: {
                "application/json": components["schemas"]["CreativeAssetRequest"];
            };
        };
        responses: {
            /** @description Created creative asset for the work request */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["CreativeAssetResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    listWorkRequestRelatedResources: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path: {
                /** @example da80cfd7bbf84959a8a981acbad996b3 */
                id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched work request related resources */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of work request related resources */
                        data: components["schemas"]["WorkRequestRelatedResourceResponse"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/work-requests/5fe4925ef8a9378056bf4e37/related-resources?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    createTaskFromWorkRequest: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example da80cfd7bbf84959a8a981acbad996b3 */
                id: string;
            };
            cookie?: never;
        };
        /** @description Payload to create a task */
        requestBody: {
            content: {
                "application/json": components["schemas"]["WorkRequestTaskRequest"];
            };
        };
        responses: {
            /** @description Created a task from a work request */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["WorkRequestTaskResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    deleteWorkRequestAttachment: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7674a2a69f8d4eb48e8d4dafaebc7238 */
                work_request_id: string;
                /** @example 62f2052962f36e010b9471f1 */
                attachment_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Deleted the attachment */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    getWorkRequestComment: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example da80cfd7bbf84959a8a981acbad996b3 */
                work_request_id: string;
                /** @example 5fe38c39574b52a62a089239 */
                comment_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched work request comment */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["WorkRequestCommentResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
    deleteWorkRequestCreativeAsset: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7674a2a69f8d4eb48e8d4dafaebc7238 */
                work_request_id: string;
                /** @example 209d2d281e2b11edbac602420ac80012 */
                creative_asset_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Deleted the creative asset */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    updateWorkRequestFormField: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 7674a2a69f8d4eb48e8d4dafaebc7238 */
                work_request_id: string;
                /** @example 209d2d281e2b11edbac602420ac80012 */
                form_field_identifier: string;
            };
            cookie?: never;
        };
        /** @description Payload to update the work request form field */
        requestBody: {
            content: {
                "application/json": components["schemas"]["WorkRequestFormFieldUpdateRequest"];
            };
        };
        responses: {
            /** @description Updated work request form field */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["WorkRequestRequestFormFieldUpdateResponse"];
                };
            };
            400: components["responses"]["ClientError"];
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
            422: components["responses"]["UnprocessableEntity"];
        };
    };
    listWorkflows: {
        parameters: {
            query?: {
                /**
                 * @description Starting index of results (zero indexed)
                 * @example 5
                 */
                offset?: components["parameters"]["offset"];
                /**
                 * @description Number of results to return per page
                 * @example 15
                 */
                page_size?: components["parameters"]["page_size"];
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description List of fetched workflows */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        /** @description List of workflows */
                        data: components["schemas"]["WorkflowListResponseItem"][];
                        pagination: components["schemas"]["Pagination"] & {
                            /** @example https://api.cmp.optimizely.com/v3/workflows?offset=10&page_size=10 */
                            next?: string | null;
                        };
                    };
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
        };
    };
    getWorkflow: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                /** @example 5fe4925ef8a9378056bf4e37 */
                workflow_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fetched workflow */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["WorkflowResponse"];
                };
            };
            401: components["responses"]["Unauthorized"];
            403: components["responses"]["Forbidden"];
            404: components["responses"]["NotFound"];
        };
    };
}
