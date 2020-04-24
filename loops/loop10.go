/* Product of factors */
package main

import (
	"fmt"
)

func main() {
	num := 24
	fac := []int{}
	for i :=1; i < num; i++ {
		if num % i == 0 {
			fac = append(fac, i)
		}
	}
	for i :=1; i < len(fac); i++ {
		for j:=1 ; j<len(fac[1:]); j++ {
			if (i * j) == num {
				fmt.Printf("%d * %d\n", i, j)
			}
		}
	}
}
