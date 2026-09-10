// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
	"net/url"
)

type PublishingService struct{ client *Client }

type BulkCreatePublishingEventMetadataParams struct {
	// Unique identifier of the publishing event
	PublishingEventID string
	Body              schema.PublishingEventMetadataBulkCreateRequest
}

// BulkCreatePublishingEventMetadata POST /v3/publishing-events/{publishing_event_id}/publishing-metadata
func (s *PublishingService) BulkCreatePublishingEventMetadata(ctx context.Context, params BulkCreatePublishingEventMetadataParams) (*Response[schema.PublishingEventMetadataBulkCreateResponse], error) {
	r := &request{method: "POST", path: "/publishing-events/{publishing_event_id}/publishing-metadata"}
	r.pathParams = map[string]string{
		"publishing_event_id": fmt.Sprint(params.PublishingEventID),
	}
	r.body = params.Body
	return do[schema.PublishingEventMetadataBulkCreateResponse](ctx, s.client, r)
}

type GetPublishingEventParams struct {
	// Unique identifier of the publishing event
	PublishingEventID string
}

// GetPublishingEvent GET /v3/publishing-events/{publishing_event_id}
func (s *PublishingService) GetPublishingEvent(ctx context.Context, params GetPublishingEventParams) (*Response[schema.PublishingEventResponse], error) {
	r := &request{method: "GET", path: "/publishing-events/{publishing_event_id}"}
	r.pathParams = map[string]string{
		"publishing_event_id": fmt.Sprint(params.PublishingEventID),
	}
	return do[schema.PublishingEventResponse](ctx, s.client, r)
}

type GetPublishingEventAssetMetadataParams struct {
	// Unique identifier of the publishing event
	PublishingEventID string
	// Unique identifier of the asset
	AssetID string
	// Unique identifier of the publishing metadata
	PublishingMetadataID string
}

// GetPublishingEventAssetMetadata GET v3/publishing-events/{publishing_event_id}/assets/{asset_id}/publishing-metadata/{publishing_metadata_id}
func (s *PublishingService) GetPublishingEventAssetMetadata(ctx context.Context, params GetPublishingEventAssetMetadataParams) (*Response[schema.PublishingEventMetadataResponse], error) {
	r := &request{method: "GET", path: "/publishing-events/{publishing_event_id}/assets/{asset_id}/publishing-metadata/{publishing_metadata_id}"}
	r.pathParams = map[string]string{
		"publishing_event_id":    fmt.Sprint(params.PublishingEventID),
		"asset_id":               fmt.Sprint(params.AssetID),
		"publishing_metadata_id": fmt.Sprint(params.PublishingMetadataID),
	}
	return do[schema.PublishingEventMetadataResponse](ctx, s.client, r)
}

type ListPublishingChannelsParams struct {
	// Number of results to return per page
	PageSize *int
}

// ListPublishingChannels GET /publishing-channels
func (s *PublishingService) ListPublishingChannels(ctx context.Context, params ListPublishingChannelsParams) (*Response[schema.PublishingChannelListResponse], error) {
	r := &request{method: "GET", path: "/publishing-channels"}
	q := url.Values{}
	if params.PageSize != nil {
		q.Set("page_size", fmt.Sprint(*params.PageSize))
	}
	if len(q) > 0 {
		r.query = q
	}
	return do[schema.PublishingChannelListResponse](ctx, s.client, r)
}

type ListPublishingEventMetadataParams struct {
	// Unique identifier of the publishing event
	PublishingEventID string
	// Publishing status of the asset
	Status *string
	// Type of asset
	AssetType *string
	// Unique identifier of the asset.
	AssetID *string
	// The locale to which the asset is being published to.
	Locale *string
}

// ListPublishingEventMetadata GET /v3/publishing-events/{publishing_event_id}/publishing-metadata
func (s *PublishingService) ListPublishingEventMetadata(ctx context.Context, params ListPublishingEventMetadataParams) (*Response[schema.PublishingEventMetadataListResponse], error) {
	r := &request{method: "GET", path: "/publishing-events/{publishing_event_id}/publishing-metadata"}
	r.pathParams = map[string]string{
		"publishing_event_id": fmt.Sprint(params.PublishingEventID),
	}
	q := url.Values{}
	if params.Status != nil {
		q.Set("status", fmt.Sprint(*params.Status))
	}
	if params.AssetType != nil {
		q.Set("asset_type", fmt.Sprint(*params.AssetType))
	}
	if params.AssetID != nil {
		q.Set("asset_id", fmt.Sprint(*params.AssetID))
	}
	if params.Locale != nil {
		q.Set("locale", fmt.Sprint(*params.Locale))
	}
	if len(q) > 0 {
		r.query = q
	}
	return do[schema.PublishingEventMetadataListResponse](ctx, s.client, r)
}
