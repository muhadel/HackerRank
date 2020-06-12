package main

import (
	"fmt"
)

func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int32, n)
	var max int64
	for _, i := range queries {
		for j := i[0] - 1; j < i[1]; j++ {
			arr[j] += i[2]
			if int64(arr[j]) > max {
				max = int64(arr[j])
			}
		}

	}
	return max
}
func main() {
	queries := [][]int32{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	}
	var n int32 = 10
	fmt.Print(arrayManipulation(n, queries))
}
