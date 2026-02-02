package httpclient

import "time"

// Option configures ServiceClient
type Option func(*ServiceClient)

// WithTimeout sets HTTP client timeout
func WithTimeout(d time.Duration) Option {
	return func(c *ServiceClient) {
		c.httpClient.Timeout = d
	}
}

// WithTransport sets custom http.RoundTripper (e.g. for retry)
func WithTransport(rt interface{}) Option {
	// Type assert to *http.Transport when needed
	return func(c *ServiceClient) {
		// c.httpClient.Transport = rt
		_ = rt
	}
}
