package usecase

import (
	"github.com/nibble-4bits/ondemand-go-bootcamp/entity"
)

type pokemonService struct {
	repo PokemonRepository
}

type PokemonService interface {
	GetByID(id int) (entity.Pokemon, error)
}

func NewPokemonService(r PokemonRepository) pokemonService {
	return pokemonService{repo: r}
}

func (s pokemonService) GetByID(id int) (entity.Pokemon, error) {
	pokemonFound, err := s.repo.GetByID(id)
	if err != nil {
		return entity.Pokemon{}, err
	}

	return pokemonFound, nil
}
