package cmd

import (
	"fmt"
	"sort"
)

//Initiater to run simulation,calls all different types of function to make
//simulation possible, like mxing the aliens and then, checking for simulation
//status and so on.
func (sim *Simulation) RunSimulation() error {

	fmt.Println("------Simulation has begun Running, buckle up Earth------")

	for ; sim.Iteration < sim.EndIteration; sim.Iteration++ {
		fmt.Println("\n _____Iteration number:", sim.Iteration+1, "_____")

		picks := LenghtMix(len(sim._aliens), sim.R)

		alienMakeMoves := true

		for _, pick := range picks {
			if err := sim.AlienMovementSimulation(sim._aliens[pick]); err != nil {
				if _, okay := err.(*AlienMovingStatusError); okay {
					continue
				}
				return err
			}
			alienMakeMoves = false
		}
		if alienMakeMoves {
			fmt.Println("Simulation Ended Early at : ", sim.Iteration+1)
			return nil
		}
	}
	return nil
}

//create a alien mobment movement check to check the operations is working properly or not
//also creates checks, creates, destroys, kill or destroy the aliens
func (sim *Simulation) AlienMovementSimulation(alien *_alien) error {

	from, to, err := sim.MakeMoveToandForm(alien)
	fmt.Println("Moving Alien: ", alien.Name)
	fmt.Println("to: ", to)
	fmt.Println("from: ", from)
	if err != nil {
		if operation, okay := err.(*AlienMovingStatusError); okay {
			switch operation.reason {
			case _cityDistroyed:
				fmt.Println(_operation, " All Cities been destroyed!")
			case _alienTrapped:
				fmt.Println(_operation, " Alien is trapped")
			case _deadAlien:
				fmt.Println(_operation, " Alien is dead")
			}
		}
		return err
	}
	//if there are Process to go, then the city invades
	alien.InvadeCity(to)

	//logic for invading and processing the city after wards
	if from != nil {
		delete(sim._defense[from.Name], alien.Name)
	}

	if sim._defense[to.Name] == nil {
		sim._defense[to.Name] = make(_occupation)
	}

	sim._defense[to.Name][alien.Name] = alien

	if len(sim._defense[to.Name]) > 1 {
		to.DestroyCity()

		output := fmt.Sprintf("%s has been destroyed by ", to.Name)
		for _, alien := range sim._defense[to.Name] {
			output += fmt.Sprintf("alien %s and ", alien.Name)
			alien.Kill()
		}

		output = output[:len(output)-5] + "!\n" //trim out the extra string that is not to be used
		fmt.Println(output)
	}
	return nil
}

//This function would let the aliens decide where should the alien go.
//if it is starting from no where, then it could be random city
// if it already in the city and there is a connection to go from
// to city to another then the alien makes the move
func (sim *Simulation) MakeMoveToandForm(alien *_alien) (*_city, *_city, error) {

	from := alien.AlienCity()
	if err := AlienStatus(alien); err != nil {
		return from, nil, err
	}

	//Start from begining
	if from == nil {
		to := sim.PickCityRandom()
		if to == nil {
			return from, to, &AlienMovingStatusError{reason: _cityDistroyed}
		}
		return from, to, nil
	}

	//need to move to the connected city
	to := sim.NextConnectedCity(alien)

	return from, to, nil

}

//check whether the city that alien is in already has a connected city or not
//that means if there is a connected city to this then a line need to make a proper move
// or start the process all over again
func (sim *Simulation) NextConnectedCity(alien *_alien) *_city {

	if !alien.AlienInvading() {
		return nil
	}

	shufflePicks := LenghtMix(len(alien.AlienCity().Links), sim.R)
	for _, pick := range shufflePicks {
		val := alien.city.Links[pick].Key
		node := alien.city.Nodes[val]
		if city := sim._world[node.Name]; !city.CityDestroyed() {
			return city
		}

	}
	return nil

}

