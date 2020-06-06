package main

import (
	"fmt"
)

// Complete the aVeryBigSum function below.
func aVeryBigSum(ar []int64) int64 {
	var total int64
	for _, value := range ar {
		total += value
	}
	return total
}

func main() {
	arr := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	res := aVeryBigSum(arr)
	fmt.Println(res)
}
