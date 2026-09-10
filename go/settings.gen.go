// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type SettingsService struct{ client *Client }

// GetSettings GET /settings
func (s *SettingsService) GetSettings(ctx context.Context) (*Response[schema.Settings], error) {
	r := &request{method: "GET", path: "/settings"}
	return do[schema.Settings](ctx, s.client, r)
}

type UpdateSettingsParams struct {
	// If `execute=true` the settings are created or updated. Otherwise, the endpoint returns only the changeset.
	Execute *bool
	// If `overwrite_workflows=true` the existing workflows are overwritten. Otherwise, a new workflow is created where a prefix `Copy of` is added to the workflow's name.
	OverwriteWorkflows *bool
	Body               schema.SettingsResources
}

// UpdateSettings POST /settings
func (s *SettingsService) UpdateSettings(ctx context.Context, params UpdateSettingsParams) (*Response[schema.SettingsUpdateResponse], error) {
	r := &request{method: "POST", path: "/settings"}
	q := url.Values{}
	if params.Execute != nil {
		q.Set("execute", fmt.Sprint(*params.Execute))
	}
	if params.OverwriteWorkflows != nil {
		q.Set("overwrite_workflows", fmt.Sprint(*params.OverwriteWorkflows))
	}
	if len(q) > 0 {
		r.query = q
	}
	r.body = params.Body
	return do[schema.SettingsUpdateResponse](ctx, s.client, r)
}
