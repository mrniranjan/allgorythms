//constants

package main

import "fmt"

func main() {

	const pi float64 = 3.14
	radius := 4.0
	area := pi * radius * radius
	fmt.Printf("area of circle with radius %f is: %f\n", radius, area)
	fmt.Printf("Address of variable radius is: %d\n", &radius)
	fmt.Printf("Address of varialbe area is: %d\n", &area)
	// this fails because we can't access constants address
	//fmt.Printf("Address of constant value pi is: %d", &pi)
}
