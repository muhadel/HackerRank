package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	var timeConverted string
	// convert hour to integer
	hour, _ := strconv.Atoi(s[0:2])
	min := s[3:5]
	sec := s[6:8]
	meridiem := s[8:10]

	// fmt.Println(hour, min, sec, meridiem)
	if strings.EqualFold(meridiem, "PM") && hour != 12 {
		hour += 12

	} else if strings.EqualFold(meridiem, "AM") && hour == 12 {
		hour -= 12
	}

	timeConverted = fmt.Sprintf("%02d", hour) + ":" + min + ":" + sec
	return timeConverted
}

func main() {
	fmt.Print(timeConversion("12:00:00AM"))
}
