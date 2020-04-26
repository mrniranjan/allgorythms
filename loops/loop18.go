// sum of digits of a number

package main

import (
	"fmt"
)

func main() {
	var n, sum int = 5213, 0
	temp := n
	for temp != 0  {
		sum += temp % 10
		temp = temp / 10
	}
	fmt.Printf("Sum of digits of number %d is %d\n", n, sum) 
}
