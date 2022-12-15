package cmd

import (
	"math/rand"
)

//--------------funcitons used for mixing arrays------------

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
