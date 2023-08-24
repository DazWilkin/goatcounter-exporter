package goatcounter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// StatsInterface is an interface that defines methods for GoatCounter Stats endpoints
type StatsInterface interface {
	Total() (*Total, error)
	Hits() (*HitsResponse, error)
	URL(string) string
}

var _ StatsInterface = (*StatsClient)(nil)

// StatsClient is a type that implements methods for GoatCounter Stats endpoints
type StatsClient struct {
	client *Client
	path   string
}

// Hits is a method that implements GoatCounter /stats/hits
func (c *StatsClient) Hits() (*HitsResponse, error) {
	ctx := context.Background()
	method := http.MethodGet
	url := fmt.Sprintf("%s/hits", c.path)

	hitsResponse := &HitsResponse{}
	resp, err := c.client.Do(ctx, method, url, nil)
	if err != nil {
		msg := "unable to get total pageview and visitor counts"
		return hitsResponse, fmt.Errorf(msg)
	}

	if err := json.Unmarshal(resp, hitsResponse); err != nil {
		msg := "unable to unmarshal response"
		return hitsResponse, fmt.Errorf(msg)
	}

	return hitsResponse, nil
}

// Total is a method that implements GoatCounter /stats/total
func (c *StatsClient) Total() (*Total, error) {
	ctx := context.Background()
	method := http.MethodGet
	url := fmt.Sprintf("%s/total", c.path)

	total := &Total{}
	resp, err := c.client.Do(ctx, method, url, nil)
	if err != nil {
		msg := "unable to get total pageview counts"
		return total, fmt.Errorf(msg)
	}

	if err := json.Unmarshal(resp, total); err != nil {
		msg := "unable to unmarshal response"
		return total, fmt.Errorf(msg)
	}

	return total, nil
}

// Url is a method that returns the endpoint's method URL
func (c *StatsClient) URL(path string) string {
	return fmt.Sprintf("%s/%s", c.path, path)
}
