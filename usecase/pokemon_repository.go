package usecase

import "github.com/nibble-4bits/ondemand-go-bootcamp/entity"

// PokemonRepository is an interface that represents basic CRUD operations
// that can be executed on a Pokemon entity
type PokemonRepository interface {
	GetByID(id int) (*entity.Pokemon, error)
	GetAll() ([]entity.Pokemon, error)
}
