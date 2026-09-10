// Auto-generated - DO NOT EDIT

package schema

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for AssetContentType.
const (
	AssetContentTypeAPIURL   AssetContentType = "api_url"
	AssetContentTypeHTMLBody AssetContentType = "html_body"
	AssetContentTypeURL      AssetContentType = "url"
)

// Defines values for AssetFieldTypes.
const (
	AssetFieldTypesCheckbox         AssetFieldTypes = "checkbox"
	AssetFieldTypesCurrencyNumber   AssetFieldTypes = "currency_number"
	AssetFieldTypesDate             AssetFieldTypes = "date"
	AssetFieldTypesDropdown         AssetFieldTypes = "dropdown"
	AssetFieldTypesLabel            AssetFieldTypes = "label"
	AssetFieldTypesPercentageNumber AssetFieldTypes = "percentage_number"
	AssetFieldTypesRadioButton      AssetFieldTypes = "radio_button"
	AssetFieldTypesRichText         AssetFieldTypes = "rich_text"
	AssetFieldTypesSimpleNumber     AssetFieldTypes = "simple_number"
	AssetFieldTypesText             AssetFieldTypes = "text"
	AssetFieldTypesTextArea         AssetFieldTypes = "text_area"
)

// Defines values for AssetLineageResponseUsedIn.
const (
	AssetLineageResponseUsedInExternal AssetLineageResponseUsedIn = "external"
)

// Defines values for AssetPermissionBulkCreateRequestPermissionsAccessType.
const (
	AssetPermissionBulkCreateRequestPermissionsAccessTypeEdit AssetPermissionBulkCreateRequestPermissionsAccessType = "edit"
	AssetPermissionBulkCreateRequestPermissionsAccessTypeView AssetPermissionBulkCreateRequestPermissionsAccessType = "view"
)

// Defines values for AssetPermissionBulkCreateRequestType.
const (
	AssetPermissionBulkCreateRequestTypeTeam AssetPermissionBulkCreateRequestType = "team"
	AssetPermissionBulkCreateRequestTypeUser AssetPermissionBulkCreateRequestType = "user"
)

// Defines values for AssetPermissionUpdateRequestAccessType.
const (
	AssetPermissionUpdateRequestAccessTypeEdit AssetPermissionUpdateRequestAccessType = "edit"
	AssetPermissionUpdateRequestAccessTypeView AssetPermissionUpdateRequestAccessType = "view"
)

// Defines values for AssetResponseType.
const (
	AssetResponseTypeArticle           AssetResponseType = "article"
	AssetResponseTypeImage             AssetResponseType = "image"
	AssetResponseTypeRawFile           AssetResponseType = "raw_file"
	AssetResponseTypeStructuredContent AssetResponseType = "structured_content"
	AssetResponseTypeVideo             AssetResponseType = "video"
)

// Defines values for BaseFieldDefinitionType.
const (
	BaseFieldDefinitionTypeBoolean  BaseFieldDefinitionType = "boolean"
	BaseFieldDefinitionTypeJSON     BaseFieldDefinitionType = "json"
	BaseFieldDefinitionTypeLocation BaseFieldDefinitionType = "location"
	BaseFieldDefinitionTypeURL      BaseFieldDefinitionType = "url"
)

// Defines values for BasePermissionsResponseSchemaAccessType.
const (
	BasePermissionsResponseSchemaAccessTypeComment BasePermissionsResponseSchemaAccessType = "comment"
	BasePermissionsResponseSchemaAccessTypeDelete  BasePermissionsResponseSchemaAccessType = "delete"
	BasePermissionsResponseSchemaAccessTypeEdit    BasePermissionsResponseSchemaAccessType = "edit"
	BasePermissionsResponseSchemaAccessTypeView    BasePermissionsResponseSchemaAccessType = "view"
)

// Defines values for BasePermissionsResponseSchemaType.
const (
	BasePermissionsResponseSchemaTypeOrganization BasePermissionsResponseSchemaType = "organization"
	BasePermissionsResponseSchemaTypeTeam         BasePermissionsResponseSchemaType = "team"
	BasePermissionsResponseSchemaTypeUser         BasePermissionsResponseSchemaType = "user"
)

// Defines values for BaseTemplateFormFieldType.
const (
	BaseTemplateFormFieldTypeBrief            BaseTemplateFormFieldType = "brief"
	BaseTemplateFormFieldTypeCheckbox         BaseTemplateFormFieldType = "checkbox"
	BaseTemplateFormFieldTypeCurrencyNumber   BaseTemplateFormFieldType = "currency_number"
	BaseTemplateFormFieldTypeDate             BaseTemplateFormFieldType = "date"
	BaseTemplateFormFieldTypeDropdown         BaseTemplateFormFieldType = "dropdown"
	BaseTemplateFormFieldTypeFile             BaseTemplateFormFieldType = "file"
	BaseTemplateFormFieldTypeInstruction      BaseTemplateFormFieldType = "instruction"
	BaseTemplateFormFieldTypeLabel            BaseTemplateFormFieldType = "label"
	BaseTemplateFormFieldTypePercentageNumber BaseTemplateFormFieldType = "percentage_number"
	BaseTemplateFormFieldTypeRadioButton      BaseTemplateFormFieldType = "radio_button"
	BaseTemplateFormFieldTypeRichtext         BaseTemplateFormFieldType = "richtext"
	BaseTemplateFormFieldTypeSection          BaseTemplateFormFieldType = "section"
	BaseTemplateFormFieldTypeSimpleNumber     BaseTemplateFormFieldType = "simple_number"
	BaseTemplateFormFieldTypeText             BaseTemplateFormFieldType = "text"
	BaseTemplateFormFieldTypeTextArea         BaseTemplateFormFieldType = "text_area"
)

// Defines values for BaseTemplateResponseApplicableTo.
const (
	BaseTemplateResponseApplicableToCampaignBrief BaseTemplateResponseApplicableTo = "campaign_brief"
	BaseTemplateResponseApplicableToTaskBrief     BaseTemplateResponseApplicableTo = "task_brief"
	BaseTemplateResponseApplicableToWorkRequest   BaseTemplateResponseApplicableTo = "work_request"
)

// Defines values for CampaignBriefResponseType.
const (
	CampaignBriefResponseTypeAttachment CampaignBriefResponseType = "attachment"
	CampaignBriefResponseTypeTemplate   CampaignBriefResponseType = "template"
	CampaignBriefResponseTypeText       CampaignBriefResponseType = "text"
)

// Defines values for CampaignCreateRequestColor.
const (
	Hash212121 CampaignCreateRequestColor = "#212121"
	Hash21B7EC CampaignCreateRequestColor = "#21B7EC"
	Hash2E2A26 CampaignCreateRequestColor = "#2e2a26"
	Hash333    CampaignCreateRequestColor = "#333"
	Hash36C5A3 CampaignCreateRequestColor = "#36c5a3"
	Hash3A97Be CampaignCreateRequestColor = "#3a97be"
	Hash3AB0C9 CampaignCreateRequestColor = "#3AB0C9"
	Hash449C6C CampaignCreateRequestColor = "#449C6C"
	Hash4C557E CampaignCreateRequestColor = "#4C557E"
	Hash4D9B93 CampaignCreateRequestColor = "#4D9B93"
	Hash4Db2Fc CampaignCreateRequestColor = "#4db2fc"
	Hash4Fd5Bd CampaignCreateRequestColor = "#4fd5bd"
	Hash5Ab5D0 CampaignCreateRequestColor = "#5ab5d0"
	Hash5F9DCE CampaignCreateRequestColor = "#5F9DCE"
	Hash63B05F CampaignCreateRequestColor = "#63B05F"
	Hash6456B7 CampaignCreateRequestColor = "#6456b7"
	Hash669BA8 CampaignCreateRequestColor = "#669BA8"
	Hash717CB9 CampaignCreateRequestColor = "#717CB9"
	Hash79C8Fd CampaignCreateRequestColor = "#79c8fd"
	Hash7ABEB2 CampaignCreateRequestColor = "#7ABEB2"
	Hash7Aa8B6 CampaignCreateRequestColor = "#7aa8b6"
	Hash7Abeb2 CampaignCreateRequestColor = "#7abeb2"
	Hash8AA872 CampaignCreateRequestColor = "#8AA872"
	Hash8Ea1A7 CampaignCreateRequestColor = "#8ea1a7"
	Hash9086Cc CampaignCreateRequestColor = "#9086cc"
	Hash908Bae CampaignCreateRequestColor = "#908bae"
	Hash908E9B CampaignCreateRequestColor = "#908e9b"
	Hash96B1Ac CampaignCreateRequestColor = "#96b1ac"
	Hash97Bed8 CampaignCreateRequestColor = "#97bed8"
	Hash99896A CampaignCreateRequestColor = "#99896A"
	Hash9A8692 CampaignCreateRequestColor = "#9a8692"
	Hash9Bc94E CampaignCreateRequestColor = "#9bc94e"
	HashA25555 CampaignCreateRequestColor = "#A25555"
	HashA29F55 CampaignCreateRequestColor = "#A29F55"
	HashA89Bac CampaignCreateRequestColor = "#a89bac"
	HashA8B9C3 CampaignCreateRequestColor = "#a8b9c3"
	HashA959Cc CampaignCreateRequestColor = "#a959cc"
	HashACC050 CampaignCreateRequestColor = "#ACC050"
	HashAF8EC3 CampaignCreateRequestColor = "#AF8EC3"
	HashAe7C98 CampaignCreateRequestColor = "#ae7c98"
	HashB05F8C CampaignCreateRequestColor = "#B05F8C"
	HashB094Be CampaignCreateRequestColor = "#b094be"
	HashB29F9F CampaignCreateRequestColor = "#b29f9f"
	HashB5A1Ae CampaignCreateRequestColor = "#b5a1ae"
	HashB8D877 CampaignCreateRequestColor = "#b8d877"
	HashB9A1Ac CampaignCreateRequestColor = "#b9a1ac"
	HashB9A29E CampaignCreateRequestColor = "#b9a29e"
	HashBd417F CampaignCreateRequestColor = "#bd417f"
	HashBece9D CampaignCreateRequestColor = "#bece9d"
	HashC0B0A1 CampaignCreateRequestColor = "#c0b0a1"
	HashC187Db CampaignCreateRequestColor = "#c187db"
	HashC1C8B5 CampaignCreateRequestColor = "#c1c8b5"
	HashC79697 CampaignCreateRequestColor = "#c79697"
	HashC897B4 CampaignCreateRequestColor = "#c897b4"
	HashCBB039 CampaignCreateRequestColor = "#CBB039"
	HashCE515B CampaignCreateRequestColor = "#CE515B"
	HashCcccdc CampaignCreateRequestColor = "#ccccdc"
	HashCf69A3 CampaignCreateRequestColor = "#cf69a3"
	HashD096Af CampaignCreateRequestColor = "#d096af"
	HashD3998F CampaignCreateRequestColor = "#d3998f"
	HashD3Ac85 CampaignCreateRequestColor = "#d3ac85"
	HashD4C8B0 CampaignCreateRequestColor = "#d4c8b0"
	HashDE7009 CampaignCreateRequestColor = "#DE7009"
	HashDe55A9 CampaignCreateRequestColor = "#de55a9"
	HashE2575B CampaignCreateRequestColor = "#e2575b"
	HashE3C68D CampaignCreateRequestColor = "#e3c68d"
	HashE4844C CampaignCreateRequestColor = "#E4844C"
	HashE785C1 CampaignCreateRequestColor = "#e785c1"
	HashEB7777 CampaignCreateRequestColor = "#EB7777"
	HashEE9808 CampaignCreateRequestColor = "#EE9808"
	HashEa868A CampaignCreateRequestColor = "#ea868a"
	HashEc2E3B CampaignCreateRequestColor = "#ec2e3b"
	HashEf803C CampaignCreateRequestColor = "#ef803c"
	HashF16281 CampaignCreateRequestColor = "#F16281"
	HashF3A35A CampaignCreateRequestColor = "#f3a35a"
	HashF45097 CampaignCreateRequestColor = "#f45097"
	HashF781B5 CampaignCreateRequestColor = "#f781b5"
	HashFCD404 CampaignCreateRequestColor = "#FCD404"
	HashFc594B CampaignCreateRequestColor = "#fc594b"
	HashFd8876 CampaignCreateRequestColor = "#fd8876"
	HashFdad3C CampaignCreateRequestColor = "#fdad3c"
	HashFec454 CampaignCreateRequestColor = "#fec454"
)

// Defines values for CampaignFieldTypes.
const (
	CampaignFieldTypesCheckbox         CampaignFieldTypes = "checkbox"
	CampaignFieldTypesCurrencyNumber   CampaignFieldTypes = "currency_number"
	CampaignFieldTypesDate             CampaignFieldTypes = "date"
	CampaignFieldTypesDropdown         CampaignFieldTypes = "dropdown"
	CampaignFieldTypesImage            CampaignFieldTypes = "image"
	CampaignFieldTypesLabel            CampaignFieldTypes = "label"
	CampaignFieldTypesPercentageNumber CampaignFieldTypes = "percentage_number"
	CampaignFieldTypesRadioButton      CampaignFieldTypes = "radio_button"
	CampaignFieldTypesRichText         CampaignFieldTypes = "rich_text"
	CampaignFieldTypesSimpleNumber     CampaignFieldTypes = "simple_number"
	CampaignFieldTypesText             CampaignFieldTypes = "text"
	CampaignFieldTypesTextArea         CampaignFieldTypes = "text_area"
	CampaignFieldTypesVideo            CampaignFieldTypes = "video"
)

// Defines values for CampaignResponseStatus.
const (
	CampaignResponseStatusAtRisk     CampaignResponseStatus = "At Risk"
	CampaignResponseStatusComplete   CampaignResponseStatus = "Complete"
	CampaignResponseStatusNotStarted CampaignResponseStatus = "Not Started"
	CampaignResponseStatusOffTrack   CampaignResponseStatus = "Off Track"
	CampaignResponseStatusOnTrack    CampaignResponseStatus = "On Track"
)

// Defines values for ChoiceDisplayOption.
const (
	ChoiceDisplayOptionCheckbox ChoiceDisplayOption = "checkbox"
	ChoiceDisplayOptionDropdown ChoiceDisplayOption = "dropdown"
	ChoiceDisplayOptionRadio    ChoiceDisplayOption = "radio"
	ChoiceDisplayOptionTag      ChoiceDisplayOption = "tag"
)

// Defines values for ContentGraphResponseInstanceCategory.
const (
	Existing ContentGraphResponseInstanceCategory = "existing"
	New      ContentGraphResponseInstanceCategory = "new"
)

// Defines values for ContentTypeFieldEmbedMixConfig.
const (
	ContentTypeFieldEmbedMixConfigN1 ContentTypeFieldEmbedMixConfig = 1
	ContentTypeFieldEmbedMixConfigN2 ContentTypeFieldEmbedMixConfig = 2
	ContentTypeFieldEmbedMixConfigN3 ContentTypeFieldEmbedMixConfig = 3
)

// Defines values for ContentTypeListingOption.
const (
	ContentTypeListingOptionN1 ContentTypeListingOption = "1"
	ContentTypeListingOptionN2 ContentTypeListingOption = "2"
	ContentTypeListingOptionN3 ContentTypeListingOption = "3"
)

// Defines values for DetailedAssetRenditionResponseStatus.
const (
	DetailedAssetRenditionResponseStatusDone       DetailedAssetRenditionResponseStatus = "Done"
	DetailedAssetRenditionResponseStatusError      DetailedAssetRenditionResponseStatus = "Error"
	DetailedAssetRenditionResponseStatusInProgress DetailedAssetRenditionResponseStatus = "InProgress"
)

// Defines values for FieldType.
const (
	FieldTypeBoolean      FieldType = "boolean"
	FieldTypeChoice       FieldType = "choice"
	FieldTypeContentType  FieldType = "content-type"
	FieldTypeDatetime     FieldType = "datetime"
	FieldTypeJSON         FieldType = "json"
	FieldTypeLibraryAsset FieldType = "library-asset"
	FieldTypeLocation     FieldType = "location"
	FieldTypeNumber       FieldType = "number"
	FieldTypeRichText     FieldType = "rich-text"
	FieldTypeTextField    FieldType = "text-field"
	FieldTypeURL          FieldType = "url"
)

// Defines values for FolderPermissionBulkCreateRequestPermissionsAccessType.
const (
	FolderPermissionBulkCreateRequestPermissionsAccessTypeEdit FolderPermissionBulkCreateRequestPermissionsAccessType = "edit"
	FolderPermissionBulkCreateRequestPermissionsAccessTypeView FolderPermissionBulkCreateRequestPermissionsAccessType = "view"
)

// Defines values for FolderPermissionBulkCreateRequestType.
const (
	FolderPermissionBulkCreateRequestTypeTeam FolderPermissionBulkCreateRequestType = "team"
	FolderPermissionBulkCreateRequestTypeUser FolderPermissionBulkCreateRequestType = "user"
)

// Defines values for FolderPermissionUpdateRequestAccessType.
const (
	Edit FolderPermissionUpdateRequestAccessType = "edit"
	View FolderPermissionUpdateRequestAccessType = "view"
)

// Defines values for GetMultipartUploadStatusResponseStatus.
const (
	UPLOADCOMPLETIONFAILED     GetMultipartUploadStatusResponseStatus = "UPLOAD_COMPLETION_FAILED"
	UPLOADCOMPLETIONINPROGRESS GetMultipartUploadStatusResponseStatus = "UPLOAD_COMPLETION_IN_PROGRESS"
	UPLOADCOMPLETIONNOTSTARTED GetMultipartUploadStatusResponseStatus = "UPLOAD_COMPLETION_NOT_STARTED"
	UPLOADCOMPLETIONSUCCEEDED  GetMultipartUploadStatusResponseStatus = "UPLOAD_COMPLETION_SUCCEEDED"
)

// Defines values for LabelGroupSourceOrgType.
const (
	Current LabelGroupSourceOrgType = "current"
	Related LabelGroupSourceOrgType = "related"
)

// Defines values for LibraryAssetType.
const (
	LibraryAssetTypeArticle           LibraryAssetType = "article"
	LibraryAssetTypeImage             LibraryAssetType = "image"
	LibraryAssetTypeRawFile           LibraryAssetType = "raw_file"
	LibraryAssetTypeStructuredContent LibraryAssetType = "structured_content"
	LibraryAssetTypeVideo             LibraryAssetType = "video"
)

// Defines values for LibraryAssetVersionResponseContentType.
const (
	LibraryAssetVersionResponseContentTypeURL LibraryAssetVersionResponseContentType = "url"
)

// Defines values for LibraryAssetVersionResponseType.
const (
	LibraryAssetVersionResponseTypeImage   LibraryAssetVersionResponseType = "image"
	LibraryAssetVersionResponseTypeRawFile LibraryAssetVersionResponseType = "raw_file"
	LibraryAssetVersionResponseTypeVideo   LibraryAssetVersionResponseType = "video"
)

// Defines values for ListTaskAssetsResponseDataType.
const (
	ListTaskAssetsResponseDataTypeArticle           ListTaskAssetsResponseDataType = "article"
	ListTaskAssetsResponseDataTypeImage             ListTaskAssetsResponseDataType = "image"
	ListTaskAssetsResponseDataTypeRawFile           ListTaskAssetsResponseDataType = "raw_file"
	ListTaskAssetsResponseDataTypeStructuredContent ListTaskAssetsResponseDataType = "structured_content"
	ListTaskAssetsResponseDataTypeVideo             ListTaskAssetsResponseDataType = "video"
)

// Defines values for PublishingEventMetadataBulkCreateResponseErrorsErrorCode.
const (
	CanonicalLinkError                       PublishingEventMetadataBulkCreateResponseErrorsErrorCode = "canonical-link-error"
	DomainNotWhitelisted                     PublishingEventMetadataBulkCreateResponseErrorsErrorCode = "domain-not-whitelisted"
	DuplicateMetadata                        PublishingEventMetadataBulkCreateResponseErrorsErrorCode = "duplicate-metadata"
	InvalidStatus                            PublishingEventMetadataBulkCreateResponseErrorsErrorCode = "invalid-status"
	LocaleNotAllowed                         PublishingEventMetadataBulkCreateResponseErrorsErrorCode = "locale-not-allowed"
	MetadataExists                           PublishingEventMetadataBulkCreateResponseErrorsErrorCode = "metadata-exists"
	MissingLocale                            PublishingEventMetadataBulkCreateResponseErrorsErrorCode = "missing-locale"
	MissingPublicURL                         PublishingEventMetadataBulkCreateResponseErrorsErrorCode = "missing-public-url"
	MissingPublishingDestinationUpdatedAt    PublishingEventMetadataBulkCreateResponseErrorsErrorCode = "missing-publishing-destination-updated-at"
	PublicURLNotAllowed                      PublishingEventMetadataBulkCreateResponseErrorsErrorCode = "public-url-not-allowed"
	PublishingDestinationUpdatedAtNotAllowed PublishingEventMetadataBulkCreateResponseErrorsErrorCode = "publishing-destination-updated-at-not-allowed"
	UnknownAsset                             PublishingEventMetadataBulkCreateResponseErrorsErrorCode = "unknown-asset"
)

// Defines values for PublishingEventMetadataCreateRequestStatus.
const (
	PublishingEventMetadataCreateRequestStatusFailed      PublishingEventMetadataCreateRequestStatus = "failed"
	PublishingEventMetadataCreateRequestStatusPublished   PublishingEventMetadataCreateRequestStatus = "published"
	PublishingEventMetadataCreateRequestStatusSynced      PublishingEventMetadataCreateRequestStatus = "synced"
	PublishingEventMetadataCreateRequestStatusUnpublished PublishingEventMetadataCreateRequestStatus = "unpublished"
)

// Defines values for PublishingEventMetadataResponseAssetType.
const (
	PublishingEventMetadataResponseAssetTypeArticle           PublishingEventMetadataResponseAssetType = "article"
	PublishingEventMetadataResponseAssetTypeImage             PublishingEventMetadataResponseAssetType = "image"
	PublishingEventMetadataResponseAssetTypeRawFile           PublishingEventMetadataResponseAssetType = "raw_file"
	PublishingEventMetadataResponseAssetTypeStructuredContent PublishingEventMetadataResponseAssetType = "structured_content"
	PublishingEventMetadataResponseAssetTypeVideo             PublishingEventMetadataResponseAssetType = "video"
)

// Defines values for PublishingEventMetadataResponseStatus.
const (
	PublishingEventMetadataResponseStatusFailed      PublishingEventMetadataResponseStatus = "failed"
	PublishingEventMetadataResponseStatusLessThanNil PublishingEventMetadataResponseStatus = "<nil>"
	PublishingEventMetadataResponseStatusPublished   PublishingEventMetadataResponseStatus = "published"
	PublishingEventMetadataResponseStatusSynced      PublishingEventMetadataResponseStatus = "synced"
	PublishingEventMetadataResponseStatusUnpublished PublishingEventMetadataResponseStatus = "unpublished"
)

// Defines values for RelatedAssetItemContentType.
const (
	RelatedAssetItemContentTypeAPIURL   RelatedAssetItemContentType = "api_url"
	RelatedAssetItemContentTypeHTMLBody RelatedAssetItemContentType = "html_body"
	RelatedAssetItemContentTypeURL      RelatedAssetItemContentType = "url"
)

// Defines values for RelatedAssetItemType.
const (
	RelatedAssetItemTypeArticle           RelatedAssetItemType = "article"
	RelatedAssetItemTypeImage             RelatedAssetItemType = "image"
	RelatedAssetItemTypeRawFile           RelatedAssetItemType = "raw_file"
	RelatedAssetItemTypeStructuredContent RelatedAssetItemType = "structured_content"
	RelatedAssetItemTypeVideo             RelatedAssetItemType = "video"
)

// Defines values for SCContentTypeManagedMigrationResponseStatus.
const (
	SCContentTypeManagedMigrationResponseStatusError      SCContentTypeManagedMigrationResponseStatus = "error"
	SCContentTypeManagedMigrationResponseStatusInProgress SCContentTypeManagedMigrationResponseStatus = "in_progress"
	SCContentTypeManagedMigrationResponseStatusNotStarted SCContentTypeManagedMigrationResponseStatus = "not_started"
	SCContentTypeManagedMigrationResponseStatusSuccess    SCContentTypeManagedMigrationResponseStatus = "success"
)

// Defines values for SettingsRoutingRuleRulesUnit.
const (
	SettingsRoutingRuleRulesUnitDays        SettingsRoutingRuleRulesUnit = "days"
	SettingsRoutingRuleRulesUnitLessThanNil SettingsRoutingRuleRulesUnit = "<nil>"
	SettingsRoutingRuleRulesUnitMonths      SettingsRoutingRuleRulesUnit = "months"
	SettingsRoutingRuleRulesUnitWeeks       SettingsRoutingRuleRulesUnit = "weeks"
	SettingsRoutingRuleRulesUnitYears       SettingsRoutingRuleRulesUnit = "years"
)

// Defines values for TaskAssetDraftBrandComplianceRequestCategoriesStatus.
const (
	TaskAssetDraftBrandComplianceRequestCategoriesStatusCompliant     TaskAssetDraftBrandComplianceRequestCategoriesStatus = "compliant"
	TaskAssetDraftBrandComplianceRequestCategoriesStatusLessThanNil   TaskAssetDraftBrandComplianceRequestCategoriesStatus = "<nil>"
	TaskAssetDraftBrandComplianceRequestCategoriesStatusNotApplicable TaskAssetDraftBrandComplianceRequestCategoriesStatus = "not applicable"
	TaskAssetDraftBrandComplianceRequestCategoriesStatusNotCompliant  TaskAssetDraftBrandComplianceRequestCategoriesStatus = "not compliant"
)

// Defines values for TaskAssetDraftBrandComplianceRequestStatus.
const (
	TaskAssetDraftBrandComplianceRequestStatusApproved    TaskAssetDraftBrandComplianceRequestStatus = "approved"
	TaskAssetDraftBrandComplianceRequestStatusDeclined    TaskAssetDraftBrandComplianceRequestStatus = "declined"
	TaskAssetDraftBrandComplianceRequestStatusNotReviewed TaskAssetDraftBrandComplianceRequestStatus = "not_reviewed"
)

// Defines values for TaskAssetDraftBrandComplianceResponseCategoriesStatus.
const (
	Compliant     TaskAssetDraftBrandComplianceResponseCategoriesStatus = "compliant"
	LessThanNil   TaskAssetDraftBrandComplianceResponseCategoriesStatus = "<nil>"
	NotApplicable TaskAssetDraftBrandComplianceResponseCategoriesStatus = "not applicable"
	NotCompliant  TaskAssetDraftBrandComplianceResponseCategoriesStatus = "not compliant"
)

// Defines values for TaskAssetDraftBrandComplianceResponseStatus.
const (
	TaskAssetDraftBrandComplianceResponseStatusApproved    TaskAssetDraftBrandComplianceResponseStatus = "approved"
	TaskAssetDraftBrandComplianceResponseStatusDeclined    TaskAssetDraftBrandComplianceResponseStatus = "declined"
	TaskAssetDraftBrandComplianceResponseStatusNotReviewed TaskAssetDraftBrandComplianceResponseStatus = "not_reviewed"
)

// Defines values for TaskAssetDraftListResponseItemContentType.
const (
	TaskAssetDraftListResponseItemContentTypeURL TaskAssetDraftListResponseItemContentType = "url"
)

// Defines values for TaskAssetDraftListResponseItemType.
const (
	TaskAssetDraftListResponseItemTypeImage   TaskAssetDraftListResponseItemType = "image"
	TaskAssetDraftListResponseItemTypeRawFile TaskAssetDraftListResponseItemType = "raw_file"
	TaskAssetDraftListResponseItemTypeVideo   TaskAssetDraftListResponseItemType = "video"
)

// Defines values for TaskAssetDraftResponseContentType.
const (
	TaskAssetDraftResponseContentTypeURL TaskAssetDraftResponseContentType = "url"
)

// Defines values for TaskAssetDraftResponseType.
const (
	TaskAssetDraftResponseTypeImage   TaskAssetDraftResponseType = "image"
	TaskAssetDraftResponseTypeRawFile TaskAssetDraftResponseType = "raw_file"
	TaskAssetDraftResponseTypeVideo   TaskAssetDraftResponseType = "video"
)

// Defines values for TaskAssetRequestForLibraryContentType.
const (
	TaskAssetRequestForLibraryContentTypeArticle           TaskAssetRequestForLibraryContentType = "article"
	TaskAssetRequestForLibraryContentTypeImage             TaskAssetRequestForLibraryContentType = "image"
	TaskAssetRequestForLibraryContentTypeRawFile           TaskAssetRequestForLibraryContentType = "raw_file"
	TaskAssetRequestForLibraryContentTypeStructuredContent TaskAssetRequestForLibraryContentType = "structured_content"
	TaskAssetRequestForLibraryContentTypeVideo             TaskAssetRequestForLibraryContentType = "video"
)

// Defines values for TaskAssetResponseType.
const (
	TaskAssetResponseTypeArticle           TaskAssetResponseType = "article"
	TaskAssetResponseTypeImage             TaskAssetResponseType = "image"
	TaskAssetResponseTypeRawFile           TaskAssetResponseType = "raw_file"
	TaskAssetResponseTypeStructuredContent TaskAssetResponseType = "structured_content"
	TaskAssetResponseTypeVideo             TaskAssetResponseType = "video"
)

// Defines values for TaskBriefResponseType.
const (
	TaskBriefResponseTypeAttachment TaskBriefResponseType = "attachment"
	TaskBriefResponseTypeTemplate   TaskBriefResponseType = "template"
	TaskBriefResponseTypeText       TaskBriefResponseType = "text"
)

// Defines values for TaskCustomFieldType.
const (
	TaskCustomFieldTypeCheckboxes          TaskCustomFieldType = "checkboxes"
	TaskCustomFieldTypeDateField           TaskCustomFieldType = "date_field"
	TaskCustomFieldTypeDropdown            TaskCustomFieldType = "dropdown"
	TaskCustomFieldTypeImage               TaskCustomFieldType = "image"
	TaskCustomFieldTypeMultiLineTextField  TaskCustomFieldType = "multi_line_text_field"
	TaskCustomFieldTypeMultiSelectDropdown TaskCustomFieldType = "multi_select_dropdown"
	TaskCustomFieldTypeMultipleChoice      TaskCustomFieldType = "multiple_choice"
	TaskCustomFieldTypeRichTextField       TaskCustomFieldType = "rich_text_field"
	TaskCustomFieldTypeTextField           TaskCustomFieldType = "text_field"
	TaskCustomFieldTypeVideo               TaskCustomFieldType = "video"
)

// Defines values for TaskListResponseItemStatus.
const (
	TaskListResponseItemStatusArchived   TaskListResponseItemStatus = "Archived"
	TaskListResponseItemStatusCompleted  TaskListResponseItemStatus = "Completed"
	TaskListResponseItemStatusInProgress TaskListResponseItemStatus = "In Progress"
	TaskListResponseItemStatusNotStarted TaskListResponseItemStatus = "Not Started"
	TaskListResponseItemStatusOnHold     TaskListResponseItemStatus = "On Hold"
	TaskListResponseItemStatusOverdue    TaskListResponseItemStatus = "Overdue"
)

// Defines values for TaskResponseStatus.
const (
	TaskResponseStatusArchived   TaskResponseStatus = "Archived"
	TaskResponseStatusCompleted  TaskResponseStatus = "Completed"
	TaskResponseStatusInProgress TaskResponseStatus = "In Progress"
	TaskResponseStatusNotStarted TaskResponseStatus = "Not Started"
	TaskResponseStatusOnHold     TaskResponseStatus = "On Hold"
	TaskResponseStatusOverdue    TaskResponseStatus = "Overdue"
)

// Defines values for TaskSubStepAssigneeType.
const (
	TaskSubStepAssigneeTypeTeam TaskSubStepAssigneeType = "team"
	TaskSubStepAssigneeTypeUser TaskSubStepAssigneeType = "user"
)

// Defines values for TaskSubStepRequestAssigneeType.
const (
	TaskSubStepRequestAssigneeTypeTeam TaskSubStepRequestAssigneeType = "team"
	TaskSubStepRequestAssigneeTypeUser TaskSubStepRequestAssigneeType = "user"
)

// Defines values for TaskSubStepRequestIsCompleted.
const (
	TaskSubStepRequestIsCompletedTrue TaskSubStepRequestIsCompleted = true
)

// Defines values for TaskSubStepRequestIsInProgress.
const (
	TaskSubStepRequestIsInProgressTrue TaskSubStepRequestIsInProgress = true
)

// Defines values for TaskSubStepRequestIsSkipped.
const (
	True TaskSubStepRequestIsSkipped = true
)

// Defines values for TemplateChoiceFormFieldType.
const (
	TemplateChoiceFormFieldTypeBrief            TemplateChoiceFormFieldType = "brief"
	TemplateChoiceFormFieldTypeCheckbox         TemplateChoiceFormFieldType = "checkbox"
	TemplateChoiceFormFieldTypeCurrencyNumber   TemplateChoiceFormFieldType = "currency_number"
	TemplateChoiceFormFieldTypeDate             TemplateChoiceFormFieldType = "date"
	TemplateChoiceFormFieldTypeDropdown         TemplateChoiceFormFieldType = "dropdown"
	TemplateChoiceFormFieldTypeFile             TemplateChoiceFormFieldType = "file"
	TemplateChoiceFormFieldTypeInstruction      TemplateChoiceFormFieldType = "instruction"
	TemplateChoiceFormFieldTypeLabel            TemplateChoiceFormFieldType = "label"
	TemplateChoiceFormFieldTypePercentageNumber TemplateChoiceFormFieldType = "percentage_number"
	TemplateChoiceFormFieldTypeRadioButton      TemplateChoiceFormFieldType = "radio_button"
	TemplateChoiceFormFieldTypeRichtext         TemplateChoiceFormFieldType = "richtext"
	TemplateChoiceFormFieldTypeSection          TemplateChoiceFormFieldType = "section"
	TemplateChoiceFormFieldTypeSimpleNumber     TemplateChoiceFormFieldType = "simple_number"
	TemplateChoiceFormFieldTypeText             TemplateChoiceFormFieldType = "text"
	TemplateChoiceFormFieldTypeTextArea         TemplateChoiceFormFieldType = "text_area"
)

// Defines values for TemplateCurrencyNumberFormFieldType.
const (
	TemplateCurrencyNumberFormFieldTypeBrief            TemplateCurrencyNumberFormFieldType = "brief"
	TemplateCurrencyNumberFormFieldTypeCheckbox         TemplateCurrencyNumberFormFieldType = "checkbox"
	TemplateCurrencyNumberFormFieldTypeCurrencyNumber   TemplateCurrencyNumberFormFieldType = "currency_number"
	TemplateCurrencyNumberFormFieldTypeDate             TemplateCurrencyNumberFormFieldType = "date"
	TemplateCurrencyNumberFormFieldTypeDropdown         TemplateCurrencyNumberFormFieldType = "dropdown"
	TemplateCurrencyNumberFormFieldTypeFile             TemplateCurrencyNumberFormFieldType = "file"
	TemplateCurrencyNumberFormFieldTypeInstruction      TemplateCurrencyNumberFormFieldType = "instruction"
	TemplateCurrencyNumberFormFieldTypeLabel            TemplateCurrencyNumberFormFieldType = "label"
	TemplateCurrencyNumberFormFieldTypePercentageNumber TemplateCurrencyNumberFormFieldType = "percentage_number"
	TemplateCurrencyNumberFormFieldTypeRadioButton      TemplateCurrencyNumberFormFieldType = "radio_button"
	TemplateCurrencyNumberFormFieldTypeRichtext         TemplateCurrencyNumberFormFieldType = "richtext"
	TemplateCurrencyNumberFormFieldTypeSection          TemplateCurrencyNumberFormFieldType = "section"
	TemplateCurrencyNumberFormFieldTypeSimpleNumber     TemplateCurrencyNumberFormFieldType = "simple_number"
	TemplateCurrencyNumberFormFieldTypeText             TemplateCurrencyNumberFormFieldType = "text"
	TemplateCurrencyNumberFormFieldTypeTextArea         TemplateCurrencyNumberFormFieldType = "text_area"
)

// Defines values for TemplateDefaultFormFieldType.
const (
	TemplateDefaultFormFieldTypeBrief            TemplateDefaultFormFieldType = "brief"
	TemplateDefaultFormFieldTypeCheckbox         TemplateDefaultFormFieldType = "checkbox"
	TemplateDefaultFormFieldTypeCurrencyNumber   TemplateDefaultFormFieldType = "currency_number"
	TemplateDefaultFormFieldTypeDate             TemplateDefaultFormFieldType = "date"
	TemplateDefaultFormFieldTypeDropdown         TemplateDefaultFormFieldType = "dropdown"
	TemplateDefaultFormFieldTypeFile             TemplateDefaultFormFieldType = "file"
	TemplateDefaultFormFieldTypeInstruction      TemplateDefaultFormFieldType = "instruction"
	TemplateDefaultFormFieldTypeLabel            TemplateDefaultFormFieldType = "label"
	TemplateDefaultFormFieldTypePercentageNumber TemplateDefaultFormFieldType = "percentage_number"
	TemplateDefaultFormFieldTypeRadioButton      TemplateDefaultFormFieldType = "radio_button"
	TemplateDefaultFormFieldTypeRichtext         TemplateDefaultFormFieldType = "richtext"
	TemplateDefaultFormFieldTypeSection          TemplateDefaultFormFieldType = "section"
	TemplateDefaultFormFieldTypeSimpleNumber     TemplateDefaultFormFieldType = "simple_number"
	TemplateDefaultFormFieldTypeText             TemplateDefaultFormFieldType = "text"
	TemplateDefaultFormFieldTypeTextArea         TemplateDefaultFormFieldType = "text_area"
)

// Defines values for TemplateInstructionFormFieldType.
const (
	TemplateInstructionFormFieldTypeBrief            TemplateInstructionFormFieldType = "brief"
	TemplateInstructionFormFieldTypeCheckbox         TemplateInstructionFormFieldType = "checkbox"
	TemplateInstructionFormFieldTypeCurrencyNumber   TemplateInstructionFormFieldType = "currency_number"
	TemplateInstructionFormFieldTypeDate             TemplateInstructionFormFieldType = "date"
	TemplateInstructionFormFieldTypeDropdown         TemplateInstructionFormFieldType = "dropdown"
	TemplateInstructionFormFieldTypeFile             TemplateInstructionFormFieldType = "file"
	TemplateInstructionFormFieldTypeInstruction      TemplateInstructionFormFieldType = "instruction"
	TemplateInstructionFormFieldTypeLabel            TemplateInstructionFormFieldType = "label"
	TemplateInstructionFormFieldTypePercentageNumber TemplateInstructionFormFieldType = "percentage_number"
	TemplateInstructionFormFieldTypeRadioButton      TemplateInstructionFormFieldType = "radio_button"
	TemplateInstructionFormFieldTypeRichtext         TemplateInstructionFormFieldType = "richtext"
	TemplateInstructionFormFieldTypeSection          TemplateInstructionFormFieldType = "section"
	TemplateInstructionFormFieldTypeSimpleNumber     TemplateInstructionFormFieldType = "simple_number"
	TemplateInstructionFormFieldTypeText             TemplateInstructionFormFieldType = "text"
	TemplateInstructionFormFieldTypeTextArea         TemplateInstructionFormFieldType = "text_area"
)

// Defines values for TemplateListResponseDataApplicableTo.
const (
	TemplateListResponseDataApplicableToCampaignBrief TemplateListResponseDataApplicableTo = "campaign_brief"
	TemplateListResponseDataApplicableToTaskBrief     TemplateListResponseDataApplicableTo = "task_brief"
	TemplateListResponseDataApplicableToWorkRequest   TemplateListResponseDataApplicableTo = "work_request"
)

// Defines values for TemplateLogicRuleActionType.
const (
	JumpTo     TemplateLogicRuleActionType = "jump_to"
	ShowValues TemplateLogicRuleActionType = "show_values"
)

// Defines values for TemplateLogicRuleConditionOperator.
const (
	AnyOf    TemplateLogicRuleConditionOperator = "any_of"
	Equal    TemplateLogicRuleConditionOperator = "equal"
	NotEqual TemplateLogicRuleConditionOperator = "not_equal"
)

// Defines values for TemplatePercentageNumberFormFieldType.
const (
	TemplatePercentageNumberFormFieldTypeBrief            TemplatePercentageNumberFormFieldType = "brief"
	TemplatePercentageNumberFormFieldTypeCheckbox         TemplatePercentageNumberFormFieldType = "checkbox"
	TemplatePercentageNumberFormFieldTypeCurrencyNumber   TemplatePercentageNumberFormFieldType = "currency_number"
	TemplatePercentageNumberFormFieldTypeDate             TemplatePercentageNumberFormFieldType = "date"
	TemplatePercentageNumberFormFieldTypeDropdown         TemplatePercentageNumberFormFieldType = "dropdown"
	TemplatePercentageNumberFormFieldTypeFile             TemplatePercentageNumberFormFieldType = "file"
	TemplatePercentageNumberFormFieldTypeInstruction      TemplatePercentageNumberFormFieldType = "instruction"
	TemplatePercentageNumberFormFieldTypeLabel            TemplatePercentageNumberFormFieldType = "label"
	TemplatePercentageNumberFormFieldTypePercentageNumber TemplatePercentageNumberFormFieldType = "percentage_number"
	TemplatePercentageNumberFormFieldTypeRadioButton      TemplatePercentageNumberFormFieldType = "radio_button"
	TemplatePercentageNumberFormFieldTypeRichtext         TemplatePercentageNumberFormFieldType = "richtext"
	TemplatePercentageNumberFormFieldTypeSection          TemplatePercentageNumberFormFieldType = "section"
	TemplatePercentageNumberFormFieldTypeSimpleNumber     TemplatePercentageNumberFormFieldType = "simple_number"
	TemplatePercentageNumberFormFieldTypeText             TemplatePercentageNumberFormFieldType = "text"
	TemplatePercentageNumberFormFieldTypeTextArea         TemplatePercentageNumberFormFieldType = "text_area"
)

// Defines values for TemplateResponseApplicableTo.
const (
	CampaignBrief TemplateResponseApplicableTo = "campaign_brief"
	TaskBrief     TemplateResponseApplicableTo = "task_brief"
	WorkRequest   TemplateResponseApplicableTo = "work_request"
)

// Defines values for TemplateSimpleNumberFormFieldType.
const (
	TemplateSimpleNumberFormFieldTypeBrief            TemplateSimpleNumberFormFieldType = "brief"
	TemplateSimpleNumberFormFieldTypeCheckbox         TemplateSimpleNumberFormFieldType = "checkbox"
	TemplateSimpleNumberFormFieldTypeCurrencyNumber   TemplateSimpleNumberFormFieldType = "currency_number"
	TemplateSimpleNumberFormFieldTypeDate             TemplateSimpleNumberFormFieldType = "date"
	TemplateSimpleNumberFormFieldTypeDropdown         TemplateSimpleNumberFormFieldType = "dropdown"
	TemplateSimpleNumberFormFieldTypeFile             TemplateSimpleNumberFormFieldType = "file"
	TemplateSimpleNumberFormFieldTypeInstruction      TemplateSimpleNumberFormFieldType = "instruction"
	TemplateSimpleNumberFormFieldTypeLabel            TemplateSimpleNumberFormFieldType = "label"
	TemplateSimpleNumberFormFieldTypePercentageNumber TemplateSimpleNumberFormFieldType = "percentage_number"
	TemplateSimpleNumberFormFieldTypeRadioButton      TemplateSimpleNumberFormFieldType = "radio_button"
	TemplateSimpleNumberFormFieldTypeRichtext         TemplateSimpleNumberFormFieldType = "richtext"
	TemplateSimpleNumberFormFieldTypeSection          TemplateSimpleNumberFormFieldType = "section"
	TemplateSimpleNumberFormFieldTypeSimpleNumber     TemplateSimpleNumberFormFieldType = "simple_number"
	TemplateSimpleNumberFormFieldTypeText             TemplateSimpleNumberFormFieldType = "text"
	TemplateSimpleNumberFormFieldTypeTextArea         TemplateSimpleNumberFormFieldType = "text_area"
)

// Defines values for WorkRequestApprovedAssetResponseContentTypeType.
const (
	WorkRequestApprovedAssetResponseContentTypeTypeHTMLBody WorkRequestApprovedAssetResponseContentTypeType = "html_body"
	WorkRequestApprovedAssetResponseContentTypeTypeURL      WorkRequestApprovedAssetResponseContentTypeType = "url"
)

// Defines values for WorkRequestApprovedAssetResponseType.
const (
	WorkRequestApprovedAssetResponseTypeArticle WorkRequestApprovedAssetResponseType = "article"
	WorkRequestApprovedAssetResponseTypeImage   WorkRequestApprovedAssetResponseType = "image"
	WorkRequestApprovedAssetResponseTypeRawFile WorkRequestApprovedAssetResponseType = "raw_file"
	WorkRequestApprovedAssetResponseTypeVideo   WorkRequestApprovedAssetResponseType = "video"
)

// Defines values for WorkRequestAttachmentBriefRequestType.
const (
	WorkRequestAttachmentBriefRequestTypeAttachmentBrief WorkRequestAttachmentBriefRequestType = "attachment_brief"
)

// Defines values for WorkRequestAttachmentBriefResponseType.
const (
	WorkRequestAttachmentBriefResponseTypeAttachmentBrief WorkRequestAttachmentBriefResponseType = "attachment_brief"
)

// Defines values for WorkRequestRelatedResourceResponseRelationType.
const (
	Linked  WorkRequestRelatedResourceResponseRelationType = "linked"
	Started WorkRequestRelatedResourceResponseRelationType = "started"
)

// Defines values for WorkRequestRequestFormFieldBriefTypePayloadType.
const (
	WorkRequestRequestFormFieldBriefTypePayloadTypeBrief WorkRequestRequestFormFieldBriefTypePayloadType = "brief"
)

// Defines values for WorkRequestRequestFormFieldBriefTypeResponseType.
const (
	WorkRequestRequestFormFieldBriefTypeResponseTypeBrief WorkRequestRequestFormFieldBriefTypeResponseType = "brief"
)

// Defines values for WorkRequestRequestFormFieldCheckboxTypePayloadType.
const (
	WorkRequestRequestFormFieldCheckboxTypePayloadTypeCheckbox WorkRequestRequestFormFieldCheckboxTypePayloadType = "checkbox"
)

// Defines values for WorkRequestRequestFormFieldCheckboxTypeResponseType.
const (
	WorkRequestRequestFormFieldCheckboxTypeResponseTypeCheckbox WorkRequestRequestFormFieldCheckboxTypeResponseType = "checkbox"
)

// Defines values for WorkRequestRequestFormFieldCurrencyNumberTypePayloadType.
const (
	CurrencyNumber WorkRequestRequestFormFieldCurrencyNumberTypePayloadType = "currency_number"
)

// Defines values for WorkRequestRequestFormFieldDateTypePayloadType.
const (
	Date WorkRequestRequestFormFieldDateTypePayloadType = "date"
)

// Defines values for WorkRequestRequestFormFieldDropdownTypePayloadType.
const (
	WorkRequestRequestFormFieldDropdownTypePayloadTypeDropdown WorkRequestRequestFormFieldDropdownTypePayloadType = "dropdown"
)

// Defines values for WorkRequestRequestFormFieldDropdownTypeResponseType.
const (
	WorkRequestRequestFormFieldDropdownTypeResponseTypeDropdown WorkRequestRequestFormFieldDropdownTypeResponseType = "dropdown"
)

// Defines values for WorkRequestRequestFormFieldFileTypePayloadType.
const (
	WorkRequestRequestFormFieldFileTypePayloadTypeFile WorkRequestRequestFormFieldFileTypePayloadType = "file"
)

// Defines values for WorkRequestRequestFormFieldFileTypeResponseType.
const (
	WorkRequestRequestFormFieldFileTypeResponseTypeFile WorkRequestRequestFormFieldFileTypeResponseType = "file"
)

// Defines values for WorkRequestRequestFormFieldLabelTypePayloadType.
const (
	WorkRequestRequestFormFieldLabelTypePayloadTypeLabel WorkRequestRequestFormFieldLabelTypePayloadType = "label"
)

// Defines values for WorkRequestRequestFormFieldLabelTypeResponseType.
const (
	WorkRequestRequestFormFieldLabelTypeResponseTypeLabel WorkRequestRequestFormFieldLabelTypeResponseType = "label"
)

// Defines values for WorkRequestRequestFormFieldPercentageNumberTypePayloadType.
const (
	PercentageNumber WorkRequestRequestFormFieldPercentageNumberTypePayloadType = "percentage_number"
)

// Defines values for WorkRequestRequestFormFieldRadioButtonTypePayloadType.
const (
	WorkRequestRequestFormFieldRadioButtonTypePayloadTypeRadioButton WorkRequestRequestFormFieldRadioButtonTypePayloadType = "radio_button"
)

// Defines values for WorkRequestRequestFormFieldRadioButtonTypeResponseType.
const (
	WorkRequestRequestFormFieldRadioButtonTypeResponseTypeRadioButton WorkRequestRequestFormFieldRadioButtonTypeResponseType = "radio_button"
)

// Defines values for WorkRequestRequestFormFieldRichtextTypePayloadType.
const (
	Richtext WorkRequestRequestFormFieldRichtextTypePayloadType = "richtext"
)

// Defines values for WorkRequestRequestFormFieldSimpleNumberTypePayloadType.
const (
	SimpleNumber WorkRequestRequestFormFieldSimpleNumberTypePayloadType = "simple_number"
)

// Defines values for WorkRequestRequestFormFieldTextAreaTypePayloadType.
const (
	TextArea WorkRequestRequestFormFieldTextAreaTypePayloadType = "text_area"
)

// Defines values for WorkRequestRequestFormFieldTextTypePayloadType.
const (
	Text WorkRequestRequestFormFieldTextTypePayloadType = "text"
)

// Defines values for WorkRequestResponseAssigneesType.
const (
	WorkRequestResponseAssigneesTypeTeam WorkRequestResponseAssigneesType = "team"
	WorkRequestResponseAssigneesTypeUser WorkRequestResponseAssigneesType = "user"
)

// Defines values for WorkRequestResponsePriority.
const (
	WorkRequestResponsePriorityHigh   WorkRequestResponsePriority = "High"
	WorkRequestResponsePriorityLow    WorkRequestResponsePriority = "Low"
	WorkRequestResponsePriorityMedium WorkRequestResponsePriority = "Medium"
)

// Defines values for WorkRequestResponseStatus.
const (
	WorkRequestResponseStatusAccepted  WorkRequestResponseStatus = "Accepted"
	WorkRequestResponseStatusCompleted WorkRequestResponseStatus = "Completed"
	WorkRequestResponseStatusDeclined  WorkRequestResponseStatus = "Declined"
	WorkRequestResponseStatusSubmitted WorkRequestResponseStatus = "Submitted"
)

// Defines values for WorkRequestTextBriefRequestType.
const (
	WorkRequestTextBriefRequestTypeTextBrief WorkRequestTextBriefRequestType = "text_brief"
)

// Defines values for WorkRequestTextBriefResponseType.
const (
	WorkRequestTextBriefResponseTypeTextBrief WorkRequestTextBriefResponseType = "text_brief"
)

// Defines values for WorkRequestUpdateRequestPriority.
const (
	WorkRequestUpdateRequestPriorityHigh   WorkRequestUpdateRequestPriority = "High"
	WorkRequestUpdateRequestPriorityLow    WorkRequestUpdateRequestPriority = "Low"
	WorkRequestUpdateRequestPriorityMedium WorkRequestUpdateRequestPriority = "Medium"
)

// Defines values for WorkRequestUpdateRequestStatus.
const (
	Accepted  WorkRequestUpdateRequestStatus = "Accepted"
	Completed WorkRequestUpdateRequestStatus = "Completed"
	Declined  WorkRequestUpdateRequestStatus = "Declined"
	Submitted WorkRequestUpdateRequestStatus = "Submitted"
)

// Defines values for WorkflowResponseStepsSubStepsAssigneesType.
const (
	WorkflowResponseStepsSubStepsAssigneesTypeTeam WorkflowResponseStepsSubStepsAssigneesType = "team"
	WorkflowResponseStepsSubStepsAssigneesTypeUser WorkflowResponseStepsSubStepsAssigneesType = "user"
)

// Defines values for WorkflowResponseStepsSubStepsType.
const (
	WorkflowResponseStepsSubStepsTypeDefault  WorkflowResponseStepsSubStepsType = "default"
	WorkflowResponseStepsSubStepsTypeExternal WorkflowResponseStepsSubStepsType = "external"
)

// Defines values for ErrorReason.
const (
	INVALIDSELFREF      ErrorReason = "INVALID_SELF_REF"
	INVALIDSTATE        ErrorReason = "INVALID_STATE"
	LISTEMPTY           ErrorReason = "LIST_EMPTY"
	MAXNOTMET           ErrorReason = "MAX_NOT_MET"
	MINNOTMET           ErrorReason = "MIN_NOT_MET"
	PATTERNERROR        ErrorReason = "PATTERN_ERROR"
	REQUIREDFIELDABSENT ErrorReason = "REQUIRED_FIELD_ABSENT"
)

// AcknowledgeSCContentPreviewResponse defines model for AcknowledgeSCContentPreviewResponse.
type AcknowledgeSCContentPreviewResponse = interface{}

// AddURLToTaskRequest defines model for AddUrlToTaskRequest.
type AddURLToTaskRequest struct {
	// Title Title of the URL.
	Title *string `json:"title,omitempty"`

	// URL The URL to be added to the task.
	URL string `json:"url"`
}

// AllowedContentTypeItem defines model for AllowedContentTypeItem.
type AllowedContentTypeItem struct {
	// Name Name of the item
	Name string `json:"name"`

	// URL URL of the item
	URL string `json:"url"`
}

// ArticleAuthorResponse Author of the article
type ArticleAuthorResponse struct {
	// Name Name of Author
	Name *string `json:"name"`
}

// AssetContent Content of the asset
type AssetContent struct {
	// Type Type of the content.
	//  - article  – `html_body`
	//  - image, raw_file, video  – `url`
	//  - structured_content  – `api_url`
	Type AssetContentType `json:"type"`

	// Value Content of the asset.  - article – the html body - image, raw_file, video – the download URL - structured_content – api url
	Value string `json:"value"`
}

// AssetContentType Type of the content.
//   - article  – `html_body`
//   - image, raw_file, video  – `url`
//   - structured_content  – `api_url`
type AssetContentType string

// AssetFieldListResponseItem defines model for AssetFieldListResponseItem.
type AssetFieldListResponseItem struct {
	union json.RawMessage
}

// AssetFieldTypeCheckbox defines model for AssetFieldTypeCheckbox.
type AssetFieldTypeCheckbox struct {
	// Choices Choices of the checkbox
	Choices []AssetFieldTypeCheckboxChoices `json:"choices"`

	// ID Identifier of the field
	ID string `json:"id"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// AssetFieldTypeCheckboxChoices defines model for .
type AssetFieldTypeCheckboxChoices struct {
	// ID Identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// AssetFieldTypeCommon defines model for AssetFieldTypeCommon.
type AssetFieldTypeCommon struct {
	// ID Identifier of the field
	ID string `json:"id"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// AssetFieldTypeCurrencyNumber defines model for AssetFieldTypeCurrencyNumber.
type AssetFieldTypeCurrencyNumber struct {
	// CurrencyCode Currency code of the numerical field
	CurrencyCode string `json:"currency_code"`

	// DecimalPlaces Decimal place of the numerical field
	DecimalPlaces *int `json:"decimal_places"`

	// HasThousandSeparator Whether the numerical field has a thousand separator
	HasThousandSeparator bool `json:"has_thousand_separator"`

	// ID Identifier of the field
	ID string `json:"id"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// AssetFieldTypeDropdown defines model for AssetFieldTypeDropdown.
type AssetFieldTypeDropdown struct {
	// Choices Choices of the dropdown
	Choices []AssetFieldTypeDropdownChoices `json:"choices"`

	// ID Identifier of the field
	ID string `json:"id"`

	// IsMultiSelect Allow users to select multiple values from the dropdown
	IsMultiSelect bool `json:"is_multi_select"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// AssetFieldTypeDropdownChoices defines model for .
type AssetFieldTypeDropdownChoices struct {
	// ID Identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// AssetFieldTypeLabel defines model for AssetFieldTypeLabel.
type AssetFieldTypeLabel struct {
	// Choices Choices of the label
	Choices []AssetFieldTypeLabelChoices `json:"choices"`

	// ID Identifier of the field
	ID string `json:"id"`

	// IsMultiSelect Select multiple values from the label
	IsMultiSelect bool `json:"is_multi_select"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// AssetFieldTypeLabelChoices defines model for .
type AssetFieldTypeLabelChoices struct {
	// ID Identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// AssetFieldTypePercentageNumber defines model for AssetFieldTypePercentageNumber.
type AssetFieldTypePercentageNumber struct {
	// DecimalPlaces Decimal place of the numerical field
	DecimalPlaces *int `json:"decimal_places"`

	// ID Identifier of the field
	ID string `json:"id"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// AssetFieldTypeRadio defines model for AssetFieldTypeRadio.
type AssetFieldTypeRadio struct {
	// Choices Choices of the radio button
	Choices []AssetFieldTypeRadioChoices `json:"choices"`

	// ID Identifier of the field
	ID string `json:"id"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// AssetFieldTypeRadioChoices defines model for .
type AssetFieldTypeRadioChoices struct {
	// ID Identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// AssetFieldTypeSimpleNumber defines model for AssetFieldTypeSimpleNumber.
type AssetFieldTypeSimpleNumber struct {
	// DecimalPlaces Decimal place of the numerical field
	DecimalPlaces *int `json:"decimal_places"`

	// HasThousandSeparator Whether the numerical field has a thousand separator
	HasThousandSeparator bool `json:"has_thousand_separator"`

	// ID Identifier of the field
	ID string `json:"id"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// AssetFieldTypes Type of the field
type AssetFieldTypes string

// AssetFieldUpdateRequest defines model for AssetFieldUpdateRequest.
type AssetFieldUpdateRequest struct {
	union json.RawMessage
}

// AssetFieldUpdateResponse defines model for AssetFieldUpdateResponse.
type AssetFieldUpdateResponse struct {
	// ID Unique identifier for the field
	ID string `json:"id"`

	// Links Meta links
	Links AssetFieldUpdateResponseLinks `json:"links"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type AssetFieldTypes `json:"type"`

	// Values List of values selected or from user input
	Values []string `json:"values"`
}

// AssetFieldUpdateResponseLinks Meta links
type AssetFieldUpdateResponseLinks struct {
	// Asset URL of the asset
	Asset string `json:"asset"`

	// AssetFields URL of the asset fields
	AssetFields string `json:"asset_fields"`
}

// AssetFieldsUpdateRequest defines model for AssetFieldsUpdateRequest.
type AssetFieldsUpdateRequest = []AssetFieldsUpdateRequest_Item

// AssetFieldsUpdateRequest_Item defines model for AssetFieldsUpdateRequest.Item.
type AssetFieldsUpdateRequest_Item struct {
	union json.RawMessage
}

// AssetLineageCreateRequest defines model for AssetLineageCreateRequest.
type AssetLineageCreateRequest struct {
	// IconURL Website favicon url
	IconURL *string `json:"icon_url,omitempty"`

	// Name Name of the source
	Name string `json:"name"`

	// RenditionID Unique identifier of a rendition of an existing version of the asset
	RenditionID *string `json:"rendition_id,omitempty"`

	// URI URI of the resource where the asset is used. If this is an external url, make sure the url starts with http:// or https://.
	URI string `json:"uri"`
}

// AssetLineageResponse defines model for AssetLineageResponse.
type AssetLineageResponse struct {
	// AssetID Unique identifier of the asset
	AssetID string `json:"asset_id"`

	// CreatedAt Date and time on which the asset lineage was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// IconURL Icon url of the asset lineage
	IconURL *string `json:"icon_url"`

	// ID Unique identifier of the asset lineage
	ID string `json:"id"`

	// Links Meta links
	Links AssetLineageResponseLinks `json:"links"`

	// Name Name of the source
	Name string `json:"name"`

	// RenditionID Unique identifier of a rendition of the asset
	RenditionID *string `json:"rendition_id"`

	// URI Unique identifier of the location where the asset is being used
	URI string `json:"uri"`

	// UsedIn Indicates where the asset is being used
	UsedIn AssetLineageResponseUsedIn `json:"used_in"`

	// VersionID Unique identifier of the asset version
	VersionID string `json:"version_id"`
}

// AssetLineageResponseLinks Meta links
type AssetLineageResponseLinks struct {
	// Asset URL of the asset
	Asset string `json:"asset"`

	// Self URL of the asset lineage
	Self string `json:"self"`
}

// AssetLineageResponseUsedIn Indicates where the asset is being used
type AssetLineageResponseUsedIn string

// AssetPermissionBulkCreateRequest defines model for AssetPermissionBulkCreateRequest.
type AssetPermissionBulkCreateRequest struct {
	// Permissions List of permissions to be granted to asset
	Permissions []AssetPermissionBulkCreateRequestPermissions `json:"permissions"`

	// Type Type of the accessor
	Type AssetPermissionBulkCreateRequestType `json:"type"`
}

// AssetPermissionBulkCreateRequestPermissionsAccessType level of access the accessor has to the asset
type AssetPermissionBulkCreateRequestPermissionsAccessType string

// AssetPermissionBulkCreateRequestPermissions defines model for .
type AssetPermissionBulkCreateRequestPermissions struct {
	// AccessType level of access the accessor has to the asset
	AccessType AssetPermissionBulkCreateRequestPermissionsAccessType `json:"access_type"`

	// ID Unique identifier of accessor
	ID string `json:"id"`

	// IsOwner Indicates if the accessor is the owner of the asset. A team cannot be an owner.
	IsOwner *bool `json:"is_owner,omitempty"`
}

// AssetPermissionBulkCreateRequestType Type of the accessor
type AssetPermissionBulkCreateRequestType string

// AssetPermissionListResponseItem defines model for AssetPermissionListResponseItem.
type AssetPermissionListResponseItem struct {
	// Data List of Permissions
	Data       []BasePermissionsResponseSchema `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// AssetPermissionUpdateRequest defines model for AssetPermissionUpdateRequest.
type AssetPermissionUpdateRequest struct {
	// AccessType Set the level of access the accessor has to the asset
	AccessType *AssetPermissionUpdateRequestAccessType `json:"access_type,omitempty"`

	// IsOwner Used to set an accessor as owner. This field can only be set to true as an owner cannot be removed without another accessor being set as the new owner. A team cannot be an owner.
	IsOwner *bool `json:"is_owner,omitempty"`
}

// AssetPermissionUpdateRequestAccessType Set the level of access the accessor has to the asset
type AssetPermissionUpdateRequestAccessType string

// AssetResponse defines model for AssetResponse.
type AssetResponse struct {
	// Content Content of the asset
	Content AssetContent `json:"content"`

	// CreatedAt Date and time on which the asset was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// FileExtension File extension of the asset
	FileExtension *string `json:"file_extension"`

	// FileLocation Location of the folder containing the asset in the library
	FileLocation string `json:"file_location"`

	// FolderID ID of the folder containing the asset in the library
	FolderID *string `json:"folder_id"`

	// ID Unique identifier for the asset
	ID string `json:"id"`

	// IsArchived Whether the asset is archived or not
	IsArchived bool `json:"is_archived"`

	// Labels Labels associated with the asset
	Labels []ResourceLabelResponse `json:"labels"`

	// Links Meta links
	Links AssetResponseLinks `json:"links"`

	// MimeType MIME type of the asset
	MimeType string `json:"mime_type"`

	// ModifiedAt Date and time of the most recent modification of the asset, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// OwnerOrganizationID ID of the asset's owner organization
	OwnerOrganizationID string `json:"owner_organization_id"`

	// ThumbnailURL URL of the asset's thumbnail
	ThumbnailURL *string `json:"thumbnail_url"`

	// Title Title of the asset
	Title string `json:"title"`

	// Type Type of the asset
	Type AssetResponseType `json:"type"`
}

// AssetResponseLinks Meta links
type AssetResponseLinks struct {
	// Self URL of the asset (only GET is supported)
	Self string `json:"self"`
}

// AssetResponseType Type of the asset
type AssetResponseType string

// AssetTypeObjectField defines model for AssetTypeObjectField.
type AssetTypeObjectField struct {
	// ID Unique identifier of the field
	ID string `json:"id"`

	// Type Type of the field
	Type string `json:"type"`

	// Values Payload for fields with attachments
	Values []AssetTypeObjectFieldUpdatePayloadValues `json:"values"`
}

// AssetTypeObjectFieldUpdatePayloadValues defines model for .
type AssetTypeObjectFieldUpdatePayloadValues struct {
	// Key Unique identifier of the file upload session. This is the `upload_meta_fields.key` field retrieved from the `/v3/upload-url` endpoint.
	Key string `json:"key"`

	// Title Title of the image or video
	Title string `json:"title"`
}

// AssetTypeObjectFieldUpdatePayload defines model for AssetTypeObjectFieldUpdatePayload.
type AssetTypeObjectFieldUpdatePayload struct {
	// Type Type of the field
	Type string `json:"type"`

	// Values Payload for fields with attachments
	Values []AssetTypeObjectFieldUpdatePayloadValues `json:"values"`
}

// AssetVersionCreateRequest defines model for AssetVersionCreateRequest.
type AssetVersionCreateRequest struct {
	// Key Unique identifier of the file upload session. This is the `upload_meta_fields.key` field retrieved from the `/v3/upload-url` endpoint.
	Key string `json:"key"`

	// Title Title of the asset
	Title string `json:"title"`
}

// AttachmentRequest defines model for AttachmentRequest.
type AttachmentRequest struct {
	// Key Unique identifier of the file upload session. This is the `upload_meta_fields.key` field retrieved from the `/v3/upload-url` endpoint.
	Key string `json:"key"`

	// Name Filename of the attachment
	Name string `json:"name"`
}

// AttachmentResponse defines model for AttachmentResponse.
type AttachmentResponse struct {
	// ID Unique identifier for the attachment
	ID string `json:"id"`

	// Name Filename of the attachment
	Name string `json:"name"`

	// URL Download URL of the attachment
	URL string `json:"url"`
}

// BaseAssetField defines model for BaseAssetField.
type BaseAssetField struct {
	// ID Identifier of the field
	ID string `json:"id"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// BaseAssetRenditionResponse defines model for BaseAssetRenditionResponse.
type BaseAssetRenditionResponse struct {
	// CreatedAt Date and time on which the rendition was created, in ISO 8601 UTC format
	CreatedAt time.Time `json:"created_at"`

	// ID Unique identifier of the rendition
	ID string `json:"id"`

	// MimeType MIME type of the rendition
	MimeType string `json:"mime_type"`

	// Name Name of the rendition
	Name string `json:"name"`

	// URL URL of the rendition
	URL string `json:"url"`
}

// BaseContentTypeModel defines model for BaseContentTypeModel.
type BaseContentTypeModel struct {
	// Component Status determining whether the content type is a component
	Component bool `json:"component"`

	// ContentTypeGUID Unique identifier of the content type
	ContentTypeGUID string `json:"content_type_guid"`

	// CreatedAt Date and time on which the content type was created, in ISO 8601 UTC format
	CreatedAt time.Time `json:"created_at"`

	// CreatedBy Unique identifier of the user who created the content type
	CreatedBy string `json:"created_by"`

	// Description Description of the content type
	Description *string `json:"description,omitempty"`

	// Disabled Disabled status of the content type
	Disabled *bool `json:"disabled,omitempty"`

	// Links Meta Links
	Links BaseContentTypeModelLinks `json:"links"`

	// Name Name of the content type
	Name string `json:"name"`

	// Source Source of the content type
	Source *string `json:"source,omitempty"`

	// SourceID Source of the content type
	SourceID *string `json:"source_id,omitempty"`

	// SourceMetadata Source metadata of the content type
	SourceMetadata *string `json:"source_metadata,omitempty"`

	// ThumbnailGUID Thumbnail GUID of the content type
	ThumbnailGUID *string `json:"thumbnail_guid,omitempty"`

	// UpdatedAt Date and time on which the content type was last updated, in ISO 8601 UTC format
	UpdatedAt time.Time `json:"updated_at"`

	// UpdatedBy Unique identifier of the user who last updated the content type
	UpdatedBy string `json:"updated_by"`
}

// BaseContentTypeModelLinks Meta Links
type BaseContentTypeModelLinks struct {
	// Self URL of the content type
	Self *string `json:"self,omitempty"`

	// Versions URL of the content type versions
	Versions *string `json:"versions,omitempty"`
}

// BaseContentTypeVersionModel defines model for BaseContentTypeVersionModel.
type BaseContentTypeVersionModel struct {
	// CreatedAt Date and time on which the version was created, in ISO 8601 UTC format
	CreatedAt time.Time `json:"created_at"`

	// CreatedBy Unique identifier of the user who created the version
	CreatedBy string `json:"created_by"`

	// ExpectedLocales Expected locales of the version
	ExpectedLocales []string `json:"expected_locales"`

	// Latest Indicates whether the version is the latest
	Latest bool `json:"latest"`

	// Links Meta links
	Links BaseContentTypeVersionModelLinks `json:"links"`

	// VersionGUID Unique identifier of the version
	VersionGUID string `json:"version_guid"`
}

// BaseContentTypeVersionModelLinks Meta links
type BaseContentTypeVersionModelLinks struct {
	// ContentType URL of the content type
	ContentType *string `json:"content_type,omitempty"`

	// Self URL of the content type version
	Self *string `json:"self,omitempty"`
}

// BaseFieldDefinition defines model for BaseFieldDefinition.
type BaseFieldDefinition struct {
	// BaseType An enumeration.
	BaseType      BaseFieldDefinitionType                   `json:"base_type"`
	Core          CoreFieldDef                              `json:"core"`
	DefaultValues *[]BaseFieldDefinition_DefaultValues_Item `json:"default_values,omitempty"`
}

// BaseFieldDefinitionDefaultValues1 defines model for .
type BaseFieldDefinitionDefaultValues1 = bool

// BaseFieldDefinitionDefaultValues2 defines model for .
type BaseFieldDefinitionDefaultValues2 = string

// BaseFieldDefinitionDefaultValues3 defines model for .
type BaseFieldDefinitionDefaultValues3 = map[string]interface{}

// BaseFieldDefinitionDefaultValues4 defines model for .
type BaseFieldDefinitionDefaultValues4 = []interface{}

// BaseFieldDefinition_DefaultValues_Item defines model for BaseFieldDefinition.default_values.Item.
type BaseFieldDefinition_DefaultValues_Item struct {
	union json.RawMessage
}

// BaseFieldDefinitionType An enumeration.
type BaseFieldDefinitionType string

// BaseFormFieldRequest defines model for BaseFormFieldRequest.
type BaseFormFieldRequest struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// Type Type of the form field
	Type string `json:"type"`

	// Values Values for the form field to populate in the work request
	Values []interface{} `json:"values"`
}

// BaseFormFieldResponse defines model for BaseFormFieldResponse.
type BaseFormFieldResponse struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// IsProtected True when this form field is read-only and only organization admins can modify it. Absent or false for editable fields.
	IsProtected *bool `json:"is_protected,omitempty"`

	// Type Type of the form field
	Type string `json:"type"`

	// Values Values of the form field
	Values []interface{} `json:"values"`
}

// BaseObjectField defines model for BaseObjectField.
type BaseObjectField struct {
	// ID Unique identifier of the field
	ID string `json:"id"`

	// Type Type of the field
	Type string `json:"type"`

	// Values Values for the campaign field
	Values []interface{} `json:"values"`
}

// BaseObjectFieldUpdatePayload defines model for BaseObjectFieldUpdatePayload.
type BaseObjectFieldUpdatePayload struct {
	// Type Type of the field
	Type string `json:"type"`

	// Values Values for the campaign field
	Values []interface{} `json:"values"`
}

// BasePermissionsResponseSchema defines model for BasePermissionsResponseSchema.
type BasePermissionsResponseSchema struct {
	// AccessType level of access the accessor has to the resource
	AccessType BasePermissionsResponseSchemaAccessType `json:"access_type"`

	// ID Unique identifier of the accessor
	ID string `json:"id"`

	// IsOwner Indicates if the accessor is the owner of the object
	IsOwner bool                               `json:"is_owner"`
	Links   BasePermissionsResponseSchemaLinks `json:"links"`

	// Name Name of the accessor
	Name string `json:"name"`

	// Type Type of the accessor
	Type BasePermissionsResponseSchemaType `json:"type"`
}

// BasePermissionsResponseSchemaAccessType level of access the accessor has to the resource
type BasePermissionsResponseSchemaAccessType string

// BasePermissionsResponseSchemaLinks defines model for .
type BasePermissionsResponseSchemaLinks struct {
	// Accessor links for user and team type objects
	Accessor *string `json:"accessor"`
}

// BasePermissionsResponseSchemaType Type of the accessor
type BasePermissionsResponseSchemaType string

// BaseSettingsFieldCreatePayload defines model for BaseSettingsFieldCreatePayload.
type BaseSettingsFieldCreatePayload struct {
	// HelperText Used to help understand the functionality of the field.
	HelperText *string `json:"helper_text,omitempty"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive bool `json:"is_active"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`
}

// BaseSettingsFieldUpdatePayload defines model for BaseSettingsFieldUpdatePayload.
type BaseSettingsFieldUpdatePayload struct {
	// HelperText Used to help understand the functionality of the field.
	HelperText *string `json:"helper_text,omitempty"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive *bool `json:"is_active,omitempty"`

	// Name Name of the field
	Name *string `json:"name,omitempty"`

	// Type Type of the field. This field is required but not updateable. Type of a field will always stay the same. Update request body must contain atleast one other updateable field.
	Type string `json:"type"`
}

// BaseSettingsFieldsResponse defines model for BaseSettingsFieldsResponse.
type BaseSettingsFieldsResponse struct {
	// HelperText Used to help understand functionality of field
	HelperText string `json:"helper_text"`

	// ID Identifier of the field
	ID string `json:"id"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive bool `json:"is_active"`

	// Links Meta links
	Links BaseSettingsFieldsResponseLinks `json:"links"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`
}

// BaseSettingsFieldsResponseLinks Meta links
type BaseSettingsFieldsResponseLinks struct {
	// Source URL of the list of fields filtered only with the field
	Source string `json:"source"`
}

// BaseTemplateFormField defines model for BaseTemplateFormField.
type BaseTemplateFormField struct {
	// HelperText Helper text of the field
	HelperText *string `json:"helper_text"`

	// Identifier Identifier of the field
	Identifier string `json:"identifier"`

	// IsProtected True when this form field is read-only and only organization admins can modify it. Absent or false for editable fields.
	IsProtected *bool `json:"is_protected,omitempty"`

	// IsReadonly Specify if the field can be used to create any resource
	IsReadonly bool `json:"is_readonly"`

	// IsRequired Required status of the field
	IsRequired bool `json:"is_required"`

	// Label Label of the field
	Label *string `json:"label"`

	// SortOrder Sort order of the field
	SortOrder int `json:"sort_order"`

	// Type Type of the field
	Type BaseTemplateFormFieldType `json:"type"`
}

// BaseTemplateFormFieldType Type of the field
type BaseTemplateFormFieldType string

// BaseTemplateResponse defines model for BaseTemplateResponse.
type BaseTemplateResponse struct {
	// ApplicableTo List of resources the template is applicable to
	ApplicableTo []BaseTemplateResponseApplicableTo `json:"applicable_to"`

	// Description Description of the template
	Description string `json:"description"`

	// ID Unique identifier of the template
	ID string `json:"id"`

	// IsActive Active status of the template
	IsActive bool `json:"is_active"`

	// Title Title of the template
	Title string `json:"title"`
}

// BaseTemplateResponseApplicableTo defines model for BaseTemplateResponse.ApplicableTo.
type BaseTemplateResponseApplicableTo string

// BatchFileURLResponse defines model for BatchFileUrlResponse.
type BatchFileURLResponse struct {
	// Urls List of urls for Files.
	Urls map[string]*string `json:"urls"`
}

// BooleanFieldValueModel defines model for BooleanFieldValueModel.
type BooleanFieldValueModel struct {
	// BoolValue Value of the field value
	BoolValue bool `json:"bool_value"`

	// OrderIndex Order index of the field value
	OrderIndex *int `json:"order_index,omitempty"`
}

// BrandComplianceCategoriesResponse defines model for BrandComplianceCategoriesResponse.
type BrandComplianceCategoriesResponse struct {
	Criteria []BrandComplianceCategoriesResponseCriteria `json:"criteria"`

	// ID Unique identifier of the category
	ID string `json:"id"`

	// Name Name of the category
	Name *string `json:"name,omitempty"`
}

// BrandComplianceCategoriesResponseCriteria defines model for .
type BrandComplianceCategoriesResponseCriteria struct {
	// Description Details about the criteria
	Description *string `json:"description,omitempty"`

	// ID Unique identifier of the criteria
	ID string `json:"id"`

	// Name Name of the criteria
	Name *string `json:"name,omitempty"`
}

// BriefTypeFormFieldRequest defines model for BriefTypeFormFieldRequest.
type BriefTypeFormFieldRequest struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// Type Type of the form field
	Type string `json:"type"`

	// Values Array of brief values
	Values []BriefTypeFormFieldRequest_Values_Item `json:"values"`
}

// BriefTypeFormFieldRequest_Values_Item defines model for BriefTypeFormFieldRequest.values.Item.
type BriefTypeFormFieldRequest_Values_Item struct {
	union json.RawMessage
}

// BriefTypeFormFieldResponse defines model for BriefTypeFormFieldResponse.
type BriefTypeFormFieldResponse struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// IsProtected True when this form field is read-only and only organization admins can modify it. Absent or false for editable fields.
	IsProtected *bool `json:"is_protected,omitempty"`

	// Type Type of the form field
	Type string `json:"type"`

	// Values Array of brief values
	Values []BriefTypeFormFieldResponse_Values_Item `json:"values"`
}

// BriefTypeFormFieldResponse_Values_Item defines model for BriefTypeFormFieldResponse.values.Item.
type BriefTypeFormFieldResponse_Values_Item struct {
	union json.RawMessage
}

// BudgetResponse Budget of the campaign
type BudgetResponse struct {
	// BudgetedAmount Planned amount for the campaign
	BudgetedAmount string `json:"budgeted_amount"`

	// CurrencyCode Currency code of the budget
	CurrencyCode string `json:"currency_code"`
}

// CampaignBriefResponse defines model for CampaignBriefResponse.
type CampaignBriefResponse struct {
	// Fields List of fields of the brief
	Fields []CampaignBriefResponseFields `json:"fields"`

	// Links Meta links
	Links    CampaignBriefResponseLinks  `json:"links"`
	Template *CampaignBriefTemplateValue `json:"template"`

	// Title Title of the campaign brief
	Title string `json:"title"`

	// Type Type of the campaign brief
	Type CampaignBriefResponseType `json:"type"`
}

// CampaignBriefResponseFields defines model for .
type CampaignBriefResponseFields struct {
	// Name Name of the brief field
	Name string `json:"name"`

	// Value List of values set for the brief field
	Value []string `json:"value"`
}

// CampaignBriefResponseLinks Meta links
type CampaignBriefResponseLinks struct {
	// Campaign URL of the campaign
	Campaign string `json:"campaign"`

	// Self URL of the brief
	Self string `json:"self"`
}

// CampaignBriefResponseType Type of the campaign brief
type CampaignBriefResponseType string

// CampaignBriefTemplateValue Template info of the brief if the brief is of template type
type CampaignBriefTemplateValue struct {
	// ID Unique identifier of the template
	ID string `json:"id"`

	// Name Name of the template
	Name string `json:"name"`
}

// CampaignCommentResponse defines model for CampaignCommentResponse.
type CampaignCommentResponse struct {
	// Attachments List of attachments of the comment
	Attachments []AttachmentResponse `json:"attachments"`

	// CommentBy Unique identifier of the user who posted the comment
	CommentBy string `json:"comment_by"`

	// CreatedAt Creation date and time of the campaign comment in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// ID Unique identifier of the campaign comment
	ID string `json:"id"`

	// IsResolved Determines if the comment is resolved in Optimizely CMP
	IsResolved bool `json:"is_resolved"`

	// Links Meta links
	Links CampaignCommentResponseLinks `json:"links"`

	// ModifiedAt Last modification date and time of the campaign comment in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// ParentCommentID Unique identifier of the parent comment
	ParentCommentID *string `json:"parent_comment_id"`

	// Value Content of the comment
	Value string `json:"value"`
}

// CampaignCommentResponseLinks Meta links
type CampaignCommentResponseLinks struct {
	// Campaign URL of the campaign
	Campaign string `json:"campaign"`

	// CommentBy URL of the user who posted the comment
	CommentBy string `json:"comment_by"`
}

// CampaignCreateRequest defines model for CampaignCreateRequest.
type CampaignCreateRequest struct {
	// Color Color of the campaign
	Color *CampaignCreateRequestColor `json:"color,omitempty"`

	// Description Description of the campaign
	Description *string `json:"description"`

	// EndDate End date of the campaign, in ISO 8601 UTC format (2026-02-10).
	EndDate *openapi_types.Date `json:"end_date,omitempty"`

	// OwnerID ID of the owner of the campaign
	OwnerID *string `json:"owner_id,omitempty"`

	// ParentCampaign ID of the parent campaign
	ParentCampaign *string `json:"parent_campaign"`

	// StartDate Start date of the campaign, in ISO 8601 UTC format (2025-02-10).
	StartDate *openapi_types.Date `json:"start_date,omitempty"`

	// Title Title of the campaign
	Title string `json:"title"`
}

// CampaignCreateRequestColor Color of the campaign
type CampaignCreateRequestColor string

// CampaignFieldTypes Type of the field
type CampaignFieldTypes string

// CampaignFieldUpdateRequest defines model for CampaignFieldUpdateRequest.
type CampaignFieldUpdateRequest struct {
	union json.RawMessage
}

// CampaignFieldUpdateResponse defines model for CampaignFieldUpdateResponse.
type CampaignFieldUpdateResponse struct {
	// ID Unique identifier for the field
	ID string `json:"id"`

	// Links Meta links
	Links CampaignFieldUpdateResponseLinks `json:"links"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type CampaignFieldTypes `json:"type"`

	// Values List of values selected or from user input
	Values []string `json:"values"`
}

// CampaignFieldUpdateResponseLinks Meta links
type CampaignFieldUpdateResponseLinks struct {
	// Campaign URL of the campaign
	Campaign string `json:"campaign"`

	// CampaignFields URL of the campaign fields
	CampaignFields string `json:"campaign_fields"`
}

// CampaignListResponseItem defines model for CampaignListResponseItem.
type CampaignListResponseItem struct {
	// Description Description of the campaign
	Description *string `json:"description"`

	// EndDate End date of the campaign in ISO 8601 format: `YYYY-MM-DD`
	EndDate *string `json:"end_date"`

	// ID Unique identifier of the campaign
	ID string `json:"id"`

	// IsHidden Whether the campaign is hidden or not
	IsHidden bool `json:"is_hidden"`

	// Links Meta links
	Links CampaignListResponseItemLinks `json:"links"`

	// Owner Owner information of the campaign
	Owner CampaignListResponseItemOwner `json:"owner"`

	// ReferenceID Reference identifier of the campaign
	ReferenceID string `json:"reference_id"`

	// StartDate Start date of the campaign in ISO 8601 format: `YYYY-MM-DD`
	StartDate *string `json:"start_date"`

	// Title Title of the campaign
	Title string `json:"title"`
}

// CampaignListResponseItemLinks Meta links
type CampaignListResponseItemLinks struct {
	// Self URL of the campaign
	Self string `json:"self"`
}

// CampaignListResponseItemOwner Owner information of the campaign
type CampaignListResponseItemOwner struct {
	// ID Unique identifier of the owner of the campaign
	ID *string `json:"id,omitempty"`
}

// CampaignResponse defines model for CampaignResponse.
type CampaignResponse struct {
	Budget *BudgetResponse `json:"budget"`

	// CreatedAt Creation date and time of the campaign,  in ISO 8601 UTC format (`2035-02-10T10:40:45Z`)
	CreatedAt time.Time `json:"created_at"`

	// Description Description of the campaign
	Description *string `json:"description"`

	// EndDate End date of the campaign in ISO 8601 format: `YYYY-MM-DD`
	EndDate *string `json:"end_date"`

	// ID Unique identifier of the campaign
	ID string `json:"id"`

	// IsHidden Whether the campaign is hidden or not
	IsHidden bool `json:"is_hidden"`

	// Labels Labels associated to the campaign
	Labels []ResourceLabelResponse `json:"labels"`

	// Links Meta links
	Links CampaignResponseLinks `json:"links"`

	// ReferenceID Reference identifier of the campaign
	ReferenceID string `json:"reference_id"`

	// StartDate Start date of the campaign in ISO 8601 format: `YYYY-MM-DD`
	StartDate *string `json:"start_date"`

	// Status Status of the campaign
	Status CampaignResponseStatus `json:"status"`

	// Title Title of the campaign
	Title string `json:"title"`
}

// CampaignResponseLinks Meta links
type CampaignResponseLinks struct {
	// Brief URL of the campaign brief
	Brief *string `json:"brief"`

	// ChildCampaigns URLs of the child campaigns
	ChildCampaigns []string `json:"child_campaigns"`

	// Owner URL of the campaign owner
	Owner *string `json:"owner"`

	// ParentCampaign URL of the parent campaign
	ParentCampaign *string `json:"parent_campaign"`

	// Self URL of the campaign
	Self string `json:"self"`
}

// CampaignResponseStatus Status of the campaign
type CampaignResponseStatus string

// CampaignUpdateRequest defines model for CampaignUpdateRequest.
type CampaignUpdateRequest struct {
	// EndDate End date of the campaign in ISO 8601 format: `YYYY-MM-DD`
	EndDate *string `json:"end_date"`

	// OwnerID Unique identifier of the campaign owner
	OwnerID *string `json:"owner_id,omitempty"`

	// StartDate Start date of the campaign in ISO 8601 format: `YYYY-MM-DD`
	StartDate *string `json:"start_date"`

	// Title Title of the campaign
	Title *string `json:"title,omitempty"`
}

// CheckboxAndRadioTypeSettingsFieldCreatePayload defines model for CheckboxAndRadioTypeSettingsFieldCreatePayload.
type CheckboxAndRadioTypeSettingsFieldCreatePayload struct {
	// Choices Choices of the field
	Choices []CheckboxAndRadioTypeSettingsFieldCreatePayloadChoices `json:"choices"`

	// HelperText Used to help understand the functionality of the field.
	HelperText *string `json:"helper_text,omitempty"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive bool `json:"is_active"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`
}

// CheckboxAndRadioTypeSettingsFieldCreatePayloadChoices defines model for .
type CheckboxAndRadioTypeSettingsFieldCreatePayloadChoices struct {
	// Name Name of the choice
	Name string `json:"name"`
}

// CheckboxAndRadioTypeSettingsFieldResponse defines model for CheckboxAndRadioTypeSettingsFieldResponse.
type CheckboxAndRadioTypeSettingsFieldResponse struct {
	// Choices Choices of the field
	Choices []CheckboxAndRadioTypeSettingsFieldResponseChoices `json:"choices"`

	// HelperText Used to help understand functionality of field
	HelperText string `json:"helper_text"`

	// ID Identifier of the field
	ID string `json:"id"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive bool `json:"is_active"`

	// Links Meta links
	Links BaseSettingsFieldsResponseLinks `json:"links"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`
}

// CheckboxAndRadioTypeSettingsFieldResponseChoices defines model for .
type CheckboxAndRadioTypeSettingsFieldResponseChoices struct {
	// ID Identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// ChoiceDisplayOption An enumeration.
type ChoiceDisplayOption string

// ChoiceFieldDefinition defines model for ChoiceFieldDefinition.
type ChoiceFieldDefinition struct {
	// Choices Choices of the field
	Choices map[string]string `json:"choices"`
	Core    CoreFieldDef      `json:"core"`

	// DefaultValues Default values of the field
	DefaultValues *[]string `json:"default_values,omitempty"`

	// DisplayOption An enumeration.
	DisplayOption *ChoiceDisplayOption `json:"display_option,omitempty"`
}

// ChoiceFieldValueModel defines model for ChoiceFieldValueModel.
type ChoiceFieldValueModel struct {
	// ChoiceKey Choice key of the field value
	ChoiceKey string `json:"choice_key"`

	// OrderIndex Order index of the field value
	OrderIndex *int `json:"order_index,omitempty"`
}

// CommentCreateRequest defines model for CommentCreateRequest.
type CommentCreateRequest struct {
	// Attachments List of comment attachments.
	Attachments *[]AttachmentRequest `json:"attachments,omitempty"`

	// Value Content of the comment. Markdown is supported. To mention a user belonging to the organization, use the `@[name](openapi-user-link)` format.
	Value string `json:"value"`
}

// CommentWithReplyCreateRequest defines model for CommentWithReplyCreateRequest.
type CommentWithReplyCreateRequest struct {
	// Attachments List of comment attachments.
	Attachments *[]AttachmentRequest `json:"attachments,omitempty"`

	// ParentCommentID Parent comment id to reply.
	ParentCommentID *string `json:"parent_comment_id,omitempty"`

	// Value Content of the comment. Markdown is supported. To mention a user belonging to the organization, use the `@[name](openapi-user-link)` format.
	Value string `json:"value"`
}

// CompleteMultipartUploadResponse defines model for CompleteMultipartUploadResponse.
type CompleteMultipartUploadResponse struct {
	// Key Unique identifier of the file combining all parts. **This key will be needed for various file / attachment addition to different CMP resources using endpoints e.g. `POST /v3/assets`, `POST /v3/campaigns/{id}/attachments` etc.**
	Key   string                               `json:"key"`
	Links CompleteMultipartUploadResponseLinks `json:"links"`
}

// CompleteMultipartUploadResponseLinks defines model for .
type CompleteMultipartUploadResponseLinks struct {
	// Status URL to check upload status
	Status string `json:"status"`
}

// CompleteSCContentPreviewResponse defines model for CompleteSCContentPreviewResponse.
type CompleteSCContentPreviewResponse = interface{}

// ContentDetailsModel defines model for ContentDetailsModel.
type ContentDetailsModel struct {
	// ContentGUID Unique identifier of the content
	ContentGUID string                     `json:"content_guid"`
	ContentType *VersionedContentTypeModel `json:"content_type,omitempty"`

	// ContentTypeGUID Unique identifier of the content type
	ContentTypeGUID string `json:"content_type_guid"`

	// ContentTypeName Name of the content type
	ContentTypeName string `json:"content_type_name"`

	// CreatedAt Date and time on which the content was created, in ISO 8601 UTC format
	CreatedAt time.Time `json:"created_at"`

	// CreatedBy Unique identifier of the user who created the content
	CreatedBy string `json:"created_by"`

	// Expired Expired status of the content
	Expired *bool `json:"expired,omitempty"`

	// ExpiryDatetime Date and time on which the content will expire, in ISO 8601 UTC format
	ExpiryDatetime      *time.Time                  `json:"expiry_datetime,omitempty"`
	LatestFieldsVersion ContentFieldsVersionDetails `json:"latest_fields_version"`

	// Links Meta links
	Links ContentDetailsModelLinks `json:"links"`

	// PrimaryLocale The primary locale of the content
	PrimaryLocale *string `json:"primary_locale,omitempty"`

	// RootContent This is true when the content is not embedded in another content through a content type reference field
	RootContent bool `json:"root_content"`

	// Source Source of the content
	Source *string `json:"source,omitempty"`

	// SourceID Source ID of the content
	SourceID *string `json:"source_id,omitempty"`

	// SourceMetadata Source metadata of the content
	SourceMetadata *string `json:"source_metadata,omitempty"`

	// TemplateGUID Template GUID of the content
	TemplateGUID *string `json:"template_guid,omitempty"`

	// Title Title of the content
	Title string `json:"title"`

	// UpdatedAt Date and time on which the content was last updated, in ISO 8601 UTC format
	UpdatedAt time.Time `json:"updated_at"`

	// UpdatedBy Unique identifier of the user who last updated the content
	UpdatedBy string `json:"updated_by"`
}

// ContentDetailsModelLinks Meta links
type ContentDetailsModelLinks struct {
	// Definition URL of the content definition
	Definition *string `json:"definition,omitempty"`

	// Self URL of the content
	Self *string `json:"self,omitempty"`
}

// ContentFieldValueExpandedModel defines model for ContentFieldValueExpandedModel.
type ContentFieldValueExpandedModel struct {
	ContentDetails ContentDetailsModel `json:"content_details"`

	// ContentGUID Content GUID of the field value
	ContentGUID string `json:"content_guid"`

	// ContentURL Content URL of the field value
	ContentURL *string `json:"content_url,omitempty"`

	// Embedded Indicates whether the field value is embedded
	Embedded bool `json:"embedded"`

	// OrderIndex Order index of the field value
	OrderIndex *int `json:"order_index,omitempty"`
}

// ContentFieldValueModel defines model for ContentFieldValueModel.
type ContentFieldValueModel struct {
	// ContentGUID Content GUID of the field value
	ContentGUID string `json:"content_guid"`

	// ContentURL Content URL of the field value
	ContentURL *string `json:"content_url,omitempty"`

	// Embedded Indicates whether the field value is embedded
	Embedded bool `json:"embedded"`

	// OrderIndex Order index of the field value
	OrderIndex *int `json:"order_index,omitempty"`
}

// ContentFieldValueWithEmbeddedModel defines model for ContentFieldValueWithEmbeddedModel.
type ContentFieldValueWithEmbeddedModel struct {
	ContentDetails RecursiveStructuredContent `json:"content_details"`
}

// ContentFieldWithEmbeddedPatchValueModel defines model for ContentFieldWithEmbeddedPatchValueModel.
type ContentFieldWithEmbeddedPatchValueModel struct {
	ContentID    string                                `json:"content_id"`
	PatchDetails RecursivePatchStructuredContentFields `json:"patch_details"`
}

// ContentFieldsVersionDetails defines model for ContentFieldsVersionDetails.
type ContentFieldsVersionDetails struct {
	// ContentHash Content hash of the version
	ContentHash string `json:"content_hash"`

	// CreatedAt Date and time on which the version was created, in ISO 8601 UTC format
	CreatedAt time.Time `json:"created_at"`

	// CreatedBy Unique identifier of the user who created the version
	CreatedBy string `json:"created_by"`

	// Fields List of fields
	Fields map[string][]LocalizedFieldValues `json:"fields"`

	// SourceID Source ID of the version
	SourceID *string `json:"source_id,omitempty"`

	// SourceMetadata Source metadata of the version
	SourceMetadata *string                         `json:"source_metadata,omitempty"`
	Validation     *ContentFieldsVersionValidation `json:"validation,omitempty"`

	// VersionGUID Unique identifier of the version
	VersionGUID string `json:"version_guid"`
}

// ContentFieldsVersionValidation defines model for ContentFieldsVersionValidation.
type ContentFieldsVersionValidation struct {
	// Fields Validation for the fields
	Fields *map[string]map[string]ErrorReason `json:"fields,omitempty"`
}

// ContentGraphListResponse defines model for ContentGraphListResponse.
type ContentGraphListResponse struct {
	// Data List of content graph instances
	Data       []ContentGraphResponse `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ContentGraphQueryRequest GraphQL request envelope. Both properties are forwarded to Optimizely Graph verbatim, so they use the GraphQL protocol names rather than the snake_case convention used elsewhere in this API.
type ContentGraphQueryRequest struct {
	// Query GraphQL query to execute against the content graph instance
	Query string `json:"query"`

	// Variables Values for the variables declared by `query`
	Variables *map[string]interface{} `json:"variables,omitempty"`
}

// ContentGraphQueryResponse The content graph's GraphQL response, passed through unchanged. Its shape is determined by the submitted query: `data` (and `extensions`) for a successful query, or `errors` when the graph rejects it.
type ContentGraphQueryResponse struct {
	// Data Query result, shaped by the submitted GraphQL query.
	Data *map[string]interface{} `json:"data,omitempty"`

	// Errors Present when the graph rejects the query.
	Errors *[]map[string]interface{} `json:"errors,omitempty"`

	// Extensions Graph metadata such as correlation id and query cost.
	Extensions           *map[string]interface{} `json:"extensions,omitempty"`
	AdditionalProperties map[string]interface{}  `json:"-"`
}

// ContentGraphResponse defines model for ContentGraphResponse.
type ContentGraphResponse struct {
	// InstanceCategory Category of the content graph instance, either `new` or `existing`. `existing` instances are used by CMS organizations.
	InstanceCategory ContentGraphResponseInstanceCategory `json:"instance_category"`

	// InstanceID Unique identifier of the content graph instance
	InstanceID string `json:"instance_id"`

	// Name Human-readable name of the content graph instance
	Name string `json:"name"`

	// SourceID Namespace within the content graph instance that holds the organization's content. `existing` instances use `cmp` for library assets and `cmpsc` for structured contents, `new` instances use `default`.
	SourceID string `json:"source_id"`
}

// ContentGraphResponseInstanceCategory Category of the content graph instance, either `new` or `existing`. `existing` instances are used by CMS organizations.
type ContentGraphResponseInstanceCategory string

// ContentMigrationSummary defines model for ContentMigrationSummary.
type ContentMigrationSummary struct {
	// Errored The number of content items that failed to migrate.
	Errored *int `json:"errored,omitempty"`

	// NotStarted The number of content items that haven't started migrating.
	NotStarted *int `json:"not_started,omitempty"`

	// Skipped The number of content items that are skipped due to content type version ID mismatch.
	Skipped *int `json:"skipped,omitempty"`

	// Succeeded The number of content items that migrated successfully.
	Succeeded *int `json:"succeeded,omitempty"`

	// Total The total number of content items.
	Total *int `json:"total,omitempty"`
}

// ContentTypeFieldDefinition defines model for ContentTypeFieldDefinition.
type ContentTypeFieldDefinition struct {
	// AllowRefEdit Whether to allow ref editing
	AllowRefEdit *bool `json:"allow_ref_edit,omitempty"`

	// AllowedContentTypes List of allowed content types
	AllowedContentTypes []string `json:"allowed_content_types"`

	// ContentTypeLinks Links related to the content type
	ContentTypeLinks *map[string]AllowedContentTypeItem `json:"content_type_links,omitempty"`
	Core             CoreFieldDef                       `json:"core"`

	// DefaultValue Default value for the field
	DefaultValue *string `json:"default_value,omitempty"`

	// RefType Ref Type:
	//   * `1` Refer only
	//   * `2` Refer and create embed
	//   * `3` Create only
	RefType ContentTypeFieldEmbedMixConfig `json:"ref_type"`
}

// ContentTypeFieldEmbedMixConfig Ref Type:
//   - `1` Refer only
//   - `2` Refer and create embed
//   - `3` Create only
type ContentTypeFieldEmbedMixConfig int

// ContentTypeListingOption List Type:
//   - `1` - Component only
//   - `2` - Content only
//   - `3` - Component and Content only
type ContentTypeListingOption string

// CoreContentType defines model for CoreContentType.
type CoreContentType struct {
	// Component Indicates whether the content type is a component
	Component bool `json:"component"`

	// Description Description of the content type
	Description *string `json:"description,omitempty"`

	// Disabled Disabled status of the content type
	Disabled *bool `json:"disabled,omitempty"`

	// Name Name of the content type
	Name string `json:"name"`

	// ThumbnailGUID Thumbnail GUID of the content type
	ThumbnailGUID *string `json:"thumbnail_guid,omitempty"`
}

// CoreFieldDef defines model for CoreFieldDef.
type CoreFieldDef struct {
	// EditorMetadata Editor metadata of the field
	EditorMetadata *CoreFieldDef_EditorMetadata `json:"editor_metadata,omitempty"`

	// FieldType An enumeration.
	FieldType *FieldType `json:"field_type,omitempty"`

	// HelpText Helptext of the field
	HelpText *string `json:"help_text,omitempty"`

	// IsList Indicates whether the field is a list
	IsList bool `json:"is_list"`

	// IsRequired Required status of the field
	IsRequired bool `json:"is_required"`

	// Key Key of the field
	Key string `json:"key"`

	// MaxListLength Maximum length of the list
	MaxListLength *int `json:"max_list_length,omitempty"`

	// MinListLength Minimum length of the list
	MinListLength *int `json:"min_list_length,omitempty"`

	// Name Name of the field
	Name string `json:"name"`

	// NeedInternationalization Indicates whether the field needs internationalization
	NeedInternationalization bool `json:"need_internationalization"`

	// OrderIndex Order index of the field
	OrderIndex *int `json:"order_index,omitempty"`

	// SourceID Source ID of the field
	SourceID *string `json:"source_id,omitempty"`

	// SourceMetadata Source metadata of the field
	SourceMetadata *string `json:"source_metadata,omitempty"`
}

// CoreFieldDefEditorMetadata0 defines model for .
type CoreFieldDefEditorMetadata0 = map[string]interface{}

// CoreFieldDefEditorMetadata1 defines model for .
type CoreFieldDefEditorMetadata1 = []interface{}

// CoreFieldDef_EditorMetadata Editor metadata of the field
type CoreFieldDef_EditorMetadata struct {
	union json.RawMessage
}

// CreateFieldResponse defines model for CreateFieldResponse.
type CreateFieldResponse struct {
	// ID Identifier of the field that has been created.
	ID string `json:"id"`
}

// CreateMultipartUploadRequest defines model for CreateMultipartUploadRequest.
type CreateMultipartUploadRequest struct {
	// FileSize Size of the file in bytes
	FileSize int `json:"file_size"`

	// PartSize The size of each part of the file to be uploaded. The size of the last part can be smaller and hence can be ignored. `upload_part_count` is calculated based on `file_size` and `part_size`. **The maximum allowed `upload_part_count` is 10000.** So use larger `part_size` for larger files to keep `upload_part_count` under 10000.
	PartSize *int `json:"part_size,omitempty"`
}

// CreateMultipartUploadResponse defines model for CreateMultipartUploadResponse.
type CreateMultipartUploadResponse struct {
	// ExpiresAt Expiration time of the upload URLs
	ExpiresAt time.Time `json:"expires_at"`

	// ID ID of the multipart upload
	ID    string                             `json:"id"`
	Links CreateMultipartUploadResponseLinks `json:"links"`

	// UploadPartCount Number of parts for the upload
	UploadPartCount int `json:"upload_part_count"`

	// UploadPartUrls Array of pre-signed URLs for each part. **Each file part must be uploaded to the URL using PUT requests following the order of the URLs**
	UploadPartUrls []string `json:"upload_part_urls"`
}

// CreateMultipartUploadResponseLinks defines model for .
type CreateMultipartUploadResponseLinks struct {
	// Complete URL to complete the upload
	Complete string `json:"complete"`

	// Status URL to check the upload status
	Status string `json:"status"`
}

// CreateSCContentTypeManagedMigrationResponse defines model for CreateSCContentTypeManagedMigrationResponse.
type CreateSCContentTypeManagedMigrationResponse struct {
	// Created False if the managed migration job is not created.
	Created bool `json:"created"`
}

// CreateTaskStructuredContentDraftResponse defines model for CreateTaskStructuredContentDraftResponse.
type CreateTaskStructuredContentDraftResponse = map[string]interface{}

// CreativeAssetRequest defines model for CreativeAssetRequest.
type CreativeAssetRequest struct {
	// Key Unique identifier of the file upload session. This is the `upload_meta_fields.key` field retrieved from the `/v3/upload-url` endpoint.
	Key string `json:"key"`

	// Name Filename of the creative asset
	Name string `json:"name"`
}

// CreativeAssetResponse defines model for CreativeAssetResponse.
type CreativeAssetResponse struct {
	// ID Unique identifier for the creative asset
	ID string `json:"id"`

	// Name Filename of the creative asset
	Name string `json:"name"`

	// URL Download URL of the creative asset
	URL string `json:"url"`
}

// CurrencyCustomField defines model for CurrencyCustomField.
type CurrencyCustomField struct {
	// CurrencyCode Curency code of custom field
	CurrencyCode string `json:"currency_code"`

	// DecimalPlaces Value of decimal places
	DecimalPlaces *int `json:"decimal_places"`

	// HasThousandSeparator Indicator of thousand separator
	HasThousandSeparator bool `json:"has_thousand_separator"`

	// Label Name of the custom field
	Label string `json:"label"`
}

// CurrencyNumberTypeSettingsFieldCreatePayload defines model for CurrencyNumberTypeSettingsFieldCreatePayload.
type CurrencyNumberTypeSettingsFieldCreatePayload struct {
	// CurrencyCode Currency code of the numerical field. This must be a valid currency code. If this field is empty, `USD` is set as the default.
	CurrencyCode *string `json:"currency_code,omitempty"`

	// DecimalPlaces Decimal place of the numerical field. This must be greather than 0. If this field is empty, 2 is set as the default.
	DecimalPlaces *int `json:"decimal_places,omitempty"`

	// HasThousandSeparator Whether the numerical field has a thousand separator
	HasThousandSeparator *bool `json:"has_thousand_separator,omitempty"`

	// HelperText Used to help understand the functionality of the field.
	HelperText *string `json:"helper_text,omitempty"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive bool `json:"is_active"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`
}

// CurrencyNumberTypeSettingsFieldResponse defines model for CurrencyNumberTypeSettingsFieldResponse.
type CurrencyNumberTypeSettingsFieldResponse struct {
	// CurrencyCode Currency code of the numerical field
	CurrencyCode string `json:"currency_code"`

	// DecimalPlaces Decimal place of the numerical field
	DecimalPlaces int `json:"decimal_places"`

	// HasThousandSeparator Whether the numerical field has a thousand separator
	HasThousandSeparator bool `json:"has_thousand_separator"`

	// HelperText Used to help understand functionality of field
	HelperText string `json:"helper_text"`

	// ID Identifier of the field
	ID string `json:"id"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive bool `json:"is_active"`

	// Links Meta links
	Links BaseSettingsFieldsResponseLinks `json:"links"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`
}

// CurrencyNumberTypeSettingsFieldUpdatePayload defines model for CurrencyNumberTypeSettingsFieldUpdatePayload.
type CurrencyNumberTypeSettingsFieldUpdatePayload struct {
	// CurrencyCode Currency code of the numerical field. This must be a valid currency code. If this field is empty, `USD` is set as the default.
	CurrencyCode *string `json:"currency_code,omitempty"`

	// DecimalPlaces Decimal place of the numerical field. This must be greather than 0. If this field is empty, 2 is set as the default.
	DecimalPlaces *int `json:"decimal_places,omitempty"`

	// HasThousandSeparator Whether the numerical field has a thousand separator
	HasThousandSeparator *bool `json:"has_thousand_separator,omitempty"`

	// HelperText Used to help understand the functionality of the field.
	HelperText *string `json:"helper_text,omitempty"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive *bool `json:"is_active,omitempty"`

	// Name Name of the field
	Name *string `json:"name,omitempty"`

	// Type Type of the field. This field is required but not updateable. Type of a field will always stay the same. Update request body must contain atleast one other updateable field.
	Type string `json:"type"`
}

// DateTypeFormFieldRequest defines model for DateTypeFormFieldRequest.
type DateTypeFormFieldRequest struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// Type Type of the form field
	Type string `json:"type"`

	// Values Array of date-time values, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	Values []time.Time `json:"values"`
}

// DateTypeFormFieldResponse defines model for DateTypeFormFieldResponse.
type DateTypeFormFieldResponse struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// IsProtected True when this form field is read-only and only organization admins can modify it. Absent or false for editable fields.
	IsProtected *bool `json:"is_protected,omitempty"`

	// Type Type of the form field
	Type string `json:"type"`

	// Values Array of date-time values, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	Values []time.Time `json:"values"`
}

// DateTypeObjectField defines model for DateTypeObjectField.
type DateTypeObjectField struct {
	// ID Unique identifier of the field
	ID string `json:"id"`

	// Type Type of the field
	Type string `json:"type"`

	// Values Array of date-time values, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	Values []time.Time `json:"values"`
}

// DateTypeObjectFieldUpdatePayload defines model for DateTypeObjectFieldUpdatePayload.
type DateTypeObjectFieldUpdatePayload struct {
	// Type Type of the field
	Type string `json:"type"`

	// Values Array of date-time values, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	Values []time.Time `json:"values"`
}

// DatetimeFieldDefinition defines model for DatetimeFieldDefinition.
type DatetimeFieldDefinition struct {
	Core CoreFieldDef `json:"core"`

	// DefaultValues Default values of the field
	DefaultValues *[]time.Time `json:"default_values,omitempty"`

	// MaxDate Maximum date of the field
	MaxDate time.Time `json:"max_date"`

	// MinDate Minimum date of the field
	MinDate time.Time `json:"min_date"`
}

// DatetimeFieldValueModel defines model for DatetimeFieldValueModel.
type DatetimeFieldValueModel struct {
	// DatetimeValue Value of the field value
	DatetimeValue time.Time `json:"datetime_value"`

	// OrderIndex Order index of the field value
	OrderIndex *int `json:"order_index,omitempty"`
}

// DeleteFieldValueModel defines model for DeleteFieldValueModel.
type DeleteFieldValueModel struct {
	// Delete Should be true if the field value needs to be deleted
	Delete bool `json:"delete"`

	// OrderIndex Order index of the field value
	OrderIndex *int `json:"order_index,omitempty"`
}

// DeleteLocaleFieldValueModel defines model for DeleteLocaleFieldValueModel.
type DeleteLocaleFieldValueModel struct {
	// Delete Should be true if the field value with locale needs to be deleted
	Delete bool   `json:"delete"`
	Locale string `json:"locale"`
}

// DetailedAssetRenditionResponse defines model for DetailedAssetRenditionResponse.
type DetailedAssetRenditionResponse struct {
	// AltText Alternative text for rendition
	AltText *string `json:"alt_text"`

	// AssetType Type of the original asset the rendition was generated from
	AssetType string `json:"asset_type"`

	// CreatedAt Date and time on which the rendition was created, in ISO 8601 UTC format
	CreatedAt time.Time `json:"created_at"`

	// Height The height (in pixels) of the rendition
	Height *int `json:"height"`

	// ID Unique identifier of the rendition
	ID string `json:"id"`

	// MimeType MIME type of the rendition
	MimeType string `json:"mime_type"`

	// ModifiedAt Date and time on which the rendition was last modified, in ISO 8601 UTC format
	ModifiedAt time.Time `json:"modified_at"`

	// Name Name of the rendition
	Name string `json:"name"`

	// OriginalAssetID ID of the original asset the rendition was generated from
	OriginalAssetID string `json:"original_asset_id"`

	// RenditionConfigID ID of rendition config used to generate the rendition
	RenditionConfigID *string `json:"rendition_config_id,omitempty"`

	// Status Status of the rendition's generation
	Status DetailedAssetRenditionResponseStatus `json:"status"`

	// URL URL of the rendition
	URL string `json:"url"`

	// Width The width (in pixels) of the rendition
	Width *int `json:"width"`
}

// DetailedAssetRenditionResponseStatus Status of the rendition's generation
type DetailedAssetRenditionResponseStatus string

// DropdownTypeObjectField defines model for DropdownTypeObjectField.
type DropdownTypeObjectField struct {
	// ID Unique identifier of the field
	ID string `json:"id"`

	// Type Type of the field
	Type string `json:"type"`

	// Values Array of choice ID. Multiple choice values are not acceptable for `is_multi_select=false`.
	Values []string `json:"values"`
}

// DropdownTypeObjectFieldUpdatePayload defines model for DropdownTypeObjectFieldUpdatePayload.
type DropdownTypeObjectFieldUpdatePayload struct {
	// Type Type of the field
	Type string `json:"type"`

	// Values Array of choice ID. Multiple choice values are not acceptable for `is_multi_select=false`.
	Values []string `json:"values"`
}

// DropdownTypeSettingsFieldCreatePayload defines model for DropdownTypeSettingsFieldCreatePayload.
type DropdownTypeSettingsFieldCreatePayload struct {
	// Choices Choices of the field
	Choices []DropdownTypeSettingsFieldCreatePayloadChoices `json:"choices"`

	// HelperText Used to help understand the functionality of the field.
	HelperText *string `json:"helper_text,omitempty"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive bool `json:"is_active"`

	// IsMultiSelect Allow users to select multiple values
	IsMultiSelect bool `json:"is_multi_select"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`
}

// DropdownTypeSettingsFieldCreatePayloadChoices defines model for .
type DropdownTypeSettingsFieldCreatePayloadChoices struct {
	// Name Name of the choice
	Name string `json:"name"`
}

// DropdownTypeSettingsFieldResponse defines model for DropdownTypeSettingsFieldResponse.
type DropdownTypeSettingsFieldResponse struct {
	// Choices Choices of the field
	Choices []DropdownTypeSettingsFieldResponseChoices `json:"choices"`

	// HelperText Used to help understand functionality of field
	HelperText string `json:"helper_text"`

	// ID Identifier of the field
	ID string `json:"id"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive bool `json:"is_active"`

	// IsMultiSelect Allow users to select multiple values
	IsMultiSelect bool `json:"is_multi_select"`

	// Links Meta links
	Links BaseSettingsFieldsResponseLinks `json:"links"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`
}

// DropdownTypeSettingsFieldResponseChoices defines model for .
type DropdownTypeSettingsFieldResponseChoices struct {
	// ID Identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// Error Error payload
type Error struct {
	// Errors Additional information
	Errors *map[string]interface{} `json:"errors,omitempty"`

	// Message Message describing the error
	Message              string                 `json:"message"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// EventCreateRequest defines model for EventCreateRequest.
type EventCreateRequest struct {
	// CampaignID ID of the campaign under which the event is to be created
	CampaignID *string `json:"campaign_id,omitempty"`

	// Description Description of the event
	Description *string `json:"description,omitempty"`

	// EndDate End date and time of the event in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`
	EndDate time.Time `json:"end_date"`

	// Fields List of fields to be associated with the event
	Fields *[]EventCreateRequest_Fields_Item `json:"fields,omitempty"`

	// IsAllDay Indicate if the event is all day long
	IsAllDay bool `json:"is_all_day"`

	// StartDate Start date and time of the event in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`
	StartDate time.Time `json:"start_date"`

	// Title Title of the event
	Title string `json:"title"`
}

// EventCreateRequest_Fields_Item defines model for EventCreateRequest.fields.Item.
type EventCreateRequest_Fields_Item struct {
	union json.RawMessage
}

// EventFieldsUpdateRequest defines model for EventFieldsUpdateRequest.
type EventFieldsUpdateRequest = []EventFieldsUpdateRequest_Item

// EventFieldsUpdateRequest_Item defines model for EventFieldsUpdateRequest.Item.
type EventFieldsUpdateRequest_Item struct {
	union json.RawMessage
}

// EventResponse defines model for EventResponse.
type EventResponse struct {
	// CampaignID Unique identifier of the campaign associated with the event
	CampaignID *string `json:"campaign_id"`

	// CreatedBy Unique identifier of the user who created the event
	CreatedBy string `json:"created_by"`

	// Description Description of the event
	Description *string `json:"description"`

	// EndDate End date of the event in ISO 8601 UTC format
	EndDate string `json:"end_date"`

	// ID Unique identifier of the event
	ID string `json:"id"`

	// IsAllDay Whether the event is a day long event or not
	IsAllDay bool `json:"is_all_day"`

	// IsArchived Whether the event is archived or not
	IsArchived bool `json:"is_archived"`

	// Links Meta links
	Links EventResponseLinks `json:"links"`

	// ReferenceID Reference ID of the event
	ReferenceID string `json:"reference_id"`

	// StartDate Start date of the event in ISO 8601 UTC format
	StartDate string `json:"start_date"`

	// Title Title of the event
	Title string `json:"title"`
}

// EventResponseLinks Meta links
type EventResponseLinks struct {
	// Campaign URL of the campaign associated with the event
	Campaign *string `json:"campaign"`

	// Self URL of the event
	Self string `json:"self"`
}

// EventUpdateRequest defines model for EventUpdateRequest.
type EventUpdateRequest struct {
	// Description Description of the event
	Description *string `json:"description,omitempty"`

	// EndDate End date and time of the event in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`
	EndDate *time.Time `json:"end_date,omitempty"`

	// IsAllDay Indicate if the event is all day long
	IsAllDay *bool `json:"is_all_day,omitempty"`

	// StartDate Start date and time of the event in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`
	StartDate *time.Time `json:"start_date,omitempty"`

	// Title Title of the event
	Title *string `json:"title,omitempty"`
}

// FeaturedImageResponse defines model for FeaturedImageResponse.
type FeaturedImageResponse struct {
	// AttributionText Attribution text of the featured image
	AttributionText *string `json:"attribution_text"`

	// Caption Caption of the featured image
	Caption string `json:"caption"`

	// Description Description of the featured image
	Description *string `json:"description"`

	// Height Height of the featured image in pixels
	Height int `json:"height"`

	// MimeType MIME type of the featured image
	MimeType string `json:"mime_type"`

	// Source Source of the featured image
	Source FeaturedImageResponseSource `json:"source"`

	// Thumbnail URL of the featured image thumbnail
	Thumbnail *string `json:"thumbnail"`

	// URL URL of the featured image
	URL string `json:"url"`

	// Width Width of the featured image in pixels
	Width int `json:"width"`
}

// FeaturedImageResponseSource Source of the featured image
type FeaturedImageResponseSource struct {
	// Name Organization or vendor name of the featured image
	Name *string `json:"name"`
}

// FieldBaseType defines model for FieldBaseType.
type FieldBaseType struct {
	// ID Identifier of the field
	ID string `json:"id"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// FieldListResponseItem defines model for FieldListResponseItem.
type FieldListResponseItem struct {
	union json.RawMessage
}

// FieldType An enumeration.
type FieldType string

// FieldTypeCheckbox defines model for FieldTypeCheckbox.
type FieldTypeCheckbox struct {
	// Choices Choices of the checkbox
	Choices []FieldTypeCheckboxChoices `json:"choices"`

	// ID Identifier of the field
	ID string `json:"id"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// FieldTypeCheckboxChoices defines model for .
type FieldTypeCheckboxChoices struct {
	// ID Identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// FieldTypeCommon defines model for FieldTypeCommon.
type FieldTypeCommon struct {
	// ID Identifier of the field
	ID string `json:"id"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// FieldTypeCurrencyNumber defines model for FieldTypeCurrencyNumber.
type FieldTypeCurrencyNumber struct {
	// CurrencyCode Currency code of the numerical field
	CurrencyCode string `json:"currency_code"`

	// DecimalPlaces Decimal place of the numerical field
	DecimalPlaces *int `json:"decimal_places"`

	// HasThousandSeparator Whether the numerical field has a thousand separator
	HasThousandSeparator bool `json:"has_thousand_separator"`

	// ID Identifier of the field
	ID string `json:"id"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// FieldTypeDropdown defines model for FieldTypeDropdown.
type FieldTypeDropdown struct {
	// Choices Choices of the dropdown
	Choices []FieldTypeDropdownChoices `json:"choices"`

	// ID Identifier of the field
	ID string `json:"id"`

	// IsMultiSelect Select multiple values from the dropdown
	IsMultiSelect bool `json:"is_multi_select"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// FieldTypeDropdownChoices defines model for .
type FieldTypeDropdownChoices struct {
	// ID Identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// FieldTypeLabel defines model for FieldTypeLabel.
type FieldTypeLabel struct {
	// Choices Choices of the label
	Choices []FieldTypeLabelChoices `json:"choices"`

	// ID Identifier of the field
	ID string `json:"id"`

	// IsMultiSelect Select multiple values from the label
	IsMultiSelect bool `json:"is_multi_select"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// FieldTypeLabelChoices defines model for .
type FieldTypeLabelChoices struct {
	// ID Identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// FieldTypePercentageNumber defines model for FieldTypePercentageNumber.
type FieldTypePercentageNumber struct {
	// DecimalPlaces Decimal place of the numerical field
	DecimalPlaces *int `json:"decimal_places"`

	// ID Identifier of the field
	ID string `json:"id"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// FieldTypeRadio defines model for FieldTypeRadio.
type FieldTypeRadio struct {
	// Choices Choices of the radio button
	Choices []FieldTypeRadioChoices `json:"choices"`

	// ID Identifier of the field
	ID string `json:"id"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// FieldTypeRadioChoices defines model for .
type FieldTypeRadioChoices struct {
	// ID Identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// FieldTypeSimpleNumber defines model for FieldTypeSimpleNumber.
type FieldTypeSimpleNumber struct {
	// DecimalPlaces Decimal place of the numerical field
	DecimalPlaces *int `json:"decimal_places"`

	// HasThousandSeparator Whether the numerical field has a thousand separator
	HasThousandSeparator bool `json:"has_thousand_separator"`

	// ID Identifier of the field
	ID string `json:"id"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`

	// Values List of selected values or from user input
	Values []string `json:"values"`
}

// FileTypeFormFieldRequest defines model for FileTypeFormFieldRequest.
type FileTypeFormFieldRequest struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// Type Type of the form field
	Type string `json:"type"`

	// Values Array of file inputs to add a file
	Values []FileTypeFormFieldValueRequest `json:"values"`
}

// FileTypeFormFieldResponse defines model for FileTypeFormFieldResponse.
type FileTypeFormFieldResponse struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// IsProtected True when this form field is read-only and only organization admins can modify it. Absent or false for editable fields.
	IsProtected *bool `json:"is_protected,omitempty"`

	// Type Type of the form field
	Type string `json:"type"`

	// Values Array of files
	Values []FileTypeFormFieldValueResponse `json:"values"`
}

// FileTypeFormFieldValueRequest defines model for FileTypeFormFieldValueRequest.
type FileTypeFormFieldValueRequest struct {
	// Key Unique identifier of the file upload session. This is the `upload_meta_fields.key` field retrieved from the `/v3/upload-url` endpoint.
	Key string `json:"key"`

	// Name Name of the file
	Name string `json:"name"`
}

// FileTypeFormFieldValueResponse defines model for FileTypeFormFieldValueResponse.
type FileTypeFormFieldValueResponse struct {
	// ID Unique identifier of the file
	ID string `json:"id"`

	// Name Name of the file
	Name string `json:"name"`

	// URL URL of the file
	URL string `json:"url"`
}

// FileURLBulkCreateRequest defines model for FileUrlBulkCreateRequest.
type FileURLBulkCreateRequest struct {
	// Guids List of guids of assets representing binary files to generate URLs for the corresponding files
	Guids []string `json:"guids"`
}

// FolderCreateRequest defines model for FolderCreateRequest.
type FolderCreateRequest struct {
	// Name Name of the folder
	Name string `json:"name"`

	// ParentFolderID ID of the parent folder
	ParentFolderID *string `json:"parent_folder_id"`
}

// FolderPermissionBulkCreateRequest defines model for FolderPermissionBulkCreateRequest.
type FolderPermissionBulkCreateRequest struct {
	// Permissions List of permissions to be granted to folder
	Permissions []FolderPermissionBulkCreateRequestPermissions `json:"permissions"`

	// Type Type of the accessor
	Type FolderPermissionBulkCreateRequestType `json:"type"`
}

// FolderPermissionBulkCreateRequestPermissionsAccessType level of access the accessor has to the folder
type FolderPermissionBulkCreateRequestPermissionsAccessType string

// FolderPermissionBulkCreateRequestPermissions defines model for .
type FolderPermissionBulkCreateRequestPermissions struct {
	// AccessType level of access the accessor has to the folder
	AccessType FolderPermissionBulkCreateRequestPermissionsAccessType `json:"access_type"`

	// ID Unique identifier of accessor
	ID string `json:"id"`

	// IsOwner Indicates if the accessor is the owner of the folder. A team cannot be an owner.
	IsOwner *bool `json:"is_owner,omitempty"`
}

// FolderPermissionBulkCreateRequestType Type of the accessor
type FolderPermissionBulkCreateRequestType string

// FolderPermissionListResponseItem defines model for FolderPermissionListResponseItem.
type FolderPermissionListResponseItem struct {
	// Data List of Permissions
	Data       []BasePermissionsResponseSchema `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// FolderPermissionUpdateRequest defines model for FolderPermissionUpdateRequest.
type FolderPermissionUpdateRequest struct {
	// AccessType Set the level of access the accessor has to the folder
	AccessType *FolderPermissionUpdateRequestAccessType `json:"access_type,omitempty"`

	// IsOwner Used to set an accessor as owner. This field can only be set to true as an owner cannot be removed without another accessor being set as the new owner. A team cannot be an owner.
	IsOwner *bool `json:"is_owner,omitempty"`
}

// FolderPermissionUpdateRequestAccessType Set the level of access the accessor has to the folder
type FolderPermissionUpdateRequestAccessType string

// FolderResponse defines model for FolderResponse.
type FolderResponse struct {
	// CreatedAt Date and time on which the folder was created,  in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// ID Unique identifier of the folder
	ID string `json:"id"`

	// Links Meta links
	Links FolderResponseLinks `json:"links"`

	// ModifiedAt Date and time of the most recent modification of the folder, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// Name Name of the folder
	Name string `json:"name"`

	// ParentFolderID ID of the parent folder
	ParentFolderID *string `json:"parent_folder_id"`

	// Path Folder location
	Path string `json:"path"`
}

// FolderResponseLinks Meta links
type FolderResponseLinks struct {
	// Assets URL of the assets inside the folder
	Assets string `json:"assets"`

	// ChildFolders URL of the children list of the folder
	ChildFolders string `json:"child_folders"`

	// ParentFolder URL of the parent of the folder
	ParentFolder *string `json:"parent_folder"`

	// Self URL of the folder
	Self string `json:"self"`
}

// FolderUpdateRequest defines model for FolderUpdateRequest.
type FolderUpdateRequest struct {
	// Name Name of the folder
	Name *string `json:"name,omitempty"`

	// ParentFolderID ID of the parent folder. If you pass null, the folder will be moved to the root.
	ParentFolderID *string `json:"parent_folder_id"`
	union          json.RawMessage
}

// FolderUpdateRequest0 defines model for .
type FolderUpdateRequest0 = interface{}

// FolderUpdateRequest1 defines model for .
type FolderUpdateRequest1 = interface{}

// FormField defines model for FormField.
type FormField struct {
	// CustomFieldType Type of the custom field when the form field is a custom field
	CustomFieldType *string `json:"custom_field_type"`

	// FieldType Type of template form field
	FieldType string `json:"field_type"`

	// Help Helper text for form field
	Help string `json:"help"`

	// Identifier Identifier of the form field
	Identifier string `json:"identifier"`

	// IsRequired Indicates whether the form field is required
	IsRequired bool `json:"is_required"`

	// LogicRules List of the template logic rules
	LogicRules []LogicRule `json:"logic_rules"`

	// SortOrder Sort order of the form field
	SortOrder int `json:"sort_order"`
}

// GenericNumberTypeSettingsFieldCreatePayload defines model for GenericNumberTypeSettingsFieldCreatePayload.
type GenericNumberTypeSettingsFieldCreatePayload struct {
	// DecimalPlaces Decimal place of the numerical field. This must be greather than 0. If this field is empty, 2 is set as the default.
	DecimalPlaces *int `json:"decimal_places,omitempty"`

	// HasThousandSeparator Whether the numerical field has a thousand separator
	HasThousandSeparator *bool `json:"has_thousand_separator,omitempty"`

	// HelperText Used to help understand the functionality of the field.
	HelperText *string `json:"helper_text,omitempty"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive bool `json:"is_active"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`
}

// GenericNumberTypeSettingsFieldResponse defines model for GenericNumberTypeSettingsFieldResponse.
type GenericNumberTypeSettingsFieldResponse struct {
	// DecimalPlaces Decimal place of the numerical field
	DecimalPlaces int `json:"decimal_places"`

	// HasThousandSeparator Whether the numerical field has a thousand separator
	HasThousandSeparator bool `json:"has_thousand_separator"`

	// HelperText Used to help understand functionality of field
	HelperText string `json:"helper_text"`

	// ID Identifier of the field
	ID string `json:"id"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive bool `json:"is_active"`

	// Links Meta links
	Links BaseSettingsFieldsResponseLinks `json:"links"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`
}

// GenericNumberTypeSettingsFieldUpdatePayload defines model for GenericNumberTypeSettingsFieldUpdatePayload.
type GenericNumberTypeSettingsFieldUpdatePayload struct {
	// DecimalPlaces Decimal place of the numerical field. This must be greather than 0. If this field is empty, 2 is set as the default.
	DecimalPlaces *int `json:"decimal_places,omitempty"`

	// HasThousandSeparator Whether the numerical field has a thousand separator
	HasThousandSeparator *bool `json:"has_thousand_separator,omitempty"`

	// HelperText Used to help understand the functionality of the field.
	HelperText *string `json:"helper_text,omitempty"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive *bool `json:"is_active,omitempty"`

	// Name Name of the field
	Name *string `json:"name,omitempty"`

	// Type Type of the field. This field is required but not updateable. Type of a field will always stay the same. Update request body must contain atleast one other updateable field.
	Type string `json:"type"`
}

// GenericSettingsFieldChoiceCreatePayload Choice for field types other than `label`
type GenericSettingsFieldChoiceCreatePayload struct {
	// Name Name of the choice
	Name string `json:"name"`
}

// GenericSettingsFieldChoiceUpdatePayload Choice for field types other than `label`
type GenericSettingsFieldChoiceUpdatePayload struct {
	// Name Name of the choice
	Name string `json:"name"`
}

// GetMultipartUploadStatusResponse defines model for GetMultipartUploadStatusResponse.
type GetMultipartUploadStatusResponse struct {
	// ExpiresAt Expiration time of the upload
	ExpiresAt time.Time `json:"expires_at"`

	// ID ID of the multipart upload
	ID string `json:"id"`

	// Key Unique identifier of the file combining all parts. **This key will be needed for various file / attachment addition to different CMP resources using endpoints e.g. `POST /v3/assets`, `POST /v3/campaigns/{id}/attachments` etc.**
	Key   string                                `json:"key"`
	Links GetMultipartUploadStatusResponseLinks `json:"links"`

	// Status Current status of the upload
	Status GetMultipartUploadStatusResponseStatus `json:"status"`

	// StatusMessage Detailed status message
	StatusMessage *string `json:"status_message"`
}

// GetMultipartUploadStatusResponseLinks defines model for .
type GetMultipartUploadStatusResponseLinks struct {
	// Self URL to this status endpoint
	Self string `json:"self"`
}

// GetMultipartUploadStatusResponseStatus Current status of the upload
type GetMultipartUploadStatusResponseStatus string

// GetUploadURLResponse defines model for GetUploadUrlResponse.
type GetUploadURLResponse struct {
	// UploadMetaFields Related meta fields to upload a file. **Items within this field must be sent along as paylod of the POST request, along with the file, to the presigned upload URL**
	UploadMetaFields GetUploadURLResponseUploadMetaFields `json:"upload_meta_fields"`

	// URL Presigned URL to use for uploading a file with an **HTTP POST** request as `multipart/form-data`
	URL string `json:"url"`
}

// GetUploadURLResponseUploadMetaFields Related meta fields to upload a file. **Items within this field must be sent along as paylod of the POST request, along with the file, to the presigned upload URL**
type GetUploadURLResponseUploadMetaFields struct {
	// Key Unique identifier of the file. **This key will be needed for various file / attachment addition to different CMP resources using endpoints e.g. `POST /v3/assets`, `POST /v3/campaigns/{id}/attachments` etc.**
	Key               string  `json:"key"`
	Policy            string  `json:"policy"`
	XAmzAlgorithm     string  `json:"x-amz-algorithm"`
	XAmzCredential    string  `json:"x-amz-credential"`
	XAmzDate          string  `json:"x-amz-date"`
	XAmzSecurityToken *string `json:"x-amz-security-token,omitempty"`
	XAmzSignature     string  `json:"x-amz-signature"`
}

// HTTPException defines model for HTTPException.
type HTTPException struct {
	// Detail Error message details.
	Detail *string `json:"detail,omitempty"`
}

// HTTPValidationError defines model for HTTPValidationError.
type HTTPValidationError struct {
	// Detail Details of the error
	Detail *[]ValidationError `json:"detail,omitempty"`
}

// Instruction defines model for Instruction.
type Instruction struct {
	// Description Description of the template instruction
	Description string `json:"description"`

	// ID Id of the template instruction
	ID string `json:"id"`
}

// JSONFieldValueModel defines model for JSONFieldValueModel.
type JSONFieldValueModel struct {
	// JSONValue Value of the field value
	JSONValue JSONFieldValueModel_JSONValue `json:"json_value"`

	// OrderIndex Order index of the field value
	OrderIndex *int `json:"order_index,omitempty"`
}

// JSONFieldValueModelJSONValue0 defines model for .
type JSONFieldValueModelJSONValue0 = map[string]interface{}

// JSONFieldValueModelJSONValue1 defines model for .
type JSONFieldValueModelJSONValue1 = []interface{}

// JSONFieldValueModel_JSONValue Value of the field value
type JSONFieldValueModel_JSONValue struct {
	union json.RawMessage
}

// KeyedPreviewCompletedModel defines model for KeyedPreviewCompletedModel.
type KeyedPreviewCompletedModel struct {
	// Completed URL for completed preview.
	Completed string `json:"completed"`

	// Locale If not specified, the `locale` will be infered from the `Content-Language` header from the `completed` URL.
	//  In case of multiple locales, only the first one will be considered and it will be converted to snake case.
	//  For example, if the `Content-Language` header from the `completed` URL returns `en-US, de, bn`, it will be inferred as `en_US`.
	//  If the `locale` is not specified and there is no `Content-Language` header, this will default to `en_US`.
	Locale *string `json:"locale,omitempty"`

	// MimeType If not specified, the `mimeType` will be inferred from the `Content-Type` header from the `completed` URL.
	//  If the `mimeType` is not specified and there is no `Content-Type` header, this will default to `text/html`.
	MimeType *string `json:"mimeType,omitempty"`

	// Name Human readable preview name to distinguish between generated preview outcomes.
	Name *string `json:"name,omitempty"`

	// OrderIndex Priority of a preview. For multiple previews, this will be used to sort the previews.
	OrderIndex *float32 `json:"orderIndex,omitempty"`
}

// KeyedPreviewErrorModel defines model for KeyedPreviewErrorModel.
type KeyedPreviewErrorModel struct {
	// Error URL for error preview.
	Error string `json:"error"`

	// Locale If not specified, the `locale` will be infered from the `Content-Language` header from the `error` URL.
	//  In case of multiple locales, only the first one will be considered and it will be converted to snake case.
	//  For example, if the `Content-Language` header from the `completed` URL returns `en-US, de, bn`, it will be inferred as `en_US`.
	//  If the `locale` is not specified and there is no `Content-Language` header, this will default to `en_US`.
	Locale *string `json:"locale,omitempty"`

	// MimeType If not specified, the `mimeType` will be inferred from the `Content-Type` header from the `error` URL.
	//  If the `mimeType` is not specified and there is no `Content-Type` header, this will default to `text/html`.
	MimeType *string `json:"mimeType,omitempty"`

	// Name Human readable preview name to distinguish between generated preview outcomes.
	Name *string `json:"name,omitempty"`

	// OrderIndex Priority of a preview. For multiple previews, this will be used to sort the previews.
	OrderIndex *float32 `json:"orderIndex,omitempty"`
}

// LabelAndDropdownTypeSettingsFieldUpdatePayload defines model for LabelAndDropdownTypeSettingsFieldUpdatePayload.
type LabelAndDropdownTypeSettingsFieldUpdatePayload struct {
	// HelperText Used to help understand the functionality of the field.
	HelperText *string `json:"helper_text,omitempty"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive *bool `json:"is_active,omitempty"`

	// IsMultiSelect Allow users to select multiple values
	IsMultiSelect *bool `json:"is_multi_select,omitempty"`

	// Name Name of the field
	Name *string `json:"name,omitempty"`

	// Type Type of the field. This field is required but not updateable. Type of a field will always stay the same. Update request body must contain atleast one other updateable field.
	Type string `json:"type"`
}

// LabelGroup defines model for LabelGroup.
type LabelGroup struct {
	// ID Unique identifier for the label group
	ID string `json:"id"`

	// Name Name of the label group
	Name string `json:"name"`

	// SourceOrgType Source organization type where the label group is listed.
	// `current` means the label group is listed in the organization with which the Open API App is associated.
	// `related` means the label group is listed in one of the organizations related to the `current` organization.
	SourceOrgType LabelGroupSourceOrgType `json:"source_org_type"`

	// Values Values of the label group
	Values []LabelGroupValue `json:"values"`
}

// LabelGroupSourceOrgType Source organization type where the label group is listed.
// `current` means the label group is listed in the organization with which the Open API App is associated.
// `related` means the label group is listed in one of the organizations related to the `current` organization.
type LabelGroupSourceOrgType string

// LabelGroupValue Values of a label group
type LabelGroupValue struct {
	// ID Unique identifier for the label group value
	ID string `json:"id"`

	// Name Name of the label group value
	Name string `json:"name"`
}

// LabelOnlyCustomField defines model for LabelOnlyCustomField.
type LabelOnlyCustomField struct {
	// Label Name of the custom field
	Label string `json:"label"`
}

// LabelTypeSettingsFieldChoiceCreatePayload defines model for LabelTypeSettingsFieldChoiceCreatePayload.
type LabelTypeSettingsFieldChoiceCreatePayload struct {
	// Color Code of choice color. Colors are limited to the following:
	// - Electric Blue (#4ECFD5)
	// - Golden Rod (#FFC700)
	// - Salmon Pink (#FF98A7)
	// - Blue Violet (#702BD5)
	// - Spring Green (#5FEEAD)
	// - Lavender Blue (#D6C4F2)
	// - Sky (#5F9FF6)
	// - Persian blue (#CB5DEB)
	// - Cool Gray (#9694B3)
	Color string `json:"color"`

	// Name Name of the choice
	Name string `json:"name"`
}

// LabelTypeSettingsFieldChoiceUpdatePayload Choice for field type of `label`
type LabelTypeSettingsFieldChoiceUpdatePayload struct {
	// Color Code of choice color. Colors are limited to the following:
	// - Electric Blue (#4ECFD5)
	// - Golden Rod (#FFC700)
	// - Salmon Pink (#FF98A7)
	// - Blue Violet (#702BD5)
	// - Spring Green (#5FEEAD)
	// - Lavender Blue (#D6C4F2)
	// - Sky (#5F9FF6)
	// - Persian blue (#CB5DEB)
	// - Cool Gray (#9694B3)
	Color *string `json:"color,omitempty"`

	// Name Name of the choice
	Name *string `json:"name,omitempty"`
}

// LabelTypeSettingsFieldCreatePayload defines model for LabelTypeSettingsFieldCreatePayload.
type LabelTypeSettingsFieldCreatePayload struct {
	// Choices Choices of the field
	Choices []LabelTypeSettingsFieldCreatePayloadChoices `json:"choices"`

	// HelperText Used to help understand the functionality of the field.
	HelperText *string `json:"helper_text,omitempty"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive bool `json:"is_active"`

	// IsMultiSelect Allow users to select multiple values
	IsMultiSelect bool `json:"is_multi_select"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`
}

// LabelTypeSettingsFieldCreatePayloadChoices defines model for .
type LabelTypeSettingsFieldCreatePayloadChoices struct {
	// Color Code of choice color. Colors are limited to the following:
	// - Electric Blue (#4ECFD5)
	// - Golden Rod (#FFC700)
	// - Salmon Pink (#FF98A7)
	// - Blue Violet (#702BD5)
	// - Spring Green (#5FEEAD)
	// - Lavender Blue (#D6C4F2)
	// - Sky (#5F9FF6)
	// - Persian blue (#CB5DEB)
	// - Cool Gray (#9694B3)
	Color string `json:"color"`

	// Name Name of the choice
	Name string `json:"name"`
}

// LabelTypeSettingsFieldResponse defines model for LabelTypeSettingsFieldResponse.
type LabelTypeSettingsFieldResponse struct {
	// Choices Choices of the field
	Choices []LabelTypeSettingsFieldResponseChoices `json:"choices"`

	// HelperText Used to help understand functionality of field
	HelperText string `json:"helper_text"`

	// ID Identifier of the field
	ID string `json:"id"`

	// IsActive Indicates active status of field. Inactive fields cannot be accessed in tasks, campaigns or workflows.
	IsActive bool `json:"is_active"`

	// IsMultiSelect Allow users to select multiple values
	IsMultiSelect bool `json:"is_multi_select"`

	// Links Meta links
	Links BaseSettingsFieldsResponseLinks `json:"links"`

	// Name Name of the field
	Name string `json:"name"`

	// Type Type of the field
	Type string `json:"type"`
}

// LabelTypeSettingsFieldResponseChoices defines model for .
type LabelTypeSettingsFieldResponseChoices struct {
	// Color Code of choice color. Colors are limited to the following:
	// - Electric Blue (#4ECFD5)
	// - Golden Rod (#FFC700)
	// - Salmon Pink (#FF98A7)
	// - Blue Violet (#702BD5)
	// - Spring Green (#5FEEAD)
	// - Lavender Blue (#D6C4F2)
	// - Sky (#5F9FF6)
	// - Persian blue (#CB5DEB)
	// - Cool Gray (#9694B3)
	Color *string `json:"color"`

	// ID Identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// LibraryArticle defines model for LibraryArticle.
type LibraryArticle struct {
	// Authors List of authors of the article
	Authors []ArticleAuthorResponse `json:"authors"`

	// CreatedAt Date and time on which the article was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// ExpiresAt Date and time for the expiration of the article, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ExpiresAt *time.Time `json:"expires_at"`

	// FileLocation Location of the folder containing the article in the library
	FileLocation string `json:"file_location"`

	// FolderID ID of the folder containing the article in the library
	FolderID *string `json:"folder_id"`

	// GroupID Unique identifier of the source article in case of copied or translated articles
	GroupID *string `json:"group_id"`

	// HTMLBody Content of the article
	HTMLBody string `json:"html_body"`

	// ID Unique identifier of the article
	ID string `json:"id"`

	// Images Featured images of the article
	Images []FeaturedImageResponse `json:"images"`

	// IsArchived Whether the article is archived or not
	IsArchived bool `json:"is_archived"`

	// Labels Labels associated with the article
	Labels []ResourceLabelResponse `json:"labels"`

	// LangCode Language code depicting the language of the article
	LangCode *string `json:"lang_code"`

	// MetaDescription Meta description of the article
	MetaDescription *string `json:"meta_description"`

	// MetaKeywords Meta Keywords of the article
	MetaKeywords []string `json:"meta_keywords"`

	// MetaTitle Meta title of the article
	MetaTitle *string `json:"meta_title"`

	// MetaURL Meta URL of the article
	MetaURL *string `json:"meta_url"`

	// ModifiedAt Date and time of the most recent modification of the article, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// OwnerID ID of the article owner
	OwnerID string `json:"owner_id"`

	// OwnerOrganizationID ID of owner organization of the article
	OwnerOrganizationID string `json:"owner_organization_id"`

	// PixelKey Unique tracking pixel of the article
	PixelKey string `json:"pixel_key"`

	// SourceArticle Link to the vendor article for Marketplace article or Task article ID for created article
	SourceArticle *string `json:"source_article"`

	// SourceName Name of the article's source
	SourceName *string `json:"source_name"`

	// Tags Tags for the article
	Tags []Tag `json:"tags"`

	// ThumbnailURL URL of the thumbnail of the article
	ThumbnailURL *string `json:"thumbnail_url"`

	// Title Title of the article
	Title string `json:"title"`

	// URL Link to the published article
	URL *string `json:"url"`

	// VersionID Version id of article
	VersionID string `json:"version_id"`

	// VersionNumber Version number of article
	VersionNumber int `json:"version_number"`
}

// LibraryAssetCreateRequest defines model for LibraryAssetCreateRequest.
type LibraryAssetCreateRequest struct {
	// FolderID ID of the folder where the asset should be added
	FolderID *string `json:"folder_id"`

	// Key Unique identifier of the file upload session. This is the `upload_meta_fields.key` field retrieved from the `/v3/upload-url` endpoint.
	Key string `json:"key"`

	// Title Title of the asset
	Title string `json:"title"`
}

// LibraryAssetDefaultValue defines model for LibraryAssetDefaultValue.
type LibraryAssetDefaultValue struct {
	// AssetGUID Unique identifier of the asset
	AssetGUID string `json:"asset_guid"`

	// AssetType An enumeration.
	AssetType LibraryAssetType `json:"asset_type"`
}

// LibraryAssetFieldDefinition defines model for LibraryAssetFieldDefinition.
type LibraryAssetFieldDefinition struct {
	AllowedTypes []LibraryAssetType `json:"allowed_types"`
	Core         CoreFieldDef       `json:"core"`

	// DefaultValues Default values for the field
	DefaultValues *[]LibraryAssetDefaultValue `json:"default_values,omitempty"`
}

// LibraryAssetFieldValueModel defines model for LibraryAssetFieldValueModel.
type LibraryAssetFieldValueModel struct {
	// AssetGUID Asset GUID of the field value
	AssetGUID string `json:"asset_guid"`

	// AssetType An enumeration.
	AssetType LibraryAssetType `json:"asset_type"`

	// Links Meta links
	Links *map[string]string `json:"links,omitempty"`

	// OrderIndex Order index of the field value
	OrderIndex *int `json:"order_index,omitempty"`
}

// LibraryAssetType An enumeration.
type LibraryAssetType string

// LibraryAssetVersionResponse defines model for LibraryAssetVersionResponse.
type LibraryAssetVersionResponse struct {
	// AssetID Unique identifier of the asset
	AssetID string `json:"asset_id"`

	// Content Content of the version
	Content LibraryAssetVersionResponseContent `json:"content"`

	// CreatedAt Date and time on which the version was created,  in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// Links Meta links
	Links LibraryAssetVersionResponseLinks `json:"links"`

	// MimeType MIME type of the version
	MimeType string `json:"mime_type"`

	// Title Title of the asset
	Title string `json:"title"`

	// Type Type of the version
	Type LibraryAssetVersionResponseType `json:"type"`

	// VersionNumber The serial number of the version
	VersionNumber int `json:"version_number"`
}

// LibraryAssetVersionResponseContentType Type of the content
type LibraryAssetVersionResponseContentType string

// LibraryAssetVersionResponseContent Content of the version
type LibraryAssetVersionResponseContent struct {
	// Type Type of the content
	Type LibraryAssetVersionResponseContentType `json:"type"`

	// Value Content of the version. It is the download URL for image, raw file, and video.
	Value string `json:"value"`
}

// LibraryAssetVersionResponseLinks Meta links
type LibraryAssetVersionResponseLinks struct {
	// Asset URL of the asset that the version is associated with
	Asset string `json:"asset"`
}

// LibraryAssetVersionResponseType Type of the version
type LibraryAssetVersionResponseType string

// LibraryImage defines model for LibraryImage.
type LibraryImage struct {
	// AllowSeIndexing Whether the image is allowed to be indexed by search engines
	AllowSeIndexing bool `json:"allow_se_indexing"`

	// AltText Alternative text for image
	AltText *string `json:"alt_text"`

	// AttributionText Attribution of Image
	AttributionText *string `json:"attribution_text"`

	// CreatedAt Date and time on which the image was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// Description Description of the image
	Description *string `json:"description"`

	// ExpiresAt Date and time for the expiration of the image, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ExpiresAt *time.Time `json:"expires_at"`

	// FileExtension File extension of the image file
	FileExtension *string `json:"file_extension"`

	// FileLocation Location of the folder containing the image in the library
	FileLocation string `json:"file_location"`

	// FileSize Size of the image in bytes
	FileSize int `json:"file_size"`

	// FocalPoint Focal point of the image
	FocalPoint *LibraryImageFocalPoint `json:"focal_point"`

	// FolderID ID of the folder containing the image in the library
	FolderID *string `json:"folder_id"`

	// ID Unique identifier of the image
	ID string `json:"id"`

	// ImageResolution Width and height of the image in pixels
	ImageResolution LibraryImageImageResolution `json:"image_resolution"`

	// IsArchived Whether the image is archived or not
	IsArchived bool `json:"is_archived"`

	// IsPublic Whether the image URL is public or not
	IsPublic bool `json:"is_public"`

	// Labels Labels associated with the image
	Labels []ResourceLabelResponse `json:"labels"`

	// MimeType MIME type of the image
	MimeType string `json:"mime_type"`

	// ModifiedAt Date and time of the most recent modification of the image, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// OwnerID ID of the image owner
	OwnerID string `json:"owner_id"`

	// OwnerOrganizationID ID of owner organization of the image
	OwnerOrganizationID string `json:"owner_organization_id"`

	// Tags Tags for the image
	Tags []Tag `json:"tags"`

	// ThumbnailURL URL of the thumbnail of the image
	ThumbnailURL *string `json:"thumbnail_url"`

	// Title Title of the image
	Title string `json:"title"`

	// URL Download the URL of the image
	URL string `json:"url"`

	// VersionID Version id of image
	VersionID string `json:"version_id"`

	// VersionNumber Version number of image
	VersionNumber int `json:"version_number"`
}

// LibraryImageFocalPoint Focal point of the image
type LibraryImageFocalPoint struct {
	// X X coordinate of the focal point
	X int `json:"x"`

	// Y Y coordinate of the focal point
	Y int `json:"y"`
}

// LibraryImageImageResolution Width and height of the image in pixels
type LibraryImageImageResolution struct {
	// Height Height of the image in pixels
	Height int `json:"height"`

	// Width Width of the image in pixels
	Width int `json:"width"`
}

// LibraryImageUpdateRequest defines model for LibraryImageUpdateRequest.
type LibraryImageUpdateRequest struct {
	// AllowSeIndexing Whether the image is allowed to be indexed by search engines
	AllowSeIndexing *bool `json:"allow_se_indexing,omitempty"`

	// AltText Alt text of the image. Set an empty string to clear the alt text.
	AltText *string `json:"alt_text,omitempty"`

	// AttributionText Attribution text of the image. Set an empty string to clear the attribution text.
	AttributionText *string `json:"attribution_text,omitempty"`

	// Description Description of the image. Set an empty string to clear the description.
	Description *string `json:"description,omitempty"`

	// ExpiresAt Date and time when the image expires, in ISO 8601 UTC format
	ExpiresAt *time.Time `json:"expires_at"`

	// FolderID ID of the folder containing the image in the library
	FolderID *string `json:"folder_id"`

	// IsArchived Whether the image should be archived
	IsArchived *bool `json:"is_archived,omitempty"`

	// IsPublic Whether the image URL should be public
	IsPublic *bool `json:"is_public,omitempty"`

	// Labels Please use [PUT /assets/{asset_id}/fields](https://docs.developers.optimizely.com/content-marketing-platform/reference/put_assets-asset-id-fields) or  [PUT /assets/{asset_id}/fields/{field_id}](https://docs.developers.optimizely.com/content-marketing-platform/reference/put_assets-asset-id-fields-field-id)  API to update label type asset fields
	// Deprecated: this property has been marked as deprecated upstream, but no `x-deprecated-reason` was set
	Labels *[]ResourceLabelRequest `json:"labels,omitempty"`

	// Tags Tags of the image. Provided tags will replace the existing ones. Set an empty array to clear the existing tags.
	Tags *[]string `json:"tags,omitempty"`

	// Title Title of the image
	Title *string `json:"title,omitempty"`
}

// LibraryRawFile defines model for LibraryRawFile.
type LibraryRawFile struct {
	// AllowSeIndexing Whether the raw file is allowed to be indexed by search engines
	AllowSeIndexing bool `json:"allow_se_indexing"`

	// AttributionText Attribution of raw file
	AttributionText *string `json:"attribution_text"`

	// CreatedAt Date and time on which the raw file was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// Description Description of the raw file
	Description *string `json:"description"`

	// ExpiresAt Date and time for the expiration of the raw file, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ExpiresAt *time.Time `json:"expires_at"`

	// FileExtension File extension of the raw file
	FileExtension *string `json:"file_extension"`

	// FileLocation Location of the folder containing the raw file in the library
	FileLocation string `json:"file_location"`

	// FileSize Size of the raw file in bytes
	FileSize int `json:"file_size"`

	// FolderID ID of the folder containing the raw file in the library
	FolderID *string `json:"folder_id"`

	// ID Unique identifier of the raw file
	ID string `json:"id"`

	// IsArchived Whether the raw file is archived or not
	IsArchived bool `json:"is_archived"`

	// IsPublic Whether the raw file URL is public or not
	IsPublic bool `json:"is_public"`

	// Labels Labels associated to the raw file
	Labels []ResourceLabelResponse `json:"labels"`

	// MimeType MIME type of the raw file
	MimeType string `json:"mime_type"`

	// ModifiedAt Date and time of the most recent modification of the raw file, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// OwnerID ID of the raw file owner
	OwnerID string `json:"owner_id"`

	// OwnerOrganizationID ID of owner organization of the raw file
	OwnerOrganizationID string `json:"owner_organization_id"`

	// Tags Tags for the raw file
	Tags []Tag `json:"tags"`

	// ThumbnailURL URL of the thumbnail of the raw file
	ThumbnailURL *string `json:"thumbnail_url"`

	// Title Title of the raw file
	Title string `json:"title"`

	// URL Download the URL of the raw file
	URL string `json:"url"`

	// VersionID Version id of raw file
	VersionID string `json:"version_id"`

	// VersionNumber Version number of raw file
	VersionNumber int `json:"version_number"`
}

// LibraryRawFileUpdateRequest defines model for LibraryRawFileUpdateRequest.
type LibraryRawFileUpdateRequest struct {
	// AllowSeIndexing Whether the raw file is allowed to be indexed by search engines
	AllowSeIndexing *bool `json:"allow_se_indexing,omitempty"`

	// AttributionText Attribution text of the raw file. Set an empty string to clear the attribution text.
	AttributionText *string `json:"attribution_text,omitempty"`

	// Description Description of the raw file. Set an empty string to clear the description.
	Description *string `json:"description,omitempty"`

	// ExpiresAt Date and time when the raw file expires, in ISO 8601 UTC format
	ExpiresAt *time.Time `json:"expires_at"`

	// FolderID ID of the folder containing the raw file in the library
	FolderID *string `json:"folder_id"`

	// IsArchived Whether the raw file should be archived
	IsArchived *bool `json:"is_archived,omitempty"`

	// IsPublic Whether the raw file URL should be public
	IsPublic *bool `json:"is_public,omitempty"`

	// Labels Please use [PUT /assets/{asset_id}/fields](https://docs.developers.optimizely.com/content-marketing-platform/reference/put_assets-asset-id-fields) or  [PUT /assets/{asset_id}/fields/{field_id}](https://docs.developers.optimizely.com/content-marketing-platform/reference/put_assets-asset-id-fields-field-id)  API to update label type asset fields
	// Deprecated: this property has been marked as deprecated upstream, but no `x-deprecated-reason` was set
	Labels *[]ResourceLabelRequest `json:"labels,omitempty"`

	// Tags Tags of the raw file. Provided tags will replace the existing ones. Set an empty array to clear the existing tags.
	Tags *[]string `json:"tags,omitempty"`

	// Title Title of the raw file
	Title *string `json:"title,omitempty"`
}

// LibraryStructuredContent defines model for LibraryStructuredContent.
type LibraryStructuredContent struct {
	ContentBody *ContentDetailsModel `json:"content_body,omitempty"`

	// CreatedAt Date and time on which the structured content was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// FileLocation Location of the folder containing the structured content in the library
	FileLocation string `json:"file_location"`

	// FolderID ID of the folder containing the structured content in the library
	FolderID *string `json:"folder_id"`

	// ID Unique identifier of the structured content
	ID string `json:"id"`

	// IsArchived Whether the structured content is archived or not
	IsArchived bool `json:"is_archived"`

	// Labels Labels associated with the structured content
	Labels []ResourceLabelResponse `json:"labels"`

	// Links Meta links
	Links LibraryStructuredContentLinks `json:"links"`

	// ModifiedAt Date and time of the most recent modification of the structured content, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// OwnerOrganizationID ID of owner organization of the structured content
	OwnerOrganizationID string `json:"owner_organization_id"`

	// Title Title of the structured content
	Title string `json:"title"`
}

// LibraryStructuredContentLinks Meta links
type LibraryStructuredContentLinks struct {
	// Self URL of the structured content
	Self string `json:"self"`
}

// LibraryStructuredContentCreateRequest defines model for LibraryStructuredContentCreateRequest.
type LibraryStructuredContentCreateRequest struct {
	// ContentBody The structured content body
	ContentBody LibraryStructuredContentCreateRequestContentBody `json:"content_body"`

	// FolderID ID of the folder containing the structured content in the library
	FolderID *string `json:"folder_id"`

	// Title Title of the structured content asset
	Title *string `json:"title,omitempty"`
}

// LibraryStructuredContentCreateRequestContentBody The structured content body
type LibraryStructuredContentCreateRequestContentBody struct {
	// ContentTypeGUID The guid of the content's content type
	ContentTypeGUID string `json:"content_type_guid"`

	// Expired Expired status of the content
	Expired *bool `json:"expired,omitempty"`

	// ExpiryDatetime Date and time on which the content will expire, in ISO 8601 UTC format
	ExpiryDatetime *time.Time `json:"expiry_datetime"`

	// Fields The fields of structured content
	Fields *StructuredContentBody `json:"fields,omitempty"`

	// RootContent This is true when the content is not embedded in another content through a content type reference field
	RootContent bool `json:"root_content"`

	// Source Source of the content
	Source *string `json:"source"`

	// SourceID Source id of the content
	SourceID *string `json:"source_id"`

	// SourceMetadata Source metadata of the content
	SourceMetadata *string `json:"source_metadata"`

	// Title Title of the structured content
	Title *string `json:"title,omitempty"`
}

// LibraryStructuredContentUpdateRequest defines model for LibraryStructuredContentUpdateRequest.
type LibraryStructuredContentUpdateRequest struct {
	// ContentBody The structured content body
	ContentBody *LibraryStructuredContentUpdateRequestContentBody `json:"content_body,omitempty"`

	// IsArchived Whether the structured content should be archived
	IsArchived *bool `json:"is_archived,omitempty"`

	// Title Title of the structured content asset
	Title *string `json:"title"`
}

// LibraryStructuredContentUpdateRequestContentBody The structured content body
type LibraryStructuredContentUpdateRequestContentBody struct {
	// Expired Expired status of the content
	Expired *bool `json:"expired"`

	// ExpiryDatetime Date and time on which the content will expire, in ISO 8601 UTC format
	ExpiryDatetime *time.Time `json:"expiry_datetime"`

	// Fields The fields of structured content
	Fields *StructuredContentFields `json:"fields,omitempty"`

	// Source Source of the content
	Source *string `json:"source"`

	// SourceID Source id of the content
	SourceID *string `json:"source_id"`

	// SourceMetadata Source metadata of the content
	SourceMetadata *string `json:"source_metadata"`

	// Title Title of the structured content
	Title *string `json:"title"`
}

// LibraryVideo defines model for LibraryVideo.
type LibraryVideo struct {
	// AllowSeIndexing Whether the video is allowed to be indexed by search engines
	AllowSeIndexing bool `json:"allow_se_indexing"`

	// AltText Alternative text for video
	AltText *string `json:"alt_text"`

	// AttributionText Attribution of Video
	AttributionText *string `json:"attribution_text"`

	// CreatedAt Date and time on which the video was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// Description Description of the video
	Description *string `json:"description"`

	// ExpiresAt Date and time for the expiration of the video, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ExpiresAt *time.Time `json:"expires_at"`

	// FileExtension File extension of the video file
	FileExtension *string `json:"file_extension"`

	// FileLocation Location of the folder containing the video in the library
	FileLocation string `json:"file_location"`

	// FileSize Size of the video in bytes
	FileSize int `json:"file_size"`

	// FolderID ID of the folder containing the video in the library
	FolderID *string `json:"folder_id"`

	// ID Unique identifier for the video
	ID string `json:"id"`

	// IsArchived Whether the video is archived or not
	IsArchived bool `json:"is_archived"`

	// IsPublic Whether the video URL is public
	IsPublic bool `json:"is_public"`

	// Labels Labels associated with the video
	Labels []ResourceLabelResponse `json:"labels"`

	// MimeType MIME type of the video
	MimeType string `json:"mime_type"`

	// ModifiedAt Date and time of the most recent modification of the video, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// OwnerID ID of the video owner
	OwnerID string `json:"owner_id"`

	// OwnerOrganizationID ID of owner organization of the video
	OwnerOrganizationID string `json:"owner_organization_id"`

	// Tags Tags for the video
	Tags []Tag `json:"tags"`

	// ThumbnailURL URL of the thumbnail of the video
	ThumbnailURL *string `json:"thumbnail_url"`

	// Title Title of the video
	Title string `json:"title"`

	// URL Download the URL of the video
	URL string `json:"url"`

	// VersionID Version id of video
	VersionID string `json:"version_id"`

	// VersionNumber Version number of video
	VersionNumber int `json:"version_number"`
}

// LibraryVideoUpdateRequest defines model for LibraryVideoUpdateRequest.
type LibraryVideoUpdateRequest struct {
	// AllowSeIndexing Whether the video is allowed to be indexed by search engines
	AllowSeIndexing *bool `json:"allow_se_indexing,omitempty"`

	// AltText Alt text of the video. Set an empty string to clear the alt text.
	AltText *string `json:"alt_text,omitempty"`

	// AttributionText Attribution text of the video. Set an empty string to clear the attribution text.
	AttributionText *string `json:"attribution_text,omitempty"`

	// Description Description of the video. Set an empty string to clear the description.
	Description *string `json:"description,omitempty"`

	// ExpiresAt Date and time when the video expires, in ISO 8601 UTC format
	ExpiresAt *time.Time `json:"expires_at"`

	// FolderID ID of the folder containing the video in the library
	FolderID *string `json:"folder_id"`

	// IsArchived Whether the video should be archived
	IsArchived *bool `json:"is_archived,omitempty"`

	// IsPublic Whether the video URL should be public
	IsPublic *bool `json:"is_public,omitempty"`

	// Labels Please use [PUT /assets/{asset_id}/fields](https://docs.developers.optimizely.com/content-marketing-platform/reference/put_assets-asset-id-fields) or  [PUT /assets/{asset_id}/fields/{field_id}](https://docs.developers.optimizely.com/content-marketing-platform/reference/put_assets-asset-id-fields-field-id)  API to update label type asset fields
	// Deprecated: this property has been marked as deprecated upstream, but no `x-deprecated-reason` was set
	Labels *[]ResourceLabelRequest `json:"labels,omitempty"`

	// Tags Tags of the video. Provided tags will replace the existing ones. Set an empty array to clear the existing tags.
	Tags *[]string `json:"tags,omitempty"`

	// Title Title of the video
	Title *string `json:"title,omitempty"`
}

// ListAssetFieldsResponse defines model for ListAssetFieldsResponse.
type ListAssetFieldsResponse struct {
	// Data List of fields
	Data       []AssetFieldListResponseItem `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListAssetLineagesResponse defines model for ListAssetLineagesResponse.
type ListAssetLineagesResponse struct {
	// Data List of asset lineage
	Data       []AssetLineageResponse `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListAssetRenditionsResponse defines model for ListAssetRenditionsResponse.
type ListAssetRenditionsResponse struct {
	// Data List of renditions
	Data       []BaseAssetRenditionResponse `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListAssetsResponse defines model for ListAssetsResponse.
type ListAssetsResponse struct {
	// Data List of assets
	Data       []AssetResponse `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`

	// TotalCount Total number of assets found based on the request
	TotalCount float32 `json:"total_count"`
}

// ListBrandComplianceCategoriesResponse defines model for ListBrandComplianceCategoriesResponse.
type ListBrandComplianceCategoriesResponse struct {
	// Data List of brand compliance categories
	Data       []BrandComplianceCategoriesResponse `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListCampaignFieldsResponse defines model for ListCampaignFieldsResponse.
type ListCampaignFieldsResponse struct {
	// Data List of fields
	Data       []FieldListResponseItem `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListCampaignsResponse defines model for ListCampaignsResponse.
type ListCampaignsResponse struct {
	// Data List of campaigns
	Data       []CampaignListResponseItem `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListEventFieldsResponse defines model for ListEventFieldsResponse.
type ListEventFieldsResponse struct {
	// Data List of fields
	Data       []FieldListResponseItem `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListEventsResponse defines model for ListEventsResponse.
type ListEventsResponse struct {
	// Data List of events
	Data       []EventResponse `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListFieldsResponse defines model for ListFieldsResponse.
type ListFieldsResponse struct {
	// Data List of fields
	Data       []ListFieldsResponse_Data_Item `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListFieldsResponse_Data_Item defines model for ListFieldsResponse.data.Item.
type ListFieldsResponse_Data_Item struct {
	union json.RawMessage
}

// ListFoldersResponse defines model for ListFoldersResponse.
type ListFoldersResponse struct {
	// Data List of folders
	Data       []FolderResponse `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListLabelGroupsResponse defines model for ListLabelGroupsResponse.
type ListLabelGroupsResponse struct {
	// Data List of label groups
	Data       []LabelGroup `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListMilestonesResponse defines model for ListMilestonesResponse.
type ListMilestonesResponse struct {
	// Data List of milestones
	Data       []MilestoneResponse `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListSCContentTypeManagedMigrationsResponse defines model for ListSCContentTypeManagedMigrationsResponse.
type ListSCContentTypeManagedMigrationsResponse = []SCContentTypeManagedMigrationResponse

// ListSCContentTypeVersionsResponse defines model for ListSCContentTypeVersionsResponse.
type ListSCContentTypeVersionsResponse = []BaseContentTypeVersionModel

// ListSCContentTypesResponse defines model for ListSCContentTypesResponse.
type ListSCContentTypesResponse = []BaseContentTypeModel

// ListTaskAssetCommentsResponse defines model for ListTaskAssetCommentsResponse.
type ListTaskAssetCommentsResponse struct {
	// Data List of comments
	Data       *[]TaskAssetCommentResponse `json:"data,omitempty"`
	Pagination *struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination,omitempty"`
}

// ListTaskAssetDraftsResponse defines model for ListTaskAssetDraftsResponse.
type ListTaskAssetDraftsResponse struct {
	// Data List of drafts
	Data       []TaskAssetDraftListResponseItem `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListTaskAssetFieldsResponse defines model for ListTaskAssetFieldsResponse.
type ListTaskAssetFieldsResponse struct {
	// Data List of fields
	Data       *[]AssetFieldListResponseItem `json:"data,omitempty"`
	Pagination *struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination,omitempty"`
}

// ListTaskAssetsResponse defines model for ListTaskAssetsResponse.
type ListTaskAssetsResponse struct {
	// Data List of assets
	Data []struct {
		// Content Content of the asset
		Content AssetContent `json:"content"`

		// CreatedAt Date and time on which the asset was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
		CreatedAt time.Time `json:"created_at"`

		// ID Unique identifier for the asset
		ID string `json:"id"`

		// Labels Labels associated with the asset
		Labels []ResourceLabelResponse `json:"labels"`

		// LibraryAssetID Unique identifier of the corresponding library asset — the library asset this task asset was forked from, or the one it was released to.
		LibraryAssetID *string `json:"library_asset_id"`

		// Links Meta links
		Links TaskAssetResponseLinks `json:"links"`

		// MimeType MIME type of the asset
		MimeType string `json:"mime_type"`

		// ModifiedAt Date and time of the most recent modification of the asset, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
		ModifiedAt time.Time `json:"modified_at"`

		// Title Title of the asset
		Title string `json:"title"`

		// Type Type of the asset
		Type ListTaskAssetsResponseDataType `json:"type"`
	} `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// TaskAssetResponseLinksWebUrls Web URLs
type TaskAssetResponseLinksWebUrls struct {
	// Drafts Web URL of the drafts of the asset - null for `article` - non-null for `image`, `video`, and `raw file`.
	Drafts *string `json:"drafts"`

	// Self Web URL of the asset
	Self string `json:"self"`

	// Task Web URL of the task that the asset is associated with
	Task string `json:"task"`
}

// TaskAssetResponseLinks Meta links
type TaskAssetResponseLinks struct {
	// Drafts URL to POST draft to the asset - null for `article` - non-null for `image`, `video`, and `raw file`.
	Drafts *string `json:"drafts"`

	// Self URL of the asset
	Self string `json:"self"`

	// Task URL of the task that the asset is associated with
	Task string `json:"task"`

	// WebUrls Web URLs
	WebUrls TaskAssetResponseLinksWebUrls `json:"web_urls"`
}

// ListTaskAssetsResponseDataType Type of the asset
type ListTaskAssetsResponseDataType string

// ListTaskAttachmentsResponse defines model for ListTaskAttachmentsResponse.
type ListTaskAttachmentsResponse struct {
	// Data List of attachments
	Data       []AttachmentResponse `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListTaskCustomFieldChoicesResponse defines model for ListTaskCustomFieldChoicesResponse.
type ListTaskCustomFieldChoicesResponse struct {
	// Data List of custom field choices
	Data       []TaskCustomFieldChoiceListResponseItem `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListTaskCustomFieldsResponse defines model for ListTaskCustomFieldsResponse.
type ListTaskCustomFieldsResponse struct {
	// Data List of custom fields
	Data       []TaskCustomField `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListTaskFieldsResponse defines model for ListTaskFieldsResponse.
type ListTaskFieldsResponse struct {
	// Data List of fields
	Data       []FieldListResponseItem `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListTaskSubStepCommentsResponse defines model for ListTaskSubStepCommentsResponse.
type ListTaskSubStepCommentsResponse struct {
	// Data List of substep comments
	Data       []TaskSubStepCommentResponse `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListTaskSubStepFieldsResponse defines model for ListTaskSubStepFieldsResponse.
type ListTaskSubStepFieldsResponse struct {
	// Data List of fields
	Data       []ListTaskSubStepFieldsResponse_Data_Item `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListTaskSubStepFieldsResponse_Data_Item defines model for ListTaskSubStepFieldsResponse.data.Item.
type ListTaskSubStepFieldsResponse_Data_Item struct {
	union json.RawMessage
}

// ListTasksResponse defines model for ListTasksResponse.
type ListTasksResponse struct {
	// Data List of tasks
	Data       []TaskListResponseItem `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListTeamsResponse defines model for ListTeamsResponse.
type ListTeamsResponse struct {
	// Data List of teams
	Data       []Team `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListWorkRequestApprovedAssetsResponse defines model for ListWorkRequestApprovedAssetsResponse.
type ListWorkRequestApprovedAssetsResponse struct {
	// Data List of work request approved assets
	Data       []WorkRequestApprovedAssetResponse `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListWorkRequestCommentsResponse defines model for ListWorkRequestCommentsResponse.
type ListWorkRequestCommentsResponse struct {
	// Data List of work request comments
	Data       []WorkRequestCommentResponse `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListWorkRequestRelatedResourcesResponse defines model for ListWorkRequestRelatedResourcesResponse.
type ListWorkRequestRelatedResourcesResponse struct {
	// Data List of work request related resources
	Data       []WorkRequestRelatedResourceResponse `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListWorkRequestsResponse defines model for ListWorkRequestsResponse.
type ListWorkRequestsResponse struct {
	// Data List of work requests
	Data       []WorkRequestResponse `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// ListWorkflowsResponse defines model for ListWorkflowsResponse.
type ListWorkflowsResponse struct {
	// Data List of workflows
	Data       []WorkflowListResponseItem `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// LocalizedFieldValues defines model for LocalizedFieldValues.
type LocalizedFieldValues struct {
	FieldValues []LocalizedFieldValues_FieldValues_Item `json:"field_values"`
	Locale      string                                  `json:"locale"`
}

// LocalizedFieldValues_FieldValues_Item defines model for LocalizedFieldValues.field_values.Item.
type LocalizedFieldValues_FieldValues_Item struct {
	union json.RawMessage
}

// LocalizedFieldValuesWithEmbeddedContent defines model for LocalizedFieldValuesWithEmbeddedContent.
type LocalizedFieldValuesWithEmbeddedContent struct {
	FieldValues []LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item `json:"field_values"`
	Locale      string                                                     `json:"locale"`
}

// LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item defines model for LocalizedFieldValuesWithEmbeddedContent.field_values.Item.
type LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item struct {
	union json.RawMessage
}

// LocationDefaultValue defines model for LocationDefaultValue.
type LocationDefaultValue struct {
	// Latitude Latitude of the location
	Latitude float32 `json:"latitude"`

	// Longitude Longitude of the location
	Longitude float32 `json:"longitude"`
}

// LocationFieldValueModel defines model for LocationFieldValueModel.
type LocationFieldValueModel struct {
	// Latitude Longitude of the field value
	Latitude float32 `json:"latitude"`

	// Longitude Longitude of the field value
	Longitude float32 `json:"longitude"`

	// OrderIndex Order index of the field value
	OrderIndex *int `json:"order_index,omitempty"`
}

// LogicRule defines model for LogicRule.
type LogicRule struct {
	// Action Template logic rule action
	Action LogicRuleAction `json:"action"`

	// Condition Logic rule action associated condition
	Condition LogicRuleCondition `json:"condition"`
}

// LogicRuleAction Template logic rule action
type LogicRuleAction struct {
	// TargetField Template logic rule action's target field
	TargetField LogicRuleTargetField `json:"target_field"`

	// Type Type of logic rule action
	Type string `json:"type"`

	// Values Values to consider on logic rule action
	Values []string `json:"values"`
}

// LogicRuleCondition Logic rule action associated condition
type LogicRuleCondition struct {
	// Operator The operator for logic rule action
	Operator string `json:"operator"`

	// Values List of labels or custom field choices
	Values []string `json:"values"`
}

// LogicRuleTargetField Template logic rule action's target field
type LogicRuleTargetField struct {
	// CustomFieldType Type of custom field if the target field is a custom field
	CustomFieldType *string `json:"custom_field_type"`

	// Identifier Name of the target field identifier
	Identifier string `json:"identifier"`

	// Type Type of the target field identifier
	Type string `json:"type"`
}

// MigrateSCContentResponse defines model for MigrateSCContentResponse.
type MigrateSCContentResponse struct {
	// ContentGUID The identifier of the content
	ContentGUID *string `json:"content_guid,omitempty"`

	// ContentHash The hashed fingerprint of the content version
	ContentHash *string `json:"content_hash,omitempty"`

	// Created Whether the version was created or not
	Created *bool `json:"created,omitempty"`

	// VersionGUID The identifier of the content version
	VersionGUID *string `json:"version_guid,omitempty"`
}

// MilestoneCreateRequest defines model for MilestoneCreateRequest.
type MilestoneCreateRequest struct {
	// CampaignID Id of the campaign to be associated with the milestone
	CampaignID *string `json:"campaign_id,omitempty"`

	// Description Description of the milestone. If provided, this should be between 1 and 250 characters.
	Description *string `json:"description"`

	// DueDate Date and time on which the milestone will expire, in ISO 8601 UTC format
	DueDate time.Time `json:"due_date"`

	// HexColor Hex color code for the milestone label
	HexColor string `json:"hex_color"`

	// Tasks List of tasks to be associated with the milestone
	Tasks *[]MilestoneCreateRequestTasks `json:"tasks,omitempty"`

	// Title Title of the milestone. This should be between 1 and 80 characters.
	Title string `json:"title"`
}

// MilestoneCreateRequestTasks defines model for .
type MilestoneCreateRequestTasks struct {
	// ID Task ID
	ID string `json:"id"`
}

// MilestoneResponse defines model for MilestoneResponse.
type MilestoneResponse struct {
	// Campaign Campaign associated with the milestone
	Campaign *MilestoneResponseCampaign `json:"campaign"`

	// Color Color of the milestone label
	Color string `json:"color"`

	// Description Description of the milestone
	Description *string `json:"description"`

	// DueDate Date and time on which the milestone will expire, in ISO 8601 UTC format
	DueDate time.Time `json:"due_date"`

	// ID Unique identifier of the milestone
	ID string `json:"id"`

	// Links Meta links
	Links MilestoneResponseLinks `json:"links"`

	// Title Title of the milestone
	Title string `json:"title"`
}

// MilestoneResponseCampaign Campaign associated with the milestone
type MilestoneResponseCampaign struct {
	// ID Unique identifier of the campaign
	ID string `json:"id"`
}

// MilestoneResponseLinks Meta links
type MilestoneResponseLinks struct {
	// Self URL of the milestone
	Self string `json:"self"`
}

// MilestoneUpdateRequest defines model for MilestoneUpdateRequest.
type MilestoneUpdateRequest struct {
	// CampaignID Id of the campaign to be associated with the milestone
	CampaignID *string `json:"campaign_id,omitempty"`

	// Description Description of the milestone. If provided, this should be between 1 and 250 characters.
	Description *string `json:"description"`

	// DueDate Date and time on which the milestone will expire, in ISO 8601 UTC format
	DueDate *time.Time `json:"due_date,omitempty"`

	// HexColor Hex color code for the milestone label
	HexColor *string `json:"hex_color,omitempty"`

	// Tasks List of tasks to be associated with the milestone. This can be an empty array, but in that case it will remove all task associations for that milestone.
	Tasks *[]MilestoneUpdateRequestTasks `json:"tasks,omitempty"`

	// Title Title of the milestone. This should be between 1 and 80 characters.
	Title *string `json:"title,omitempty"`
}

// MilestoneUpdateRequestTasks defines model for .
type MilestoneUpdateRequestTasks struct {
	// ID Task ID
	ID string `json:"id"`
}

// MultiChoiceTypeFormFieldRequest defines model for MultiChoiceTypeFormFieldRequest.
type MultiChoiceTypeFormFieldRequest struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// Type Type of the form field
	Type string `json:"type"`

	// Values Array of choice IDs
	Values []string `json:"values"`
}

// MultiChoiceTypeFormFieldResponse defines model for MultiChoiceTypeFormFieldResponse.
type MultiChoiceTypeFormFieldResponse struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// IsProtected True when this form field is read-only and only organization admins can modify it. Absent or false for editable fields.
	IsProtected *bool `json:"is_protected,omitempty"`

	// Type Type of the form field
	Type string `json:"type"`

	// Values Array of choices
	Values []MultiChoiceTypeFormFieldResponseValues `json:"values"`
}

// MultiChoiceTypeFormFieldResponseValues defines model for .
type MultiChoiceTypeFormFieldResponseValues struct {
	// ID Unique identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// MultiChoiceTypeObjectField defines model for MultiChoiceTypeObjectField.
type MultiChoiceTypeObjectField struct {
	// ID Unique identifier of the field
	ID string `json:"id"`

	// Type Type of the field
	Type string `json:"type"`

	// Values Array of choice IDs
	Values []string `json:"values"`
}

// MultiChoiceTypeObjectFieldUpdatePayload defines model for MultiChoiceTypeObjectFieldUpdatePayload.
type MultiChoiceTypeObjectFieldUpdatePayload struct {
	// Type Type of the field
	Type string `json:"type"`

	// Values Array of choice IDs
	Values []string `json:"values"`
}

// MultipleOptionCustomField defines model for MultipleOptionCustomField.
type MultipleOptionCustomField struct {
	// Choices Custom field choices
	Choices []string `json:"choices"`

	// Label Name of the custom field
	Label string `json:"label"`
}

// NumberCustomField defines model for NumberCustomField.
type NumberCustomField struct {
	// DecimalPlaces Value of decimal places
	DecimalPlaces *int `json:"decimal_places"`

	// HasThousandSeparator Indicator of thousand separator
	HasThousandSeparator bool `json:"has_thousand_separator"`

	// Label Name of the custom field
	Label string `json:"label"`
}

// NumberFieldDefinition defines model for NumberFieldDefinition.
type NumberFieldDefinition struct {
	Core CoreFieldDef `json:"core"`

	// DefaultValues Default values of the field
	DefaultValues *[]float32 `json:"default_values,omitempty"`

	// MaxValue Maximum value of the field
	MaxValue *float32 `json:"max_value,omitempty"`

	// MinValue Minimum value of the field
	MinValue *float32 `json:"min_value,omitempty"`
}

// NumberFieldValueModel defines model for NumberFieldValueModel.
type NumberFieldValueModel struct {
	// NumValue Value of the field value
	NumValue float32 `json:"num_value"`

	// OrderIndex Order index of the field value
	OrderIndex *int `json:"order_index,omitempty"`
}

// NumberTypeFormFieldRequest defines model for NumberTypeFormFieldRequest.
type NumberTypeFormFieldRequest struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// Type Type of the form field
	Type string `json:"type"`

	// Values Array of numbers
	Values []float32 `json:"values"`
}

// NumberTypeFormFieldResponse defines model for NumberTypeFormFieldResponse.
type NumberTypeFormFieldResponse struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// IsProtected True when this form field is read-only and only organization admins can modify it. Absent or false for editable fields.
	IsProtected *bool `json:"is_protected,omitempty"`

	// Type Type of the form field
	Type string `json:"type"`

	// Values Array of numeric strings
	Values []string `json:"values"`
}

// NumberTypeObjectField defines model for NumberTypeObjectField.
type NumberTypeObjectField struct {
	// ID Unique identifier of the field
	ID string `json:"id"`

	// Type Type of the field
	Type string `json:"type"`

	// Values Accepts a single numeric value
	Values []float32 `json:"values"`
}

// NumberTypeObjectFieldUpdatePayload defines model for NumberTypeObjectFieldUpdatePayload.
type NumberTypeObjectFieldUpdatePayload struct {
	// Type Type of the field
	Type string `json:"type"`

	// Values Accepts a single numeric value
	Values []float32 `json:"values"`
}

// ObjectFieldCreateRequest defines model for ObjectFieldCreateRequest.
type ObjectFieldCreateRequest struct {
	union json.RawMessage
}

// ObjectFieldCreateResponse defines model for ObjectFieldCreateResponse.
type ObjectFieldCreateResponse struct {
	union json.RawMessage
}

// Pagination Pagination related information
type Pagination struct {
	// Next URL to the next page
	Next *string `json:"next"`

	// Previous URL to the previous page
	Previous *string `json:"previous"`
}

// PatchLocalizedFieldValuesWithEmbeddedContent defines model for PatchLocalizedFieldValuesWithEmbeddedContent.
type PatchLocalizedFieldValuesWithEmbeddedContent struct {
	FieldValues []PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item `json:"field_values"`

	// Locale The default value should be en_US
	Locale string `json:"locale"`
}

// PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item defines model for PatchLocalizedFieldValuesWithEmbeddedContent.field_values.Item.
type PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item struct {
	union json.RawMessage
}

// PercentageCustomField defines model for PercentageCustomField.
type PercentageCustomField struct {
	// DecimalPlaces Value of decimal places
	DecimalPlaces *int `json:"decimal_places"`

	// Label Name of the custom field
	Label string `json:"label"`
}

// PublishingChannelListResponse defines model for PublishingChannelListResponse.
type PublishingChannelListResponse struct {
	// Data List of publishing channels
	Data       []PublishingChannelResponse `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// PublishingChannelResponse defines model for PublishingChannelResponse.
type PublishingChannelResponse struct {
	// Disabled Whether the publishing channel is currently disabled
	Disabled *bool `json:"disabled"`

	// ID Unique identifier of the publishing channel
	ID string `json:"id"`

	// Name Human-readable name of the publishing channel
	Name *string `json:"name"`
}

// PublishingEventAssets defines model for PublishingEventAssets.
type PublishingEventAssets struct {
	// ID Unique identifier of the asset
	ID string `json:"id"`

	// Links Meta links
	Links PublishingEventAssetsLinks `json:"links"`

	// PublishingMetadata List of information related to publishing metadata for the asset
	PublishingMetadata []PublishingEventAssetsPublishingMetadata `json:"publishing_metadata"`

	// Type Type of the asset: `article`, `image`, `video`, `raw_file`, `structured_content`
	Type string `json:"type"`
}

// PublishingEventAssetsLinks Meta links
type PublishingEventAssetsLinks struct {
	// Self URL of asset
	Self string `json:"self"`
}

// PublishingEventAssetsPublishingMetadataLinks Meta links
type PublishingEventAssetsPublishingMetadataLinks struct {
	// Self URL of the publishing metadata of the asset
	Self string `json:"self"`
}

// PublishingEventAssetsPublishingMetadata defines model for .
type PublishingEventAssetsPublishingMetadata struct {
	// ID Unique identifier of the publishing metadata
	ID string `json:"id"`

	// Links Meta links
	Links PublishingEventAssetsPublishingMetadataLinks `json:"links"`
}

// PublishingEventMetadataBulkCreateRequest defines model for PublishingEventMetadataBulkCreateRequest.
type PublishingEventMetadataBulkCreateRequest struct {
	// Data List publishing metadata to be posted
	Data []PublishingEventMetadataCreateRequest `json:"data"`
}

// PublishingEventMetadataBulkCreateResponse defines model for PublishingEventMetadataBulkCreateResponse.
type PublishingEventMetadataBulkCreateResponse struct {
	// Data List of successfully posted publishing metadata
	Data []PublishingEventMetadataResponse `json:"data"`

	// Errors List of errors of assets for which publishing metadata failed to post
	Errors []PublishingEventMetadataBulkCreateResponseErrors `json:"errors"`
}

// PublishingEventMetadataBulkCreateResponseErrorsErrorCode Custom error code
type PublishingEventMetadataBulkCreateResponseErrorsErrorCode string

// PublishingEventMetadataBulkCreateResponseErrors defines model for .
type PublishingEventMetadataBulkCreateResponseErrors struct {
	// AssetID Unique identifier of the asset that failed to update its publishing metadata
	AssetID string `json:"asset_id"`

	// ErrorCode Custom error code
	ErrorCode PublishingEventMetadataBulkCreateResponseErrorsErrorCode `json:"error_code"`

	// Locale locale of the asset
	Locale *string `json:"locale"`

	// Message Description of the error
	Message string `json:"message"`
}

// PublishingEventMetadataCreateRequest defines model for PublishingEventMetadataCreateRequest.
type PublishingEventMetadataCreateRequest struct {
	// AssetID Unique identifier of the asset
	AssetID string `json:"asset_id"`

	// Locale The locale to which the asset is being published
	Locale *string `json:"locale,omitempty"`

	// PublicURL public url of asset
	PublicURL *string `json:"public_url,omitempty"`

	// PublishingDestinationUpdatedAt Timestamp of when the publishing destination of the asset was updated
	PublishingDestinationUpdatedAt *time.Time `json:"publishing_destination_updated_at,omitempty"`

	// Status Publishing status of the asset
	Status PublishingEventMetadataCreateRequestStatus `json:"status"`

	// StatusMessage Any message patched by the integrator about the asset's status
	StatusMessage *string `json:"status_message,omitempty"`
}

// PublishingEventMetadataCreateRequestStatus Publishing status of the asset
type PublishingEventMetadataCreateRequestStatus string

// PublishingEventMetadataListResponse defines model for PublishingEventMetadataListResponse.
type PublishingEventMetadataListResponse struct {
	// Data List of successfully posted publishing metadata
	Data *[]PublishingEventMetadataResponse `json:"data,omitempty"`
}

// PublishingEventMetadataResponse defines model for PublishingEventMetadataResponse.
type PublishingEventMetadataResponse struct {
	// AssetID Unique identifier of asset
	AssetID string `json:"asset_id"`

	// AssetType Type of the asset
	AssetType PublishingEventMetadataResponseAssetType `json:"asset_type"`

	// ID Unique identifier of publishing metadata
	ID string `json:"id"`

	// Links Meta links
	Links PublishingEventMetadataResponseLinks `json:"links"`

	// Locale The locale to which the asset is being published.
	Locale *string `json:"locale"`

	// PublicURL Public URL of the asset
	PublicURL *string `json:"public_url"`

	// PublishingDestinationUpdatedAt Timestamp of when the publishing destination of the asset was updated
	PublishingDestinationUpdatedAt *time.Time `json:"publishing_destination_updated_at"`

	// Status Publishing status of the asset
	Status *PublishingEventMetadataResponseStatus `json:"status"`

	// StatusMessage Any message patched by the integrator about the asset's status
	StatusMessage *string `json:"status_message"`
}

// PublishingEventMetadataResponseAssetType Type of the asset
type PublishingEventMetadataResponseAssetType string

// PublishingEventMetadataResponseLinks Meta links
type PublishingEventMetadataResponseLinks struct {
	// Asset URL of the asset
	Asset string `json:"asset"`

	// PublishingEvent URL of the publishing event
	PublishingEvent string `json:"publishing_event"`

	// Self URL of the publishing metadata
	Self string `json:"self"`
}

// PublishingEventMetadataResponseStatus Publishing status of the asset
type PublishingEventMetadataResponseStatus string

// PublishingEventResponse defines model for PublishingEventResponse.
type PublishingEventResponse struct {
	// Assets List of assets associated with the publishing event.
	Assets []PublishingEventAssets `json:"assets"`

	// ID Unique identifier of the publishing event
	ID string `json:"id"`

	// Links Meta links
	Links PublishingEventResponseLinks `json:"links"`
}

// PublishingEventResponseLinks Meta links
type PublishingEventResponseLinks struct {
	// PublishingMetadata URL of list of publishing metadata
	PublishingMetadata string `json:"publishing_metadata"`

	// Self URL of the publishing event
	Self string `json:"self"`
}

// RadioButtonTypeFormFieldRequest defines model for RadioButtonTypeFormFieldRequest.
type RadioButtonTypeFormFieldRequest struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// Type Type of the form field
	Type string `json:"type"`

	// Values Array of choice IDs
	Values []string `json:"values"`
}

// RadioButtonTypeFormFieldResponse defines model for RadioButtonTypeFormFieldResponse.
type RadioButtonTypeFormFieldResponse struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// IsProtected True when this form field is read-only and only organization admins can modify it. Absent or false for editable fields.
	IsProtected *bool `json:"is_protected,omitempty"`

	// Type Type of the form field
	Type string `json:"type"`

	// Values Array of choices
	Values []RadioButtonTypeFormFieldResponseValues `json:"values"`
}

// RadioButtonTypeFormFieldResponseValues defines model for .
type RadioButtonTypeFormFieldResponseValues struct {
	// ID Unique identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// RadioButtonTypeObjectField defines model for RadioButtonTypeObjectField.
type RadioButtonTypeObjectField struct {
	// ID Unique identifier of the field
	ID string `json:"id"`

	// Type Type of the field
	Type string `json:"type"`

	// Values Accepts a single choice Id.
	Values []string `json:"values"`
}

// RadioButtonTypeObjectFieldUpdatePayload defines model for RadioButtonTypeObjectFieldUpdatePayload.
type RadioButtonTypeObjectFieldUpdatePayload struct {
	// Type Type of the field
	Type string `json:"type"`

	// Values Accepts a single choice Id.
	Values []string `json:"values"`
}

// RecursivePatchStructuredContentFields defines model for RecursivePatchStructuredContentFields.
type RecursivePatchStructuredContentFields struct {
	// ContentBody The fields of structured content
	ContentBody *StructuredContentFields `json:"content_body,omitempty"`

	// Expired Expired status of the content
	Expired *bool `json:"expired,omitempty"`

	// ExpiryDatetime Date and time on which the content will expire, in ISO 8601 UTC format
	ExpiryDatetime *time.Time `json:"expiry_datetime"`

	// RootContent This is true when the content is not embedded in another content through a content type reference field
	RootContent *bool `json:"root_content,omitempty"`

	// Source Source of the content
	Source *string `json:"source"`

	// SourceID Source id of the content
	SourceID       *string `json:"source_id"`
	SourceMetadata *string `json:"source_metadata,omitempty"`

	// Title Title of the structured content
	Title *string `json:"title,omitempty"`
}

// RecursiveStructuredContent defines model for RecursiveStructuredContent.
type RecursiveStructuredContent struct {
	// ContentTypeGUID The guid of the content's content type
	ContentTypeGUID string `json:"content_type_guid"`

	// Expired Expired status of the content
	Expired *bool `json:"expired,omitempty"`

	// ExpiryDatetime Date and time on which the content will expire, in ISO 8601 UTC format
	ExpiryDatetime *time.Time `json:"expiry_datetime"`

	// RootContent This is true when the content is not embedded in another content through a content type reference field
	RootContent bool `json:"root_content"`

	// Source Source of the content
	Source *string `json:"source"`

	// SourceID Source id of the content
	SourceID *string `json:"source_id"`

	// SourceMetadata Source metadata of the content
	SourceMetadata *string `json:"source_metadata"`

	// Title Title of the structured content
	Title                *string                                              `json:"title,omitempty"`
	AdditionalProperties map[string][]LocalizedFieldValuesWithEmbeddedContent `json:"-"`
}

// RelatedAssetItem defines model for RelatedAssetItem.
type RelatedAssetItem struct {
	Content RelatedAssetItemContent `json:"content"`

	// FileExtension File extension of the asset
	FileExtension *string `json:"file_extension"`

	// ID Unique identifier for the asset
	ID string `json:"id"`

	// Links Meta links
	Links RelatedAssetItemLinks `json:"links"`

	// MimeType MIME type of the asset
	MimeType *string `json:"mime_type"`

	// Title Title of the asset
	Title string `json:"title"`

	// Type Type of the asset
	Type RelatedAssetItemType `json:"type"`
}

// RelatedAssetItemContentType Type of the content.
//   - article  – `html_body`
//   - image, raw_file, video  – `url`
//   - structured_content  – `api_url`
type RelatedAssetItemContentType string

// RelatedAssetItemContent defines model for .
type RelatedAssetItemContent struct {
	// Type Type of the content.
	//  - article  – `html_body`
	//  - image, raw_file, video  – `url`
	//  - structured_content  – `api_url`
	Type RelatedAssetItemContentType `json:"type"`

	// Value Content of the asset. - article – the html body - image, raw_file, video – the download URL - structured_content – api url
	Value string `json:"value"`
}

// RelatedAssetItemLinks Meta links
type RelatedAssetItemLinks struct {
	// Self URL of the asset (only GET is supported)
	Self string `json:"self"`
}

// RelatedAssetItemType Type of the asset
type RelatedAssetItemType string

// RelatedAssetsListResponse defines model for RelatedAssetsListResponse.
type RelatedAssetsListResponse struct {
	// Data List of related assets
	Data       []RelatedAssetItem `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// RenditionConfigResponse defines model for RenditionConfigResponse.
type RenditionConfigResponse struct {
	// AssetType The asset type this config applies to
	AssetType string `json:"asset_type"`

	// Height The expected height (in pixels) of the renditions generated using this config
	Height *int `json:"height"`

	// ID Unique identifier of the rendition config
	ID string `json:"id"`

	// InputFormat The file format of the target asset
	InputFormat *string `json:"input_format"`

	// Name Name of the rendition config
	Name string `json:"name"`

	// OutputFormat The file format of the generated rendition
	OutputFormat *string `json:"output_format"`

	// Quality The quality (percent) of the renditions generated using this config
	Quality *float32 `json:"quality"`

	// Width The expected width (in pixels) of the renditions generated using this config
	Width *int `json:"width"`
}

// ReplaceRelatedAssetsRequest defines model for ReplaceRelatedAssetsRequest.
type ReplaceRelatedAssetsRequest struct {
	RelatedAssets []ReplaceRelatedAssetsRequestRelatedAssets `json:"related_assets"`
}

// ReplaceRelatedAssetsRequestRelatedAssets defines model for .
type ReplaceRelatedAssetsRequestRelatedAssets struct {
	// ID Unique identifier for the asset
	ID string `json:"id"`
}

// ReplaceRelatedAssetsResponse defines model for ReplaceRelatedAssetsResponse.
type ReplaceRelatedAssetsResponse struct {
	// ID Unique identifier for the asset
	ID            string                                      `json:"id"`
	RelatedAssets []ReplaceRelatedAssetsResponseRelatedAssets `json:"related_assets"`
}

// ReplaceRelatedAssetsResponseRelatedAssets defines model for .
type ReplaceRelatedAssetsResponseRelatedAssets struct {
	// ID Unique identifier for the asset
	ID string `json:"id"`
}

// ResourceChoiceResponse defines model for ResourceChoiceResponse.
type ResourceChoiceResponse struct {
	// ID Unique identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// ResourceLabelRequest Payload to associate labels to a resource
type ResourceLabelRequest struct {
	// Group Label group
	Group string `json:"group"`

	// Values List of values of the label
	Values []string `json:"values"`
}

// ResourceLabelResponse Label associated to a fetched resource
type ResourceLabelResponse struct {
	// Group Label group
	Group ResourceLabelResponseGroup `json:"group"`

	// Values List of values of the label
	Values []ResourceLabelResponseValues `json:"values"`
}

// ResourceLabelResponseGroup Label group
type ResourceLabelResponseGroup struct {
	// ID Identifier of the label group
	ID string `json:"id"`

	// Name Name of the label group
	Name string `json:"name"`
}

// ResourceLabelResponseValues defines model for .
type ResourceLabelResponseValues struct {
	// ID Identifier of the label value
	ID string `json:"id"`

	// Name Name of the label value
	Name string `json:"name"`
}

// RichTextFieldDefinition defines model for RichTextFieldDefinition.
type RichTextFieldDefinition struct {
	Core CoreFieldDef `json:"core"`

	// DefaultValues Default values of the field
	DefaultValues *[]string `json:"default_values,omitempty"`

	// MaxVisualTextLength Maximum length of the field
	MaxVisualTextLength *int `json:"max_visual_text_length,omitempty"`

	// MinVisualTextLength Minimum length of the field
	MinVisualTextLength int `json:"min_visual_text_length"`
}

// RichTextFieldValueModel defines model for RichTextFieldValueModel.
type RichTextFieldValueModel struct {
	// OrderIndex Order index of the field value
	OrderIndex *int `json:"order_index,omitempty"`

	// RichTextValue Value of the field value
	RichTextValue string `json:"rich_text_value"`
}

// SCContentMigrationCreateRequest defines model for SCContentMigrationCreateRequest.
type SCContentMigrationCreateRequest struct {
	// CreatedBy Unique identifier of the user who migrated the content
	CreatedBy string `json:"created_by"`

	// Fields List of fields
	Fields *map[string][]LocalizedFieldValues `json:"fields,omitempty"`

	// NewContentTypeVersionID Unique identifier of the content type version to migrate the content to
	NewContentTypeVersionID *string `json:"new_content_type_version_id,omitempty"`

	// Source Source of the content
	Source *string `json:"source"`

	// SourceID Source id of the content
	SourceID *string `json:"source_id"`
}

// SCContentPreviewAcknowledgeRequest defines model for SCContentPreviewAcknowledgeRequest.
type SCContentPreviewAcknowledgeRequest struct {
	// AcknowledgedBy Unique identifier of the user who acknowledged the preview request
	AcknowledgedBy string `json:"acknowledged_by"`

	// ContentHash Content hash of the preview request
	ContentHash string `json:"content_hash"`
}

// SCContentPreviewCompleteRequest defines model for SCContentPreviewCompleteRequest.
type SCContentPreviewCompleteRequest struct {
	KeyedPreviews map[string]SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties `json:"keyed_previews"`
}

// SCContentPreviewCompleteRequestKeyedPreviews0 defines model for .
type SCContentPreviewCompleteRequestKeyedPreviews0 = string

// SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties defines model for SCContentPreviewCompleteRequest.keyed_previews.AdditionalProperties.
type SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties struct {
	union json.RawMessage
}

// SCContentType defines model for SCContentType.
type SCContentType struct {
	// Component Status determining whether the content type is a component
	Component bool `json:"component"`

	// ContentTypeGUID Unique identifier of the content type
	ContentTypeGUID string `json:"content_type_guid"`

	// CreatedAt Date and time on which the content type was created, in ISO 8601 UTC format
	CreatedAt time.Time `json:"created_at"`

	// CreatedBy Unique identifier of the user who created the content type
	CreatedBy string `json:"created_by"`

	// Description Description of the content type
	Description *string `json:"description,omitempty"`

	// Disabled Disabled status of the content type
	Disabled      *bool                `json:"disabled,omitempty"`
	LatestVersion SCContentTypeVersion `json:"latest_version"`

	// Links Meta Links
	Links BaseContentTypeModelLinks `json:"links"`

	// Name Name of the content type
	Name string `json:"name"`

	// Source Source of the content type
	Source *string `json:"source,omitempty"`

	// SourceID Source of the content type
	SourceID *string `json:"source_id,omitempty"`

	// SourceMetadata Source metadata of the content type
	SourceMetadata *string `json:"source_metadata,omitempty"`

	// ThumbnailGUID Thumbnail GUID of the content type
	ThumbnailGUID *string `json:"thumbnail_guid,omitempty"`

	// UpdatedAt Date and time on which the content type was last updated, in ISO 8601 UTC format
	UpdatedAt time.Time `json:"updated_at"`

	// UpdatedBy Unique identifier of the user who last updated the content type
	UpdatedBy string `json:"updated_by"`
}

// SCContentTypeCreateRequest defines model for SCContentTypeCreateRequest.
type SCContentTypeCreateRequest struct {
	// CreatedBy Unique identifier of the user who is creating the content type
	CreatedBy string          `json:"created_by"`
	Details   CoreContentType `json:"details"`

	// ExpectedLocales Expected locales for the content type
	ExpectedLocales  *[]string                                          `json:"expected_locales,omitempty"`
	FieldDefinitions []SCContentTypeCreateRequest_FieldDefinitions_Item `json:"field_definitions"`

	// Source Source for the content type
	Source *string `json:"source,omitempty"`

	// SourceID Source ID for the content type
	SourceID *string `json:"source_id,omitempty"`

	// SourceMetadata Source metadata for the content type
	SourceMetadata *string `json:"source_metadata,omitempty"`
}

// SCContentTypeCreateRequest_FieldDefinitions_Item defines model for SCContentTypeCreateRequest.field_definitions.Item.
type SCContentTypeCreateRequest_FieldDefinitions_Item struct {
	union json.RawMessage
}

// SCContentTypeCreateResponse defines model for SCContentTypeCreateResponse.
type SCContentTypeCreateResponse struct {
	// ContentTypeGUID Unique identifier of the content type
	ContentTypeGUID *string `json:"content_type_guid,omitempty"`

	// ContentTypeVersionGUID Unique identifier of the content type version
	ContentTypeVersionGUID *string `json:"content_type_version_guid,omitempty"`

	// Created Created status of the requested content type
	Created bool `json:"created"`
}

// SCContentTypeManagedMigrationCreateRequest defines model for SCContentTypeManagedMigrationCreateRequest.
type SCContentTypeManagedMigrationCreateRequest struct {
	// CreatedBy Unique identifier of the user who created the job.
	CreatedBy     string                `json:"created_by"`
	DefaultValues *LocalizedFieldValues `json:"default_values,omitempty"`

	// SourceContentTypeVersionID The ID of the source content type version to migrate from.
	SourceContentTypeVersionID string `json:"source_content_type_version_id"`
}

// SCContentTypeManagedMigrationResponse defines model for SCContentTypeManagedMigrationResponse.
type SCContentTypeManagedMigrationResponse struct {
	ContentMigrationSummary *ContentMigrationSummary `json:"content_migration_summary,omitempty"`

	// ContentTypeID The ID of the content type.
	ContentTypeID *string `json:"content_type_id,omitempty"`

	// CreatedAt Date and time on which the managed migration job was created,  in ISO 8601 UTC format (2020-02-10T10:40:45Z).
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// DefaultValues The default values for the content type.
	DefaultValues *map[string]interface{} `json:"default_values,omitempty"`

	// ID The ID of the managed migration job.
	ID *string `json:"id,omitempty"`

	// InstanceID The ID of the instance.
	InstanceID *string `json:"instance_id,omitempty"`

	// SourceContentTypeVersionID The ID of the source content type version.
	SourceContentTypeVersionID *string `json:"source_content_type_version_id,omitempty"`

	// Status The status of the managed migration job.
	Status *SCContentTypeManagedMigrationResponseStatus `json:"status,omitempty"`

	// TargetContentTypeVersionID The ID of the target content type version.
	TargetContentTypeVersionID *string `json:"target_content_type_version_id,omitempty"`

	// UpdatedAt Date and time on which the managed migration job was last updated,  in ISO 8601 UTC format (2020-02-10T10:40:45Z).
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// SCContentTypeManagedMigrationResponseStatus The status of the managed migration job.
type SCContentTypeManagedMigrationResponseStatus string

// SCContentTypeManagedMigrationStartResponse defines model for SCContentTypeManagedMigrationStartResponse.
type SCContentTypeManagedMigrationStartResponse struct {
	// JobID The unique identifier of job.
	JobID *string `json:"job_id,omitempty"`

	// Started True if job started, otherwise false.
	Started *bool `json:"started,omitempty"`
}

// SCContentTypeManagedMigrationValidateRequest defines model for SCContentTypeManagedMigrationValidateRequest.
type SCContentTypeManagedMigrationValidateRequest struct {
	DefaultValues *LocalizedFieldValues `json:"default_values,omitempty"`

	// SourceContentTypeVersionID The ID of the source content type version to migrate from.
	SourceContentTypeVersionID string `json:"source_content_type_version_id"`
}

// SCContentTypeUpdateRequest defines model for SCContentTypeUpdateRequest.
type SCContentTypeUpdateRequest struct {
	Details CoreContentType `json:"details"`

	// SourceMetadata Source metadata of the content type
	SourceMetadata *string `json:"source_metadata,omitempty"`

	// UpdatedBy Unique identifier of the user who updated the content type
	UpdatedBy string `json:"updated_by"`
}

// SCContentTypeUpdateResponse defines model for SCContentTypeUpdateResponse.
type SCContentTypeUpdateResponse struct {
	// Updated Updated status of the request
	Updated bool `json:"updated"`
}

// SCContentTypeVersion defines model for SCContentTypeVersion.
type SCContentTypeVersion struct {
	// CreatedAt Date and time on which the version was created, in ISO 8601 UTC format
	CreatedAt time.Time `json:"created_at"`

	// CreatedBy Unique identifier of the user who created the version
	CreatedBy string `json:"created_by"`

	// ExpectedLocales Expected locales of the version
	ExpectedLocales  []string                                     `json:"expected_locales"`
	FieldDefinitions []SCContentTypeVersion_FieldDefinitions_Item `json:"field_definitions"`

	// Latest Indicates whether the version is the latest
	Latest bool `json:"latest"`

	// Links Meta links
	Links BaseContentTypeVersionModelLinks `json:"links"`

	// VersionGUID Unique identifier of the version
	VersionGUID string `json:"version_guid"`
}

// SCContentTypeVersion_FieldDefinitions_Item defines model for SCContentTypeVersion.field_definitions.Item.
type SCContentTypeVersion_FieldDefinitions_Item struct {
	union json.RawMessage
}

// SCContentTypeVersionCreateRequest defines model for SCContentTypeVersionCreateRequest.
type SCContentTypeVersionCreateRequest struct {
	// CreatedBy Unique identifier of the user who created the version
	CreatedBy string `json:"created_by"`

	// ExpectedLocales Expected locales of the version
	ExpectedLocales  *[]string                                                 `json:"expected_locales,omitempty"`
	FieldDefinitions []SCContentTypeVersionCreateRequest_FieldDefinitions_Item `json:"field_definitions"`
}

// SCContentTypeVersionCreateRequest_FieldDefinitions_Item defines model for SCContentTypeVersionCreateRequest.field_definitions.Item.
type SCContentTypeVersionCreateRequest_FieldDefinitions_Item struct {
	union json.RawMessage
}

// Section defines model for Section.
type Section struct {
	// ID Id of the template section
	ID string `json:"id"`

	// Label Name of the template section
	Label string `json:"label"`
}

// Settings defines model for Settings.
type Settings struct {
	Resources SettingsResources `json:"resources"`
}

// SettingsApp defines model for SettingsApp.
type SettingsApp struct {
	// AuthorizationCallbackUrls List of the authorization callback URLs
	AuthorizationCallbackUrls []string `json:"authorization_callback_urls"`

	// Description Description of the app
	Description string `json:"description"`

	// Env Environment of the app
	Env string `json:"env"`

	// ExposeEmail Indicates whether an email should be exposed in an API response
	ExposeEmail bool `json:"expose_email"`

	// HomepageURL Home page URL for app
	HomepageURL string `json:"homepage_url"`

	// Name Name of the app
	Name string `json:"name"`
}

// SettingsChangeset defines model for SettingsChangeset.
type SettingsChangeset struct {
	Create SettingsChangesetBase `json:"create"`
	Update SettingsChangesetBase `json:"update"`
}

// SettingsChangesetBase defines model for SettingsChangesetBase.
type SettingsChangesetBase struct {
	// Apps List of the app names
	Apps []string `json:"apps"`

	// CustomFields List of the custom field names
	CustomFields []string `json:"custom_fields"`

	// Labels List of the label names
	Labels []string `json:"labels"`

	// RoutingRules List of the routing rule names
	RoutingRules []string `json:"routing_rules"`

	// Templates List of the template names
	Templates []string `json:"templates"`

	// Webhooks List of the webhook names
	Webhooks []string `json:"webhooks"`

	// Workflows List of the workflow names
	Workflows []string `json:"workflows"`
}

// SettingsCustomField Settings for custom fields
type SettingsCustomField struct {
	// Checkbox List of checkbox custom fields
	Checkbox []MultipleOptionCustomField `json:"checkbox"`

	// Currency List of currency custom fields
	Currency []CurrencyCustomField `json:"currency"`

	// Date List of date custom fields
	Date []LabelOnlyCustomField `json:"date"`

	// Dropdown List of dropdown custom fields
	Dropdown []MultipleOptionCustomField `json:"dropdown"`

	// Image List of image custom fields
	Image []LabelOnlyCustomField `json:"image"`

	// MultiSelectDropdown List of multi-select dropdown custom fields
	MultiSelectDropdown []MultipleOptionCustomField `json:"multi_select_dropdown"`

	// Multichoice List of multichoice custom fields
	Multichoice []MultipleOptionCustomField `json:"multichoice"`

	// Number List of number custom fields
	Number []NumberCustomField `json:"number"`

	// Percentage List of percentage custom fields
	Percentage []PercentageCustomField `json:"percentage"`

	// Richtext List of rich text custom fields
	Richtext []LabelOnlyCustomField `json:"richtext"`

	// String List of string custom fields
	String []LabelOnlyCustomField `json:"string"`

	// Textarea List of textarea custom fields
	Textarea []LabelOnlyCustomField `json:"textarea"`

	// Video List of video custom fields
	Video []LabelOnlyCustomField `json:"video"`
}

// SettingsFieldChoiceCreateRequest Choices of the field
type SettingsFieldChoiceCreateRequest = []SettingsFieldChoiceCreateRequest_Item

// SettingsFieldChoiceCreateRequest_Item defines model for SettingsFieldChoiceCreateRequest.Item.
type SettingsFieldChoiceCreateRequest_Item struct {
	union json.RawMessage
}

// SettingsFieldChoiceCreateResponse Choices of the field
type SettingsFieldChoiceCreateResponse = []SettingsFieldChoiceCreateResponseItem

// SettingsFieldChoiceCreateResponseItem defines model for .
type SettingsFieldChoiceCreateResponseItem struct {
	// ID Identifier of the field choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// SettingsFieldChoiceUpdateRequest Payload for updating choice
type SettingsFieldChoiceUpdateRequest struct {
	union json.RawMessage
}

// SettingsFieldCreateRequest defines model for SettingsFieldCreateRequest.
type SettingsFieldCreateRequest struct {
	union json.RawMessage
}

// SettingsFieldUpdateRequest defines model for SettingsFieldUpdateRequest.
type SettingsFieldUpdateRequest struct {
	union json.RawMessage
}

// SettingsLabelGroup defines model for SettingsLabelGroup.
type SettingsLabelGroup struct {
	// HasSingleValue Single or multiple value support indicator for the label group
	HasSingleValue bool `json:"has_single_value"`

	// LabelType Name of the label group
	LabelType string `json:"label_type"`

	// Labels List of labels of the label group
	Labels []SettingsLabelGroupLabel `json:"labels"`
}

// SettingsLabelGroupLabel defines model for SettingsLabelGroupLabel.
type SettingsLabelGroupLabel struct {
	// Color Color of the label
	Color *string `json:"color"`

	// Name Name of the label
	Name string `json:"name"`
}

// SettingsResources defines model for SettingsResources.
type SettingsResources struct {
	// Apps Settings apps
	Apps []SettingsApp `json:"apps"`

	// CustomFields Settings for custom fields
	CustomFields SettingsCustomField `json:"custom_fields"`

	// Labels Settings labels
	Labels []SettingsLabelGroup `json:"labels"`

	// RoutingRules Settings routing rules
	RoutingRules []SettingsRoutingRule `json:"routing_rules"`

	// Templates Settings templates
	Templates []SettingsTemplate `json:"templates"`

	// Webhooks Settings webhooks
	Webhooks []SettingsWebhook `json:"webhooks"`

	// Workflows Settings workflows
	Workflows []SettingsWorkflow `json:"workflows"`
}

// SettingsRoutingRule defines model for SettingsRoutingRule.
type SettingsRoutingRule struct {
	// Description Description of the routing rule
	Description string `json:"description"`

	// Name Name of the routing rule
	Name string `json:"name"`

	// Rules List of the associated rules
	Rules []SettingsRoutingRuleRules `json:"rules"`

	// TemplateName Name of the associated template
	TemplateName *string `json:"template_name"`

	// Watchers List of the routing rule watchers email
	Watchers []string `json:"watchers"`
}

// SettingsRoutingRuleRulesUnit Unit of rule's field. Not null when the rule's `field_type` value is `dueDate`
type SettingsRoutingRuleRulesUnit string

// SettingsRoutingRuleRules defines model for .
type SettingsRoutingRuleRules struct {
	// CustomFieldType Associated custom field type
	CustomFieldType *string `json:"custom_field_type"`

	// FieldType Name of the rule field type
	FieldType string `json:"field_type"`

	// Identifier Name of the rule identifier
	Identifier *string `json:"identifier"`

	// Operator Name of the rule operator
	Operator string `json:"operator"`

	// Unit Unit of rule's field. Not null when the rule's `field_type` value is `dueDate`
	Unit *SettingsRoutingRuleRulesUnit `json:"unit"`

	// Value List of the associated label, custom field, or due date choice
	Value []string `json:"value"`
}

// SettingsTemplate defines model for SettingsTemplate.
type SettingsTemplate struct {
	// Description Description of the template
	Description string `json:"description"`

	// FormFields List of the template form_fields
	FormFields []FormField `json:"form_fields"`

	// Instructions List of the template instructions
	Instructions []Instruction `json:"instructions"`

	// Name Name of the template
	Name string `json:"name"`

	// PublicDescription Public description of the template
	PublicDescription string `json:"public_description"`

	// Sections List of the template sections
	Sections []Section `json:"sections"`

	// Types List of areas where you can use the template
	Types []string `json:"types"`
}

// SettingsUpdateResponse defines model for SettingsUpdateResponse.
type SettingsUpdateResponse struct {
	Changeset SettingsChangeset `json:"changeset"`
	Resources SettingsResources `json:"resources"`
}

// SettingsWebhook defines model for SettingsWebhook.
type SettingsWebhook struct {
	// CallbackURL Callback URL of the webhook
	CallbackURL string `json:"callback_url"`

	// Description Description of the webhook
	Description string `json:"description"`

	// EventNames List of the subscribed event names for the webhook
	EventNames []string `json:"event_names"`

	// Name Name of the webhook
	Name string `json:"name"`

	// Secret Webhook callback URL secret
	Secret string `json:"secret"`
}

// SettingsWorkflow defines model for SettingsWorkflow.
type SettingsWorkflow struct {
	// BlacklistedChannels List of the disallowed channels
	BlacklistedChannels []WorkflowChannel `json:"blacklisted_channels"`

	// CustomFields List of the associated custom field
	CustomFields []WorkflowCustomfield `json:"custom_fields"`

	// DefaultChannels List of the associated channels
	DefaultChannels []WorkflowChannel `json:"default_channels"`

	// Description Description of the workflow
	Description *string `json:"description"`

	// IsFlexibleWorkflow Indicates whether the workflow is flexible
	IsFlexibleWorkflow bool `json:"is_flexible_workflow"`

	// IsNotPushedToLibraryByDefault Indicates whether the contents of the task created from this workflow are not pushed to the library by default
	IsNotPushedToLibraryByDefault bool `json:"is_not_pushed_to_library_by_default"`

	// IsReleasedAsAsset Indicates whether the contents of the task created from this workflow are released as an asset
	IsReleasedAsAsset bool `json:"is_released_as_asset"`

	// Labels List of the associated labels
	Labels []WorkflowLabel `json:"labels"`

	// Name Name of the workflow
	Name string `json:"name"`

	// Steps List of the workflow steps
	Steps []WorkflowStep `json:"steps"`
}

// StringTypeObjectField defines model for StringTypeObjectField.
type StringTypeObjectField struct {
	// ID Unique identifier of the field
	ID string `json:"id"`

	// Type Type of the field
	Type string `json:"type"`

	// Values Any single string
	Values []string `json:"values"`
}

// StringTypeObjectFieldUpdatePayload defines model for StringTypeObjectFieldUpdatePayload.
type StringTypeObjectFieldUpdatePayload struct {
	// Type Type of the field
	Type string `json:"type"`

	// Values Any single string
	Values []string `json:"values"`
}

// StructuredContentBody The fields of structured content
type StructuredContentBody struct {
	// HasEmbedded Whether the structured content has embedded content reference field values
	HasEmbedded          *bool                                                `json:"has_embedded"`
	AdditionalProperties map[string][]LocalizedFieldValuesWithEmbeddedContent `json:"-"`
}

// StructuredContentFields The fields of structured content
type StructuredContentFields map[string][]StructuredContentFields_Item

// StructuredContentFields_Item defines model for StructuredContentFields.Item.
type StructuredContentFields_Item struct {
	union json.RawMessage
}

// Tag defines model for Tag.
type Tag struct {
	// GUID Unique identifier of the tag
	GUID string `json:"guid"`

	// Name Name of the tag
	Name string `json:"name"`
}

// TaskArticle defines model for TaskArticle.
type TaskArticle struct {
	// CreatedAt Date and time on which the article was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// HTMLBody Content of the article
	HTMLBody string `json:"html_body"`

	// ID Unique identifier of the article
	ID string `json:"id"`

	// Labels Labels associated with the article
	Labels []ResourceLabelResponse `json:"labels"`

	// LibraryAssetID Unique identifier of the corresponding library asset — the library asset this task asset was forked from, or the one it was released to.
	LibraryAssetID *string `json:"library_asset_id"`

	// Links Meta links
	Links TaskArticleLinks `json:"links"`

	// ModifiedAt Date and time of the most recent modification of the article, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// Title Title of the article
	Title string `json:"title"`

	// URL Public URL of the article
	URL *string `json:"url"`
}

// TaskArticleLinksWebUrls Web URLs
type TaskArticleLinksWebUrls struct {
	// Drafts Web URL of the drafts of the article – not supported for `article` so the value will be `null`
	Drafts *string `json:"drafts"`

	// Self Web URL of the article
	Self string `json:"self"`

	// Task Web URL of the task that the article is associated with
	Task string `json:"task"`
}

// TaskArticleLinks Meta links
type TaskArticleLinks struct {
	// Drafts URL to POST draft to a task asset – not supported for `article` so the value will be `null`
	Drafts *string `json:"drafts"`

	// Self URL of the article
	Self string `json:"self"`

	// Task URL of the task that the article is associated with
	Task string `json:"task"`

	// WebUrls Web URLs
	WebUrls TaskArticleLinksWebUrls `json:"web_urls"`
}

// TaskAssetCommentResponse defines model for TaskAssetCommentResponse.
type TaskAssetCommentResponse struct {
	// Attachments List of attachments of the comment
	Attachments []AttachmentResponse `json:"attachments"`

	// CreatedAt Creation date and time of the comment, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// ID Unique identifier of the comment
	ID string `json:"id"`

	// IsResolved Determines if the comment is resolved in Optimizely CMP
	IsResolved bool `json:"is_resolved"`

	// Links Meta links
	Links TaskAssetCommentResponseLinks `json:"links"`

	// ModifiedAt Last modification date and time of the comment, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// Value Content of the comment
	Value string `json:"value"`
}

// TaskAssetCommentResponseLinks Meta links
type TaskAssetCommentResponseLinks struct {
	// Asset URL of the asset
	Asset string `json:"asset"`

	// Task URL of the task
	Task string `json:"task"`
}

// TaskAssetDraftBrandComplianceRequest defines model for TaskAssetDraftBrandComplianceRequest.
type TaskAssetDraftBrandComplianceRequest struct {
	// Categories List of compliance categories
	Categories []TaskAssetDraftBrandComplianceRequestCategories `json:"categories"`

	// Status Brand compliance status of the draft
	Status TaskAssetDraftBrandComplianceRequestStatus `json:"status"`
}

// TaskAssetDraftBrandComplianceRequestCategoriesCriteria defines model for .
type TaskAssetDraftBrandComplianceRequestCategoriesCriteria struct {
	// Description Description of the criterion
	Description *string `json:"description"`

	// ID Unique identifier of the criterion
	ID string `json:"id"`

	// Selected Whether the criterion is selected or not
	Selected *bool `json:"selected"`
}

// TaskAssetDraftBrandComplianceRequestCategoriesStatus Brand compliance status of the category
type TaskAssetDraftBrandComplianceRequestCategoriesStatus string

// TaskAssetDraftBrandComplianceRequestCategories defines model for .
type TaskAssetDraftBrandComplianceRequestCategories struct {
	// Criteria List of criteria
	Criteria []TaskAssetDraftBrandComplianceRequestCategoriesCriteria `json:"criteria"`

	// ID Unique identifier of the category
	ID string `json:"id"`

	// Notes Notes left by the reviewer for the category
	Notes *string `json:"notes"`

	// Status Brand compliance status of the category
	Status *TaskAssetDraftBrandComplianceRequestCategoriesStatus `json:"status"`
}

// TaskAssetDraftBrandComplianceRequestStatus Brand compliance status of the draft
type TaskAssetDraftBrandComplianceRequestStatus string

// TaskAssetDraftBrandComplianceResponse defines model for TaskAssetDraftBrandComplianceResponse.
type TaskAssetDraftBrandComplianceResponse struct {
	// Categories List of compliance categories
	Categories []TaskAssetDraftBrandComplianceResponseCategories `json:"categories"`

	// ReviewedAt Date and time of when the draft was reviewed, in ISO 8601 UTC format
	ReviewedAt *time.Time `json:"reviewed_at"`

	// ReviewedBy Unique identifier of the user who reviewed the draft
	ReviewedBy *string `json:"reviewed_by"`

	// Status Brand compliance status of the draft
	Status TaskAssetDraftBrandComplianceResponseStatus `json:"status"`
}

// TaskAssetDraftBrandComplianceResponseCategoriesCriteria defines model for .
type TaskAssetDraftBrandComplianceResponseCategoriesCriteria struct {
	// Description Description of the criterion
	Description *string `json:"description"`

	// ID Unique identifier of the criterion
	ID string `json:"id"`

	// Name Name of the criterion
	Name string `json:"name"`

	// Selected Whether the criterion is selected or not
	Selected *bool `json:"selected"`
}

// TaskAssetDraftBrandComplianceResponseCategoriesStatus Brand compliance status of the category
type TaskAssetDraftBrandComplianceResponseCategoriesStatus string

// TaskAssetDraftBrandComplianceResponseCategories defines model for .
type TaskAssetDraftBrandComplianceResponseCategories struct {
	// Criteria List of criteria
	Criteria []TaskAssetDraftBrandComplianceResponseCategoriesCriteria `json:"criteria"`

	// ID Unique identifier of the category
	ID string `json:"id"`

	// Name Name of the category
	Name string `json:"name"`

	// Notes Notes left by the reviewer for the category
	Notes *string `json:"notes"`

	// Status Brand compliance status of the category
	Status *TaskAssetDraftBrandComplianceResponseCategoriesStatus `json:"status"`
}

// TaskAssetDraftBrandComplianceResponseStatus Brand compliance status of the draft
type TaskAssetDraftBrandComplianceResponseStatus string

// TaskAssetDraftListResponseItem defines model for TaskAssetDraftListResponseItem.
type TaskAssetDraftListResponseItem struct {
	// Content Content of the draft
	Content TaskAssetDraftListResponseItemContent `json:"content"`

	// CreatedAt Date and time on which the draft was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// DraftNumber The serial number of the draft
	DraftNumber int `json:"draft_number"`

	// ID Unique identifier of the draft
	ID string `json:"id"`

	// Links Meta links
	Links TaskAssetDraftListResponseItemLinks `json:"links"`

	// MimeType MIME type of the draft
	MimeType string `json:"mime_type"`

	// Title Title of the draft
	Title string `json:"title"`

	// Type Type of the draft
	Type TaskAssetDraftListResponseItemType `json:"type"`
}

// TaskAssetDraftListResponseItemContentType Type of the content
type TaskAssetDraftListResponseItemContentType string

// TaskAssetDraftListResponseItemContent Content of the draft
type TaskAssetDraftListResponseItemContent struct {
	// Type Type of the content
	Type TaskAssetDraftListResponseItemContentType `json:"type"`

	// Value Content of the draft; download URL for `image`, `raw file`, and `video`
	Value string `json:"value"`
}

// TaskAssetDraftListResponseItemLinks Meta links
type TaskAssetDraftListResponseItemLinks struct {
	// Asset URL of the asset that the draft is associated with
	Asset string `json:"asset"`

	// Task URL of the task that the asset is associated with
	Task string `json:"task"`
}

// TaskAssetDraftListResponseItemType Type of the draft
type TaskAssetDraftListResponseItemType string

// TaskAssetDraftResponse defines model for TaskAssetDraftResponse.
type TaskAssetDraftResponse struct {
	// AssetID Unique identifier of the asset
	AssetID string `json:"asset_id"`

	// Content Content of the draft
	Content TaskAssetDraftResponseContent `json:"content"`

	// CreatedAt Date and time on which the draft was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// DraftNumber The serial number of the draft
	DraftNumber int `json:"draft_number"`

	// Links Meta links
	Links TaskAssetDraftResponseLinks `json:"links"`

	// MimeType MIME type of the draft
	MimeType string `json:"mime_type"`

	// Title Title of the draft
	Title string `json:"title"`

	// Type Type of the draft
	Type TaskAssetDraftResponseType `json:"type"`
}

// TaskAssetDraftResponseContentType Type of the content
type TaskAssetDraftResponseContentType string

// TaskAssetDraftResponseContent Content of the draft
type TaskAssetDraftResponseContent struct {
	// Type Type of the content
	Type TaskAssetDraftResponseContentType `json:"type"`

	// Value Content of the draft; download URL for `image`, `raw file`, and `video`
	Value string `json:"value"`
}

// TaskAssetDraftResponseLinks Meta links
type TaskAssetDraftResponseLinks struct {
	// Asset URL of the asset that the draft is associated with
	Asset string `json:"asset"`

	// Task URL of the task that the asset is associated with
	Task string `json:"task"`
}

// TaskAssetDraftResponseType Type of the draft
type TaskAssetDraftResponseType string

// TaskAssetFieldsUpdateRequest defines model for TaskAssetFieldsUpdateRequest.
type TaskAssetFieldsUpdateRequest = []TaskAssetFieldsUpdateRequest_Item

// TaskAssetFieldsUpdateRequest_Item defines model for TaskAssetFieldsUpdateRequest.Item.
type TaskAssetFieldsUpdateRequest_Item struct {
	union json.RawMessage
}

// TaskAssetRequest defines model for TaskAssetRequest.
type TaskAssetRequest struct {
	union json.RawMessage
}

// TaskAssetRequestForDirectUpload defines model for TaskAssetRequestForDirectUpload.
type TaskAssetRequestForDirectUpload struct {
	// Key Unique identifier of the file upload session. This is the `upload_meta_fields.key` field retrieved from the `/v3/upload-url` endpoint.
	Key string `json:"key"`

	// Title Title of the asset
	Title string `json:"title"`

	// Type The mechanism for adding the asset to the task
	Type *string `json:"type,omitempty"`
}

// TaskAssetRequestForLibrary defines model for TaskAssetRequestForLibrary.
type TaskAssetRequestForLibrary struct {
	// ContentID Unique identifier of the asset
	ContentID openapi_types.UUID `json:"content_id"`

	// ContentType The type of asset
	ContentType TaskAssetRequestForLibraryContentType `json:"content_type"`

	// Type The mechanism for adding the asset to the task
	Type string `json:"type"`
}

// TaskAssetRequestForLibraryContentType The type of asset
type TaskAssetRequestForLibraryContentType string

// TaskAssetResponse defines model for TaskAssetResponse.
type TaskAssetResponse struct {
	// Content Content of the asset
	Content AssetContent `json:"content"`

	// CreatedAt Date and time on which the asset was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// ID Unique identifier for the asset
	ID string `json:"id"`

	// Labels Labels associated with the asset
	Labels []ResourceLabelResponse `json:"labels"`

	// Links Meta links
	Links TaskAssetResponseLinks `json:"links"`

	// MimeType MIME type of the asset
	MimeType string `json:"mime_type"`

	// ModifiedAt Date and time of the most recent modification of the asset, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// Title Title of the asset
	Title string `json:"title"`

	// Type Type of the asset
	Type TaskAssetResponseType `json:"type"`
}

// TaskAssetResponseType Type of the asset
type TaskAssetResponseType string

// TaskBriefResponse defines model for TaskBriefResponse.
type TaskBriefResponse struct {
	// Fields List of fields of the brief
	Fields []TaskBriefResponseFields `json:"fields"`

	// Links Meta links
	Links    TaskBriefResponseLinks  `json:"links"`
	Template *TaskBriefTemplateValue `json:"template"`

	// Title Title of the task brief
	Title string `json:"title"`

	// Type Type of the task brief
	Type TaskBriefResponseType `json:"type"`
}

// TaskBriefResponseFields defines model for .
type TaskBriefResponseFields struct {
	// Name Name of the brief field
	Name string `json:"name"`

	// SettingsFieldID Unique identifier of the settings field. If available, can be used to look up the field details using settings API.
	SettingsFieldID *string `json:"settings_field_id"`

	// Value List of values set for the brief field
	Value []string `json:"value"`
}

// TaskBriefResponseLinks Meta links
type TaskBriefResponseLinks struct {
	// Self URL of the brief
	Self string `json:"self"`

	// Task URL of the task
	Task string `json:"task"`
}

// TaskBriefResponseType Type of the task brief
type TaskBriefResponseType string

// TaskBriefTemplateValue Template info of the brief if the brief is of template type
type TaskBriefTemplateValue struct {
	// ID Unique identifier of the template
	ID string `json:"id"`

	// Name Name of the template
	Name string `json:"name"`
}

// TaskCommentResponse defines model for TaskCommentResponse.
type TaskCommentResponse struct {
	// Attachments List of attachments of the comment
	Attachments []AttachmentResponse `json:"attachments"`

	// CommentBy Unique identifier of the user who posted the comment
	CommentBy string `json:"comment_by"`

	// CreatedAt Creation date and time of the task comment, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// ID Unique identifier of the task comment
	ID string `json:"id"`

	// IsResolved Determines if the comment is resolved in Optimizely CMP
	IsResolved bool `json:"is_resolved"`

	// Links Meta links
	Links TaskCommentResponseLinks `json:"links"`

	// ModifiedAt Last modification date and time of the task comment, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// ParentCommentID Unique identifier of the parent comment
	ParentCommentID *string `json:"parent_comment_id"`

	// Value Content of the comment
	Value string `json:"value"`
}

// TaskCommentResponseLinks Meta links
type TaskCommentResponseLinks struct {
	// CommentBy URL of the user who posted the comment
	CommentBy string `json:"comment_by"`

	// Task URL of the task
	Task string `json:"task"`
}

// TaskCreateRequest defines model for TaskCreateRequest.
type TaskCreateRequest struct {
	// CampaignID Id of the campaign to be associated with the task
	CampaignID *string `json:"campaign_id,omitempty"`

	// DueAt Due date and time of the task, in ISO 8601 UTC format (2020-02-10T10:40:45Z). The `due_at` field is required if `workflow_id` is present and the workflow does not have step duration. If `workflow_id` is present and the workflow has step duration, then either `start_at` or `due_at` is required.
	DueAt *time.Time `json:"due_at,omitempty"`

	// OwnerID Id of the owner of the task
	OwnerID *string `json:"owner_id,omitempty"`

	// StartAt Start date and time of the task, in ISO 8601 UTC format (2020-02-10T10:40:45Z). If `workflow_id` is present and the workflow has step duration, then either `start_at` or `due_at` is required. If both `start_at` and `due_at` is provided with a `workflow_id` that has step duration, then the `start_at` field will be ignored.
	StartAt *time.Time `json:"start_at,omitempty"`

	// Title Title of the task
	Title string `json:"title"`

	// WorkflowID Id of the workflow to be used in the task
	WorkflowID *string `json:"workflow_id,omitempty"`
}

// TaskCustomField defines model for TaskCustomField.
type TaskCustomField struct {
	// ID Unique identifier for the custom field
	ID string `json:"id"`

	// Links Meta links
	Links TaskCustomFieldLinks `json:"links"`

	// Name Name of the custom field
	Name string `json:"name"`

	// Type Type of the custom field
	Type TaskCustomFieldType `json:"type"`

	// Values List of values of the custom field
	Values []TaskCustomFieldValues `json:"values"`
}

// TaskCustomFieldLinks Meta links
type TaskCustomFieldLinks struct {
	// Choices URL of the available custom field choices
	Choices *string `json:"choices"`

	// Self URL of the custom field
	Self string `json:"self"`
}

// TaskCustomFieldType Type of the custom field
type TaskCustomFieldType string

// TaskCustomFieldValues defines model for .
type TaskCustomFieldValues struct {
	// ID Identifier of the custom field value
	ID *string `json:"id"`

	// Name Name of the custom field value
	Name string `json:"name"`
}

// TaskCustomFieldChoiceListResponseItem defines model for TaskCustomFieldChoiceListResponseItem.
type TaskCustomFieldChoiceListResponseItem struct {
	// ID Unique identifier of the custom field choice
	ID string `json:"id"`

	// Name Name of the custom field choice
	Name string `json:"name"`
}

// TaskCustomFieldUpdateRequest defines model for TaskCustomFieldUpdateRequest.
type TaskCustomFieldUpdateRequest struct {
	// Values An array of choice IDs, if the custom field has choices. Otherwise, text values to update the custom field values with. Must contain at least one id/value in the array. For custom fields of type `checkboxes`, you can provide an array of multiple-choice IDs as values.
	// Value for the custom field of type `date_field` must follow this format: `YYYY-MM-DD`.
	Values []string `json:"values"`
}

// TaskExternalWorkRequest defines model for TaskExternalWorkRequest.
type TaskExternalWorkRequest struct {
	// Identifier Identifier of the external work
	Identifier *string `json:"identifier"`

	// Status Status of the external work
	Status *string `json:"status"`

	// Title Title of the external work
	Title *string `json:"title"`

	// URL Link to the external work (must be a valid URL with HTTPS scheme)
	URL *string `json:"url"`
}

// TaskExternalWorkResponse defines model for TaskExternalWorkResponse.
type TaskExternalWorkResponse struct {
	// ExternalSystem Name of the external system
	ExternalSystem string `json:"external_system"`

	// Identifier Identifier of the external work
	Identifier *string `json:"identifier"`

	// Links Meta links
	Links TaskExternalWorkResponseLinks `json:"links"`

	// Status Status of the external work
	Status *string `json:"status"`

	// Title Title of the external work
	Title *string `json:"title"`

	// URL Link to the external work
	URL *string `json:"url"`
}

// TaskExternalWorkResponseLinks Meta links
type TaskExternalWorkResponseLinks struct {
	// Self URL to get or update the external work information
	Self string `json:"self"`
}

// TaskFieldUpdateRequest defines model for TaskFieldUpdateRequest.
type TaskFieldUpdateRequest struct {
	union json.RawMessage
}

// TaskImage defines model for TaskImage.
type TaskImage struct {
	// CreatedAt Date and time on which the image was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// FileSize Size of the image in bytes
	FileSize int `json:"file_size"`

	// ID Unique identifier of the image
	ID string `json:"id"`

	// ImageResolution Width and height of the image in pixels
	ImageResolution TaskImageImageResolution `json:"image_resolution"`

	// Labels Labels associated with the image
	Labels []ResourceLabelResponse `json:"labels"`

	// LibraryAssetID Unique identifier of the corresponding library asset — the library asset this task asset was forked from, or the one it was released to.
	LibraryAssetID *string `json:"library_asset_id"`

	// Links Meta links
	Links TaskImageLinks `json:"links"`

	// MimeType MIME type of the image
	MimeType string `json:"mime_type"`

	// ModifiedAt Date and time of the most recent modification of the image, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// Title Title of the image
	Title string `json:"title"`

	// URL Download the URL of the image
	URL string `json:"url"`
}

// TaskImageImageResolution Width and height of the image in pixels
type TaskImageImageResolution struct {
	// Height Height of the image in pixels
	Height int `json:"height"`

	// Width Width of the image in pixels
	Width int `json:"width"`
}

// TaskImageLinksWebUrls Web URLs
type TaskImageLinksWebUrls struct {
	// Drafts Web URL of the drafts of the image
	Drafts string `json:"drafts"`

	// Self Web URL of the image
	Self string `json:"self"`

	// Task Web URL of the task that the image is associated with
	Task string `json:"task"`
}

// TaskImageLinks Meta links
type TaskImageLinks struct {
	// Drafts URL to POST draft to the image
	Drafts string `json:"drafts"`

	// Self URL of the image
	Self string `json:"self"`

	// Task URL of the task that the image is associated with
	Task string `json:"task"`

	// WebUrls Web URLs
	WebUrls TaskImageLinksWebUrls `json:"web_urls"`
}

// TaskListResponseItem defines model for TaskListResponseItem.
type TaskListResponseItem struct {
	// CampaignID Unique identifier of the campaign
	CampaignID *string `json:"campaign_id"`

	// DueAt Due date of the task in ISO 8601 UTC format
	DueAt *string `json:"due_at"`

	// ID Unique identifier of the task
	ID string `json:"id"`

	// IsArchived Whether the task is archived or not
	IsArchived bool `json:"is_archived"`

	// IsCompleted Whether the task is completed or not
	IsCompleted bool `json:"is_completed"`

	// Links Meta links
	Links TaskListResponseItemLinks `json:"links"`

	// MilestoneID Unique identifier of the milestone
	MilestoneID *string `json:"milestone_id"`

	// ModifiedAt Modified date of the task in ISO 8601 UTC format
	ModifiedAt *string `json:"modified_at"`

	// ReferenceID Reference identifier of the task
	ReferenceID string `json:"reference_id"`

	// StartAt Start date of the task in ISO 8601 UTC format
	StartAt *string `json:"start_at"`

	// Status Current status of the task
	Status TaskListResponseItemStatus `json:"status"`

	// Title Title of the task
	Title string `json:"title"`

	// WorkflowID Unique identifier of the workflow
	WorkflowID *string `json:"workflow_id"`
}

// TaskListResponseItemLinksWebUrls Web URLs
type TaskListResponseItemLinksWebUrls struct {
	// Self Web URL of the task
	Self string `json:"self"`
}

// TaskListResponseItemLinks Meta links
type TaskListResponseItemLinks struct {
	// Campaign URL of the campaign that the task is associated with
	Campaign *string `json:"campaign"`

	// Milestone URL of the milestone that the task is associated with
	Milestone *string `json:"milestone"`

	// Self URL of the task
	Self string `json:"self"`

	// WebUrls Web URLs
	WebUrls TaskListResponseItemLinksWebUrls `json:"web_urls"`

	// Workflow URL of the workflow the task
	Workflow *string `json:"workflow"`
}

// TaskListResponseItemStatus Current status of the task
type TaskListResponseItemStatus string

// TaskPublishingIntentCreateRequest defines model for TaskPublishingIntentCreateRequest.
type TaskPublishingIntentCreateRequest struct {
	// ChannelID Unique identifier of the target publishing channel
	ChannelID string `json:"channel_id"`
}

// TaskPublishingIntentResponse defines model for TaskPublishingIntentResponse.
type TaskPublishingIntentResponse struct {
	// ID Unique identifier of the publishing intent
	ID string `json:"id"`
}

// TaskRawFile defines model for TaskRawFile.
type TaskRawFile struct {
	// CreatedAt Date and time on which the raw file was created, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// FileSize Size of the raw file in bytes
	FileSize int `json:"file_size"`

	// ID Unique identifier of the raw file
	ID string `json:"id"`

	// Labels Labels associated to the raw file
	Labels []ResourceLabelResponse `json:"labels"`

	// LibraryAssetID Unique identifier of the corresponding library asset — the library asset this task asset was forked from, or the one it was released to.
	LibraryAssetID *string `json:"library_asset_id"`

	// Links Meta links
	Links TaskRawFileLinks `json:"links"`

	// MimeType MIME type of the raw file
	MimeType string `json:"mime_type"`

	// ModifiedAt Date and time of the most recent modification of the raw file,  in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// Title Title of the raw file
	Title string `json:"title"`

	// URL Download the URL of the raw file
	URL string `json:"url"`
}

// TaskRawFileLinksWebUrls Web URLs
type TaskRawFileLinksWebUrls struct {
	// Drafts Web URL of the drafts of the raw file
	Drafts string `json:"drafts"`

	// Self Web URL of the raw file
	Self string `json:"self"`

	// Task Web URL of the task that the raw file is associated with
	Task string `json:"task"`
}

// TaskRawFileLinks Meta links
type TaskRawFileLinks struct {
	// Drafts URL to POST draft to the raw file
	Drafts string `json:"drafts"`

	// Self URL of the raw file
	Self string `json:"self"`

	// Task URL of the task that the raw file is associated with
	Task string `json:"task"`

	// WebUrls Web URLs
	WebUrls TaskRawFileLinksWebUrls `json:"web_urls"`
}

// TaskResponse defines model for TaskResponse.
type TaskResponse struct {
	// CampaignID Unique identifier of the campaign
	CampaignID *string `json:"campaign_id"`

	// DueAt Due date and time of the task, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	DueAt *time.Time `json:"due_at"`

	// ID Unique identifier of the task
	ID string `json:"id"`

	// IsArchived Determines whether the task is archived
	IsArchived bool `json:"is_archived"`

	// IsCompleted Determines whether the task is completed
	IsCompleted bool `json:"is_completed"`

	// Labels Labels associated with the task
	Labels []ResourceLabelResponse `json:"labels"`

	// Links Meta links
	Links TaskResponseLinks `json:"links"`

	// MilestoneID Unique identifier of the milestone
	MilestoneID *string `json:"milestone_id"`

	// ModifiedAt Modified date of the task in ISO 8601 UTC format
	ModifiedAt *string `json:"modified_at"`

	// ReferenceID Reference ID of the task
	ReferenceID string `json:"reference_id"`

	// StartAt Start date and time of the task, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	StartAt *time.Time `json:"start_at"`

	// Status Current status of the task
	Status TaskResponseStatus `json:"status"`

	// Steps Steps of the task
	Steps []TaskStep `json:"steps"`

	// Title Title of the task
	Title string `json:"title"`

	// WorkflowID Unique identifier of workflow
	WorkflowID *string `json:"workflow_id"`
}

// TaskResponseLinksWebUrls Web URLs
type TaskResponseLinksWebUrls struct {
	// Brief Web URL of the task brief
	Brief string `json:"brief"`

	// Self Web URL of the task
	Self string `json:"self"`
}

// TaskResponseLinks Meta links
type TaskResponseLinks struct {
	// Assets URL for getting the list of the task assets or adding a new asset to the task
	Assets string `json:"assets"`

	// Attachments URL for getting the list of the task attachments
	Attachments string `json:"attachments"`

	// Brief URL of the task brief
	Brief *string `json:"brief"`

	// Campaign URL of the campaign that the task is associated with
	Campaign string `json:"campaign"`

	// CustomFields URL of the list of custom fields added to the task
	CustomFields *string `json:"custom_fields"`

	// Milestone URL of the milestone that the task is associated with
	Milestone *string `json:"milestone"`

	// Self URL of the task
	Self string `json:"self"`

	// WebUrls Web URLs
	WebUrls TaskResponseLinksWebUrls `json:"web_urls"`
}

// TaskResponseStatus Current status of the task
type TaskResponseStatus string

// TaskStep defines model for TaskStep.
type TaskStep struct {
	// Description Description of the step
	Description *string `json:"description"`

	// DueAt Due date and time of the step containing the substep,  in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	DueAt *time.Time `json:"due_at"`

	// ID Unique identifier of the step
	ID string `json:"id"`

	// IsCompleted Determines whether the step is completed in Optimizely CMP
	IsCompleted bool `json:"is_completed"`

	// SubSteps Substeps of the step
	SubSteps []TaskSubStep `json:"sub_steps"`

	// Title Title of the step
	Title string `json:"title"`
}

// TaskStepRequest defines model for TaskStepRequest.
type TaskStepRequest struct {
	// DueAt Due date and time of the task step, in ISO 8601 UTC format (2020-02-10T10:40:45Z). Point to note: in the CMP web interface, a popup is shown while updating a step's due date to inform the user of temporal or sequential constraints based on the due date of other steps. However, this API will update the step's due date while superceding any validation for the workflow.
	DueAt *time.Time `json:"due_at,omitempty"`
}

// TaskStructuredContentCreateRequest defines model for TaskStructuredContentCreateRequest.
type TaskStructuredContentCreateRequest struct {
	// ContentTypeGUID Unique identifier of the content type
	ContentTypeGUID string `json:"content_type_guid"`

	// Expired Expired status of the content
	Expired *bool `json:"expired,omitempty"`

	// ExpiryDatetime Date and time on which the content will expire, in ISO 8601 UTC format
	ExpiryDatetime *time.Time `json:"expiry_datetime"`

	// Fields The fields of structured content
	Fields *StructuredContentBody `json:"fields,omitempty"`

	// PrimaryLocale The primary locale of the content
	PrimaryLocale *string `json:"primary_locale,omitempty"`

	// RootContent This is true when the content is not embedded in another content through a content type reference field
	RootContent *bool `json:"root_content,omitempty"`

	// Source Source of the content
	Source *string `json:"source"`

	// SourceID Source id of the content
	SourceID *string `json:"source_id"`

	// SourceMetadata Source metadata of the content
	SourceMetadata *string `json:"source_metadata"`

	// TemplateGUID Unique identifier of the template
	TemplateGUID *string `json:"template_guid,omitempty"`

	// Title Title of the structured content
	Title *string `json:"title,omitempty"`
}

// TaskStructuredContentDraftRequest defines model for TaskStructuredContentDraftRequest.
type TaskStructuredContentDraftRequest struct {
	ContentBody LocalizedFieldValues `json:"content_body"`
}

// TaskStructuredContentUpdateRequest defines model for TaskStructuredContentUpdateRequest.
type TaskStructuredContentUpdateRequest struct {
	// Expired Expired status of the content
	Expired *bool `json:"expired,omitempty"`

	// ExpiryDatetime Date and time on which the content will expire, in ISO 8601 UTC format
	ExpiryDatetime *time.Time `json:"expiry_datetime"`

	// Fields The fields of structured content
	Fields *StructuredContentFields `json:"fields,omitempty"`

	// PrimaryLocale The primary locale of the content
	PrimaryLocale *string `json:"primary_locale,omitempty"`

	// Source Source of the content
	Source *string `json:"source"`

	// SourceID Source id of the content
	SourceID *string `json:"source_id"`

	// SourceMetadata Source metadata of the content
	SourceMetadata *string `json:"source_metadata"`

	// Title Title of the structured content
	Title *string `json:"title,omitempty"`
}

// TaskSubStep defines model for TaskSubStep.
type TaskSubStep struct {
	// AssigneeID ID of the assignee of the step containing the substep
	AssigneeID *string `json:"assignee_id"`

	// AssigneeType Type of the assignee of the step containing the substep.
	AssigneeType TaskSubStepAssigneeType `json:"assignee_type"`

	// ID Unique identifier of the substep
	ID string `json:"id"`

	// IsCompleted Determines whether the substep is completed in Optimizely CMP
	IsCompleted bool `json:"is_completed"`

	// IsExternal Determines whether the substep is external in Optimizely CMP
	IsExternal bool `json:"is_external"`

	// IsInProgress Determines whether the substep is In Progress in Optimizely CMP
	IsInProgress bool `json:"is_in_progress"`

	// IsSkipped Determines whether the substep is skipped in Optimizely CMP
	IsSkipped bool `json:"is_skipped"`

	// Links Meta links
	Links TaskSubStepLinks `json:"links"`

	// Title Title of the substep
	Title string `json:"title"`
}

// TaskSubStepAssigneeType Type of the assignee of the step containing the substep.
type TaskSubStepAssigneeType string

// TaskSubStepLinks Meta links
type TaskSubStepLinks struct {
	// Assignee Assignee of the step containing the substep
	Assignee *string `json:"assignee"`

	// ExternalWork URL of the external work associated with substep if substep is external in _Optimizely CMP_
	ExternalWork *string `json:"external_work"`

	// Self URL of the substep
	Self string `json:"self"`

	// Task URL of the task
	Task string `json:"task"`
}

// TaskSubStepCommentResponse defines model for TaskSubStepCommentResponse.
type TaskSubStepCommentResponse struct {
	// Attachments List of attachments of the comment
	Attachments []AttachmentResponse `json:"attachments"`

	// CreatedAt Creation date and time of the task substep comment,  in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// ID Unique identifier of the substep comment
	ID string `json:"id"`

	// IsResolved Determines if the comment is resolved in Optimizely CMP
	IsResolved bool `json:"is_resolved"`

	// Links Meta links
	Links *TaskSubStepCommentResponseLinks `json:"links,omitempty"`

	// ModifiedAt Last modification date and time of the task substep comment,  in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// Value Content of the comment
	Value string `json:"value"`
}

// TaskSubStepCommentResponseLinks Meta links
type TaskSubStepCommentResponseLinks struct {
	// CommentBy Creator of the substep comment
	CommentBy string `json:"comment_by"`

	// Self URL of the substep comment
	Self string `json:"self"`

	// SubStep URL of the substep
	SubStep string `json:"sub_step"`

	// Task URL of the task
	Task string `json:"task"`
}

// TaskSubStepCommentUpdateRequest defines model for TaskSubStepCommentUpdateRequest.
type TaskSubStepCommentUpdateRequest struct {
	// Value Updated comment value. Markdown is supported. To mention a user belonging to the organization, use `@[name](openapi-user-link)` format.
	Value string `json:"value"`
}

// TaskSubStepRequest defines model for TaskSubStepRequest.
type TaskSubStepRequest struct {
	// AssigneeID ID of the assignee for the substep; cannot be passed together with `is_completed`, `is_in_progress` or `is_skipped`
	AssigneeID *string `json:"assignee_id"`

	// AssigneeType Type of the assignee for the substep. If not provided with `assignee_id`, defaults to `user`.
	AssigneeType *TaskSubStepRequestAssigneeType `json:"assignee_type,omitempty"`

	// IsCompleted Whether the substep should be completed; cannot be passed together with `is_in_progress`, `is_skipped` or `assignee_id`
	IsCompleted *TaskSubStepRequestIsCompleted `json:"is_completed,omitempty"`

	// IsInProgress Whether the substep should be progressed; cannot be passed together with `is_completed`, `is_skipped` or `assignee_id`
	IsInProgress *TaskSubStepRequestIsInProgress `json:"is_in_progress,omitempty"`

	// IsSkipped Whether the substep should be skipped; cannot be passed together with `is_in_progress`, `is_completed` or `assignee_id`
	IsSkipped *TaskSubStepRequestIsSkipped `json:"is_skipped,omitempty"`
}

// TaskSubStepRequestAssigneeType Type of the assignee for the substep. If not provided with `assignee_id`, defaults to `user`.
type TaskSubStepRequestAssigneeType string

// TaskSubStepRequestIsCompleted Whether the substep should be completed; cannot be passed together with `is_in_progress`, `is_skipped` or `assignee_id`
type TaskSubStepRequestIsCompleted bool

// TaskSubStepRequestIsInProgress Whether the substep should be progressed; cannot be passed together with `is_completed`, `is_skipped` or `assignee_id`
type TaskSubStepRequestIsInProgress bool

// TaskSubStepRequestIsSkipped Whether the substep should be skipped; cannot be passed together with `is_in_progress`, `is_completed` or `assignee_id`
type TaskSubStepRequestIsSkipped bool

// TaskUpdateRequest defines model for TaskUpdateRequest.
type TaskUpdateRequest struct {
	// CampaignID ID of the campaign to be associated with the task
	CampaignID *string `json:"campaign_id,omitempty"`

	// DueAt Due date and time of the task
	DueAt *time.Time `json:"due_at,omitempty"`

	// Labels Labels to associate
	Labels *[]ResourceLabelRequest `json:"labels,omitempty"`

	// OwnerID ID of the owner for the task
	OwnerID *string `json:"owner_id,omitempty"`

	// StartAt Start date of the task
	StartAt *time.Time `json:"start_at,omitempty"`

	// Title Title of the task
	Title *string `json:"title,omitempty"`

	// WorkflowID ID of the workflow to be used in the task
	WorkflowID *string `json:"workflow_id,omitempty"`
}

// TaskURLResponse defines model for TaskUrlResponse.
type TaskURLResponse struct {
	// CreatedAt Timestamp when the URL was added to the task
	CreatedAt time.Time `json:"created_at"`

	// ID Unique identifier of the URL added to the task
	ID string `json:"id"`
}

// TaskVideo defines model for TaskVideo.
type TaskVideo struct {
	// CreatedAt Date and time when the video was created,  in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// FileSize Size of the video in bytes
	FileSize int `json:"file_size"`

	// ID Unique identifier for the video
	ID string `json:"id"`

	// Labels Labels associated with the video
	Labels []ResourceLabelResponse `json:"labels"`

	// LibraryAssetID Unique identifier of the corresponding library asset — the library asset this task asset was forked from, or the one it was released to.
	LibraryAssetID *string `json:"library_asset_id"`

	// Links Meta links
	Links TaskVideoLinks `json:"links"`

	// MimeType MIME type of the video
	MimeType string `json:"mime_type"`

	// ModifiedAt Date and time of the most recent modification of the video, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// Title Title of the video
	Title string `json:"title"`

	// URL Download URL of the video
	URL string `json:"url"`
}

// TaskVideoLinksWebUrls Web URLs
type TaskVideoLinksWebUrls struct {
	// Drafts Web URL of the drafts of the video
	Drafts string `json:"drafts"`

	// Self Web URL of the video
	Self string `json:"self"`

	// Task Web URL of the task that the video is associated with
	Task string `json:"task"`
}

// TaskVideoLinks Meta links
type TaskVideoLinks struct {
	// Drafts URL to POST draft to the video
	Drafts string `json:"drafts"`

	// Self URL of the video
	Self string `json:"self"`

	// Task URL of the task that the video is associated with
	Task string `json:"task"`

	// WebUrls Web URLs
	WebUrls TaskVideoLinksWebUrls `json:"web_urls"`
}

// Team defines model for Team.
type Team struct {
	// ID Unique identifier of the team
	ID string `json:"id"`

	// Links Meta links
	Links TeamLinks `json:"links"`

	// Name Name of the team
	Name string `json:"name"`
}

// TeamLinks Meta links
type TeamLinks struct {
	// Self URL of the team
	Self string `json:"self"`
}

// TeamWithUsers defines model for TeamWithUsers.
type TeamWithUsers struct {
	// ID Unique identifier of the team
	ID string `json:"id"`

	// Links Meta links
	Links TeamWithUsersLinks `json:"links"`

	// Name Name of the team
	Name string `json:"name"`

	// Users List of users in the team
	Users []TeamWithUsersUsers `json:"users"`
}

// TeamWithUsersLinks Meta links
type TeamWithUsersLinks struct {
	// Self URL of the team
	Self string `json:"self"`
}

// TeamWithUsersUsers defines model for .
type TeamWithUsersUsers struct {
	// Email Email address of the user
	Email string `json:"email"`

	// ID Unique identifier of the user
	ID string `json:"id"`

	// Name Name of the user
	Name string `json:"name"`
}

// TemplateChoiceFormField defines model for TemplateChoiceFormField.
type TemplateChoiceFormField struct {
	// HelperText Helper text of the field
	HelperText *string `json:"helper_text"`

	// Identifier Identifier of the field
	Identifier string `json:"identifier"`

	// IsProtected True when this form field is read-only and only organization admins can modify it. Absent or false for editable fields.
	IsProtected *bool `json:"is_protected,omitempty"`

	// IsReadonly Specify if the field can be used to create any resource
	IsReadonly bool `json:"is_readonly"`

	// IsRequired Required status of the field
	IsRequired bool `json:"is_required"`

	// Label Label of the field
	Label *string `json:"label"`

	// LogicRules List of logic rules
	LogicRules TemplateLogicRule `json:"logic_rules"`

	// SortOrder Sort order of the field
	SortOrder int `json:"sort_order"`

	// Type Type of the field
	Type TemplateChoiceFormFieldType `json:"type"`

	// TypeSpecificMeta Type-specific meta-information
	TypeSpecificMeta TemplateChoiceFormFieldTypeSpecificMeta `json:"type_specific_meta"`
}

// TemplateChoiceFormFieldType Type of the field
type TemplateChoiceFormFieldType string

// TemplateChoiceFormFieldTypeSpecificMeta Type-specific meta-information
type TemplateChoiceFormFieldTypeSpecificMeta struct {
	// Choices Choices for the field
	Choices []ResourceChoiceResponse `json:"choices"`

	// IsMultiSelect Whether multiple selection is allowed
	IsMultiSelect bool `json:"is_multi_select"`
}

// TemplateCurrencyNumberFormField defines model for TemplateCurrencyNumberFormField.
type TemplateCurrencyNumberFormField struct {
	// HelperText Helper text of the field
	HelperText *string `json:"helper_text"`

	// Identifier Identifier of the field
	Identifier string `json:"identifier"`

	// IsProtected True when this form field is read-only and only organization admins can modify it. Absent or false for editable fields.
	IsProtected *bool `json:"is_protected,omitempty"`

	// IsReadonly Specify if the field can be used to create any resource
	IsReadonly bool `json:"is_readonly"`

	// IsRequired Required status of the field
	IsRequired bool `json:"is_required"`

	// Label Label of the field
	Label *string `json:"label"`

	// LogicRules List of logic rules
	LogicRules TemplateLogicRule `json:"logic_rules"`

	// SortOrder Sort order of the field
	SortOrder int `json:"sort_order"`

	// Type Type of the field
	Type TemplateCurrencyNumberFormFieldType `json:"type"`

	// TypeSpecificMeta Type-specific meta-information
	TypeSpecificMeta TemplateCurrencyNumberFormFieldTypeSpecificMeta `json:"type_specific_meta"`
}

// TemplateCurrencyNumberFormFieldType Type of the field
type TemplateCurrencyNumberFormFieldType string

// TemplateCurrencyNumberFormFieldTypeSpecificMeta Type-specific meta-information
type TemplateCurrencyNumberFormFieldTypeSpecificMeta struct {
	// CurrencyCode Currency code of the field
	CurrencyCode string `json:"currency_code"`

	// DecimalPlaces Decimal places allowed for the field
	DecimalPlaces *int `json:"decimal_places"`

	// HasThousandSeparator Whether the field has thousand separators
	HasThousandSeparator bool `json:"has_thousand_separator"`
}

// TemplateDefaultFormField defines model for TemplateDefaultFormField.
type TemplateDefaultFormField struct {
	// HelperText Helper text of the field
	HelperText *string `json:"helper_text"`

	// Identifier Identifier of the field
	Identifier string `json:"identifier"`

	// IsProtected True when this form field is read-only and only organization admins can modify it. Absent or false for editable fields.
	IsProtected *bool `json:"is_protected,omitempty"`

	// IsReadonly Specify if the field can be used to create any resource
	IsReadonly bool `json:"is_readonly"`

	// IsRequired Required status of the field
	IsRequired bool `json:"is_required"`

	// Label Label of the field
	Label *string `json:"label"`

	// LogicRules List of logic rules
	LogicRules TemplateLogicRule `json:"logic_rules"`

	// SortOrder Sort order of the field
	SortOrder int `json:"sort_order"`

	// Type Type of the field
	Type TemplateDefaultFormFieldType `json:"type"`

	// TypeSpecificMeta Type-specific meta-information
	TypeSpecificMeta *map[string]interface{} `json:"type_specific_meta"`
}

// TemplateDefaultFormFieldType Type of the field
type TemplateDefaultFormFieldType string

// TemplateInstructionFormField defines model for TemplateInstructionFormField.
type TemplateInstructionFormField struct {
	// HelperText Helper text of the field
	HelperText *string `json:"helper_text"`

	// Identifier Identifier of the field
	Identifier string `json:"identifier"`

	// IsProtected True when this form field is read-only and only organization admins can modify it. Absent or false for editable fields.
	IsProtected *bool `json:"is_protected,omitempty"`

	// IsReadonly Specify if the field can be used to create any resource
	IsReadonly bool `json:"is_readonly"`

	// IsRequired Required status of the field
	IsRequired bool `json:"is_required"`

	// Label Label of the field
	Label *string `json:"label"`

	// LogicRules List of logic rules
	LogicRules TemplateLogicRule `json:"logic_rules"`

	// SortOrder Sort order of the field
	SortOrder int `json:"sort_order"`

	// Type Type of the field
	Type TemplateInstructionFormFieldType `json:"type"`

	// TypeSpecificMeta Type-specific meta-information
	TypeSpecificMeta TemplateInstructionFormFieldTypeSpecificMeta `json:"type_specific_meta"`
}

// TemplateInstructionFormFieldType Type of the field
type TemplateInstructionFormFieldType string

// TemplateInstructionFormFieldTypeSpecificMeta Type-specific meta-information
type TemplateInstructionFormFieldTypeSpecificMeta struct {
	// Description Description of the field
	Description string `json:"description"`
}

// TemplateListResponse defines model for TemplateListResponse.
type TemplateListResponse struct {
	// Data List of templates
	Data []struct {
		// ApplicableTo List of resources the template is applicable to
		ApplicableTo []TemplateListResponseDataApplicableTo `json:"applicable_to"`

		// Description Description of the template
		Description string `json:"description"`

		// ID Unique identifier of the template
		ID string `json:"id"`

		// IsActive Active status of the template
		IsActive bool `json:"is_active"`

		// Links Meta links
		Links struct {
			// Self URL of the template
			Self string `json:"self"`
		} `json:"links"`

		// Title Title of the template
		Title string `json:"title"`
	} `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// TemplateListResponseDataApplicableTo defines model for TemplateListResponse.Data.ApplicableTo.
type TemplateListResponseDataApplicableTo string

// TemplateLogicRule List of logic rules
type TemplateLogicRule = []TemplateLogicRuleItem

// TemplateLogicRuleItem defines model for .
type TemplateLogicRuleItem struct {
	// Action Action to perform for the rule
	Action TemplateLogicRuleAction `json:"action"`

	// Condition Condition of the rule
	Condition TemplateLogicRuleCondition `json:"condition"`
}

// TemplateLogicRuleAction Action to perform for the rule
type TemplateLogicRuleAction struct {
	// TargetField Target field for the action
	TargetField TemplateLogicRuleActionTargetField `json:"target_field"`

	// Type Type of the action
	Type TemplateLogicRuleActionType `json:"type"`

	// Values Values for the action
	Values []string `json:"values"`
}

// TemplateLogicRuleActionTargetField Target field for the action
type TemplateLogicRuleActionTargetField struct {
	// Identifier Identifier of the target field
	Identifier string `json:"identifier"`
}

// TemplateLogicRuleActionType Type of the action
type TemplateLogicRuleActionType string

// TemplateLogicRuleCondition Condition of the rule
type TemplateLogicRuleCondition struct {
	// Operator Operator of the condition
	Operator TemplateLogicRuleConditionOperator `json:"operator"`

	// Values Values of the condition
	Values []string `json:"values"`
}

// TemplateLogicRuleConditionOperator Operator of the condition
type TemplateLogicRuleConditionOperator string

// TemplatePercentageNumberFormField defines model for TemplatePercentageNumberFormField.
type TemplatePercentageNumberFormField struct {
	// HelperText Helper text of the field
	HelperText *string `json:"helper_text"`

	// Identifier Identifier of the field
	Identifier string `json:"identifier"`

	// IsProtected True when this form field is read-only and only organization admins can modify it. Absent or false for editable fields.
	IsProtected *bool `json:"is_protected,omitempty"`

	// IsReadonly Specify if the field can be used to create any resource
	IsReadonly bool `json:"is_readonly"`

	// IsRequired Required status of the field
	IsRequired bool `json:"is_required"`

	// Label Label of the field
	Label *string `json:"label"`

	// LogicRules List of logic rules
	LogicRules TemplateLogicRule `json:"logic_rules"`

	// SortOrder Sort order of the field
	SortOrder int `json:"sort_order"`

	// Type Type of the field
	Type TemplatePercentageNumberFormFieldType `json:"type"`

	// TypeSpecificMeta Type-specific meta-information
	TypeSpecificMeta TemplatePercentageNumberFormFieldTypeSpecificMeta `json:"type_specific_meta"`
}

// TemplatePercentageNumberFormFieldType Type of the field
type TemplatePercentageNumberFormFieldType string

// TemplatePercentageNumberFormFieldTypeSpecificMeta Type-specific meta-information
type TemplatePercentageNumberFormFieldTypeSpecificMeta struct {
	// DecimalPlaces Decimal places allowed for the field
	DecimalPlaces *int `json:"decimal_places"`
}

// TemplateResponse defines model for TemplateResponse.
type TemplateResponse struct {
	// ApplicableTo List of resources the template is applicable to
	ApplicableTo []TemplateResponseApplicableTo `json:"applicable_to"`

	// Description Description of the template
	Description string `json:"description"`

	// FormFields List of form fields
	FormFields []TemplateResponse_FormFields_Item `json:"form_fields"`

	// ID Unique identifier of the template
	ID string `json:"id"`

	// IsActive Active status of the template
	IsActive bool `json:"is_active"`

	// Links Meta links
	Links TemplateResponseLinks `json:"links"`

	// Title Title of the template
	Title string `json:"title"`
}

// TemplateResponseApplicableTo defines model for TemplateResponse.ApplicableTo.
type TemplateResponseApplicableTo string

// TemplateResponse_FormFields_Item defines model for TemplateResponse.form_fields.Item.
type TemplateResponse_FormFields_Item struct {
	union json.RawMessage
}

// TemplateResponseLinks Meta links
type TemplateResponseLinks struct {
	// Self URL of the template
	Self string `json:"self"`
}

// TemplateSimpleNumberFormField defines model for TemplateSimpleNumberFormField.
type TemplateSimpleNumberFormField struct {
	// HelperText Helper text of the field
	HelperText *string `json:"helper_text"`

	// Identifier Identifier of the field
	Identifier string `json:"identifier"`

	// IsProtected True when this form field is read-only and only organization admins can modify it. Absent or false for editable fields.
	IsProtected *bool `json:"is_protected,omitempty"`

	// IsReadonly Specify if the field can be used to create any resource
	IsReadonly bool `json:"is_readonly"`

	// IsRequired Required status of the field
	IsRequired bool `json:"is_required"`

	// Label Label of the field
	Label *string `json:"label"`

	// LogicRules List of logic rules
	LogicRules TemplateLogicRule `json:"logic_rules"`

	// SortOrder Sort order of the field
	SortOrder int `json:"sort_order"`

	// Type Type of the field
	Type TemplateSimpleNumberFormFieldType `json:"type"`

	// TypeSpecificMeta Type-specific meta-information
	TypeSpecificMeta TemplateSimpleNumberFormFieldTypeSpecificMeta `json:"type_specific_meta"`
}

// TemplateSimpleNumberFormFieldType Type of the field
type TemplateSimpleNumberFormFieldType string

// TemplateSimpleNumberFormFieldTypeSpecificMeta Type-specific meta-information
type TemplateSimpleNumberFormFieldTypeSpecificMeta struct {
	// DecimalPlaces Decimal places allowed for the field
	DecimalPlaces *int `json:"decimal_places"`

	// HasThousandSeparator Whether the field has thousand separators
	HasThousandSeparator bool `json:"has_thousand_separator"`
}

// TextFieldDefinition defines model for TextFieldDefinition.
type TextFieldDefinition struct {
	Core CoreFieldDef `json:"core"`

	// DefaultValues Default values of the field
	DefaultValues *[]string `json:"default_values,omitempty"`

	// MaxLength Maximum length of the field
	MaxLength *int `json:"max_length,omitempty"`

	// MinLength Minimum length of the field
	MinLength int `json:"min_length"`

	// ValidationPattern Validation pattern of the field
	ValidationPattern string `json:"validation_pattern"`
}

// TextFieldValueModel defines model for TextFieldValueModel.
type TextFieldValueModel struct {
	// OrderIndex Order index of the field value
	OrderIndex *int `json:"order_index,omitempty"`

	// TextValue Value of the field value
	TextValue string `json:"text_value"`
}

// TextTypeFormFieldRequest defines model for TextTypeFormFieldRequest.
type TextTypeFormFieldRequest struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// Type Type of the form field
	Type   string   `json:"type"`
	Values []string `json:"values"`
}

// TextTypeFormFieldResponse defines model for TextTypeFormFieldResponse.
type TextTypeFormFieldResponse struct {
	// Identifier Identifier of the form field collected from the template form field
	Identifier string `json:"identifier"`

	// IsProtected True when this form field is read-only and only organization admins can modify it. Absent or false for editable fields.
	IsProtected *bool `json:"is_protected,omitempty"`

	// Type Type of the form field
	Type   string   `json:"type"`
	Values []string `json:"values"`
}

// URLFieldValueModel defines model for URLFieldValueModel.
type URLFieldValueModel struct {
	// OrderIndex Order index of the field value
	OrderIndex *int `json:"order_index,omitempty"`

	// URL URL of the field value
	URL string `json:"url"`
}

// UpdateAssetFieldsResponse defines model for UpdateAssetFieldsResponse.
type UpdateAssetFieldsResponse struct {
	// Data List of fields
	Data []AssetFieldListResponseItem `json:"data"`

	// Links Meta links
	Links UpdateAssetFieldsResponseLinks `json:"links"`
}

// UpdateAssetFieldsResponseLinks Meta links
type UpdateAssetFieldsResponseLinks struct {
	// AssetFields URL to get the asset fields
	AssetFields string `json:"asset_fields"`

	// AssetURL URL of the asset
	AssetURL string `json:"asset_url"`
}

// UpdateSCContentTypeManagedMigrationRequest Default values for the managed migration job content.
type UpdateSCContentTypeManagedMigrationRequest struct {
	DefaultValues LocalizedFieldValues `json:"default_values"`
}

// UpdateSCContentTypeManagedMigrationResponse defines model for UpdateSCContentTypeManagedMigrationResponse.
type UpdateSCContentTypeManagedMigrationResponse struct {
	// Updated True if the managed migration job is updated.
	Updated *bool `json:"updated,omitempty"`
}

// UserListResponse defines model for UserListResponse.
type UserListResponse struct {
	// Data List of users
	Data []struct {
		// FirstName First name of the user
		FirstName string `json:"first_name"`

		// FullName Full name of the user
		FullName string `json:"full_name"`

		// ID Unique identifier of the user
		ID string `json:"id"`

		// ImageURL URL of the profile picture of the user - null if the user has not provided an image
		ImageURL *string `json:"image_url"`

		// LastName Last name of the user
		LastName string `json:"last_name"`

		// Links Meta links
		Links struct {
			// Self URL of the user
			Self string `json:"self"`
		} `json:"links"`

		// Roles List of roles assigned to the user
		Roles []UserListResponseItemRoles `json:"roles"`
	} `json:"data"`
	Pagination struct {
		Next *string `json:"next"`

		// Previous URL to the previous page
		Previous *string `json:"previous"`
	} `json:"pagination"`
}

// UserListResponseItemRoles defines model for .
type UserListResponseItemRoles struct {
	// Name Name of the role
	Name string `json:"name"`
}

// UserListResponseItem defines model for UserListResponseItem.
type UserListResponseItem struct {
	// FirstName First name of the user
	FirstName string `json:"first_name"`

	// FullName Full name of the user
	FullName string `json:"full_name"`

	// ID Unique identifier of the user
	ID string `json:"id"`

	// ImageURL URL of the profile picture of the user - null if the user has not provided an image
	ImageURL *string `json:"image_url"`

	// LastName Last name of the user
	LastName string `json:"last_name"`

	// Links Meta links
	Links UserListResponseItemLinks `json:"links"`

	// Roles List of roles assigned to the user
	Roles []UserListResponseItemRoles `json:"roles"`
}

// UserListResponseItemLinks Meta links
type UserListResponseItemLinks struct {
	// Self URL of the user
	Self string `json:"self"`
}

// UserResponse defines model for UserResponse.
type UserResponse struct {
	// Email Email of the user – null unless the application is configured to expose the user email
	Email *string `json:"email"`

	// FirstName First name of the user
	FirstName string `json:"first_name"`

	// FullName Full name of the user
	FullName string `json:"full_name"`

	// ID Unique identifier of the user
	ID string `json:"id"`

	// ImageURL URL of the profile picture of the user – null if the user has not provided an image
	ImageURL *string `json:"image_url"`

	// LastName Last name of the user
	LastName string `json:"last_name"`

	// Links Meta links
	Links *UserResponseLinks `json:"links,omitempty"`
}

// UserResponseLinks Meta links
type UserResponseLinks struct {
	// Self URL of the user
	Self *string `json:"self,omitempty"`
}

// ValidateSCContentTypeManagedMigrationResponse defines model for ValidateSCContentTypeManagedMigrationResponse.
type ValidateSCContentTypeManagedMigrationResponse struct {
	// IsManagedMigrationPossible False if the managed migration not possible
	IsManagedMigrationPossible bool `json:"is_managed_migration_possible"`
}

// ValidationError defines model for ValidationError.
type ValidationError struct {
	// Loc Locations where the error occurred
	Loc []string `json:"loc"`

	// Msg Message describing the error
	Msg string `json:"msg"`

	// Type Type of the error
	Type string `json:"type"`
}

// VersionedContentTypeModel defines model for VersionedContentTypeModel.
type VersionedContentTypeModel struct {
	// Component Indicates whether the content type is a component
	Component bool `json:"component"`

	// ContentTypeGUID Unique identifier of the content type
	ContentTypeGUID string `json:"content_type_guid"`

	// CreatedAt Date and time on which the content type was created, in ISO 8601 UTC format
	CreatedAt time.Time `json:"created_at"`

	// CreatedBy Unique identifier of the user who created the content type
	CreatedBy string `json:"created_by"`

	// Description Description of the content type
	Description *string `json:"description,omitempty"`

	// Disabled Disabled status of the content type
	Disabled *bool `json:"disabled,omitempty"`

	// Links Meta links
	Links VersionedContentTypeModelLinks `json:"links"`

	// Name Name of the content type
	Name string `json:"name"`

	// Source Source of the content type
	Source *string `json:"source,omitempty"`

	// SourceID Source of the content type
	SourceID *string `json:"source_id,omitempty"`

	// SourceMetadata Source metadata of the content type
	SourceMetadata *string `json:"source_metadata,omitempty"`

	// ThumbnailGUID Thumbnail GUID of the content type
	ThumbnailGUID *string `json:"thumbnail_guid,omitempty"`

	// UpdatedAt Date and time on which the content type was last updated, in ISO 8601 UTC format
	UpdatedAt time.Time `json:"updated_at"`

	// UpdatedBy Unique identifier of the user who last updated the content type
	UpdatedBy string               `json:"updated_by"`
	Version   SCContentTypeVersion `json:"version"`
}

// VersionedContentTypeModelLinks Meta links
type VersionedContentTypeModelLinks struct {
	// Self URL of the content type
	Self *string `json:"self,omitempty"`

	// Versions URL of the content type versions
	Versions *string `json:"versions,omitempty"`
}

// WorkRequestApprovedAssetResponse defines model for WorkRequestApprovedAssetResponse.
type WorkRequestApprovedAssetResponse struct {
	// ContentType Content of the approved asset
	ContentType WorkRequestApprovedAssetResponseContentType `json:"content_type"`

	// ID Unique identifier of the work request's approved asset
	ID string `json:"id"`

	// MimeType Mime type of the work request's approved asset
	MimeType string `json:"mime_type"`

	// Title Title of the work request's approved asset
	Title string `json:"title"`

	// Type Type of the work request's approved asset
	Type WorkRequestApprovedAssetResponseType `json:"type"`
}

// WorkRequestApprovedAssetResponseContentTypeType Type of the content. 'html_body for 'article' and 'url' for `image`, `raw file`, and `video`
type WorkRequestApprovedAssetResponseContentTypeType string

// WorkRequestApprovedAssetResponseContentType Content of the approved asset
type WorkRequestApprovedAssetResponseContentType struct {
	// Type Type of the content. 'html_body for 'article' and 'url' for `image`, `raw file`, and `video`
	Type WorkRequestApprovedAssetResponseContentTypeType `json:"type"`

	// Value Content of the asset. 'html_body for 'article' and 'url' for `image`, `raw file`, and `video`
	Value string `json:"value"`
}

// WorkRequestApprovedAssetResponseType Type of the work request's approved asset
type WorkRequestApprovedAssetResponseType string

// WorkRequestAttachmentBriefRequest Brief from uploaded attachment
type WorkRequestAttachmentBriefRequest struct {
	Type  WorkRequestAttachmentBriefRequestType  `json:"type"`
	Value WorkRequestAttachmentBriefRequestValue `json:"value"`
}

// WorkRequestAttachmentBriefRequestType defines model for WorkRequestAttachmentBriefRequest.Type.
type WorkRequestAttachmentBriefRequestType string

// WorkRequestAttachmentBriefRequestValue defines model for .
type WorkRequestAttachmentBriefRequestValue struct {
	// Key Unique identifier of the file upload session. This should be the value of `upload_meta_fields.key` retrieved using the `/v3/upload-url` endpoint
	Key string `json:"key"`

	// Name Name of the attachment brief
	Name string `json:"name"`
}

// WorkRequestAttachmentBriefResponse Brief from uploaded attachment
type WorkRequestAttachmentBriefResponse struct {
	Type  WorkRequestAttachmentBriefResponseType  `json:"type"`
	Value WorkRequestAttachmentBriefResponseValue `json:"value"`
}

// WorkRequestAttachmentBriefResponseType defines model for WorkRequestAttachmentBriefResponse.Type.
type WorkRequestAttachmentBriefResponseType string

// WorkRequestAttachmentBriefResponseValue defines model for .
type WorkRequestAttachmentBriefResponseValue struct {
	// ID Unique identifier of the attachment brief
	ID string `json:"id"`

	// Name Name of the attachment brief
	Name string `json:"name"`

	// URL URL of the attachment brief
	URL string `json:"url"`
}

// WorkRequestCampaignRequest defines model for WorkRequestCampaignRequest.
type WorkRequestCampaignRequest struct {
	// Description Description of the campaign
	Description *string `json:"description"`

	// EndDate End date of the campaign, in ISO 8601 UTC format (2022-08-31T06:00:00Z)
	EndDate *time.Time `json:"end_date"`

	// OwnerID Unique identifier of the campaign owner
	OwnerID *string `json:"owner_id,omitempty"`

	// ParentCampaignID Unique identifier of a campaign
	ParentCampaignID *string `json:"parent_campaign_id,omitempty"`

	// StartDate Start date of the campaign, in ISO 8601 UTC format (2022-08-31T06:00:00Z)
	StartDate *time.Time `json:"start_date"`

	// Title Title of the campaign
	Title string `json:"title"`
}

// WorkRequestCampaignResponse defines model for WorkRequestCampaignResponse.
type WorkRequestCampaignResponse struct {
	// Description Description of the campaign
	Description *string `json:"description"`

	// EndDate End date of the campaign, in ISO 8601 UTC format (2022-08-31T06:00:00Z)
	EndDate *time.Time `json:"end_date"`

	// ID Unique identifier of the campaign
	ID string `json:"id"`

	// Links Meta links
	Links WorkRequestCampaignResponseLinks `json:"links"`

	// OwnerID Unique identifier of the campaign owner
	OwnerID string `json:"owner_id"`

	// ParentCampaignID Unique identifier of the parent campaign
	ParentCampaignID string `json:"parent_campaign_id"`

	// ReferenceID Reference identifier of the campaign
	ReferenceID string `json:"reference_id"`

	// StartDate Start date of the campaign, in ISO 8601 UTC format (2022-08-31T06:00:00Z)
	StartDate *time.Time `json:"start_date"`

	// Title Title of the campaign
	Title string `json:"title"`
}

// WorkRequestCampaignResponseLinks Meta links
type WorkRequestCampaignResponseLinks struct {
	// Brief URL of the campaign brief
	Brief *string `json:"brief"`

	// ParentCampaign URL of the parent campaign
	ParentCampaign *string `json:"parent_campaign"`

	// Self URL of the campaign
	Self string `json:"self"`
}

// WorkRequestCommentResponse defines model for WorkRequestCommentResponse.
type WorkRequestCommentResponse struct {
	// Attachments List of attachments of the comment
	Attachments []AttachmentResponse `json:"attachments"`

	// CommentBy Unique identifier of the user who posted the comment
	CommentBy string `json:"comment_by"`

	// CreatedAt Creation date and time of the work request comment, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	CreatedAt time.Time `json:"created_at"`

	// ID Unique identifier of the work request comment
	ID string `json:"id"`

	// IsResolved Determines if the comment is resolved
	IsResolved bool `json:"is_resolved"`

	// Links Meta links
	Links WorkRequestCommentResponseLinks `json:"links"`

	// ModifiedAt Last modification date and time of the work request comment, in ISO 8601 UTC format (2020-02-10T10:40:45Z)
	ModifiedAt time.Time `json:"modified_at"`

	// ParentCommentID Unique identifier of the parent comment
	ParentCommentID *string `json:"parent_comment_id"`

	// Value Content of the comment
	Value string `json:"value"`
}

// WorkRequestCommentResponseLinks Meta links
type WorkRequestCommentResponseLinks struct {
	// CommentBy URL of the user who posted the comment
	CommentBy string `json:"comment_by"`

	// WorkRequest URL of the work request
	WorkRequest string `json:"work_request"`
}

// WorkRequestCreateRequest defines model for WorkRequestCreateRequest.
type WorkRequestCreateRequest struct {
	// Assignees List of assigned users to the work request
	Assignees *[]string `json:"assignees,omitempty"`

	// FormFields Form fields from the work request template to populate in the work request. **NOTE** You must provide values for all the required form fields.
	FormFields []WorkRequestCreateRequest_FormFields_Item `json:"form_fields"`

	// TemplateID The template on which the work request should be created
	TemplateID string `json:"template_id"`
}

// WorkRequestCreateRequest_FormFields_Item defines model for WorkRequestCreateRequest.form_fields.Item.
type WorkRequestCreateRequest_FormFields_Item struct {
	union json.RawMessage
}

// WorkRequestFormFieldUpdateRequest defines model for WorkRequestFormFieldUpdateRequest.
type WorkRequestFormFieldUpdateRequest struct {
	union json.RawMessage
}

// WorkRequestRelatedResourceResponse defines model for WorkRequestRelatedResourceResponse.
type WorkRequestRelatedResourceResponse struct {
	// CreatedAt Date and time on which the work request related resource was created, in ISO 8601 UTC format (2022-04-30T06:00:00.000Z)
	CreatedAt time.Time `json:"created_at"`

	// ID Unique identifier of the related resource
	ID string `json:"id"`

	// Links Meta links
	Links WorkRequestRelatedResourceResponseLinks `json:"links"`

	// RelationType Related resource relation type
	RelationType WorkRequestRelatedResourceResponseRelationType `json:"relation_type"`

	// ResourceType Related resource type. task/campaign/event etc
	ResourceType string `json:"resource_type"`
}

// WorkRequestRelatedResourceResponseLinks Meta links
type WorkRequestRelatedResourceResponseLinks struct {
	// Self URL of the related resource
	Self *string `json:"self"`
}

// WorkRequestRelatedResourceResponseRelationType Related resource relation type
type WorkRequestRelatedResourceResponseRelationType string

// WorkRequestRequestFormFieldBriefTypePayload defines model for WorkRequestRequestFormFieldBriefTypePayload.
type WorkRequestRequestFormFieldBriefTypePayload struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldBriefTypePayloadType `json:"type"`

	// Values Brief value payload
	Values []WorkRequestRequestFormFieldBriefTypePayload_Values_Item `json:"values"`
}

// WorkRequestRequestFormFieldBriefTypePayloadType Type of the form field
type WorkRequestRequestFormFieldBriefTypePayloadType string

// WorkRequestRequestFormFieldBriefTypePayload_Values_Item defines model for WorkRequestRequestFormFieldBriefTypePayload.values.Item.
type WorkRequestRequestFormFieldBriefTypePayload_Values_Item struct {
	union json.RawMessage
}

// WorkRequestRequestFormFieldBriefTypeResponse defines model for WorkRequestRequestFormFieldBriefTypeResponse.
type WorkRequestRequestFormFieldBriefTypeResponse struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldBriefTypeResponseType `json:"type"`

	// Values Brief value payload
	Values []WorkRequestRequestFormFieldBriefTypeResponse_Values_Item `json:"values"`
}

// WorkRequestRequestFormFieldBriefTypeResponseType Type of the form field
type WorkRequestRequestFormFieldBriefTypeResponseType string

// WorkRequestRequestFormFieldBriefTypeResponse_Values_Item defines model for WorkRequestRequestFormFieldBriefTypeResponse.values.Item.
type WorkRequestRequestFormFieldBriefTypeResponse_Values_Item struct {
	union json.RawMessage
}

// WorkRequestRequestFormFieldCheckboxTypePayload defines model for WorkRequestRequestFormFieldCheckboxTypePayload.
type WorkRequestRequestFormFieldCheckboxTypePayload struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldCheckboxTypePayloadType `json:"type"`

	// Values Unique identifiers of the choices
	Values []string `json:"values"`
}

// WorkRequestRequestFormFieldCheckboxTypePayloadType Type of the form field
type WorkRequestRequestFormFieldCheckboxTypePayloadType string

// WorkRequestRequestFormFieldCheckboxTypeResponse defines model for WorkRequestRequestFormFieldCheckboxTypeResponse.
type WorkRequestRequestFormFieldCheckboxTypeResponse struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldCheckboxTypeResponseType `json:"type"`

	// Values Values for the form field
	Values []WorkRequestRequestFormFieldCheckboxTypeResponseValues `json:"values"`
}

// WorkRequestRequestFormFieldCheckboxTypeResponseType Type of the form field
type WorkRequestRequestFormFieldCheckboxTypeResponseType string

// WorkRequestRequestFormFieldCheckboxTypeResponseValues defines model for .
type WorkRequestRequestFormFieldCheckboxTypeResponseValues struct {
	// ID Unique identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// WorkRequestRequestFormFieldCurrencyNumberTypePayload defines model for WorkRequestRequestFormFieldCurrencyNumberTypePayload.
type WorkRequestRequestFormFieldCurrencyNumberTypePayload struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldCurrencyNumberTypePayloadType `json:"type"`

	// Values Currency amount for the form field
	Values []string `json:"values"`
}

// WorkRequestRequestFormFieldCurrencyNumberTypePayloadType Type of the form field
type WorkRequestRequestFormFieldCurrencyNumberTypePayloadType string

// WorkRequestRequestFormFieldDateTypePayload defines model for WorkRequestRequestFormFieldDateTypePayload.
type WorkRequestRequestFormFieldDateTypePayload struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldDateTypePayloadType `json:"type"`

	// Values Date and time, in ISO 8601 UTC format
	Values []string `json:"values"`
}

// WorkRequestRequestFormFieldDateTypePayloadType Type of the form field
type WorkRequestRequestFormFieldDateTypePayloadType string

// WorkRequestRequestFormFieldDropdownTypePayload defines model for WorkRequestRequestFormFieldDropdownTypePayload.
type WorkRequestRequestFormFieldDropdownTypePayload struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldDropdownTypePayloadType `json:"type"`

	// Values Unique identifier of the choice for the form field
	Values []string `json:"values"`
}

// WorkRequestRequestFormFieldDropdownTypePayloadType Type of the form field
type WorkRequestRequestFormFieldDropdownTypePayloadType string

// WorkRequestRequestFormFieldDropdownTypeResponse defines model for WorkRequestRequestFormFieldDropdownTypeResponse.
type WorkRequestRequestFormFieldDropdownTypeResponse struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldDropdownTypeResponseType `json:"type"`

	// Values Value for the form field
	Values []WorkRequestRequestFormFieldDropdownTypeResponseValues `json:"values"`
}

// WorkRequestRequestFormFieldDropdownTypeResponseType Type of the form field
type WorkRequestRequestFormFieldDropdownTypeResponseType string

// WorkRequestRequestFormFieldDropdownTypeResponseValues defines model for .
type WorkRequestRequestFormFieldDropdownTypeResponseValues struct {
	// ID Unique identifier of the choice
	ID string `json:"id"`

	// Name Name of the choice
	Name string `json:"name"`
}

// WorkRequestRequestFormFieldFileTypePayload defines model for WorkRequestRequestFormFieldFileTypePayload.
type WorkRequestRequestFormFieldFileTypePayload struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldFileTypePayloadType `json:"type"`

	// Values uploaded file values
	Values []WorkRequestRequestFormFieldFileTypePayloadValues `json:"values"`
}

// WorkRequestRequestFormFieldFileTypePayloadType Type of the form field
type WorkRequestRequestFormFieldFileTypePayloadType string

// WorkRequestRequestFormFieldFileTypePayloadValues defines model for .
type WorkRequestRequestFormFieldFileTypePayloadValues struct {
	// Key Unique identifier of the file upload session. This is the `upload_meta_fields.key` field was retrieved from the `/v3/upload-url` endpoint.
	Key string `json:"key"`

	// Name Name for the file
	Name string `json:"name"`
}

// WorkRequestRequestFormFieldFileTypeResponse defines model for WorkRequestRequestFormFieldFileTypeResponse.
type WorkRequestRequestFormFieldFileTypeResponse struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldFileTypeResponseType `json:"type"`

	// Values Values for the form field
	Values []WorkRequestRequestFormFieldFileTypeResponseValues `json:"values"`
}

// WorkRequestRequestFormFieldFileTypeResponseType Type of the form field
type WorkRequestRequestFormFieldFileTypeResponseType string

// WorkRequestRequestFormFieldFileTypeResponseValues defines model for .
type WorkRequestRequestFormFieldFileTypeResponseValues struct {
	// ID Unique identifier of the file
	ID string `json:"id"`

	// Name Name for the file
	Name string `json:"name"`

	// URL Download URL of the file
	URL string `json:"url"`
}

// WorkRequestRequestFormFieldLabelTypePayload defines model for WorkRequestRequestFormFieldLabelTypePayload.
type WorkRequestRequestFormFieldLabelTypePayload struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldLabelTypePayloadType `json:"type"`

	// Values Unique identifiers of label values for the form field
	Values []string `json:"values"`
}

// WorkRequestRequestFormFieldLabelTypePayloadType Type of the form field
type WorkRequestRequestFormFieldLabelTypePayloadType string

// WorkRequestRequestFormFieldLabelTypeResponse defines model for WorkRequestRequestFormFieldLabelTypeResponse.
type WorkRequestRequestFormFieldLabelTypeResponse struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldLabelTypeResponseType `json:"type"`

	// Values Values for the form field
	Values []WorkRequestRequestFormFieldLabelTypeResponseValues `json:"values"`
}

// WorkRequestRequestFormFieldLabelTypeResponseType Type of the form field
type WorkRequestRequestFormFieldLabelTypeResponseType string

// WorkRequestRequestFormFieldLabelTypeResponseValues defines model for .
type WorkRequestRequestFormFieldLabelTypeResponseValues struct {
	// ID Unique identifiers of label value
	ID string `json:"id"`

	// Name Name of the label value
	Name string `json:"name"`
}

// WorkRequestRequestFormFieldPercentageNumberTypePayload defines model for WorkRequestRequestFormFieldPercentageNumberTypePayload.
type WorkRequestRequestFormFieldPercentageNumberTypePayload struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldPercentageNumberTypePayloadType `json:"type"`

	// Values Value for the form field
	Values []string `json:"values"`
}

// WorkRequestRequestFormFieldPercentageNumberTypePayloadType Type of the form field
type WorkRequestRequestFormFieldPercentageNumberTypePayloadType string

// WorkRequestRequestFormFieldRadioButtonTypePayload defines model for WorkRequestRequestFormFieldRadioButtonTypePayload.
type WorkRequestRequestFormFieldRadioButtonTypePayload struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldRadioButtonTypePayloadType `json:"type"`

	// Values Unique identifier of the choice for the form field
	Values []string `json:"values"`
}

// WorkRequestRequestFormFieldRadioButtonTypePayloadType Type of the form field
type WorkRequestRequestFormFieldRadioButtonTypePayloadType string

// WorkRequestRequestFormFieldRadioButtonTypeResponse defines model for WorkRequestRequestFormFieldRadioButtonTypeResponse.
type WorkRequestRequestFormFieldRadioButtonTypeResponse struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldRadioButtonTypeResponseType `json:"type"`

	// Values Value for the form field
	Values []WorkRequestRequestFormFieldRadioButtonTypeResponseValues `json:"values"`
}

// WorkRequestRequestFormFieldRadioButtonTypeResponseType Type of the form field
type WorkRequestRequestFormFieldRadioButtonTypeResponseType string

// WorkRequestRequestFormFieldRadioButtonTypeResponseValues defines model for .
type WorkRequestRequestFormFieldRadioButtonTypeResponseValues struct {
	// ID Unique identifier of the choice
	ID *string `json:"id,omitempty"`

	// Name Name of the choice value
	Name *string `json:"name,omitempty"`
}

// WorkRequestRequestFormFieldRichtextTypePayload defines model for WorkRequestRequestFormFieldRichtextTypePayload.
type WorkRequestRequestFormFieldRichtextTypePayload struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldRichtextTypePayloadType `json:"type"`

	// Values Value of the form field
	Values []string `json:"values"`
}

// WorkRequestRequestFormFieldRichtextTypePayloadType Type of the form field
type WorkRequestRequestFormFieldRichtextTypePayloadType string

// WorkRequestRequestFormFieldSimpleNumberTypePayload defines model for WorkRequestRequestFormFieldSimpleNumberTypePayload.
type WorkRequestRequestFormFieldSimpleNumberTypePayload struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldSimpleNumberTypePayloadType `json:"type"`

	// Values Value of the form field
	Values []string `json:"values"`
}

// WorkRequestRequestFormFieldSimpleNumberTypePayloadType Type of the form field
type WorkRequestRequestFormFieldSimpleNumberTypePayloadType string

// WorkRequestRequestFormFieldTextAreaTypePayload defines model for WorkRequestRequestFormFieldTextAreaTypePayload.
type WorkRequestRequestFormFieldTextAreaTypePayload struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldTextAreaTypePayloadType `json:"type"`

	// Values Value of the form field
	Values []string `json:"values"`
}

// WorkRequestRequestFormFieldTextAreaTypePayloadType Type of the form field
type WorkRequestRequestFormFieldTextAreaTypePayloadType string

// WorkRequestRequestFormFieldTextTypePayload defines model for WorkRequestRequestFormFieldTextTypePayload.
type WorkRequestRequestFormFieldTextTypePayload struct {
	// Type Type of the form field
	Type WorkRequestRequestFormFieldTextTypePayloadType `json:"type"`

	// Values Value of the form field
	Values []string `json:"values"`
}

// WorkRequestRequestFormFieldTextTypePayloadType Type of the form field
type WorkRequestRequestFormFieldTextTypePayloadType string

// WorkRequestRequestFormFieldUpdateResponse defines model for WorkRequestRequestFormFieldUpdateResponse.
type WorkRequestRequestFormFieldUpdateResponse struct {
	union json.RawMessage
}

// WorkRequestResponse defines model for WorkRequestResponse.
type WorkRequestResponse struct {
	// Assignees List of assigned users to the work request
	Assignees []WorkRequestResponseAssignees `json:"assignees"`

	// CreatedAt Date and time on which the work request was created, in ISO 8601 UTC format (2022-04-30T06:00:00.000Z)
	CreatedAt time.Time `json:"created_at"`

	// CreatedBy Creator of the work request
	CreatedBy WorkRequestResponseCreatedBy `json:"created_by"`

	// FormFields Populated form fields following the template
	FormFields []WorkRequestResponse_FormFields_Item `json:"form_fields"`

	// ID Unique identifier of the work request
	ID string `json:"id"`

	// Links Meta links
	Links WorkRequestResponseLinks `json:"links"`

	// ModifiedAt Date and time on which the work request was last modified, in ISO 8601 UTC format (2022-04-30T08:00:00.000Z)
	ModifiedAt time.Time `json:"modified_at"`

	// Priority Priority of the work request
	Priority WorkRequestResponsePriority `json:"priority"`

	// ReferenceID Reference identifier of the work request
	ReferenceID string `json:"reference_id"`

	// Status Current status of the work request
	Status WorkRequestResponseStatus `json:"status"`

	// Template The template, based on the work request that was created
	Template WorkRequestResponseTemplate `json:"template"`
}

// WorkRequestResponseAssigneesLinks Meta links
type WorkRequestResponseAssigneesLinks struct {
	// Self URL of the assigned user
	Self string `json:"self"`
}

// WorkRequestResponseAssigneesType Type of the assignee
type WorkRequestResponseAssigneesType string

// WorkRequestResponseAssignees defines model for .
type WorkRequestResponseAssignees struct {
	// FullName Full name of the assignee
	FullName string `json:"full_name"`

	// ID Unique identifier of the assignee
	ID string `json:"id"`

	// ImageURL URL of the profile picture of the assignee. `null` if the assignee has not provided an image
	ImageURL *string `json:"image_url"`

	// Links Meta links
	Links WorkRequestResponseAssigneesLinks `json:"links"`

	// Type Type of the assignee
	Type WorkRequestResponseAssigneesType `json:"type"`
}

// WorkRequestResponseCreatedByLinks Meta links
type WorkRequestResponseCreatedByLinks struct {
	// Self URL of the creator
	Self string `json:"self"`
}

// WorkRequestResponseCreatedBy Creator of the work request
type WorkRequestResponseCreatedBy struct {
	// ID Unique identifier of the creator of the work request
	ID string `json:"id"`

	// Links Meta links
	Links WorkRequestResponseCreatedByLinks `json:"links"`
}

// WorkRequestResponse_FormFields_Item defines model for WorkRequestResponse.form_fields.Item.
type WorkRequestResponse_FormFields_Item struct {
	union json.RawMessage
}

// WorkRequestResponseLinks Meta links
type WorkRequestResponseLinks struct {
	// Attachments URL to add attachments to work request
	Attachments string `json:"attachments"`

	// Campaigns URL to create a campaign from the work request
	Campaigns string `json:"campaigns"`

	// Comments URL to create or get work request comments
	Comments string `json:"comments"`

	// CreativeAssets URL to add creative assets to the work request
	CreativeAssets string `json:"creative_assets"`

	// RelatedResources URL to get related resources
	RelatedResources string `json:"related_resources"`

	// Self URL of the work request
	Self string `json:"self"`

	// Tasks URL to create a task from the work request
	Tasks string `json:"tasks"`
}

// WorkRequestResponsePriority Priority of the work request
type WorkRequestResponsePriority string

// WorkRequestResponseStatus Current status of the work request
type WorkRequestResponseStatus string

// WorkRequestResponseTemplateLinks Meta links
type WorkRequestResponseTemplateLinks struct {
	// Self URL of the template
	Self string `json:"self"`
}

// WorkRequestResponseTemplate The template, based on the work request that was created
type WorkRequestResponseTemplate struct {
	// ID Unique identifier of the template
	ID string `json:"id"`

	// Links Meta links
	Links WorkRequestResponseTemplateLinks `json:"links"`

	// Title Title of the template
	Title string `json:"title"`
}

// WorkRequestTaskRequest defines model for WorkRequestTaskRequest.
type WorkRequestTaskRequest struct {
	// AssetIdsWithoutFieldInheritance List of asset ids which will not inherit any fields.
	AssetIdsWithoutFieldInheritance *[]string `json:"asset_ids_without_field_inheritance,omitempty"`

	// CampaignID Unique identifier of parent campaign
	CampaignID *string `json:"campaign_id"`

	// DueAt Due date and time of the task, in ISO 8601 UTC format (2020-02-10T10:40:45Z). The `due_at` field is required if `workflow_id` is present and the workflow does not have step duration. If `workflow_id` is present and the workflow has step duration, then either `start_at` or `due_at` is required.
	DueAt *time.Time `json:"due_at"`

	// InheritFieldsFrom Specifies the source from which to inherit fields.
	//
	// Supported values:
	// - `work_request` (default) : Fields are created from the work request
	// - `workflow` : Fields are inherited from the workflow specified by `workflow_id`
	InheritFieldsFrom *string `json:"inherit_fields_from,omitempty"`

	// OwnerID Unique identifier of the task owner. If passed as null, then the owner will be the token user.
	OwnerID *string `json:"owner_id"`

	// StartAt Start date and time of the task, in ISO 8601 UTC format (2020-02-10T10:40:45Z). If `workflow_id` is present and the workflow has step duration, then either `start_at` or `due_at` is required. If both `start_at` and `due_at` is provided with a `workflow_id` that has step duration, then the `start_at` field will be ignored.
	StartAt *time.Time `json:"start_at"`

	// Title Title of the task. If not provided. it will be populated by the work request title.
	Title *string `json:"title,omitempty"`

	// WorkflowID Unique identifier of workflow
	WorkflowID *string `json:"workflow_id"`
}

// WorkRequestTaskResponse defines model for WorkRequestTaskResponse.
type WorkRequestTaskResponse struct {
	// DueAt Due date of the task, in ISO 8601 UTC format (2022-12-31T06:00:00Z)
	DueAt *time.Time `json:"due_at"`

	// ID Unique identifier of the task
	ID string `json:"id"`

	// Links Meta links
	Links WorkRequestTaskResponseLinks `json:"links"`

	// OwnerID Unique identifier of the owner
	OwnerID string `json:"owner_id"`

	// ReferenceID Reference Id of the task
	ReferenceID *string `json:"reference_id,omitempty"`

	// StartAt Start date of the task, in ISO 8601 UTC format (2022-12-01T06:00:00Z)
	StartAt *time.Time `json:"start_at"`

	// Title Title of the task
	Title string `json:"title"`
}

// WorkRequestTaskResponseLinks Meta links
type WorkRequestTaskResponseLinks struct {
	// Owner URL of the task owner
	Owner *string `json:"owner"`

	// Self URL of the task
	Self string `json:"self"`
}

// WorkRequestTextBriefRequest Written brief
type WorkRequestTextBriefRequest struct {
	Type  WorkRequestTextBriefRequestType `json:"type"`
	Value string                          `json:"value"`
}

// WorkRequestTextBriefRequestType defines model for WorkRequestTextBriefRequest.Type.
type WorkRequestTextBriefRequestType string

// WorkRequestTextBriefResponse Written brief
type WorkRequestTextBriefResponse struct {
	Type  WorkRequestTextBriefResponseType `json:"type"`
	Value string                           `json:"value"`
}

// WorkRequestTextBriefResponseType defines model for WorkRequestTextBriefResponse.Type.
type WorkRequestTextBriefResponseType string

// WorkRequestUpdateRequest defines model for WorkRequestUpdateRequest.
type WorkRequestUpdateRequest struct {
	// Assignees List of users assigned to the work request
	Assignees *[]string `json:"assignees,omitempty"`

	// Priority Update the priority of the work request
	Priority *WorkRequestUpdateRequestPriority `json:"priority,omitempty"`

	// Status Update the status of the work request
	Status *WorkRequestUpdateRequestStatus `json:"status,omitempty"`
}

// WorkRequestUpdateRequestPriority Update the priority of the work request
type WorkRequestUpdateRequestPriority string

// WorkRequestUpdateRequestStatus Update the status of the work request
type WorkRequestUpdateRequestStatus string

// WorkflowChannel defines model for WorkflowChannel.
type WorkflowChannel struct {
	// DisplayName Name of the channel
	DisplayName string `json:"display_name"`

	// Type Type of channel
	Type string `json:"type"`
}

// WorkflowCustomfield defines model for WorkflowCustomfield.
type WorkflowCustomfield struct {
	// DefaultValue List of the custom field choices
	DefaultValue []string `json:"default_value"`

	// Field Name of the custom field
	Field string `json:"field"`

	// IsRequired Indicates whether a custom field is required
	IsRequired bool `json:"is_required"`

	// Type Type of custom field
	Type string `json:"type"`
}

// WorkflowLabel defines model for WorkflowLabel.
type WorkflowLabel struct {
	// DefaultValue List of label names
	DefaultValue []string `json:"default_value"`

	// IsRequired Indicates whether a label is required
	IsRequired bool `json:"is_required"`

	// IsRequiredAtTaskCreation Indicates whether a label is required at task creation
	IsRequiredAtTaskCreation bool `json:"is_required_at_task_creation"`

	// LabelType Name of the Label Group
	LabelType string `json:"label_type"`
}

// WorkflowListResponseItem defines model for WorkflowListResponseItem.
type WorkflowListResponseItem struct {
	// Description Description of the workflow
	Description *string `json:"description"`

	// ID Unique identifier of the workflow
	ID string `json:"id"`

	// IsFlexible Indicates whether the workflow is flexible
	IsFlexible bool `json:"is_flexible"`

	// Name Name of the workflow
	Name string `json:"name"`
}

// WorkflowResponse defines model for WorkflowResponse.
type WorkflowResponse struct {
	// Description Description of the workflow
	Description *string `json:"description"`

	// Fields List of fields
	Fields []ObjectFieldCreateResponse `json:"fields"`

	// ID Unique identifier of the workflow
	ID string `json:"id"`

	// IsActive Indicates whether the workflow is active
	IsActive bool `json:"is_active"`

	// IsAssetApprovalEnabled Indicates whether the asset approval of a workflow is enabled
	IsAssetApprovalEnabled bool `json:"is_asset_approval_enabled"`

	// IsFlexible Indicates whether the workflow is flexible
	IsFlexible bool `json:"is_flexible"`

	// IsResourceManagementEnabled Indicates whether the resource management of a workflow is enabled
	IsResourceManagementEnabled bool `json:"is_resource_management_enabled"`

	// IsSmartDurationEnabled Indicates whether the smart duration of a workflow is enabled
	IsSmartDurationEnabled bool `json:"is_smart_duration_enabled"`

	// Name Name of the workflow
	Name string `json:"name"`

	// Steps List of steps
	Steps []WorkflowResponseSteps `json:"steps"`
}

// WorkflowResponseStepsSubStepsActions defines model for .
type WorkflowResponseStepsSubStepsActions struct {
	// Name Name of the action
	Name string `json:"name"`
}

// WorkflowResponseStepsSubStepsAssigneesLinks Meta links
type WorkflowResponseStepsSubStepsAssigneesLinks struct {
	// Self URL of the assignee
	Self string `json:"self"`
}

// WorkflowResponseStepsSubStepsAssigneesType Type of the assignee
type WorkflowResponseStepsSubStepsAssigneesType string

// WorkflowResponseStepsSubStepsAssignees defines model for .
type WorkflowResponseStepsSubStepsAssignees struct {
	// ID Identifier of the sub-step assignee
	ID string `json:"id"`

	// Links Meta links
	Links WorkflowResponseStepsSubStepsAssigneesLinks `json:"links"`

	// Name Name of the assignee
	Name string `json:"name"`

	// Type Type of the assignee
	Type WorkflowResponseStepsSubStepsAssigneesType `json:"type"`
}

// WorkflowResponseStepsSubStepsType Type of the sub-step
type WorkflowResponseStepsSubStepsType string

// WorkflowResponseStepsSubSteps defines model for .
type WorkflowResponseStepsSubSteps struct {
	// Actions List of actions of a sub-step
	Actions []WorkflowResponseStepsSubStepsActions `json:"actions"`

	// Assignees List of assignees of a sub-step
	Assignees []WorkflowResponseStepsSubStepsAssignees `json:"assignees"`

	// ID Identifier of the sub-step
	ID string `json:"id"`

	// Name Name of the sub-step
	Name string `json:"name"`

	// Type Type of the sub-step
	Type WorkflowResponseStepsSubStepsType `json:"type"`
}

// WorkflowResponseSteps defines model for .
type WorkflowResponseSteps struct {
	// ID Identifier of the step
	ID string `json:"id"`

	// Name Name of the step
	Name string `json:"name"`

	// SubSteps List of sub-steps
	SubSteps []WorkflowResponseStepsSubSteps `json:"sub_steps"`
}

// WorkflowStep defines model for WorkflowStep.
type WorkflowStep struct {
	// Description Description of the workflow step
	Description *string `json:"description"`

	// Duration Duration of the workflow step
	Duration *int `json:"duration"`

	// Label Name of the workflow step
	Label string `json:"label"`

	// Substeps List of the substeps
	Substeps []WorkflowSubStep `json:"substeps"`
}

// WorkflowSubStep defines model for WorkflowSubStep.
type WorkflowSubStep struct {
	// Actions List of actions
	Actions []string `json:"actions"`

	// Description Description of the substep
	Description *string `json:"description"`

	// ExternalSubStepConfig External substep configuration
	ExternalSubStepConfig *WorkflowSubStepExternalSubStepConfig `json:"external_sub_step_config,omitempty"`

	// ExternalSystem Name of the external system if the substep is external
	ExternalSystem *string `json:"external_system"`

	// IsExternal Indicates whether a substep is external
	IsExternal *bool `json:"is_external"`

	// Label Name of the substep
	Label *string `json:"label"`
}

// WorkflowSubStepExternalSubStepConfig External substep configuration
type WorkflowSubStepExternalSubStepConfig struct {
	// IsUserInteractionAllowed Indicates whether the user interaction is allowed
	IsUserInteractionAllowed *bool `json:"is_user_interaction_allowed,omitempty"`
}

// ErrorReason An enumeration.
type ErrorReason string

// Offset defines model for offset.
type Offset = int

// PageSize defines model for page_size.
type PageSize = int

// BadGateway Error payload
type BadGateway = Error

// ClientError Error payload
type ClientError = Error

// Forbidden Error payload
type Forbidden = Error

// GatewayTimeout Error payload
type GatewayTimeout = Error

// NotFound Error payload
type NotFound = Error

// Unauthorized Error payload
type Unauthorized = Error

// UnprocessableEntity Error payload
type UnprocessableEntity = Error

// Getter for additional properties for ContentGraphQueryResponse. Returns the specified
// element and whether it was found
func (a ContentGraphQueryResponse) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ContentGraphQueryResponse
func (a *ContentGraphQueryResponse) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ContentGraphQueryResponse to handle AdditionalProperties
func (a *ContentGraphQueryResponse) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["data"]; found {
		err = json.Unmarshal(raw, &a.Data)
		if err != nil {
			return fmt.Errorf("error reading 'data': %w", err)
		}
		delete(object, "data")
	}

	if raw, found := object["errors"]; found {
		err = json.Unmarshal(raw, &a.Errors)
		if err != nil {
			return fmt.Errorf("error reading 'errors': %w", err)
		}
		delete(object, "errors")
	}

	if raw, found := object["extensions"]; found {
		err = json.Unmarshal(raw, &a.Extensions)
		if err != nil {
			return fmt.Errorf("error reading 'extensions': %w", err)
		}
		delete(object, "extensions")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ContentGraphQueryResponse to handle AdditionalProperties
func (a ContentGraphQueryResponse) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Data != nil {
		object["data"], err = json.Marshal(a.Data)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'data': %w", err)
		}
	}

	if a.Errors != nil {
		object["errors"], err = json.Marshal(a.Errors)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'errors': %w", err)
		}
	}

	if a.Extensions != nil {
		object["extensions"], err = json.Marshal(a.Extensions)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'extensions': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for Error. Returns the specified
// element and whether it was found
func (a Error) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for Error
func (a *Error) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for Error to handle AdditionalProperties
func (a *Error) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["errors"]; found {
		err = json.Unmarshal(raw, &a.Errors)
		if err != nil {
			return fmt.Errorf("error reading 'errors': %w", err)
		}
		delete(object, "errors")
	}

	if raw, found := object["message"]; found {
		err = json.Unmarshal(raw, &a.Message)
		if err != nil {
			return fmt.Errorf("error reading 'message': %w", err)
		}
		delete(object, "message")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for Error to handle AdditionalProperties
func (a Error) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Errors != nil {
		object["errors"], err = json.Marshal(a.Errors)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'errors': %w", err)
		}
	}

	object["message"], err = json.Marshal(a.Message)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'message': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for RecursiveStructuredContent. Returns the specified
// element and whether it was found
func (a RecursiveStructuredContent) Get(fieldName string) (value []LocalizedFieldValuesWithEmbeddedContent, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for RecursiveStructuredContent
func (a *RecursiveStructuredContent) Set(fieldName string, value []LocalizedFieldValuesWithEmbeddedContent) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string][]LocalizedFieldValuesWithEmbeddedContent)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for RecursiveStructuredContent to handle AdditionalProperties
func (a *RecursiveStructuredContent) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["content_type_guid"]; found {
		err = json.Unmarshal(raw, &a.ContentTypeGUID)
		if err != nil {
			return fmt.Errorf("error reading 'content_type_guid': %w", err)
		}
		delete(object, "content_type_guid")
	}

	if raw, found := object["expired"]; found {
		err = json.Unmarshal(raw, &a.Expired)
		if err != nil {
			return fmt.Errorf("error reading 'expired': %w", err)
		}
		delete(object, "expired")
	}

	if raw, found := object["expiry_datetime"]; found {
		err = json.Unmarshal(raw, &a.ExpiryDatetime)
		if err != nil {
			return fmt.Errorf("error reading 'expiry_datetime': %w", err)
		}
		delete(object, "expiry_datetime")
	}

	if raw, found := object["root_content"]; found {
		err = json.Unmarshal(raw, &a.RootContent)
		if err != nil {
			return fmt.Errorf("error reading 'root_content': %w", err)
		}
		delete(object, "root_content")
	}

	if raw, found := object["source"]; found {
		err = json.Unmarshal(raw, &a.Source)
		if err != nil {
			return fmt.Errorf("error reading 'source': %w", err)
		}
		delete(object, "source")
	}

	if raw, found := object["source_id"]; found {
		err = json.Unmarshal(raw, &a.SourceID)
		if err != nil {
			return fmt.Errorf("error reading 'source_id': %w", err)
		}
		delete(object, "source_id")
	}

	if raw, found := object["source_metadata"]; found {
		err = json.Unmarshal(raw, &a.SourceMetadata)
		if err != nil {
			return fmt.Errorf("error reading 'source_metadata': %w", err)
		}
		delete(object, "source_metadata")
	}

	if raw, found := object["title"]; found {
		err = json.Unmarshal(raw, &a.Title)
		if err != nil {
			return fmt.Errorf("error reading 'title': %w", err)
		}
		delete(object, "title")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string][]LocalizedFieldValuesWithEmbeddedContent)
		for fieldName, fieldBuf := range object {
			var fieldVal []LocalizedFieldValuesWithEmbeddedContent
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for RecursiveStructuredContent to handle AdditionalProperties
func (a RecursiveStructuredContent) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["content_type_guid"], err = json.Marshal(a.ContentTypeGUID)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'content_type_guid': %w", err)
	}

	if a.Expired != nil {
		object["expired"], err = json.Marshal(a.Expired)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'expired': %w", err)
		}
	}

	if a.ExpiryDatetime != nil {
		object["expiry_datetime"], err = json.Marshal(a.ExpiryDatetime)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'expiry_datetime': %w", err)
		}
	}

	object["root_content"], err = json.Marshal(a.RootContent)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'root_content': %w", err)
	}

	if a.Source != nil {
		object["source"], err = json.Marshal(a.Source)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'source': %w", err)
		}
	}

	if a.SourceID != nil {
		object["source_id"], err = json.Marshal(a.SourceID)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'source_id': %w", err)
		}
	}

	if a.SourceMetadata != nil {
		object["source_metadata"], err = json.Marshal(a.SourceMetadata)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'source_metadata': %w", err)
		}
	}

	if a.Title != nil {
		object["title"], err = json.Marshal(a.Title)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'title': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for StructuredContentBody. Returns the specified
// element and whether it was found
func (a StructuredContentBody) Get(fieldName string) (value []LocalizedFieldValuesWithEmbeddedContent, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for StructuredContentBody
func (a *StructuredContentBody) Set(fieldName string, value []LocalizedFieldValuesWithEmbeddedContent) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string][]LocalizedFieldValuesWithEmbeddedContent)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for StructuredContentBody to handle AdditionalProperties
func (a *StructuredContentBody) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["has_embedded"]; found {
		err = json.Unmarshal(raw, &a.HasEmbedded)
		if err != nil {
			return fmt.Errorf("error reading 'has_embedded': %w", err)
		}
		delete(object, "has_embedded")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string][]LocalizedFieldValuesWithEmbeddedContent)
		for fieldName, fieldBuf := range object {
			var fieldVal []LocalizedFieldValuesWithEmbeddedContent
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for StructuredContentBody to handle AdditionalProperties
func (a StructuredContentBody) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.HasEmbedded != nil {
		object["has_embedded"], err = json.Marshal(a.HasEmbedded)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'has_embedded': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// AsAssetFieldTypeCommon returns the union data inside the AssetFieldListResponseItem as a AssetFieldTypeCommon
func (t AssetFieldListResponseItem) AsAssetFieldTypeCommon() (AssetFieldTypeCommon, error) {
	var body AssetFieldTypeCommon
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetFieldTypeCommon overwrites any union data inside the AssetFieldListResponseItem as the provided AssetFieldTypeCommon
func (t *AssetFieldListResponseItem) FromAssetFieldTypeCommon(v AssetFieldTypeCommon) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetFieldTypeCommon performs a merge with any union data inside the AssetFieldListResponseItem, using the provided AssetFieldTypeCommon
func (t *AssetFieldListResponseItem) MergeAssetFieldTypeCommon(v AssetFieldTypeCommon) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAssetFieldTypeLabel returns the union data inside the AssetFieldListResponseItem as a AssetFieldTypeLabel
func (t AssetFieldListResponseItem) AsAssetFieldTypeLabel() (AssetFieldTypeLabel, error) {
	var body AssetFieldTypeLabel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetFieldTypeLabel overwrites any union data inside the AssetFieldListResponseItem as the provided AssetFieldTypeLabel
func (t *AssetFieldListResponseItem) FromAssetFieldTypeLabel(v AssetFieldTypeLabel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetFieldTypeLabel performs a merge with any union data inside the AssetFieldListResponseItem, using the provided AssetFieldTypeLabel
func (t *AssetFieldListResponseItem) MergeAssetFieldTypeLabel(v AssetFieldTypeLabel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAssetFieldTypeDropdown returns the union data inside the AssetFieldListResponseItem as a AssetFieldTypeDropdown
func (t AssetFieldListResponseItem) AsAssetFieldTypeDropdown() (AssetFieldTypeDropdown, error) {
	var body AssetFieldTypeDropdown
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetFieldTypeDropdown overwrites any union data inside the AssetFieldListResponseItem as the provided AssetFieldTypeDropdown
func (t *AssetFieldListResponseItem) FromAssetFieldTypeDropdown(v AssetFieldTypeDropdown) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetFieldTypeDropdown performs a merge with any union data inside the AssetFieldListResponseItem, using the provided AssetFieldTypeDropdown
func (t *AssetFieldListResponseItem) MergeAssetFieldTypeDropdown(v AssetFieldTypeDropdown) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAssetFieldTypeRadio returns the union data inside the AssetFieldListResponseItem as a AssetFieldTypeRadio
func (t AssetFieldListResponseItem) AsAssetFieldTypeRadio() (AssetFieldTypeRadio, error) {
	var body AssetFieldTypeRadio
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetFieldTypeRadio overwrites any union data inside the AssetFieldListResponseItem as the provided AssetFieldTypeRadio
func (t *AssetFieldListResponseItem) FromAssetFieldTypeRadio(v AssetFieldTypeRadio) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetFieldTypeRadio performs a merge with any union data inside the AssetFieldListResponseItem, using the provided AssetFieldTypeRadio
func (t *AssetFieldListResponseItem) MergeAssetFieldTypeRadio(v AssetFieldTypeRadio) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAssetFieldTypeCheckbox returns the union data inside the AssetFieldListResponseItem as a AssetFieldTypeCheckbox
func (t AssetFieldListResponseItem) AsAssetFieldTypeCheckbox() (AssetFieldTypeCheckbox, error) {
	var body AssetFieldTypeCheckbox
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetFieldTypeCheckbox overwrites any union data inside the AssetFieldListResponseItem as the provided AssetFieldTypeCheckbox
func (t *AssetFieldListResponseItem) FromAssetFieldTypeCheckbox(v AssetFieldTypeCheckbox) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetFieldTypeCheckbox performs a merge with any union data inside the AssetFieldListResponseItem, using the provided AssetFieldTypeCheckbox
func (t *AssetFieldListResponseItem) MergeAssetFieldTypeCheckbox(v AssetFieldTypeCheckbox) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAssetFieldTypeSimpleNumber returns the union data inside the AssetFieldListResponseItem as a AssetFieldTypeSimpleNumber
func (t AssetFieldListResponseItem) AsAssetFieldTypeSimpleNumber() (AssetFieldTypeSimpleNumber, error) {
	var body AssetFieldTypeSimpleNumber
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetFieldTypeSimpleNumber overwrites any union data inside the AssetFieldListResponseItem as the provided AssetFieldTypeSimpleNumber
func (t *AssetFieldListResponseItem) FromAssetFieldTypeSimpleNumber(v AssetFieldTypeSimpleNumber) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetFieldTypeSimpleNumber performs a merge with any union data inside the AssetFieldListResponseItem, using the provided AssetFieldTypeSimpleNumber
func (t *AssetFieldListResponseItem) MergeAssetFieldTypeSimpleNumber(v AssetFieldTypeSimpleNumber) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAssetFieldTypePercentageNumber returns the union data inside the AssetFieldListResponseItem as a AssetFieldTypePercentageNumber
func (t AssetFieldListResponseItem) AsAssetFieldTypePercentageNumber() (AssetFieldTypePercentageNumber, error) {
	var body AssetFieldTypePercentageNumber
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetFieldTypePercentageNumber overwrites any union data inside the AssetFieldListResponseItem as the provided AssetFieldTypePercentageNumber
func (t *AssetFieldListResponseItem) FromAssetFieldTypePercentageNumber(v AssetFieldTypePercentageNumber) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetFieldTypePercentageNumber performs a merge with any union data inside the AssetFieldListResponseItem, using the provided AssetFieldTypePercentageNumber
func (t *AssetFieldListResponseItem) MergeAssetFieldTypePercentageNumber(v AssetFieldTypePercentageNumber) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAssetFieldTypeCurrencyNumber returns the union data inside the AssetFieldListResponseItem as a AssetFieldTypeCurrencyNumber
func (t AssetFieldListResponseItem) AsAssetFieldTypeCurrencyNumber() (AssetFieldTypeCurrencyNumber, error) {
	var body AssetFieldTypeCurrencyNumber
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetFieldTypeCurrencyNumber overwrites any union data inside the AssetFieldListResponseItem as the provided AssetFieldTypeCurrencyNumber
func (t *AssetFieldListResponseItem) FromAssetFieldTypeCurrencyNumber(v AssetFieldTypeCurrencyNumber) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetFieldTypeCurrencyNumber performs a merge with any union data inside the AssetFieldListResponseItem, using the provided AssetFieldTypeCurrencyNumber
func (t *AssetFieldListResponseItem) MergeAssetFieldTypeCurrencyNumber(v AssetFieldTypeCurrencyNumber) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t AssetFieldListResponseItem) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AssetFieldListResponseItem) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsStringTypeObjectFieldUpdatePayload returns the union data inside the AssetFieldUpdateRequest as a StringTypeObjectFieldUpdatePayload
func (t AssetFieldUpdateRequest) AsStringTypeObjectFieldUpdatePayload() (StringTypeObjectFieldUpdatePayload, error) {
	var body StringTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromStringTypeObjectFieldUpdatePayload overwrites any union data inside the AssetFieldUpdateRequest as the provided StringTypeObjectFieldUpdatePayload
func (t *AssetFieldUpdateRequest) FromStringTypeObjectFieldUpdatePayload(v StringTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeStringTypeObjectFieldUpdatePayload performs a merge with any union data inside the AssetFieldUpdateRequest, using the provided StringTypeObjectFieldUpdatePayload
func (t *AssetFieldUpdateRequest) MergeStringTypeObjectFieldUpdatePayload(v StringTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMultiChoiceTypeObjectFieldUpdatePayload returns the union data inside the AssetFieldUpdateRequest as a MultiChoiceTypeObjectFieldUpdatePayload
func (t AssetFieldUpdateRequest) AsMultiChoiceTypeObjectFieldUpdatePayload() (MultiChoiceTypeObjectFieldUpdatePayload, error) {
	var body MultiChoiceTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMultiChoiceTypeObjectFieldUpdatePayload overwrites any union data inside the AssetFieldUpdateRequest as the provided MultiChoiceTypeObjectFieldUpdatePayload
func (t *AssetFieldUpdateRequest) FromMultiChoiceTypeObjectFieldUpdatePayload(v MultiChoiceTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMultiChoiceTypeObjectFieldUpdatePayload performs a merge with any union data inside the AssetFieldUpdateRequest, using the provided MultiChoiceTypeObjectFieldUpdatePayload
func (t *AssetFieldUpdateRequest) MergeMultiChoiceTypeObjectFieldUpdatePayload(v MultiChoiceTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDropdownTypeObjectFieldUpdatePayload returns the union data inside the AssetFieldUpdateRequest as a DropdownTypeObjectFieldUpdatePayload
func (t AssetFieldUpdateRequest) AsDropdownTypeObjectFieldUpdatePayload() (DropdownTypeObjectFieldUpdatePayload, error) {
	var body DropdownTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDropdownTypeObjectFieldUpdatePayload overwrites any union data inside the AssetFieldUpdateRequest as the provided DropdownTypeObjectFieldUpdatePayload
func (t *AssetFieldUpdateRequest) FromDropdownTypeObjectFieldUpdatePayload(v DropdownTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDropdownTypeObjectFieldUpdatePayload performs a merge with any union data inside the AssetFieldUpdateRequest, using the provided DropdownTypeObjectFieldUpdatePayload
func (t *AssetFieldUpdateRequest) MergeDropdownTypeObjectFieldUpdatePayload(v DropdownTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRadioButtonTypeObjectFieldUpdatePayload returns the union data inside the AssetFieldUpdateRequest as a RadioButtonTypeObjectFieldUpdatePayload
func (t AssetFieldUpdateRequest) AsRadioButtonTypeObjectFieldUpdatePayload() (RadioButtonTypeObjectFieldUpdatePayload, error) {
	var body RadioButtonTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRadioButtonTypeObjectFieldUpdatePayload overwrites any union data inside the AssetFieldUpdateRequest as the provided RadioButtonTypeObjectFieldUpdatePayload
func (t *AssetFieldUpdateRequest) FromRadioButtonTypeObjectFieldUpdatePayload(v RadioButtonTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRadioButtonTypeObjectFieldUpdatePayload performs a merge with any union data inside the AssetFieldUpdateRequest, using the provided RadioButtonTypeObjectFieldUpdatePayload
func (t *AssetFieldUpdateRequest) MergeRadioButtonTypeObjectFieldUpdatePayload(v RadioButtonTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsNumberTypeObjectFieldUpdatePayload returns the union data inside the AssetFieldUpdateRequest as a NumberTypeObjectFieldUpdatePayload
func (t AssetFieldUpdateRequest) AsNumberTypeObjectFieldUpdatePayload() (NumberTypeObjectFieldUpdatePayload, error) {
	var body NumberTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberTypeObjectFieldUpdatePayload overwrites any union data inside the AssetFieldUpdateRequest as the provided NumberTypeObjectFieldUpdatePayload
func (t *AssetFieldUpdateRequest) FromNumberTypeObjectFieldUpdatePayload(v NumberTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberTypeObjectFieldUpdatePayload performs a merge with any union data inside the AssetFieldUpdateRequest, using the provided NumberTypeObjectFieldUpdatePayload
func (t *AssetFieldUpdateRequest) MergeNumberTypeObjectFieldUpdatePayload(v NumberTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDateTypeObjectFieldUpdatePayload returns the union data inside the AssetFieldUpdateRequest as a DateTypeObjectFieldUpdatePayload
func (t AssetFieldUpdateRequest) AsDateTypeObjectFieldUpdatePayload() (DateTypeObjectFieldUpdatePayload, error) {
	var body DateTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDateTypeObjectFieldUpdatePayload overwrites any union data inside the AssetFieldUpdateRequest as the provided DateTypeObjectFieldUpdatePayload
func (t *AssetFieldUpdateRequest) FromDateTypeObjectFieldUpdatePayload(v DateTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDateTypeObjectFieldUpdatePayload performs a merge with any union data inside the AssetFieldUpdateRequest, using the provided DateTypeObjectFieldUpdatePayload
func (t *AssetFieldUpdateRequest) MergeDateTypeObjectFieldUpdatePayload(v DateTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAssetTypeObjectFieldUpdatePayload returns the union data inside the AssetFieldUpdateRequest as a AssetTypeObjectFieldUpdatePayload
func (t AssetFieldUpdateRequest) AsAssetTypeObjectFieldUpdatePayload() (AssetTypeObjectFieldUpdatePayload, error) {
	var body AssetTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetTypeObjectFieldUpdatePayload overwrites any union data inside the AssetFieldUpdateRequest as the provided AssetTypeObjectFieldUpdatePayload
func (t *AssetFieldUpdateRequest) FromAssetTypeObjectFieldUpdatePayload(v AssetTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetTypeObjectFieldUpdatePayload performs a merge with any union data inside the AssetFieldUpdateRequest, using the provided AssetTypeObjectFieldUpdatePayload
func (t *AssetFieldUpdateRequest) MergeAssetTypeObjectFieldUpdatePayload(v AssetTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t AssetFieldUpdateRequest) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AssetFieldUpdateRequest) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsStringTypeObjectField returns the union data inside the AssetFieldsUpdateRequest_Item as a StringTypeObjectField
func (t AssetFieldsUpdateRequest_Item) AsStringTypeObjectField() (StringTypeObjectField, error) {
	var body StringTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromStringTypeObjectField overwrites any union data inside the AssetFieldsUpdateRequest_Item as the provided StringTypeObjectField
func (t *AssetFieldsUpdateRequest_Item) FromStringTypeObjectField(v StringTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeStringTypeObjectField performs a merge with any union data inside the AssetFieldsUpdateRequest_Item, using the provided StringTypeObjectField
func (t *AssetFieldsUpdateRequest_Item) MergeStringTypeObjectField(v StringTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMultiChoiceTypeObjectField returns the union data inside the AssetFieldsUpdateRequest_Item as a MultiChoiceTypeObjectField
func (t AssetFieldsUpdateRequest_Item) AsMultiChoiceTypeObjectField() (MultiChoiceTypeObjectField, error) {
	var body MultiChoiceTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMultiChoiceTypeObjectField overwrites any union data inside the AssetFieldsUpdateRequest_Item as the provided MultiChoiceTypeObjectField
func (t *AssetFieldsUpdateRequest_Item) FromMultiChoiceTypeObjectField(v MultiChoiceTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMultiChoiceTypeObjectField performs a merge with any union data inside the AssetFieldsUpdateRequest_Item, using the provided MultiChoiceTypeObjectField
func (t *AssetFieldsUpdateRequest_Item) MergeMultiChoiceTypeObjectField(v MultiChoiceTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDropdownTypeObjectField returns the union data inside the AssetFieldsUpdateRequest_Item as a DropdownTypeObjectField
func (t AssetFieldsUpdateRequest_Item) AsDropdownTypeObjectField() (DropdownTypeObjectField, error) {
	var body DropdownTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDropdownTypeObjectField overwrites any union data inside the AssetFieldsUpdateRequest_Item as the provided DropdownTypeObjectField
func (t *AssetFieldsUpdateRequest_Item) FromDropdownTypeObjectField(v DropdownTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDropdownTypeObjectField performs a merge with any union data inside the AssetFieldsUpdateRequest_Item, using the provided DropdownTypeObjectField
func (t *AssetFieldsUpdateRequest_Item) MergeDropdownTypeObjectField(v DropdownTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRadioButtonTypeObjectField returns the union data inside the AssetFieldsUpdateRequest_Item as a RadioButtonTypeObjectField
func (t AssetFieldsUpdateRequest_Item) AsRadioButtonTypeObjectField() (RadioButtonTypeObjectField, error) {
	var body RadioButtonTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRadioButtonTypeObjectField overwrites any union data inside the AssetFieldsUpdateRequest_Item as the provided RadioButtonTypeObjectField
func (t *AssetFieldsUpdateRequest_Item) FromRadioButtonTypeObjectField(v RadioButtonTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRadioButtonTypeObjectField performs a merge with any union data inside the AssetFieldsUpdateRequest_Item, using the provided RadioButtonTypeObjectField
func (t *AssetFieldsUpdateRequest_Item) MergeRadioButtonTypeObjectField(v RadioButtonTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsNumberTypeObjectField returns the union data inside the AssetFieldsUpdateRequest_Item as a NumberTypeObjectField
func (t AssetFieldsUpdateRequest_Item) AsNumberTypeObjectField() (NumberTypeObjectField, error) {
	var body NumberTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberTypeObjectField overwrites any union data inside the AssetFieldsUpdateRequest_Item as the provided NumberTypeObjectField
func (t *AssetFieldsUpdateRequest_Item) FromNumberTypeObjectField(v NumberTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberTypeObjectField performs a merge with any union data inside the AssetFieldsUpdateRequest_Item, using the provided NumberTypeObjectField
func (t *AssetFieldsUpdateRequest_Item) MergeNumberTypeObjectField(v NumberTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDateTypeObjectField returns the union data inside the AssetFieldsUpdateRequest_Item as a DateTypeObjectField
func (t AssetFieldsUpdateRequest_Item) AsDateTypeObjectField() (DateTypeObjectField, error) {
	var body DateTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDateTypeObjectField overwrites any union data inside the AssetFieldsUpdateRequest_Item as the provided DateTypeObjectField
func (t *AssetFieldsUpdateRequest_Item) FromDateTypeObjectField(v DateTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDateTypeObjectField performs a merge with any union data inside the AssetFieldsUpdateRequest_Item, using the provided DateTypeObjectField
func (t *AssetFieldsUpdateRequest_Item) MergeDateTypeObjectField(v DateTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAssetTypeObjectField returns the union data inside the AssetFieldsUpdateRequest_Item as a AssetTypeObjectField
func (t AssetFieldsUpdateRequest_Item) AsAssetTypeObjectField() (AssetTypeObjectField, error) {
	var body AssetTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetTypeObjectField overwrites any union data inside the AssetFieldsUpdateRequest_Item as the provided AssetTypeObjectField
func (t *AssetFieldsUpdateRequest_Item) FromAssetTypeObjectField(v AssetTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetTypeObjectField performs a merge with any union data inside the AssetFieldsUpdateRequest_Item, using the provided AssetTypeObjectField
func (t *AssetFieldsUpdateRequest_Item) MergeAssetTypeObjectField(v AssetTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t AssetFieldsUpdateRequest_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AssetFieldsUpdateRequest_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsLocationDefaultValue returns the union data inside the BaseFieldDefinition_DefaultValues_Item as a LocationDefaultValue
func (t BaseFieldDefinition_DefaultValues_Item) AsLocationDefaultValue() (LocationDefaultValue, error) {
	var body LocationDefaultValue
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLocationDefaultValue overwrites any union data inside the BaseFieldDefinition_DefaultValues_Item as the provided LocationDefaultValue
func (t *BaseFieldDefinition_DefaultValues_Item) FromLocationDefaultValue(v LocationDefaultValue) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLocationDefaultValue performs a merge with any union data inside the BaseFieldDefinition_DefaultValues_Item, using the provided LocationDefaultValue
func (t *BaseFieldDefinition_DefaultValues_Item) MergeLocationDefaultValue(v LocationDefaultValue) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBaseFieldDefinitionDefaultValues1 returns the union data inside the BaseFieldDefinition_DefaultValues_Item as a BaseFieldDefinitionDefaultValues1
func (t BaseFieldDefinition_DefaultValues_Item) AsBaseFieldDefinitionDefaultValues1() (BaseFieldDefinitionDefaultValues1, error) {
	var body BaseFieldDefinitionDefaultValues1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBaseFieldDefinitionDefaultValues1 overwrites any union data inside the BaseFieldDefinition_DefaultValues_Item as the provided BaseFieldDefinitionDefaultValues1
func (t *BaseFieldDefinition_DefaultValues_Item) FromBaseFieldDefinitionDefaultValues1(v BaseFieldDefinitionDefaultValues1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBaseFieldDefinitionDefaultValues1 performs a merge with any union data inside the BaseFieldDefinition_DefaultValues_Item, using the provided BaseFieldDefinitionDefaultValues1
func (t *BaseFieldDefinition_DefaultValues_Item) MergeBaseFieldDefinitionDefaultValues1(v BaseFieldDefinitionDefaultValues1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBaseFieldDefinitionDefaultValues2 returns the union data inside the BaseFieldDefinition_DefaultValues_Item as a BaseFieldDefinitionDefaultValues2
func (t BaseFieldDefinition_DefaultValues_Item) AsBaseFieldDefinitionDefaultValues2() (BaseFieldDefinitionDefaultValues2, error) {
	var body BaseFieldDefinitionDefaultValues2
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBaseFieldDefinitionDefaultValues2 overwrites any union data inside the BaseFieldDefinition_DefaultValues_Item as the provided BaseFieldDefinitionDefaultValues2
func (t *BaseFieldDefinition_DefaultValues_Item) FromBaseFieldDefinitionDefaultValues2(v BaseFieldDefinitionDefaultValues2) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBaseFieldDefinitionDefaultValues2 performs a merge with any union data inside the BaseFieldDefinition_DefaultValues_Item, using the provided BaseFieldDefinitionDefaultValues2
func (t *BaseFieldDefinition_DefaultValues_Item) MergeBaseFieldDefinitionDefaultValues2(v BaseFieldDefinitionDefaultValues2) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBaseFieldDefinitionDefaultValues3 returns the union data inside the BaseFieldDefinition_DefaultValues_Item as a BaseFieldDefinitionDefaultValues3
func (t BaseFieldDefinition_DefaultValues_Item) AsBaseFieldDefinitionDefaultValues3() (BaseFieldDefinitionDefaultValues3, error) {
	var body BaseFieldDefinitionDefaultValues3
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBaseFieldDefinitionDefaultValues3 overwrites any union data inside the BaseFieldDefinition_DefaultValues_Item as the provided BaseFieldDefinitionDefaultValues3
func (t *BaseFieldDefinition_DefaultValues_Item) FromBaseFieldDefinitionDefaultValues3(v BaseFieldDefinitionDefaultValues3) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBaseFieldDefinitionDefaultValues3 performs a merge with any union data inside the BaseFieldDefinition_DefaultValues_Item, using the provided BaseFieldDefinitionDefaultValues3
func (t *BaseFieldDefinition_DefaultValues_Item) MergeBaseFieldDefinitionDefaultValues3(v BaseFieldDefinitionDefaultValues3) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBaseFieldDefinitionDefaultValues4 returns the union data inside the BaseFieldDefinition_DefaultValues_Item as a BaseFieldDefinitionDefaultValues4
func (t BaseFieldDefinition_DefaultValues_Item) AsBaseFieldDefinitionDefaultValues4() (BaseFieldDefinitionDefaultValues4, error) {
	var body BaseFieldDefinitionDefaultValues4
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBaseFieldDefinitionDefaultValues4 overwrites any union data inside the BaseFieldDefinition_DefaultValues_Item as the provided BaseFieldDefinitionDefaultValues4
func (t *BaseFieldDefinition_DefaultValues_Item) FromBaseFieldDefinitionDefaultValues4(v BaseFieldDefinitionDefaultValues4) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBaseFieldDefinitionDefaultValues4 performs a merge with any union data inside the BaseFieldDefinition_DefaultValues_Item, using the provided BaseFieldDefinitionDefaultValues4
func (t *BaseFieldDefinition_DefaultValues_Item) MergeBaseFieldDefinitionDefaultValues4(v BaseFieldDefinitionDefaultValues4) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t BaseFieldDefinition_DefaultValues_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *BaseFieldDefinition_DefaultValues_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkRequestTextBriefRequest returns the union data inside the BriefTypeFormFieldRequest_Values_Item as a WorkRequestTextBriefRequest
func (t BriefTypeFormFieldRequest_Values_Item) AsWorkRequestTextBriefRequest() (WorkRequestTextBriefRequest, error) {
	var body WorkRequestTextBriefRequest
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestTextBriefRequest overwrites any union data inside the BriefTypeFormFieldRequest_Values_Item as the provided WorkRequestTextBriefRequest
func (t *BriefTypeFormFieldRequest_Values_Item) FromWorkRequestTextBriefRequest(v WorkRequestTextBriefRequest) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestTextBriefRequest performs a merge with any union data inside the BriefTypeFormFieldRequest_Values_Item, using the provided WorkRequestTextBriefRequest
func (t *BriefTypeFormFieldRequest_Values_Item) MergeWorkRequestTextBriefRequest(v WorkRequestTextBriefRequest) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestAttachmentBriefRequest returns the union data inside the BriefTypeFormFieldRequest_Values_Item as a WorkRequestAttachmentBriefRequest
func (t BriefTypeFormFieldRequest_Values_Item) AsWorkRequestAttachmentBriefRequest() (WorkRequestAttachmentBriefRequest, error) {
	var body WorkRequestAttachmentBriefRequest
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestAttachmentBriefRequest overwrites any union data inside the BriefTypeFormFieldRequest_Values_Item as the provided WorkRequestAttachmentBriefRequest
func (t *BriefTypeFormFieldRequest_Values_Item) FromWorkRequestAttachmentBriefRequest(v WorkRequestAttachmentBriefRequest) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestAttachmentBriefRequest performs a merge with any union data inside the BriefTypeFormFieldRequest_Values_Item, using the provided WorkRequestAttachmentBriefRequest
func (t *BriefTypeFormFieldRequest_Values_Item) MergeWorkRequestAttachmentBriefRequest(v WorkRequestAttachmentBriefRequest) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t BriefTypeFormFieldRequest_Values_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *BriefTypeFormFieldRequest_Values_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkRequestTextBriefResponse returns the union data inside the BriefTypeFormFieldResponse_Values_Item as a WorkRequestTextBriefResponse
func (t BriefTypeFormFieldResponse_Values_Item) AsWorkRequestTextBriefResponse() (WorkRequestTextBriefResponse, error) {
	var body WorkRequestTextBriefResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestTextBriefResponse overwrites any union data inside the BriefTypeFormFieldResponse_Values_Item as the provided WorkRequestTextBriefResponse
func (t *BriefTypeFormFieldResponse_Values_Item) FromWorkRequestTextBriefResponse(v WorkRequestTextBriefResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestTextBriefResponse performs a merge with any union data inside the BriefTypeFormFieldResponse_Values_Item, using the provided WorkRequestTextBriefResponse
func (t *BriefTypeFormFieldResponse_Values_Item) MergeWorkRequestTextBriefResponse(v WorkRequestTextBriefResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestAttachmentBriefResponse returns the union data inside the BriefTypeFormFieldResponse_Values_Item as a WorkRequestAttachmentBriefResponse
func (t BriefTypeFormFieldResponse_Values_Item) AsWorkRequestAttachmentBriefResponse() (WorkRequestAttachmentBriefResponse, error) {
	var body WorkRequestAttachmentBriefResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestAttachmentBriefResponse overwrites any union data inside the BriefTypeFormFieldResponse_Values_Item as the provided WorkRequestAttachmentBriefResponse
func (t *BriefTypeFormFieldResponse_Values_Item) FromWorkRequestAttachmentBriefResponse(v WorkRequestAttachmentBriefResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestAttachmentBriefResponse performs a merge with any union data inside the BriefTypeFormFieldResponse_Values_Item, using the provided WorkRequestAttachmentBriefResponse
func (t *BriefTypeFormFieldResponse_Values_Item) MergeWorkRequestAttachmentBriefResponse(v WorkRequestAttachmentBriefResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t BriefTypeFormFieldResponse_Values_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *BriefTypeFormFieldResponse_Values_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsStringTypeObjectFieldUpdatePayload returns the union data inside the CampaignFieldUpdateRequest as a StringTypeObjectFieldUpdatePayload
func (t CampaignFieldUpdateRequest) AsStringTypeObjectFieldUpdatePayload() (StringTypeObjectFieldUpdatePayload, error) {
	var body StringTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromStringTypeObjectFieldUpdatePayload overwrites any union data inside the CampaignFieldUpdateRequest as the provided StringTypeObjectFieldUpdatePayload
func (t *CampaignFieldUpdateRequest) FromStringTypeObjectFieldUpdatePayload(v StringTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeStringTypeObjectFieldUpdatePayload performs a merge with any union data inside the CampaignFieldUpdateRequest, using the provided StringTypeObjectFieldUpdatePayload
func (t *CampaignFieldUpdateRequest) MergeStringTypeObjectFieldUpdatePayload(v StringTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMultiChoiceTypeObjectFieldUpdatePayload returns the union data inside the CampaignFieldUpdateRequest as a MultiChoiceTypeObjectFieldUpdatePayload
func (t CampaignFieldUpdateRequest) AsMultiChoiceTypeObjectFieldUpdatePayload() (MultiChoiceTypeObjectFieldUpdatePayload, error) {
	var body MultiChoiceTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMultiChoiceTypeObjectFieldUpdatePayload overwrites any union data inside the CampaignFieldUpdateRequest as the provided MultiChoiceTypeObjectFieldUpdatePayload
func (t *CampaignFieldUpdateRequest) FromMultiChoiceTypeObjectFieldUpdatePayload(v MultiChoiceTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMultiChoiceTypeObjectFieldUpdatePayload performs a merge with any union data inside the CampaignFieldUpdateRequest, using the provided MultiChoiceTypeObjectFieldUpdatePayload
func (t *CampaignFieldUpdateRequest) MergeMultiChoiceTypeObjectFieldUpdatePayload(v MultiChoiceTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDropdownTypeObjectFieldUpdatePayload returns the union data inside the CampaignFieldUpdateRequest as a DropdownTypeObjectFieldUpdatePayload
func (t CampaignFieldUpdateRequest) AsDropdownTypeObjectFieldUpdatePayload() (DropdownTypeObjectFieldUpdatePayload, error) {
	var body DropdownTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDropdownTypeObjectFieldUpdatePayload overwrites any union data inside the CampaignFieldUpdateRequest as the provided DropdownTypeObjectFieldUpdatePayload
func (t *CampaignFieldUpdateRequest) FromDropdownTypeObjectFieldUpdatePayload(v DropdownTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDropdownTypeObjectFieldUpdatePayload performs a merge with any union data inside the CampaignFieldUpdateRequest, using the provided DropdownTypeObjectFieldUpdatePayload
func (t *CampaignFieldUpdateRequest) MergeDropdownTypeObjectFieldUpdatePayload(v DropdownTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRadioButtonTypeObjectFieldUpdatePayload returns the union data inside the CampaignFieldUpdateRequest as a RadioButtonTypeObjectFieldUpdatePayload
func (t CampaignFieldUpdateRequest) AsRadioButtonTypeObjectFieldUpdatePayload() (RadioButtonTypeObjectFieldUpdatePayload, error) {
	var body RadioButtonTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRadioButtonTypeObjectFieldUpdatePayload overwrites any union data inside the CampaignFieldUpdateRequest as the provided RadioButtonTypeObjectFieldUpdatePayload
func (t *CampaignFieldUpdateRequest) FromRadioButtonTypeObjectFieldUpdatePayload(v RadioButtonTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRadioButtonTypeObjectFieldUpdatePayload performs a merge with any union data inside the CampaignFieldUpdateRequest, using the provided RadioButtonTypeObjectFieldUpdatePayload
func (t *CampaignFieldUpdateRequest) MergeRadioButtonTypeObjectFieldUpdatePayload(v RadioButtonTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsNumberTypeObjectFieldUpdatePayload returns the union data inside the CampaignFieldUpdateRequest as a NumberTypeObjectFieldUpdatePayload
func (t CampaignFieldUpdateRequest) AsNumberTypeObjectFieldUpdatePayload() (NumberTypeObjectFieldUpdatePayload, error) {
	var body NumberTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberTypeObjectFieldUpdatePayload overwrites any union data inside the CampaignFieldUpdateRequest as the provided NumberTypeObjectFieldUpdatePayload
func (t *CampaignFieldUpdateRequest) FromNumberTypeObjectFieldUpdatePayload(v NumberTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberTypeObjectFieldUpdatePayload performs a merge with any union data inside the CampaignFieldUpdateRequest, using the provided NumberTypeObjectFieldUpdatePayload
func (t *CampaignFieldUpdateRequest) MergeNumberTypeObjectFieldUpdatePayload(v NumberTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDateTypeObjectFieldUpdatePayload returns the union data inside the CampaignFieldUpdateRequest as a DateTypeObjectFieldUpdatePayload
func (t CampaignFieldUpdateRequest) AsDateTypeObjectFieldUpdatePayload() (DateTypeObjectFieldUpdatePayload, error) {
	var body DateTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDateTypeObjectFieldUpdatePayload overwrites any union data inside the CampaignFieldUpdateRequest as the provided DateTypeObjectFieldUpdatePayload
func (t *CampaignFieldUpdateRequest) FromDateTypeObjectFieldUpdatePayload(v DateTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDateTypeObjectFieldUpdatePayload performs a merge with any union data inside the CampaignFieldUpdateRequest, using the provided DateTypeObjectFieldUpdatePayload
func (t *CampaignFieldUpdateRequest) MergeDateTypeObjectFieldUpdatePayload(v DateTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAssetTypeObjectFieldUpdatePayload returns the union data inside the CampaignFieldUpdateRequest as a AssetTypeObjectFieldUpdatePayload
func (t CampaignFieldUpdateRequest) AsAssetTypeObjectFieldUpdatePayload() (AssetTypeObjectFieldUpdatePayload, error) {
	var body AssetTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetTypeObjectFieldUpdatePayload overwrites any union data inside the CampaignFieldUpdateRequest as the provided AssetTypeObjectFieldUpdatePayload
func (t *CampaignFieldUpdateRequest) FromAssetTypeObjectFieldUpdatePayload(v AssetTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetTypeObjectFieldUpdatePayload performs a merge with any union data inside the CampaignFieldUpdateRequest, using the provided AssetTypeObjectFieldUpdatePayload
func (t *CampaignFieldUpdateRequest) MergeAssetTypeObjectFieldUpdatePayload(v AssetTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t CampaignFieldUpdateRequest) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CampaignFieldUpdateRequest) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCoreFieldDefEditorMetadata0 returns the union data inside the CoreFieldDef_EditorMetadata as a CoreFieldDefEditorMetadata0
func (t CoreFieldDef_EditorMetadata) AsCoreFieldDefEditorMetadata0() (CoreFieldDefEditorMetadata0, error) {
	var body CoreFieldDefEditorMetadata0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCoreFieldDefEditorMetadata0 overwrites any union data inside the CoreFieldDef_EditorMetadata as the provided CoreFieldDefEditorMetadata0
func (t *CoreFieldDef_EditorMetadata) FromCoreFieldDefEditorMetadata0(v CoreFieldDefEditorMetadata0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCoreFieldDefEditorMetadata0 performs a merge with any union data inside the CoreFieldDef_EditorMetadata, using the provided CoreFieldDefEditorMetadata0
func (t *CoreFieldDef_EditorMetadata) MergeCoreFieldDefEditorMetadata0(v CoreFieldDefEditorMetadata0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCoreFieldDefEditorMetadata1 returns the union data inside the CoreFieldDef_EditorMetadata as a CoreFieldDefEditorMetadata1
func (t CoreFieldDef_EditorMetadata) AsCoreFieldDefEditorMetadata1() (CoreFieldDefEditorMetadata1, error) {
	var body CoreFieldDefEditorMetadata1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCoreFieldDefEditorMetadata1 overwrites any union data inside the CoreFieldDef_EditorMetadata as the provided CoreFieldDefEditorMetadata1
func (t *CoreFieldDef_EditorMetadata) FromCoreFieldDefEditorMetadata1(v CoreFieldDefEditorMetadata1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCoreFieldDefEditorMetadata1 performs a merge with any union data inside the CoreFieldDef_EditorMetadata, using the provided CoreFieldDefEditorMetadata1
func (t *CoreFieldDef_EditorMetadata) MergeCoreFieldDefEditorMetadata1(v CoreFieldDefEditorMetadata1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t CoreFieldDef_EditorMetadata) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CoreFieldDef_EditorMetadata) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsStringTypeObjectField returns the union data inside the EventCreateRequest_Fields_Item as a StringTypeObjectField
func (t EventCreateRequest_Fields_Item) AsStringTypeObjectField() (StringTypeObjectField, error) {
	var body StringTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromStringTypeObjectField overwrites any union data inside the EventCreateRequest_Fields_Item as the provided StringTypeObjectField
func (t *EventCreateRequest_Fields_Item) FromStringTypeObjectField(v StringTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeStringTypeObjectField performs a merge with any union data inside the EventCreateRequest_Fields_Item, using the provided StringTypeObjectField
func (t *EventCreateRequest_Fields_Item) MergeStringTypeObjectField(v StringTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMultiChoiceTypeObjectField returns the union data inside the EventCreateRequest_Fields_Item as a MultiChoiceTypeObjectField
func (t EventCreateRequest_Fields_Item) AsMultiChoiceTypeObjectField() (MultiChoiceTypeObjectField, error) {
	var body MultiChoiceTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMultiChoiceTypeObjectField overwrites any union data inside the EventCreateRequest_Fields_Item as the provided MultiChoiceTypeObjectField
func (t *EventCreateRequest_Fields_Item) FromMultiChoiceTypeObjectField(v MultiChoiceTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMultiChoiceTypeObjectField performs a merge with any union data inside the EventCreateRequest_Fields_Item, using the provided MultiChoiceTypeObjectField
func (t *EventCreateRequest_Fields_Item) MergeMultiChoiceTypeObjectField(v MultiChoiceTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDropdownTypeObjectField returns the union data inside the EventCreateRequest_Fields_Item as a DropdownTypeObjectField
func (t EventCreateRequest_Fields_Item) AsDropdownTypeObjectField() (DropdownTypeObjectField, error) {
	var body DropdownTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDropdownTypeObjectField overwrites any union data inside the EventCreateRequest_Fields_Item as the provided DropdownTypeObjectField
func (t *EventCreateRequest_Fields_Item) FromDropdownTypeObjectField(v DropdownTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDropdownTypeObjectField performs a merge with any union data inside the EventCreateRequest_Fields_Item, using the provided DropdownTypeObjectField
func (t *EventCreateRequest_Fields_Item) MergeDropdownTypeObjectField(v DropdownTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRadioButtonTypeObjectField returns the union data inside the EventCreateRequest_Fields_Item as a RadioButtonTypeObjectField
func (t EventCreateRequest_Fields_Item) AsRadioButtonTypeObjectField() (RadioButtonTypeObjectField, error) {
	var body RadioButtonTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRadioButtonTypeObjectField overwrites any union data inside the EventCreateRequest_Fields_Item as the provided RadioButtonTypeObjectField
func (t *EventCreateRequest_Fields_Item) FromRadioButtonTypeObjectField(v RadioButtonTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRadioButtonTypeObjectField performs a merge with any union data inside the EventCreateRequest_Fields_Item, using the provided RadioButtonTypeObjectField
func (t *EventCreateRequest_Fields_Item) MergeRadioButtonTypeObjectField(v RadioButtonTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsNumberTypeObjectField returns the union data inside the EventCreateRequest_Fields_Item as a NumberTypeObjectField
func (t EventCreateRequest_Fields_Item) AsNumberTypeObjectField() (NumberTypeObjectField, error) {
	var body NumberTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberTypeObjectField overwrites any union data inside the EventCreateRequest_Fields_Item as the provided NumberTypeObjectField
func (t *EventCreateRequest_Fields_Item) FromNumberTypeObjectField(v NumberTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberTypeObjectField performs a merge with any union data inside the EventCreateRequest_Fields_Item, using the provided NumberTypeObjectField
func (t *EventCreateRequest_Fields_Item) MergeNumberTypeObjectField(v NumberTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDateTypeObjectField returns the union data inside the EventCreateRequest_Fields_Item as a DateTypeObjectField
func (t EventCreateRequest_Fields_Item) AsDateTypeObjectField() (DateTypeObjectField, error) {
	var body DateTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDateTypeObjectField overwrites any union data inside the EventCreateRequest_Fields_Item as the provided DateTypeObjectField
func (t *EventCreateRequest_Fields_Item) FromDateTypeObjectField(v DateTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDateTypeObjectField performs a merge with any union data inside the EventCreateRequest_Fields_Item, using the provided DateTypeObjectField
func (t *EventCreateRequest_Fields_Item) MergeDateTypeObjectField(v DateTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAssetTypeObjectField returns the union data inside the EventCreateRequest_Fields_Item as a AssetTypeObjectField
func (t EventCreateRequest_Fields_Item) AsAssetTypeObjectField() (AssetTypeObjectField, error) {
	var body AssetTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetTypeObjectField overwrites any union data inside the EventCreateRequest_Fields_Item as the provided AssetTypeObjectField
func (t *EventCreateRequest_Fields_Item) FromAssetTypeObjectField(v AssetTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetTypeObjectField performs a merge with any union data inside the EventCreateRequest_Fields_Item, using the provided AssetTypeObjectField
func (t *EventCreateRequest_Fields_Item) MergeAssetTypeObjectField(v AssetTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t EventCreateRequest_Fields_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *EventCreateRequest_Fields_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsMultiChoiceTypeObjectField returns the union data inside the EventFieldsUpdateRequest_Item as a MultiChoiceTypeObjectField
func (t EventFieldsUpdateRequest_Item) AsMultiChoiceTypeObjectField() (MultiChoiceTypeObjectField, error) {
	var body MultiChoiceTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMultiChoiceTypeObjectField overwrites any union data inside the EventFieldsUpdateRequest_Item as the provided MultiChoiceTypeObjectField
func (t *EventFieldsUpdateRequest_Item) FromMultiChoiceTypeObjectField(v MultiChoiceTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMultiChoiceTypeObjectField performs a merge with any union data inside the EventFieldsUpdateRequest_Item, using the provided MultiChoiceTypeObjectField
func (t *EventFieldsUpdateRequest_Item) MergeMultiChoiceTypeObjectField(v MultiChoiceTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsStringTypeObjectField returns the union data inside the EventFieldsUpdateRequest_Item as a StringTypeObjectField
func (t EventFieldsUpdateRequest_Item) AsStringTypeObjectField() (StringTypeObjectField, error) {
	var body StringTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromStringTypeObjectField overwrites any union data inside the EventFieldsUpdateRequest_Item as the provided StringTypeObjectField
func (t *EventFieldsUpdateRequest_Item) FromStringTypeObjectField(v StringTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeStringTypeObjectField performs a merge with any union data inside the EventFieldsUpdateRequest_Item, using the provided StringTypeObjectField
func (t *EventFieldsUpdateRequest_Item) MergeStringTypeObjectField(v StringTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDropdownTypeObjectField returns the union data inside the EventFieldsUpdateRequest_Item as a DropdownTypeObjectField
func (t EventFieldsUpdateRequest_Item) AsDropdownTypeObjectField() (DropdownTypeObjectField, error) {
	var body DropdownTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDropdownTypeObjectField overwrites any union data inside the EventFieldsUpdateRequest_Item as the provided DropdownTypeObjectField
func (t *EventFieldsUpdateRequest_Item) FromDropdownTypeObjectField(v DropdownTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDropdownTypeObjectField performs a merge with any union data inside the EventFieldsUpdateRequest_Item, using the provided DropdownTypeObjectField
func (t *EventFieldsUpdateRequest_Item) MergeDropdownTypeObjectField(v DropdownTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRadioButtonTypeObjectField returns the union data inside the EventFieldsUpdateRequest_Item as a RadioButtonTypeObjectField
func (t EventFieldsUpdateRequest_Item) AsRadioButtonTypeObjectField() (RadioButtonTypeObjectField, error) {
	var body RadioButtonTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRadioButtonTypeObjectField overwrites any union data inside the EventFieldsUpdateRequest_Item as the provided RadioButtonTypeObjectField
func (t *EventFieldsUpdateRequest_Item) FromRadioButtonTypeObjectField(v RadioButtonTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRadioButtonTypeObjectField performs a merge with any union data inside the EventFieldsUpdateRequest_Item, using the provided RadioButtonTypeObjectField
func (t *EventFieldsUpdateRequest_Item) MergeRadioButtonTypeObjectField(v RadioButtonTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsNumberTypeObjectField returns the union data inside the EventFieldsUpdateRequest_Item as a NumberTypeObjectField
func (t EventFieldsUpdateRequest_Item) AsNumberTypeObjectField() (NumberTypeObjectField, error) {
	var body NumberTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberTypeObjectField overwrites any union data inside the EventFieldsUpdateRequest_Item as the provided NumberTypeObjectField
func (t *EventFieldsUpdateRequest_Item) FromNumberTypeObjectField(v NumberTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberTypeObjectField performs a merge with any union data inside the EventFieldsUpdateRequest_Item, using the provided NumberTypeObjectField
func (t *EventFieldsUpdateRequest_Item) MergeNumberTypeObjectField(v NumberTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDateTypeObjectField returns the union data inside the EventFieldsUpdateRequest_Item as a DateTypeObjectField
func (t EventFieldsUpdateRequest_Item) AsDateTypeObjectField() (DateTypeObjectField, error) {
	var body DateTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDateTypeObjectField overwrites any union data inside the EventFieldsUpdateRequest_Item as the provided DateTypeObjectField
func (t *EventFieldsUpdateRequest_Item) FromDateTypeObjectField(v DateTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDateTypeObjectField performs a merge with any union data inside the EventFieldsUpdateRequest_Item, using the provided DateTypeObjectField
func (t *EventFieldsUpdateRequest_Item) MergeDateTypeObjectField(v DateTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAssetTypeObjectField returns the union data inside the EventFieldsUpdateRequest_Item as a AssetTypeObjectField
func (t EventFieldsUpdateRequest_Item) AsAssetTypeObjectField() (AssetTypeObjectField, error) {
	var body AssetTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetTypeObjectField overwrites any union data inside the EventFieldsUpdateRequest_Item as the provided AssetTypeObjectField
func (t *EventFieldsUpdateRequest_Item) FromAssetTypeObjectField(v AssetTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetTypeObjectField performs a merge with any union data inside the EventFieldsUpdateRequest_Item, using the provided AssetTypeObjectField
func (t *EventFieldsUpdateRequest_Item) MergeAssetTypeObjectField(v AssetTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t EventFieldsUpdateRequest_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *EventFieldsUpdateRequest_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFieldTypeCommon returns the union data inside the FieldListResponseItem as a FieldTypeCommon
func (t FieldListResponseItem) AsFieldTypeCommon() (FieldTypeCommon, error) {
	var body FieldTypeCommon
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeCommon overwrites any union data inside the FieldListResponseItem as the provided FieldTypeCommon
func (t *FieldListResponseItem) FromFieldTypeCommon(v FieldTypeCommon) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeCommon performs a merge with any union data inside the FieldListResponseItem, using the provided FieldTypeCommon
func (t *FieldListResponseItem) MergeFieldTypeCommon(v FieldTypeCommon) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypeLabel returns the union data inside the FieldListResponseItem as a FieldTypeLabel
func (t FieldListResponseItem) AsFieldTypeLabel() (FieldTypeLabel, error) {
	var body FieldTypeLabel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeLabel overwrites any union data inside the FieldListResponseItem as the provided FieldTypeLabel
func (t *FieldListResponseItem) FromFieldTypeLabel(v FieldTypeLabel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeLabel performs a merge with any union data inside the FieldListResponseItem, using the provided FieldTypeLabel
func (t *FieldListResponseItem) MergeFieldTypeLabel(v FieldTypeLabel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypeDropdown returns the union data inside the FieldListResponseItem as a FieldTypeDropdown
func (t FieldListResponseItem) AsFieldTypeDropdown() (FieldTypeDropdown, error) {
	var body FieldTypeDropdown
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeDropdown overwrites any union data inside the FieldListResponseItem as the provided FieldTypeDropdown
func (t *FieldListResponseItem) FromFieldTypeDropdown(v FieldTypeDropdown) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeDropdown performs a merge with any union data inside the FieldListResponseItem, using the provided FieldTypeDropdown
func (t *FieldListResponseItem) MergeFieldTypeDropdown(v FieldTypeDropdown) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypeRadio returns the union data inside the FieldListResponseItem as a FieldTypeRadio
func (t FieldListResponseItem) AsFieldTypeRadio() (FieldTypeRadio, error) {
	var body FieldTypeRadio
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeRadio overwrites any union data inside the FieldListResponseItem as the provided FieldTypeRadio
func (t *FieldListResponseItem) FromFieldTypeRadio(v FieldTypeRadio) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeRadio performs a merge with any union data inside the FieldListResponseItem, using the provided FieldTypeRadio
func (t *FieldListResponseItem) MergeFieldTypeRadio(v FieldTypeRadio) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypeCheckbox returns the union data inside the FieldListResponseItem as a FieldTypeCheckbox
func (t FieldListResponseItem) AsFieldTypeCheckbox() (FieldTypeCheckbox, error) {
	var body FieldTypeCheckbox
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeCheckbox overwrites any union data inside the FieldListResponseItem as the provided FieldTypeCheckbox
func (t *FieldListResponseItem) FromFieldTypeCheckbox(v FieldTypeCheckbox) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeCheckbox performs a merge with any union data inside the FieldListResponseItem, using the provided FieldTypeCheckbox
func (t *FieldListResponseItem) MergeFieldTypeCheckbox(v FieldTypeCheckbox) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypeSimpleNumber returns the union data inside the FieldListResponseItem as a FieldTypeSimpleNumber
func (t FieldListResponseItem) AsFieldTypeSimpleNumber() (FieldTypeSimpleNumber, error) {
	var body FieldTypeSimpleNumber
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeSimpleNumber overwrites any union data inside the FieldListResponseItem as the provided FieldTypeSimpleNumber
func (t *FieldListResponseItem) FromFieldTypeSimpleNumber(v FieldTypeSimpleNumber) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeSimpleNumber performs a merge with any union data inside the FieldListResponseItem, using the provided FieldTypeSimpleNumber
func (t *FieldListResponseItem) MergeFieldTypeSimpleNumber(v FieldTypeSimpleNumber) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypePercentageNumber returns the union data inside the FieldListResponseItem as a FieldTypePercentageNumber
func (t FieldListResponseItem) AsFieldTypePercentageNumber() (FieldTypePercentageNumber, error) {
	var body FieldTypePercentageNumber
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypePercentageNumber overwrites any union data inside the FieldListResponseItem as the provided FieldTypePercentageNumber
func (t *FieldListResponseItem) FromFieldTypePercentageNumber(v FieldTypePercentageNumber) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypePercentageNumber performs a merge with any union data inside the FieldListResponseItem, using the provided FieldTypePercentageNumber
func (t *FieldListResponseItem) MergeFieldTypePercentageNumber(v FieldTypePercentageNumber) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypeCurrencyNumber returns the union data inside the FieldListResponseItem as a FieldTypeCurrencyNumber
func (t FieldListResponseItem) AsFieldTypeCurrencyNumber() (FieldTypeCurrencyNumber, error) {
	var body FieldTypeCurrencyNumber
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeCurrencyNumber overwrites any union data inside the FieldListResponseItem as the provided FieldTypeCurrencyNumber
func (t *FieldListResponseItem) FromFieldTypeCurrencyNumber(v FieldTypeCurrencyNumber) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeCurrencyNumber performs a merge with any union data inside the FieldListResponseItem, using the provided FieldTypeCurrencyNumber
func (t *FieldListResponseItem) MergeFieldTypeCurrencyNumber(v FieldTypeCurrencyNumber) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FieldListResponseItem) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FieldListResponseItem) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFolderUpdateRequest0 returns the union data inside the FolderUpdateRequest as a FolderUpdateRequest0
func (t FolderUpdateRequest) AsFolderUpdateRequest0() (FolderUpdateRequest0, error) {
	var body FolderUpdateRequest0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFolderUpdateRequest0 overwrites any union data inside the FolderUpdateRequest as the provided FolderUpdateRequest0
func (t *FolderUpdateRequest) FromFolderUpdateRequest0(v FolderUpdateRequest0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFolderUpdateRequest0 performs a merge with any union data inside the FolderUpdateRequest, using the provided FolderUpdateRequest0
func (t *FolderUpdateRequest) MergeFolderUpdateRequest0(v FolderUpdateRequest0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFolderUpdateRequest1 returns the union data inside the FolderUpdateRequest as a FolderUpdateRequest1
func (t FolderUpdateRequest) AsFolderUpdateRequest1() (FolderUpdateRequest1, error) {
	var body FolderUpdateRequest1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFolderUpdateRequest1 overwrites any union data inside the FolderUpdateRequest as the provided FolderUpdateRequest1
func (t *FolderUpdateRequest) FromFolderUpdateRequest1(v FolderUpdateRequest1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFolderUpdateRequest1 performs a merge with any union data inside the FolderUpdateRequest, using the provided FolderUpdateRequest1
func (t *FolderUpdateRequest) MergeFolderUpdateRequest1(v FolderUpdateRequest1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FolderUpdateRequest) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	if err != nil {
		return nil, err
	}
	object := make(map[string]json.RawMessage)
	if t.union != nil {
		err = json.Unmarshal(b, &object)
		if err != nil {
			return nil, err
		}
	}

	if t.Name != nil {
		object["name"], err = json.Marshal(t.Name)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'name': %w", err)
		}
	}

	if t.ParentFolderID != nil {
		object["parent_folder_id"], err = json.Marshal(t.ParentFolderID)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'parent_folder_id': %w", err)
		}
	}
	b, err = json.Marshal(object)
	return b, err
}

func (t *FolderUpdateRequest) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	if err != nil {
		return err
	}
	object := make(map[string]json.RawMessage)
	err = json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &t.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
	}

	if raw, found := object["parent_folder_id"]; found {
		err = json.Unmarshal(raw, &t.ParentFolderID)
		if err != nil {
			return fmt.Errorf("error reading 'parent_folder_id': %w", err)
		}
	}

	return err
}

// AsJSONFieldValueModelJSONValue0 returns the union data inside the JSONFieldValueModel_JSONValue as a JSONFieldValueModelJSONValue0
func (t JSONFieldValueModel_JSONValue) AsJSONFieldValueModelJSONValue0() (JSONFieldValueModelJSONValue0, error) {
	var body JSONFieldValueModelJSONValue0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromJSONFieldValueModelJSONValue0 overwrites any union data inside the JSONFieldValueModel_JSONValue as the provided JSONFieldValueModelJSONValue0
func (t *JSONFieldValueModel_JSONValue) FromJSONFieldValueModelJSONValue0(v JSONFieldValueModelJSONValue0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeJSONFieldValueModelJSONValue0 performs a merge with any union data inside the JSONFieldValueModel_JSONValue, using the provided JSONFieldValueModelJSONValue0
func (t *JSONFieldValueModel_JSONValue) MergeJSONFieldValueModelJSONValue0(v JSONFieldValueModelJSONValue0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsJSONFieldValueModelJSONValue1 returns the union data inside the JSONFieldValueModel_JSONValue as a JSONFieldValueModelJSONValue1
func (t JSONFieldValueModel_JSONValue) AsJSONFieldValueModelJSONValue1() (JSONFieldValueModelJSONValue1, error) {
	var body JSONFieldValueModelJSONValue1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromJSONFieldValueModelJSONValue1 overwrites any union data inside the JSONFieldValueModel_JSONValue as the provided JSONFieldValueModelJSONValue1
func (t *JSONFieldValueModel_JSONValue) FromJSONFieldValueModelJSONValue1(v JSONFieldValueModelJSONValue1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeJSONFieldValueModelJSONValue1 performs a merge with any union data inside the JSONFieldValueModel_JSONValue, using the provided JSONFieldValueModelJSONValue1
func (t *JSONFieldValueModel_JSONValue) MergeJSONFieldValueModelJSONValue1(v JSONFieldValueModelJSONValue1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t JSONFieldValueModel_JSONValue) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *JSONFieldValueModel_JSONValue) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsBaseSettingsFieldsResponse returns the union data inside the ListFieldsResponse_Data_Item as a BaseSettingsFieldsResponse
func (t ListFieldsResponse_Data_Item) AsBaseSettingsFieldsResponse() (BaseSettingsFieldsResponse, error) {
	var body BaseSettingsFieldsResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBaseSettingsFieldsResponse overwrites any union data inside the ListFieldsResponse_Data_Item as the provided BaseSettingsFieldsResponse
func (t *ListFieldsResponse_Data_Item) FromBaseSettingsFieldsResponse(v BaseSettingsFieldsResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBaseSettingsFieldsResponse performs a merge with any union data inside the ListFieldsResponse_Data_Item, using the provided BaseSettingsFieldsResponse
func (t *ListFieldsResponse_Data_Item) MergeBaseSettingsFieldsResponse(v BaseSettingsFieldsResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsLabelTypeSettingsFieldResponse returns the union data inside the ListFieldsResponse_Data_Item as a LabelTypeSettingsFieldResponse
func (t ListFieldsResponse_Data_Item) AsLabelTypeSettingsFieldResponse() (LabelTypeSettingsFieldResponse, error) {
	var body LabelTypeSettingsFieldResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLabelTypeSettingsFieldResponse overwrites any union data inside the ListFieldsResponse_Data_Item as the provided LabelTypeSettingsFieldResponse
func (t *ListFieldsResponse_Data_Item) FromLabelTypeSettingsFieldResponse(v LabelTypeSettingsFieldResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLabelTypeSettingsFieldResponse performs a merge with any union data inside the ListFieldsResponse_Data_Item, using the provided LabelTypeSettingsFieldResponse
func (t *ListFieldsResponse_Data_Item) MergeLabelTypeSettingsFieldResponse(v LabelTypeSettingsFieldResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDropdownTypeSettingsFieldResponse returns the union data inside the ListFieldsResponse_Data_Item as a DropdownTypeSettingsFieldResponse
func (t ListFieldsResponse_Data_Item) AsDropdownTypeSettingsFieldResponse() (DropdownTypeSettingsFieldResponse, error) {
	var body DropdownTypeSettingsFieldResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDropdownTypeSettingsFieldResponse overwrites any union data inside the ListFieldsResponse_Data_Item as the provided DropdownTypeSettingsFieldResponse
func (t *ListFieldsResponse_Data_Item) FromDropdownTypeSettingsFieldResponse(v DropdownTypeSettingsFieldResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDropdownTypeSettingsFieldResponse performs a merge with any union data inside the ListFieldsResponse_Data_Item, using the provided DropdownTypeSettingsFieldResponse
func (t *ListFieldsResponse_Data_Item) MergeDropdownTypeSettingsFieldResponse(v DropdownTypeSettingsFieldResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCheckboxAndRadioTypeSettingsFieldResponse returns the union data inside the ListFieldsResponse_Data_Item as a CheckboxAndRadioTypeSettingsFieldResponse
func (t ListFieldsResponse_Data_Item) AsCheckboxAndRadioTypeSettingsFieldResponse() (CheckboxAndRadioTypeSettingsFieldResponse, error) {
	var body CheckboxAndRadioTypeSettingsFieldResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCheckboxAndRadioTypeSettingsFieldResponse overwrites any union data inside the ListFieldsResponse_Data_Item as the provided CheckboxAndRadioTypeSettingsFieldResponse
func (t *ListFieldsResponse_Data_Item) FromCheckboxAndRadioTypeSettingsFieldResponse(v CheckboxAndRadioTypeSettingsFieldResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCheckboxAndRadioTypeSettingsFieldResponse performs a merge with any union data inside the ListFieldsResponse_Data_Item, using the provided CheckboxAndRadioTypeSettingsFieldResponse
func (t *ListFieldsResponse_Data_Item) MergeCheckboxAndRadioTypeSettingsFieldResponse(v CheckboxAndRadioTypeSettingsFieldResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsGenericNumberTypeSettingsFieldResponse returns the union data inside the ListFieldsResponse_Data_Item as a GenericNumberTypeSettingsFieldResponse
func (t ListFieldsResponse_Data_Item) AsGenericNumberTypeSettingsFieldResponse() (GenericNumberTypeSettingsFieldResponse, error) {
	var body GenericNumberTypeSettingsFieldResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGenericNumberTypeSettingsFieldResponse overwrites any union data inside the ListFieldsResponse_Data_Item as the provided GenericNumberTypeSettingsFieldResponse
func (t *ListFieldsResponse_Data_Item) FromGenericNumberTypeSettingsFieldResponse(v GenericNumberTypeSettingsFieldResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGenericNumberTypeSettingsFieldResponse performs a merge with any union data inside the ListFieldsResponse_Data_Item, using the provided GenericNumberTypeSettingsFieldResponse
func (t *ListFieldsResponse_Data_Item) MergeGenericNumberTypeSettingsFieldResponse(v GenericNumberTypeSettingsFieldResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCurrencyNumberTypeSettingsFieldResponse returns the union data inside the ListFieldsResponse_Data_Item as a CurrencyNumberTypeSettingsFieldResponse
func (t ListFieldsResponse_Data_Item) AsCurrencyNumberTypeSettingsFieldResponse() (CurrencyNumberTypeSettingsFieldResponse, error) {
	var body CurrencyNumberTypeSettingsFieldResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCurrencyNumberTypeSettingsFieldResponse overwrites any union data inside the ListFieldsResponse_Data_Item as the provided CurrencyNumberTypeSettingsFieldResponse
func (t *ListFieldsResponse_Data_Item) FromCurrencyNumberTypeSettingsFieldResponse(v CurrencyNumberTypeSettingsFieldResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCurrencyNumberTypeSettingsFieldResponse performs a merge with any union data inside the ListFieldsResponse_Data_Item, using the provided CurrencyNumberTypeSettingsFieldResponse
func (t *ListFieldsResponse_Data_Item) MergeCurrencyNumberTypeSettingsFieldResponse(v CurrencyNumberTypeSettingsFieldResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t ListFieldsResponse_Data_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ListFieldsResponse_Data_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFieldTypeLabel returns the union data inside the ListTaskSubStepFieldsResponse_Data_Item as a FieldTypeLabel
func (t ListTaskSubStepFieldsResponse_Data_Item) AsFieldTypeLabel() (FieldTypeLabel, error) {
	var body FieldTypeLabel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeLabel overwrites any union data inside the ListTaskSubStepFieldsResponse_Data_Item as the provided FieldTypeLabel
func (t *ListTaskSubStepFieldsResponse_Data_Item) FromFieldTypeLabel(v FieldTypeLabel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeLabel performs a merge with any union data inside the ListTaskSubStepFieldsResponse_Data_Item, using the provided FieldTypeLabel
func (t *ListTaskSubStepFieldsResponse_Data_Item) MergeFieldTypeLabel(v FieldTypeLabel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypeDropdown returns the union data inside the ListTaskSubStepFieldsResponse_Data_Item as a FieldTypeDropdown
func (t ListTaskSubStepFieldsResponse_Data_Item) AsFieldTypeDropdown() (FieldTypeDropdown, error) {
	var body FieldTypeDropdown
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeDropdown overwrites any union data inside the ListTaskSubStepFieldsResponse_Data_Item as the provided FieldTypeDropdown
func (t *ListTaskSubStepFieldsResponse_Data_Item) FromFieldTypeDropdown(v FieldTypeDropdown) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeDropdown performs a merge with any union data inside the ListTaskSubStepFieldsResponse_Data_Item, using the provided FieldTypeDropdown
func (t *ListTaskSubStepFieldsResponse_Data_Item) MergeFieldTypeDropdown(v FieldTypeDropdown) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypeRadio returns the union data inside the ListTaskSubStepFieldsResponse_Data_Item as a FieldTypeRadio
func (t ListTaskSubStepFieldsResponse_Data_Item) AsFieldTypeRadio() (FieldTypeRadio, error) {
	var body FieldTypeRadio
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeRadio overwrites any union data inside the ListTaskSubStepFieldsResponse_Data_Item as the provided FieldTypeRadio
func (t *ListTaskSubStepFieldsResponse_Data_Item) FromFieldTypeRadio(v FieldTypeRadio) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeRadio performs a merge with any union data inside the ListTaskSubStepFieldsResponse_Data_Item, using the provided FieldTypeRadio
func (t *ListTaskSubStepFieldsResponse_Data_Item) MergeFieldTypeRadio(v FieldTypeRadio) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypeCheckbox returns the union data inside the ListTaskSubStepFieldsResponse_Data_Item as a FieldTypeCheckbox
func (t ListTaskSubStepFieldsResponse_Data_Item) AsFieldTypeCheckbox() (FieldTypeCheckbox, error) {
	var body FieldTypeCheckbox
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeCheckbox overwrites any union data inside the ListTaskSubStepFieldsResponse_Data_Item as the provided FieldTypeCheckbox
func (t *ListTaskSubStepFieldsResponse_Data_Item) FromFieldTypeCheckbox(v FieldTypeCheckbox) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeCheckbox performs a merge with any union data inside the ListTaskSubStepFieldsResponse_Data_Item, using the provided FieldTypeCheckbox
func (t *ListTaskSubStepFieldsResponse_Data_Item) MergeFieldTypeCheckbox(v FieldTypeCheckbox) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t ListTaskSubStepFieldsResponse_Data_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ListTaskSubStepFieldsResponse_Data_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsNumberFieldValueModel returns the union data inside the LocalizedFieldValues_FieldValues_Item as a NumberFieldValueModel
func (t LocalizedFieldValues_FieldValues_Item) AsNumberFieldValueModel() (NumberFieldValueModel, error) {
	var body NumberFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberFieldValueModel overwrites any union data inside the LocalizedFieldValues_FieldValues_Item as the provided NumberFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) FromNumberFieldValueModel(v NumberFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberFieldValueModel performs a merge with any union data inside the LocalizedFieldValues_FieldValues_Item, using the provided NumberFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) MergeNumberFieldValueModel(v NumberFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBooleanFieldValueModel returns the union data inside the LocalizedFieldValues_FieldValues_Item as a BooleanFieldValueModel
func (t LocalizedFieldValues_FieldValues_Item) AsBooleanFieldValueModel() (BooleanFieldValueModel, error) {
	var body BooleanFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBooleanFieldValueModel overwrites any union data inside the LocalizedFieldValues_FieldValues_Item as the provided BooleanFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) FromBooleanFieldValueModel(v BooleanFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBooleanFieldValueModel performs a merge with any union data inside the LocalizedFieldValues_FieldValues_Item, using the provided BooleanFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) MergeBooleanFieldValueModel(v BooleanFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDatetimeFieldValueModel returns the union data inside the LocalizedFieldValues_FieldValues_Item as a DatetimeFieldValueModel
func (t LocalizedFieldValues_FieldValues_Item) AsDatetimeFieldValueModel() (DatetimeFieldValueModel, error) {
	var body DatetimeFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDatetimeFieldValueModel overwrites any union data inside the LocalizedFieldValues_FieldValues_Item as the provided DatetimeFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) FromDatetimeFieldValueModel(v DatetimeFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDatetimeFieldValueModel performs a merge with any union data inside the LocalizedFieldValues_FieldValues_Item, using the provided DatetimeFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) MergeDatetimeFieldValueModel(v DatetimeFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTextFieldValueModel returns the union data inside the LocalizedFieldValues_FieldValues_Item as a TextFieldValueModel
func (t LocalizedFieldValues_FieldValues_Item) AsTextFieldValueModel() (TextFieldValueModel, error) {
	var body TextFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTextFieldValueModel overwrites any union data inside the LocalizedFieldValues_FieldValues_Item as the provided TextFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) FromTextFieldValueModel(v TextFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTextFieldValueModel performs a merge with any union data inside the LocalizedFieldValues_FieldValues_Item, using the provided TextFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) MergeTextFieldValueModel(v TextFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRichTextFieldValueModel returns the union data inside the LocalizedFieldValues_FieldValues_Item as a RichTextFieldValueModel
func (t LocalizedFieldValues_FieldValues_Item) AsRichTextFieldValueModel() (RichTextFieldValueModel, error) {
	var body RichTextFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRichTextFieldValueModel overwrites any union data inside the LocalizedFieldValues_FieldValues_Item as the provided RichTextFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) FromRichTextFieldValueModel(v RichTextFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRichTextFieldValueModel performs a merge with any union data inside the LocalizedFieldValues_FieldValues_Item, using the provided RichTextFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) MergeRichTextFieldValueModel(v RichTextFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsLibraryAssetFieldValueModel returns the union data inside the LocalizedFieldValues_FieldValues_Item as a LibraryAssetFieldValueModel
func (t LocalizedFieldValues_FieldValues_Item) AsLibraryAssetFieldValueModel() (LibraryAssetFieldValueModel, error) {
	var body LibraryAssetFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLibraryAssetFieldValueModel overwrites any union data inside the LocalizedFieldValues_FieldValues_Item as the provided LibraryAssetFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) FromLibraryAssetFieldValueModel(v LibraryAssetFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLibraryAssetFieldValueModel performs a merge with any union data inside the LocalizedFieldValues_FieldValues_Item, using the provided LibraryAssetFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) MergeLibraryAssetFieldValueModel(v LibraryAssetFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsContentFieldValueExpandedModel returns the union data inside the LocalizedFieldValues_FieldValues_Item as a ContentFieldValueExpandedModel
func (t LocalizedFieldValues_FieldValues_Item) AsContentFieldValueExpandedModel() (ContentFieldValueExpandedModel, error) {
	var body ContentFieldValueExpandedModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentFieldValueExpandedModel overwrites any union data inside the LocalizedFieldValues_FieldValues_Item as the provided ContentFieldValueExpandedModel
func (t *LocalizedFieldValues_FieldValues_Item) FromContentFieldValueExpandedModel(v ContentFieldValueExpandedModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentFieldValueExpandedModel performs a merge with any union data inside the LocalizedFieldValues_FieldValues_Item, using the provided ContentFieldValueExpandedModel
func (t *LocalizedFieldValues_FieldValues_Item) MergeContentFieldValueExpandedModel(v ContentFieldValueExpandedModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsContentFieldValueModel returns the union data inside the LocalizedFieldValues_FieldValues_Item as a ContentFieldValueModel
func (t LocalizedFieldValues_FieldValues_Item) AsContentFieldValueModel() (ContentFieldValueModel, error) {
	var body ContentFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentFieldValueModel overwrites any union data inside the LocalizedFieldValues_FieldValues_Item as the provided ContentFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) FromContentFieldValueModel(v ContentFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentFieldValueModel performs a merge with any union data inside the LocalizedFieldValues_FieldValues_Item, using the provided ContentFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) MergeContentFieldValueModel(v ContentFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsURLFieldValueModel returns the union data inside the LocalizedFieldValues_FieldValues_Item as a URLFieldValueModel
func (t LocalizedFieldValues_FieldValues_Item) AsURLFieldValueModel() (URLFieldValueModel, error) {
	var body URLFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromURLFieldValueModel overwrites any union data inside the LocalizedFieldValues_FieldValues_Item as the provided URLFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) FromURLFieldValueModel(v URLFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeURLFieldValueModel performs a merge with any union data inside the LocalizedFieldValues_FieldValues_Item, using the provided URLFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) MergeURLFieldValueModel(v URLFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsJSONFieldValueModel returns the union data inside the LocalizedFieldValues_FieldValues_Item as a JSONFieldValueModel
func (t LocalizedFieldValues_FieldValues_Item) AsJSONFieldValueModel() (JSONFieldValueModel, error) {
	var body JSONFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromJSONFieldValueModel overwrites any union data inside the LocalizedFieldValues_FieldValues_Item as the provided JSONFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) FromJSONFieldValueModel(v JSONFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeJSONFieldValueModel performs a merge with any union data inside the LocalizedFieldValues_FieldValues_Item, using the provided JSONFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) MergeJSONFieldValueModel(v JSONFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsChoiceFieldValueModel returns the union data inside the LocalizedFieldValues_FieldValues_Item as a ChoiceFieldValueModel
func (t LocalizedFieldValues_FieldValues_Item) AsChoiceFieldValueModel() (ChoiceFieldValueModel, error) {
	var body ChoiceFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromChoiceFieldValueModel overwrites any union data inside the LocalizedFieldValues_FieldValues_Item as the provided ChoiceFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) FromChoiceFieldValueModel(v ChoiceFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeChoiceFieldValueModel performs a merge with any union data inside the LocalizedFieldValues_FieldValues_Item, using the provided ChoiceFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) MergeChoiceFieldValueModel(v ChoiceFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsLocationFieldValueModel returns the union data inside the LocalizedFieldValues_FieldValues_Item as a LocationFieldValueModel
func (t LocalizedFieldValues_FieldValues_Item) AsLocationFieldValueModel() (LocationFieldValueModel, error) {
	var body LocationFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLocationFieldValueModel overwrites any union data inside the LocalizedFieldValues_FieldValues_Item as the provided LocationFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) FromLocationFieldValueModel(v LocationFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLocationFieldValueModel performs a merge with any union data inside the LocalizedFieldValues_FieldValues_Item, using the provided LocationFieldValueModel
func (t *LocalizedFieldValues_FieldValues_Item) MergeLocationFieldValueModel(v LocationFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t LocalizedFieldValues_FieldValues_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *LocalizedFieldValues_FieldValues_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsNumberFieldValueModel returns the union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a NumberFieldValueModel
func (t LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsNumberFieldValueModel() (NumberFieldValueModel, error) {
	var body NumberFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberFieldValueModel overwrites any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided NumberFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromNumberFieldValueModel(v NumberFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberFieldValueModel performs a merge with any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided NumberFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeNumberFieldValueModel(v NumberFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBooleanFieldValueModel returns the union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a BooleanFieldValueModel
func (t LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsBooleanFieldValueModel() (BooleanFieldValueModel, error) {
	var body BooleanFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBooleanFieldValueModel overwrites any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided BooleanFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromBooleanFieldValueModel(v BooleanFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBooleanFieldValueModel performs a merge with any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided BooleanFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeBooleanFieldValueModel(v BooleanFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDatetimeFieldValueModel returns the union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a DatetimeFieldValueModel
func (t LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsDatetimeFieldValueModel() (DatetimeFieldValueModel, error) {
	var body DatetimeFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDatetimeFieldValueModel overwrites any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided DatetimeFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromDatetimeFieldValueModel(v DatetimeFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDatetimeFieldValueModel performs a merge with any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided DatetimeFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeDatetimeFieldValueModel(v DatetimeFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTextFieldValueModel returns the union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a TextFieldValueModel
func (t LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsTextFieldValueModel() (TextFieldValueModel, error) {
	var body TextFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTextFieldValueModel overwrites any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided TextFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromTextFieldValueModel(v TextFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTextFieldValueModel performs a merge with any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided TextFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeTextFieldValueModel(v TextFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRichTextFieldValueModel returns the union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a RichTextFieldValueModel
func (t LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsRichTextFieldValueModel() (RichTextFieldValueModel, error) {
	var body RichTextFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRichTextFieldValueModel overwrites any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided RichTextFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromRichTextFieldValueModel(v RichTextFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRichTextFieldValueModel performs a merge with any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided RichTextFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeRichTextFieldValueModel(v RichTextFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsLibraryAssetFieldValueModel returns the union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a LibraryAssetFieldValueModel
func (t LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsLibraryAssetFieldValueModel() (LibraryAssetFieldValueModel, error) {
	var body LibraryAssetFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLibraryAssetFieldValueModel overwrites any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided LibraryAssetFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromLibraryAssetFieldValueModel(v LibraryAssetFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLibraryAssetFieldValueModel performs a merge with any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided LibraryAssetFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeLibraryAssetFieldValueModel(v LibraryAssetFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsContentFieldValueWithEmbeddedModel returns the union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a ContentFieldValueWithEmbeddedModel
func (t LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsContentFieldValueWithEmbeddedModel() (ContentFieldValueWithEmbeddedModel, error) {
	var body ContentFieldValueWithEmbeddedModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentFieldValueWithEmbeddedModel overwrites any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided ContentFieldValueWithEmbeddedModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromContentFieldValueWithEmbeddedModel(v ContentFieldValueWithEmbeddedModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentFieldValueWithEmbeddedModel performs a merge with any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided ContentFieldValueWithEmbeddedModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeContentFieldValueWithEmbeddedModel(v ContentFieldValueWithEmbeddedModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsContentFieldValueModel returns the union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a ContentFieldValueModel
func (t LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsContentFieldValueModel() (ContentFieldValueModel, error) {
	var body ContentFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentFieldValueModel overwrites any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided ContentFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromContentFieldValueModel(v ContentFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentFieldValueModel performs a merge with any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided ContentFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeContentFieldValueModel(v ContentFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsURLFieldValueModel returns the union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a URLFieldValueModel
func (t LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsURLFieldValueModel() (URLFieldValueModel, error) {
	var body URLFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromURLFieldValueModel overwrites any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided URLFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromURLFieldValueModel(v URLFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeURLFieldValueModel performs a merge with any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided URLFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeURLFieldValueModel(v URLFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsJSONFieldValueModel returns the union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a JSONFieldValueModel
func (t LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsJSONFieldValueModel() (JSONFieldValueModel, error) {
	var body JSONFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromJSONFieldValueModel overwrites any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided JSONFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromJSONFieldValueModel(v JSONFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeJSONFieldValueModel performs a merge with any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided JSONFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeJSONFieldValueModel(v JSONFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsChoiceFieldValueModel returns the union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a ChoiceFieldValueModel
func (t LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsChoiceFieldValueModel() (ChoiceFieldValueModel, error) {
	var body ChoiceFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromChoiceFieldValueModel overwrites any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided ChoiceFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromChoiceFieldValueModel(v ChoiceFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeChoiceFieldValueModel performs a merge with any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided ChoiceFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeChoiceFieldValueModel(v ChoiceFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsLocationFieldValueModel returns the union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a LocationFieldValueModel
func (t LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsLocationFieldValueModel() (LocationFieldValueModel, error) {
	var body LocationFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLocationFieldValueModel overwrites any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided LocationFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromLocationFieldValueModel(v LocationFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLocationFieldValueModel performs a merge with any union data inside the LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided LocationFieldValueModel
func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeLocationFieldValueModel(v LocationFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *LocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsStringTypeObjectField returns the union data inside the ObjectFieldCreateRequest as a StringTypeObjectField
func (t ObjectFieldCreateRequest) AsStringTypeObjectField() (StringTypeObjectField, error) {
	var body StringTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromStringTypeObjectField overwrites any union data inside the ObjectFieldCreateRequest as the provided StringTypeObjectField
func (t *ObjectFieldCreateRequest) FromStringTypeObjectField(v StringTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeStringTypeObjectField performs a merge with any union data inside the ObjectFieldCreateRequest, using the provided StringTypeObjectField
func (t *ObjectFieldCreateRequest) MergeStringTypeObjectField(v StringTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMultiChoiceTypeObjectField returns the union data inside the ObjectFieldCreateRequest as a MultiChoiceTypeObjectField
func (t ObjectFieldCreateRequest) AsMultiChoiceTypeObjectField() (MultiChoiceTypeObjectField, error) {
	var body MultiChoiceTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMultiChoiceTypeObjectField overwrites any union data inside the ObjectFieldCreateRequest as the provided MultiChoiceTypeObjectField
func (t *ObjectFieldCreateRequest) FromMultiChoiceTypeObjectField(v MultiChoiceTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMultiChoiceTypeObjectField performs a merge with any union data inside the ObjectFieldCreateRequest, using the provided MultiChoiceTypeObjectField
func (t *ObjectFieldCreateRequest) MergeMultiChoiceTypeObjectField(v MultiChoiceTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDropdownTypeObjectField returns the union data inside the ObjectFieldCreateRequest as a DropdownTypeObjectField
func (t ObjectFieldCreateRequest) AsDropdownTypeObjectField() (DropdownTypeObjectField, error) {
	var body DropdownTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDropdownTypeObjectField overwrites any union data inside the ObjectFieldCreateRequest as the provided DropdownTypeObjectField
func (t *ObjectFieldCreateRequest) FromDropdownTypeObjectField(v DropdownTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDropdownTypeObjectField performs a merge with any union data inside the ObjectFieldCreateRequest, using the provided DropdownTypeObjectField
func (t *ObjectFieldCreateRequest) MergeDropdownTypeObjectField(v DropdownTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRadioButtonTypeObjectField returns the union data inside the ObjectFieldCreateRequest as a RadioButtonTypeObjectField
func (t ObjectFieldCreateRequest) AsRadioButtonTypeObjectField() (RadioButtonTypeObjectField, error) {
	var body RadioButtonTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRadioButtonTypeObjectField overwrites any union data inside the ObjectFieldCreateRequest as the provided RadioButtonTypeObjectField
func (t *ObjectFieldCreateRequest) FromRadioButtonTypeObjectField(v RadioButtonTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRadioButtonTypeObjectField performs a merge with any union data inside the ObjectFieldCreateRequest, using the provided RadioButtonTypeObjectField
func (t *ObjectFieldCreateRequest) MergeRadioButtonTypeObjectField(v RadioButtonTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsNumberTypeObjectField returns the union data inside the ObjectFieldCreateRequest as a NumberTypeObjectField
func (t ObjectFieldCreateRequest) AsNumberTypeObjectField() (NumberTypeObjectField, error) {
	var body NumberTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberTypeObjectField overwrites any union data inside the ObjectFieldCreateRequest as the provided NumberTypeObjectField
func (t *ObjectFieldCreateRequest) FromNumberTypeObjectField(v NumberTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberTypeObjectField performs a merge with any union data inside the ObjectFieldCreateRequest, using the provided NumberTypeObjectField
func (t *ObjectFieldCreateRequest) MergeNumberTypeObjectField(v NumberTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDateTypeObjectField returns the union data inside the ObjectFieldCreateRequest as a DateTypeObjectField
func (t ObjectFieldCreateRequest) AsDateTypeObjectField() (DateTypeObjectField, error) {
	var body DateTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDateTypeObjectField overwrites any union data inside the ObjectFieldCreateRequest as the provided DateTypeObjectField
func (t *ObjectFieldCreateRequest) FromDateTypeObjectField(v DateTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDateTypeObjectField performs a merge with any union data inside the ObjectFieldCreateRequest, using the provided DateTypeObjectField
func (t *ObjectFieldCreateRequest) MergeDateTypeObjectField(v DateTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAssetTypeObjectField returns the union data inside the ObjectFieldCreateRequest as a AssetTypeObjectField
func (t ObjectFieldCreateRequest) AsAssetTypeObjectField() (AssetTypeObjectField, error) {
	var body AssetTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetTypeObjectField overwrites any union data inside the ObjectFieldCreateRequest as the provided AssetTypeObjectField
func (t *ObjectFieldCreateRequest) FromAssetTypeObjectField(v AssetTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetTypeObjectField performs a merge with any union data inside the ObjectFieldCreateRequest, using the provided AssetTypeObjectField
func (t *ObjectFieldCreateRequest) MergeAssetTypeObjectField(v AssetTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t ObjectFieldCreateRequest) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ObjectFieldCreateRequest) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFieldTypeCommon returns the union data inside the ObjectFieldCreateResponse as a FieldTypeCommon
func (t ObjectFieldCreateResponse) AsFieldTypeCommon() (FieldTypeCommon, error) {
	var body FieldTypeCommon
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeCommon overwrites any union data inside the ObjectFieldCreateResponse as the provided FieldTypeCommon
func (t *ObjectFieldCreateResponse) FromFieldTypeCommon(v FieldTypeCommon) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeCommon performs a merge with any union data inside the ObjectFieldCreateResponse, using the provided FieldTypeCommon
func (t *ObjectFieldCreateResponse) MergeFieldTypeCommon(v FieldTypeCommon) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypeLabel returns the union data inside the ObjectFieldCreateResponse as a FieldTypeLabel
func (t ObjectFieldCreateResponse) AsFieldTypeLabel() (FieldTypeLabel, error) {
	var body FieldTypeLabel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeLabel overwrites any union data inside the ObjectFieldCreateResponse as the provided FieldTypeLabel
func (t *ObjectFieldCreateResponse) FromFieldTypeLabel(v FieldTypeLabel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeLabel performs a merge with any union data inside the ObjectFieldCreateResponse, using the provided FieldTypeLabel
func (t *ObjectFieldCreateResponse) MergeFieldTypeLabel(v FieldTypeLabel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypeDropdown returns the union data inside the ObjectFieldCreateResponse as a FieldTypeDropdown
func (t ObjectFieldCreateResponse) AsFieldTypeDropdown() (FieldTypeDropdown, error) {
	var body FieldTypeDropdown
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeDropdown overwrites any union data inside the ObjectFieldCreateResponse as the provided FieldTypeDropdown
func (t *ObjectFieldCreateResponse) FromFieldTypeDropdown(v FieldTypeDropdown) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeDropdown performs a merge with any union data inside the ObjectFieldCreateResponse, using the provided FieldTypeDropdown
func (t *ObjectFieldCreateResponse) MergeFieldTypeDropdown(v FieldTypeDropdown) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypeRadio returns the union data inside the ObjectFieldCreateResponse as a FieldTypeRadio
func (t ObjectFieldCreateResponse) AsFieldTypeRadio() (FieldTypeRadio, error) {
	var body FieldTypeRadio
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeRadio overwrites any union data inside the ObjectFieldCreateResponse as the provided FieldTypeRadio
func (t *ObjectFieldCreateResponse) FromFieldTypeRadio(v FieldTypeRadio) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeRadio performs a merge with any union data inside the ObjectFieldCreateResponse, using the provided FieldTypeRadio
func (t *ObjectFieldCreateResponse) MergeFieldTypeRadio(v FieldTypeRadio) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypeCheckbox returns the union data inside the ObjectFieldCreateResponse as a FieldTypeCheckbox
func (t ObjectFieldCreateResponse) AsFieldTypeCheckbox() (FieldTypeCheckbox, error) {
	var body FieldTypeCheckbox
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeCheckbox overwrites any union data inside the ObjectFieldCreateResponse as the provided FieldTypeCheckbox
func (t *ObjectFieldCreateResponse) FromFieldTypeCheckbox(v FieldTypeCheckbox) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeCheckbox performs a merge with any union data inside the ObjectFieldCreateResponse, using the provided FieldTypeCheckbox
func (t *ObjectFieldCreateResponse) MergeFieldTypeCheckbox(v FieldTypeCheckbox) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypeSimpleNumber returns the union data inside the ObjectFieldCreateResponse as a FieldTypeSimpleNumber
func (t ObjectFieldCreateResponse) AsFieldTypeSimpleNumber() (FieldTypeSimpleNumber, error) {
	var body FieldTypeSimpleNumber
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeSimpleNumber overwrites any union data inside the ObjectFieldCreateResponse as the provided FieldTypeSimpleNumber
func (t *ObjectFieldCreateResponse) FromFieldTypeSimpleNumber(v FieldTypeSimpleNumber) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeSimpleNumber performs a merge with any union data inside the ObjectFieldCreateResponse, using the provided FieldTypeSimpleNumber
func (t *ObjectFieldCreateResponse) MergeFieldTypeSimpleNumber(v FieldTypeSimpleNumber) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypePercentageNumber returns the union data inside the ObjectFieldCreateResponse as a FieldTypePercentageNumber
func (t ObjectFieldCreateResponse) AsFieldTypePercentageNumber() (FieldTypePercentageNumber, error) {
	var body FieldTypePercentageNumber
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypePercentageNumber overwrites any union data inside the ObjectFieldCreateResponse as the provided FieldTypePercentageNumber
func (t *ObjectFieldCreateResponse) FromFieldTypePercentageNumber(v FieldTypePercentageNumber) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypePercentageNumber performs a merge with any union data inside the ObjectFieldCreateResponse, using the provided FieldTypePercentageNumber
func (t *ObjectFieldCreateResponse) MergeFieldTypePercentageNumber(v FieldTypePercentageNumber) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFieldTypeCurrencyNumber returns the union data inside the ObjectFieldCreateResponse as a FieldTypeCurrencyNumber
func (t ObjectFieldCreateResponse) AsFieldTypeCurrencyNumber() (FieldTypeCurrencyNumber, error) {
	var body FieldTypeCurrencyNumber
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFieldTypeCurrencyNumber overwrites any union data inside the ObjectFieldCreateResponse as the provided FieldTypeCurrencyNumber
func (t *ObjectFieldCreateResponse) FromFieldTypeCurrencyNumber(v FieldTypeCurrencyNumber) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFieldTypeCurrencyNumber performs a merge with any union data inside the ObjectFieldCreateResponse, using the provided FieldTypeCurrencyNumber
func (t *ObjectFieldCreateResponse) MergeFieldTypeCurrencyNumber(v FieldTypeCurrencyNumber) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t ObjectFieldCreateResponse) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ObjectFieldCreateResponse) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsNumberFieldValueModel returns the union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a NumberFieldValueModel
func (t PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsNumberFieldValueModel() (NumberFieldValueModel, error) {
	var body NumberFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberFieldValueModel overwrites any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided NumberFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromNumberFieldValueModel(v NumberFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberFieldValueModel performs a merge with any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided NumberFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeNumberFieldValueModel(v NumberFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBooleanFieldValueModel returns the union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a BooleanFieldValueModel
func (t PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsBooleanFieldValueModel() (BooleanFieldValueModel, error) {
	var body BooleanFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBooleanFieldValueModel overwrites any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided BooleanFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromBooleanFieldValueModel(v BooleanFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBooleanFieldValueModel performs a merge with any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided BooleanFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeBooleanFieldValueModel(v BooleanFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDatetimeFieldValueModel returns the union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a DatetimeFieldValueModel
func (t PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsDatetimeFieldValueModel() (DatetimeFieldValueModel, error) {
	var body DatetimeFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDatetimeFieldValueModel overwrites any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided DatetimeFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromDatetimeFieldValueModel(v DatetimeFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDatetimeFieldValueModel performs a merge with any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided DatetimeFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeDatetimeFieldValueModel(v DatetimeFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTextFieldValueModel returns the union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a TextFieldValueModel
func (t PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsTextFieldValueModel() (TextFieldValueModel, error) {
	var body TextFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTextFieldValueModel overwrites any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided TextFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromTextFieldValueModel(v TextFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTextFieldValueModel performs a merge with any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided TextFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeTextFieldValueModel(v TextFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRichTextFieldValueModel returns the union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a RichTextFieldValueModel
func (t PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsRichTextFieldValueModel() (RichTextFieldValueModel, error) {
	var body RichTextFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRichTextFieldValueModel overwrites any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided RichTextFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromRichTextFieldValueModel(v RichTextFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRichTextFieldValueModel performs a merge with any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided RichTextFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeRichTextFieldValueModel(v RichTextFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsLibraryAssetFieldValueModel returns the union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a LibraryAssetFieldValueModel
func (t PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsLibraryAssetFieldValueModel() (LibraryAssetFieldValueModel, error) {
	var body LibraryAssetFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLibraryAssetFieldValueModel overwrites any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided LibraryAssetFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromLibraryAssetFieldValueModel(v LibraryAssetFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLibraryAssetFieldValueModel performs a merge with any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided LibraryAssetFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeLibraryAssetFieldValueModel(v LibraryAssetFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsContentFieldValueWithEmbeddedModel returns the union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a ContentFieldValueWithEmbeddedModel
func (t PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsContentFieldValueWithEmbeddedModel() (ContentFieldValueWithEmbeddedModel, error) {
	var body ContentFieldValueWithEmbeddedModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentFieldValueWithEmbeddedModel overwrites any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided ContentFieldValueWithEmbeddedModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromContentFieldValueWithEmbeddedModel(v ContentFieldValueWithEmbeddedModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentFieldValueWithEmbeddedModel performs a merge with any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided ContentFieldValueWithEmbeddedModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeContentFieldValueWithEmbeddedModel(v ContentFieldValueWithEmbeddedModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsContentFieldValueModel returns the union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a ContentFieldValueModel
func (t PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsContentFieldValueModel() (ContentFieldValueModel, error) {
	var body ContentFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentFieldValueModel overwrites any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided ContentFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromContentFieldValueModel(v ContentFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentFieldValueModel performs a merge with any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided ContentFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeContentFieldValueModel(v ContentFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsURLFieldValueModel returns the union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a URLFieldValueModel
func (t PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsURLFieldValueModel() (URLFieldValueModel, error) {
	var body URLFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromURLFieldValueModel overwrites any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided URLFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromURLFieldValueModel(v URLFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeURLFieldValueModel performs a merge with any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided URLFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeURLFieldValueModel(v URLFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsJSONFieldValueModel returns the union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a JSONFieldValueModel
func (t PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsJSONFieldValueModel() (JSONFieldValueModel, error) {
	var body JSONFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromJSONFieldValueModel overwrites any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided JSONFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromJSONFieldValueModel(v JSONFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeJSONFieldValueModel performs a merge with any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided JSONFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeJSONFieldValueModel(v JSONFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsChoiceFieldValueModel returns the union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a ChoiceFieldValueModel
func (t PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsChoiceFieldValueModel() (ChoiceFieldValueModel, error) {
	var body ChoiceFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromChoiceFieldValueModel overwrites any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided ChoiceFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromChoiceFieldValueModel(v ChoiceFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeChoiceFieldValueModel performs a merge with any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided ChoiceFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeChoiceFieldValueModel(v ChoiceFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsLocationFieldValueModel returns the union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a LocationFieldValueModel
func (t PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsLocationFieldValueModel() (LocationFieldValueModel, error) {
	var body LocationFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLocationFieldValueModel overwrites any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided LocationFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromLocationFieldValueModel(v LocationFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLocationFieldValueModel performs a merge with any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided LocationFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeLocationFieldValueModel(v LocationFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDeleteFieldValueModel returns the union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a DeleteFieldValueModel
func (t PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsDeleteFieldValueModel() (DeleteFieldValueModel, error) {
	var body DeleteFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDeleteFieldValueModel overwrites any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided DeleteFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromDeleteFieldValueModel(v DeleteFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDeleteFieldValueModel performs a merge with any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided DeleteFieldValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeDeleteFieldValueModel(v DeleteFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsContentFieldWithEmbeddedPatchValueModel returns the union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as a ContentFieldWithEmbeddedPatchValueModel
func (t PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) AsContentFieldWithEmbeddedPatchValueModel() (ContentFieldWithEmbeddedPatchValueModel, error) {
	var body ContentFieldWithEmbeddedPatchValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentFieldWithEmbeddedPatchValueModel overwrites any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item as the provided ContentFieldWithEmbeddedPatchValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) FromContentFieldWithEmbeddedPatchValueModel(v ContentFieldWithEmbeddedPatchValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentFieldWithEmbeddedPatchValueModel performs a merge with any union data inside the PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item, using the provided ContentFieldWithEmbeddedPatchValueModel
func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MergeContentFieldWithEmbeddedPatchValueModel(v ContentFieldWithEmbeddedPatchValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PatchLocalizedFieldValuesWithEmbeddedContent_FieldValues_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsSCContentPreviewCompleteRequestKeyedPreviews0 returns the union data inside the SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties as a SCContentPreviewCompleteRequestKeyedPreviews0
func (t SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties) AsSCContentPreviewCompleteRequestKeyedPreviews0() (SCContentPreviewCompleteRequestKeyedPreviews0, error) {
	var body SCContentPreviewCompleteRequestKeyedPreviews0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSCContentPreviewCompleteRequestKeyedPreviews0 overwrites any union data inside the SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties as the provided SCContentPreviewCompleteRequestKeyedPreviews0
func (t *SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties) FromSCContentPreviewCompleteRequestKeyedPreviews0(v SCContentPreviewCompleteRequestKeyedPreviews0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSCContentPreviewCompleteRequestKeyedPreviews0 performs a merge with any union data inside the SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties, using the provided SCContentPreviewCompleteRequestKeyedPreviews0
func (t *SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties) MergeSCContentPreviewCompleteRequestKeyedPreviews0(v SCContentPreviewCompleteRequestKeyedPreviews0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsKeyedPreviewCompletedModel returns the union data inside the SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties as a KeyedPreviewCompletedModel
func (t SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties) AsKeyedPreviewCompletedModel() (KeyedPreviewCompletedModel, error) {
	var body KeyedPreviewCompletedModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromKeyedPreviewCompletedModel overwrites any union data inside the SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties as the provided KeyedPreviewCompletedModel
func (t *SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties) FromKeyedPreviewCompletedModel(v KeyedPreviewCompletedModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeKeyedPreviewCompletedModel performs a merge with any union data inside the SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties, using the provided KeyedPreviewCompletedModel
func (t *SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties) MergeKeyedPreviewCompletedModel(v KeyedPreviewCompletedModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsKeyedPreviewErrorModel returns the union data inside the SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties as a KeyedPreviewErrorModel
func (t SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties) AsKeyedPreviewErrorModel() (KeyedPreviewErrorModel, error) {
	var body KeyedPreviewErrorModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromKeyedPreviewErrorModel overwrites any union data inside the SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties as the provided KeyedPreviewErrorModel
func (t *SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties) FromKeyedPreviewErrorModel(v KeyedPreviewErrorModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeKeyedPreviewErrorModel performs a merge with any union data inside the SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties, using the provided KeyedPreviewErrorModel
func (t *SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties) MergeKeyedPreviewErrorModel(v KeyedPreviewErrorModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *SCContentPreviewCompleteRequest_KeyedPreviews_AdditionalProperties) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsContentTypeFieldDefinition returns the union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as a ContentTypeFieldDefinition
func (t SCContentTypeCreateRequest_FieldDefinitions_Item) AsContentTypeFieldDefinition() (ContentTypeFieldDefinition, error) {
	var body ContentTypeFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentTypeFieldDefinition overwrites any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as the provided ContentTypeFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) FromContentTypeFieldDefinition(v ContentTypeFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentTypeFieldDefinition performs a merge with any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item, using the provided ContentTypeFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) MergeContentTypeFieldDefinition(v ContentTypeFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsLibraryAssetFieldDefinition returns the union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as a LibraryAssetFieldDefinition
func (t SCContentTypeCreateRequest_FieldDefinitions_Item) AsLibraryAssetFieldDefinition() (LibraryAssetFieldDefinition, error) {
	var body LibraryAssetFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLibraryAssetFieldDefinition overwrites any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as the provided LibraryAssetFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) FromLibraryAssetFieldDefinition(v LibraryAssetFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLibraryAssetFieldDefinition performs a merge with any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item, using the provided LibraryAssetFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) MergeLibraryAssetFieldDefinition(v LibraryAssetFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTextFieldDefinition returns the union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as a TextFieldDefinition
func (t SCContentTypeCreateRequest_FieldDefinitions_Item) AsTextFieldDefinition() (TextFieldDefinition, error) {
	var body TextFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTextFieldDefinition overwrites any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as the provided TextFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) FromTextFieldDefinition(v TextFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTextFieldDefinition performs a merge with any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item, using the provided TextFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) MergeTextFieldDefinition(v TextFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBaseFieldDefinition returns the union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as a BaseFieldDefinition
func (t SCContentTypeCreateRequest_FieldDefinitions_Item) AsBaseFieldDefinition() (BaseFieldDefinition, error) {
	var body BaseFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBaseFieldDefinition overwrites any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as the provided BaseFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) FromBaseFieldDefinition(v BaseFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBaseFieldDefinition performs a merge with any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item, using the provided BaseFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) MergeBaseFieldDefinition(v BaseFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDatetimeFieldDefinition returns the union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as a DatetimeFieldDefinition
func (t SCContentTypeCreateRequest_FieldDefinitions_Item) AsDatetimeFieldDefinition() (DatetimeFieldDefinition, error) {
	var body DatetimeFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDatetimeFieldDefinition overwrites any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as the provided DatetimeFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) FromDatetimeFieldDefinition(v DatetimeFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDatetimeFieldDefinition performs a merge with any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item, using the provided DatetimeFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) MergeDatetimeFieldDefinition(v DatetimeFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRichTextFieldDefinition returns the union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as a RichTextFieldDefinition
func (t SCContentTypeCreateRequest_FieldDefinitions_Item) AsRichTextFieldDefinition() (RichTextFieldDefinition, error) {
	var body RichTextFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRichTextFieldDefinition overwrites any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as the provided RichTextFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) FromRichTextFieldDefinition(v RichTextFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRichTextFieldDefinition performs a merge with any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item, using the provided RichTextFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) MergeRichTextFieldDefinition(v RichTextFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsChoiceFieldDefinition returns the union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as a ChoiceFieldDefinition
func (t SCContentTypeCreateRequest_FieldDefinitions_Item) AsChoiceFieldDefinition() (ChoiceFieldDefinition, error) {
	var body ChoiceFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromChoiceFieldDefinition overwrites any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as the provided ChoiceFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) FromChoiceFieldDefinition(v ChoiceFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeChoiceFieldDefinition performs a merge with any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item, using the provided ChoiceFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) MergeChoiceFieldDefinition(v ChoiceFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsNumberFieldDefinition returns the union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as a NumberFieldDefinition
func (t SCContentTypeCreateRequest_FieldDefinitions_Item) AsNumberFieldDefinition() (NumberFieldDefinition, error) {
	var body NumberFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberFieldDefinition overwrites any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item as the provided NumberFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) FromNumberFieldDefinition(v NumberFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberFieldDefinition performs a merge with any union data inside the SCContentTypeCreateRequest_FieldDefinitions_Item, using the provided NumberFieldDefinition
func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) MergeNumberFieldDefinition(v NumberFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t SCContentTypeCreateRequest_FieldDefinitions_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *SCContentTypeCreateRequest_FieldDefinitions_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsContentTypeFieldDefinition returns the union data inside the SCContentTypeVersion_FieldDefinitions_Item as a ContentTypeFieldDefinition
func (t SCContentTypeVersion_FieldDefinitions_Item) AsContentTypeFieldDefinition() (ContentTypeFieldDefinition, error) {
	var body ContentTypeFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentTypeFieldDefinition overwrites any union data inside the SCContentTypeVersion_FieldDefinitions_Item as the provided ContentTypeFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) FromContentTypeFieldDefinition(v ContentTypeFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentTypeFieldDefinition performs a merge with any union data inside the SCContentTypeVersion_FieldDefinitions_Item, using the provided ContentTypeFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) MergeContentTypeFieldDefinition(v ContentTypeFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsLibraryAssetFieldDefinition returns the union data inside the SCContentTypeVersion_FieldDefinitions_Item as a LibraryAssetFieldDefinition
func (t SCContentTypeVersion_FieldDefinitions_Item) AsLibraryAssetFieldDefinition() (LibraryAssetFieldDefinition, error) {
	var body LibraryAssetFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLibraryAssetFieldDefinition overwrites any union data inside the SCContentTypeVersion_FieldDefinitions_Item as the provided LibraryAssetFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) FromLibraryAssetFieldDefinition(v LibraryAssetFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLibraryAssetFieldDefinition performs a merge with any union data inside the SCContentTypeVersion_FieldDefinitions_Item, using the provided LibraryAssetFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) MergeLibraryAssetFieldDefinition(v LibraryAssetFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTextFieldDefinition returns the union data inside the SCContentTypeVersion_FieldDefinitions_Item as a TextFieldDefinition
func (t SCContentTypeVersion_FieldDefinitions_Item) AsTextFieldDefinition() (TextFieldDefinition, error) {
	var body TextFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTextFieldDefinition overwrites any union data inside the SCContentTypeVersion_FieldDefinitions_Item as the provided TextFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) FromTextFieldDefinition(v TextFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTextFieldDefinition performs a merge with any union data inside the SCContentTypeVersion_FieldDefinitions_Item, using the provided TextFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) MergeTextFieldDefinition(v TextFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBaseFieldDefinition returns the union data inside the SCContentTypeVersion_FieldDefinitions_Item as a BaseFieldDefinition
func (t SCContentTypeVersion_FieldDefinitions_Item) AsBaseFieldDefinition() (BaseFieldDefinition, error) {
	var body BaseFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBaseFieldDefinition overwrites any union data inside the SCContentTypeVersion_FieldDefinitions_Item as the provided BaseFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) FromBaseFieldDefinition(v BaseFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBaseFieldDefinition performs a merge with any union data inside the SCContentTypeVersion_FieldDefinitions_Item, using the provided BaseFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) MergeBaseFieldDefinition(v BaseFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDatetimeFieldDefinition returns the union data inside the SCContentTypeVersion_FieldDefinitions_Item as a DatetimeFieldDefinition
func (t SCContentTypeVersion_FieldDefinitions_Item) AsDatetimeFieldDefinition() (DatetimeFieldDefinition, error) {
	var body DatetimeFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDatetimeFieldDefinition overwrites any union data inside the SCContentTypeVersion_FieldDefinitions_Item as the provided DatetimeFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) FromDatetimeFieldDefinition(v DatetimeFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDatetimeFieldDefinition performs a merge with any union data inside the SCContentTypeVersion_FieldDefinitions_Item, using the provided DatetimeFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) MergeDatetimeFieldDefinition(v DatetimeFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRichTextFieldDefinition returns the union data inside the SCContentTypeVersion_FieldDefinitions_Item as a RichTextFieldDefinition
func (t SCContentTypeVersion_FieldDefinitions_Item) AsRichTextFieldDefinition() (RichTextFieldDefinition, error) {
	var body RichTextFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRichTextFieldDefinition overwrites any union data inside the SCContentTypeVersion_FieldDefinitions_Item as the provided RichTextFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) FromRichTextFieldDefinition(v RichTextFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRichTextFieldDefinition performs a merge with any union data inside the SCContentTypeVersion_FieldDefinitions_Item, using the provided RichTextFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) MergeRichTextFieldDefinition(v RichTextFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsChoiceFieldDefinition returns the union data inside the SCContentTypeVersion_FieldDefinitions_Item as a ChoiceFieldDefinition
func (t SCContentTypeVersion_FieldDefinitions_Item) AsChoiceFieldDefinition() (ChoiceFieldDefinition, error) {
	var body ChoiceFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromChoiceFieldDefinition overwrites any union data inside the SCContentTypeVersion_FieldDefinitions_Item as the provided ChoiceFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) FromChoiceFieldDefinition(v ChoiceFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeChoiceFieldDefinition performs a merge with any union data inside the SCContentTypeVersion_FieldDefinitions_Item, using the provided ChoiceFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) MergeChoiceFieldDefinition(v ChoiceFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsNumberFieldDefinition returns the union data inside the SCContentTypeVersion_FieldDefinitions_Item as a NumberFieldDefinition
func (t SCContentTypeVersion_FieldDefinitions_Item) AsNumberFieldDefinition() (NumberFieldDefinition, error) {
	var body NumberFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberFieldDefinition overwrites any union data inside the SCContentTypeVersion_FieldDefinitions_Item as the provided NumberFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) FromNumberFieldDefinition(v NumberFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberFieldDefinition performs a merge with any union data inside the SCContentTypeVersion_FieldDefinitions_Item, using the provided NumberFieldDefinition
func (t *SCContentTypeVersion_FieldDefinitions_Item) MergeNumberFieldDefinition(v NumberFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t SCContentTypeVersion_FieldDefinitions_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *SCContentTypeVersion_FieldDefinitions_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsContentTypeFieldDefinition returns the union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as a ContentTypeFieldDefinition
func (t SCContentTypeVersionCreateRequest_FieldDefinitions_Item) AsContentTypeFieldDefinition() (ContentTypeFieldDefinition, error) {
	var body ContentTypeFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentTypeFieldDefinition overwrites any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as the provided ContentTypeFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) FromContentTypeFieldDefinition(v ContentTypeFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentTypeFieldDefinition performs a merge with any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item, using the provided ContentTypeFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) MergeContentTypeFieldDefinition(v ContentTypeFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsLibraryAssetFieldDefinition returns the union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as a LibraryAssetFieldDefinition
func (t SCContentTypeVersionCreateRequest_FieldDefinitions_Item) AsLibraryAssetFieldDefinition() (LibraryAssetFieldDefinition, error) {
	var body LibraryAssetFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLibraryAssetFieldDefinition overwrites any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as the provided LibraryAssetFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) FromLibraryAssetFieldDefinition(v LibraryAssetFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLibraryAssetFieldDefinition performs a merge with any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item, using the provided LibraryAssetFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) MergeLibraryAssetFieldDefinition(v LibraryAssetFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTextFieldDefinition returns the union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as a TextFieldDefinition
func (t SCContentTypeVersionCreateRequest_FieldDefinitions_Item) AsTextFieldDefinition() (TextFieldDefinition, error) {
	var body TextFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTextFieldDefinition overwrites any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as the provided TextFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) FromTextFieldDefinition(v TextFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTextFieldDefinition performs a merge with any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item, using the provided TextFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) MergeTextFieldDefinition(v TextFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBaseFieldDefinition returns the union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as a BaseFieldDefinition
func (t SCContentTypeVersionCreateRequest_FieldDefinitions_Item) AsBaseFieldDefinition() (BaseFieldDefinition, error) {
	var body BaseFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBaseFieldDefinition overwrites any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as the provided BaseFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) FromBaseFieldDefinition(v BaseFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBaseFieldDefinition performs a merge with any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item, using the provided BaseFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) MergeBaseFieldDefinition(v BaseFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDatetimeFieldDefinition returns the union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as a DatetimeFieldDefinition
func (t SCContentTypeVersionCreateRequest_FieldDefinitions_Item) AsDatetimeFieldDefinition() (DatetimeFieldDefinition, error) {
	var body DatetimeFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDatetimeFieldDefinition overwrites any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as the provided DatetimeFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) FromDatetimeFieldDefinition(v DatetimeFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDatetimeFieldDefinition performs a merge with any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item, using the provided DatetimeFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) MergeDatetimeFieldDefinition(v DatetimeFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRichTextFieldDefinition returns the union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as a RichTextFieldDefinition
func (t SCContentTypeVersionCreateRequest_FieldDefinitions_Item) AsRichTextFieldDefinition() (RichTextFieldDefinition, error) {
	var body RichTextFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRichTextFieldDefinition overwrites any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as the provided RichTextFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) FromRichTextFieldDefinition(v RichTextFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRichTextFieldDefinition performs a merge with any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item, using the provided RichTextFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) MergeRichTextFieldDefinition(v RichTextFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsChoiceFieldDefinition returns the union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as a ChoiceFieldDefinition
func (t SCContentTypeVersionCreateRequest_FieldDefinitions_Item) AsChoiceFieldDefinition() (ChoiceFieldDefinition, error) {
	var body ChoiceFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromChoiceFieldDefinition overwrites any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as the provided ChoiceFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) FromChoiceFieldDefinition(v ChoiceFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeChoiceFieldDefinition performs a merge with any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item, using the provided ChoiceFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) MergeChoiceFieldDefinition(v ChoiceFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsNumberFieldDefinition returns the union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as a NumberFieldDefinition
func (t SCContentTypeVersionCreateRequest_FieldDefinitions_Item) AsNumberFieldDefinition() (NumberFieldDefinition, error) {
	var body NumberFieldDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberFieldDefinition overwrites any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item as the provided NumberFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) FromNumberFieldDefinition(v NumberFieldDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberFieldDefinition performs a merge with any union data inside the SCContentTypeVersionCreateRequest_FieldDefinitions_Item, using the provided NumberFieldDefinition
func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) MergeNumberFieldDefinition(v NumberFieldDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t SCContentTypeVersionCreateRequest_FieldDefinitions_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *SCContentTypeVersionCreateRequest_FieldDefinitions_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGenericSettingsFieldChoiceCreatePayload returns the union data inside the SettingsFieldChoiceCreateRequest_Item as a GenericSettingsFieldChoiceCreatePayload
func (t SettingsFieldChoiceCreateRequest_Item) AsGenericSettingsFieldChoiceCreatePayload() (GenericSettingsFieldChoiceCreatePayload, error) {
	var body GenericSettingsFieldChoiceCreatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGenericSettingsFieldChoiceCreatePayload overwrites any union data inside the SettingsFieldChoiceCreateRequest_Item as the provided GenericSettingsFieldChoiceCreatePayload
func (t *SettingsFieldChoiceCreateRequest_Item) FromGenericSettingsFieldChoiceCreatePayload(v GenericSettingsFieldChoiceCreatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGenericSettingsFieldChoiceCreatePayload performs a merge with any union data inside the SettingsFieldChoiceCreateRequest_Item, using the provided GenericSettingsFieldChoiceCreatePayload
func (t *SettingsFieldChoiceCreateRequest_Item) MergeGenericSettingsFieldChoiceCreatePayload(v GenericSettingsFieldChoiceCreatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsLabelTypeSettingsFieldChoiceCreatePayload returns the union data inside the SettingsFieldChoiceCreateRequest_Item as a LabelTypeSettingsFieldChoiceCreatePayload
func (t SettingsFieldChoiceCreateRequest_Item) AsLabelTypeSettingsFieldChoiceCreatePayload() (LabelTypeSettingsFieldChoiceCreatePayload, error) {
	var body LabelTypeSettingsFieldChoiceCreatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLabelTypeSettingsFieldChoiceCreatePayload overwrites any union data inside the SettingsFieldChoiceCreateRequest_Item as the provided LabelTypeSettingsFieldChoiceCreatePayload
func (t *SettingsFieldChoiceCreateRequest_Item) FromLabelTypeSettingsFieldChoiceCreatePayload(v LabelTypeSettingsFieldChoiceCreatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLabelTypeSettingsFieldChoiceCreatePayload performs a merge with any union data inside the SettingsFieldChoiceCreateRequest_Item, using the provided LabelTypeSettingsFieldChoiceCreatePayload
func (t *SettingsFieldChoiceCreateRequest_Item) MergeLabelTypeSettingsFieldChoiceCreatePayload(v LabelTypeSettingsFieldChoiceCreatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t SettingsFieldChoiceCreateRequest_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *SettingsFieldChoiceCreateRequest_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGenericSettingsFieldChoiceUpdatePayload returns the union data inside the SettingsFieldChoiceUpdateRequest as a GenericSettingsFieldChoiceUpdatePayload
func (t SettingsFieldChoiceUpdateRequest) AsGenericSettingsFieldChoiceUpdatePayload() (GenericSettingsFieldChoiceUpdatePayload, error) {
	var body GenericSettingsFieldChoiceUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGenericSettingsFieldChoiceUpdatePayload overwrites any union data inside the SettingsFieldChoiceUpdateRequest as the provided GenericSettingsFieldChoiceUpdatePayload
func (t *SettingsFieldChoiceUpdateRequest) FromGenericSettingsFieldChoiceUpdatePayload(v GenericSettingsFieldChoiceUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGenericSettingsFieldChoiceUpdatePayload performs a merge with any union data inside the SettingsFieldChoiceUpdateRequest, using the provided GenericSettingsFieldChoiceUpdatePayload
func (t *SettingsFieldChoiceUpdateRequest) MergeGenericSettingsFieldChoiceUpdatePayload(v GenericSettingsFieldChoiceUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsLabelTypeSettingsFieldChoiceUpdatePayload returns the union data inside the SettingsFieldChoiceUpdateRequest as a LabelTypeSettingsFieldChoiceUpdatePayload
func (t SettingsFieldChoiceUpdateRequest) AsLabelTypeSettingsFieldChoiceUpdatePayload() (LabelTypeSettingsFieldChoiceUpdatePayload, error) {
	var body LabelTypeSettingsFieldChoiceUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLabelTypeSettingsFieldChoiceUpdatePayload overwrites any union data inside the SettingsFieldChoiceUpdateRequest as the provided LabelTypeSettingsFieldChoiceUpdatePayload
func (t *SettingsFieldChoiceUpdateRequest) FromLabelTypeSettingsFieldChoiceUpdatePayload(v LabelTypeSettingsFieldChoiceUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLabelTypeSettingsFieldChoiceUpdatePayload performs a merge with any union data inside the SettingsFieldChoiceUpdateRequest, using the provided LabelTypeSettingsFieldChoiceUpdatePayload
func (t *SettingsFieldChoiceUpdateRequest) MergeLabelTypeSettingsFieldChoiceUpdatePayload(v LabelTypeSettingsFieldChoiceUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t SettingsFieldChoiceUpdateRequest) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *SettingsFieldChoiceUpdateRequest) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsBaseSettingsFieldCreatePayload returns the union data inside the SettingsFieldCreateRequest as a BaseSettingsFieldCreatePayload
func (t SettingsFieldCreateRequest) AsBaseSettingsFieldCreatePayload() (BaseSettingsFieldCreatePayload, error) {
	var body BaseSettingsFieldCreatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBaseSettingsFieldCreatePayload overwrites any union data inside the SettingsFieldCreateRequest as the provided BaseSettingsFieldCreatePayload
func (t *SettingsFieldCreateRequest) FromBaseSettingsFieldCreatePayload(v BaseSettingsFieldCreatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBaseSettingsFieldCreatePayload performs a merge with any union data inside the SettingsFieldCreateRequest, using the provided BaseSettingsFieldCreatePayload
func (t *SettingsFieldCreateRequest) MergeBaseSettingsFieldCreatePayload(v BaseSettingsFieldCreatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsLabelTypeSettingsFieldCreatePayload returns the union data inside the SettingsFieldCreateRequest as a LabelTypeSettingsFieldCreatePayload
func (t SettingsFieldCreateRequest) AsLabelTypeSettingsFieldCreatePayload() (LabelTypeSettingsFieldCreatePayload, error) {
	var body LabelTypeSettingsFieldCreatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLabelTypeSettingsFieldCreatePayload overwrites any union data inside the SettingsFieldCreateRequest as the provided LabelTypeSettingsFieldCreatePayload
func (t *SettingsFieldCreateRequest) FromLabelTypeSettingsFieldCreatePayload(v LabelTypeSettingsFieldCreatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLabelTypeSettingsFieldCreatePayload performs a merge with any union data inside the SettingsFieldCreateRequest, using the provided LabelTypeSettingsFieldCreatePayload
func (t *SettingsFieldCreateRequest) MergeLabelTypeSettingsFieldCreatePayload(v LabelTypeSettingsFieldCreatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDropdownTypeSettingsFieldCreatePayload returns the union data inside the SettingsFieldCreateRequest as a DropdownTypeSettingsFieldCreatePayload
func (t SettingsFieldCreateRequest) AsDropdownTypeSettingsFieldCreatePayload() (DropdownTypeSettingsFieldCreatePayload, error) {
	var body DropdownTypeSettingsFieldCreatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDropdownTypeSettingsFieldCreatePayload overwrites any union data inside the SettingsFieldCreateRequest as the provided DropdownTypeSettingsFieldCreatePayload
func (t *SettingsFieldCreateRequest) FromDropdownTypeSettingsFieldCreatePayload(v DropdownTypeSettingsFieldCreatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDropdownTypeSettingsFieldCreatePayload performs a merge with any union data inside the SettingsFieldCreateRequest, using the provided DropdownTypeSettingsFieldCreatePayload
func (t *SettingsFieldCreateRequest) MergeDropdownTypeSettingsFieldCreatePayload(v DropdownTypeSettingsFieldCreatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCheckboxAndRadioTypeSettingsFieldCreatePayload returns the union data inside the SettingsFieldCreateRequest as a CheckboxAndRadioTypeSettingsFieldCreatePayload
func (t SettingsFieldCreateRequest) AsCheckboxAndRadioTypeSettingsFieldCreatePayload() (CheckboxAndRadioTypeSettingsFieldCreatePayload, error) {
	var body CheckboxAndRadioTypeSettingsFieldCreatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCheckboxAndRadioTypeSettingsFieldCreatePayload overwrites any union data inside the SettingsFieldCreateRequest as the provided CheckboxAndRadioTypeSettingsFieldCreatePayload
func (t *SettingsFieldCreateRequest) FromCheckboxAndRadioTypeSettingsFieldCreatePayload(v CheckboxAndRadioTypeSettingsFieldCreatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCheckboxAndRadioTypeSettingsFieldCreatePayload performs a merge with any union data inside the SettingsFieldCreateRequest, using the provided CheckboxAndRadioTypeSettingsFieldCreatePayload
func (t *SettingsFieldCreateRequest) MergeCheckboxAndRadioTypeSettingsFieldCreatePayload(v CheckboxAndRadioTypeSettingsFieldCreatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsGenericNumberTypeSettingsFieldCreatePayload returns the union data inside the SettingsFieldCreateRequest as a GenericNumberTypeSettingsFieldCreatePayload
func (t SettingsFieldCreateRequest) AsGenericNumberTypeSettingsFieldCreatePayload() (GenericNumberTypeSettingsFieldCreatePayload, error) {
	var body GenericNumberTypeSettingsFieldCreatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGenericNumberTypeSettingsFieldCreatePayload overwrites any union data inside the SettingsFieldCreateRequest as the provided GenericNumberTypeSettingsFieldCreatePayload
func (t *SettingsFieldCreateRequest) FromGenericNumberTypeSettingsFieldCreatePayload(v GenericNumberTypeSettingsFieldCreatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGenericNumberTypeSettingsFieldCreatePayload performs a merge with any union data inside the SettingsFieldCreateRequest, using the provided GenericNumberTypeSettingsFieldCreatePayload
func (t *SettingsFieldCreateRequest) MergeGenericNumberTypeSettingsFieldCreatePayload(v GenericNumberTypeSettingsFieldCreatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCurrencyNumberTypeSettingsFieldCreatePayload returns the union data inside the SettingsFieldCreateRequest as a CurrencyNumberTypeSettingsFieldCreatePayload
func (t SettingsFieldCreateRequest) AsCurrencyNumberTypeSettingsFieldCreatePayload() (CurrencyNumberTypeSettingsFieldCreatePayload, error) {
	var body CurrencyNumberTypeSettingsFieldCreatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCurrencyNumberTypeSettingsFieldCreatePayload overwrites any union data inside the SettingsFieldCreateRequest as the provided CurrencyNumberTypeSettingsFieldCreatePayload
func (t *SettingsFieldCreateRequest) FromCurrencyNumberTypeSettingsFieldCreatePayload(v CurrencyNumberTypeSettingsFieldCreatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCurrencyNumberTypeSettingsFieldCreatePayload performs a merge with any union data inside the SettingsFieldCreateRequest, using the provided CurrencyNumberTypeSettingsFieldCreatePayload
func (t *SettingsFieldCreateRequest) MergeCurrencyNumberTypeSettingsFieldCreatePayload(v CurrencyNumberTypeSettingsFieldCreatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t SettingsFieldCreateRequest) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *SettingsFieldCreateRequest) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsBaseSettingsFieldUpdatePayload returns the union data inside the SettingsFieldUpdateRequest as a BaseSettingsFieldUpdatePayload
func (t SettingsFieldUpdateRequest) AsBaseSettingsFieldUpdatePayload() (BaseSettingsFieldUpdatePayload, error) {
	var body BaseSettingsFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBaseSettingsFieldUpdatePayload overwrites any union data inside the SettingsFieldUpdateRequest as the provided BaseSettingsFieldUpdatePayload
func (t *SettingsFieldUpdateRequest) FromBaseSettingsFieldUpdatePayload(v BaseSettingsFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBaseSettingsFieldUpdatePayload performs a merge with any union data inside the SettingsFieldUpdateRequest, using the provided BaseSettingsFieldUpdatePayload
func (t *SettingsFieldUpdateRequest) MergeBaseSettingsFieldUpdatePayload(v BaseSettingsFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsLabelAndDropdownTypeSettingsFieldUpdatePayload returns the union data inside the SettingsFieldUpdateRequest as a LabelAndDropdownTypeSettingsFieldUpdatePayload
func (t SettingsFieldUpdateRequest) AsLabelAndDropdownTypeSettingsFieldUpdatePayload() (LabelAndDropdownTypeSettingsFieldUpdatePayload, error) {
	var body LabelAndDropdownTypeSettingsFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLabelAndDropdownTypeSettingsFieldUpdatePayload overwrites any union data inside the SettingsFieldUpdateRequest as the provided LabelAndDropdownTypeSettingsFieldUpdatePayload
func (t *SettingsFieldUpdateRequest) FromLabelAndDropdownTypeSettingsFieldUpdatePayload(v LabelAndDropdownTypeSettingsFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLabelAndDropdownTypeSettingsFieldUpdatePayload performs a merge with any union data inside the SettingsFieldUpdateRequest, using the provided LabelAndDropdownTypeSettingsFieldUpdatePayload
func (t *SettingsFieldUpdateRequest) MergeLabelAndDropdownTypeSettingsFieldUpdatePayload(v LabelAndDropdownTypeSettingsFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsGenericNumberTypeSettingsFieldUpdatePayload returns the union data inside the SettingsFieldUpdateRequest as a GenericNumberTypeSettingsFieldUpdatePayload
func (t SettingsFieldUpdateRequest) AsGenericNumberTypeSettingsFieldUpdatePayload() (GenericNumberTypeSettingsFieldUpdatePayload, error) {
	var body GenericNumberTypeSettingsFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGenericNumberTypeSettingsFieldUpdatePayload overwrites any union data inside the SettingsFieldUpdateRequest as the provided GenericNumberTypeSettingsFieldUpdatePayload
func (t *SettingsFieldUpdateRequest) FromGenericNumberTypeSettingsFieldUpdatePayload(v GenericNumberTypeSettingsFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGenericNumberTypeSettingsFieldUpdatePayload performs a merge with any union data inside the SettingsFieldUpdateRequest, using the provided GenericNumberTypeSettingsFieldUpdatePayload
func (t *SettingsFieldUpdateRequest) MergeGenericNumberTypeSettingsFieldUpdatePayload(v GenericNumberTypeSettingsFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCurrencyNumberTypeSettingsFieldUpdatePayload returns the union data inside the SettingsFieldUpdateRequest as a CurrencyNumberTypeSettingsFieldUpdatePayload
func (t SettingsFieldUpdateRequest) AsCurrencyNumberTypeSettingsFieldUpdatePayload() (CurrencyNumberTypeSettingsFieldUpdatePayload, error) {
	var body CurrencyNumberTypeSettingsFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCurrencyNumberTypeSettingsFieldUpdatePayload overwrites any union data inside the SettingsFieldUpdateRequest as the provided CurrencyNumberTypeSettingsFieldUpdatePayload
func (t *SettingsFieldUpdateRequest) FromCurrencyNumberTypeSettingsFieldUpdatePayload(v CurrencyNumberTypeSettingsFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCurrencyNumberTypeSettingsFieldUpdatePayload performs a merge with any union data inside the SettingsFieldUpdateRequest, using the provided CurrencyNumberTypeSettingsFieldUpdatePayload
func (t *SettingsFieldUpdateRequest) MergeCurrencyNumberTypeSettingsFieldUpdatePayload(v CurrencyNumberTypeSettingsFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t SettingsFieldUpdateRequest) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *SettingsFieldUpdateRequest) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsPatchLocalizedFieldValuesWithEmbeddedContent returns the union data inside the StructuredContentFields_Item as a PatchLocalizedFieldValuesWithEmbeddedContent
func (t StructuredContentFields_Item) AsPatchLocalizedFieldValuesWithEmbeddedContent() (PatchLocalizedFieldValuesWithEmbeddedContent, error) {
	var body PatchLocalizedFieldValuesWithEmbeddedContent
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPatchLocalizedFieldValuesWithEmbeddedContent overwrites any union data inside the StructuredContentFields_Item as the provided PatchLocalizedFieldValuesWithEmbeddedContent
func (t *StructuredContentFields_Item) FromPatchLocalizedFieldValuesWithEmbeddedContent(v PatchLocalizedFieldValuesWithEmbeddedContent) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePatchLocalizedFieldValuesWithEmbeddedContent performs a merge with any union data inside the StructuredContentFields_Item, using the provided PatchLocalizedFieldValuesWithEmbeddedContent
func (t *StructuredContentFields_Item) MergePatchLocalizedFieldValuesWithEmbeddedContent(v PatchLocalizedFieldValuesWithEmbeddedContent) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDeleteLocaleFieldValueModel returns the union data inside the StructuredContentFields_Item as a DeleteLocaleFieldValueModel
func (t StructuredContentFields_Item) AsDeleteLocaleFieldValueModel() (DeleteLocaleFieldValueModel, error) {
	var body DeleteLocaleFieldValueModel
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDeleteLocaleFieldValueModel overwrites any union data inside the StructuredContentFields_Item as the provided DeleteLocaleFieldValueModel
func (t *StructuredContentFields_Item) FromDeleteLocaleFieldValueModel(v DeleteLocaleFieldValueModel) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDeleteLocaleFieldValueModel performs a merge with any union data inside the StructuredContentFields_Item, using the provided DeleteLocaleFieldValueModel
func (t *StructuredContentFields_Item) MergeDeleteLocaleFieldValueModel(v DeleteLocaleFieldValueModel) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t StructuredContentFields_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *StructuredContentFields_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsStringTypeObjectField returns the union data inside the TaskAssetFieldsUpdateRequest_Item as a StringTypeObjectField
func (t TaskAssetFieldsUpdateRequest_Item) AsStringTypeObjectField() (StringTypeObjectField, error) {
	var body StringTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromStringTypeObjectField overwrites any union data inside the TaskAssetFieldsUpdateRequest_Item as the provided StringTypeObjectField
func (t *TaskAssetFieldsUpdateRequest_Item) FromStringTypeObjectField(v StringTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeStringTypeObjectField performs a merge with any union data inside the TaskAssetFieldsUpdateRequest_Item, using the provided StringTypeObjectField
func (t *TaskAssetFieldsUpdateRequest_Item) MergeStringTypeObjectField(v StringTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMultiChoiceTypeObjectField returns the union data inside the TaskAssetFieldsUpdateRequest_Item as a MultiChoiceTypeObjectField
func (t TaskAssetFieldsUpdateRequest_Item) AsMultiChoiceTypeObjectField() (MultiChoiceTypeObjectField, error) {
	var body MultiChoiceTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMultiChoiceTypeObjectField overwrites any union data inside the TaskAssetFieldsUpdateRequest_Item as the provided MultiChoiceTypeObjectField
func (t *TaskAssetFieldsUpdateRequest_Item) FromMultiChoiceTypeObjectField(v MultiChoiceTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMultiChoiceTypeObjectField performs a merge with any union data inside the TaskAssetFieldsUpdateRequest_Item, using the provided MultiChoiceTypeObjectField
func (t *TaskAssetFieldsUpdateRequest_Item) MergeMultiChoiceTypeObjectField(v MultiChoiceTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDropdownTypeObjectField returns the union data inside the TaskAssetFieldsUpdateRequest_Item as a DropdownTypeObjectField
func (t TaskAssetFieldsUpdateRequest_Item) AsDropdownTypeObjectField() (DropdownTypeObjectField, error) {
	var body DropdownTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDropdownTypeObjectField overwrites any union data inside the TaskAssetFieldsUpdateRequest_Item as the provided DropdownTypeObjectField
func (t *TaskAssetFieldsUpdateRequest_Item) FromDropdownTypeObjectField(v DropdownTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDropdownTypeObjectField performs a merge with any union data inside the TaskAssetFieldsUpdateRequest_Item, using the provided DropdownTypeObjectField
func (t *TaskAssetFieldsUpdateRequest_Item) MergeDropdownTypeObjectField(v DropdownTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRadioButtonTypeObjectField returns the union data inside the TaskAssetFieldsUpdateRequest_Item as a RadioButtonTypeObjectField
func (t TaskAssetFieldsUpdateRequest_Item) AsRadioButtonTypeObjectField() (RadioButtonTypeObjectField, error) {
	var body RadioButtonTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRadioButtonTypeObjectField overwrites any union data inside the TaskAssetFieldsUpdateRequest_Item as the provided RadioButtonTypeObjectField
func (t *TaskAssetFieldsUpdateRequest_Item) FromRadioButtonTypeObjectField(v RadioButtonTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRadioButtonTypeObjectField performs a merge with any union data inside the TaskAssetFieldsUpdateRequest_Item, using the provided RadioButtonTypeObjectField
func (t *TaskAssetFieldsUpdateRequest_Item) MergeRadioButtonTypeObjectField(v RadioButtonTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsNumberTypeObjectField returns the union data inside the TaskAssetFieldsUpdateRequest_Item as a NumberTypeObjectField
func (t TaskAssetFieldsUpdateRequest_Item) AsNumberTypeObjectField() (NumberTypeObjectField, error) {
	var body NumberTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberTypeObjectField overwrites any union data inside the TaskAssetFieldsUpdateRequest_Item as the provided NumberTypeObjectField
func (t *TaskAssetFieldsUpdateRequest_Item) FromNumberTypeObjectField(v NumberTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberTypeObjectField performs a merge with any union data inside the TaskAssetFieldsUpdateRequest_Item, using the provided NumberTypeObjectField
func (t *TaskAssetFieldsUpdateRequest_Item) MergeNumberTypeObjectField(v NumberTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDateTypeObjectField returns the union data inside the TaskAssetFieldsUpdateRequest_Item as a DateTypeObjectField
func (t TaskAssetFieldsUpdateRequest_Item) AsDateTypeObjectField() (DateTypeObjectField, error) {
	var body DateTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDateTypeObjectField overwrites any union data inside the TaskAssetFieldsUpdateRequest_Item as the provided DateTypeObjectField
func (t *TaskAssetFieldsUpdateRequest_Item) FromDateTypeObjectField(v DateTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDateTypeObjectField performs a merge with any union data inside the TaskAssetFieldsUpdateRequest_Item, using the provided DateTypeObjectField
func (t *TaskAssetFieldsUpdateRequest_Item) MergeDateTypeObjectField(v DateTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAssetTypeObjectField returns the union data inside the TaskAssetFieldsUpdateRequest_Item as a AssetTypeObjectField
func (t TaskAssetFieldsUpdateRequest_Item) AsAssetTypeObjectField() (AssetTypeObjectField, error) {
	var body AssetTypeObjectField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetTypeObjectField overwrites any union data inside the TaskAssetFieldsUpdateRequest_Item as the provided AssetTypeObjectField
func (t *TaskAssetFieldsUpdateRequest_Item) FromAssetTypeObjectField(v AssetTypeObjectField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetTypeObjectField performs a merge with any union data inside the TaskAssetFieldsUpdateRequest_Item, using the provided AssetTypeObjectField
func (t *TaskAssetFieldsUpdateRequest_Item) MergeAssetTypeObjectField(v AssetTypeObjectField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t TaskAssetFieldsUpdateRequest_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *TaskAssetFieldsUpdateRequest_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsTaskAssetRequestForDirectUpload returns the union data inside the TaskAssetRequest as a TaskAssetRequestForDirectUpload
func (t TaskAssetRequest) AsTaskAssetRequestForDirectUpload() (TaskAssetRequestForDirectUpload, error) {
	var body TaskAssetRequestForDirectUpload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTaskAssetRequestForDirectUpload overwrites any union data inside the TaskAssetRequest as the provided TaskAssetRequestForDirectUpload
func (t *TaskAssetRequest) FromTaskAssetRequestForDirectUpload(v TaskAssetRequestForDirectUpload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTaskAssetRequestForDirectUpload performs a merge with any union data inside the TaskAssetRequest, using the provided TaskAssetRequestForDirectUpload
func (t *TaskAssetRequest) MergeTaskAssetRequestForDirectUpload(v TaskAssetRequestForDirectUpload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTaskAssetRequestForLibrary returns the union data inside the TaskAssetRequest as a TaskAssetRequestForLibrary
func (t TaskAssetRequest) AsTaskAssetRequestForLibrary() (TaskAssetRequestForLibrary, error) {
	var body TaskAssetRequestForLibrary
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTaskAssetRequestForLibrary overwrites any union data inside the TaskAssetRequest as the provided TaskAssetRequestForLibrary
func (t *TaskAssetRequest) FromTaskAssetRequestForLibrary(v TaskAssetRequestForLibrary) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTaskAssetRequestForLibrary performs a merge with any union data inside the TaskAssetRequest, using the provided TaskAssetRequestForLibrary
func (t *TaskAssetRequest) MergeTaskAssetRequestForLibrary(v TaskAssetRequestForLibrary) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t TaskAssetRequest) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *TaskAssetRequest) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsStringTypeObjectFieldUpdatePayload returns the union data inside the TaskFieldUpdateRequest as a StringTypeObjectFieldUpdatePayload
func (t TaskFieldUpdateRequest) AsStringTypeObjectFieldUpdatePayload() (StringTypeObjectFieldUpdatePayload, error) {
	var body StringTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromStringTypeObjectFieldUpdatePayload overwrites any union data inside the TaskFieldUpdateRequest as the provided StringTypeObjectFieldUpdatePayload
func (t *TaskFieldUpdateRequest) FromStringTypeObjectFieldUpdatePayload(v StringTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeStringTypeObjectFieldUpdatePayload performs a merge with any union data inside the TaskFieldUpdateRequest, using the provided StringTypeObjectFieldUpdatePayload
func (t *TaskFieldUpdateRequest) MergeStringTypeObjectFieldUpdatePayload(v StringTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMultiChoiceTypeObjectFieldUpdatePayload returns the union data inside the TaskFieldUpdateRequest as a MultiChoiceTypeObjectFieldUpdatePayload
func (t TaskFieldUpdateRequest) AsMultiChoiceTypeObjectFieldUpdatePayload() (MultiChoiceTypeObjectFieldUpdatePayload, error) {
	var body MultiChoiceTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMultiChoiceTypeObjectFieldUpdatePayload overwrites any union data inside the TaskFieldUpdateRequest as the provided MultiChoiceTypeObjectFieldUpdatePayload
func (t *TaskFieldUpdateRequest) FromMultiChoiceTypeObjectFieldUpdatePayload(v MultiChoiceTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMultiChoiceTypeObjectFieldUpdatePayload performs a merge with any union data inside the TaskFieldUpdateRequest, using the provided MultiChoiceTypeObjectFieldUpdatePayload
func (t *TaskFieldUpdateRequest) MergeMultiChoiceTypeObjectFieldUpdatePayload(v MultiChoiceTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDropdownTypeObjectFieldUpdatePayload returns the union data inside the TaskFieldUpdateRequest as a DropdownTypeObjectFieldUpdatePayload
func (t TaskFieldUpdateRequest) AsDropdownTypeObjectFieldUpdatePayload() (DropdownTypeObjectFieldUpdatePayload, error) {
	var body DropdownTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDropdownTypeObjectFieldUpdatePayload overwrites any union data inside the TaskFieldUpdateRequest as the provided DropdownTypeObjectFieldUpdatePayload
func (t *TaskFieldUpdateRequest) FromDropdownTypeObjectFieldUpdatePayload(v DropdownTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDropdownTypeObjectFieldUpdatePayload performs a merge with any union data inside the TaskFieldUpdateRequest, using the provided DropdownTypeObjectFieldUpdatePayload
func (t *TaskFieldUpdateRequest) MergeDropdownTypeObjectFieldUpdatePayload(v DropdownTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRadioButtonTypeObjectFieldUpdatePayload returns the union data inside the TaskFieldUpdateRequest as a RadioButtonTypeObjectFieldUpdatePayload
func (t TaskFieldUpdateRequest) AsRadioButtonTypeObjectFieldUpdatePayload() (RadioButtonTypeObjectFieldUpdatePayload, error) {
	var body RadioButtonTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRadioButtonTypeObjectFieldUpdatePayload overwrites any union data inside the TaskFieldUpdateRequest as the provided RadioButtonTypeObjectFieldUpdatePayload
func (t *TaskFieldUpdateRequest) FromRadioButtonTypeObjectFieldUpdatePayload(v RadioButtonTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRadioButtonTypeObjectFieldUpdatePayload performs a merge with any union data inside the TaskFieldUpdateRequest, using the provided RadioButtonTypeObjectFieldUpdatePayload
func (t *TaskFieldUpdateRequest) MergeRadioButtonTypeObjectFieldUpdatePayload(v RadioButtonTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsNumberTypeObjectFieldUpdatePayload returns the union data inside the TaskFieldUpdateRequest as a NumberTypeObjectFieldUpdatePayload
func (t TaskFieldUpdateRequest) AsNumberTypeObjectFieldUpdatePayload() (NumberTypeObjectFieldUpdatePayload, error) {
	var body NumberTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberTypeObjectFieldUpdatePayload overwrites any union data inside the TaskFieldUpdateRequest as the provided NumberTypeObjectFieldUpdatePayload
func (t *TaskFieldUpdateRequest) FromNumberTypeObjectFieldUpdatePayload(v NumberTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberTypeObjectFieldUpdatePayload performs a merge with any union data inside the TaskFieldUpdateRequest, using the provided NumberTypeObjectFieldUpdatePayload
func (t *TaskFieldUpdateRequest) MergeNumberTypeObjectFieldUpdatePayload(v NumberTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDateTypeObjectFieldUpdatePayload returns the union data inside the TaskFieldUpdateRequest as a DateTypeObjectFieldUpdatePayload
func (t TaskFieldUpdateRequest) AsDateTypeObjectFieldUpdatePayload() (DateTypeObjectFieldUpdatePayload, error) {
	var body DateTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDateTypeObjectFieldUpdatePayload overwrites any union data inside the TaskFieldUpdateRequest as the provided DateTypeObjectFieldUpdatePayload
func (t *TaskFieldUpdateRequest) FromDateTypeObjectFieldUpdatePayload(v DateTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDateTypeObjectFieldUpdatePayload performs a merge with any union data inside the TaskFieldUpdateRequest, using the provided DateTypeObjectFieldUpdatePayload
func (t *TaskFieldUpdateRequest) MergeDateTypeObjectFieldUpdatePayload(v DateTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsAssetTypeObjectFieldUpdatePayload returns the union data inside the TaskFieldUpdateRequest as a AssetTypeObjectFieldUpdatePayload
func (t TaskFieldUpdateRequest) AsAssetTypeObjectFieldUpdatePayload() (AssetTypeObjectFieldUpdatePayload, error) {
	var body AssetTypeObjectFieldUpdatePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAssetTypeObjectFieldUpdatePayload overwrites any union data inside the TaskFieldUpdateRequest as the provided AssetTypeObjectFieldUpdatePayload
func (t *TaskFieldUpdateRequest) FromAssetTypeObjectFieldUpdatePayload(v AssetTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAssetTypeObjectFieldUpdatePayload performs a merge with any union data inside the TaskFieldUpdateRequest, using the provided AssetTypeObjectFieldUpdatePayload
func (t *TaskFieldUpdateRequest) MergeAssetTypeObjectFieldUpdatePayload(v AssetTypeObjectFieldUpdatePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t TaskFieldUpdateRequest) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *TaskFieldUpdateRequest) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsTemplateDefaultFormField returns the union data inside the TemplateResponse_FormFields_Item as a TemplateDefaultFormField
func (t TemplateResponse_FormFields_Item) AsTemplateDefaultFormField() (TemplateDefaultFormField, error) {
	var body TemplateDefaultFormField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTemplateDefaultFormField overwrites any union data inside the TemplateResponse_FormFields_Item as the provided TemplateDefaultFormField
func (t *TemplateResponse_FormFields_Item) FromTemplateDefaultFormField(v TemplateDefaultFormField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTemplateDefaultFormField performs a merge with any union data inside the TemplateResponse_FormFields_Item, using the provided TemplateDefaultFormField
func (t *TemplateResponse_FormFields_Item) MergeTemplateDefaultFormField(v TemplateDefaultFormField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTemplateCurrencyNumberFormField returns the union data inside the TemplateResponse_FormFields_Item as a TemplateCurrencyNumberFormField
func (t TemplateResponse_FormFields_Item) AsTemplateCurrencyNumberFormField() (TemplateCurrencyNumberFormField, error) {
	var body TemplateCurrencyNumberFormField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTemplateCurrencyNumberFormField overwrites any union data inside the TemplateResponse_FormFields_Item as the provided TemplateCurrencyNumberFormField
func (t *TemplateResponse_FormFields_Item) FromTemplateCurrencyNumberFormField(v TemplateCurrencyNumberFormField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTemplateCurrencyNumberFormField performs a merge with any union data inside the TemplateResponse_FormFields_Item, using the provided TemplateCurrencyNumberFormField
func (t *TemplateResponse_FormFields_Item) MergeTemplateCurrencyNumberFormField(v TemplateCurrencyNumberFormField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTemplateSimpleNumberFormField returns the union data inside the TemplateResponse_FormFields_Item as a TemplateSimpleNumberFormField
func (t TemplateResponse_FormFields_Item) AsTemplateSimpleNumberFormField() (TemplateSimpleNumberFormField, error) {
	var body TemplateSimpleNumberFormField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTemplateSimpleNumberFormField overwrites any union data inside the TemplateResponse_FormFields_Item as the provided TemplateSimpleNumberFormField
func (t *TemplateResponse_FormFields_Item) FromTemplateSimpleNumberFormField(v TemplateSimpleNumberFormField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTemplateSimpleNumberFormField performs a merge with any union data inside the TemplateResponse_FormFields_Item, using the provided TemplateSimpleNumberFormField
func (t *TemplateResponse_FormFields_Item) MergeTemplateSimpleNumberFormField(v TemplateSimpleNumberFormField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTemplatePercentageNumberFormField returns the union data inside the TemplateResponse_FormFields_Item as a TemplatePercentageNumberFormField
func (t TemplateResponse_FormFields_Item) AsTemplatePercentageNumberFormField() (TemplatePercentageNumberFormField, error) {
	var body TemplatePercentageNumberFormField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTemplatePercentageNumberFormField overwrites any union data inside the TemplateResponse_FormFields_Item as the provided TemplatePercentageNumberFormField
func (t *TemplateResponse_FormFields_Item) FromTemplatePercentageNumberFormField(v TemplatePercentageNumberFormField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTemplatePercentageNumberFormField performs a merge with any union data inside the TemplateResponse_FormFields_Item, using the provided TemplatePercentageNumberFormField
func (t *TemplateResponse_FormFields_Item) MergeTemplatePercentageNumberFormField(v TemplatePercentageNumberFormField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTemplateChoiceFormField returns the union data inside the TemplateResponse_FormFields_Item as a TemplateChoiceFormField
func (t TemplateResponse_FormFields_Item) AsTemplateChoiceFormField() (TemplateChoiceFormField, error) {
	var body TemplateChoiceFormField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTemplateChoiceFormField overwrites any union data inside the TemplateResponse_FormFields_Item as the provided TemplateChoiceFormField
func (t *TemplateResponse_FormFields_Item) FromTemplateChoiceFormField(v TemplateChoiceFormField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTemplateChoiceFormField performs a merge with any union data inside the TemplateResponse_FormFields_Item, using the provided TemplateChoiceFormField
func (t *TemplateResponse_FormFields_Item) MergeTemplateChoiceFormField(v TemplateChoiceFormField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTemplateInstructionFormField returns the union data inside the TemplateResponse_FormFields_Item as a TemplateInstructionFormField
func (t TemplateResponse_FormFields_Item) AsTemplateInstructionFormField() (TemplateInstructionFormField, error) {
	var body TemplateInstructionFormField
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTemplateInstructionFormField overwrites any union data inside the TemplateResponse_FormFields_Item as the provided TemplateInstructionFormField
func (t *TemplateResponse_FormFields_Item) FromTemplateInstructionFormField(v TemplateInstructionFormField) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTemplateInstructionFormField performs a merge with any union data inside the TemplateResponse_FormFields_Item, using the provided TemplateInstructionFormField
func (t *TemplateResponse_FormFields_Item) MergeTemplateInstructionFormField(v TemplateInstructionFormField) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t TemplateResponse_FormFields_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *TemplateResponse_FormFields_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsBriefTypeFormFieldRequest returns the union data inside the WorkRequestCreateRequest_FormFields_Item as a BriefTypeFormFieldRequest
func (t WorkRequestCreateRequest_FormFields_Item) AsBriefTypeFormFieldRequest() (BriefTypeFormFieldRequest, error) {
	var body BriefTypeFormFieldRequest
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBriefTypeFormFieldRequest overwrites any union data inside the WorkRequestCreateRequest_FormFields_Item as the provided BriefTypeFormFieldRequest
func (t *WorkRequestCreateRequest_FormFields_Item) FromBriefTypeFormFieldRequest(v BriefTypeFormFieldRequest) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBriefTypeFormFieldRequest performs a merge with any union data inside the WorkRequestCreateRequest_FormFields_Item, using the provided BriefTypeFormFieldRequest
func (t *WorkRequestCreateRequest_FormFields_Item) MergeBriefTypeFormFieldRequest(v BriefTypeFormFieldRequest) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDateTypeFormFieldRequest returns the union data inside the WorkRequestCreateRequest_FormFields_Item as a DateTypeFormFieldRequest
func (t WorkRequestCreateRequest_FormFields_Item) AsDateTypeFormFieldRequest() (DateTypeFormFieldRequest, error) {
	var body DateTypeFormFieldRequest
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDateTypeFormFieldRequest overwrites any union data inside the WorkRequestCreateRequest_FormFields_Item as the provided DateTypeFormFieldRequest
func (t *WorkRequestCreateRequest_FormFields_Item) FromDateTypeFormFieldRequest(v DateTypeFormFieldRequest) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDateTypeFormFieldRequest performs a merge with any union data inside the WorkRequestCreateRequest_FormFields_Item, using the provided DateTypeFormFieldRequest
func (t *WorkRequestCreateRequest_FormFields_Item) MergeDateTypeFormFieldRequest(v DateTypeFormFieldRequest) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFileTypeFormFieldRequest returns the union data inside the WorkRequestCreateRequest_FormFields_Item as a FileTypeFormFieldRequest
func (t WorkRequestCreateRequest_FormFields_Item) AsFileTypeFormFieldRequest() (FileTypeFormFieldRequest, error) {
	var body FileTypeFormFieldRequest
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFileTypeFormFieldRequest overwrites any union data inside the WorkRequestCreateRequest_FormFields_Item as the provided FileTypeFormFieldRequest
func (t *WorkRequestCreateRequest_FormFields_Item) FromFileTypeFormFieldRequest(v FileTypeFormFieldRequest) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFileTypeFormFieldRequest performs a merge with any union data inside the WorkRequestCreateRequest_FormFields_Item, using the provided FileTypeFormFieldRequest
func (t *WorkRequestCreateRequest_FormFields_Item) MergeFileTypeFormFieldRequest(v FileTypeFormFieldRequest) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMultiChoiceTypeFormFieldRequest returns the union data inside the WorkRequestCreateRequest_FormFields_Item as a MultiChoiceTypeFormFieldRequest
func (t WorkRequestCreateRequest_FormFields_Item) AsMultiChoiceTypeFormFieldRequest() (MultiChoiceTypeFormFieldRequest, error) {
	var body MultiChoiceTypeFormFieldRequest
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMultiChoiceTypeFormFieldRequest overwrites any union data inside the WorkRequestCreateRequest_FormFields_Item as the provided MultiChoiceTypeFormFieldRequest
func (t *WorkRequestCreateRequest_FormFields_Item) FromMultiChoiceTypeFormFieldRequest(v MultiChoiceTypeFormFieldRequest) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMultiChoiceTypeFormFieldRequest performs a merge with any union data inside the WorkRequestCreateRequest_FormFields_Item, using the provided MultiChoiceTypeFormFieldRequest
func (t *WorkRequestCreateRequest_FormFields_Item) MergeMultiChoiceTypeFormFieldRequest(v MultiChoiceTypeFormFieldRequest) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsNumberTypeFormFieldRequest returns the union data inside the WorkRequestCreateRequest_FormFields_Item as a NumberTypeFormFieldRequest
func (t WorkRequestCreateRequest_FormFields_Item) AsNumberTypeFormFieldRequest() (NumberTypeFormFieldRequest, error) {
	var body NumberTypeFormFieldRequest
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberTypeFormFieldRequest overwrites any union data inside the WorkRequestCreateRequest_FormFields_Item as the provided NumberTypeFormFieldRequest
func (t *WorkRequestCreateRequest_FormFields_Item) FromNumberTypeFormFieldRequest(v NumberTypeFormFieldRequest) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberTypeFormFieldRequest performs a merge with any union data inside the WorkRequestCreateRequest_FormFields_Item, using the provided NumberTypeFormFieldRequest
func (t *WorkRequestCreateRequest_FormFields_Item) MergeNumberTypeFormFieldRequest(v NumberTypeFormFieldRequest) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRadioButtonTypeFormFieldRequest returns the union data inside the WorkRequestCreateRequest_FormFields_Item as a RadioButtonTypeFormFieldRequest
func (t WorkRequestCreateRequest_FormFields_Item) AsRadioButtonTypeFormFieldRequest() (RadioButtonTypeFormFieldRequest, error) {
	var body RadioButtonTypeFormFieldRequest
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRadioButtonTypeFormFieldRequest overwrites any union data inside the WorkRequestCreateRequest_FormFields_Item as the provided RadioButtonTypeFormFieldRequest
func (t *WorkRequestCreateRequest_FormFields_Item) FromRadioButtonTypeFormFieldRequest(v RadioButtonTypeFormFieldRequest) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRadioButtonTypeFormFieldRequest performs a merge with any union data inside the WorkRequestCreateRequest_FormFields_Item, using the provided RadioButtonTypeFormFieldRequest
func (t *WorkRequestCreateRequest_FormFields_Item) MergeRadioButtonTypeFormFieldRequest(v RadioButtonTypeFormFieldRequest) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTextTypeFormFieldRequest returns the union data inside the WorkRequestCreateRequest_FormFields_Item as a TextTypeFormFieldRequest
func (t WorkRequestCreateRequest_FormFields_Item) AsTextTypeFormFieldRequest() (TextTypeFormFieldRequest, error) {
	var body TextTypeFormFieldRequest
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTextTypeFormFieldRequest overwrites any union data inside the WorkRequestCreateRequest_FormFields_Item as the provided TextTypeFormFieldRequest
func (t *WorkRequestCreateRequest_FormFields_Item) FromTextTypeFormFieldRequest(v TextTypeFormFieldRequest) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTextTypeFormFieldRequest performs a merge with any union data inside the WorkRequestCreateRequest_FormFields_Item, using the provided TextTypeFormFieldRequest
func (t *WorkRequestCreateRequest_FormFields_Item) MergeTextTypeFormFieldRequest(v TextTypeFormFieldRequest) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t WorkRequestCreateRequest_FormFields_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *WorkRequestCreateRequest_FormFields_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkRequestRequestFormFieldBriefTypePayload returns the union data inside the WorkRequestFormFieldUpdateRequest as a WorkRequestRequestFormFieldBriefTypePayload
func (t WorkRequestFormFieldUpdateRequest) AsWorkRequestRequestFormFieldBriefTypePayload() (WorkRequestRequestFormFieldBriefTypePayload, error) {
	var body WorkRequestRequestFormFieldBriefTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldBriefTypePayload overwrites any union data inside the WorkRequestFormFieldUpdateRequest as the provided WorkRequestRequestFormFieldBriefTypePayload
func (t *WorkRequestFormFieldUpdateRequest) FromWorkRequestRequestFormFieldBriefTypePayload(v WorkRequestRequestFormFieldBriefTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldBriefTypePayload performs a merge with any union data inside the WorkRequestFormFieldUpdateRequest, using the provided WorkRequestRequestFormFieldBriefTypePayload
func (t *WorkRequestFormFieldUpdateRequest) MergeWorkRequestRequestFormFieldBriefTypePayload(v WorkRequestRequestFormFieldBriefTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldCheckboxTypePayload returns the union data inside the WorkRequestFormFieldUpdateRequest as a WorkRequestRequestFormFieldCheckboxTypePayload
func (t WorkRequestFormFieldUpdateRequest) AsWorkRequestRequestFormFieldCheckboxTypePayload() (WorkRequestRequestFormFieldCheckboxTypePayload, error) {
	var body WorkRequestRequestFormFieldCheckboxTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldCheckboxTypePayload overwrites any union data inside the WorkRequestFormFieldUpdateRequest as the provided WorkRequestRequestFormFieldCheckboxTypePayload
func (t *WorkRequestFormFieldUpdateRequest) FromWorkRequestRequestFormFieldCheckboxTypePayload(v WorkRequestRequestFormFieldCheckboxTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldCheckboxTypePayload performs a merge with any union data inside the WorkRequestFormFieldUpdateRequest, using the provided WorkRequestRequestFormFieldCheckboxTypePayload
func (t *WorkRequestFormFieldUpdateRequest) MergeWorkRequestRequestFormFieldCheckboxTypePayload(v WorkRequestRequestFormFieldCheckboxTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldCurrencyNumberTypePayload returns the union data inside the WorkRequestFormFieldUpdateRequest as a WorkRequestRequestFormFieldCurrencyNumberTypePayload
func (t WorkRequestFormFieldUpdateRequest) AsWorkRequestRequestFormFieldCurrencyNumberTypePayload() (WorkRequestRequestFormFieldCurrencyNumberTypePayload, error) {
	var body WorkRequestRequestFormFieldCurrencyNumberTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldCurrencyNumberTypePayload overwrites any union data inside the WorkRequestFormFieldUpdateRequest as the provided WorkRequestRequestFormFieldCurrencyNumberTypePayload
func (t *WorkRequestFormFieldUpdateRequest) FromWorkRequestRequestFormFieldCurrencyNumberTypePayload(v WorkRequestRequestFormFieldCurrencyNumberTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldCurrencyNumberTypePayload performs a merge with any union data inside the WorkRequestFormFieldUpdateRequest, using the provided WorkRequestRequestFormFieldCurrencyNumberTypePayload
func (t *WorkRequestFormFieldUpdateRequest) MergeWorkRequestRequestFormFieldCurrencyNumberTypePayload(v WorkRequestRequestFormFieldCurrencyNumberTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldDateTypePayload returns the union data inside the WorkRequestFormFieldUpdateRequest as a WorkRequestRequestFormFieldDateTypePayload
func (t WorkRequestFormFieldUpdateRequest) AsWorkRequestRequestFormFieldDateTypePayload() (WorkRequestRequestFormFieldDateTypePayload, error) {
	var body WorkRequestRequestFormFieldDateTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldDateTypePayload overwrites any union data inside the WorkRequestFormFieldUpdateRequest as the provided WorkRequestRequestFormFieldDateTypePayload
func (t *WorkRequestFormFieldUpdateRequest) FromWorkRequestRequestFormFieldDateTypePayload(v WorkRequestRequestFormFieldDateTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldDateTypePayload performs a merge with any union data inside the WorkRequestFormFieldUpdateRequest, using the provided WorkRequestRequestFormFieldDateTypePayload
func (t *WorkRequestFormFieldUpdateRequest) MergeWorkRequestRequestFormFieldDateTypePayload(v WorkRequestRequestFormFieldDateTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldDropdownTypePayload returns the union data inside the WorkRequestFormFieldUpdateRequest as a WorkRequestRequestFormFieldDropdownTypePayload
func (t WorkRequestFormFieldUpdateRequest) AsWorkRequestRequestFormFieldDropdownTypePayload() (WorkRequestRequestFormFieldDropdownTypePayload, error) {
	var body WorkRequestRequestFormFieldDropdownTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldDropdownTypePayload overwrites any union data inside the WorkRequestFormFieldUpdateRequest as the provided WorkRequestRequestFormFieldDropdownTypePayload
func (t *WorkRequestFormFieldUpdateRequest) FromWorkRequestRequestFormFieldDropdownTypePayload(v WorkRequestRequestFormFieldDropdownTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldDropdownTypePayload performs a merge with any union data inside the WorkRequestFormFieldUpdateRequest, using the provided WorkRequestRequestFormFieldDropdownTypePayload
func (t *WorkRequestFormFieldUpdateRequest) MergeWorkRequestRequestFormFieldDropdownTypePayload(v WorkRequestRequestFormFieldDropdownTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldFileTypePayload returns the union data inside the WorkRequestFormFieldUpdateRequest as a WorkRequestRequestFormFieldFileTypePayload
func (t WorkRequestFormFieldUpdateRequest) AsWorkRequestRequestFormFieldFileTypePayload() (WorkRequestRequestFormFieldFileTypePayload, error) {
	var body WorkRequestRequestFormFieldFileTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldFileTypePayload overwrites any union data inside the WorkRequestFormFieldUpdateRequest as the provided WorkRequestRequestFormFieldFileTypePayload
func (t *WorkRequestFormFieldUpdateRequest) FromWorkRequestRequestFormFieldFileTypePayload(v WorkRequestRequestFormFieldFileTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldFileTypePayload performs a merge with any union data inside the WorkRequestFormFieldUpdateRequest, using the provided WorkRequestRequestFormFieldFileTypePayload
func (t *WorkRequestFormFieldUpdateRequest) MergeWorkRequestRequestFormFieldFileTypePayload(v WorkRequestRequestFormFieldFileTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldLabelTypePayload returns the union data inside the WorkRequestFormFieldUpdateRequest as a WorkRequestRequestFormFieldLabelTypePayload
func (t WorkRequestFormFieldUpdateRequest) AsWorkRequestRequestFormFieldLabelTypePayload() (WorkRequestRequestFormFieldLabelTypePayload, error) {
	var body WorkRequestRequestFormFieldLabelTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldLabelTypePayload overwrites any union data inside the WorkRequestFormFieldUpdateRequest as the provided WorkRequestRequestFormFieldLabelTypePayload
func (t *WorkRequestFormFieldUpdateRequest) FromWorkRequestRequestFormFieldLabelTypePayload(v WorkRequestRequestFormFieldLabelTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldLabelTypePayload performs a merge with any union data inside the WorkRequestFormFieldUpdateRequest, using the provided WorkRequestRequestFormFieldLabelTypePayload
func (t *WorkRequestFormFieldUpdateRequest) MergeWorkRequestRequestFormFieldLabelTypePayload(v WorkRequestRequestFormFieldLabelTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldPercentageNumberTypePayload returns the union data inside the WorkRequestFormFieldUpdateRequest as a WorkRequestRequestFormFieldPercentageNumberTypePayload
func (t WorkRequestFormFieldUpdateRequest) AsWorkRequestRequestFormFieldPercentageNumberTypePayload() (WorkRequestRequestFormFieldPercentageNumberTypePayload, error) {
	var body WorkRequestRequestFormFieldPercentageNumberTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldPercentageNumberTypePayload overwrites any union data inside the WorkRequestFormFieldUpdateRequest as the provided WorkRequestRequestFormFieldPercentageNumberTypePayload
func (t *WorkRequestFormFieldUpdateRequest) FromWorkRequestRequestFormFieldPercentageNumberTypePayload(v WorkRequestRequestFormFieldPercentageNumberTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldPercentageNumberTypePayload performs a merge with any union data inside the WorkRequestFormFieldUpdateRequest, using the provided WorkRequestRequestFormFieldPercentageNumberTypePayload
func (t *WorkRequestFormFieldUpdateRequest) MergeWorkRequestRequestFormFieldPercentageNumberTypePayload(v WorkRequestRequestFormFieldPercentageNumberTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldRadioButtonTypePayload returns the union data inside the WorkRequestFormFieldUpdateRequest as a WorkRequestRequestFormFieldRadioButtonTypePayload
func (t WorkRequestFormFieldUpdateRequest) AsWorkRequestRequestFormFieldRadioButtonTypePayload() (WorkRequestRequestFormFieldRadioButtonTypePayload, error) {
	var body WorkRequestRequestFormFieldRadioButtonTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldRadioButtonTypePayload overwrites any union data inside the WorkRequestFormFieldUpdateRequest as the provided WorkRequestRequestFormFieldRadioButtonTypePayload
func (t *WorkRequestFormFieldUpdateRequest) FromWorkRequestRequestFormFieldRadioButtonTypePayload(v WorkRequestRequestFormFieldRadioButtonTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldRadioButtonTypePayload performs a merge with any union data inside the WorkRequestFormFieldUpdateRequest, using the provided WorkRequestRequestFormFieldRadioButtonTypePayload
func (t *WorkRequestFormFieldUpdateRequest) MergeWorkRequestRequestFormFieldRadioButtonTypePayload(v WorkRequestRequestFormFieldRadioButtonTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldRichtextTypePayload returns the union data inside the WorkRequestFormFieldUpdateRequest as a WorkRequestRequestFormFieldRichtextTypePayload
func (t WorkRequestFormFieldUpdateRequest) AsWorkRequestRequestFormFieldRichtextTypePayload() (WorkRequestRequestFormFieldRichtextTypePayload, error) {
	var body WorkRequestRequestFormFieldRichtextTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldRichtextTypePayload overwrites any union data inside the WorkRequestFormFieldUpdateRequest as the provided WorkRequestRequestFormFieldRichtextTypePayload
func (t *WorkRequestFormFieldUpdateRequest) FromWorkRequestRequestFormFieldRichtextTypePayload(v WorkRequestRequestFormFieldRichtextTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldRichtextTypePayload performs a merge with any union data inside the WorkRequestFormFieldUpdateRequest, using the provided WorkRequestRequestFormFieldRichtextTypePayload
func (t *WorkRequestFormFieldUpdateRequest) MergeWorkRequestRequestFormFieldRichtextTypePayload(v WorkRequestRequestFormFieldRichtextTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldSimpleNumberTypePayload returns the union data inside the WorkRequestFormFieldUpdateRequest as a WorkRequestRequestFormFieldSimpleNumberTypePayload
func (t WorkRequestFormFieldUpdateRequest) AsWorkRequestRequestFormFieldSimpleNumberTypePayload() (WorkRequestRequestFormFieldSimpleNumberTypePayload, error) {
	var body WorkRequestRequestFormFieldSimpleNumberTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldSimpleNumberTypePayload overwrites any union data inside the WorkRequestFormFieldUpdateRequest as the provided WorkRequestRequestFormFieldSimpleNumberTypePayload
func (t *WorkRequestFormFieldUpdateRequest) FromWorkRequestRequestFormFieldSimpleNumberTypePayload(v WorkRequestRequestFormFieldSimpleNumberTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldSimpleNumberTypePayload performs a merge with any union data inside the WorkRequestFormFieldUpdateRequest, using the provided WorkRequestRequestFormFieldSimpleNumberTypePayload
func (t *WorkRequestFormFieldUpdateRequest) MergeWorkRequestRequestFormFieldSimpleNumberTypePayload(v WorkRequestRequestFormFieldSimpleNumberTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldTextAreaTypePayload returns the union data inside the WorkRequestFormFieldUpdateRequest as a WorkRequestRequestFormFieldTextAreaTypePayload
func (t WorkRequestFormFieldUpdateRequest) AsWorkRequestRequestFormFieldTextAreaTypePayload() (WorkRequestRequestFormFieldTextAreaTypePayload, error) {
	var body WorkRequestRequestFormFieldTextAreaTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldTextAreaTypePayload overwrites any union data inside the WorkRequestFormFieldUpdateRequest as the provided WorkRequestRequestFormFieldTextAreaTypePayload
func (t *WorkRequestFormFieldUpdateRequest) FromWorkRequestRequestFormFieldTextAreaTypePayload(v WorkRequestRequestFormFieldTextAreaTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldTextAreaTypePayload performs a merge with any union data inside the WorkRequestFormFieldUpdateRequest, using the provided WorkRequestRequestFormFieldTextAreaTypePayload
func (t *WorkRequestFormFieldUpdateRequest) MergeWorkRequestRequestFormFieldTextAreaTypePayload(v WorkRequestRequestFormFieldTextAreaTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldTextTypePayload returns the union data inside the WorkRequestFormFieldUpdateRequest as a WorkRequestRequestFormFieldTextTypePayload
func (t WorkRequestFormFieldUpdateRequest) AsWorkRequestRequestFormFieldTextTypePayload() (WorkRequestRequestFormFieldTextTypePayload, error) {
	var body WorkRequestRequestFormFieldTextTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldTextTypePayload overwrites any union data inside the WorkRequestFormFieldUpdateRequest as the provided WorkRequestRequestFormFieldTextTypePayload
func (t *WorkRequestFormFieldUpdateRequest) FromWorkRequestRequestFormFieldTextTypePayload(v WorkRequestRequestFormFieldTextTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldTextTypePayload performs a merge with any union data inside the WorkRequestFormFieldUpdateRequest, using the provided WorkRequestRequestFormFieldTextTypePayload
func (t *WorkRequestFormFieldUpdateRequest) MergeWorkRequestRequestFormFieldTextTypePayload(v WorkRequestRequestFormFieldTextTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t WorkRequestFormFieldUpdateRequest) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *WorkRequestFormFieldUpdateRequest) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkRequestTextBriefRequest returns the union data inside the WorkRequestRequestFormFieldBriefTypePayload_Values_Item as a WorkRequestTextBriefRequest
func (t WorkRequestRequestFormFieldBriefTypePayload_Values_Item) AsWorkRequestTextBriefRequest() (WorkRequestTextBriefRequest, error) {
	var body WorkRequestTextBriefRequest
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestTextBriefRequest overwrites any union data inside the WorkRequestRequestFormFieldBriefTypePayload_Values_Item as the provided WorkRequestTextBriefRequest
func (t *WorkRequestRequestFormFieldBriefTypePayload_Values_Item) FromWorkRequestTextBriefRequest(v WorkRequestTextBriefRequest) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestTextBriefRequest performs a merge with any union data inside the WorkRequestRequestFormFieldBriefTypePayload_Values_Item, using the provided WorkRequestTextBriefRequest
func (t *WorkRequestRequestFormFieldBriefTypePayload_Values_Item) MergeWorkRequestTextBriefRequest(v WorkRequestTextBriefRequest) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestAttachmentBriefRequest returns the union data inside the WorkRequestRequestFormFieldBriefTypePayload_Values_Item as a WorkRequestAttachmentBriefRequest
func (t WorkRequestRequestFormFieldBriefTypePayload_Values_Item) AsWorkRequestAttachmentBriefRequest() (WorkRequestAttachmentBriefRequest, error) {
	var body WorkRequestAttachmentBriefRequest
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestAttachmentBriefRequest overwrites any union data inside the WorkRequestRequestFormFieldBriefTypePayload_Values_Item as the provided WorkRequestAttachmentBriefRequest
func (t *WorkRequestRequestFormFieldBriefTypePayload_Values_Item) FromWorkRequestAttachmentBriefRequest(v WorkRequestAttachmentBriefRequest) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestAttachmentBriefRequest performs a merge with any union data inside the WorkRequestRequestFormFieldBriefTypePayload_Values_Item, using the provided WorkRequestAttachmentBriefRequest
func (t *WorkRequestRequestFormFieldBriefTypePayload_Values_Item) MergeWorkRequestAttachmentBriefRequest(v WorkRequestAttachmentBriefRequest) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t WorkRequestRequestFormFieldBriefTypePayload_Values_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *WorkRequestRequestFormFieldBriefTypePayload_Values_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkRequestTextBriefResponse returns the union data inside the WorkRequestRequestFormFieldBriefTypeResponse_Values_Item as a WorkRequestTextBriefResponse
func (t WorkRequestRequestFormFieldBriefTypeResponse_Values_Item) AsWorkRequestTextBriefResponse() (WorkRequestTextBriefResponse, error) {
	var body WorkRequestTextBriefResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestTextBriefResponse overwrites any union data inside the WorkRequestRequestFormFieldBriefTypeResponse_Values_Item as the provided WorkRequestTextBriefResponse
func (t *WorkRequestRequestFormFieldBriefTypeResponse_Values_Item) FromWorkRequestTextBriefResponse(v WorkRequestTextBriefResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestTextBriefResponse performs a merge with any union data inside the WorkRequestRequestFormFieldBriefTypeResponse_Values_Item, using the provided WorkRequestTextBriefResponse
func (t *WorkRequestRequestFormFieldBriefTypeResponse_Values_Item) MergeWorkRequestTextBriefResponse(v WorkRequestTextBriefResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestAttachmentBriefResponse returns the union data inside the WorkRequestRequestFormFieldBriefTypeResponse_Values_Item as a WorkRequestAttachmentBriefResponse
func (t WorkRequestRequestFormFieldBriefTypeResponse_Values_Item) AsWorkRequestAttachmentBriefResponse() (WorkRequestAttachmentBriefResponse, error) {
	var body WorkRequestAttachmentBriefResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestAttachmentBriefResponse overwrites any union data inside the WorkRequestRequestFormFieldBriefTypeResponse_Values_Item as the provided WorkRequestAttachmentBriefResponse
func (t *WorkRequestRequestFormFieldBriefTypeResponse_Values_Item) FromWorkRequestAttachmentBriefResponse(v WorkRequestAttachmentBriefResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestAttachmentBriefResponse performs a merge with any union data inside the WorkRequestRequestFormFieldBriefTypeResponse_Values_Item, using the provided WorkRequestAttachmentBriefResponse
func (t *WorkRequestRequestFormFieldBriefTypeResponse_Values_Item) MergeWorkRequestAttachmentBriefResponse(v WorkRequestAttachmentBriefResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t WorkRequestRequestFormFieldBriefTypeResponse_Values_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *WorkRequestRequestFormFieldBriefTypeResponse_Values_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkRequestRequestFormFieldBriefTypeResponse returns the union data inside the WorkRequestRequestFormFieldUpdateResponse as a WorkRequestRequestFormFieldBriefTypeResponse
func (t WorkRequestRequestFormFieldUpdateResponse) AsWorkRequestRequestFormFieldBriefTypeResponse() (WorkRequestRequestFormFieldBriefTypeResponse, error) {
	var body WorkRequestRequestFormFieldBriefTypeResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldBriefTypeResponse overwrites any union data inside the WorkRequestRequestFormFieldUpdateResponse as the provided WorkRequestRequestFormFieldBriefTypeResponse
func (t *WorkRequestRequestFormFieldUpdateResponse) FromWorkRequestRequestFormFieldBriefTypeResponse(v WorkRequestRequestFormFieldBriefTypeResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldBriefTypeResponse performs a merge with any union data inside the WorkRequestRequestFormFieldUpdateResponse, using the provided WorkRequestRequestFormFieldBriefTypeResponse
func (t *WorkRequestRequestFormFieldUpdateResponse) MergeWorkRequestRequestFormFieldBriefTypeResponse(v WorkRequestRequestFormFieldBriefTypeResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldCheckboxTypeResponse returns the union data inside the WorkRequestRequestFormFieldUpdateResponse as a WorkRequestRequestFormFieldCheckboxTypeResponse
func (t WorkRequestRequestFormFieldUpdateResponse) AsWorkRequestRequestFormFieldCheckboxTypeResponse() (WorkRequestRequestFormFieldCheckboxTypeResponse, error) {
	var body WorkRequestRequestFormFieldCheckboxTypeResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldCheckboxTypeResponse overwrites any union data inside the WorkRequestRequestFormFieldUpdateResponse as the provided WorkRequestRequestFormFieldCheckboxTypeResponse
func (t *WorkRequestRequestFormFieldUpdateResponse) FromWorkRequestRequestFormFieldCheckboxTypeResponse(v WorkRequestRequestFormFieldCheckboxTypeResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldCheckboxTypeResponse performs a merge with any union data inside the WorkRequestRequestFormFieldUpdateResponse, using the provided WorkRequestRequestFormFieldCheckboxTypeResponse
func (t *WorkRequestRequestFormFieldUpdateResponse) MergeWorkRequestRequestFormFieldCheckboxTypeResponse(v WorkRequestRequestFormFieldCheckboxTypeResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldCurrencyNumberTypePayload returns the union data inside the WorkRequestRequestFormFieldUpdateResponse as a WorkRequestRequestFormFieldCurrencyNumberTypePayload
func (t WorkRequestRequestFormFieldUpdateResponse) AsWorkRequestRequestFormFieldCurrencyNumberTypePayload() (WorkRequestRequestFormFieldCurrencyNumberTypePayload, error) {
	var body WorkRequestRequestFormFieldCurrencyNumberTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldCurrencyNumberTypePayload overwrites any union data inside the WorkRequestRequestFormFieldUpdateResponse as the provided WorkRequestRequestFormFieldCurrencyNumberTypePayload
func (t *WorkRequestRequestFormFieldUpdateResponse) FromWorkRequestRequestFormFieldCurrencyNumberTypePayload(v WorkRequestRequestFormFieldCurrencyNumberTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldCurrencyNumberTypePayload performs a merge with any union data inside the WorkRequestRequestFormFieldUpdateResponse, using the provided WorkRequestRequestFormFieldCurrencyNumberTypePayload
func (t *WorkRequestRequestFormFieldUpdateResponse) MergeWorkRequestRequestFormFieldCurrencyNumberTypePayload(v WorkRequestRequestFormFieldCurrencyNumberTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldDateTypePayload returns the union data inside the WorkRequestRequestFormFieldUpdateResponse as a WorkRequestRequestFormFieldDateTypePayload
func (t WorkRequestRequestFormFieldUpdateResponse) AsWorkRequestRequestFormFieldDateTypePayload() (WorkRequestRequestFormFieldDateTypePayload, error) {
	var body WorkRequestRequestFormFieldDateTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldDateTypePayload overwrites any union data inside the WorkRequestRequestFormFieldUpdateResponse as the provided WorkRequestRequestFormFieldDateTypePayload
func (t *WorkRequestRequestFormFieldUpdateResponse) FromWorkRequestRequestFormFieldDateTypePayload(v WorkRequestRequestFormFieldDateTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldDateTypePayload performs a merge with any union data inside the WorkRequestRequestFormFieldUpdateResponse, using the provided WorkRequestRequestFormFieldDateTypePayload
func (t *WorkRequestRequestFormFieldUpdateResponse) MergeWorkRequestRequestFormFieldDateTypePayload(v WorkRequestRequestFormFieldDateTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldDropdownTypeResponse returns the union data inside the WorkRequestRequestFormFieldUpdateResponse as a WorkRequestRequestFormFieldDropdownTypeResponse
func (t WorkRequestRequestFormFieldUpdateResponse) AsWorkRequestRequestFormFieldDropdownTypeResponse() (WorkRequestRequestFormFieldDropdownTypeResponse, error) {
	var body WorkRequestRequestFormFieldDropdownTypeResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldDropdownTypeResponse overwrites any union data inside the WorkRequestRequestFormFieldUpdateResponse as the provided WorkRequestRequestFormFieldDropdownTypeResponse
func (t *WorkRequestRequestFormFieldUpdateResponse) FromWorkRequestRequestFormFieldDropdownTypeResponse(v WorkRequestRequestFormFieldDropdownTypeResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldDropdownTypeResponse performs a merge with any union data inside the WorkRequestRequestFormFieldUpdateResponse, using the provided WorkRequestRequestFormFieldDropdownTypeResponse
func (t *WorkRequestRequestFormFieldUpdateResponse) MergeWorkRequestRequestFormFieldDropdownTypeResponse(v WorkRequestRequestFormFieldDropdownTypeResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldFileTypeResponse returns the union data inside the WorkRequestRequestFormFieldUpdateResponse as a WorkRequestRequestFormFieldFileTypeResponse
func (t WorkRequestRequestFormFieldUpdateResponse) AsWorkRequestRequestFormFieldFileTypeResponse() (WorkRequestRequestFormFieldFileTypeResponse, error) {
	var body WorkRequestRequestFormFieldFileTypeResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldFileTypeResponse overwrites any union data inside the WorkRequestRequestFormFieldUpdateResponse as the provided WorkRequestRequestFormFieldFileTypeResponse
func (t *WorkRequestRequestFormFieldUpdateResponse) FromWorkRequestRequestFormFieldFileTypeResponse(v WorkRequestRequestFormFieldFileTypeResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldFileTypeResponse performs a merge with any union data inside the WorkRequestRequestFormFieldUpdateResponse, using the provided WorkRequestRequestFormFieldFileTypeResponse
func (t *WorkRequestRequestFormFieldUpdateResponse) MergeWorkRequestRequestFormFieldFileTypeResponse(v WorkRequestRequestFormFieldFileTypeResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldLabelTypeResponse returns the union data inside the WorkRequestRequestFormFieldUpdateResponse as a WorkRequestRequestFormFieldLabelTypeResponse
func (t WorkRequestRequestFormFieldUpdateResponse) AsWorkRequestRequestFormFieldLabelTypeResponse() (WorkRequestRequestFormFieldLabelTypeResponse, error) {
	var body WorkRequestRequestFormFieldLabelTypeResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldLabelTypeResponse overwrites any union data inside the WorkRequestRequestFormFieldUpdateResponse as the provided WorkRequestRequestFormFieldLabelTypeResponse
func (t *WorkRequestRequestFormFieldUpdateResponse) FromWorkRequestRequestFormFieldLabelTypeResponse(v WorkRequestRequestFormFieldLabelTypeResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldLabelTypeResponse performs a merge with any union data inside the WorkRequestRequestFormFieldUpdateResponse, using the provided WorkRequestRequestFormFieldLabelTypeResponse
func (t *WorkRequestRequestFormFieldUpdateResponse) MergeWorkRequestRequestFormFieldLabelTypeResponse(v WorkRequestRequestFormFieldLabelTypeResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldPercentageNumberTypePayload returns the union data inside the WorkRequestRequestFormFieldUpdateResponse as a WorkRequestRequestFormFieldPercentageNumberTypePayload
func (t WorkRequestRequestFormFieldUpdateResponse) AsWorkRequestRequestFormFieldPercentageNumberTypePayload() (WorkRequestRequestFormFieldPercentageNumberTypePayload, error) {
	var body WorkRequestRequestFormFieldPercentageNumberTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldPercentageNumberTypePayload overwrites any union data inside the WorkRequestRequestFormFieldUpdateResponse as the provided WorkRequestRequestFormFieldPercentageNumberTypePayload
func (t *WorkRequestRequestFormFieldUpdateResponse) FromWorkRequestRequestFormFieldPercentageNumberTypePayload(v WorkRequestRequestFormFieldPercentageNumberTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldPercentageNumberTypePayload performs a merge with any union data inside the WorkRequestRequestFormFieldUpdateResponse, using the provided WorkRequestRequestFormFieldPercentageNumberTypePayload
func (t *WorkRequestRequestFormFieldUpdateResponse) MergeWorkRequestRequestFormFieldPercentageNumberTypePayload(v WorkRequestRequestFormFieldPercentageNumberTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldRadioButtonTypeResponse returns the union data inside the WorkRequestRequestFormFieldUpdateResponse as a WorkRequestRequestFormFieldRadioButtonTypeResponse
func (t WorkRequestRequestFormFieldUpdateResponse) AsWorkRequestRequestFormFieldRadioButtonTypeResponse() (WorkRequestRequestFormFieldRadioButtonTypeResponse, error) {
	var body WorkRequestRequestFormFieldRadioButtonTypeResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldRadioButtonTypeResponse overwrites any union data inside the WorkRequestRequestFormFieldUpdateResponse as the provided WorkRequestRequestFormFieldRadioButtonTypeResponse
func (t *WorkRequestRequestFormFieldUpdateResponse) FromWorkRequestRequestFormFieldRadioButtonTypeResponse(v WorkRequestRequestFormFieldRadioButtonTypeResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldRadioButtonTypeResponse performs a merge with any union data inside the WorkRequestRequestFormFieldUpdateResponse, using the provided WorkRequestRequestFormFieldRadioButtonTypeResponse
func (t *WorkRequestRequestFormFieldUpdateResponse) MergeWorkRequestRequestFormFieldRadioButtonTypeResponse(v WorkRequestRequestFormFieldRadioButtonTypeResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldRichtextTypePayload returns the union data inside the WorkRequestRequestFormFieldUpdateResponse as a WorkRequestRequestFormFieldRichtextTypePayload
func (t WorkRequestRequestFormFieldUpdateResponse) AsWorkRequestRequestFormFieldRichtextTypePayload() (WorkRequestRequestFormFieldRichtextTypePayload, error) {
	var body WorkRequestRequestFormFieldRichtextTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldRichtextTypePayload overwrites any union data inside the WorkRequestRequestFormFieldUpdateResponse as the provided WorkRequestRequestFormFieldRichtextTypePayload
func (t *WorkRequestRequestFormFieldUpdateResponse) FromWorkRequestRequestFormFieldRichtextTypePayload(v WorkRequestRequestFormFieldRichtextTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldRichtextTypePayload performs a merge with any union data inside the WorkRequestRequestFormFieldUpdateResponse, using the provided WorkRequestRequestFormFieldRichtextTypePayload
func (t *WorkRequestRequestFormFieldUpdateResponse) MergeWorkRequestRequestFormFieldRichtextTypePayload(v WorkRequestRequestFormFieldRichtextTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldSimpleNumberTypePayload returns the union data inside the WorkRequestRequestFormFieldUpdateResponse as a WorkRequestRequestFormFieldSimpleNumberTypePayload
func (t WorkRequestRequestFormFieldUpdateResponse) AsWorkRequestRequestFormFieldSimpleNumberTypePayload() (WorkRequestRequestFormFieldSimpleNumberTypePayload, error) {
	var body WorkRequestRequestFormFieldSimpleNumberTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldSimpleNumberTypePayload overwrites any union data inside the WorkRequestRequestFormFieldUpdateResponse as the provided WorkRequestRequestFormFieldSimpleNumberTypePayload
func (t *WorkRequestRequestFormFieldUpdateResponse) FromWorkRequestRequestFormFieldSimpleNumberTypePayload(v WorkRequestRequestFormFieldSimpleNumberTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldSimpleNumberTypePayload performs a merge with any union data inside the WorkRequestRequestFormFieldUpdateResponse, using the provided WorkRequestRequestFormFieldSimpleNumberTypePayload
func (t *WorkRequestRequestFormFieldUpdateResponse) MergeWorkRequestRequestFormFieldSimpleNumberTypePayload(v WorkRequestRequestFormFieldSimpleNumberTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldTextAreaTypePayload returns the union data inside the WorkRequestRequestFormFieldUpdateResponse as a WorkRequestRequestFormFieldTextAreaTypePayload
func (t WorkRequestRequestFormFieldUpdateResponse) AsWorkRequestRequestFormFieldTextAreaTypePayload() (WorkRequestRequestFormFieldTextAreaTypePayload, error) {
	var body WorkRequestRequestFormFieldTextAreaTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldTextAreaTypePayload overwrites any union data inside the WorkRequestRequestFormFieldUpdateResponse as the provided WorkRequestRequestFormFieldTextAreaTypePayload
func (t *WorkRequestRequestFormFieldUpdateResponse) FromWorkRequestRequestFormFieldTextAreaTypePayload(v WorkRequestRequestFormFieldTextAreaTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldTextAreaTypePayload performs a merge with any union data inside the WorkRequestRequestFormFieldUpdateResponse, using the provided WorkRequestRequestFormFieldTextAreaTypePayload
func (t *WorkRequestRequestFormFieldUpdateResponse) MergeWorkRequestRequestFormFieldTextAreaTypePayload(v WorkRequestRequestFormFieldTextAreaTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsWorkRequestRequestFormFieldTextTypePayload returns the union data inside the WorkRequestRequestFormFieldUpdateResponse as a WorkRequestRequestFormFieldTextTypePayload
func (t WorkRequestRequestFormFieldUpdateResponse) AsWorkRequestRequestFormFieldTextTypePayload() (WorkRequestRequestFormFieldTextTypePayload, error) {
	var body WorkRequestRequestFormFieldTextTypePayload
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkRequestRequestFormFieldTextTypePayload overwrites any union data inside the WorkRequestRequestFormFieldUpdateResponse as the provided WorkRequestRequestFormFieldTextTypePayload
func (t *WorkRequestRequestFormFieldUpdateResponse) FromWorkRequestRequestFormFieldTextTypePayload(v WorkRequestRequestFormFieldTextTypePayload) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkRequestRequestFormFieldTextTypePayload performs a merge with any union data inside the WorkRequestRequestFormFieldUpdateResponse, using the provided WorkRequestRequestFormFieldTextTypePayload
func (t *WorkRequestRequestFormFieldUpdateResponse) MergeWorkRequestRequestFormFieldTextTypePayload(v WorkRequestRequestFormFieldTextTypePayload) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t WorkRequestRequestFormFieldUpdateResponse) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *WorkRequestRequestFormFieldUpdateResponse) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsBriefTypeFormFieldResponse returns the union data inside the WorkRequestResponse_FormFields_Item as a BriefTypeFormFieldResponse
func (t WorkRequestResponse_FormFields_Item) AsBriefTypeFormFieldResponse() (BriefTypeFormFieldResponse, error) {
	var body BriefTypeFormFieldResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBriefTypeFormFieldResponse overwrites any union data inside the WorkRequestResponse_FormFields_Item as the provided BriefTypeFormFieldResponse
func (t *WorkRequestResponse_FormFields_Item) FromBriefTypeFormFieldResponse(v BriefTypeFormFieldResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBriefTypeFormFieldResponse performs a merge with any union data inside the WorkRequestResponse_FormFields_Item, using the provided BriefTypeFormFieldResponse
func (t *WorkRequestResponse_FormFields_Item) MergeBriefTypeFormFieldResponse(v BriefTypeFormFieldResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsDateTypeFormFieldResponse returns the union data inside the WorkRequestResponse_FormFields_Item as a DateTypeFormFieldResponse
func (t WorkRequestResponse_FormFields_Item) AsDateTypeFormFieldResponse() (DateTypeFormFieldResponse, error) {
	var body DateTypeFormFieldResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDateTypeFormFieldResponse overwrites any union data inside the WorkRequestResponse_FormFields_Item as the provided DateTypeFormFieldResponse
func (t *WorkRequestResponse_FormFields_Item) FromDateTypeFormFieldResponse(v DateTypeFormFieldResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDateTypeFormFieldResponse performs a merge with any union data inside the WorkRequestResponse_FormFields_Item, using the provided DateTypeFormFieldResponse
func (t *WorkRequestResponse_FormFields_Item) MergeDateTypeFormFieldResponse(v DateTypeFormFieldResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFileTypeFormFieldResponse returns the union data inside the WorkRequestResponse_FormFields_Item as a FileTypeFormFieldResponse
func (t WorkRequestResponse_FormFields_Item) AsFileTypeFormFieldResponse() (FileTypeFormFieldResponse, error) {
	var body FileTypeFormFieldResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFileTypeFormFieldResponse overwrites any union data inside the WorkRequestResponse_FormFields_Item as the provided FileTypeFormFieldResponse
func (t *WorkRequestResponse_FormFields_Item) FromFileTypeFormFieldResponse(v FileTypeFormFieldResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFileTypeFormFieldResponse performs a merge with any union data inside the WorkRequestResponse_FormFields_Item, using the provided FileTypeFormFieldResponse
func (t *WorkRequestResponse_FormFields_Item) MergeFileTypeFormFieldResponse(v FileTypeFormFieldResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMultiChoiceTypeFormFieldResponse returns the union data inside the WorkRequestResponse_FormFields_Item as a MultiChoiceTypeFormFieldResponse
func (t WorkRequestResponse_FormFields_Item) AsMultiChoiceTypeFormFieldResponse() (MultiChoiceTypeFormFieldResponse, error) {
	var body MultiChoiceTypeFormFieldResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMultiChoiceTypeFormFieldResponse overwrites any union data inside the WorkRequestResponse_FormFields_Item as the provided MultiChoiceTypeFormFieldResponse
func (t *WorkRequestResponse_FormFields_Item) FromMultiChoiceTypeFormFieldResponse(v MultiChoiceTypeFormFieldResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMultiChoiceTypeFormFieldResponse performs a merge with any union data inside the WorkRequestResponse_FormFields_Item, using the provided MultiChoiceTypeFormFieldResponse
func (t *WorkRequestResponse_FormFields_Item) MergeMultiChoiceTypeFormFieldResponse(v MultiChoiceTypeFormFieldResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsNumberTypeFormFieldResponse returns the union data inside the WorkRequestResponse_FormFields_Item as a NumberTypeFormFieldResponse
func (t WorkRequestResponse_FormFields_Item) AsNumberTypeFormFieldResponse() (NumberTypeFormFieldResponse, error) {
	var body NumberTypeFormFieldResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNumberTypeFormFieldResponse overwrites any union data inside the WorkRequestResponse_FormFields_Item as the provided NumberTypeFormFieldResponse
func (t *WorkRequestResponse_FormFields_Item) FromNumberTypeFormFieldResponse(v NumberTypeFormFieldResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNumberTypeFormFieldResponse performs a merge with any union data inside the WorkRequestResponse_FormFields_Item, using the provided NumberTypeFormFieldResponse
func (t *WorkRequestResponse_FormFields_Item) MergeNumberTypeFormFieldResponse(v NumberTypeFormFieldResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRadioButtonTypeFormFieldResponse returns the union data inside the WorkRequestResponse_FormFields_Item as a RadioButtonTypeFormFieldResponse
func (t WorkRequestResponse_FormFields_Item) AsRadioButtonTypeFormFieldResponse() (RadioButtonTypeFormFieldResponse, error) {
	var body RadioButtonTypeFormFieldResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRadioButtonTypeFormFieldResponse overwrites any union data inside the WorkRequestResponse_FormFields_Item as the provided RadioButtonTypeFormFieldResponse
func (t *WorkRequestResponse_FormFields_Item) FromRadioButtonTypeFormFieldResponse(v RadioButtonTypeFormFieldResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRadioButtonTypeFormFieldResponse performs a merge with any union data inside the WorkRequestResponse_FormFields_Item, using the provided RadioButtonTypeFormFieldResponse
func (t *WorkRequestResponse_FormFields_Item) MergeRadioButtonTypeFormFieldResponse(v RadioButtonTypeFormFieldResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTextTypeFormFieldResponse returns the union data inside the WorkRequestResponse_FormFields_Item as a TextTypeFormFieldResponse
func (t WorkRequestResponse_FormFields_Item) AsTextTypeFormFieldResponse() (TextTypeFormFieldResponse, error) {
	var body TextTypeFormFieldResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTextTypeFormFieldResponse overwrites any union data inside the WorkRequestResponse_FormFields_Item as the provided TextTypeFormFieldResponse
func (t *WorkRequestResponse_FormFields_Item) FromTextTypeFormFieldResponse(v TextTypeFormFieldResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTextTypeFormFieldResponse performs a merge with any union data inside the WorkRequestResponse_FormFields_Item, using the provided TextTypeFormFieldResponse
func (t *WorkRequestResponse_FormFields_Item) MergeTextTypeFormFieldResponse(v TextTypeFormFieldResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t WorkRequestResponse_FormFields_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *WorkRequestResponse_FormFields_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
