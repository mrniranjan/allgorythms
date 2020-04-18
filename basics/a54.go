// tuple assignment x = y = y, x or multiple variable assignment
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b ,c, d int = 10, 5, 25, 100
	fmt.Println("Initial value of a, b, c, d are", a, b, c, d)
	// default way
	a , b = c , d
	fmt.Println("After tuple assignment a, b = c,d\nvalue of a, b, c, d are", a, b, c, d)
	a = 10
	b = 5
	// right hand side can be a mathematical operation
	a, b = c + 1, d + 2
	fmt.Println("value of a, b are ", a, b)
	a , b = 10, 5
	a, b = int(math.Sqrt(float64(c))), int(math.Sqrt(float64(d)))
	fmt.Println("value of a, b are ", a, b)
	// another declaration
	x, y, z := a, b, c
	fmt.Println("values of x, y, z are:", x, y, z)
}
