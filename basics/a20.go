// difference between sprintf and sprintln
package main

import "fmt"

func main() {
	const pi = 3.14
	radius := 15
	area := pi * radius * radius
	s1 := fmt.Sprintf("Area of circle with %d radius is %d\n", radius, area)
	s2 := fmt.Sprintln("Area of circle with", radius, "radius is", area)
	fmt.Println(s1)
	fmt.Println(s2)
}
	
	
