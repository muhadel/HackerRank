package main

import (
	"fmt"
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	var valleys, seaLevel int32 = 0, 0
	
	for _, rune := range s {
		ch := string(rune)
		if ch == "D" { 
			if seaLevel == 0 { valleys++ }
			seaLevel--
		 }
		if ch == "U" { seaLevel++ }
	}
	return valleys
}

func main()  {
	// DDUUUUDD
	fmt.Println(countingValleys(8,"DDUUUUDD"))

}