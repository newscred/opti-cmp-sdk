# Auto-generated - DO NOT EDIT

from __future__ import annotations

from typing import Any

from .asset import AssetNamespace
from .brand_compliance import BrandComplianceNamespace
from .campaign import CampaignNamespace
from .event import EventNamespace
from .field import FieldNamespace
from .label import LabelNamespace
from .library import LibraryNamespace
from .milestone import MilestoneNamespace
from .publishing import PublishingNamespace
from .settings import SettingsNamespace
from .structured_content import StructuredContentNamespace
from .task import TaskNamespace
from .task_step import TaskStepNamespace
from .team import TeamNamespace
from .template import TemplateNamespace
from .uploader import UploaderNamespace
from .user import UserNamespace
from .work_request import WorkRequestNamespace
from .workflow import WorkflowNamespace

NAMESPACES: dict[str, type[Any]] = {
    "asset": AssetNamespace,
    "brand_compliance": BrandComplianceNamespace,
    "campaign": CampaignNamespace,
    "event": EventNamespace,
    "field": FieldNamespace,
    "label": LabelNamespace,
    "library": LibraryNamespace,
    "milestone": MilestoneNamespace,
    "publishing": PublishingNamespace,
    "settings": SettingsNamespace,
    "structured_content": StructuredContentNamespace,
    "task": TaskNamespace,
    "task_step": TaskStepNamespace,
    "team": TeamNamespace,
    "template": TemplateNamespace,
    "uploader": UploaderNamespace,
    "user": UserNamespace,
    "workflow": WorkflowNamespace,
    "work_request": WorkRequestNamespace,
}


class AsyncNamespaces:
    """Namespaces attached by `register_api_endpoints_plugin`."""

    asset: AssetNamespace
    brand_compliance: BrandComplianceNamespace
    campaign: CampaignNamespace
    event: EventNamespace
    field: FieldNamespace
    label: LabelNamespace
    library: LibraryNamespace
    milestone: MilestoneNamespace
    publishing: PublishingNamespace
    settings: SettingsNamespace
    structured_content: StructuredContentNamespace
    task: TaskNamespace
    task_step: TaskStepNamespace
    team: TeamNamespace
    template: TemplateNamespace
    uploader: UploaderNamespace
    user: UserNamespace
    workflow: WorkflowNamespace
    work_request: WorkRequestNamespace


__all__ = [
    "NAMESPACES",
    "AssetNamespace",
    "AsyncNamespaces",
    "BrandComplianceNamespace",
    "CampaignNamespace",
    "EventNamespace",
    "FieldNamespace",
    "LabelNamespace",
    "LibraryNamespace",
    "MilestoneNamespace",
    "PublishingNamespace",
    "SettingsNamespace",
    "StructuredContentNamespace",
    "TaskNamespace",
    "TaskStepNamespace",
    "TeamNamespace",
    "TemplateNamespace",
    "UploaderNamespace",
    "UserNamespace",
    "WorkRequestNamespace",
    "WorkflowNamespace",
]
