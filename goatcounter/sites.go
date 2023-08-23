package goatcounter

type SitesInterface interface{}

var _ SitesInterface = (*SitesClient)(nil)

type SitesClient struct {
	client *Client
}
