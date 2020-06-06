package main

import (
	"fmt"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	var minimum, maximum = arr[0], arr[0]
	var sumOfArray int64 = 0
	for _, value := range arr {
		if value < minimum {
			minimum = value
		}
		if value > maximum {
			maximum = value
		}
		sumOfArray += int64(value)
	}
	fmt.Print(sumOfArray-int64(maximum), sumOfArray-int64(minimum))
}
func main() {
	arr := []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr)
}
