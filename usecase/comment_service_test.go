package usecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/nibble-4bits/ondemand-go-bootcamp/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockComments = []entity.Comment{
	{ID: 1, PostID: 1},
	{ID: 2, PostID: 1},
	{ID: 3, PostID: 3},
}

type mockCommentRepo struct {
	mock.Mock
}

func (c mockCommentRepo) GetByID(id int) (*entity.Comment, error) {
	args := c.Called(id)
	return args.Get(0).(*entity.Comment), args.Error(1)
}

func TestCommentService_GetByID(t *testing.T) {
	tests := []struct {
		id   int
		want *entity.Comment
		err  error
	}{
		{id: 1, want: &mockComments[0], err: nil},
		{id: 2, want: &mockComments[1], err: nil},
		{id: 3, want: &mockComments[2], err: nil},
		{id: 4, want: nil, err: errors.New("not found")},
	}

	for _, test := range tests {
		repo := mockCommentRepo{}
		repo.On("GetByID", test.id).Return(test.want, test.err)
		service := NewCommentService(repo)

		testname := fmt.Sprintf("Get comment by ID %v", test.id)
		t.Run(testname, func(t *testing.T) {
			comment, err := service.GetByID(test.id)

			if err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.Equal(t, comment, test.want)
			}
		})
	}
}
