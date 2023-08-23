package goatcounter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type StatsInterface interface {
	Total() (*StatsTotal, error)
	Hits() (*StatsHits, error)
}

var _ StatsInterface = (*StatsClient)(nil)

type StatsClient struct {
	client *Client
}

func (c *StatsClient) Hits() (*StatsHits, error) {
	ctx := context.Background()

	endpoint := Stats
	url := fmt.Sprintf("%s/hits", c.client.Url(endpoint))
	method := http.MethodGet

	hits := &StatsHits{}
	resp, err := c.client.Do(ctx, method, url, nil)
	if err != nil {
		msg := "unable to get total pageview and visitor counts"
		return hits, fmt.Errorf(msg)
	}

	if err := json.Unmarshal(resp, hits); err != nil {
		msg := "unable to unmarshal response"
		return hits, fmt.Errorf(msg)
	}

	return hits, nil
}
func (c *StatsClient) Total() (*StatsTotal, error) {
	ctx := context.Background()

	endpoint := Stats
	url := fmt.Sprintf("%s/total", c.client.Url(endpoint))
	method := http.MethodGet

	total := &StatsTotal{}
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
