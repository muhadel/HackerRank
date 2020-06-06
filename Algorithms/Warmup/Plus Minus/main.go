package main

import (
	"fmt"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	arrLength := float64(len(arr))
	var positive, negative, zero int
	for _, value := range arr {
		if value > 0 {
			positive++
		} else if value < 0 {
			negative++
		} else {
			zero++
		}
	}

	fmt.Println(float64(positive) / arrLength)
	fmt.Println(float64(negative) / arrLength)
	fmt.Println(float64(zero) / arrLength)
}

func main() {
	arr := []int32{1, 1, 0, -1, -1}
	plusMinus(arr)

}
