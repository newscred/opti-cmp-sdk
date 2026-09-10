// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type TeamService struct{ client *Client }

type GetTeamParams struct {
	ID string
}

// GetTeam GET /teams/{id}
func (s *TeamService) GetTeam(ctx context.Context, params GetTeamParams) (*Response[schema.TeamWithUsers], error) {
	r := &request{method: "GET", path: "/teams/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.TeamWithUsers](ctx, s.client, r)
}

type ListTeamsParams struct {
	Offset   *int
	PageSize *int
}

// ListTeams GET /teams
func (s *TeamService) ListTeams(ctx context.Context, params ListTeamsParams) (*Response[schema.ListTeamsResponse], error) {
	r := &request{method: "GET", path: "/teams"}
	q := url.Values{}
	if params.Offset != nil {
		q.Set("offset", fmt.Sprint(*params.Offset))
	}
	if params.PageSize != nil {
		q.Set("page_size", fmt.Sprint(*params.PageSize))
	}
	if len(q) > 0 {
		r.query = q
	}
	return do[schema.ListTeamsResponse](ctx, s.client, r)
}
