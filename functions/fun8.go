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
	temp_num := *num
	for i :=1; i < *num; {
		fmt.Printf("\n%d place in %d is %d\n", i, *num, temp_num % 10)
		temp_num = temp_num / 10
		i = i * 10
	}
}
