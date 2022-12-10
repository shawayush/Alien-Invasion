package cmd

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
