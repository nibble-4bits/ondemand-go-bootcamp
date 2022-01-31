package adapter

import "github.com/nibble-4bits/ondemand-go-bootcamp/httpClient"

// CSVDataSource represents a data source that reads from a CSV file.
type CSVDataSource interface {
	// ReadCollection reads the collection that belongs to the
	// current CSVDataSource and returns the items as a
	// slice of slices of strings.
	ReadCollection() ([][]string, error)
}

// HTTPDataSource represents a data source that fetches data
// from an HTTP endpoint.
type HTTPDataSource interface {
	// ReadItem sends a request to the specified endpoint
	// and returns an *httpClient.Response.
	ReadItem(endpoint string) (*httpClient.Response, error)
}
