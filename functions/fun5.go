/* Passing integer pointers to functions */
package main

import (
	"fmt"
)

const (
	km = 1000
	gm = 1000
)

func main() {
	var num, distance, weight, m = 20,3, 50,  "mm"
	a := mmToCm(&num)
	b := distanceInMeters(&distance)
	c := weightInMilligrams(&weight)
	fmt.Printf("%d %s is %d cm\n", num, m, a)
	fmt.Printf("%d Meters was travelled\n", b)
	fmt.Printf("%d Milligrams is %d grams\n", c , weight)
}

func mmToCm(p * int) int {
	return *p / 10
}

func distanceInMeters(p * int) int {
	return *p * km 
}

func weightInMilligrams(p * int) int {
	return *p * gm
}
