package usecase

import (
	"fmt"

	"github.com/nibble-4bits/ondemand-go-bootcamp/entity"
)

type commentService struct {
	repo CommentRepository
}

// CommentService is an interface that represents basic CRUD operations
// that can be executed on a Comment entity
type CommentService interface {
	GetByID(id int) (*entity.Comment, error)
}

// NewCommentService receives a CommentRepository and returns an instance of commentService
func NewCommentService(r CommentRepository) commentService {
	return commentService{repo: r}
}

// GetByID returns a comment by ID from the underlying service repository.
func (s commentService) GetByID(id int) (*entity.Comment, error) {
	commentFound, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return commentFound, nil
}
