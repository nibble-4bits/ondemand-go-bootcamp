package adapter

import (
	"github.com/nibble-4bits/ondemand-go-bootcamp/httpClient"

	"github.com/stretchr/testify/mock"
)

type mockCSVDataSource struct {
	mock.Mock
}

func (m mockCSVDataSource) ReadCollection() ([][]string, error) {
	args := m.Called()
	return args.Get(0).([][]string), args.Error(1)
}

type mockHTTPDataSource struct {
	mock.Mock
}

func (m mockHTTPDataSource) ReadItem(endpoint string) (*httpClient.Response, error) {
	// We don't care about the actual endpoint, since any HTTP API
	// should return similar results.
	args := m.Called()
	return args.Get(0).(*httpClient.Response), args.Error(1)
}

type mockDataStore struct {
	mock.Mock
}

func (m mockDataStore) SaveRecord(record []string) error {
	return nil
}
