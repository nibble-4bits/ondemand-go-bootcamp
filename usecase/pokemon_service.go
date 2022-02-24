package usecase

import (
	"errors"
	"fmt"

	"github.com/nibble-4bits/ondemand-go-bootcamp/entity"
)

var (
	ErrUnsupportedParityType = errors.New("unsupported parity type. Must be one of 'even' or 'odd'")
)

type pokemonService struct {
	repo PokemonRepository
}

// PokemonService is an interface that represents basic CRUD operations
// that can be executed on a Pokemon entity
type PokemonService interface {
	GetByID(id int) (*entity.Pokemon, error)
	GetAll() ([]entity.Pokemon, error)
	GetByParity(parity string, workers int, itemCount int, quota int) ([]entity.Pokemon, error)
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

func (s pokemonService) GetByParity(parity string, workers int, itemCount int, quota int) ([]entity.Pokemon, error) {
	if parity != "even" && parity != "odd" {
		return nil, ErrUnsupportedParityType
	}

	workerFunc := func(id int, parity string, quota int, jobs <-chan entity.Pokemon, results chan<- entity.Pokemon) {
		found := 0

		for p := range jobs {
			if (parity == "even" && p.ID%2 == 0) || (parity == "odd" && p.ID%2 == 1) {
				found++
				results <- p
			}

			if found == quota {
				return
			}
		}

	}

	pokemons, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	jobs := make(chan entity.Pokemon, len(pokemons))
	results := make(chan entity.Pokemon, (len(pokemons)/2)+1)
	remainingItems := itemCount
	for i := 0; i < workers; i++ {
		actualQuota := quota
		if remainingItems < quota {
			actualQuota = remainingItems
		}
		go workerFunc(i+1, parity, actualQuota, jobs, results)
		remainingItems -= quota
	}

	for _, pokemon := range pokemons {
		jobs <- pokemon
	}
	close(jobs)

	var filteredPokemons []entity.Pokemon
	var resultsEnd int
	if itemCount <= (len(pokemons)/2)+1 {
		resultsEnd = itemCount
	} else {
		resultsEnd = (len(pokemons) / 2) + 1
	}
	for i := 0; i < resultsEnd; i++ {
		filteredPokemons = append(filteredPokemons, <-results)
	}

	return filteredPokemons, nil
}
