package usecase

import (
	"fmt"

	"github.com/nibble-4bits/ondemand-go-bootcamp/entity"
)

type pokemonService struct {
	repo PokemonRepository
}

type PokemonService interface {
	GetByID(id int) (*entity.Pokemon, error)
	GetAll() ([]entity.Pokemon, error)
}

func NewPokemonService(r PokemonRepository) pokemonService {
	return pokemonService{repo: r}
}

func (s pokemonService) GetByID(id int) (*entity.Pokemon, error) {
	pokemonFound, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return pokemonFound, nil
}

func (s pokemonService) GetAll() ([]entity.Pokemon, error) {
	pokemons, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return pokemons, nil
}
