// difference between sprint, sprintf, sprintln
package main

import "fmt"

func main() {
	const pi = 3.14
	radius := 15.0
	area := pi * radius * radius
	s1 := fmt.Sprint("Area of circle with ", radius, " radius is", area)
	s2 := fmt.Sprintf("Area of circle with %f radius is %f\n", radius, area)
	s3 := fmt.Sprintln("Area of circle with", radius, "radius is", area)
	fmt.Print(s1)
	fmt.Print(s2)
	fmt.Print(s3)
}
