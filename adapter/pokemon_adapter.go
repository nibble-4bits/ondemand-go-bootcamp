package adapter

import (
	"errors"
	"fmt"

	"github.com/nibble-4bits/ondemand-go-bootcamp/entity"
)

var (
	ErrPokemonsNotFound    = errors.New("no pokemons found")
	ErrPokemonNotFoundByID = errors.New("no pokemon found by ID")
)

type pokemonAdapter struct {
	dataSource DataSource
	pokemons   []entity.Pokemon
}

// NewPokemonAdapter receives a data source and will try to fetch the
// list of pokemons from a data source.
//
// If successful, an instance of *pokemonAdapter will be returned.
// Otherwise and error will be returned.
func NewPokemonAdapter(ds DataSource) (*pokemonAdapter, error) {
	adapter := &pokemonAdapter{dataSource: ds}

	if err := adapter.getPokemons(); err != nil {
		return nil, err
	}

	return adapter, nil
}

func (a *pokemonAdapter) getPokemons() error {
	csvRecords, err := a.dataSource.ReadCollection()
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

// GetAll returns an slice of all pokemons.
//
// In case no pokemons are found at all, an ErrPokemonsNotFound error is returned.
func (a *pokemonAdapter) GetAll() ([]entity.Pokemon, error) {
	if len(a.pokemons) == 0 {
		return nil, ErrPokemonsNotFound
	}

	return a.pokemons, nil
}
