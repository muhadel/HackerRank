package main

import (
	"fmt"
)

/*
 * Complete the getMoneySpent function below.
 */
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var maximumTotalPrice int32 = -1
	for _, keyboard := range keyboards {
		for _, driver := range drives {
			sumOfItems := keyboard + driver
			if sumOfItems <= b && sumOfItems > maximumTotalPrice {
				maximumTotalPrice = sumOfItems
			}
		}
	}
	return maximumTotalPrice
}
func main() {
	// Test case 1
	keyboards := []int32{3, 1}
	drives := []int32{5, 2, 8}
	var budget int32 = 10

	// Test case 2
	// keyboards := []int32{4}
	// drives := []int32{5}
	// var budget int32 = 5

	fmt.Println(getMoneySpent(keyboards, drives, budget))
}
