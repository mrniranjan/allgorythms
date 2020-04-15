package main

import "fmt"

func main() {
	var r = 5.1
	var pi = 3.15
	var area = r * r * pi
	s1 := fmt.Sprintf("Area of circle with %f radius is %f\n", r, area)
	fmt.Print(s1)
}
