package cmd

var _directions = map[string]string{
	"north": "south",
	"south": "north",
	"east":  "west",
	"west":  "east",
}

const (
	_numberOfRandonAlien int = 10

	_dead string = "dead"

	_deadAlien AlienMovingStatus = iota

	_alienTrapped

	_cityDistroyed

	_destroyed string = "destroyed"

	_operation string = "stopped"
)

type (
	_occupation map[string]*Alien

	_defense map[string]_occupation

	_cityTxtFile []*City //putting city in a array/ list

	_world map[string]*City

	_aliens []*Alien

	_alien = Alien

	_city = City

	_cityMapFile []*_city
)
