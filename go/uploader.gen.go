// Auto-generated - DO NOT EDIT

package opticmp

import (
	"context"
	"fmt"
	"github.com/newscred/opti-cmp-sdk/go/schema"
)

type UploaderService struct{ client *Client }

type CompleteMultipartUploadParams struct {
	// ID of the multipart upload to complete
	ID string
}

// CompleteMultipartUpload POST /v3/multipart-uploads/{id}/complete
func (s *UploaderService) CompleteMultipartUpload(ctx context.Context, params CompleteMultipartUploadParams) (*Response[schema.CompleteMultipartUploadResponse], error) {
	r := &request{method: "POST", path: "/multipart-uploads/{id}/complete"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.CompleteMultipartUploadResponse](ctx, s.client, r)
}

type CreateMultipartUploadParams struct {
	Body schema.CreateMultipartUploadRequest
}

// CreateMultipartUpload POST /v3/multipart-uploads
func (s *UploaderService) CreateMultipartUpload(ctx context.Context, params CreateMultipartUploadParams) (*Response[schema.CreateMultipartUploadResponse], error) {
	r := &request{method: "POST", path: "/multipart-uploads"}
	r.body = params.Body
	return do[schema.CreateMultipartUploadResponse](ctx, s.client, r)
}

type GetMultipartUploadStatusParams struct {
	// ID of the multipart upload to check status for
	ID string
}

// GetMultipartUploadStatus GET /v3/multipart-uploads/{id}/status
func (s *UploaderService) GetMultipartUploadStatus(ctx context.Context, params GetMultipartUploadStatusParams) (*Response[schema.GetMultipartUploadStatusResponse], error) {
	r := &request{method: "GET", path: "/multipart-uploads/{id}/status"}
	r.pathParams = map[string]string{
		"id": fmt.Sprint(params.ID),
	}
	return do[schema.GetMultipartUploadStatusResponse](ctx, s.client, r)
}

// GetUploadURL GET /upload-url
func (s *UploaderService) GetUploadURL(ctx context.Context) (*Response[schema.GetUploadURLResponse], error) {
	r := &request{method: "GET", path: "/upload-url"}
	return do[schema.GetUploadURLResponse](ctx, s.client, r)
}
