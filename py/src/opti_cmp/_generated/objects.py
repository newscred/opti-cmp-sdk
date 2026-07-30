# Auto-generated - DO NOT EDIT
# mypy: disable-error-code="misc,assignment"
"""Response shapes, declaring attributes for the type checker.

Derived from `schema.py`. These classes are never instantiated: every response
body is an `APIObject` holding exactly what the API returned. Declaring the
fields here is what makes `response.data.title` type-check, while leaving fields
the specification does not yet know about reachable at runtime.

Optional fields are declared `T | None`. A field the API omits entirely raises
`AttributeError` rather than returning `None`; use `.get("field")` for those.

A few shapes keep the functional `TypedDict` form because their keys are not
valid Python identifiers (`x-amz-signature`). Those stay subscript-only, which
is the only thing that could work.
"""

from __future__ import annotations

from typing import Any, Literal, TypeAlias, TypedDict

from typing_extensions import NotRequired

from .._object import APIObject


class AddUrlToTaskRequest(APIObject):
    title: str | None
    url: str


class AllowedContentTypeItem(APIObject):
    name: str
    url: str


class ArticleAuthorResponse(APIObject):
    name: str | None


class AssetContent(APIObject):
    type: Literal["url", "api_url", "html_body"]
    value: str


class Choice(APIObject):
    id: str
    name: str


AssetFieldTypes: TypeAlias = Literal[
    "checkbox",
    "currency_number",
    "date",
    "dropdown",
    "label",
    "radio_button",
    "text_area",
    "percentage_number",
    "rich_text",
    "simple_number",
    "text",
]


class Links(APIObject):
    asset: str
    asset_fields: str


class AssetFieldUpdateResponse(APIObject):
    id: str
    links: Links
    name: str
    type: AssetFieldTypes
    values: list[str]


class AssetLineageCreateRequest(APIObject):
    icon_url: str | None
    name: str
    rendition_id: str | None
    uri: str


class Links1(APIObject):
    asset: str
    self: str


class AssetLineageResponse(APIObject):
    asset_id: str
    created_at: str
    icon_url: str | None
    id: str
    links: Links1
    name: str
    rendition_id: str | None
    uri: str
    used_in: Literal["external"]
    version_id: str


class Permission(APIObject):
    access_type: Literal["view", "edit"]
    id: str
    is_owner: bool | None


class AssetPermissionBulkCreateRequest(APIObject):
    permissions: list[Permission]
    type: Literal["user", "team"]


class AssetPermissionUpdateRequest(APIObject):
    access_type: Literal["view", "edit"] | None
    is_owner: bool | None


class Links2(APIObject):
    self: str


class Value(APIObject):
    key: str
    title: str


class AssetVersionCreateRequest(APIObject):
    key: str
    title: str


class AttachmentRequest(APIObject):
    key: str
    name: str


class AttachmentResponse(APIObject):
    id: str
    name: str
    url: str


class BaseAssetField(APIObject):
    id: str
    name: str
    type: str
    values: list[str]


class BaseAssetRenditionResponse(APIObject):
    created_at: str
    id: str
    mime_type: str
    name: str
    url: str


class Links3(APIObject):
    self: str | None
    versions: str | None


class BaseContentTypeModel(APIObject):
    component: bool
    content_type_guid: str
    created_at: str
    created_by: str
    description: str | None
    disabled: bool | None
    links: Links3
    name: str
    source: str | None
    source_id: str | None
    source_metadata: str | None
    thumbnail_guid: str | None
    updated_at: str
    updated_by: str


class Links4(APIObject):
    content_type: str | None
    self: str | None


class BaseContentTypeVersionModel(APIObject):
    created_at: str
    created_by: str
    expected_locales: list[str]
    latest: bool
    links: Links4
    version_guid: str


BaseFieldDefinitionType: TypeAlias = Literal["boolean", "json", "url", "location"]


class BaseFormFieldRequest(APIObject):
    identifier: str
    type: str
    values: list[Any]


class BaseFormFieldResponse(APIObject):
    identifier: str
    is_protected: bool | None
    type: str
    values: list[Any]


class BaseObjectFieldUpdatePayload(APIObject):
    type: str
    values: list[Any]


class Links5(APIObject):
    accessor: str | None


class BasePermissionsResponseSchema(APIObject):
    access_type: Literal["view", "edit", "comment", "delete"]
    id: str
    is_owner: bool
    links: Links5
    name: str
    type: Literal["user", "team", "organization"]


class BaseSettingsFieldCreatePayload(APIObject):
    helper_text: str | None
    is_active: bool
    name: str
    type: str


class BaseSettingsFieldUpdatePayload(APIObject):
    helper_text: str | None
    is_active: bool | None
    name: str | None
    type: str


class Links6(APIObject):
    source: str


class BaseSettingsFieldsResponse(APIObject):
    helper_text: str
    id: str
    is_active: bool
    links: Links6
    name: str
    type: str


class BaseTemplateFormField(APIObject):
    helper_text: str | None
    identifier: str
    is_protected: bool | None
    is_readonly: bool
    is_required: bool
    label: str | None
    sort_order: int
    type: Literal[
        "brief",
        "checkbox",
        "currency_number",
        "date",
        "dropdown",
        "file",
        "instruction",
        "label",
        "percentage_number",
        "radio_button",
        "richtext",
        "section",
        "simple_number",
        "text",
        "text_area",
    ]


class BaseTemplateResponse(APIObject):
    applicable_to: list[Literal["work_request", "task_brief", "campaign_brief"]]
    description: str
    id: str
    is_active: bool
    title: str


class BatchFileUrlResponse(APIObject):
    urls: dict[str, str | None]


class BooleanFieldValueModel(APIObject):
    bool_value: bool
    order_index: int | None


class Criterion(APIObject):
    description: str | None
    id: str
    name: str | None


class BrandComplianceCategoriesResponse(APIObject):
    criteria: list[Criterion]
    id: str
    name: str | None


class BudgetResponse(APIObject):
    budgeted_amount: str
    currency_code: str


class Field(APIObject):
    name: str
    value: list[str]


class Links7(APIObject):
    campaign: str
    self: str


class CampaignBriefTemplateValue(APIObject):
    id: str
    name: str


class Links8(APIObject):
    campaign: str
    comment_by: str


class CampaignCommentResponse(APIObject):
    attachments: list[AttachmentResponse]
    comment_by: str
    created_at: str
    id: str
    is_resolved: bool
    links: Links8
    modified_at: str
    parent_comment_id: str | None
    value: str


class CampaignCreateRequest(APIObject):
    color: (
        Literal[
            "#333",
            "#ec2e3b",
            "#fc594b",
            "#e2575b",
            "#f45097",
            "#de55a9",
            "#bd417f",
            "#a959cc",
            "#6456b7",
            "#4db2fc",
            "#3a97be",
            "#36c5a3",
            "#9bc94e",
            "#fdad3c",
            "#ef803c",
            "#fd8876",
            "#ea868a",
            "#f781b5",
            "#e785c1",
            "#cf69a3",
            "#c187db",
            "#9086cc",
            "#79c8fd",
            "#5ab5d0",
            "#4fd5bd",
            "#b8d877",
            "#fec454",
            "#f3a35a",
            "#d3998f",
            "#c79697",
            "#d096af",
            "#c897b4",
            "#ae7c98",
            "#b094be",
            "#908bae",
            "#97bed8",
            "#7aa8b6",
            "#7abeb2",
            "#bece9d",
            "#e3c68d",
            "#d3ac85",
            "#b9a29e",
            "#b29f9f",
            "#b9a1ac",
            "#b5a1ae",
            "#9a8692",
            "#a89bac",
            "#908e9b",
            "#a8b9c3",
            "#8ea1a7",
            "#96b1ac",
            "#c1c8b5",
            "#d4c8b0",
            "#c0b0a1",
            "#212121",
            "#CE515B",
            "#F16281",
            "#EB7777",
            "#A25555",
            "#B05F8C",
            "#AF8EC3",
            "#21B7EC",
            "#3AB0C9",
            "#669BA8",
            "#7ABEB2",
            "#5F9DCE",
            "#717CB9",
            "#4C557E",
            "#449C6C",
            "#63B05F",
            "#4D9B93",
            "#ACC050",
            "#A29F55",
            "#8AA872",
            "#E4844C",
            "#EE9808",
            "#DE7009",
            "#FCD404",
            "#CBB039",
            "#99896A",
            "#2e2a26",
            "#ccccdc",
        ]
        | None
    )
    description: str | None
    end_date: str | None
    owner_id: str | None
    parent_campaign: str | None
    start_date: str | None
    title: str


CampaignFieldTypes: TypeAlias = Literal[
    "checkbox",
    "currency_number",
    "date",
    "dropdown",
    "label",
    "radio_button",
    "text_area",
    "percentage_number",
    "rich_text",
    "simple_number",
    "text",
    "image",
    "video",
]


class Links9(APIObject):
    campaign: str
    campaign_fields: str


class CampaignFieldUpdateResponse(APIObject):
    id: str
    links: Links9
    name: str
    type: CampaignFieldTypes
    values: list[str]


class Links10(APIObject):
    self: str


class Owner(APIObject):
    id: str | None


class CampaignListResponseItem(APIObject):
    description: str | None
    end_date: str | None
    id: str
    is_hidden: bool
    links: Links10
    owner: Owner
    reference_id: str
    start_date: str | None
    title: str


class Links11(APIObject):
    brief: str | None
    child_campaigns: list[str]
    owner: str | None
    parent_campaign: str | None
    self: str


class CampaignUpdateRequest(APIObject):
    end_date: str | None
    owner_id: str | None
    start_date: str | None
    title: str | None


class Choice4(APIObject):
    name: str


class CheckboxAndRadioTypeSettingsFieldCreatePayload(BaseSettingsFieldCreatePayload):
    choices: list[Choice4]


class Choice5(APIObject):
    id: str
    name: str


class CheckboxAndRadioTypeSettingsFieldResponse(BaseSettingsFieldsResponse):
    choices: list[Choice5]


ChoiceDisplayOption: TypeAlias = Literal["radio", "dropdown", "tag", "checkbox"]


class ChoiceFieldValueModel(APIObject):
    choice_key: str
    order_index: int | None


class CommentCreateRequest(APIObject):
    attachments: list[AttachmentRequest] | None
    value: str


class CommentWithReplyCreateRequest(APIObject):
    attachments: list[AttachmentRequest] | None
    parent_comment_id: str | None
    value: str


class Links12(APIObject):
    definition: str | None
    self: str | None


class ContentFieldValueModel(APIObject):
    content_guid: str
    content_url: str | None
    embedded: bool
    order_index: int | None


class ContentMigrationSummary(APIObject):
    errored: int | None
    not_started: int | None
    skipped: int | None
    succeeded: int | None
    total: int | None


ContentTypeFieldEmbedMixConfig: TypeAlias = Literal[1, 2, 3]
ContentTypeListingOption: TypeAlias = Literal["1", "2", "3"]


class CoreContentType(APIObject):
    component: bool
    description: str | None
    disabled: bool | None
    name: str
    thumbnail_guid: str | None


class CreativeAssetRequest(APIObject):
    key: str
    name: str


class CreativeAssetResponse(APIObject):
    id: str
    name: str
    url: str


class CurrencyCustomField(APIObject):
    currency_code: str
    decimal_places: int | None
    has_thousand_separator: bool
    label: str


class DateTypeFormFieldRequest(BaseFormFieldRequest):
    values: list[str] | None


class DateTypeFormFieldResponse(BaseFormFieldResponse):
    values: list[str] | None


class DateTypeObjectFieldUpdatePayload(BaseObjectFieldUpdatePayload):
    values: list[str] | None


class DatetimeFieldValueModel(APIObject):
    datetime_value: str
    order_index: int | None


class DeleteFieldValueModel(APIObject):
    delete: bool
    order_index: int | None


class DeleteLocaleFieldValueModel(APIObject):
    delete: bool
    locale: str


