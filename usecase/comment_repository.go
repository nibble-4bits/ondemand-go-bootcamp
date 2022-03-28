package usecase

import "github.com/nibble-4bits/ondemand-go-bootcamp/entity"

// CommentRepository is an interface that represents basic CRUD operations
// that can be executed on a Comment entity
type CommentRepository interface {
	GetByID(id int) (*entity.Comment, error)
}
