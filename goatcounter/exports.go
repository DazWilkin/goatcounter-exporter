package goatcounter

import "net/http"

type ExportsInterface interface{}

type ExportsClient struct {
	client *http.Client
}

func NewExportsClient(client *http.Client) ExportsInterface {
	return &ExportsClient{
		client: client,
	}
}