class DetailedAssetRenditionResponse(BaseAssetRenditionResponse):
    alt_text: str | None
    asset_type: str
    height: int | None
    modified_at: str
    original_asset_id: str
    rendition_config_id: str | None
    status: Literal["Done", "InProgress", "Error"]
    width: int | None


class DropdownTypeObjectFieldUpdatePayload(BaseObjectFieldUpdatePayload):
    values: list[str] | None


class Choice6(APIObject):
    name: str


class DropdownTypeSettingsFieldCreatePayload(BaseSettingsFieldCreatePayload):
    choices: list[Choice6]
    is_multi_select: bool


class Choice7(APIObject):
    id: str
    name: str


class DropdownTypeSettingsFieldResponse(BaseSettingsFieldsResponse):
    choices: list[Choice7]
    is_multi_select: bool


class Error(APIObject):
    errors: dict[str, Any] | None
    message: str


class Links13(APIObject):
    campaign: str | None
    self: str


class EventResponse(APIObject):
    campaign_id: str | None
    created_by: str
    description: str | None
    end_date: str
    id: str
    is_all_day: bool
    is_archived: bool
    links: Links13
    reference_id: str
    start_date: str
    title: str


class EventUpdateRequest(APIObject):
    description: str | None
    end_date: str | None
    is_all_day: bool | None
    start_date: str | None
    title: str | None


class Source(APIObject):
    name: str | None


class FeaturedImageResponse(APIObject):
    attribution_text: str | None
    caption: str
    description: str | None
    height: int
    mime_type: str
    source: Source
    thumbnail: str | None
    url: str
    width: int


class FieldBaseType(APIObject):
    id: str
    name: str
    type: str
    values: list[str]


FieldType: TypeAlias = Literal[
    "boolean",
    "number",
    "text-field",
    "rich-text",
    "datetime",
    "library-asset",
    "content-type",
    "json",
    "url",
    "choice",
    "location",
]


class FieldTypeCheckbox(FieldBaseType):
    choices: list[Choice7]


class FieldTypeCommon(FieldBaseType):
    pass


class FieldTypeCurrencyNumber(FieldBaseType):
    currency_code: str
    decimal_places: int | None
    has_thousand_separator: bool


class FieldTypeDropdown(FieldBaseType):
    choices: list[Choice7]
    is_multi_select: bool


class FieldTypeLabel(FieldBaseType):
    choices: list[Choice7]
    is_multi_select: bool


class FieldTypePercentageNumber(FieldBaseType):
    decimal_places: int | None


class FieldTypeRadio(FieldBaseType):
    choices: list[Choice7]


class FieldTypeSimpleNumber(FieldBaseType):
    decimal_places: int | None
    has_thousand_separator: bool


class FileTypeFormFieldValueRequest(APIObject):
    key: str
    name: str


class FileTypeFormFieldValueResponse(APIObject):
    id: str
    name: str
    url: str


class FileUrlBulkCreateRequest(APIObject):
    guids: list[str]


class FolderCreateRequest(APIObject):
    name: str
    parent_folder_id: str | None


class FolderPermissionBulkCreateRequest(APIObject):
    permissions: list[Permission]
    type: Literal["user", "team"]


class FolderPermissionUpdateRequest(APIObject):
    access_type: Literal["view", "edit"] | None
    is_owner: bool | None


class Links14(APIObject):
    assets: str
    child_folders: str
    parent_folder: str | None
    self: str


class FolderResponse(APIObject):
    created_at: str
    id: str
    links: Links14
    modified_at: str
    name: str
    parent_folder_id: str | None
    path: str


class FolderUpdateRequest1(APIObject):
    name: str
    parent_folder_id: str | None


class FolderUpdateRequest2(APIObject):
    name: str | None
    parent_folder_id: str | None


FolderUpdateRequest: TypeAlias = FolderUpdateRequest1 | FolderUpdateRequest2


class GenericNumberTypeSettingsFieldCreatePayload(BaseSettingsFieldCreatePayload):
    decimal_places: int | None
    has_thousand_separator: bool | None


class GenericNumberTypeSettingsFieldResponse(BaseSettingsFieldsResponse):
    decimal_places: int
    has_thousand_separator: bool


class GenericNumberTypeSettingsFieldUpdatePayload(BaseSettingsFieldUpdatePayload):
    decimal_places: int | None
    has_thousand_separator: bool | None


class GenericSettingsFieldChoiceCreatePayload(APIObject):
    name: str


class GenericSettingsFieldChoiceUpdatePayload(APIObject):
    name: str


class HTTPException(APIObject):
    detail: str | None


class Instruction(APIObject):
    description: str
    id: str


class JSONFieldValueModel(APIObject):
    json_value: dict[str, Any] | list[Any]
    order_index: int | None


class KeyedPreviewCompletedModel(APIObject):
    completed: str
    locale: str | None
    mimeType: str | None
    name: str | None
    orderIndex: float | None


class KeyedPreviewErrorModel(APIObject):
    error: str
    locale: str | None
    mimeType: str | None
    name: str | None
    orderIndex: float | None


class LabelAndDropdownTypeSettingsFieldUpdatePayload(BaseSettingsFieldUpdatePayload):
    is_multi_select: bool | None


class LabelGroupValue(APIObject):
    id: str
    name: str


class LabelOnlyCustomField(APIObject):
    label: str


class LabelTypeSettingsFieldChoiceCreatePayload(
    GenericSettingsFieldChoiceCreatePayload
):
    color: str


class LabelTypeSettingsFieldChoiceUpdatePayload(APIObject):
    color: str | None
    name: str | None


class Choice12(APIObject):
    color: str
    name: str


class LabelTypeSettingsFieldCreatePayload(BaseSettingsFieldCreatePayload):
    choices: list[Choice12]
    is_multi_select: bool


class Choice13(APIObject):
    color: str | None
    id: str
    name: str


class LabelTypeSettingsFieldResponse(BaseSettingsFieldsResponse):
    choices: list[Choice13]
    is_multi_select: bool


class LibraryAssetCreateRequest(APIObject):
    folder_id: str | None
    key: str
    title: str


LibraryAssetType: TypeAlias = Literal[
    "article", "image", "video", "raw_file", "structured_content"
]


class Content(APIObject):
    type: Literal["url"]
    value: str


class Links15(APIObject):
    asset: str


class LibraryAssetVersionResponse(APIObject):
    asset_id: str
    content: Content
    created_at: str
    links: Links15
    mime_type: str
    title: str
    type: Literal["image", "video", "raw_file"]
    version_number: int


class FocalPoint(APIObject):
    x: int
    y: int


class ImageResolution(APIObject):
    height: int
    width: int


Tag1: TypeAlias = str


class Links16(APIObject):
    self: str


class Fields1(APIObject):
    has_embedded: bool | None


class ContentBody(APIObject):
    content_type_guid: str
    expired: bool | None
    expiry_datetime: str | None
    fields: Fields1 | None
    root_content: bool
    source: str | None
    source_id: str | None
    source_metadata: str | None
    title: str | None


class LibraryStructuredContentCreateRequest(APIObject):
    content_body: ContentBody
    folder_id: str | None
    title: str | None


class LocationDefaultValue(APIObject):
    latitude: float
    longitude: float


class LocationFieldValueModel(APIObject):
    latitude: float
    longitude: float
    order_index: int | None


class TargetField(APIObject):
    custom_field_type: str | None
    identifier: str
    type: str


class Action(APIObject):
    target_field: TargetField
    type: str
    values: list[str]


class Condition(APIObject):
    operator: str
    values: list[str]


class LogicRule(APIObject):
    action: Action
    condition: Condition


class LogicRuleAction(APIObject):
    target_field: TargetField
    type: str
    values: list[str]


class LogicRuleCondition(APIObject):
    operator: str
    values: list[str]


class LogicRuleTargetField(APIObject):
    custom_field_type: str | None
    identifier: str
    type: str


class Task(APIObject):
    id: str


class MilestoneCreateRequest(APIObject):
    campaign_id: str | None
    description: str | None
    due_date: str
    hex_color: str
    tasks: list[Task] | None
    title: str


class Campaign(APIObject):
    id: str


class MilestoneResponse(APIObject):
    campaign: Campaign | None
    color: str
    description: str | None
    due_date: str
    id: str
    links: Links16
    title: str


class MilestoneUpdateRequest(APIObject):
    campaign_id: str | None
    description: str | None
    due_date: str | None
    hex_color: str | None
    tasks: list[Task] | None
    title: str | None


class MultiChoiceTypeFormFieldRequest(BaseFormFieldRequest):
    values: list[str] | None


class Value1(APIObject):
    id: str
    name: str


class MultiChoiceTypeFormFieldResponse(BaseFormFieldResponse):
    values: list[Value1] | None


class MultiChoiceTypeObjectFieldUpdatePayload(BaseObjectFieldUpdatePayload):
    values: list[str] | None


class MultipleOptionCustomField(APIObject):
    choices: list[str]
    label: str


class NumberCustomField(APIObject):
    decimal_places: int | None
    has_thousand_separator: bool
    label: str


class NumberFieldValueModel(APIObject):
    num_value: float
    order_index: int | None


Value2: TypeAlias = float


class NumberTypeFormFieldRequest(BaseFormFieldRequest):
    values: list[Value2] | None


class NumberTypeFormFieldResponse(BaseFormFieldResponse):
    values: list[str] | None


class NumberTypeObjectFieldUpdatePayload(BaseObjectFieldUpdatePayload):
    values: list[float] | None


ObjectFieldCreateResponse: TypeAlias = (
    FieldTypeCommon
    | FieldTypeLabel
    | FieldTypeDropdown
    | FieldTypeRadio
    | FieldTypeCheckbox
    | FieldTypeSimpleNumber
    | FieldTypePercentageNumber
    | FieldTypeCurrencyNumber
)


class Pagination(APIObject):
    next: str | None
    previous: str | None


class PercentageCustomField(APIObject):
    decimal_places: int | None
    label: str


class Pagination2(Pagination):
    next: str | None


class PublishingChannelResponse(APIObject):
    disabled: bool | None
    id: str
    name: str | None


class PublishingMetadatum(APIObject):
    id: str
    links: Links16


class PublishingEventAssets(APIObject):
    id: str
    links: Links16
    publishing_metadata: list[PublishingMetadatum]
    type: str


class Error1(APIObject):
    asset_id: str
    error_code: Literal[
        "canonical-link-error",
        "unknown-asset",
        "metadata-exists",
        "duplicate-metadata",
        "invalid-status",
        "missing-public-url",
        "public-url-not-allowed",
        "domain-not-whitelisted",
        "missing-publishing-destination-updated-at",
        "publishing-destination-updated-at-not-allowed",
        "missing-locale",
        "locale-not-allowed",
    ]
    locale: str | None
    message: str


class PublishingEventMetadataCreateRequest(APIObject):
    asset_id: str
    locale: str | None
    public_url: str | None
    publishing_destination_updated_at: str | None
    status: Literal["published", "unpublished", "synced", "failed"]
    status_message: str | None


class Links20(APIObject):
    asset: str
    publishing_event: str
    self: str


class PublishingEventMetadataResponse(APIObject):
    asset_id: str
    asset_type: Literal["article", "image", "video", "raw_file", "structured_content"]
    id: str
    links: Links20
    locale: str | None
    public_url: str | None
    publishing_destination_updated_at: str | None
    status: Literal["published", "unpublished", "synced", "failed"] | None
    status_message: str | None


class Links21(APIObject):
    publishing_metadata: str
    self: str


class PublishingEventResponse(APIObject):
    assets: list[PublishingEventAssets]
    id: str
    links: Links21


class RadioButtonTypeFormFieldRequest(BaseFormFieldRequest):
    values: list[str] | None


class Value3(APIObject):
    id: str
    name: str


class RadioButtonTypeFormFieldResponse(MultiChoiceTypeFormFieldResponse):
    values: list[Value3] | None


class RadioButtonTypeObjectFieldUpdatePayload(BaseObjectFieldUpdatePayload):
    values: list[str] | None


