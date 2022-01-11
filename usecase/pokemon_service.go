package usecase

import (
	"log"

	"github.com/nibble-4bits/ondemand-go-bootcamp/entity"
)

type pokemonService struct {
	repo PokemonRepository
}

type PokemonService interface {
	GetByID(id int) entity.Pokemon
}

func NewPokemonService(r PokemonRepository) pokemonService {
	return pokemonService{repo: r}
}

func (s pokemonService) GetByID(id int) entity.Pokemon {
	pokemonFound, err := s.repo.GetByID(id)
	if err != nil {
		log.Fatal(err)
	}

	return pokemonFound
}
