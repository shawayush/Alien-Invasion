package cmd

import (
	"fmt"
	"math/rand"
)

func Mix(vals []int, r *rand.Rand) {

	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func LenghtMix(length int, r *rand.Rand) []int {

	vals := MakeRange(0, length)
	Mix(vals, r)
	return vals
}

func MakeRange(min, max int) []int {

	vals := make([]int, max-min)
	for i := range vals {
		vals[i] = min + i
	}
	return vals
}

//used to check and print for the city that is not destroyed
func (err *AlienMovingStatusError) Error() string {
	return fmt.Sprintf("Simulator stopped as :", err.reason)
}
