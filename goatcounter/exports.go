package goatcounter

import "fmt"

// ExportsInterface is an interface that defines methods for GoatCounter Exports endpoints
type ExportsInterface interface {
	URL(string) string
}

// ExportsClient is a type that implements methods for GoatCounter Exports endpoints
type ExportsClient struct {
	client *Client
	path   string
}

// Url is a method that returns the endpoint's method URL
func (c *ExportsClient) URL(path string) string {
	return fmt.Sprintf("%s/%s", c.path, path)
}
