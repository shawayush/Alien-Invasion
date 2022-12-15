package cli

import (
	"Alien-Invasion/cmd"
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	_iterationsInput, _numberOfAliens      int
	_inputCityTxtFile, _inputalienNameFile string
)

const (
	_aliens        int = 10    // using 10 Aleins if not specified
	_itterations   int = 10000 // using 10000 itterations as specified (or) could input yours if required
	_cityTxtFile       = "./test/example4.txt"
	_alienNameFile     = "./test/aliens.txt"
)

func Execute() {

	if err := CheckCliInputs(); err != nil {

		fmt.Println("Error while checking flags: %s\n", err)
		os.Exit(1)
	}

	//writng logic for taking the files input then mapping into the world simulation
	//take the input from the file, parse it --> parse the map, (Like connecting the nodes)
	// build the city with nodes, connecting Noth, East, West, South
	createdCity, inputFile, err := cmd.ReadAndMakeWorldMap(_cityTxtFile)
	if err != nil {
		fmt.Println("checking")
		fmt.Println("Error while reading city Input File ", err)
		os.Exit(1)
	}

	// writing a logic for building a simulatior
	//use random from to create itterations both random numbers and random alines
	CreateAttackingAliens := cmd.CreateAttackingAliens()                              //creates random intigers of number to be used by the aliens
	attackingAliens := cmd.CreateRandomAliens(_numberOfAliens, CreateAttackingAliens) //assign those values to the aliens, name provided in alien.txt file
	if _alienNameFile != "" {
		if err := cmd.NameAliensFromFile(attackingAliens, _alienNameFile); err != nil {
			fmt.Println("Error while reading alien Input File ", err)
			os.Exit(1)
		}
	}

	//A variable used to initiate simulation, which creates a value for the itteration throughouts, where it could be stored and
	//used to moniotr and definitely run the simulation in the required way as possible
	startSimulation := cmd.IntiateNewSimulation(CreateAttackingAliens, _iterationsInput, createdCity, attackingAliens) //start simulation

	//run the simultion
	//create a loop for itteration
	//make the alines shuffle
	//create another loop
	//make alien moves
	//delete if the city is destroyed
	//add a checker if the simulation completes before
	if err := startSimulation.RunSimulation(); err != nil {
		fmt.Errorf("Error while running simulation: ", err)
		os.Exit(1)
	}

	fmt.Println("\n<><><><><><><><><> SIMULATION SUCCESS! <><><><><><><><><>")

	//print the city remainig
	remainingCity := inputFile.FilterCitiesDestroyed(createdCity)
	if len(remainingCity) != 0 {
		fmt.Println("\nEarth has been hanging there :) Remaining Cities are: ")
		fmt.Println(remainingCity)
	} else {
		fmt.Println("\nEarth has been destroyed :((")
	}

}

//Initilaize the cli parameters need to be given
//could be used to after using go.main.go -help
func init() {

	flag.IntVar(&_iterationsInput, "iterations", _itterations, "number of iterations to simulate")
	flag.IntVar(&_numberOfAliens, "aliens", _aliens, "number of aliens invading during an invasion")
	flag.StringVar(&_inputCityTxtFile, "city", _cityTxtFile, "a file used as city input to make simulation")
	flag.StringVar(&_inputalienNameFile, "Alien", _alienNameFile, "a file used to identify aliens")
	flag.Parse()

}

//func for checking the if the inputs in the cli
//is right or not!
func CheckCliInputs() error {

	if _numberOfAliens <= 0 {
		return errors.New("Number of Aliens should be greater than 0")
	}

	if _iterationsInput <= 0 {
		return errors.New("Iterations should be a positive number!")
	}

	if len(_inputCityTxtFile) == 0 {
		return errors.New("World map file path is empty")
	}

	return nil

}
