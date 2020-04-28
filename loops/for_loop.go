package main

import (
	"fmt"
)

// for is Go’s only looping construct.
func main() {
	sum := 0
	for i := 0; i < 4; i++ {
		sum += 1
	}
	fmt.Println(sum) // Prints 4

	for j := 0; j < 8; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j) // Prints 1 3 5 7
	}

	// for without a condition will loop repeatedly until you break out of the loop
	// or return from the enclosing function.
	for k := 0; k < 10; k++ {
		if k == 4 {
			break
		}
		fmt.Println(k) // Print 0 1 2 3
	}
}
