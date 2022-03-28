package adapter

import (
	"errors"
	"fmt"
	"math"
	"sync"

	"github.com/nibble-4bits/ondemand-go-bootcamp/entity"
)

const (
	// Max number of goroutines that can be spawned by a worker pool
	MAX_WORKERS = 20
)

var (
	ErrPokemonsNotFound      = errors.New("no pokemons found")
	ErrPokemonNotFoundByID   = errors.New("no pokemon found by ID")
	ErrUnsupportedParityType = errors.New("unsupported parity type. Must be one of 'even' or 'odd'")
	ErrMaxNumberOfWorkers    = errors.New("max number of workers excedeed")
)

type pokemonAdapter struct {
	csvDataSource CSVDataSource
	pokemons      []entity.Pokemon
}

// NewPokemonAdapter receives a data source and will try to fetch the
// list of pokemons from a data source.
//
// If successful, an instance of *pokemonAdapter will be returned.
// Otherwise and error will be returned.
func NewPokemonAdapter(ds CSVDataSource) (*pokemonAdapter, error) {
	adapter := &pokemonAdapter{csvDataSource: ds}

	if err := adapter.loadPokemons(); err != nil {
		return nil, err
	}

	return adapter, nil
}

func (a *pokemonAdapter) loadPokemons() error {
	csvRecords, err := a.csvDataSource.ReadCollection()
	if err != nil {
		return err
	}

	for _, v := range csvRecords {
		p := entity.Pokemon{}

		p.ID.ParseInt(v[0], -1)
		p.Name = v[1]
		p.Type1 = v[2]
		p.Type2 = v[3]
		p.Total.ParseInt(v[4], -1)
		p.HP.ParseInt(v[5], -1)
		p.Attack.ParseInt(v[6], -1)
		p.Defense.ParseInt(v[7], -1)
		p.SpAtk.ParseInt(v[8], -1)
		p.SpDef.ParseInt(v[9], -1)
		p.Speed.ParseInt(v[10], -1)
		p.Generation.ParseInt(v[11], -1)
		p.Legendary.ParseBool(v[12], false)

		a.pokemons = append(a.pokemons, p)
	}

	return nil
}

// GetByID searches for a pokemon with the given id parameter.
//
// If the search is successful, a pointer to the found Pokemon is returned.
// Otherwise and ErrPokemonNotFoundByID error is returned.
func (a *pokemonAdapter) GetByID(id int) (*entity.Pokemon, error) {
	for _, pokemon := range a.pokemons {
		if id == int(pokemon.ID) {
			return &pokemon, nil
		}
	}

	return nil, fmt.Errorf("%w %v", ErrPokemonNotFoundByID, id)
}

// GetAll returns a slice of all pokemons.
//
// In case no pokemons are found at all, an ErrPokemonsNotFound error is returned.
func (a *pokemonAdapter) GetAll() ([]entity.Pokemon, error) {
	if len(a.pokemons) == 0 {
		return nil, ErrPokemonsNotFound
	}

	return a.pokemons, nil
}

// GetByParity returns an slice of pokemons filtered by "even" or "odd" parity,
// using the worker pool pattern.
//
// The following arguments have to be passed:
//
// - parity: Must be "even" or "odd".
//
// - itemCount: The number of pokemons that must appear in the resulting slice.
//
// - quota: The number of pokemons each goroutine will process at most to verify if they match
// the corresponding parity.
func (a *pokemonAdapter) GetByParity(parity string, itemCount int, quota int) ([]entity.Pokemon, error) {
	if parity != "even" && parity != "odd" {
		return nil, ErrUnsupportedParityType
	}

	// The number of workers is calculated as the itemCount divided by the quota (items per worker)
	// We use the math.Ceil function to ensure the number of workers is at least 1
	// in case itemCount < quota
	workers := int(math.Ceil(float64(itemCount) / float64(quota)))
	if workers > MAX_WORKERS {
		return nil, fmt.Errorf("%w. Consider incrementing the items per worker or decreasing the number of items to filter", ErrMaxNumberOfWorkers)
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

	pokemons, err := a.GetAll()
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
