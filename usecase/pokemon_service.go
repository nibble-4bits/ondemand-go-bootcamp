package usecase

import (
	"fmt"

	"github.com/nibble-4bits/ondemand-go-bootcamp/entity"
)

type pokemonService struct {
	repo PokemonRepository
}

// PokemonService is an interface that represents basic CRUD operations
// that can be executed on a Pokemon entity
type PokemonService interface {
	GetByID(id int) (*entity.Pokemon, error)
	GetAll() ([]entity.Pokemon, error)
	GetByParity(parity string, itemCount int, quota int) ([]entity.Pokemon, error)
}

// NewPokemonService receives a PokemonRepository and returns an instance of pokemonService
func NewPokemonService(r PokemonRepository) pokemonService {
	return pokemonService{repo: r}
}

// GetByID returns a pokemon by ID from the underlying service repository.
func (s pokemonService) GetByID(id int) (*entity.Pokemon, error) {
	pokemonFound, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return pokemonFound, nil
}

// GetByID returns all pokemons from the underlying service repository.
func (s pokemonService) GetAll() ([]entity.Pokemon, error) {
	pokemons, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return pokemons, nil
}

// GetByParity returns an slice of pokemons filtered by "even" or "odd" parity.
func (s pokemonService) GetByParity(parity string, itemCount int, quota int) ([]entity.Pokemon, error) {
	pokemons, err := s.repo.GetByParity(parity, itemCount, quota)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return pokemons, nil
}
