package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type BasicRequest struct {
	httpRequest *http.Request
	err error
}

func NewRequest(method string, addr string) *BasicRequest {
	newUrl, err := url.Parse(addr)

	if err != nil {
		return &BasicRequest{
			nil,
			fmt.Errorf("http.NewRequest: %w", err),
		}
	}

	req := &http.Request{
		Method: method,
		URL: newUrl,
	}

	req.Header = make(http.Header)
	req.Header.Set("Content-Type", "application/json")

	return &BasicRequest{req, nil}
}

func (r *BasicRequest) Http() *http.Request {
	return r.httpRequest
}

func (r *BasicRequest) Err() error {
	return r.err
}

func (r *BasicRequest) WithBody(body any) *BasicRequest {
	if r.err != nil {
		return r
	}

	data, err := json.Marshal(body)

	if err != nil {
		r.err = fmt.Errorf("BasicRequest.WithBody: %s", err)
		return r
	}

	r.httpRequest.Body = io.NopCloser(bytes.NewReader(data))

	return r
}

func (r *BasicRequest) WithApiKeyAuthorization(key string) *BasicRequest {
	if r.err != nil {
		return r
	}

	r.httpRequest.Header.Set("Authorization", fmt.Sprintf("Api-key %s", key))

	return r
}

func (r *BasicRequest) WithContext(ctx context.Context) *BasicRequest {
	if r.err != nil {
		return r
	}

	r.httpRequest.WithContext(ctx)

	return r
}