class RecursiveStructuredContent(APIObject):
    content_type_guid: str
    expired: bool | None
    expiry_datetime: str | None
    root_content: bool
    source: str | None
    source_id: str | None
    source_metadata: str | None
    title: str | None


class Content1(APIObject):
    type: Literal["url", "api_url", "html_body"]
    value: str


class Links22(APIObject):
    self: str


class RelatedAssetItem(APIObject):
    content: Content1
    file_extension: str | None
    id: str
    links: Links22
    mime_type: str | None
    title: str
    type: Literal["article", "image", "video", "raw_file", "structured_content"]


class RelatedAssetsListResponse(APIObject):
    data: list[RelatedAssetItem]
    pagination: Pagination2


class RenditionConfigResponse(APIObject):
    asset_type: str
    height: int | None
    id: str
    input_format: str | None
    name: str
    output_format: str | None
    quality: float | None
    width: int | None


class RelatedAsset(APIObject):
    id: str


class ReplaceRelatedAssetsRequest(APIObject):
    related_assets: list[RelatedAsset]


class ReplaceRelatedAssetsResponse(APIObject):
    id: str
    related_assets: list[RelatedAsset]


class ResourceChoiceResponse(APIObject):
    id: str
    name: str


class ResourceLabelRequest(APIObject):
    group: str
    values: list[str]


class Group(APIObject):
    id: str
    name: str


class ResourceLabelResponse(APIObject):
    group: Group
    values: list[Value3]


class RichTextFieldValueModel(APIObject):
    order_index: int | None
    rich_text_value: str


class SCContentPreviewAcknowledgeRequest(APIObject):
    acknowledged_by: str
    content_hash: str


class SCContentPreviewCompleteRequest(APIObject):
    keyed_previews: dict[str, str | KeyedPreviewCompletedModel | KeyedPreviewErrorModel]


class SCContentTypeCreateResponse(APIObject):
    content_type_guid: str | None
    content_type_version_guid: str | None
    created: bool


class SCContentTypeManagedMigrationResponse(APIObject):
    content_migration_summary: ContentMigrationSummary | None
    content_type_id: str | None
    created_at: str | None
    default_values: dict[str, Any] | None
    id: str | None
    instance_id: str | None
    source_content_type_version_id: str | None
    status: Literal["not_started", "in_progress", "success", "error"] | None
    target_content_type_version_id: str | None
    updated_at: str | None


class SCContentTypeManagedMigrationStartResponse(APIObject):
    job_id: str | None
    started: bool | None


class SCContentTypeUpdateRequest(APIObject):
    details: CoreContentType
    source_metadata: str | None
    updated_by: str


class SCContentTypeUpdateResponse(APIObject):
    updated: bool


class Section(APIObject):
    id: str
    label: str


class CustomFields(APIObject):
    checkbox: list[MultipleOptionCustomField]
    currency: list[CurrencyCustomField]
    date: list[LabelOnlyCustomField]
    dropdown: list[MultipleOptionCustomField]
    image: list[LabelOnlyCustomField]
    multi_select_dropdown: list[MultipleOptionCustomField]
    multichoice: list[MultipleOptionCustomField]
    number: list[NumberCustomField]
    percentage: list[PercentageCustomField]
    richtext: list[LabelOnlyCustomField]
    string: list[LabelOnlyCustomField]
    textarea: list[LabelOnlyCustomField]
    video: list[LabelOnlyCustomField]


class SettingsApp(APIObject):
    authorization_callback_urls: list[str]
    description: str
    env: str
    expose_email: bool
    homepage_url: str
    name: str


class Create(APIObject):
    apps: list[str]
    custom_fields: list[str]
    labels: list[str]
    routing_rules: list[str]
    templates: list[str]
    webhooks: list[str]
    workflows: list[str]


class Update(APIObject):
    apps: list[str]
    custom_fields: list[str]
    labels: list[str]
    routing_rules: list[str]
    templates: list[str]
    webhooks: list[str]
    workflows: list[str]


class SettingsChangeset(APIObject):
    create: Create
    update: Update


class SettingsChangesetBase(APIObject):
    apps: list[str]
    custom_fields: list[str]
    labels: list[str]
    routing_rules: list[str]
    templates: list[str]
    webhooks: list[str]
    workflows: list[str]


class SettingsCustomField(APIObject):
    checkbox: list[MultipleOptionCustomField]
    currency: list[CurrencyCustomField]
    date: list[LabelOnlyCustomField]
    dropdown: list[MultipleOptionCustomField]
    image: list[LabelOnlyCustomField]
    multi_select_dropdown: list[MultipleOptionCustomField]
    multichoice: list[MultipleOptionCustomField]
    number: list[NumberCustomField]
    percentage: list[PercentageCustomField]
    richtext: list[LabelOnlyCustomField]
    string: list[LabelOnlyCustomField]
    textarea: list[LabelOnlyCustomField]
    video: list[LabelOnlyCustomField]


SettingsFieldChoiceCreateRequest: TypeAlias = list[
    GenericSettingsFieldChoiceCreatePayload | LabelTypeSettingsFieldChoiceCreatePayload
]


class SettingsFieldChoiceCreateResponseItem(APIObject):
    id: str
    name: str


SettingsFieldChoiceCreateResponse: TypeAlias = list[
    SettingsFieldChoiceCreateResponseItem
]
SettingsFieldChoiceUpdateRequest: TypeAlias = (
    GenericSettingsFieldChoiceUpdatePayload | LabelTypeSettingsFieldChoiceUpdatePayload
)


class SettingsLabelGroupLabel(APIObject):
    color: str | None
    name: str


class Rule(APIObject):
    custom_field_type: str | None
    field_type: str
    identifier: str | None
    operator: str
    unit: Literal["days", "weeks", "months", "years"] | None
    value: list[str]


class SettingsRoutingRule(APIObject):
    description: str
    name: str
    rules: list[Rule]
    template_name: str | None
    watchers: list[str]


class Changeset(APIObject):
    create: Create
    update: Update


class SettingsWebhook(APIObject):
    callback_url: str
    description: str
    event_names: list[str]
    name: str
    secret: str


class StringTypeObjectFieldUpdatePayload(BaseObjectFieldUpdatePayload):
    values: list[str] | None


class StructuredContentBody(APIObject):
    has_embedded: bool | None


class Tag(APIObject):
    guid: str
    name: str


class WebUrls(APIObject):
    drafts: str | None
    self: str
    task: str


class Links23(APIObject):
    drafts: str | None
    self: str
    task: str
    web_urls: WebUrls


class TaskArticle(APIObject):
    created_at: str
    html_body: str
    id: str
    labels: list[ResourceLabelResponse]
    library_asset_id: str | None
    links: Links23
    modified_at: str
    title: str
    url: str | None


class Links24(APIObject):
    asset: str
    task: str


class TaskAssetCommentResponse(APIObject):
    attachments: list[AttachmentResponse]
    created_at: str
    id: str
    is_resolved: bool
    links: Links24
    modified_at: str
    value: str


class Criterion1(APIObject):
    description: str | None
    id: str
    selected: bool | None


class Category(APIObject):
    criteria: list[Criterion1]
    id: str
    notes: str | None
    status: Literal["compliant", "not compliant", "not applicable"] | None


class TaskAssetDraftBrandComplianceRequest(APIObject):
    categories: list[Category]
    status: Literal["approved", "declined", "not_reviewed"]


class Criterion2(APIObject):
    description: str | None
    id: str
    name: str
    selected: bool | None


class Category1(APIObject):
    criteria: list[Criterion2]
    id: str
    name: str
    notes: str | None
    status: Literal["compliant", "not compliant", "not applicable"] | None


class TaskAssetDraftBrandComplianceResponse(APIObject):
    categories: list[Category1]
    reviewed_at: str | None
    reviewed_by: str | None
    status: Literal["approved", "declined", "not_reviewed"]


class Content2(APIObject):
    type: Literal["url"]
    value: str


class TaskAssetDraftListResponseItem(APIObject):
    content: Content2
    created_at: str
    draft_number: int
    id: str
    links: Links24
    mime_type: str
    title: str
    type: Literal["image", "video", "raw_file"]


class TaskAssetDraftResponse(APIObject):
    asset_id: str
    content: Content2
    created_at: str
    draft_number: int
    links: Links24
    mime_type: str
    title: str
    type: Literal["image", "video", "raw_file"]


class TaskAssetRequestForDirectUpload(APIObject):
    key: str
    title: str
    type: str | None


class TaskAssetRequestForLibrary(APIObject):
    content_id: str
    content_type: Literal["article", "image", "raw_file", "structured_content", "video"]
    type: str


class Links27(APIObject):
    drafts: str | None
    self: str
    task: str
    web_urls: WebUrls


class TaskAssetResponse(APIObject):
    content: AssetContent
    created_at: str
    id: str
    labels: list[ResourceLabelResponse]
    links: Links27
    mime_type: str
    modified_at: str
    title: str
    type: Literal["article", "image", "video", "raw_file", "structured_content"]


class Field1(APIObject):
    name: str
    settings_field_id: str | None
    value: list[str]


class Links28(APIObject):
    self: str
    task: str


class TaskBriefTemplateValue(APIObject):
    id: str
    name: str


class Links29(APIObject):
    comment_by: str
    task: str


class TaskCommentResponse(APIObject):
    attachments: list[AttachmentResponse]
    comment_by: str
    created_at: str
    id: str
    is_resolved: bool
    links: Links29
    modified_at: str
    parent_comment_id: str | None
    value: str


class TaskCreateRequest(APIObject):
    campaign_id: str | None
    due_at: str | None
    owner_id: str | None
    start_at: str | None
    title: str
    workflow_id: str | None


class Links30(APIObject):
    choices: str | None
    self: str


class Value5(APIObject):
    id: str | None
    name: str


class TaskCustomField(APIObject):
    id: str
    links: Links30
    name: str
    type: Literal[
        "text_field",
        "multi_line_text_field",
        "checkboxes",
        "dropdown",
        "multi_select_dropdown",
        "multiple_choice",
        "date_field",
        "image",
        "video",
        "rich_text_field",
    ]
    values: list[Value5]


class TaskCustomFieldChoiceListResponseItem(APIObject):
    id: str
    name: str


class TaskCustomFieldUpdateRequest(APIObject):
    values: list[str]


class TaskExternalWorkRequest(APIObject):
    identifier: str | None
    status: str | None
    title: str | None
    url: str | None


class Links31(APIObject):
    self: str


class TaskExternalWorkResponse(APIObject):
    external_system: str
    identifier: str | None
    links: Links31
    status: str | None
    title: str | None
    url: str | None


class WebUrls2(APIObject):
    drafts: str
    self: str
    task: str


class Links32(APIObject):
    drafts: str
    self: str
    task: str
    web_urls: WebUrls2


class TaskImage(APIObject):
    created_at: str
    file_size: int
    id: str
    image_resolution: ImageResolution
    labels: list[ResourceLabelResponse]
    library_asset_id: str | None
    links: Links32
    mime_type: str
    modified_at: str
    title: str
    url: str


class WebUrls3(APIObject):
    self: str


class Links33(APIObject):
    campaign: str | None
    milestone: str | None
    self: str
    web_urls: WebUrls3
    workflow: str | None


class TaskListResponseItem(APIObject):
    campaign_id: str | None
    due_at: str | None
    id: str
    is_archived: bool
    is_completed: bool
    links: Links33
    milestone_id: str | None
    modified_at: str | None
    reference_id: str
    start_at: str | None
    status: Literal[
        "Archived", "Completed", "Overdue", "Not Started", "In Progress", "On Hold"
    ]
    title: str
    workflow_id: str | None


class TaskPublishingIntentCreateRequest(APIObject):
    channel_id: str


class TaskPublishingIntentResponse(APIObject):
    id: str


class WebUrls4(APIObject):
    drafts: str
    self: str
    task: str


class Links34(APIObject):
    drafts: str
    self: str
    task: str
    web_urls: WebUrls4


class TaskRawFile(APIObject):
    created_at: str
    file_size: int
    id: str
    labels: list[ResourceLabelResponse]
    library_asset_id: str | None
    links: Links34
    mime_type: str
    modified_at: str
    title: str
    url: str


