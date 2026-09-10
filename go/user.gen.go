// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type UserService struct{ client *Client }

type FindUserByEmailParams struct {
	Email string
}

// FindUserByEmail GET /users
func (s *UserService) FindUserByEmail(ctx context.Context, params FindUserByEmailParams) (*Response[any], error) {
	r := &request{method: "GET", path: "/users"}
	q := url.Values{}
	q.Set("email", fmt.Sprint(params.Email))
	if len(q) > 0 {
		r.query = q
	}
	return do[any](ctx, s.client, r)
}

type GetUserParams struct {
	ID string
}

// GetUser GET /users/{id}
func (s *UserService) GetUser(ctx context.Context, params GetUserParams) (*Response[schema.UserResponse], error) {
	r := &request{method: "GET", path: "/users/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.UserResponse](ctx, s.client, r)
}

// ListUsers GET /userlist
func (s *UserService) ListUsers(ctx context.Context) (*Response[schema.UserListResponse], error) {
	r := &request{method: "GET", path: "/userlist"}
	return do[schema.UserListResponse](ctx, s.client, r)
}
