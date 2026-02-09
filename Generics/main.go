package main

import "fmt"

func main() {
	// Intitialize a map for the integer values
	ints := map[string]int64{
		"first":  54,
		"second": 32,
	}

	// Intitialize a map for the float values
	floats := map[string]float64{
		"first":  43.32,
		"second": 65.9,
	}

	fmt.Printf("Non-Generic Sums: %v & %v\n",
		SumInt(ints),
		SumFloat(floats))

	fmt.Printf("Generic Sums: %v & %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))
}

// SumInt adds together the values of m
func SumInt(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloat adds together the values of m
func SumFloat(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both int64 & float64
// as types for the map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
