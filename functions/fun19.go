// This program is to illustrate the use of switch statements 
package main

import (
	"fmt"
	"time"
	"os"
)

type weekday int

const (
	Sunday weekday = iota 
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)


func main() {
	var day int 
	fmt.Printf("\nWhat is the day today(1-7): ")
	if _, err := fmt.Scan(&day); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if getDay(&day) {
		fmt.Println("Right answer")
	} else {
		fmt.Println("Wrong answer")
	}
}
// This function accepts the pointer to an integer and returns bool 
func getDay(d *int) bool {
	var result bool 
	switch *d {
	case 1:
		result = (time.Now().Weekday() == time.Weekday(Monday))
	case 2:
		result = (time.Now().Weekday() == time.Weekday(Tuesday))
	case 3:
		result = (time.Now().Weekday() == time.Weekday(Wednesday))
	case 4:
		result = (time.Now().Weekday() == time.Weekday(Thursday))
	case 5:
		result = (time.Now().Weekday() == time.Weekday(Friday))
	case 6:
		result = (time.Now().Weekday() == time.Weekday(Saturday))
	case 7:
		result = (time.Now().Weekday() == time.Weekday(Sunday))
	default:
		fmt.Println("Wrong Input")
		return false
	}
	return result
}
