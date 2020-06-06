package main

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(ar []int32) int32 {
	var maxOfArray int32
	birth := make(map[int32]int32)

	for _, value := range ar {
		if birth[value] == 0 {
			birth[value] = 1
		} else {
			birth[value]++
		}
		// get largest age
		if value > maxOfArray {
			maxOfArray = value
		}

	}

	return birth[maxOfArray]
}

func main() {
	arr := []int32{3, 2, 1, 3}
	birthdayCakeCandles(arr)
}
