package usecase

import (
	"errors"
	"fmt"
	"sync"

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

// GetByParity returns an slice of pokemons filtered by "even" or "odd" parity,
// using the worker pool pattern.
//
// The following arguments have to be passed:
//
// - parity: Must be "even" or "odd".
//
// - workers: The number of goroutines to spawn.
//
// - itemCount: The number of pokemons that must appear in the resulting slice.
//
// - quota: The number of pokemons each goroutine will process at most to verify if they match
// the corresponding parity.
func (s pokemonService) GetByParity(parity string, workers int, itemCount int, quota int) ([]entity.Pokemon, error) {
	if parity != "even" && parity != "odd" {
		return nil, ErrUnsupportedParityType
	}

	workerFunc := func(parity string, quota int, jobs <-chan entity.Pokemon, results chan<- entity.Pokemon) {
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
	wg := sync.WaitGroup{}
	remainingItems := itemCount
	for i := 0; i < workers; i++ {
		actualQuota := quota
		if remainingItems < quota {
			actualQuota = remainingItems
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			workerFunc(parity, actualQuota, jobs, results)
		}()
		remainingItems -= quota
	}

	for _, pokemon := range pokemons {
		jobs <- pokemon
	}
	close(jobs)

	wg.Wait()
	close(results)

	var filteredPokemons []entity.Pokemon
	for result := range results {
		filteredPokemons = append(filteredPokemons, result)
	}

	return filteredPokemons, nil
}
