package main

import (
	"fmt"
	"strings"
)

// Complete the staircase function below.
func staircase(n int32) {
	var i int32 = 1
    for i <= n {
        fmt.Print(strings.Repeat(" ", int(n - i)), strings.Repeat("#", int(i)), "\n")
        i++
    }
}

func main() {
	staircase(6)
}
