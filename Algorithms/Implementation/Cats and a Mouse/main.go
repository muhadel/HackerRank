package main

import (
	"fmt"
	"math"
)

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
	var diffA = math.Abs(float64(z - x))
	var diffB = math.Abs(float64(z - y))
	if diffA < diffB {
		return "Cat A"
	} else if diffB < diffA {
		return "Cat B"
	}
	return "Mouse C"
}
func main() {
	// var catA, catB, Mouse int32 = 1, 2, 3
	var catA, catB, Mouse int32 = 1, 3, 2
	fmt.Print(catAndMouse(catA, catB, Mouse))
}
