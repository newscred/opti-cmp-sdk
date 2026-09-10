// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type TemplateService struct{ client *Client }

type GetTemplateParams struct {
	TemplateID string
}

// GetTemplate GET /templates/{template_id}
func (s *TemplateService) GetTemplate(ctx context.Context, params GetTemplateParams) (*Response[schema.TemplateResponse], error) {
	r := &request{method: "GET", path: "/templates/{template_id}"}
	r.pathParams = map[string]string{
		"template_id": fmt.Sprint(params.TemplateID),
	}
	return do[schema.TemplateResponse](ctx, s.client, r)
}

type ListTemplatesParams struct {
	Search          *string
	ApplicableTo    *string
	IncludeInactive *bool
}

// ListTemplates GET /templates
func (s *TemplateService) ListTemplates(ctx context.Context, params ListTemplatesParams) (*Response[schema.TemplateListResponse], error) {
	r := &request{method: "GET", path: "/templates"}
	q := url.Values{}
	if params.Search != nil {
		q.Set("search", fmt.Sprint(*params.Search))
	}
	if params.ApplicableTo != nil {
		q.Set("applicable_to", fmt.Sprint(*params.ApplicableTo))
	}
	if params.IncludeInactive != nil {
		q.Set("include_inactive", fmt.Sprint(*params.IncludeInactive))
	}
	if len(q) > 0 {
		r.query = q
	}
	return do[schema.TemplateListResponse](ctx, s.client, r)
}
