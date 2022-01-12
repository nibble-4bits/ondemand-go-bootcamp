package adapter

import (
	"fmt"
	"strconv"

	"github.com/nibble-4bits/ondemand-go-bootcamp/entity"
)

type pokemonAdapter struct {
	dataSource DataSource
	pokemons   []entity.Pokemon
}

func NewPokemonAdapter(ds DataSource) *pokemonAdapter {
	adapter := &pokemonAdapter{dataSource: ds}

	adapter.getPokemons()

	return adapter
}

func (a *pokemonAdapter) getPokemons() {
	csvRecords := a.dataSource.ReadCollection()

	// Remove header from slice of records
	csvRecords = csvRecords[1:]

	for _, v := range csvRecords {
		p := entity.Pokemon{}

		p.Id, _ = strconv.Atoi(v[0])
		p.Name = v[1]
		p.Type1 = v[2]
		p.Type2 = v[3]
		p.Total, _ = strconv.Atoi(v[4])
		p.HP, _ = strconv.Atoi(v[5])
		p.Attack, _ = strconv.Atoi(v[6])
		p.Defense, _ = strconv.Atoi(v[7])
		p.SpAtk, _ = strconv.Atoi(v[8])
		p.SpDef, _ = strconv.Atoi(v[9])
		p.Speed, _ = strconv.Atoi(v[10])
		p.Generation, _ = strconv.Atoi(v[11])
		p.Legendary, _ = strconv.ParseBool(v[12])

		a.pokemons = append(a.pokemons, p)
	}
}

func (a *pokemonAdapter) GetByID(id int) (entity.Pokemon, error) {
	for _, pokemon := range a.pokemons {
		if id == pokemon.Id {
			return pokemon, nil
		}
	}

	return entity.Pokemon{}, fmt.Errorf("pokemon with ID %v not found", id)
}
