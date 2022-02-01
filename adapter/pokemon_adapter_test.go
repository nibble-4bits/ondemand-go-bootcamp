package adapter

import (
	"fmt"
	"testing"

	"github.com/nibble-4bits/ondemand-go-bootcamp/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockPokemonCSVData = [][]string{
	{"1", "Bulbasaur", "Grass", "Poison", "318", "45", "49", "49", "65", "65", "45", "1", "False"},
	{"2", "Ivysaur", "Grass", "Poison", "405", "60", "62", "63", "80", "80", "60", "1", "False"},
	{"3", "Venusaur", "Grass", "Poison", "525", "80", "82", "83", "100", "100", "80", "1", "False"},
}

var mockPokemons = []entity.Pokemon{
	{
		ID:         1,
		Name:       "Bulbasaur",
		Type1:      "Grass",
		Type2:      "Poison",
		Total:      318,
		HP:         45,
		Attack:     49,
		Defense:    49,
		SpAtk:      65,
		SpDef:      65,
		Speed:      45,
		Generation: 1,
		Legendary:  false,
	},
	{
		ID:         2,
		Name:       "Ivysaur",
		Type1:      "Grass",
		Type2:      "Poison",
		Total:      405,
		HP:         60,
		Attack:     62,
		Defense:    63,
		SpAtk:      80,
		SpDef:      80,
		Speed:      60,
		Generation: 1,
		Legendary:  false,
	},
	{
		ID:         3,
		Name:       "Venusaur",
		Type1:      "Grass",
		Type2:      "Poison",
		Total:      525,
		HP:         80,
		Attack:     82,
		Defense:    83,
		SpAtk:      100,
		SpDef:      100,
		Speed:      80,
		Generation: 1,
		Legendary:  false,
	},
}

type mockCSVDataSource struct {
	mock.Mock
}

func (m mockCSVDataSource) ReadCollection() ([][]string, error) {
	args := m.Called()
	return args.Get(0).([][]string), args.Error(1)
}

func TestPokemonAdapter_GetByID(t *testing.T) {
	tests := []struct {
		id   int
		want *entity.Pokemon
		err  error
	}{
		{id: 1, want: &mockPokemons[0], err: nil},
		{id: 2, want: &mockPokemons[1], err: nil},
		{id: 3, want: &mockPokemons[2], err: nil},
		{id: 4, want: nil, err: ErrPokemonNotFoundByID},
	}

	for _, test := range tests {
		csvDataSource := mockCSVDataSource{}
		csvDataSource.On("ReadCollection").Return(mockPokemonCSVData, nil)
		adapter, err := NewPokemonAdapter(csvDataSource)

		assert.Nil(t, err)

		testname := fmt.Sprintf("Get comment by ID %v", test.id)
		t.Run(testname, func(t *testing.T) {
			pokemon, err := adapter.GetByID(test.id)

			if err != nil {
				assert.ErrorIs(t, err, test.err)
			} else {
				assert.Equal(t, pokemon, test.want)
			}
		})
	}
}

func TestPokemonService_GetAll(t *testing.T) {
	tests := []struct {
		name    string
		csvData [][]string
		want    []entity.Pokemon
		err     error
	}{
		{name: "Get all successfully", csvData: mockPokemonCSVData, want: mockPokemons, err: nil},
		{name: "Error on empty data", csvData: nil, want: nil, err: ErrPokemonsNotFound},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			csvDataSource := mockCSVDataSource{}
			csvDataSource.On("ReadCollection").Return(test.csvData, nil)
			adapter, err := NewPokemonAdapter(csvDataSource)

			assert.Nil(t, err)

			pokemons, err := adapter.GetAll()

			if err != nil {
				assert.ErrorIs(t, err, test.err)
			} else {
				assert.Equal(t, pokemons, test.want)
			}
		})
	}
}
