package routableclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/DekaCube/go-routable/internal/types"
)

type SimpleRoutableClient struct {
	bearerToken string
	host        *url.URL
	httpClient  *http.Client
}

func NewSimpleRoutableClient(bearerToken string, host string, timeout time.Duration) (*SimpleRoutableClient, error) {
	parsedURL, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("failed to parse host: %w", err)
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return nil, fmt.Errorf("invalid host %q: scheme and host are required", host)
	}

	return &SimpleRoutableClient{
		bearerToken: bearerToken,
		host:        parsedURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

// ListCompanies retrieves a list of companies from Routable.
func (c *SimpleRoutableClient) ListCompanies(opts ...ListCompaniesOption) (*types.RoutableListCompaniesResponse, error) {
	u := c.host.ResolveReference(&url.URL{Path: "/v1/companies"})

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	if len(params) > 0 {
		u.RawQuery = params.Encode()
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.bearerToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	var result types.RoutableListCompaniesResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}
