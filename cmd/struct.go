package cmd

import "math/rand"

type City struct {
	Node
	RoadsName map[string]string
}

//TODO : Needs to build a flag which can say that survived or not during the invasion
type Node struct {
	Name  string
	Links []*Link
	Nodes map[string]*Node
}

//Represents the connection between the node used above
type Link struct {
	Key   string
	Nodes []string
}

type Status struct {
	Name  string
	Flags map[string]bool
	Node  *Node
}

type Alien struct {
	Status
	city *City
}

type Simulation struct {
	R            *rand.Rand
	Iteration    int
	EndIteration int
	_world
	_aliens
	_defense
}
