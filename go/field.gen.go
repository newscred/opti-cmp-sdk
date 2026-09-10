// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type FieldService struct{ client *Client }

type CreateFieldParams struct {
	Body schema.SettingsFieldCreateRequest
}

// CreateField POST /fields
func (s *FieldService) CreateField(ctx context.Context, params CreateFieldParams) (*Response[schema.CreateFieldResponse], error) {
	r := &request{method: "POST", path: "/fields"}
	r.body = params.Body
	return do[schema.CreateFieldResponse](ctx, s.client, r)
}

type CreateFieldChoicesParams struct {
	ID   string
	Body schema.SettingsFieldChoiceCreateRequest
}

// CreateFieldChoices POST /fields/{id}/choices
func (s *FieldService) CreateFieldChoices(ctx context.Context, params CreateFieldChoicesParams) (*Response[schema.SettingsFieldChoiceCreateResponse], error) {
	r := &request{method: "POST", path: "/fields/{id}/choices"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.SettingsFieldChoiceCreateResponse](ctx, s.client, r)
}

type DeleteFieldChoiceParams struct {
	FieldID  string
	ChoiceID string
}

// DeleteFieldChoice DELETE /fields/{field_id}/choices/{choice_id}
func (s *FieldService) DeleteFieldChoice(ctx context.Context, params DeleteFieldChoiceParams) (*Response[any], error) {
	r := &request{method: "DELETE", path: "/fields/{field_id}/choices/{choice_id}"}
	r.pathParams = map[string]string{
		"field_id":  fmt.Sprint(params.FieldID),
		"choice_id": fmt.Sprint(params.ChoiceID),
	}
	return do[any](ctx, s.client, r)
}

type ListFieldsParams struct {
	Ids      *string
	Offset   *int
	PageSize *int
}

// ListFields GET /fields
func (s *FieldService) ListFields(ctx context.Context, params ListFieldsParams) (*Response[schema.ListFieldsResponse], error) {
	r := &request{method: "GET", path: "/fields"}
	q := url.Values{}
	if params.Ids != nil {
		q.Set("ids", fmt.Sprint(*params.Ids))
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
	return do[schema.ListFieldsResponse](ctx, s.client, r)
}

type UpdateFieldParams struct {
	ID   string
	Body schema.SettingsFieldUpdateRequest
}

// UpdateField PATCH /fields/{id}
func (s *FieldService) UpdateField(ctx context.Context, params UpdateFieldParams) (*Response[any], error) {
	r := &request{method: "PATCH", path: "/fields/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[any](ctx, s.client, r)
}

type UpdateFieldChoiceParams struct {
	FieldID  string
	ChoiceID string
	Body     schema.SettingsFieldChoiceUpdateRequest
}

// UpdateFieldChoice PATCH /fields/{field_id}/choices/{choice_id}
func (s *FieldService) UpdateFieldChoice(ctx context.Context, params UpdateFieldChoiceParams) (*Response[any], error) {
	r := &request{method: "PATCH", path: "/fields/{field_id}/choices/{choice_id}"}
	r.pathParams = map[string]string{
		"field_id":  fmt.Sprint(params.FieldID),
		"choice_id": fmt.Sprint(params.ChoiceID),
	}
	r.body = params.Body
	return do[any](ctx, s.client, r)
}
