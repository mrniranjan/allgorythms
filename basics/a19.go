package main

import "fmt"

func main() {
	radius := 15
	area := radius * radius
	s := fmt.Sprintf("Area of circle with %d radius is %d", radius, area)
	fmt.Println(s)
}
