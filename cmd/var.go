package cmd

//used for logic for connecting the roads between each
var _directions = map[string]string{
	"north": "south",
	"south": "north",
	"east":  "west",
	"west":  "east",
}

const (
	_numberOfRandomAlien int = 10 //used for creating random aliesn

	_dead string = "dead" //used for marking the alien is dead

	_deadAlien AlienMovingStatus = iota //default dead alien, while checking operations

	_alienTrapped //variable used for checking the aline is trapped

	_cityDistroyed //to check if city is destroyed

	_destroyed string = "destroyed" //used to put flag in the city is destroyed

	_operation string = "<Process stopped>" //used to check for the process like, alien dead, city destroyed, etc
)

type (
	_occupation map[string]*Alien //check for occupation used in _defence

	_defense map[string]_occupation //city defence

	_cityTxtFile []*City //putting city in a array/ list

	_world map[string]*City //creating a world for the whole city

	_aliens []*Alien //list of aliens

	_alien = Alien //variable used to do operation in alien

	_city = City //variable ised to operation in city

	_cityMapFile []*_city //list of city
)
