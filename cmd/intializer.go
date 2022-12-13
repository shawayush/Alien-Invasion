package cmd

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

func ReadAndMakeWorldMap(file string) (_world, _cityTxtFile, error) {

	readableFile, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	defer readableFile.Close()

	//using scanner to read the contect of the files

	scanner := bufio.NewScanner(readableFile)
	scanner.Split(bufio.ScanLines)

	worlds := make(_world)
	textInput := make(_cityTxtFile, 0)

	for scanner.Scan() {
		sentences := strings.Split(scanner.Text(), " ")
		cities := worlds.AddNewCity(sentences[0])

		for _, words := range sentences[1:] {
			road, city := BuildLink(words)

			otherCity, Cityexists := worlds[city]
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
func (n *Node) ConnectTwoCities(other *Node) *Link {

	link := FormLinkBetweenCities(n.Name, other.Name)
	return n.ConnectRoad(&link, other)

}

func FormLinkBetweenCities(nodes ...string) Link {

	sort.Strings(nodes)
	key := strings.Join(nodes, "_")
	return Link{key, nodes}
}

func (n *Node) ConnectRoad(link *Link, other *Node) *Link {

	if n.Nodes[link.Key] == nil {
		n.Links = append(n.Links, link)
		n.Nodes[link.Key] = other
	}
	return link
}

func BuildLink(words string) (string, string) {

	word := strings.Split(words, "=")
	road, connectingCity := word[0], word[1]
	return road, connectingCity
}

func (w _world) AddNewCity(name string) *City {

	return w.AddCity(City{
		Node: Node{
			Name: name,
			//Links: make([]*Link, 0),
			Nodes: make(map[string]*Node),
		},
		RoadsName: make(map[string]string),
	})
}

func (w _world) AddCity(city City) *City {

	w[city.Name] = &city
	return &city
}