class WebUrls5(APIObject):
    brief: str
    self: str


class Links35(APIObject):
    assets: str
    attachments: str
    brief: str | None
    campaign: str
    custom_fields: str | None
    milestone: str | None
    self: str
    web_urls: WebUrls5


class TaskStepRequest(APIObject):
    due_at: str | None


class TaskStructuredContentCreateRequest(APIObject):
    content_type_guid: str
    expired: bool | None
    expiry_datetime: str | None
    fields: Fields1 | None
    primary_locale: str | None
    root_content: bool | None
    source: str | None
    source_id: str | None
    source_metadata: str | None
    template_guid: str | None
    title: str | None


class Links36(APIObject):
    assignee: str | None
    external_work: str | None
    self: str
    task: str


class TaskSubStep(APIObject):
    assignee_id: str | None
    assignee_type: Literal["user", "team"]
    id: str
    is_completed: bool
    is_external: bool
    is_in_progress: bool
    is_skipped: bool
    links: Links36
    title: str


class Links37(APIObject):
    comment_by: str
    self: str
    sub_step: str
    task: str


class TaskSubStepCommentResponse(APIObject):
    attachments: list[AttachmentResponse]
    created_at: str
    id: str
    is_resolved: bool
    links: Links37 | None
    modified_at: str
    value: str


class TaskSubStepCommentUpdateRequest(APIObject):
    value: str


class TaskSubStepRequest(APIObject):
    assignee_id: str | None
    assignee_type: Literal["user", "team"] | None
    is_completed: Literal[True] | None
    is_in_progress: Literal[True] | None
    is_skipped: Literal[True] | None


class TaskUpdateRequest(APIObject):
    campaign_id: str | None
    due_at: str | None
    labels: list[ResourceLabelRequest] | None
    owner_id: str | None
    start_at: str | None
    title: str | None
    workflow_id: str | None


class TaskUrlResponse(APIObject):
    created_at: str
    id: str


class WebUrls6(APIObject):
    drafts: str
    self: str
    task: str


class Links38(APIObject):
    drafts: str
    self: str
    task: str
    web_urls: WebUrls6


class TaskVideo(APIObject):
    created_at: str
    file_size: int
    id: str
    labels: list[ResourceLabelResponse]
    library_asset_id: str | None
    links: Links38
    mime_type: str
    modified_at: str
    title: str
    url: str


class Links39(APIObject):
    self: str


class Team(APIObject):
    id: str
    links: Links39
    name: str


class User(APIObject):
    email: str
    id: str
    name: str


class TeamWithUsers(APIObject):
    id: str
    links: Links39
    name: str
    users: list[User]


class TypeSpecificMeta(APIObject):
    choices: list[ResourceChoiceResponse]
    is_multi_select: bool


class TypeSpecificMeta1(APIObject):
    currency_code: str
    decimal_places: int | None
    has_thousand_separator: bool


class TypeSpecificMeta2(APIObject):
    description: str


class Datum(BaseTemplateResponse):
    links: Links39


class TemplateListResponse(APIObject):
    data: list[Datum]
    pagination: Pagination2


class TargetField2(APIObject):
    identifier: str


class TemplateLogicRuleAction(APIObject):
    target_field: TargetField2
    type: Literal["jump_to", "show_values"]
    values: list[str]


class TemplateLogicRuleCondition(APIObject):
    operator: Literal["any_of", "equal", "not_equal"]
    values: list[str]


class TypeSpecificMeta3(APIObject):
    decimal_places: int | None


class TypeSpecificMeta4(APIObject):
    decimal_places: int | None
    has_thousand_separator: bool


class TextFieldValueModel(APIObject):
    order_index: int | None
    text_value: str


class TextTypeFormFieldRequest(BaseFormFieldRequest):
    values: list[str] | None


class TextTypeFormFieldResponse(BaseFormFieldResponse):
    values: list[str] | None


class URLFieldValueModel(APIObject):
    order_index: int | None
    url: str


class Role(APIObject):
    name: str


class UserListResponseItem(APIObject):
    first_name: str
    full_name: str
    id: str
    image_url: str | None
    last_name: str
    links: Links39
    roles: list[Role]


class Links45(APIObject):
    self: str | None


class UserResponse(APIObject):
    email: str | None
    first_name: str
    full_name: str
    id: str
    image_url: str | None
    last_name: str
    links: Links45 | None


class ValidationError(APIObject):
    loc: list[str]
    msg: str
    type: str


class Links46(APIObject):
    self: str | None
    versions: str | None


class ContentType(APIObject):
    type: Literal["html_body", "url"]
    value: str


class WorkRequestApprovedAssetResponse(APIObject):
    content_type: ContentType
    id: str
    mime_type: str
    title: str
    type: Literal["article", "image", "video", "raw_file"]


class Value6(APIObject):
    key: str
    name: str


class WorkRequestAttachmentBriefRequest(APIObject):
    type: Literal["attachment_brief"]
    value: Value6


class Value7(APIObject):
    id: str
    name: str
    url: str


class WorkRequestAttachmentBriefResponse(APIObject):
    type: Literal["attachment_brief"]
    value: Value7


class WorkRequestCampaignRequest(APIObject):
    description: str | None
    end_date: str | None
    owner_id: str | None
    parent_campaign_id: str | None
    start_date: str | None
    title: str


class Links47(APIObject):
    brief: str | None
    parent_campaign: str | None
    self: str


class WorkRequestCampaignResponse(APIObject):
    description: str | None
    end_date: str | None
    id: str
    links: Links47
    owner_id: str
    parent_campaign_id: str
    reference_id: str
    start_date: str | None
    title: str


class Links48(APIObject):
    comment_by: str
    work_request: str


class WorkRequestCommentResponse(APIObject):
    attachments: list[AttachmentResponse]
    comment_by: str
    created_at: str
    id: str
    is_resolved: bool
    links: Links48
    modified_at: str
    parent_comment_id: str | None
    value: str


class Links49(APIObject):
    self: str | None


class WorkRequestRelatedResourceResponse(APIObject):
    created_at: str
    id: str
    links: Links49
    relation_type: Literal["started", "linked"]
    resource_type: str


class WorkRequestRequestFormFieldCheckboxTypePayload(APIObject):
    type: Literal["checkbox"]
    values: list[str]


class Value8(APIObject):
    id: str
    name: str


class WorkRequestRequestFormFieldCheckboxTypeResponse(APIObject):
    type: Literal["checkbox"]
    values: list[Value8]


class WorkRequestRequestFormFieldCurrencyNumberTypePayload(APIObject):
    type: Literal["currency_number"]
    values: list[str]


class WorkRequestRequestFormFieldDateTypePayload(APIObject):
    type: Literal["date"]
    values: list[str]


class WorkRequestRequestFormFieldDropdownTypePayload(APIObject):
    type: Literal["dropdown"]
    values: list[str]


class WorkRequestRequestFormFieldDropdownTypeResponse(APIObject):
    type: Literal["dropdown"]
    values: list[Value8]


class Value10(APIObject):
    key: str
    name: str


class WorkRequestRequestFormFieldFileTypePayload(APIObject):
    type: Literal["file"]
    values: list[Value10]


class Value11(APIObject):
    id: str
    name: str
    url: str


class WorkRequestRequestFormFieldFileTypeResponse(APIObject):
    type: Literal["file"]
    values: list[Value11]


class WorkRequestRequestFormFieldLabelTypePayload(APIObject):
    type: Literal["label"]
    values: list[str]


class Value12(APIObject):
    id: str
    name: str


class WorkRequestRequestFormFieldLabelTypeResponse(APIObject):
    type: Literal["label"]
    values: list[Value12]


class WorkRequestRequestFormFieldPercentageNumberTypePayload(APIObject):
    type: Literal["percentage_number"]
    values: list[str]


class WorkRequestRequestFormFieldRadioButtonTypePayload(APIObject):
    type: Literal["radio_button"]
    values: list[str]


class Value13(APIObject):
    id: str | None
    name: str | None


class WorkRequestRequestFormFieldRadioButtonTypeResponse(APIObject):
    type: Literal["radio_button"]
    values: list[Value13]


class WorkRequestRequestFormFieldRichtextTypePayload(APIObject):
    type: Literal["richtext"]
    values: list[str]


class WorkRequestRequestFormFieldSimpleNumberTypePayload(APIObject):
    type: Literal["simple_number"]
    values: list[str]


class WorkRequestRequestFormFieldTextAreaTypePayload(APIObject):
    type: Literal["text_area"]
    values: list[str]


class WorkRequestRequestFormFieldTextTypePayload(APIObject):
    type: Literal["text"]
    values: list[str]


class Links50(APIObject):
    self: str


class Assignee(APIObject):
    full_name: str
    id: str
    image_url: str | None
    links: Links50
    type: Literal["user", "team"]


class CreatedBy(APIObject):
    id: str
    links: Links50


class Links52(APIObject):
    attachments: str
    campaigns: str
    comments: str
    creative_assets: str
    related_resources: str
    self: str
    tasks: str


class Links53(APIObject):
    self: str


class Template(APIObject):
    id: str
    links: Links53
    title: str


class WorkRequestTaskRequest(APIObject):
    asset_ids_without_field_inheritance: list[str] | None
    campaign_id: str | None
    due_at: str | None
    inherit_fields_from: str | None
    owner_id: str | None
    start_at: str | None
    title: str | None
    workflow_id: str | None


class Links54(APIObject):
    owner: str | None
    self: str


class WorkRequestTaskResponse(APIObject):
    due_at: str | None
    id: str
    links: Links54
    owner_id: str
    reference_id: str | None
    start_at: str | None
    title: str


class WorkRequestTextBriefRequest(APIObject):
    type: Literal["text_brief"]
    value: str


class WorkRequestTextBriefResponse(APIObject):
    type: Literal["text_brief"]
    value: str


class WorkRequestUpdateRequest(APIObject):
    assignees: list[str] | None
    priority: Literal["Low", "Medium", "High"] | None
    status: Literal["Accepted", "Declined", "Submitted", "Completed"] | None


class WorkflowChannel(APIObject):
    display_name: str
    type: str


class WorkflowCustomfield(APIObject):
    default_value: list[str]
    field: str
    is_required: bool
    type: str


class WorkflowLabel(APIObject):
    default_value: list[str]
    is_required: bool
    is_required_at_task_creation: bool
    label_type: str


class WorkflowListResponseItem(APIObject):
    description: str | None
    id: str
    is_flexible: bool
    name: str


class Action1(APIObject):
    name: str


class Links55(APIObject):
    self: str


class Assignee1(APIObject):
    id: str
    links: Links55
    name: str
    type: Literal["user", "team"]


class SubStep(APIObject):
    actions: list[Action1]
    assignees: list[Assignee1]
    id: str
    name: str
    type: Literal["default", "external"]


class Step(APIObject):
    id: str
    name: str
    sub_steps: list[SubStep]


class WorkflowResponse(APIObject):
    description: str | None
    fields: list[ObjectFieldCreateResponse]
    id: str
    is_active: bool
    is_asset_approval_enabled: bool
    is_flexible: bool
    is_resource_management_enabled: bool
    is_smart_duration_enabled: bool
    name: str
    steps: list[Step]


class ExternalSubStepConfig(APIObject):
    is_user_interaction_allowed: bool | None


class WorkflowSubStep(APIObject):
    actions: list[str]
    description: str | None
    external_sub_step_config: ExternalSubStepConfig | None
    external_system: str | None
    is_external: bool | None
    label: str | None


ErrorReason: TypeAlias = Literal[
    "REQUIRED_FIELD_ABSENT",
    "LIST_EMPTY",
    "INVALID_STATE",
    "INVALID_SELF_REF",
    "MIN_NOT_MET",
    "MAX_NOT_MET",
    "PATTERN_ERROR",
]


class ListAssetLineagesResponse(APIObject):
    data: list[AssetLineageResponse]
    pagination: Pagination2


class Links56(APIObject):
    asset_fields: str
    asset_url: str


