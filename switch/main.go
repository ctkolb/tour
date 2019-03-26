package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Thursday

	fmt.Println("When's", day, "?")
	now := time.Now()
	switch day {
	case now.Weekday() + 0:
		fmt.Println("Today.")
	case now.Weekday() + 1:
		fmt.Println("Tomorrow.")
	case now.Weekday() + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
		var testDate time.Time
		i := 1
		for ; testDate.Weekday() != day; i++ {
			testDate = now.AddDate(0, 0, i)
		}
		fmt.Println("(Which is ", i, "days.)")

	}
}
