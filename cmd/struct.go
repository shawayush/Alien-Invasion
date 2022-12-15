package cmd

import "math/rand"

//used for city values and attributes
type City struct {
	Node
	RoadsName map[string]string
}

//used for city attributes
type Node struct {
	Name  string
	Links []*Link
	Nodes map[string]*Node
	Flags map[string]bool
}

//Represents the connection between the node used above
type Link struct {
	Key   string
	Nodes []string
}

//Representation for alien status
type Status struct {
	Name  string
	Flags map[string]bool
	Node  *Node
}

//struct for Alien and there attributes
type Alien struct {
	Status
	city *City
}

//type used to run simulation
type Simulation struct {
	R            *rand.Rand
	Iteration    int
	EndIteration int
	_world
	_aliens
	_defense
}

//
type AlienMovingStatus uint8

type AlienMovingStatusError struct {
	reason AlienMovingStatus
}
