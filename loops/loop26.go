// Prime factors

package main

import (
	"fmt"
)

func main() {
	number := 980
 	var factors, primeFactors = []int{}, []int{}
	for i :=1; i <= number; i++ {
		if number % i == 0 {
			factors = append(factors, i)
		}
	}
	for i :=0; i < len(factors); i++ {
		count := 0
		for j:=1; j <= factors[i]; j++ {
			if factors[i] % j == 0 {
				count++
			}
		}
		if count == 2 {
			primeFactors = append(primeFactors, factors[i])
		}
	}
	fmt.Printf("Factors of number %d are: %v\n", number, factors)
	fmt.Printf("Prime Factors of %d are: %v\n", number, primeFactors)
}
