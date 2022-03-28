package adapter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nibble-4bits/ondemand-go-bootcamp/entity"
	"github.com/nibble-4bits/ondemand-go-bootcamp/httpClient"

	"github.com/stretchr/testify/assert"
)

var mockCommentCSVData = [][]string{
	{"1", "1", "Eliseo", "eliseo@gardner.biz", "lorem"},
	{"2", "1", "John", "john@gardner.biz", "ipsum"},
	{"3", "2", "Sophia", "sophia@gardner.biz", "dolor"},
}

var mockCommentHTTPData = []*httpClient.Response{
	{
		Body: []byte(`{
		"id":     4,
		"postId": 3,
		"name":   "Mike",
		"email":  "mike@alysha.tv",
		"body":   "sit"
	}`),
		StatusCode: http.StatusOK,
	},
	{
		Body:       []byte("{}"),
		StatusCode: http.StatusNotFound,
	},
}

var mockComments = []entity.Comment{
	{
		ID:     1,
		PostID: 1,
		Name:   "Eliseo",
		Email:  "eliseo@gardner.biz",
		Body:   "lorem",
	},
	{
		ID:     2,
		PostID: 1,
		Name:   "John",
		Email:  "john@gardner.biz",
		Body:   "ipsum",
	},
	{
		ID:     3,
		PostID: 2,
		Name:   "Sophia",
		Email:  "sophia@gardner.biz",
		Body:   "dolor",
	},
	{
		ID:     4,
		PostID: 3,
		Name:   "Mike",
		Email:  "mike@alysha.tv",
		Body:   "sit",
	},
}

func TestCommentAdapter_GetByID(t *testing.T) {
	tests := []struct {
		id       int
		httpData *httpClient.Response
		want     *entity.Comment
		err      error
	}{
		{id: 1, httpData: nil, want: &mockComments[0], err: nil},
		{id: 2, httpData: nil, want: &mockComments[1], err: nil},
		{id: 3, httpData: nil, want: &mockComments[2], err: nil},
		{id: 4, httpData: mockCommentHTTPData[0], want: &mockComments[3], err: nil},
		{id: 5, httpData: mockCommentHTTPData[1], want: nil, err: ErrCommentNotFoundByID},
	}

	for _, test := range tests {
		csvDataSource := mockCSVDataSource{}
		csvDataSource.On("ReadCollection").Return(mockCommentCSVData, nil)
		httpDataSource := mockHTTPDataSource{}
		httpDataSource.On("ReadItem").Return(test.httpData, nil)
		dataStore := mockDataStore{}
		adapter, err := NewCommentAdapter(csvDataSource, httpDataSource, dataStore)

		assert.Nil(t, err)

		testname := fmt.Sprintf("Get comment by ID %v", test.id)
		t.Run(testname, func(t *testing.T) {
			comment, err := adapter.GetByID(test.id)

			if err != nil {
				assert.ErrorIs(t, err, test.err)
			} else {
				assert.Equal(t, comment, test.want)
			}
		})
	}
}
