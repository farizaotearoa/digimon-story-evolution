package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// HTTPClient is a wrapper around http.Client with additional functionality
type HTTPClient struct {
	client  *http.Client
	baseURL string
}

// NewHTTPClient creates a new HTTP client for microservice communication
func NewHTTPClient(baseURL string, timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
	}
}

// Request represents an HTTP request
type Request struct {
	Method  string
	Path    string
	Body    interface{}
	Headers map[string]string
}

// Response represents an HTTP response
type Response struct {
	StatusCode int
	Body       []byte
	Headers    http.Header
}

// Do performs an HTTP request
func (h *HTTPClient) Do(ctx context.Context, req Request) (*Response, error) {
	var body io.Reader
	if req.Body != nil {
		jsonBody, err := json.Marshal(req.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		body = bytes.NewBuffer(jsonBody)
	}

	httpReq, err := http.NewRequestWithContext(ctx, req.Method, h.baseURL+req.Path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set default headers
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("User-Agent", "digimon-api-gateway/1.0")

	// Set custom headers
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	// Add request ID if available
	if requestID := ctx.Value("request_id"); requestID != nil {
		httpReq.Header.Set("X-Request-ID", requestID.(string))
	}

	Logger.Debug("Making HTTP request",
		zap.String("method", req.Method),
		zap.String("url", httpReq.URL.String()),
		zap.String("request_id", httpReq.Header.Get("X-Request-ID")),
	)

	resp, err := h.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	Logger.Debug("Received HTTP response",
		zap.String("method", req.Method),
		zap.String("url", httpReq.URL.String()),
		zap.Int("status_code", resp.StatusCode),
		zap.String("request_id", httpReq.Header.Get("X-Request-ID")),
	)

	return &Response{
		StatusCode: resp.StatusCode,
		Body:       respBody,
		Headers:    resp.Header,
	}, nil
}

// Get performs a GET request
func (h *HTTPClient) Get(ctx context.Context, path string, headers map[string]string) (*Response, error) {
	return h.Do(ctx, Request{
		Method:  "GET",
		Path:    path,
		Headers: headers,
	})
}

// Post performs a POST request
func (h *HTTPClient) Post(ctx context.Context, path string, body interface{}, headers map[string]string) (*Response, error) {
	return h.Do(ctx, Request{
		Method:  "POST",
		Path:    path,
		Body:    body,
		Headers: headers,
	})
}

// Put performs a PUT request
func (h *HTTPClient) Put(ctx context.Context, path string, body interface{}, headers map[string]string) (*Response, error) {
	return h.Do(ctx, Request{
		Method:  "PUT",
		Path:    path,
		Body:    body,
		Headers: headers,
	})
}

// Delete performs a DELETE request
func (h *HTTPClient) Delete(ctx context.Context, path string, headers map[string]string) (*Response, error) {
	return h.Do(ctx, Request{
		Method:  "DELETE",
		Path:    path,
		Headers: headers,
	})
}
