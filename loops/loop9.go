/* Finding factors of a number */
package main

import (
	"fmt"
)

func main() {
	num := 24
	fac := []int{}
	for i:=1; i <= num; i++ {
		if num % i == 0 {
			fac = append(fac, i)
		}
	}
	fmt.Printf("Factors of number %d are: %v\n", num, fac)
}
