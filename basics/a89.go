/* if else with Logical Operators */
package main

import (
	"fmt"
)

func main() {
	num := 15
	if num <= 50 {
		fmt.Printf("Number %d is less than or equal to 50\n", num)
	} else if num >= 51 && num <= 100 {
		fmt.Printf("Number %d is between 51 and 100\n", num)
	} else {
		fmt.Printf("Number %d is greater than 100\n", num)
	}
}
