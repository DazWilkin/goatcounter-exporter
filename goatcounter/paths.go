package goatcounter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// PathsInterface is an interface that defines methods on GoatCounter Paths endpoints
type PathsInterface interface {
	List() (*PathsResponse, error)
	URL(string) string
}

var _ PathsInterface = (*PathsClient)(nil)

// PathsClient is a type that implements methods for GoatCounter Path endpoints
type PathsClient struct {
	client *Client
	path   string
}

// Url is a method that returns the endpoint's method URL
func (c *PathsClient) URL(path string) string {
	return fmt.Sprintf("%s/%s", c.path, path)
}

// List is a method that implements GoatCounter's /paths method
func (c *PathsClient) List() (*PathsResponse, error) {
	ctx := context.Background()
	method := http.MethodGet

	// Retain the base URL without the query string
	// The query string will differ for each pagination
	// But will be based upon the base URL
	// /paths?after={after}
	baseUrl := c.URL("")
	// Initially the URL to use will be the base URL
	url := baseUrl

	pathsResponses := &PathsResponse{}
	for {
		pathsResponse := &PathsResponse{}
		resp, err := c.client.Do(ctx, method, url, nil)
		if err != nil {
			msg := "unable to list paths"

			if errResponseor, ok := err.(*ErrorResponse); ok {
				return pathsResponse, fmt.Errorf("%s\n%+v", msg, errResponseor)
			}

			return pathsResponse, fmt.Errorf(msg)
		}

		if err := json.Unmarshal(resp, pathsResponse); err != nil {
			msg := "unable to unmarshal response"
			return pathsResponse, fmt.Errorf(msg)
		}

		// Append this page's Paths to the overall tally
		pathsResponses.Paths = append(
			pathsResponses.Paths,
			pathsResponse.Paths...,
		)

		// If there are no further pages, we're done
		if !pathsResponse.More {
			break
		}

		// Retrieve the next page
		// Ensure there are Paths
		if len(pathsResponse.Paths) == 0 {
			break
		}

		// Grab last ID
		after := pathsResponse.Paths[len(pathsResponse.Paths)-1].ID
		// Append it to the URL
		url = fmt.Sprintf("%s?after=%d", baseUrl, after)
	}
	return pathsResponses, nil
}
