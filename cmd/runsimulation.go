package cmd

import (
	"fmt"
	"math/rand"
	"sort"
)

func (sim *Simulation) RunSimulation() error {

	fmt.Println("------Simulation has begun Running, buckle up Earth------")

	for ; sim.Iteration < sim.EndIteration; sim.Iteration++ {
		fmt.Println()
		fmt.Println(" Itteration number: ", sim.Iteration)
		picks := RemixArray(len(sim._aliens), sim.R)

		alienMakeMoves := true

		for _, pick := range picks {
			if err := sim.AlienMovementSimulation(sim._aliens[pick]); err != nil {
				if _, okay := err.(*AlienMovingStatusError); okay {
					continue
				}
				return err
			}
			alienMakeMoves = true
		}
		if alienMakeMoves {
			fmt.Println("Simulation Ended Early at : ", sim.Iteration)
			return nil
		}
	}
	return nil
}

func RemixArray(l int, r *rand.Rand) []int {

	rangeValue := make([]int, l)

	for i := range rangeValue {
		rangeValue[i] = i
	}

	for len(rangeValue) > 0 {
		n := len(rangeValue)
		randomNumber := r.Intn(n)
		rangeValue[n-1], rangeValue[randomNumber] = rangeValue[randomNumber], rangeValue[n-1]
		rangeValue = rangeValue[:n-1]
	}

	return rangeValue
}

func (sim *Simulation) AlienMovementSimulation(alien *_alien) error {

	from, to, err := sim.MakeMoveToandForm(alien)
	fmt.Println(" Moving Alien: ", alien.Name)
	fmt.Println("to: ", to)
	fmt.Println("from: ", from)
	if err != nil {
		if operation, okay := err.(*AlienMovingStatusError); okay {
			switch operation.reason {
			case _cityDistroyed:
				fmt.Println(_operation, "All Cities been destroyed!")
			case _alienTrapped:
				fmt.Println(_operation, "Alien is trapped")
			case _deadAlien:
				fmt.Println(_operation, "Alien is dead")
			}
		}
	}
	return err
}

//Decide
func (sim *Simulation) MakeMoveToandForm(alien *_alien) (*_city, *_city, error) {

	from := alien.city
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

func (sim *Simulation) NextConnectedCity(alien *_alien) *_city {

	if !alien.AlienInvading() {
		return nil
	}

	shufflePicks := LenghtMix(len(alien.city.Links), sim.R)

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

func AlienStatus(alien *_alien) *AlienMovingStatusError {

	if alien.AlienDead() {
		return &AlienMovingStatusError{_deadAlien}
	}
	if alien.AlienTrapped() {
		return &AlienMovingStatusError{_alienTrapped}
	}
	return nil

}

func (alien *_alien) AlienDead() bool {
	return alien.Flags[_dead]
}

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

func (alien *_alien) AlienInvading() bool {
	return alien.Node != nil
}

func (city *City) CityDestroyed() bool {
	return city.Flags[_destroyed]
}

func (city *City) DestroyCity() {
	city.Flags[_destroyed] = true
}
