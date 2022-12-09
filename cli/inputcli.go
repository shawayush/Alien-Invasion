package cli

import (
	"errors"
	"fmt"
)

var (
	_iterations, _numberOfAliens           int
	_inputCityTxtFile, _inputalienNameFile string
)

const (
	_aliens        int = 10    // using 10 Aleins if not scepified
	_itterations   int = 10000 // using 10000 itterations as specified (or) could input yours if required
	_cityTxtFile       = "./test/"
	_alienNameFile     = "./test/"
)

func Execute() {

	if err := CheckCliInputs(); err != nil {
		fmt.Println("Error while checking flags: %s\n", err)

	}

}

func init() {
	// to do : add flag for cli
}

func CheckCliInputs() error {

	if _numberOfAliens <= 0 {
		return errors.New("Number of Aliens should be greater than 0")
	}

	if _iterations <= 0 {
		return errors.New("Iterations should be a positive number!")
	}

	if len(_inputCityTxtFile) == 0 {
		return errors.New("World map file path is empty")
	}

	return nil

}
