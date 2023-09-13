package goatcounter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

const (
	domain  string = "goatcounter.com"
	version string = "api/v0"
)

// ClientInterface is an interface that defines methods for GoatCounter API client
type ClientInterface interface {
	Do(context.Context, string, string, io.Reader) ([]byte, error)

	// This method is not part of a resource type
	Count() (*Count, error)

	// These correspond to resource types defined on the API
	// /exports
	// /paths
	// /sites
	// /stats
	// /users
	Exports() ExportsInterface
	Paths() PathsInterface
	Sites() SitesInterface
	Stats() StatsInterface
	Users() UsersInterface
}

var _ ClientInterface = (*Client)(nil)

// Client is a type that implements methods for GoatCounter API client
type Client struct {
	Code string

	client      *http.Client
	ratelimiter *rate.Limiter

	path  string
	token string
}

// NewClient is a function that creates a new GoatCounter client
func NewClient(code, token string) *Client {
	host := fmt.Sprintf("%s.%s", code, domain)
	path := fmt.Sprintf("https://%s/%s", host, version)

	return &Client{
		Code: code,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		ratelimiter: rate.NewLimiter(rate.Every(time.Second), 4),

		path:  path,
		token: token,
	}
}

// Do is a method that makes a GoatCounter API request
func (c *Client) Do(ctx context.Context, method, url string, body io.Reader) ([]byte, error) {
	rqst, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		msg := "unable to create GoatCounter HTTP request"
		return []byte{}, fmt.Errorf(msg)
	}

	// Add Authorization header
	rqst.Header.Set(
		"Authorization",
		fmt.Sprintf("Bearer %s", c.token),
	)

	// Add Content-Type header
	rqst.Header.Set(
		"Content-Type",
		"application/json",
	)
	// Add Accept header
	rqst.Header.Set(
		"Accept",
		"application/json",
	)

	// Always apply rate limiter to request
	if err := c.ratelimiter.Wait(ctx); err != nil {
		msg := "GoatCounter API request canceled or timed out"
		return nil, fmt.Errorf(msg)
	}

	resp, err := c.client.Do(rqst)
	if err != nil {
		msg := "unable to send GoatCounter HTTP request"
		if errResponseor, ok := err.(*ErrorResponse); ok {
			return nil, fmt.Errorf("%s\n%+v", msg, errResponseor)
		}
		return nil, fmt.Errorf("%s\n%+v", msg, err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		msg := "unable to read response body"
		return nil, fmt.Errorf(msg)
	}

	if resp.StatusCode != http.StatusOK {
		errResponseor := ErrorResponse{}
		if err := json.Unmarshal(respBody, &errResponseor); err == nil {
			return nil, errResponseor
		}

		msg := "unable to unmarshal error message"
		return nil, fmt.Errorf(msg)
	}

	return respBody, nil
}

// Count is a method that implements GoatCounter /count endpoint
func (c *Client) Count() (*Count, error) {
	ctx := context.Background()
	method := http.MethodPost
	url := fmt.Sprintf("%s/count", c.path)

	count := &Count{}

	resp, err := c.Do(ctx, method, url, nil)
	if err != nil {
		msg := "unable to get pageview count"

		if errResponseor, ok := err.(*ErrorResponse); ok {
			return count, fmt.Errorf("%s\n%+v", msg, errResponseor)
		}

		return count, fmt.Errorf("%s\n%+v", msg, err)
	}

	if err := json.Unmarshal(resp, count); err != nil {
		msg := "unable to marshal response as count"
		return count, fmt.Errorf(msg)
	}

	return count, nil
}

// Exports is a method that implements GoatCounter Exports endpoints
// Exports returns an implementation of ExportsInterface
func (c *Client) Exports() ExportsInterface {
	return &ExportsClient{
		client: c,
		path:   fmt.Sprintf("%s/export", c.path),
	}
}

// Paths is a method that implements GoatCounter Paths endpoints
// Paths returns an implementation of PathsInterface
func (c *Client) Paths() PathsInterface {
	return &PathsClient{
		client: c,
		path:   fmt.Sprintf("%s/paths", c.path),
	}
}

// Sites is a method that implements GoatCounter Sites endpoints
// Sites returns an implementation of SitesInterface
func (c *Client) Sites() SitesInterface {
	return &SitesClient{
		client: c,
		path:   fmt.Sprintf("%s/sites", c.path),
	}
}

// Stats is a method that implements GoatCounter Stats endpoints
// Stats returns an implementation of StatsInterface
func (c *Client) Stats() StatsInterface {
	return &StatsClient{
		client: c,
		path:   fmt.Sprintf("%s/stats", c.path),
	}
}

// Users is a method that implements GoatCounter Users endpoints
// Users returns an implementation of UsersInterface
func (c *Client) Users() UsersInterface {
	return &UsersClient{
		client: c,
		path:   c.path,
	}
}
