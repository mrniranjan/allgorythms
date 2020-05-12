/* Passing integer pointer to function */
package main

import (
	"fmt"
)

func main() {
	num := 53323
	digit(&num)
}

func digit(num *int) {
	tempNum := *num
	for i :=1; i < *num; {
		fmt.Printf("\n%d place in %d is %d\n", i, *num, tempNum % 10)
		tempNum = temp_num / 10
		i = i * 10
	}
}