class ListAssetRenditionsResponse(APIObject):
    data: list[BaseAssetRenditionResponse]
    pagination: Pagination2


class ListBrandComplianceCategoriesResponse(APIObject):
    data: list[BrandComplianceCategoriesResponse]
    pagination: Pagination2


class ListCampaignsResponse(APIObject):
    data: list[CampaignListResponseItem]
    pagination: Pagination2


class ListEventsResponse(APIObject):
    data: list[EventResponse]
    pagination: Pagination2


class CreateFieldResponse(APIObject):
    id: str


class ListFoldersResponse(APIObject):
    data: list[FolderResponse]
    pagination: Pagination2


class ListMilestonesResponse(APIObject):
    data: list[MilestoneResponse]
    pagination: Pagination2


class Links57(APIObject):
    complete: str
    status: str


class CreateMultipartUploadResponse(APIObject):
    expires_at: str
    id: str
    links: Links57
    upload_part_count: int
    upload_part_urls: list[str]


class Links58(APIObject):
    status: str


class CompleteMultipartUploadResponse(APIObject):
    key: str
    links: Links58


class Links59(APIObject):
    self: str


class GetMultipartUploadStatusResponse(APIObject):
    expires_at: str
    id: str
    key: str
    links: Links59
    status: Literal[
        "UPLOAD_COMPLETION_NOT_STARTED",
        "UPLOAD_COMPLETION_IN_PROGRESS",
        "UPLOAD_COMPLETION_SUCCEEDED",
        "UPLOAD_COMPLETION_FAILED",
    ]
    status_message: str | None


ListSCContentTypesResponse: TypeAlias = list[BaseContentTypeModel]
ListSCContentTypeManagedMigrationsResponse: TypeAlias = list[
    SCContentTypeManagedMigrationResponse
]


class CreateSCContentTypeManagedMigrationResponse(APIObject):
    created: bool


class ValidateSCContentTypeManagedMigrationResponse(APIObject):
    is_managed_migration_possible: bool


class UpdateSCContentTypeManagedMigrationResponse(APIObject):
    updated: bool | None


ListSCContentTypeVersionsResponse: TypeAlias = list[BaseContentTypeVersionModel]


class MigrateSCContentResponse(APIObject):
    content_guid: str | None
    content_hash: str | None
    created: bool | None
    version_guid: str | None


class ListTasksResponse(APIObject):
    data: list[TaskListResponseItem]
    pagination: Pagination2


class Datum2(TaskAssetResponse):
    library_asset_id: str | None


class ListTaskAssetsResponse(APIObject):
    data: list[Datum2]
    pagination: Pagination2


class ListTaskAttachmentsResponse(APIObject):
    data: list[AttachmentResponse]
    pagination: Pagination2


class ListTaskCustomFieldsResponse(APIObject):
    data: list[TaskCustomField]
    pagination: Pagination2


class ListTaskAssetCommentsResponse(APIObject):
    data: list[TaskAssetCommentResponse] | None
    pagination: Pagination2 | None


class ListTaskAssetDraftsResponse(APIObject):
    data: list[TaskAssetDraftListResponseItem]
    pagination: Pagination2


class ListTaskCustomFieldChoicesResponse(APIObject):
    data: list[TaskCustomFieldChoiceListResponseItem]
    pagination: Pagination2


class ListTaskSubStepCommentsResponse(APIObject):
    data: list[TaskSubStepCommentResponse]
    pagination: Pagination2


Data1: TypeAlias = (
    FieldTypeLabel | FieldTypeDropdown | FieldTypeRadio | FieldTypeCheckbox
)


class ListTaskSubStepFieldsResponse(APIObject):
    data: list[Data1]
    pagination: Pagination2


class CreateTaskStructuredContentDraftResponse(APIObject):
    pass


class ListTeamsResponse(APIObject):
    data: list[Team]
    pagination: Pagination2


UploadMetaFields = TypedDict(
    "UploadMetaFields",
    {
        "key": str,
        "policy": str,
        "x-amz-algorithm": str,
        "x-amz-credential": str,
        "x-amz-date": str,
        "x-amz-security-token": NotRequired[str],
        "x-amz-signature": str,
    },
)


class GetUploadUrlResponse(APIObject):
    upload_meta_fields: UploadMetaFields
    url: str


class ListWorkRequestApprovedAssetsResponse(APIObject):
    data: list[WorkRequestApprovedAssetResponse]
    pagination: Pagination2


class ListWorkRequestCommentsResponse(APIObject):
    data: list[WorkRequestCommentResponse]
    pagination: Pagination2


class ListWorkRequestRelatedResourcesResponse(APIObject):
    data: list[WorkRequestRelatedResourceResponse]
    pagination: Pagination2


class ListWorkflowsResponse(APIObject):
    data: list[WorkflowListResponseItem]
    pagination: Pagination2


class AssetFieldTypeCheckbox(BaseAssetField):
    choices: list[Choice]


class AssetFieldTypeCommon(BaseAssetField):
    pass


class AssetFieldTypeCurrencyNumber(BaseAssetField):
    currency_code: str
    decimal_places: int | None
    has_thousand_separator: bool


class AssetFieldTypeDropdown(BaseAssetField):
    choices: list[Choice]
    is_multi_select: bool


class AssetFieldTypeLabel(BaseAssetField):
    choices: list[Choice]
    is_multi_select: bool


class AssetFieldTypePercentageNumber(BaseAssetField):
    decimal_places: int | None


class AssetFieldTypeRadio(BaseAssetField):
    choices: list[Choice]


class AssetFieldTypeSimpleNumber(BaseAssetField):
    decimal_places: int | None
    has_thousand_separator: bool


class AssetPermissionListResponseItem(APIObject):
    data: list[BasePermissionsResponseSchema]
    pagination: Pagination2


class AssetResponse(APIObject):
    content: AssetContent
    created_at: str
    file_extension: str | None
    file_location: str
    folder_id: str | None
    id: str
    is_archived: bool
    labels: list[ResourceLabelResponse]
    links: Links2
    mime_type: str
    modified_at: str
    owner_organization_id: str
    thumbnail_url: str | None
    title: str
    type: Literal["article", "image", "video", "raw_file", "structured_content"]


class AssetTypeObjectFieldUpdatePayload(BaseObjectFieldUpdatePayload):
    values: list[Value] | None


class BaseObjectField(BaseObjectFieldUpdatePayload):
    id: str


class BriefTypeFormFieldRequest(BaseFormFieldRequest):
    values: list[WorkRequestTextBriefRequest | WorkRequestAttachmentBriefRequest] | None


class BriefTypeFormFieldResponse(BaseFormFieldResponse):
    values: (
        list[WorkRequestTextBriefResponse | WorkRequestAttachmentBriefResponse] | None
    )


class CampaignBriefResponse(APIObject):
    fields: list[Field]
    links: Links7
    template: CampaignBriefTemplateValue | None
    title: str
    type: Literal["template", "attachment", "text"]


CampaignFieldUpdateRequest: TypeAlias = (
    StringTypeObjectFieldUpdatePayload
    | MultiChoiceTypeObjectFieldUpdatePayload
    | DropdownTypeObjectFieldUpdatePayload
    | RadioButtonTypeObjectFieldUpdatePayload
    | NumberTypeObjectFieldUpdatePayload
    | DateTypeObjectFieldUpdatePayload
    | AssetTypeObjectFieldUpdatePayload
)


class CampaignResponse(APIObject):
    budget: BudgetResponse | None
    created_at: str
    description: str | None
    end_date: str | None
    id: str
    is_hidden: bool
    labels: list[ResourceLabelResponse]
    links: Links11
    reference_id: str
    start_date: str | None
    status: Literal["Not Started", "Off Track", "On Track", "Complete", "At Risk"]
    title: str


class ContentFieldValueWithEmbeddedModel(APIObject):
    content_details: RecursiveStructuredContent


class ContentFieldsVersionValidation(APIObject):
    fields: dict[str, dict[str, ErrorReason]] | None


class CoreFieldDef(APIObject):
    editor_metadata: dict[str, Any] | list[Any] | None
    field_type: FieldType | None
    help_text: str | None
    is_list: bool
    is_required: bool
    key: str
    max_list_length: int | None
    min_list_length: int | None
    name: str
    need_internationalization: bool
    order_index: int | None
    source_id: str | None
    source_metadata: str | None


class CurrencyNumberTypeSettingsFieldCreatePayload(
    GenericNumberTypeSettingsFieldCreatePayload
):
    currency_code: str | None


class CurrencyNumberTypeSettingsFieldResponse(GenericNumberTypeSettingsFieldResponse):
    currency_code: str


class CurrencyNumberTypeSettingsFieldUpdatePayload(
    GenericNumberTypeSettingsFieldUpdatePayload
):
    currency_code: str | None


class DateTypeObjectField(BaseObjectField, DateTypeObjectFieldUpdatePayload):
    pass


class DatetimeFieldDefinition(APIObject):
    core: CoreFieldDef
    default_values: list[str] | None
    max_date: str
    min_date: str


class DropdownTypeObjectField(BaseObjectField, DropdownTypeObjectFieldUpdatePayload):
    pass


FieldListResponseItem: TypeAlias = (
    FieldTypeCommon
    | FieldTypeLabel
    | FieldTypeDropdown
    | FieldTypeRadio
    | FieldTypeCheckbox
    | FieldTypeSimpleNumber
    | FieldTypePercentageNumber
    | FieldTypeCurrencyNumber
)


class FileTypeFormFieldRequest(BaseFormFieldRequest):
    values: list[FileTypeFormFieldValueRequest] | None


class FileTypeFormFieldResponse(BaseFormFieldResponse):
    values: list[FileTypeFormFieldValueResponse] | None


class FolderPermissionListResponseItem(APIObject):
    data: list[BasePermissionsResponseSchema]
    pagination: Pagination2


class FormField(APIObject):
    custom_field_type: str | None
    field_type: str
    help: str
    identifier: str
    is_required: bool
    logic_rules: list[LogicRule]
    sort_order: int


class HTTPValidationError(APIObject):
    detail: list[ValidationError] | None


class LabelGroup(APIObject):
    id: str
    name: str
    source_org_type: Literal["current", "related"]
    values: list[LabelGroupValue]


class LibraryArticle(APIObject):
    authors: list[ArticleAuthorResponse]
    created_at: str
    expires_at: str | None
    file_location: str
    folder_id: str | None
    group_id: str | None
    html_body: str
    id: str
    images: list[FeaturedImageResponse]
    is_archived: bool
    labels: list[ResourceLabelResponse]
    lang_code: str | None
    meta_description: str | None
    meta_keywords: list[str]
    meta_title: str | None
    meta_url: str | None
    modified_at: str
    owner_id: str
    owner_organization_id: str
    pixel_key: str
    source_article: str | None
    source_name: str | None
    tags: list[Tag]
    thumbnail_url: str | None
    title: str
    url: str | None
    version_id: str
    version_number: int


class LibraryAssetDefaultValue(APIObject):
    asset_guid: str
    asset_type: LibraryAssetType


class LibraryAssetFieldDefinition(APIObject):
    allowed_types: list[LibraryAssetType]
    core: CoreFieldDef
    default_values: list[LibraryAssetDefaultValue] | None


class LibraryAssetFieldValueModel(APIObject):
    asset_guid: str
    asset_type: LibraryAssetType
    links: dict[str, str] | None
    order_index: int | None


class LibraryImage(APIObject):
    allow_se_indexing: bool
    alt_text: str | None
    attribution_text: str | None
    created_at: str
    description: str | None
    expires_at: str | None
    file_extension: str | None
    file_location: str
    file_size: int
    focal_point: FocalPoint | None
    folder_id: str | None
    id: str
    image_resolution: ImageResolution
    is_archived: bool
    is_public: bool
    labels: list[ResourceLabelResponse]
    mime_type: str
    modified_at: str
    owner_id: str
    owner_organization_id: str
    tags: list[Tag]
    thumbnail_url: str | None
    title: str
    url: str
    version_id: str
    version_number: int


