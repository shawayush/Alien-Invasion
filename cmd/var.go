package cmd

var _directions = map[string]string{
	"north": "south",
	"south": "north",
	"east":  "west",
	"west":  "east",
}

const _numberOfRandonAlien int = 10

type (
	_occupation map[string]*Alien

	_defense map[string]_occupation

	_cityTxtFile []*City //putting city in a array/ list

	_world map[string]*City

	_aliens []*Alien
)
