// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type LabelService struct{ client *Client }

type ListLabelGroupsParams struct {
	// Source organization type to filter by
	SourceOrgType *string
	Offset        *int
	PageSize      *int
}

// ListLabelGroups GET /label-groups
func (s *LabelService) ListLabelGroups(ctx context.Context, params ListLabelGroupsParams) (*Response[schema.ListLabelGroupsResponse], error) {
	r := &request{method: "GET", path: "/label-groups"}
	q := url.Values{}
	if params.SourceOrgType != nil {
		q.Set("source_org_type", fmt.Sprint(*params.SourceOrgType))
	}
	if params.Offset != nil {
		q.Set("offset", fmt.Sprint(*params.Offset))
	}
	if params.PageSize != nil {
		q.Set("page_size", fmt.Sprint(*params.PageSize))
	}
	if len(q) > 0 {
		r.query = q
	}
	return do[schema.ListLabelGroupsResponse](ctx, s.client, r)
}
