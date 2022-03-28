package data

import (
	"github.com/nibble-4bits/ondemand-go-bootcamp/httpClient"
)

// httpDataSource represents a data source that fetches data
// from an HTTP endpoint.
type httpDataSource struct {
	// client is an instance of an HTTP client.
	client *httpClient.Client
}

// NewHTTPDataSource returns and instance of httpDataSource
func NewHTTPDataSource() httpDataSource {
	return httpDataSource{
		client: httpClient.New(),
	}
}

// ReadItem sends an HTTP GET request to the specified endpoint
// and returns an *httpClient.Response.
func (h httpDataSource) ReadItem(endpoint string) (*httpClient.Response, error) {
	resp, err := h.client.Get(endpoint)

	return resp, err
}
