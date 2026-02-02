package httpclient

import (
	"net/http"
	"time"
)

// RetryRoundTripper wraps http.RoundTripper with retry logic
type RetryRoundTripper struct {
	Next          http.RoundTripper
	MaxRetries    int
	RetryDelay    time.Duration
	RetryableCode func(statusCode int) bool
}

// RoundTrip implements http.RoundTripper
func (r *RetryRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error
	for attempt := 0; attempt <= r.MaxRetries; attempt++ {
		if r.Next == nil {
			r.Next = http.DefaultTransport
		}
		resp, err = r.Next.RoundTrip(req)
		if err != nil {
			if attempt < r.MaxRetries {
				time.Sleep(r.retryDelay(attempt))
				continue
			}
			return nil, err
		}
		if r.RetryableCode != nil && r.RetryableCode(resp.StatusCode) && attempt < r.MaxRetries {
			_ = resp.Body.Close()
			time.Sleep(r.retryDelay(attempt))
			continue
		}
		return resp, nil
	}
	return resp, err
}

func (r *RetryRoundTripper) retryDelay(attempt int) time.Duration {
	if r.RetryDelay > 0 {
		return r.RetryDelay
	}
	return time.Duration(attempt+1) * 100 * time.Millisecond
}
