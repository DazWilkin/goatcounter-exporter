package goatcounter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// SitesInterface is an interface that defines methods for GoatCounter Sites endpoints
type SitesInterface interface {
	Get(string) (*Site, error)
	List() (*SitesResponse, error)
	URL(string) string
}

// SitesInterface is an interface that implements methods for GoatCounter Sites endpoints
var _ SitesInterface = (*SitesClient)(nil)

type SitesClient struct {
	client *Client
	path   string
}

func (c *SitesClient) URL(path string) string {
	return fmt.Sprintf("%s/%s", c.path, path)
}
func (c *SitesClient) Get(ID string) (*Site, error) {
	ctx := context.Background()
	method := http.MethodGet
	url := c.URL(ID)

	site := &Site{}
	resp, err := c.client.Do(ctx, method, url, nil)
	if err != nil {
		msg := "unable to list sites"

		if errResponseor, ok := err.(*ErrorResponse); ok {
			return site, fmt.Errorf("%s\n%+v", msg, errResponseor)
		}

		return site, fmt.Errorf(msg)
	}

	if err := json.Unmarshal(resp, site); err != nil {
		msg := "unable to unmarshal response"
		return site, fmt.Errorf(msg)
	}

	return site, nil
}
func (c *SitesClient) List() (*SitesResponse, error) {
	ctx := context.Background()
	method := http.MethodGet
	url := c.URL("")

	sitesResponse := &SitesResponse{}
	resp, err := c.client.Do(ctx, method, url, nil)
	if err != nil {
		msg := "unable to list sites"

		if errResponseor, ok := err.(*ErrorResponse); ok {
			return sitesResponse, fmt.Errorf("%s\n%+v", msg, errResponseor)
		}

		return sitesResponse, fmt.Errorf(msg)
	}

	if err := json.Unmarshal(resp, sitesResponse); err != nil {
		msg := "unable to unmarshal response"
		return sitesResponse, fmt.Errorf(msg)
	}

	return sitesResponse, nil
}
