package usecase

import "github.com/nibble-4bits/ondemand-go-bootcamp/entity"

type PokemonRepository interface {
	GetByID(id int) (entity.Pokemon, error)
	GetAll() ([]entity.Pokemon, error)
}
