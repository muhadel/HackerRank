package main

import (
	"fmt"
	"math"
)


func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var d1 , d2 int32
	arrLength := len(arr)
	for index, value := range arr {
		d1 += value[index]
		d2 += value[arrLength - index - 1]
	}
	return  int32(math.Abs(float64(d1 - d2)))
}


func main()  {
	arr := [][]int32{
		{1, 2, 3},
		{4, 5, 6,},
		{9, 8, 9},
	}
	result := diagonalDifference(arr)
	fmt.Println(result)
}