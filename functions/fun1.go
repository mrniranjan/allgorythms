/* Using arrays and passing array as pointer to function */

package main

import (
	"fmt"
)

func main() {
	num := []int{2, 6, 7, 9, 13, 15, 31}
	fmt.Println("Original array: ", num)
	checkEven(&num)
}

// Function to print even numbers 
func checkEven(p *[]int) {
	fmt.Println("Even numbers from above array:")
	for _, v := range *p {
		if v % 2 == 0 {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()
}
