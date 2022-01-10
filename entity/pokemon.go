package entity

type Pokemon struct {
	Id                                                          int
	Name, Type1, Type2                                          string
	Total, HP, Attack, Defense, SpAtk, SpDef, Speed, Generation int
	Legendary                                                   bool
}
