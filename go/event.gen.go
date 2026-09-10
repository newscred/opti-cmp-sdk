// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type EventService struct{ client *Client }

type CreateEventParams struct {
	Body schema.EventCreateRequest
}

// CreateEvent POST /events
func (s *EventService) CreateEvent(ctx context.Context, params CreateEventParams) (*Response[schema.EventResponse], error) {
	r := &request{method: "POST", path: "/events"}
	r.body = params.Body
	return do[schema.EventResponse](ctx, s.client, r)
}

type GetEventParams struct {
	ID string
}

// GetEvent GET /events/{id}
func (s *EventService) GetEvent(ctx context.Context, params GetEventParams) (*Response[schema.EventResponse], error) {
	r := &request{method: "GET", path: "/events/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.EventResponse](ctx, s.client, r)
}

type ListEventFieldsParams struct {
	ID       string
	Offset   *int
	PageSize *int
}

// ListEventFields GET events/{id}/fields
func (s *EventService) ListEventFields(ctx context.Context, params ListEventFieldsParams) (*Response[schema.ListEventFieldsResponse], error) {
	r := &request{method: "GET", path: "/events/{id}/fields"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
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
	return do[schema.ListEventFieldsResponse](ctx, s.client, r)
}

type ListEventsParams struct {
	StartDate  *string
	EndDate    *string
	CampaignID *string
	Offset     *int
	PageSize   *int
}

// ListEvents GET /events
func (s *EventService) ListEvents(ctx context.Context, params ListEventsParams) (*Response[schema.ListEventsResponse], error) {
	r := &request{method: "GET", path: "/events"}
	q := url.Values{}
	if params.StartDate != nil {
		q.Set("start_date", fmt.Sprint(*params.StartDate))
	}
	if params.EndDate != nil {
		q.Set("end_date", fmt.Sprint(*params.EndDate))
	}
	if params.CampaignID != nil {
		q.Set("campaign_id", fmt.Sprint(*params.CampaignID))
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
	return do[schema.ListEventsResponse](ctx, s.client, r)
}

type UpdateEventParams struct {
	ID   string
	Body schema.EventUpdateRequest
}

// UpdateEvent PATCH /events/{id}
func (s *EventService) UpdateEvent(ctx context.Context, params UpdateEventParams) (*Response[schema.EventResponse], error) {
	r := &request{method: "PATCH", path: "/events/{id}"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[schema.EventResponse](ctx, s.client, r)
}

type UpdateEventFieldsParams struct {
	ID   string
	Body schema.EventFieldsUpdateRequest
}

// UpdateEventFields PUT /events/{id}/fields
func (s *EventService) UpdateEventFields(ctx context.Context, params UpdateEventFieldsParams) (*Response[any], error) {
	r := &request{method: "PUT", path: "/events/{id}/fields"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	r.body = params.Body
	return do[any](ctx, s.client, r)
}
