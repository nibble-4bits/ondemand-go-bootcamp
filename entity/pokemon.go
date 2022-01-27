package entity

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
