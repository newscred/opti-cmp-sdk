// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
)

type AssetService struct{ client *Client }

type GetAssetURLParams struct {
	AssetID string
}

// GetAssetURL GET /asset-urls/{asset_id}
func (s *AssetService) GetAssetURL(ctx context.Context, params GetAssetURLParams) (*Response[any], error) {
	r := &request{method: "GET", path: "/asset-urls/{asset_id}"}
	r.pathParams = map[string]string{
		"asset_id": fmt.Sprint(params.AssetID),
	}
	return do[any](ctx, s.client, r)
}
