package config

import (
	"log"

	"github.com/spf13/viper"
)

// configuration is a set of properties that are loaded when the program runs
type configuration struct {
	PokemonCSVPath string `mapstructure:"pokemon_csv_path"`
}

// Config is a variable that holds the program configuration
var Config configuration

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalln(err)
	}
}
