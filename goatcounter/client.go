package goatcounter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

const (
	Version string = "api/v0"
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
	Code     string
	Instance string

	client      *http.Client
	ratelimiter *rate.Limiter
	logger      *slog.Logger

	path  string
	token string
}

// NewClient is a function that creates a new GoatCounter client
func NewClient(code, instance, token string, logger *slog.Logger) *Client {
	logger = logger.With("goatcounter", "client")
	path := fmt.Sprintf("https://%s.%s/%s", code, instance, Version)

	return &Client{
		Code:     code,
		Instance: instance,

		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		ratelimiter: rate.NewLimiter(rate.Every(time.Second), 4),
		logger:      logger,

		path:  path,
		token: token,
	}
}

// NewTestClient is a function that create a new GoatCounter test client
// The test client is used by the tests that use prometheus/testutil
// These require a local http endpoint
// The code parameter is not used to form the API endpoint but it is used by metrics
func NewTestClient(code, endpoint string, logger *slog.Logger) *Client {
	logger = logger.With("goatcounter", "testclient")
	path := fmt.Sprintf("%s/%s", endpoint, Version)

	return &Client{
		Code:     code,
		Instance: endpoint,

		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		ratelimiter: rate.NewLimiter(rate.Every(time.Second), 4),
		logger:      logger,

		path: path,
	}
}

// Do is a method that makes a GoatCounter API request
func (c *Client) Do(ctx context.Context, method, url string, body io.Reader) ([]byte, error) {
	logger := c.logger.With("method", "do")
	rqst, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		msg := "unable to create GoatCounter HTTP request"
		logger.Error(msg, "err", err)
		return []byte{}, errors.New(msg)
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
		logger.Error(msg, "err", err)
		return nil, errors.New(msg)
	}

	resp, err := c.client.Do(rqst)
	if err != nil {
		msg := "unable to send GoatCounter HTTP request"
		logger.Error(msg, "err", err)

		if errResponseor, ok := err.(*ErrorResponse); ok {
			return nil, fmt.Errorf("%s\n%+v", msg, errResponseor)
		}

		return nil, fmt.Errorf("%s\n%+v", msg, err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			logger.Error("error closing response body", "err", err)
		}
	}()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		msg := "unable to read response body"
		logger.Error(msg, "err", err)
		return nil, errors.New(msg)
	}

	if resp.StatusCode != http.StatusOK {
		errResponseor := ErrorResponse{}
		if err := json.Unmarshal(respBody, &errResponseor); err == nil {
			msg := "unable to unmarshal error message"
			logger.Error(msg, "err", err)
			return nil, errResponseor
		}

		msg := "unexpected response status code"
		logger.Error(msg, "code", resp.StatusCode)
		return nil, errors.New(msg)
	}

	return respBody, nil
}

// Count is a method that implements GoatCounter /count endpoint
func (c *Client) Count() (*Count, error) {
	logger := c.logger.With("method", "count")

	ctx := context.Background()
	method := http.MethodPost
	url := fmt.Sprintf("%s/count", c.path)

	count := &Count{}

	resp, err := c.Do(ctx, method, url, nil)
	if err != nil {
		msg := "unable to get pageview count"
		logger.Error(msg, "err", err)

		if errResponseor, ok := err.(*ErrorResponse); ok {
			return count, fmt.Errorf("%s\n%+v", msg, errResponseor)
		}

		return count, fmt.Errorf("%s\n%+v", msg, err)
	}

	if err := json.Unmarshal(resp, count); err != nil {
		msg := "unable to marshal response as count"
		logger.Error(msg, "err", err)
		return count, errors.New(msg)
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