class LibraryImageUpdateRequest(APIObject):
    allow_se_indexing: bool | None
    alt_text: str | None
    attribution_text: str | None
    description: str | None
    expires_at: str | None
    folder_id: str | None
    is_archived: bool | None
    is_public: bool | None
    labels: list[ResourceLabelRequest] | None
    tags: list[Tag1] | None
    title: str | None


class LibraryRawFile(APIObject):
    allow_se_indexing: bool
    attribution_text: str | None
    created_at: str
    description: str | None
    expires_at: str | None
    file_extension: str | None
    file_location: str
    file_size: int
    folder_id: str | None
    id: str
    is_archived: bool
    is_public: bool
    labels: list[ResourceLabelResponse]
    mime_type: str
    modified_at: str
    owner_id: str
    owner_organization_id: str
    tags: list[Tag]
    thumbnail_url: str | None
    title: str
    url: str
    version_id: str
    version_number: int


class LibraryRawFileUpdateRequest(APIObject):
    allow_se_indexing: bool | None
    attribution_text: str | None
    description: str | None
    expires_at: str | None
    folder_id: str | None
    is_archived: bool | None
    is_public: bool | None
    labels: list[ResourceLabelRequest] | None
    tags: list[Tag1] | None
    title: str | None


class LibraryVideo(APIObject):
    allow_se_indexing: bool
    alt_text: str | None
    attribution_text: str | None
    created_at: str
    description: str | None
    expires_at: str | None
    file_extension: str | None
    file_location: str
    file_size: int
    folder_id: str | None
    id: str
    is_archived: bool
    is_public: bool
    labels: list[ResourceLabelResponse]
    mime_type: str
    modified_at: str
    owner_id: str
    owner_organization_id: str
    tags: list[Tag]
    thumbnail_url: str | None
    title: str
    url: str
    version_id: str
    version_number: int


class LibraryVideoUpdateRequest(APIObject):
    allow_se_indexing: bool | None
    alt_text: str | None
    attribution_text: str | None
    description: str | None
    expires_at: str | None
    folder_id: str | None
    is_archived: bool | None
    is_public: bool | None
    labels: list[ResourceLabelRequest] | None
    tags: list[Tag1] | None
    title: str | None


class LocalizedFieldValuesWithEmbeddedContent(APIObject):
    field_values: list[
        NumberFieldValueModel
        | BooleanFieldValueModel
        | DatetimeFieldValueModel
        | TextFieldValueModel
        | RichTextFieldValueModel
        | LibraryAssetFieldValueModel
        | ContentFieldValueWithEmbeddedModel
        | ContentFieldValueModel
        | URLFieldValueModel
        | JSONFieldValueModel
        | ChoiceFieldValueModel
        | LocationFieldValueModel
    ]
    locale: str


class MultiChoiceTypeObjectField(
    BaseObjectField, MultiChoiceTypeObjectFieldUpdatePayload
):
    pass


class NumberFieldDefinition(APIObject):
    core: CoreFieldDef
    default_values: list[float] | None
    max_value: float | None
    min_value: float | None


class NumberTypeObjectField(BaseObjectField, NumberTypeObjectFieldUpdatePayload):
    pass


class PublishingChannelListResponse(APIObject):
    data: list[PublishingChannelResponse]
    pagination: Pagination2


class PublishingEventMetadataBulkCreateRequest(APIObject):
    data: list[PublishingEventMetadataCreateRequest]


class PublishingEventMetadataBulkCreateResponse(APIObject):
    data: list[PublishingEventMetadataResponse]
    errors: list[Error1]


class PublishingEventMetadataListResponse(APIObject):
    data: list[PublishingEventMetadataResponse] | None


class RadioButtonTypeObjectField(
    BaseObjectField, RadioButtonTypeObjectFieldUpdatePayload
):
    pass


class RichTextFieldDefinition(APIObject):
    core: CoreFieldDef
    default_values: list[str] | None
    max_visual_text_length: int | None
    min_visual_text_length: int


SettingsFieldCreateRequest: TypeAlias = (
    BaseSettingsFieldCreatePayload
    | LabelTypeSettingsFieldCreatePayload
    | DropdownTypeSettingsFieldCreatePayload
    | CheckboxAndRadioTypeSettingsFieldCreatePayload
    | GenericNumberTypeSettingsFieldCreatePayload
    | CurrencyNumberTypeSettingsFieldCreatePayload
)
SettingsFieldUpdateRequest: TypeAlias = (
    BaseSettingsFieldUpdatePayload
    | LabelAndDropdownTypeSettingsFieldUpdatePayload
    | GenericNumberTypeSettingsFieldUpdatePayload
    | CurrencyNumberTypeSettingsFieldUpdatePayload
)


class SettingsLabelGroup(APIObject):
    has_single_value: bool
    label_type: str
    labels: list[SettingsLabelGroupLabel]


class SettingsTemplate(APIObject):
    description: str
    form_fields: list[FormField]
    instructions: list[Instruction]
    name: str
    public_description: str
    sections: list[Section]
    types: list[str]


class StringTypeObjectField(BaseObjectField, StringTypeObjectFieldUpdatePayload):
    pass


TaskAssetRequest: TypeAlias = (
    TaskAssetRequestForDirectUpload | TaskAssetRequestForLibrary
)


class TaskBriefResponse(APIObject):
    fields: list[Field1]
    links: Links28
    template: TaskBriefTemplateValue | None
    title: str
    type: Literal["template", "attachment", "text"]


TaskFieldUpdateRequest: TypeAlias = (
    StringTypeObjectFieldUpdatePayload
    | MultiChoiceTypeObjectFieldUpdatePayload
    | DropdownTypeObjectFieldUpdatePayload
    | RadioButtonTypeObjectFieldUpdatePayload
    | NumberTypeObjectFieldUpdatePayload
    | DateTypeObjectFieldUpdatePayload
    | AssetTypeObjectFieldUpdatePayload
)


class TaskStep(APIObject):
    description: str | None
    due_at: str | None
    id: str
    is_completed: bool
    sub_steps: list[TaskSubStep]
    title: str


class TemplateLogicRuleItem(APIObject):
    action: TemplateLogicRuleAction
    condition: TemplateLogicRuleCondition


TemplateLogicRule: TypeAlias = list[TemplateLogicRuleItem]


class TemplatePercentageNumberFormField(BaseTemplateFormField):
    logic_rules: TemplateLogicRule
    type_specific_meta: TypeSpecificMeta3


class TemplateSimpleNumberFormField(BaseTemplateFormField):
    logic_rules: TemplateLogicRule
    type_specific_meta: TypeSpecificMeta4


class TextFieldDefinition(APIObject):
    core: CoreFieldDef
    default_values: list[str] | None
    max_length: int | None
    min_length: int
    validation_pattern: str


class Datum1(UserListResponseItem):
    links: Links39


class UserListResponse(APIObject):
    data: list[Datum1]
    pagination: Pagination2


FormFields1: TypeAlias = (
    BriefTypeFormFieldRequest
    | DateTypeFormFieldRequest
    | FileTypeFormFieldRequest
    | MultiChoiceTypeFormFieldRequest
    | NumberTypeFormFieldRequest
    | RadioButtonTypeFormFieldRequest
    | TextTypeFormFieldRequest
)


class WorkRequestCreateRequest(APIObject):
    assignees: list[str] | None
    form_fields: list[FormFields1]
    template_id: str


Values: TypeAlias = WorkRequestTextBriefRequest | WorkRequestAttachmentBriefRequest


class WorkRequestRequestFormFieldBriefTypePayload(APIObject):
    type: Literal["brief"]
    values: list[Values]


Values1: TypeAlias = WorkRequestTextBriefResponse | WorkRequestAttachmentBriefResponse


class WorkRequestRequestFormFieldBriefTypeResponse(APIObject):
    type: Literal["brief"]
    values: list[Values1]


WorkRequestRequestFormFieldUpdateResponse: TypeAlias = (
    WorkRequestRequestFormFieldBriefTypeResponse
    | WorkRequestRequestFormFieldCheckboxTypeResponse
    | WorkRequestRequestFormFieldCurrencyNumberTypePayload
    | WorkRequestRequestFormFieldDateTypePayload
    | WorkRequestRequestFormFieldDropdownTypeResponse
    | WorkRequestRequestFormFieldFileTypeResponse
    | WorkRequestRequestFormFieldLabelTypeResponse
    | WorkRequestRequestFormFieldPercentageNumberTypePayload
    | WorkRequestRequestFormFieldRadioButtonTypeResponse
    | WorkRequestRequestFormFieldRichtextTypePayload
    | WorkRequestRequestFormFieldSimpleNumberTypePayload
    | WorkRequestRequestFormFieldTextAreaTypePayload
    | WorkRequestRequestFormFieldTextTypePayload
)
FormFields2: TypeAlias = (
    BriefTypeFormFieldResponse
    | DateTypeFormFieldResponse
    | FileTypeFormFieldResponse
    | MultiChoiceTypeFormFieldResponse
    | NumberTypeFormFieldResponse
    | RadioButtonTypeFormFieldResponse
    | TextTypeFormFieldResponse
)


class WorkRequestResponse(APIObject):
    assignees: list[Assignee]
    created_at: str
    created_by: CreatedBy
    form_fields: list[FormFields2]
    id: str
    links: Links52
    modified_at: str
    priority: Literal["Low", "Medium", "High"]
    reference_id: str
    status: Literal["Accepted", "Declined", "Submitted", "Completed"]
    template: Template


class WorkflowStep(APIObject):
    description: str | None
    duration: int | None
    label: str
    substeps: list[WorkflowSubStep]


class ListAssetsResponse(APIObject):
    data: list[AssetResponse]
    pagination: Pagination2
    total_count: float


class ListCampaignFieldsResponse(APIObject):
    data: list[FieldListResponseItem]
    pagination: Pagination2


class ListEventFieldsResponse(APIObject):
    data: list[FieldListResponseItem]
    pagination: Pagination2


Data: TypeAlias = (
    BaseSettingsFieldsResponse
    | LabelTypeSettingsFieldResponse
    | DropdownTypeSettingsFieldResponse
    | CheckboxAndRadioTypeSettingsFieldResponse
    | GenericNumberTypeSettingsFieldResponse
    | CurrencyNumberTypeSettingsFieldResponse
)


class ListFieldsResponse(APIObject):
    data: list[Data]
    pagination: Pagination2


class ListLabelGroupsResponse(APIObject):
    data: list[LabelGroup]
    pagination: Pagination2


class ListTaskFieldsResponse(APIObject):
    data: list[FieldListResponseItem]
    pagination: Pagination2


class ListWorkRequestsResponse(APIObject):
    data: list[WorkRequestResponse]
    pagination: Pagination2


AssetFieldListResponseItem: TypeAlias = (
    AssetFieldTypeCommon
    | AssetFieldTypeLabel
    | AssetFieldTypeDropdown
    | AssetFieldTypeRadio
    | AssetFieldTypeCheckbox
    | AssetFieldTypeSimpleNumber
    | AssetFieldTypePercentageNumber
    | AssetFieldTypeCurrencyNumber
)
AssetFieldUpdateRequest: TypeAlias = (
    StringTypeObjectFieldUpdatePayload
    | MultiChoiceTypeObjectFieldUpdatePayload
    | DropdownTypeObjectFieldUpdatePayload
    | RadioButtonTypeObjectFieldUpdatePayload
    | NumberTypeObjectFieldUpdatePayload
    | DateTypeObjectFieldUpdatePayload
    | AssetTypeObjectFieldUpdatePayload
)


class AssetTypeObjectField(BaseObjectField, AssetTypeObjectFieldUpdatePayload):
    pass


class BaseFieldDefinition(APIObject):
    base_type: BaseFieldDefinitionType
    core: CoreFieldDef
    default_values: (
        list[LocationDefaultValue | bool | str | dict[str, Any] | list[Any]] | None
    )


class ChoiceFieldDefinition(APIObject):
    choices: dict[str, str]
    core: CoreFieldDef
    default_values: list[str] | None
    display_option: ChoiceDisplayOption | None


