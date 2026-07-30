# Auto-generated - DO NOT EDIT

from __future__ import annotations

from collections.abc import Mapping
from typing import TYPE_CHECKING, Literal

from ..._types import AsyncDispatcher, RequestConfig, Response

if TYPE_CHECKING:
    from .. import objects, schema


class TaskNamespace:
    """`task` endpoints."""

    def __init__(self, client: AsyncDispatcher) -> None:
        self._client = client

    async def add_asset_to_task(
        self,
        id: str,
        *,
        body: schema.TaskAssetRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskAssetResponse]:
        """This API allows you to add an asset to a task. There currently two mechanisms to add an asset - either by uploading an asset first and then calling this API (for only `images`, `videos`, and `raw files`), or by adding an asset from the library. The 'type' parameter in the request body will determine the mechanism used to add the asset.

        POST /tasks/{id}/assets
        """
        return await self._client._dispatch(
            namespace="task",
            operation="add_asset_to_task",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def add_comment_to_task(
        self,
        task_id: str,
        *,
        body: schema.CommentCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskCommentResponse]:
        """Post a comment on a task.

        POST /tasks/{task_id}/comments

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="add_comment_to_task",
            params={
                "task_id": task_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def add_comment_to_task_asset(
        self,
        task_id: str,
        asset_id: str,
        *,
        body: schema.CommentCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskAssetCommentResponse]:
        """Add a comment to the latest version of a task asset. Supports adding comments to only **images**, **videos** and **raw files** type assets.

        POST /tasks/{task_id}/assets/{asset_id}/comments
        """
        return await self._client._dispatch(
            namespace="task",
            operation="add_comment_to_task_asset",
            params={
                "task_id": task_id,
                "asset_id": asset_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def add_comment_to_task_sub_step(
        self,
        task_id: str,
        step_id: str,
        sub_step_id: str,
        *,
        body: schema.CommentWithReplyCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskSubStepCommentResponse]:
        """Create a task substep comment

        POST /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments
        """
        return await self._client._dispatch(
            namespace="task",
            operation="add_comment_to_task_sub_step",
            params={
                "task_id": task_id,
                "step_id": step_id,
                "sub_step_id": sub_step_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def add_draft_to_task_asset(
        self,
        task_id: str,
        asset_id: str,
        *,
        body: schema.TaskAssetRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskAssetDraftResponse]:
        """Add a new draft to a task asset. You can add a draft to only `images`, `videos`, and `raw files` type task assets. See [Upload assets](https://docs.developers.optimizely.com/content-marketing-platform/docs/upload-assets) to upload an asset draft.

        POST /tasks/{task_id}/assets/{asset_id}/drafts
        """
        return await self._client._dispatch(
            namespace="task",
            operation="add_draft_to_task_asset",
            params={
                "task_id": task_id,
                "asset_id": asset_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def add_field_to_task(
        self,
        task_id: str,
        *,
        body: schema.ObjectFieldCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ObjectFieldCreateResponse]:
        """Add a field to a task.

        POST tasks/{task_id}/fields

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="add_field_to_task",
            params={
                "task_id": task_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def add_structured_content_to_task(
        self,
        task_id: str,
        *,
        body: schema.TaskStructuredContentCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ContentDetailsModel]:
        """Add a new structured content to task

        POST /tasks/{task_id}/structured-contents

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="add_structured_content_to_task",
            params={
                "task_id": task_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def add_url_to_task(
        self,
        task_id: str,
        *,
        body: schema.AddUrlToTaskRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskUrlResponse]:
        """Adds a URL to a task.

        POST /tasks/{task_id}/urls

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="add_url_to_task",
            params={
                "task_id": task_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def create_task(
        self,
        *,
        body: schema.TaskCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskResponse]:
        """Create a task

        POST /tasks

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="create_task",
            params={
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def create_task_publishing_intent(
        self,
        task_id: str,
        *,
        body: schema.TaskPublishingIntentCreateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskPublishingIntentResponse]:
        """Create a publishing intent for a task on the specified publishing channel.

        POST /tasks/{task_id}/publishing-intents

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="create_task_publishing_intent",
            params={
                "task_id": task_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def create_task_structured_content_draft(
        self,
        task_id: str,
        content_id: str,
        *,
        body: schema.TaskStructuredContentDraftRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.CreateTaskStructuredContentDraftResponse]:
        """Create structured content draft using a task.

        POST /tasks/{task_id}/structured-contents/{content_id}/drafts

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="create_task_structured_content_draft",
            params={
                "task_id": task_id,
                "content_id": content_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def delete_task_structured_content(
        self,
        task_id: str,
        content_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Delete a structured content from task

        DELETE /tasks/{task_id}/structured-contents/{content_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="delete_task_structured_content",
            params={
                "task_id": task_id,
                "content_id": content_id,
                "headers": headers,
                "request": request,
            },
        )

    async def delete_task_sub_step_comment(
        self,
        task_id: str,
        step_id: str,
        sub_step_id: str,
        comment_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Delete a task substep comment

        DELETE /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}
        """
        return await self._client._dispatch(
            namespace="task",
            operation="delete_task_sub_step_comment",
            params={
                "task_id": task_id,
                "step_id": step_id,
                "sub_step_id": sub_step_id,
                "comment_id": comment_id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_task(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskResponse]:
        """Get a task

        GET /tasks/{id}
        """
        return await self._client._dispatch(
            namespace="task",
            operation="get_task",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_task_article(
        self,
        task_id: str,
        article_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskArticle]:
        """Get an article associated with the task identified by `task_id`.

        GET /tasks/{task_id}/articles/{article_id}
        """
        return await self._client._dispatch(
            namespace="task",
            operation="get_task_article",
            params={
                "task_id": task_id,
                "article_id": article_id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_task_brief(
        self,
        id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskBriefResponse]:
        """Get brief of the task

        GET /tasks/{id}/brief
        """
        return await self._client._dispatch(
            namespace="task",
            operation="get_task_brief",
            params={
                "id": id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_task_custom_field(
        self,
        task_id: str,
        custom_field_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskCustomField]:
        """Get a custom field of a task

        GET /tasks/{task_id}/custom-fields/{custom_field_id}
        """
        return await self._client._dispatch(
            namespace="task",
            operation="get_task_custom_field",
            params={
                "task_id": task_id,
                "custom_field_id": custom_field_id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_task_image(
        self,
        task_id: str,
        image_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskImage]:
        """Get an image associated with the task identified by `task_id`.

        GET /tasks/{task_id}/images/{image_id}
        """
        return await self._client._dispatch(
            namespace="task",
            operation="get_task_image",
            params={
                "task_id": task_id,
                "image_id": image_id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_task_raw_file(
        self,
        task_id: str,
        raw_file_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskRawFile]:
        """Get a raw file associated with the task identified by `task_id`.

        GET /tasks/{task_id}/raw-files/{raw_file_id}
        """
        return await self._client._dispatch(
            namespace="task",
            operation="get_task_raw_file",
            params={
                "task_id": task_id,
                "raw_file_id": raw_file_id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_task_structured_content(
        self,
        task_id: str,
        content_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ContentDetailsModel]:
        """Get a task structured content by its guid.

        GET /tasks/{task_id}/structured-contents/{content_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="get_task_structured_content",
            params={
                "task_id": task_id,
                "content_id": content_id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_task_sub_step(
        self,
        task_id: str,
        step_id: str,
        sub_step_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskSubStep]:
        """Get the substep of a task

        GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}
        """
        return await self._client._dispatch(
            namespace="task",
            operation="get_task_sub_step",
            params={
                "task_id": task_id,
                "step_id": step_id,
                "sub_step_id": sub_step_id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_task_sub_step_comment(
        self,
        task_id: str,
        step_id: str,
        sub_step_id: str,
        comment_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskSubStepCommentResponse]:
        """Get a task substep comment

        GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}
        """
        return await self._client._dispatch(
            namespace="task",
            operation="get_task_sub_step_comment",
            params={
                "task_id": task_id,
                "step_id": step_id,
                "sub_step_id": sub_step_id,
                "comment_id": comment_id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_task_sub_step_external_work(
        self,
        task_id: str,
        step_id: str,
        sub_step_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskExternalWorkResponse]:
        """Get the external work information of a task external substep

        GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/external-work
        """
        return await self._client._dispatch(
            namespace="task",
            operation="get_task_sub_step_external_work",
            params={
                "task_id": task_id,
                "step_id": step_id,
                "sub_step_id": sub_step_id,
                "headers": headers,
                "request": request,
            },
        )

    async def get_task_video(
        self,
        task_id: str,
        video_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskVideo]:
        """Get a video associated with the task identified by task_id

        GET /tasks/{task_id}/videos/{video_id}
        """
        return await self._client._dispatch(
            namespace="task",
            operation="get_task_video",
            params={
                "task_id": task_id,
                "video_id": video_id,
                "headers": headers,
                "request": request,
            },
        )

    async def list_task_asset_comments(
        self,
        task_id: str,
        asset_id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListTaskAssetCommentsResponse]:
        """List all the comments for the latest version of a task asset.

        GET /tasks/{task_id}/assets/{asset_id}/comments
        """
        return await self._client._dispatch(
            namespace="task",
            operation="list_task_asset_comments",
            params={
                "task_id": task_id,
                "asset_id": asset_id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def list_task_asset_drafts(
        self,
        task_id: str,
        asset_id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListTaskAssetDraftsResponse]:
        """Get the list of drafts on an asset of a task.

        GET /tasks/{task_id}/assets/{asset_id}/drafts

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="list_task_asset_drafts",
            params={
                "task_id": task_id,
                "asset_id": asset_id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def list_task_asset_fields(
        self,
        task_id: str,
        asset_id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListTaskAssetFieldsResponse]:
        """List all the fields of a task asset.

        GET /tasks/{task_id}/assets/{asset_id}/fields

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="list_task_asset_fields",
            params={
                "task_id": task_id,
                "asset_id": asset_id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def list_task_assets(
        self,
        id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListTaskAssetsResponse]:
        """Get the list of the assets or contents of a task

        GET /tasks/{id}/assets
        """
        return await self._client._dispatch(
            namespace="task",
            operation="list_task_assets",
            params={
                "id": id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def list_task_attachments(
        self,
        id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListTaskAttachmentsResponse]:
        """Get the list of the attachments of a task

        GET /tasks/{id}/attachments
        """
        return await self._client._dispatch(
            namespace="task",
            operation="list_task_attachments",
            params={
                "id": id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def list_task_custom_field_choices(
        self,
        task_id: str,
        custom_field_id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListTaskCustomFieldChoicesResponse]:
        """Get the list of the choices of a custom field in a task.

        GET /tasks/{task_id}/custom-fields/{custom_field_id}/choices
        """
        return await self._client._dispatch(
            namespace="task",
            operation="list_task_custom_field_choices",
            params={
                "task_id": task_id,
                "custom_field_id": custom_field_id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def list_task_custom_fields(
        self,
        id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListTaskCustomFieldsResponse]:
        """Get the list of custom fields added to a task.

        GET tasks/{id}/custom-fields
        """
        return await self._client._dispatch(
            namespace="task",
            operation="list_task_custom_fields",
            params={
                "id": id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def list_task_fields(
        self,
        task_id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListTaskFieldsResponse]:
        """Get the list of fields of a task.

        GET tasks/{task_id}/fields

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="list_task_fields",
            params={
                "task_id": task_id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def list_task_sub_step_comments(
        self,
        task_id: str,
        step_id: str,
        sub_step_id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListTaskSubStepCommentsResponse]:
        """Get the list of the comments of a task substep

        GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments
        """
        return await self._client._dispatch(
            namespace="task",
            operation="list_task_sub_step_comments",
            params={
                "task_id": task_id,
                "step_id": step_id,
                "sub_step_id": sub_step_id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def list_task_sub_step_fields(
        self,
        task_id: str,
        step_id: str,
        sub_step_id: str,
        *,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListTaskSubStepFieldsResponse]:
        """Get the list of fields of a substep of a task.

        GET /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/fields

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="list_task_sub_step_fields",
            params={
                "task_id": task_id,
                "step_id": step_id,
                "sub_step_id": sub_step_id,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def list_tasks(
        self,
        *,
        search_key: str | None = None,
        campaign: str | None = None,
        workflow: str | None = None,
        milestone: str | None = None,
        start_date: str | None = None,
        due_date: str | None = None,
        status: list[
            Literal[
                "Archived",
                "Completed",
                "Overdue",
                "Not Started",
                "In Progress",
                "On Hold",
            ]
        ]
        | None = None,
        offset: int | None = None,
        page_size: int | None = None,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ListTasksResponse]:
        """Get a list of tasks.

        GET /tasks

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="list_tasks",
            params={
                "search_key": search_key,
                "campaign": campaign,
                "workflow": workflow,
                "milestone": milestone,
                "start_date": start_date,
                "due_date": due_date,
                "status": status,
                "offset": offset,
                "page_size": page_size,
                "headers": headers,
                "request": request,
            },
        )

    async def remove_task_field(
        self,
        task_id: str,
        field_id: str,
        *,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Remove a field of a task.

        DELETE tasks/{task_id}/fields/{field_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="remove_task_field",
            params={
                "task_id": task_id,
                "field_id": field_id,
                "headers": headers,
                "request": request,
            },
        )

    async def update_task(
        self,
        id: str,
        *,
        body: schema.TaskUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskResponse]:
        """Update a task

        PATCH /tasks/{id}
        """
        return await self._client._dispatch(
            namespace="task",
            operation="update_task",
            params={
                "id": id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def update_task_asset_fields(
        self,
        task_id: str,
        asset_id: str,
        *,
        body: schema.TaskAssetFieldsUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Replace fields of an asset in a task.

        PUT tasks/{task_id}/assets/{asset_id}/fields

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="update_task_asset_fields",
            params={
                "task_id": task_id,
                "asset_id": asset_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def update_task_custom_field(
        self,
        task_id: str,
        custom_field_id: str,
        *,
        body: schema.TaskCustomFieldUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskCustomField]:
        """Update a custom field of a task. You can update the following types: `text_field`, `multi_line_text_field`, `checkboxes`, `dropdown`, `multi_select_dropdown`, `multiple_choice`, `date_field`, `rich_text_field`

        PATCH /tasks/{task_id}/custom-fields/{custom_field_id}
        """
        return await self._client._dispatch(
            namespace="task",
            operation="update_task_custom_field",
            params={
                "task_id": task_id,
                "custom_field_id": custom_field_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def update_task_field(
        self,
        task_id: str,
        field_id: str,
        *,
        body: schema.TaskFieldUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[None]:
        """Update a field of a task.

        PUT tasks/{taks_id}/fields/{field_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="update_task_field",
            params={
                "task_id": task_id,
                "field_id": field_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def update_task_structured_content(
        self,
        task_id: str,
        content_id: str,
        *,
        body: schema.TaskStructuredContentUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.ContentDetailsModel]:
        """Update structured content using a task.

        PATCH /tasks/{task_id}/structured-contents/{content_id}

        Experimental: this endpoint is experimental upstream.
        """
        return await self._client._dispatch(
            namespace="task",
            operation="update_task_structured_content",
            params={
                "task_id": task_id,
                "content_id": content_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def update_task_sub_step(
        self,
        task_id: str,
        step_id: str,
        sub_step_id: str,
        *,
        body: schema.TaskSubStepRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskSubStep]:
        """Update a substep of a step in a task

        PATCH /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}
        """
        return await self._client._dispatch(
            namespace="task",
            operation="update_task_sub_step",
            params={
                "task_id": task_id,
                "step_id": step_id,
                "sub_step_id": sub_step_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def update_task_sub_step_comment(
        self,
        task_id: str,
        step_id: str,
        sub_step_id: str,
        comment_id: str,
        *,
        body: schema.TaskSubStepCommentUpdateRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskSubStepCommentResponse]:
        """Update a task substep comment

        PATCH /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/comments/{comment_id}
        """
        return await self._client._dispatch(
            namespace="task",
            operation="update_task_sub_step_comment",
            params={
                "task_id": task_id,
                "step_id": step_id,
                "sub_step_id": sub_step_id,
                "comment_id": comment_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )

    async def update_task_sub_step_external_work(
        self,
        task_id: str,
        step_id: str,
        sub_step_id: str,
        *,
        body: schema.TaskExternalWorkRequest,
        headers: Mapping[str, str] | None = None,
        request: RequestConfig | None = None,
    ) -> Response[objects.TaskExternalWorkResponse]:
        """Update the external work information of a task external substep

        PATCH /tasks/{task_id}/steps/{step_id}/sub-steps/{sub_step_id}/external-work
        """
        return await self._client._dispatch(
            namespace="task",
            operation="update_task_sub_step_external_work",
            params={
                "task_id": task_id,
                "step_id": step_id,
                "sub_step_id": sub_step_id,
                "body": body,
                "headers": headers,
                "request": request,
            },
        )
