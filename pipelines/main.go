package main

import "fmt"

func main() {
	// Stage 1: Takes in a slice of integers, multiplies them by a multiplier and returns a slice of multiplied integers
	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, value := range values {
			multipliedValues[i] = value * multiplier
		}
		return multipliedValues
	}

	//Stage 2: Takes in a slice of integers, adds them to an additive and returns a slice of added integers
	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, value := range values {
			addedValues[i] = value + additive
		}
		return addedValues
	}

	// 3. Combining add and multiply stages within a range clause
	integers := []int{1, 2, 3, 4}
	for _, v := range add(multiply(integers, 2), 1) {
		fmt.Println(v)
	}

}
