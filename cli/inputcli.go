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
	_cityTxtFile       = "./test/ex.txt"
	_alienNameFile     = "./test/"
)

func Execute() {

	if err := CheckCliInputs(); err != nil {
		fmt.Println("Error while checking flags: %s\n", err)
		os.Exit(1)
	}
	/*
		//PLAN : -
		//writng logic for taking the files input then mapping into the world simulation
			//take the input from the file, parse it --> parse the map, (Like connecting the nodes)
			// build the city with nodes, connecting Noth, East, West, South
	*/
	world, inputfile, err := cmd.ReadAndMakeWorldMap(_cityTxtFile)
	if err != nil {
		//add error handling
	}
	fmt.Println(world)
	fmt.Println(inputfile)
	/*

		// writing a logic for building a simulatior
			//use random from to create itterations both random numbers and random alines

		//run the simultion
			//create a loop for itteration
			//make the alines shuffle
				//create another loop
				//make alien moves
				//delete if the city is destroyed
				//add a checker if the simulation completes before
				//

		//print the city remainig
	*/

}

func init() {

	flag.IntVar(&_iterationsInput, "iterations", _itterations, "number of iterations to simulate")
	flag.IntVar(&_numberOfAliens, "aliens", _aliens, "number of aliens invading during an invasion")
	flag.StringVar(&_inputCityTxtFile, "world", _cityTxtFile, "a file used as world map input to make simulation")
	flag.StringVar(&_inputalienNameFile, "Alien", _alienNameFile, "a file used to identify aliens")
	flag.Parse()

}

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
