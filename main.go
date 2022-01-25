package main

import (
	"log"

	"github.com/nibble-4bits/ondemand-go-bootcamp/adapter"
	"github.com/nibble-4bits/ondemand-go-bootcamp/data"
	"github.com/nibble-4bits/ondemand-go-bootcamp/httpAPI/router"
	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"

	"github.com/spf13/viper"
)

// configuration is a set of properties that are loaded when the program runs
type configuration struct {
	// CSVDataPath is a path to the CSV file that contains the data that will be served by the HTTP API
	CSVDataPath string
}

// config is a variable that holds the program configuration
var config configuration

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	csvDataSource := data.NewCSVDataSource(config.CSVDataPath)

	pokemonAdapter, err := adapter.NewPokemonAdapter(csvDataSource)
	if err != nil {
		// Could handle the error here more gracefully, for example
		// we could try and fetch the pokemons from another data source.
		// For now, just exit fatally and log the error.
		log.Fatalln(err)
	}

	pokemonService := usecase.NewPokemonService(pokemonAdapter)

	router.StartServer(pokemonService)
}
