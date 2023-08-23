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

type Client struct {
	client      *http.Client
	ratelimiter *rate.Limiter

	code  string
	token string
}

func NewClient(code, token string) *Client {
	return &Client{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		ratelimiter: rate.NewLimiter(rate.Every(time.Second), 4),

		code:  code,
		token: token,
	}
}

func (c *Client) Url(endpoint Endpoint) string {
	return fmt.Sprintf("https://%s.%s/%s/%s",
		c.code,
		domain,
		version,
		endpoint.String(),
	)
}
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

	// Always apply rate limiter to request
	if err := c.ratelimiter.Wait(ctx); err != nil {
		msg := "GoatCounter API request canceled or timed out"
		return []byte{}, fmt.Errorf(msg)
	}

	resp, err := c.client.Do(rqst)
	if err != nil {
		msg := "unable to send GoatCounter HTTP request"
		return []byte{}, fmt.Errorf(msg)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		msg := "unable to read response body"
		return []byte{}, fmt.Errorf(msg)
	}

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}
		return []byte{}, errorResp
	}

	return respBody, nil

}
func (c *Client) Count() {}
func (c *Client) Exports() ExportsInterface {
	return &ExportsClient{}
}
func (c *Client) Me() (*User, error) {
	ctx := context.Background()

	endpoint := Me
	url := c.Url(endpoint)
	method := http.MethodGet

	user := &User{}

	resp, err := c.Do(ctx, method, url, nil)
	if err != nil {
		msg := "unable to get information about the current user"
		return user, fmt.Errorf(msg)
	}

	if err := json.Unmarshal(resp, user); err != nil {
		msg := "unable to unmarshal response"
		return user, fmt.Errorf(msg)
	}

	return user, nil
}
func (c *Client) Paths() {
}
func (c *Client) Sites() SitesInterface {
	return &SitesClient{
		client: c,
	}
}
func (c *Client) Stats() StatsInterface {
	return &StatsClient{
		client: c,
	}
}