//PickCityRandom picks any undestroyed city the in world
//Catch is to make deterministically, and sort it out so that
//it could be as deterministic as it could be, so idea
//behind it would be pick at random got some idea from the
//consensus mechanisum that COSMOS SDK follows
func (sim *Simulation) PickCityRandom() *_city {

	var keys []string
	for key := range sim._world {
		if city := sim._world[key]; !city.CityDestroyed() {
			keys = append(keys, key)
		}
	}

	if len(keys) == 0 {
		return nil
	}

	sort.Strings(keys)
	pick := sim.R.Intn(len(keys))
	return sim._world[keys[pick]]
}

//-------------------------------Alien Actions-------------------------------

//check alien status if it is trapped or dead!
func AlienStatus(alien *_alien) *AlienMovingStatusError {

	if alien.AlienDead() {
		return &AlienMovingStatusError{_deadAlien}
	}
	if alien.AlienTrapped() {
		return &AlienMovingStatusError{_alienTrapped}
	}
	return nil

}

// Check if the alien is dead or not used in AlienStatus
func (alien *Alien) AlienDead() bool {

	return alien.Flags[_dead]
}

// Checks the  status of alien at which city it is
func (alien *Alien) AlienCity() *City {

	return alien.city
}

// alien invades the city
func (alien *Alien) InvadeCity(city *City) {

	alien.Node = &city.Node
	alien.city = city
}

//kill the alien
func (a *Alien) Kill() {

	a.Flags[_dead] = true
}

//check whether is Alien is Invading or not
func (alien *Alien) AlienInvading() bool {
	return alien.Node != nil
}

//Proper formatting for alien string
func (alien *Alien) String() string {
	return fmt.Sprintf("name=%s city={%s}\n", alien.Name, alien.city)
}

// Check if the alien is trapped or not used in AlienStatus
func (alien *Alien) AlienTrapped() bool {
	if !alien.AlienInvading() {
		return false
	}
	for _, n := range alien.city.Nodes {
		city := City{Node: *n}
		if !city.Flags[_destroyed] {
			return false
		}
	}
	return true
}

//-------------------------------City Actions-------------------------------

//check whether the city is destroyed or not
func (city *City) CityDestroyed() bool {

	return city.Flags[_destroyed]
}

//function use to destroy city
func (city *City) DestroyCity() {

	city.Flags[_destroyed] = true
}

//putting the city in the string format
func (city *City) String() string {

	var links string
	for _, link := range city.Links {
		newcity := city.Nodes[link.Key]
		otherCity := City{Node: *newcity}

		if otherCity.CityDestroyed() {
			continue
		}

		links += fmt.Sprintf("%s=%s ", city.RoadsName[link.Key], otherCity.Name)
	}

	if len(links) == 0 {
		return city.Name
	}
	return fmt.Sprintf("%s %s", city.Name, links[:len(links)-1])
}

//-------------------------------cityMapFile Actions-------------------------------

//checks for the city that has been destroyed and filter the city from the original
//text file so that it can be used print the ciuty remainings
func (cityMapFile _cityMapFile) FilterCitiesDestroyed(cities _world) _cityMapFile {

	cityOutput := make(_cityTxtFile, 0, len(cityMapFile))
	checkCityStatus := make(map[string]bool)

	for _, city := range cityMapFile {

		if checkCityStatus[city.Name] {
			continue
		}

		if !city.CityDestroyed() {
			cityOutput = append(cityOutput, city)
			checkCityStatus[city.Name] = true
			continue
		}

		for _, link := range city.Links {
			new := city.Nodes[link.Key]
			differentCity := cities[new.Name]
			if checkCityStatus[differentCity.Name] || differentCity.CityDestroyed() {
				continue
			}

			cityOutput = append(cityOutput, differentCity)
			checkCityStatus[differentCity.Name] = true
		}
	}
	return _cityMapFile(cityOutput)
}

func (in _cityMapFile) String() string {
	var str string
	for _, city := range in {
		if city.CityDestroyed() {
			continue //city destroyed, so required no action
		}

		str += fmt.Sprintf("%s\n", city) //city is there need to give the name
	}
	return str
}

//-------------------------------------------------------------------------------

func (w _world) String() string {
	var str string
	for _, city := range w {
		str += fmt.Sprintf("%s\n", city)
	}
	return str
}

//used to check and print for the city that is not destroyed
func (err *AlienMovingStatusError) Error() string {
	return fmt.Sprintf("Simulator stopped as :", err.reason)
}