class ContentTypeFieldDefinition(APIObject):
    allow_ref_edit: bool | None
    allowed_content_types: list[str]
    content_type_links: dict[str, AllowedContentTypeItem] | None
    core: CoreFieldDef
    default_value: str | None
    ref_type: ContentTypeFieldEmbedMixConfig


Fields: TypeAlias = (
    StringTypeObjectField
    | MultiChoiceTypeObjectField
    | DropdownTypeObjectField
    | RadioButtonTypeObjectField
    | NumberTypeObjectField
    | DateTypeObjectField
    | AssetTypeObjectField
)


class EventCreateRequest(APIObject):
    campaign_id: str | None
    description: str | None
    end_date: str
    fields: list[Fields] | None
    is_all_day: bool
    start_date: str
    title: str


EventFieldsUpdateRequest1: TypeAlias = (
    MultiChoiceTypeObjectField
    | StringTypeObjectField
    | DropdownTypeObjectField
    | RadioButtonTypeObjectField
    | NumberTypeObjectField
    | DateTypeObjectField
    | AssetTypeObjectField
)
EventFieldsUpdateRequest: TypeAlias = list[EventFieldsUpdateRequest1]
ObjectFieldCreateRequest: TypeAlias = (
    StringTypeObjectField
    | MultiChoiceTypeObjectField
    | DropdownTypeObjectField
    | RadioButtonTypeObjectField
    | NumberTypeObjectField
    | DateTypeObjectField
    | AssetTypeObjectField
)


class SCContentTypeCreateRequest(APIObject):
    created_by: str
    details: CoreContentType
    expected_locales: list[str] | None
    field_definitions: list[
        ContentTypeFieldDefinition
        | LibraryAssetFieldDefinition
        | TextFieldDefinition
        | BaseFieldDefinition
        | DatetimeFieldDefinition
        | RichTextFieldDefinition
        | ChoiceFieldDefinition
        | NumberFieldDefinition
    ]
    source: str | None
    source_id: str | None
    source_metadata: str | None


class SCContentTypeVersion(BaseContentTypeVersionModel):
    field_definitions: list[
        ContentTypeFieldDefinition
        | LibraryAssetFieldDefinition
        | TextFieldDefinition
        | BaseFieldDefinition
        | DatetimeFieldDefinition
        | RichTextFieldDefinition
        | ChoiceFieldDefinition
        | NumberFieldDefinition
    ]


class SCContentTypeVersionCreateRequest(APIObject):
    created_by: str
    expected_locales: list[str] | None
    field_definitions: list[
        ContentTypeFieldDefinition
        | LibraryAssetFieldDefinition
        | TextFieldDefinition
        | BaseFieldDefinition
        | DatetimeFieldDefinition
        | RichTextFieldDefinition
        | ChoiceFieldDefinition
        | NumberFieldDefinition
    ]


class SettingsWorkflow(APIObject):
    blacklisted_channels: list[WorkflowChannel]
    custom_fields: list[WorkflowCustomfield]
    default_channels: list[WorkflowChannel]
    description: str | None
    is_flexible_workflow: bool
    is_not_pushed_to_library_by_default: bool
    is_released_as_asset: bool
    labels: list[WorkflowLabel]
    name: str
    steps: list[WorkflowStep]


TaskAssetFieldsUpdateRequest1: TypeAlias = (
    StringTypeObjectField
    | MultiChoiceTypeObjectField
    | DropdownTypeObjectField
    | RadioButtonTypeObjectField
    | NumberTypeObjectField
    | DateTypeObjectField
    | AssetTypeObjectField
)
TaskAssetFieldsUpdateRequest: TypeAlias = list[TaskAssetFieldsUpdateRequest1]


class TaskResponse(APIObject):
    campaign_id: str | None
    due_at: str | None
    id: str
    is_archived: bool
    is_completed: bool
    labels: list[ResourceLabelResponse]
    links: Links35
    milestone_id: str | None
    modified_at: str | None
    reference_id: str
    start_at: str | None
    status: Literal[
        "Archived", "Completed", "Overdue", "Not Started", "In Progress", "On Hold"
    ]
    steps: list[TaskStep]
    title: str
    workflow_id: str | None


class TemplateChoiceFormField(BaseTemplateFormField):
    logic_rules: TemplateLogicRule
    type_specific_meta: TypeSpecificMeta


class TemplateCurrencyNumberFormField(BaseTemplateFormField):
    logic_rules: TemplateLogicRule
    type_specific_meta: TypeSpecificMeta1


class TemplateDefaultFormField(BaseTemplateFormField):
    logic_rules: TemplateLogicRule
    type_specific_meta: dict[str, Any] | None


class TemplateInstructionFormField(BaseTemplateFormField):
    logic_rules: TemplateLogicRule
    type_specific_meta: TypeSpecificMeta2


FormFields: TypeAlias = (
    TemplateDefaultFormField
    | TemplateCurrencyNumberFormField
    | TemplateSimpleNumberFormField
    | TemplatePercentageNumberFormField
    | TemplateChoiceFormField
    | TemplateInstructionFormField
)


class TemplateResponse(BaseTemplateResponse):
    form_fields: list[FormFields]
    links: Links39


class VersionedContentTypeModel(APIObject):
    component: bool
    content_type_guid: str
    created_at: str
    created_by: str
    description: str | None
    disabled: bool | None
    links: Links46
    name: str
    source: str | None
    source_id: str | None
    source_metadata: str | None
    thumbnail_guid: str | None
    updated_at: str
    updated_by: str
    version: SCContentTypeVersion


WorkRequestFormFieldUpdateRequest: TypeAlias = (
    WorkRequestRequestFormFieldBriefTypePayload
    | WorkRequestRequestFormFieldCheckboxTypePayload
    | WorkRequestRequestFormFieldCurrencyNumberTypePayload
    | WorkRequestRequestFormFieldDateTypePayload
    | WorkRequestRequestFormFieldDropdownTypePayload
    | WorkRequestRequestFormFieldFileTypePayload
    | WorkRequestRequestFormFieldLabelTypePayload
    | WorkRequestRequestFormFieldPercentageNumberTypePayload
    | WorkRequestRequestFormFieldRadioButtonTypePayload
    | WorkRequestRequestFormFieldRichtextTypePayload
    | WorkRequestRequestFormFieldSimpleNumberTypePayload
    | WorkRequestRequestFormFieldTextAreaTypePayload
    | WorkRequestRequestFormFieldTextTypePayload
)


class ListAssetFieldsResponse(APIObject):
    data: list[AssetFieldListResponseItem]
    pagination: Pagination2


class UpdateAssetFieldsResponse(APIObject):
    data: list[AssetFieldListResponseItem]
    links: Links56


class ListTaskAssetFieldsResponse(APIObject):
    data: list[AssetFieldListResponseItem] | None
    pagination: Pagination2 | None


AssetFieldsUpdateRequest1: TypeAlias = (
    StringTypeObjectField
    | MultiChoiceTypeObjectField
    | DropdownTypeObjectField
    | RadioButtonTypeObjectField
    | NumberTypeObjectField
    | DateTypeObjectField
    | AssetTypeObjectField
)
AssetFieldsUpdateRequest: TypeAlias = list[AssetFieldsUpdateRequest1]


class SCContentType(BaseContentTypeModel):
    latest_version: SCContentTypeVersion


class Resources(APIObject):
    apps: list[SettingsApp]
    custom_fields: CustomFields
    labels: list[SettingsLabelGroup]
    routing_rules: list[SettingsRoutingRule]
    templates: list[SettingsTemplate]
    webhooks: list[SettingsWebhook]
    workflows: list[SettingsWorkflow]


class Settings(APIObject):
    resources: Resources


class SettingsResources(APIObject):
    apps: list[SettingsApp]
    custom_fields: CustomFields
    labels: list[SettingsLabelGroup]
    routing_rules: list[SettingsRoutingRule]
    templates: list[SettingsTemplate]
    webhooks: list[SettingsWebhook]
    workflows: list[SettingsWorkflow]


class Resources1(APIObject):
    apps: list[SettingsApp]
    custom_fields: CustomFields
    labels: list[SettingsLabelGroup]
    routing_rules: list[SettingsRoutingRule]
    templates: list[SettingsTemplate]
    webhooks: list[SettingsWebhook]
    workflows: list[SettingsWorkflow]


class SettingsUpdateResponse(APIObject):
    changeset: Changeset
    resources: Resources1


class ContentDetailsModel(APIObject):
    content_guid: str
    content_type: VersionedContentTypeModel | None
    content_type_guid: str
    content_type_name: str
    created_at: str
    created_by: str
    expired: bool | None
    expiry_datetime: str | None
    latest_fields_version: ContentFieldsVersionDetails
    links: Links12
    primary_locale: str | None
    root_content: bool
    source: str | None
    source_id: str | None
    source_metadata: str | None
    template_guid: str | None
    title: str
    updated_at: str
    updated_by: str


class ContentFieldValueExpandedModel(APIObject):
    content_details: ContentDetailsModel
    content_guid: str
    content_url: str | None
    embedded: bool
    order_index: int | None


class ContentFieldWithEmbeddedPatchValueModel(APIObject):
    content_id: str
    patch_details: RecursivePatchStructuredContentFields


class ContentFieldsVersionDetails(APIObject):
    content_hash: str
    created_at: str
    created_by: str
    fields: dict[str, list[LocalizedFieldValues]]
    source_id: str | None
    source_metadata: str | None
    validation: ContentFieldsVersionValidation | None
    version_guid: str


class LibraryStructuredContent(APIObject):
    content_body: ContentDetailsModel | None
    created_at: str
    file_location: str
    folder_id: str | None
    id: str
    is_archived: bool
    labels: list[ResourceLabelResponse]
    links: Links16
    modified_at: str
    owner_organization_id: str
    title: str


class ContentBody1(APIObject):
    expired: bool | None
    expiry_datetime: str | None
    fields: StructuredContentFields | None
    source: str | None
    source_id: str | None
    source_metadata: str | None
    title: str | None


class LibraryStructuredContentUpdateRequest(APIObject):
    content_body: ContentBody1 | None
    is_archived: bool | None
    title: str | None


class LocalizedFieldValues(APIObject):
    field_values: list[
        NumberFieldValueModel
        | BooleanFieldValueModel
        | DatetimeFieldValueModel
        | TextFieldValueModel
        | RichTextFieldValueModel
        | LibraryAssetFieldValueModel
        | ContentFieldValueExpandedModel
        | ContentFieldValueModel
        | URLFieldValueModel
        | JSONFieldValueModel
        | ChoiceFieldValueModel
        | LocationFieldValueModel
    ]
    locale: str


class PatchLocalizedFieldValuesWithEmbeddedContent(APIObject):
    field_values: list[
        NumberFieldValueModel
        | BooleanFieldValueModel
        | DatetimeFieldValueModel
        | TextFieldValueModel
        | RichTextFieldValueModel
        | LibraryAssetFieldValueModel
        | ContentFieldValueWithEmbeddedModel
        | ContentFieldValueModel
        | URLFieldValueModel
        | JSONFieldValueModel
        | ChoiceFieldValueModel
        | LocationFieldValueModel
        | DeleteFieldValueModel
        | ContentFieldWithEmbeddedPatchValueModel
    ]
    locale: str


class RecursivePatchStructuredContentFields(APIObject):
    content_body: StructuredContentFields | None
    expired: bool | None
    expiry_datetime: str | None
    root_content: bool | None
    source: str | None
    source_id: str | None
    source_metadata: str | None
    title: str | None


class SCContentMigrationCreateRequest(APIObject):
    created_by: str
    fields: dict[str, list[LocalizedFieldValues]] | None
    new_content_type_version_id: str | None
    source: str | None
    source_id: str | None


class SCContentTypeManagedMigrationCreateRequest(APIObject):
    created_by: str
    default_values: LocalizedFieldValues | None
    source_content_type_version_id: str


class SCContentTypeManagedMigrationValidateRequest(APIObject):
    default_values: LocalizedFieldValues | None
    source_content_type_version_id: str


