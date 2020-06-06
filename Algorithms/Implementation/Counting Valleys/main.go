package main

import (
	"fmt"
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	var valleys, seaLevel int32 = 0, 0
	
	for _, rune := range s {
		ch := string(rune)
		fmt.Println("seaLevel" , seaLevel)
		if ch == "D" { 
			if seaLevel == 0 { valleys++ }
			seaLevel--
			
		 }
		if ch == "U" { seaLevel++ }
	}
	return valleys
}

func main()  {
	// my current rank 985005
	// var steps1 = 8
	// var steps string = "UDDDUDUU"
	// DDUUUUDD
	fmt.Println(countingValleys(8,"DDUUUUDD"))

}