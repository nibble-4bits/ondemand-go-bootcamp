package entity

import (
	"log"
	"strconv"
)

// intProperty represents an integer property of a Pokemon
type intProperty int

// ParseInt tries to parse the string parameter to an int.
// If the parsing succeeds, the result is assigned to the receiver. Otherwise, defaultVal gets assigned.
func (i *intProperty) ParseInt(s string, defaultVal int) {
	parsed, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Error parsing string \"%v\" to int, using default value %v\n", s, defaultVal)
		parsed = defaultVal
	}

	*i = intProperty(parsed)
}

// boolProperty represents a boolean property of a Pokemon
type boolProperty bool

// ParseInt tries to parse the string parameter to a bool.
// If the parsing succeeds, the result is assigned to the receiver. Otherwise, defaultVal gets assigned.
func (b *boolProperty) ParseBool(s string, defaultVal bool) {
	parsed, err := strconv.ParseBool(s)
	if err != nil {
		log.Printf("Error parsing string \"%v\" to bool, using default value %v\n", s, defaultVal)
		parsed = defaultVal
	}

	*b = boolProperty(parsed)
}

// Pokemon is a struct that represents the features of a pokemon
type Pokemon struct {
	ID         intProperty  `json:"id"`
	Name       string       `json:"name"`
	Type1      string       `json:"type1"`
	Type2      string       `json:"type2"`
	Total      intProperty  `json:"total"`
	HP         intProperty  `json:"hp"`
	Attack     intProperty  `json:"attack"`
	Defense    intProperty  `json:"defense"`
	SpAtk      intProperty  `json:"sp_attack"`
	SpDef      intProperty  `json:"sp_defense"`
	Speed      intProperty  `json:"speed"`
	Generation intProperty  `json:"generation"`
	Legendary  boolProperty `json:"legendary"`
}