StructuredContentFields: TypeAlias = dict[
    str,
    list[PatchLocalizedFieldValuesWithEmbeddedContent | DeleteLocaleFieldValueModel],
]


class TaskStructuredContentDraftRequest(APIObject):
    content_body: LocalizedFieldValues


class TaskStructuredContentUpdateRequest(APIObject):
    expired: bool | None
    expiry_datetime: str | None
    fields: StructuredContentFields | None
    primary_locale: str | None
    source: str | None
    source_id: str | None
    source_metadata: str | None
    title: str | None


__all__ = [
    "AddUrlToTaskRequest",
    "ArticleAuthorResponse",
    "AssetFieldListResponseItem",
    "AssetFieldUpdateRequest",
    "AssetFieldUpdateResponse",
    "AssetFieldsUpdateRequest",
    "AssetLineageCreateRequest",
    "AssetLineageResponse",
    "AssetPermissionBulkCreateRequest",
    "AssetPermissionListResponseItem",
    "AssetPermissionUpdateRequest",
    "AssetResponse",
    "AssetTypeObjectFieldUpdatePayload",
    "AssetVersionCreateRequest",
    "AttachmentRequest",
    "AttachmentResponse",
    "BaseAssetRenditionResponse",
    "BaseFormFieldRequest",
    "BaseFormFieldResponse",
    "BaseObjectFieldUpdatePayload",
    "BaseSettingsFieldCreatePayload",
    "BaseSettingsFieldUpdatePayload",
    "BaseSettingsFieldsResponse",
    "BaseTemplateResponse",
    "BatchFileUrlResponse",
    "BrandComplianceCategoriesResponse",
    "BriefTypeFormFieldRequest",
    "BriefTypeFormFieldResponse",
    "BudgetResponse",
    "CampaignBriefResponse",
    "CampaignCommentResponse",
    "CampaignCreateRequest",
    "CampaignFieldUpdateRequest",
    "CampaignFieldUpdateResponse",
    "CampaignListResponseItem",
    "CampaignResponse",
    "CampaignUpdateRequest",
    "CheckboxAndRadioTypeSettingsFieldCreatePayload",
    "CheckboxAndRadioTypeSettingsFieldResponse",
    "CommentCreateRequest",
    "CommentWithReplyCreateRequest",
    "CompleteMultipartUploadResponse",
    "CreateFieldResponse",
    "CreateMultipartUploadResponse",
    "CreateSCContentTypeManagedMigrationResponse",
    "CreateTaskStructuredContentDraftResponse",
    "CreativeAssetRequest",
    "CreativeAssetResponse",
    "CurrencyNumberTypeSettingsFieldCreatePayload",
    "CurrencyNumberTypeSettingsFieldResponse",
    "CurrencyNumberTypeSettingsFieldUpdatePayload",
    "DateTypeFormFieldRequest",
    "DateTypeFormFieldResponse",
    "DateTypeObjectFieldUpdatePayload",
    "DetailedAssetRenditionResponse",
    "DropdownTypeObjectFieldUpdatePayload",
    "DropdownTypeSettingsFieldCreatePayload",
    "DropdownTypeSettingsFieldResponse",
    "EventCreateRequest",
    "EventFieldsUpdateRequest",
    "EventResponse",
    "EventUpdateRequest",
    "FeaturedImageResponse",
    "FieldListResponseItem",
    "FileTypeFormFieldRequest",
    "FileTypeFormFieldResponse",
    "FileTypeFormFieldValueRequest",
    "FileTypeFormFieldValueResponse",
    "FileUrlBulkCreateRequest",
    "FolderCreateRequest",
    "FolderPermissionBulkCreateRequest",
    "FolderPermissionListResponseItem",
    "FolderPermissionUpdateRequest",
    "FolderResponse",
    "FolderUpdateRequest",
    "GenericNumberTypeSettingsFieldCreatePayload",
    "GenericNumberTypeSettingsFieldResponse",
    "GenericNumberTypeSettingsFieldUpdatePayload",
    "GenericSettingsFieldChoiceCreatePayload",
    "GenericSettingsFieldChoiceUpdatePayload",
    "GetMultipartUploadStatusResponse",
    "GetUploadUrlResponse",
    "LabelAndDropdownTypeSettingsFieldUpdatePayload",
    "LabelTypeSettingsFieldChoiceCreatePayload",
    "LabelTypeSettingsFieldChoiceUpdatePayload",
    "LabelTypeSettingsFieldCreatePayload",
    "LabelTypeSettingsFieldResponse",
    "LibraryAssetCreateRequest",
    "LibraryAssetVersionResponse",
    "LibraryImageUpdateRequest",
    "LibraryRawFileUpdateRequest",
    "LibraryStructuredContentCreateRequest",
    "LibraryStructuredContentUpdateRequest",
    "LibraryVideoUpdateRequest",
    "ListAssetFieldsResponse",
    "ListAssetLineagesResponse",
    "ListAssetRenditionsResponse",
    "ListAssetsResponse",
    "ListBrandComplianceCategoriesResponse",
    "ListCampaignFieldsResponse",
    "ListCampaignsResponse",
    "ListEventFieldsResponse",
    "ListEventsResponse",
    "ListFieldsResponse",
    "ListFoldersResponse",
    "ListLabelGroupsResponse",
    "ListMilestonesResponse",
    "ListSCContentTypeManagedMigrationsResponse",
    "ListSCContentTypeVersionsResponse",
    "ListSCContentTypesResponse",
    "ListTaskAssetCommentsResponse",
    "ListTaskAssetDraftsResponse",
    "ListTaskAssetFieldsResponse",
    "ListTaskAssetsResponse",
    "ListTaskAttachmentsResponse",
    "ListTaskCustomFieldChoicesResponse",
    "ListTaskCustomFieldsResponse",
    "ListTaskFieldsResponse",
    "ListTaskSubStepCommentsResponse",
    "ListTaskSubStepFieldsResponse",
    "ListTasksResponse",
    "ListTeamsResponse",
    "ListWorkRequestApprovedAssetsResponse",
    "ListWorkRequestCommentsResponse",
    "ListWorkRequestRelatedResourcesResponse",
    "ListWorkRequestsResponse",
    "ListWorkflowsResponse",
    "MigrateSCContentResponse",
    "MilestoneCreateRequest",
    "MilestoneResponse",
    "MilestoneUpdateRequest",
    "MultiChoiceTypeFormFieldRequest",
    "MultiChoiceTypeFormFieldResponse",
    "MultiChoiceTypeObjectFieldUpdatePayload",
    "NumberTypeFormFieldRequest",
    "NumberTypeFormFieldResponse",
    "NumberTypeObjectFieldUpdatePayload",
    "ObjectFieldCreateRequest",
    "ObjectFieldCreateResponse",
    "PublishingChannelListResponse",
    "PublishingChannelResponse",
    "PublishingEventMetadataBulkCreateRequest",
    "PublishingEventMetadataBulkCreateResponse",
    "PublishingEventMetadataCreateRequest",
    "PublishingEventMetadataListResponse",
    "PublishingEventMetadataResponse",
    "PublishingEventResponse",
    "RadioButtonTypeFormFieldRequest",
    "RadioButtonTypeFormFieldResponse",
    "RadioButtonTypeObjectFieldUpdatePayload",
    "RelatedAssetsListResponse",
    "RenditionConfigResponse",
    "ReplaceRelatedAssetsRequest",
    "ReplaceRelatedAssetsResponse",
    "ResourceChoiceResponse",
    "ResourceLabelRequest",
    "ResourceLabelResponse",
    "SCContentMigrationCreateRequest",
    "SCContentPreviewAcknowledgeRequest",
    "SCContentPreviewCompleteRequest",
    "SCContentTypeCreateRequest",
    "SCContentTypeCreateResponse",
    "SCContentTypeManagedMigrationCreateRequest",
    "SCContentTypeManagedMigrationResponse",
    "SCContentTypeManagedMigrationStartResponse",
    "SCContentTypeManagedMigrationValidateRequest",
    "SCContentTypeUpdateRequest",
    "SCContentTypeUpdateResponse",
    "SCContentTypeVersionCreateRequest",
    "SettingsFieldChoiceCreateRequest",
    "SettingsFieldChoiceCreateResponse",
    "SettingsFieldChoiceUpdateRequest",
    "SettingsFieldCreateRequest",
    "SettingsFieldUpdateRequest",
    "SettingsUpdateResponse",
    "StringTypeObjectFieldUpdatePayload",
    "TaskAssetCommentResponse",
    "TaskAssetDraftBrandComplianceRequest",
    "TaskAssetDraftBrandComplianceResponse",
    "TaskAssetDraftListResponseItem",
    "TaskAssetDraftResponse",
    "TaskAssetFieldsUpdateRequest",
    "TaskAssetRequest",
    "TaskAssetResponse",
    "TaskBriefResponse",
    "TaskCommentResponse",
    "TaskCreateRequest",
    "TaskCustomFieldChoiceListResponseItem",
    "TaskCustomFieldUpdateRequest",
    "TaskExternalWorkRequest",
    "TaskExternalWorkResponse",
    "TaskFieldUpdateRequest",
    "TaskListResponseItem",
    "TaskPublishingIntentCreateRequest",
    "TaskPublishingIntentResponse",
    "TaskResponse",
    "TaskStepRequest",
    "TaskStructuredContentCreateRequest",
    "TaskStructuredContentDraftRequest",
    "TaskStructuredContentUpdateRequest",
    "TaskSubStepCommentResponse",
    "TaskSubStepCommentUpdateRequest",
    "TaskSubStepRequest",
    "TaskUpdateRequest",
    "TaskUrlResponse",
    "TemplateListResponse",
    "TemplateResponse",
    "TextTypeFormFieldRequest",
    "TextTypeFormFieldResponse",
    "UpdateAssetFieldsResponse",
    "UpdateSCContentTypeManagedMigrationResponse",
    "UserListResponse",
    "UserListResponseItem",
    "UserResponse",
    "ValidateSCContentTypeManagedMigrationResponse",
    "WorkRequestApprovedAssetResponse",
    "WorkRequestAttachmentBriefRequest",
    "WorkRequestAttachmentBriefResponse",
    "WorkRequestCampaignRequest",
    "WorkRequestCampaignResponse",
    "WorkRequestCommentResponse",
    "WorkRequestCreateRequest",
    "WorkRequestFormFieldUpdateRequest",
    "WorkRequestRelatedResourceResponse",
    "WorkRequestRequestFormFieldBriefTypePayload",
    "WorkRequestRequestFormFieldBriefTypeResponse",
    "WorkRequestRequestFormFieldCheckboxTypePayload",
    "WorkRequestRequestFormFieldCheckboxTypeResponse",
    "WorkRequestRequestFormFieldCurrencyNumberTypePayload",
    "WorkRequestRequestFormFieldDateTypePayload",
    "WorkRequestRequestFormFieldDropdownTypePayload",
    "WorkRequestRequestFormFieldDropdownTypeResponse",
    "WorkRequestRequestFormFieldFileTypePayload",
    "WorkRequestRequestFormFieldFileTypeResponse",
    "WorkRequestRequestFormFieldLabelTypePayload",
    "WorkRequestRequestFormFieldLabelTypeResponse",
    "WorkRequestRequestFormFieldPercentageNumberTypePayload",
    "WorkRequestRequestFormFieldRadioButtonTypePayload",
    "WorkRequestRequestFormFieldRadioButtonTypeResponse",
    "WorkRequestRequestFormFieldRichtextTypePayload",
    "WorkRequestRequestFormFieldSimpleNumberTypePayload",
    "WorkRequestRequestFormFieldTextAreaTypePayload",
    "WorkRequestRequestFormFieldTextTypePayload",
    "WorkRequestRequestFormFieldUpdateResponse",
    "WorkRequestResponse",
    "WorkRequestTaskRequest",
    "WorkRequestTaskResponse",
    "WorkRequestTextBriefRequest",
    "WorkRequestTextBriefResponse",
    "WorkRequestUpdateRequest",
    "WorkflowListResponseItem",
    "WorkflowResponse",
]
