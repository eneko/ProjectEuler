// problem19.go
package main

import (
	"fmt"
)

//1:Monday
//2
//3
//4
//5
//6
//7:Sunday

func isLeap(year int) bool {
	return (year%100 == 0 && year%400 == 0) || (year%100 != 0 && year%4 == 0)
}

func contains(list []int, value int) bool {
	for _, a := range list {
		if a == value {
			return true
		}
	}
	return false
}

func main() {
	// Days of year that are first of the month
	nonLeapYear := []int{1, 32, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335}
	leapYear := []int{1, 32, 61, 92, 122, 153, 183, 214, 245, 275, 306, 336}

	sundaysOnFirst := 0
	year := 1900
	day := 7 // First Sunday of 1900
	leap := isLeap(year)

	for year < 2001 {
		// Check if day is first of the month
		if year > 1900 {
			if (leap && contains(leapYear, day)) || (!leap && contains(nonLeapYear, day)) {
				fmt.Printf("Year: %v, Day: %v is first day of month!\n", year, day)
				sundaysOnFirst++
				// } else {
				// fmt.Printf("Year: %v, Day: %v\n", year, day)
			}
		}

		// Next Sunday
		day += 7

		// Increase year
		if !leap && day > 365 {
			// fmt.Println("Non leap year ended")
			day = day % 365
			year++
			leap = isLeap(year)
		} else if leap && day > 366 {
			// fmt.Println("Leap year ended")
			day = day % 366
			year++
			leap = isLeap(year)
		}
	}

	fmt.Printf("Sundays on first of the month: %v\n", sundaysOnFirst)
}
