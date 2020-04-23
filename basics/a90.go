/* Another way to initialize maps with some data*/
package main

import (
	"fmt"
)

func main() {
	primeNumbers := map[int]int {
		1: 2,
		2: 3,
		3: 5,
		4: 7,
		5: 11,
	}
	for i:=1; i < len(primeNumbers); i++ {
		prime_numbers[i]++
	}
	// using range to print key, value pairs
	for key, value := range primeNumbers {
		fmt.Printf("\nkey: %d value: %d", key, value)
	}
	fmt.Println()
}
