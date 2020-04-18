//illustrate floating point numbers
package main

import "fmt"

func main() {
	mass := 15 // kg's
	height := 2 // 2 meters
	radius := 0.35
	pi := 3.14
	/*
       equation: mass/pi * radius * height
        */
	density := float64(mass) / pi * radius * float64(height)
	fmt.Printf("Density of cylinder :%.2f\n", density)
}
