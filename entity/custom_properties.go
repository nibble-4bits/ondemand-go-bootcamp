package entity

import (
	"log"
	"strconv"
)

// intProperty represents an integer property of an entity
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

// boolProperty represents a boolean property of an entity
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
