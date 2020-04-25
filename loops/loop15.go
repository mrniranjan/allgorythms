//checking the digits of a number
package main

import (
	"fmt"
)

func main() {
	n := 532
	temp := n
	digits := make([]int, 0, 1) // holding number of digits
	for temp != 0 {
		digits = append(digits, temp % 10)
		temp = temp / 10
	}
	fmt.Printf("Digits of the number %d are %v\n", n, digits)
}
