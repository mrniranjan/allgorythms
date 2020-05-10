//Greatest common Divisor
package main

import (
	"fmt"
)

func main() {
	var n1, n2 int =  75, 20
	gcd := GCD(n1, n2)
	fmt.Println(gcd)
}

func GCD(n1, n2 int) int {
	for n2 != 0 {
		t := n2
		n2 = n1 % n2
		n1 = t
	}
	return n1
}
