package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

//ReadAndMakeWorldMap function reads the city file and create links between
//the city so that it could create a world for aliens to Invade
func ReadAndMakeWorldMap(file string) (_world, _cityMapFile, error) {

	readableFile, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	defer readableFile.Close()

	//using scanner to read the contect of the files

	scanner := bufio.NewScanner(readableFile)
	scanner.Split(bufio.ScanLines)

	worlds := make(_world)
	textInput := make(_cityMapFile, 0)

	for scanner.Scan() {
		//break the file into sentences
		sentences := strings.Split(scanner.Text(), " ")
		//pick city for
		cities := worlds.AddNewCity(sentences[0])
		//start building raods/links that is need to used throughout
		for _, words := range sentences[1:] {
			road, city := BuildLink(words)
			//checks if there is other city is connected well enough or not
			otherCity, Cityexists := worlds[city]
			fmt.Println("")
			if !Cityexists {
				otherCity = worlds.AddNewCity(city)
			}
			//logic for linking roads between cities
			link := cities.ConnectTwoCities(&otherCity.Node)
			otherCity.ConnectRoad(link, &cities.Node)
			cities.RoadsName[link.Key] = road
			otherCity.RoadsName[link.Key] = _directions[road]
		}

		textInput = append(textInput, cities)

	}

	return worlds, textInput, nil

}

//Makes connection between two cities
func (n *Node) ConnectTwoCities(other *Node) *Link {

	link := FormLinkBetweenCities(n.Name, other.Name)
	return n.ConnectRoad(&link, other)

}

//used in ConnectTwoCities function. This function helps to create
//link between two cities or nodes
func FormLinkBetweenCities(nodes ...string) Link {

	sort.Strings(nodes)
	key := strings.Join(nodes, "_")
	return Link{key, nodes}
}

//connect one city through a path/road
func (n *Node) ConnectRoad(link *Link, other *Node) *Link {

	if n.Nodes[link.Key] == nil {
		n.Links = append(n.Links, link)
		n.Nodes[link.Key] = other
	}
	return link
}

//Read the input from the file that it could be used to create
//takes the input from the file, make connections as required by the
// inputs given
func BuildLink(words string) (string, string) {

	word := strings.Split(words, "=")
	road, connectingCity := word[0], word[1]
	return road, connectingCity
}

//Creates a graph like strucutre for adding the city in the worldmap
func (w _world) AddNewCity(name string) *City {

	return w.AddCity(City{
		Node: Node{
			Name:  name,
			Links: make([]*Link, 0),
			Nodes: make(map[string]*Node),
			Flags: make(map[string]bool),
		},
		RoadsName: make(map[string]string),
	})
}

//function used in AddNewCity for adding city the strucutre that has been created
func (w _world) AddCity(city City) *City {

	w[city.Name] = &city
	return &city
}

//  String representation for a Node
func (n *Node) String() string {
	var links string
	for k, n := range n.Nodes {
		links += fmt.Sprintf("%s:%s ", k, n.Name)
	}
	if len(n.Nodes) > 0 {
		links = links[:len(links)-1]
	}
	return fmt.Sprintf("name=%s links=map[%s]", n.Name, links)
}

func (l Link) String() string {
	return fmt.Sprintf("key=%s nodes=%s\n", l.Key, l.Nodes)
}
