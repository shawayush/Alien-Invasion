package cmd

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func CreateRandomAliens(n int, r *rand.Rand) []*Alien {

	aliensList := make([]*Alien, n)
	for i := 0; i < n; i++ {
		name := strconv.Itoa(r.Int())[:_numberOfRandonAlien]
		alienNameAndAttributes := CreateNewAlien(name)
		aliensList[i] = &alienNameAndAttributes

	}
	return aliensList
}

func CreateNewAlien(name string) Alien {

	return Alien{
		Status: Status{Name: name,
			Flags: make(map[string]bool)},
	}
}

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

func CreateAttackingAliens() *rand.Rand {

	seed := time.Now().UnixNano()
	randomNumber := rand.NewSource(seed)
	randomNumberSource := rand.New(randomNumber)
	return randomNumberSource
}

func IntiateNewSimulation(r *rand.Rand, lastItteration int, world _world, numberOfAlien _aliens) Simulation {

	return Simulation{
		R:            r,
		Iteration:    0,
		EndIteration: lastItteration,
		_world:       world,
		_aliens:      numberOfAlien,
		_defense:     make(_defense),
	}
}
