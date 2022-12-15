package cmd

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//Creates Random Aliens used thorughout the invasion
func CreateRandomAliens(n int, r *rand.Rand) []*Alien {

	aliensList := make([]*Alien, n)

	for i := 0; i < n; i++ {
		name := strconv.Itoa(r.Int())[:_numberOfRandomAlien]
		alienNameAndAttributes := CreateNewAlien(name)
		aliensList[i] = &alienNameAndAttributes
	}
	return aliensList
}

//Create new alien also put a var(status) so that there status could be monitered
func CreateNewAlien(name string) Alien {

	return Alien{
		Status: Status{Name: name,
			Flags: make(map[string]bool)},
	}
}

//function used to get alien name from the file
func NameAliensFromFile(aliens []*Alien, file string) error {

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for i := 0; i < len(aliens) && scanner.Scan(); i++ {
		aliens[i].Name = scanner.Text()
	}
	return nil
}

//Create a randmoness of n number of string so that it
//could be used for aliens that is not named
func CreateAttackingAliens() *rand.Rand {

	seed := time.Now().UnixNano()
	randomNumber := rand.NewSource(seed)
	randomNumberSource := rand.New(randomNumber)
	return randomNumberSource
}

// function used to create a variable that is used to initiate a new simulator
// whenever a function is called for.
func IntiateNewSimulation(r *rand.Rand, lastIteration int, world _world, numberOfAlien _aliens) Simulation {

	return Simulation{
		R:            r,
		Iteration:    0,
		EndIteration: lastIteration,
		_world:       world,
		_aliens:      numberOfAlien,
		_defense:     make(_defense),
	}
}
