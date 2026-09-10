package opticmp

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// request is the internal description of a call the generated methods build.
type request struct {
	method     string
	path       string // URL template, e.g. "/assets/{asset_id}/fields", or an absolute URL
	pathParams map[string]string
	query      url.Values
	header     http.Header
	body       any // nil, or a value to JSON-encode
}

// expandPath replaces {name} segments in a URL template with escaped values.
func expandPath(tmpl string, params map[string]string) string {
	if len(params) == 0 {
		return tmpl
	}
	out := tmpl
	for name, value := range params {
		out = strings.ReplaceAll(out, "{"+name+"}", url.PathEscape(value))
	}
	return out
}

func (c *Client) buildURL(r *request) string {
	u := expandPath(r.path, r.pathParams)
	if !strings.HasPrefix(u, "http") {
		u = c.baseURL + u
	}
	if len(r.query) > 0 {
		sep := "?"
		if strings.Contains(u, "?") {
			sep = "&"
		}
		u = u + sep + r.query.Encode()
	}
	return u
}

// do sends a request and decodes the response body into T.
func do[T any](ctx context.Context, c *Client, r *request) (*Response[T], error) {
	u := c.buildURL(r)

	header := c.header.Clone()
	for k, vs := range r.header {
		for _, v := range vs {
			header.Add(k, v)
		}
	}

	var bodyReader io.Reader
	var bodyForErr any
	if r.body != nil {
		encoded, err := json.Marshal(r.body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(encoded)
		bodyForErr = r.body
		header.Set("Content-Type", "application/json; charset=utf-8")
	}

	req, err := http.NewRequestWithContext(ctx, r.method, u, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header = header

	if err := c.auth.apply(ctx, req); err != nil {
		return nil, err
	}

	spec := &RequestSpec{Method: r.method, URL: u, Header: req.Header.Clone(), Body: bodyForErr}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, &HTTPError{Message: err.Error(), Request: spec}
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &HTTPError{Status: resp.StatusCode, Message: err.Error(), Header: resp.Header, Request: spec}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var data any
		if len(raw) > 0 {
			if json.Unmarshal(raw, &data) != nil {
				data = string(raw)
			}
		}
		return nil, &HTTPError{
			Status:  resp.StatusCode,
			Message: resp.Status,
			Data:    data,
			Header:  resp.Header,
			Request: spec,
		}
	}

	out := &Response[T]{
		Header: resp.Header,
		Status: resp.StatusCode,
		URL:    u,
		client: c,
		raw:    raw,
	}
	if resp.Request != nil && resp.Request.URL != nil {
		out.URL = resp.Request.URL.String()
	}
	if resp.StatusCode != http.StatusNoContent && len(raw) > 0 {
		if err := json.Unmarshal(raw, &out.Data); err != nil {
			return nil, err
		}
	}
	return out, nil
}
