// Checking if a number is odd or even using bitwise operator

package main

import (
	"fmt"
)

func main() {
	numbers := []int{3,1,5,2,4,6,8,15,23,55}
	for _, value := range numbers {
		if value & 1 == 1 {
			fmt.Printf("%d is odd\n", value)
		} else {
			fmt.Printf("%d is even\n", value)
		}
	}
}
