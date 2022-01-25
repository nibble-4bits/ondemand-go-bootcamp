package entity

// Pokemon is a struct that represents the features of a pokemon
type Pokemon struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Type1      string `json:"type1"`
	Type2      string `json:"type2"`
	Total      int    `json:"total"`
	HP         int    `json:"hp"`
	Attack     int    `json:"attack"`
	Defense    int    `json:"defense"`
	SpAtk      int    `json:"sp_attack"`
	SpDef      int    `json:"sp_defense"`
	Speed      int    `json:"speed"`
	Generation int    `json:"generation"`
	Legendary  bool   `json:"legendary"`
}
