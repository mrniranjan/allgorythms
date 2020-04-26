package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loop 1 breaks the inner loop")
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break
			}
		}
	}
	fmt.Println()

	// This is where labels come to our rescue. A label can be used to break from an outer loop.
	fmt.Println("Loop 2 breaks the outer loop")
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}
	}
}
