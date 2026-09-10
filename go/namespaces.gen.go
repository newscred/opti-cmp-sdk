// Auto-generated - DO NOT EDIT

package opticmp

// namespaces holds the endpoint namespace fields, embedded into Client.
type namespaces struct {
	Asset             *AssetService
	BrandCompliance   *BrandComplianceService
	Campaign          *CampaignService
	ContentGraph      *ContentGraphService
	Event             *EventService
	Field             *FieldService
	Label             *LabelService
	Library           *LibraryService
	Milestone         *MilestoneService
	Publishing        *PublishingService
	Settings          *SettingsService
	StructuredContent *StructuredContentService
	Task              *TaskService
	TaskStep          *TaskStepService
	Team              *TeamService
	Template          *TemplateService
	Uploader          *UploaderService
	User              *UserService
	WorkRequest       *WorkRequestService
	Workflow          *WorkflowService
}

func (c *Client) initNamespaces() {
	c.Asset = &AssetService{client: c}
	c.BrandCompliance = &BrandComplianceService{client: c}
	c.Campaign = &CampaignService{client: c}
	c.ContentGraph = &ContentGraphService{client: c}
	c.Event = &EventService{client: c}
	c.Field = &FieldService{client: c}
	c.Label = &LabelService{client: c}
	c.Library = &LibraryService{client: c}
	c.Milestone = &MilestoneService{client: c}
	c.Publishing = &PublishingService{client: c}
	c.Settings = &SettingsService{client: c}
	c.StructuredContent = &StructuredContentService{client: c}
	c.Task = &TaskService{client: c}
	c.TaskStep = &TaskStepService{client: c}
	c.Team = &TeamService{client: c}
	c.Template = &TemplateService{client: c}
	c.Uploader = &UploaderService{client: c}
	c.User = &UserService{client: c}
	c.WorkRequest = &WorkRequestService{client: c}
	c.Workflow = &WorkflowService{client: c}
}
