package usecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/nibble-4bits/ondemand-go-bootcamp/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockPokemons = []entity.Pokemon{
	{ID: 1},
	{ID: 2},
	{ID: 3},
}

type mockPokemonRepo struct {
	mock.Mock
}

func (c mockPokemonRepo) GetByID(id int) (*entity.Pokemon, error) {
	args := c.Called(id)
	return args.Get(0).(*entity.Pokemon), args.Error(1)
}

func (c mockPokemonRepo) GetAll() ([]entity.Pokemon, error) {
	args := c.Called()
	return args.Get(0).([]entity.Pokemon), args.Error(1)
}

func (c mockPokemonRepo) GetByParity(parity string, itemCount int, quota int) ([]entity.Pokemon, error) {
	args := c.Called()
	return args.Get(0).([]entity.Pokemon), args.Error(1)
}

func TestPokemonService_GetByID(t *testing.T) {
	tests := []struct {
		id   int
		want *entity.Pokemon
		err  error
	}{
		{id: 1, want: &mockPokemons[0], err: nil},
		{id: 2, want: &mockPokemons[1], err: nil},
		{id: 3, want: &mockPokemons[2], err: nil},
		{id: 4, want: nil, err: errors.New("not found")},
	}

	for _, test := range tests {
		repo := mockPokemonRepo{}
		repo.On("GetByID", test.id).Return(test.want, test.err)
		service := NewPokemonService(repo)

		testname := fmt.Sprintf("Get comment by ID %v", test.id)
		t.Run(testname, func(t *testing.T) {
			pokemon, err := service.GetByID(test.id)

			if err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.Equal(t, pokemon, test.want)
			}
		})
	}
}

func TestPokemonService_GetAll(t *testing.T) {
	tests := []struct {
		name string
		want []entity.Pokemon
		err  error
	}{
		{name: "Get all successfully", want: mockPokemons, err: nil},
		{name: "Error on get all", want: nil, err: errors.New("not found")},
	}

	for _, test := range tests {
		repo := mockPokemonRepo{}
		repo.On("GetAll").Return(test.want, test.err)
		service := NewPokemonService(repo)

		t.Run(test.name, func(t *testing.T) {
			pokemons, err := service.GetAll()

			if err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.Equal(t, pokemons, test.want)
			}
		})
	}
}

func TestPokemonService_GetByParity(t *testing.T) {
	tests := []struct {
		name      string
		parity    string
		itemCount int
		quota     int
		want      []entity.Pokemon
		err       error
	}{
		{
			name:      "Get even pokemons successfully",
			parity:    "even",
			itemCount: 5,
			quota:     1,
			want:      []entity.Pokemon{mockPokemons[1]},
			err:       nil,
		},
		{
			name:      "Get odd pokemons successfully",
			parity:    "odd",
			itemCount: 5,
			quota:     1,
			want:      []entity.Pokemon{mockPokemons[0], mockPokemons[2]},
			err:       nil,
		},
		{
			name: "Error on get by parity",
			want: nil,
			err:  errors.New("an error has occurred"),
		},
	}

	for _, test := range tests {
		repo := mockPokemonRepo{}
		repo.On("GetByParity").Return(test.want, test.err)
		service := NewPokemonService(repo)

		t.Run(test.name, func(t *testing.T) {
			pokemons, err := service.GetByParity(test.parity, test.itemCount, test.quota)

			if err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.Equal(t, pokemons, test.want)
			}
		})
	}
}
