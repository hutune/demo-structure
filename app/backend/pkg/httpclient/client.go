package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ServiceClient internal HTTP client for service-to-service calls
type ServiceClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewServiceClient creates a client for calling another service
func NewServiceClient(baseURL string, opts ...Option) *ServiceClient {
	c := &ServiceClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

// Get performs GET request and decodes JSON response into v
func (c *ServiceClient) Get(ctx context.Context, path string, v interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("http %d: %s", resp.StatusCode, string(body))
	}
	if v != nil {
		return json.NewDecoder(resp.Body).Decode(v)
	}
	return nil
}

// Do executes the request (for custom methods)
func (c *ServiceClient) Do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}
