# Auto-generated - DO NOT EDIT

from __future__ import annotations

from .._types import Routes

ROUTES: Routes = {
    "asset": {"get_asset_url": {"method": "GET", "url": "/asset-urls/{asset_id}"}},
    "brand_compliance": {
        "get_task_asset_draft_brand_compliance": {
            "method": "GET",
            "url": "/tasks/{task_id}/assets/{asset_id}/drafts/{draft_id}/brand-compliance",
        },
        "list_brand_compliance_categories": {
            "method": "GET",
            "url": "/brand-compliance/categories",
        },
        "update_task_asset_draft_brand_compliance": {
            "method": "PUT",
            "url": "/tasks/{task_id}/assets/{asset_id}/drafts/{draft_id}/brand-compliance",
        },
    },
    "campaign": {
        "add_attachment_to_campaign": {
            "method": "POST",
            "url": "/campaigns/{id}/attachments",
        },
        "add_comment_to_campaign": {
            "method": "POST",
            "url": "/campaigns/{id}/comments",
        },
        "add_field_to_campaign": {"method": "POST", "url": "/campaigns/{id}/fields"},
        "create_campaign": {"method": "POST", "url": "/campaigns"},
        "get_campaign": {"method": "GET", "url": "/campaigns/{id}"},
        "get_campaign_brief": {"method": "GET", "url": "/campaigns/{id}/brief"},
        "list_campaign_fields": {"method": "GET", "url": "/campaigns/{id}/fields"},
        "list_campaigns": {"method": "GET", "url": "/campaigns"},
        "update_campaign": {"method": "PATCH", "url": "/campaigns/{id}"},
        "update_campaign_field": {
            "method": "PUT",
            "url": "/campaigns/{campaign_id}/fields/{field_id}",
        },
    },
    "event": {
        "create_event": {"method": "POST", "url": "/events"},
        "get_event": {"method": "GET", "url": "/events/{id}"},
        "list_event_fields": {"method": "GET", "url": "/events/{id}/fields"},
        "list_events": {"method": "GET", "url": "/events"},
        "update_event": {"method": "PATCH", "url": "/events/{id}"},
        "update_event_fields": {"method": "PUT", "url": "/events/{id}/fields"},
    },
    "field": {
        "create_field": {"method": "POST", "url": "/fields"},
        "create_field_choices": {"method": "POST", "url": "/fields/{id}/choices"},
        "delete_field_choice": {
            "method": "DELETE",
            "url": "/fields/{field_id}/choices/{choice_id}",
        },
        "list_fields": {"method": "GET", "url": "/fields"},
        "update_field": {"method": "PATCH", "url": "/fields/{id}"},
        "update_field_choice": {
            "method": "PATCH",
            "url": "/fields/{field_id}/choices/{choice_id}",
        },
    },
    "label": {"list_label_groups": {"method": "GET", "url": "/label-groups"}},
    "library": {
        "add_asset_permissions": {
            "method": "POST",
            "url": "/assets/{asset_id}/permissions",
        },
        "add_folder_permissions": {
            "method": "POST",
            "url": "/folders/{id}/permissions",
        },
        "create_asset": {"method": "POST", "url": "/assets"},
        "create_asset_lineage": {
            "method": "POST",
            "url": "/assets/{asset_id}/lineages",
        },
        "create_asset_version": {
            "method": "POST",
            "url": "/assets/{asset_id}/versions",
        },
        "create_file_urls": {"method": "POST", "url": "/file-urls"},
        "create_folder": {"method": "POST", "url": "/folders"},
        "create_structured_content": {"method": "POST", "url": "/structured-contents"},
        "delete_asset_lineage": {
            "method": "DELETE",
            "url": "/assets/{asset_id}/lineages/{lineage_id}",
        },
        "delete_folder": {"method": "DELETE", "url": "/folders/{id}"},
        "delete_image": {"method": "DELETE", "url": "/images/{id}"},
        "delete_raw_file": {"method": "DELETE", "url": "/raw-files/{id}"},
        "delete_video": {"method": "DELETE", "url": "/videos/{id}"},
        "get_article": {"method": "GET", "url": "/articles/{id}"},
        "get_folder": {"method": "GET", "url": "/folders/{id}"},
        "get_image": {"method": "GET", "url": "/images/{id}"},
        "get_raw_file": {"method": "GET", "url": "/raw-files/{id}"},
        "get_rendition": {"method": "GET", "url": "/renditions/{id}"},
        "get_rendition_config": {"method": "GET", "url": "/rendition-configs/{id}"},
        "get_structured_content": {"method": "GET", "url": "/structured-contents/{id}"},
        "get_video": {"method": "GET", "url": "/videos/{id}"},
        "list_asset_fields": {"method": "GET", "url": "/assets/{asset_id}/fields"},
        "list_asset_lineages": {"method": "GET", "url": "/asset-lineages"},
        "list_asset_permissions": {
            "method": "GET",
            "url": "/assets/{asset_id}/permissions",
        },
        "list_asset_renditions": {
            "method": "GET",
            "url": "/assets/{asset_id}/renditions",
        },
        "list_assets": {"method": "GET", "url": "/assets"},
        "list_folder_permissions": {
            "method": "GET",
            "url": "/folders/{id}/permissions",
        },
        "list_folders": {"method": "GET", "url": "/folders"},
        "list_related_assets": {
            "method": "GET",
            "url": "/assets/{asset_id}/related-assets",
        },
        "remove_asset_permission": {
            "method": "DELETE",
            "url": "/assets/{asset_id}/permissions/{accessor_id}",
        },
        "remove_folder_permission": {
            "method": "DELETE",
            "url": "/folders/{id}/permissions/{accessor_id}",
        },
        "replace_related_assets": {
            "method": "PUT",
            "url": "/assets/{asset_id}/related-assets",
        },
        "update_asset_field": {
            "method": "PUT",
            "url": "/assets/{asset_id}/fields/{field_id}",
        },
        "update_asset_fields": {"method": "PUT", "url": "/assets/{asset_id}/fields"},
        "update_asset_permission": {
            "method": "PATCH",
            "url": "/assets/{asset_id}/permissions/{accessor_id}",
        },
        "update_folder": {"method": "PATCH", "url": "/folders/{id}"},
        "update_folder_permission": {
            "method": "PATCH",
            "url": "/folders/{id}/permissions/{accessor_id}",
        },
        "update_image": {"method": "PATCH", "url": "/images/{id}"},
        "update_raw_file": {"method": "PATCH", "url": "/raw-files/{id}"},
        "update_structured_content": {
            "method": "PATCH",
            "url": "/structured-contents/{id}",
        },
        "update_video": {"method": "PATCH", "url": "/videos/{id}"},
    },
    "milestone": {
        "create_milestone": {"method": "POST", "url": "/milestones"},
        "get_milestone": {"method": "GET", "url": "/milestones/{id}"},
        "list_milestones": {"method": "GET", "url": "/milestones"},
        "update_milestone": {"method": "PATCH", "url": "/milestones/{id}"},
    },
    "publishing": {
        "bulk_create_publishing_event_metadata": {
            "method": "POST",
            "url": "/publishing-events/{publishing_event_id}/publishing-metadata",
        },
        "get_publishing_event": {
            "method": "GET",
            "url": "/publishing-events/{publishing_event_id}",
        },
        "get_publishing_event_asset_metadata": {
            "method": "GET",
            "url": "/publishing-events/{publishing_event_id}/assets/{asset_id}/publishing-metadata/{publishing_metadata_id}",
        },
        "list_publishing_channels": {"method": "GET", "url": "/publishing-channels"},
        "list_publishing_event_metadata": {
            "method": "GET",
            "url": "/publishing-events/{publishing_event_id}/publishing-metadata",
        },
    },
    "settings": {
        "get_settings": {"method": "GET", "url": "/settings"},
        "update_settings": {"method": "POST", "url": "/settings"},
    },
    "structured_content": {
        "acknowledge_sccontent_preview": {
            "method": "POST",
            "url": "/structured-content/contents/{content_id}/versions/{version_id}/previews/{preview_id}/acknowledge",
        },
        "complete_sccontent_preview": {
            "method": "POST",
            "url": "/structured-content/contents/{content_id}/versions/{version_id}/previews/{preview_id}/complete",
        },
        "create_sccontent_type": {
            "method": "POST",
            "url": "/structured-content/content-types",
        },
        "create_sccontent_type_managed_migration": {
            "method": "POST",
            "url": "/structured-content/content-types/{content_type_id}/managed-migrations",
        },
        "create_sccontent_type_version": {
            "method": "POST",
            "url": "/structured-content/content-types/{content_type_id}/versions",
        },
        "delete_sccontent_type_managed_migration": {
            "method": "DELETE",
            "url": "/structured-content/content-types/{content_type_id}/managed-migrations/{job_id}",
        },
        "get_sccontent_type": {
            "method": "GET",
            "url": "/structured-content/content-types/{content_type_id}",
        },
        "get_sccontent_type_managed_migration": {
            "method": "GET",
            "url": "/structured-content/content-types/{content_type_id}/managed-migrations/{job_id}",
        },
        "get_sccontent_type_version": {
            "method": "GET",
            "url": "/structured-content/content-types/{content_type_id}/versions/{version_id}",
        },
        "list_sccontent_type_managed_migrations": {
            "method": "GET",
            "url": "/structured-content/content-types/{content_type_id}/managed-migrations",
        },
        "list_sccontent_type_versions": {
            "method": "GET",
            "url": "/structured-content/content-types/{content_type_id}/versions",
        },
        "list_sccontent_types": {
            "method": "GET",
            "url": "/structured-content/content-types",
        },
        "migrate_sccontent": {
            "method": "POST",
            "url": "/structured-content/contents/{content_id}/migration",
        },
        "start_sccontent_type_managed_migration": {
            "method": "POST",
            "url": "/structured-content/content-types/{content_type_id}/managed-migrations/{job_id}/start",
        },
        "update_sccontent_type": {
            "method": "POST",
            "url": "/structured-content/content-types/{content_type_id}",
        },
        "update_sccontent_type_managed_migration": {
            "method": "PATCH",
            "url": "/structured-content/content-types/{content_type_id}/managed-migrations/{job_id}",
        },
        "validate_sccontent_type_managed_migration": {
            "method": "POST",
            "url": "/structured-content/content-types/{content_type_id}/managed-migrations/validate",
        },
    },
    "task": {
        "add_asset_to_task": {"method": "POST", "url": "/tasks/{id}/assets"},
        "add_comment_to_task": {"method": "POST", "url": "/tasks/{task_id}/comments"},
        "add_comment_to_task_asset": {
            "method": "POST",
            "url": "/tasks/{task_id}/assets/{asset_id}/comments",
        },
        "add_comment_to_task_sub_step": {
            "method": "POST",
            "url": "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments",
        },
        "add_draft_to_task_asset": {
            "method": "POST",
            "url": "/tasks/{task_id}/assets/{asset_id}/drafts",
        },
        "add_field_to_task": {"method": "POST", "url": "/tasks/{task_id}/fields"},
        "add_structured_content_to_task": {
            "method": "POST",
            "url": "/tasks/{task_id}/structured-contents",
        },
        "add_url_to_task": {"method": "POST", "url": "/tasks/{task_id}/urls"},
        "create_task": {"method": "POST", "url": "/tasks"},
        "create_task_publishing_intent": {
            "method": "POST",
            "url": "/tasks/{task_id}/publishing-intents",
        },
        "create_task_structured_content_draft": {
            "method": "POST",
            "url": "/tasks/{task_id}/structured-contents/{content_id}/drafts",
        },
        "delete_task_structured_content": {
            "method": "DELETE",
            "url": "/tasks/{task_id}/structured-contents/{content_id}",
        },
        "delete_task_sub_step_comment": {
            "method": "DELETE",
            "url": "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}",
        },
        "get_task": {"method": "GET", "url": "/tasks/{id}"},
        "get_task_article": {
            "method": "GET",
            "url": "/tasks/{task_id}/articles/{article_id}",
        },
        "get_task_brief": {"method": "GET", "url": "/tasks/{id}/brief"},
        "get_task_custom_field": {
            "method": "GET",
            "url": "/tasks/{task_id}/custom-fields/{custom_field_id}",
        },
        "get_task_image": {
            "method": "GET",
            "url": "/tasks/{task_id}/images/{image_id}",
        },
        "get_task_raw_file": {
            "method": "GET",
            "url": "/tasks/{task_id}/raw-files/{raw_file_id}",
        },
        "get_task_structured_content": {
            "method": "GET",
            "url": "/tasks/{task_id}/structured-contents/{content_id}",
        },
        "get_task_sub_step": {
            "method": "GET",
            "url": "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}",
        },
        "get_task_sub_step_comment": {
            "method": "GET",
            "url": "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}",
        },
        "get_task_sub_step_external_work": {
            "method": "GET",
            "url": "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/external-work",
        },
        "get_task_video": {
            "method": "GET",
            "url": "/tasks/{task_id}/videos/{video_id}",
        },
        "list_task_asset_comments": {
            "method": "GET",
            "url": "/tasks/{task_id}/assets/{asset_id}/comments",
        },
        "list_task_asset_drafts": {
            "method": "GET",
            "url": "/tasks/{task_id}/assets/{asset_id}/drafts",
        },
        "list_task_asset_fields": {
            "method": "GET",
            "url": "/tasks/{task_id}/assets/{asset_id}/fields",
        },
        "list_task_assets": {"method": "GET", "url": "/tasks/{id}/assets"},
        "list_task_attachments": {"method": "GET", "url": "/tasks/{id}/attachments"},
        "list_task_custom_field_choices": {
            "method": "GET",
            "url": "/tasks/{task_id}/custom-fields/{custom_field_id}/choices",
        },
        "list_task_custom_fields": {
            "method": "GET",
            "url": "/tasks/{id}/custom-fields",
        },
        "list_task_fields": {"method": "GET", "url": "/tasks/{task_id}/fields"},
        "list_task_sub_step_comments": {
            "method": "GET",
            "url": "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments",
        },
        "list_task_sub_step_fields": {
            "method": "GET",
            "url": "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/fields",
        },
        "list_tasks": {"method": "GET", "url": "/tasks"},
        "remove_task_field": {
            "method": "DELETE",
            "url": "/tasks/{task_id}/fields/{field_id}",
        },
        "update_task": {"method": "PATCH", "url": "/tasks/{id}"},
        "update_task_asset_fields": {
            "method": "PUT",
            "url": "/tasks/{task_id}/assets/{asset_id}/fields",
        },
        "update_task_custom_field": {
            "method": "PATCH",
            "url": "/tasks/{task_id}/custom-fields/{custom_field_id}",
        },
        "update_task_field": {
            "method": "PUT",
            "url": "/tasks/{task_id}/fields/{field_id}",
        },
        "update_task_structured_content": {
            "method": "PATCH",
            "url": "/tasks/{task_id}/structured-contents/{content_id}",
        },
        "update_task_sub_step": {
            "method": "PATCH",
            "url": "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}",
        },
        "update_task_sub_step_comment": {
            "method": "PATCH",
            "url": "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}",
        },
        "update_task_sub_step_external_work": {
            "method": "PATCH",
            "url": "/tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/external-work",
        },
    },
    "task_step": {
        "update_task_step": {
            "method": "PATCH",
            "url": "/tasks/{task_id}/steps/{step_id}",
        }
    },
    "team": {
        "get_team": {"method": "GET", "url": "/teams/{id}"},
        "list_teams": {"method": "GET", "url": "/teams"},
    },
    "template": {
        "get_template": {"method": "GET", "url": "/templates/{template_id}"},
        "list_templates": {"method": "GET", "url": "/templates"},
    },
    "uploader": {
        "complete_multipart_upload": {
            "method": "POST",
            "url": "/multipart-uploads/{id}/complete",
        },
        "create_multipart_upload": {"method": "POST", "url": "/multipart-uploads"},
        "get_multipart_upload_status": {
            "method": "GET",
            "url": "/multipart-uploads/{id}/status",
        },
        "get_upload_url": {"method": "GET", "url": "/upload-url"},
    },
    "user": {
        "find_user_by_email": {"method": "GET", "url": "/users"},
        "get_user": {"method": "GET", "url": "/users/{id}"},
        "list_users": {"method": "GET", "url": "/userlist"},
    },
    "work_request": {
        "add_attachment_to_work_request": {
            "method": "POST",
            "url": "/work-requests/{id}/attachments",
        },
        "add_comment_to_work_request": {
            "method": "POST",
            "url": "/work-requests/{id}/comments",
        },
        "create_campaign_from_work_request": {
            "method": "POST",
            "url": "/work-requests/{id}/campaigns",
        },
        "create_task_from_work_request": {
            "method": "POST",
            "url": "/work-requests/{id}/tasks",
        },
        "create_work_request": {"method": "POST", "url": "/work-requests"},
        "create_work_request_creative_asset": {
            "method": "POST",
            "url": "/work-requests/{id}/creative-assets",
        },
        "delete_work_request_attachment": {
            "method": "DELETE",
            "url": "/work-requests/{work_request_id}/attachments/{attachment_id}",
        },
        "delete_work_request_creative_asset": {
            "method": "DELETE",
            "url": "/work-requests/{work_request_id}/creative-assets/{creative_asset_id}",
        },
        "get_work_request": {"method": "GET", "url": "/work-requests/{id}"},
        "get_work_request_comment": {
            "method": "GET",
            "url": "/work-requests/{work_request_id}/comments/{comment_id}",
        },
        "list_work_request_approved_assets": {
            "method": "GET",
            "url": "/work-requests/{id}/approved-assets",
        },
        "list_work_request_comments": {
            "method": "GET",
            "url": "/work-requests/{id}/comments",
        },
        "list_work_request_related_resources": {
            "method": "GET",
            "url": "/work-requests/{id}/related-resources",
        },
        "list_work_requests": {"method": "GET", "url": "/work-requests"},
        "update_work_request": {"method": "PATCH", "url": "/work-requests/{id}"},
        "update_work_request_form_field": {
            "method": "PUT",
            "url": "/work-requests/{work_request_id}/form-fields/{form_field_identifier}",
        },
    },
    "workflow": {
        "get_workflow": {"method": "GET", "url": "/workflows/{workflow_id}"},
        "list_workflows": {"method": "GET", "url": "/workflows"},
    },
}
