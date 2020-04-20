// Loops and if conditions
// basic if condition
package main

import "fmt"

func main() {
	var num1, num2 int
	num1 = 52
	num2 = 53
	if (num1 < num2) {
		fmt.Printf("%d is less than %d\n", num1, num2)
	}
	if (num2 < num1) {
		fmt.Printf("%d is less than %d\n", num2, num1)
	}
}

