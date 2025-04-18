package goatcounter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// UsersInterface is an interface that defines GoatCounter Users endpoints
type UsersInterface interface {
	Me() (*User, error)
	URL(string) string
}

var _ UsersInterface = (*UsersClient)(nil)

// UsersClient is a type that implements GoatCounter Users endpoints
type UsersClient struct {
	client *Client
	path   string
}

// Me is a method that implements GoatCounter /me
func (c *UsersClient) Me() (*User, error) {
	ctx := context.Background()
	method := http.MethodGet
	url := fmt.Sprintf("%s/me", c.path)

	user := &User{}

	resp, err := c.client.Do(ctx, method, url, nil)
	if err != nil {
		msg := "unable to get information about the current user"
		return user, errors.New(msg)
	}

	if err := json.Unmarshal(resp, user); err != nil {
		msg := "unable to unmarshal response as user"
		return user, errors.New(msg)
	}

	return user, nil
}

// Url is a method that returns the endpoint's method URL
func (c *UsersClient) URL(path string) string {
	return fmt.Sprintf("%s/%s", c.path, path)
}
