// common factors 
package main

import "fmt"

func main() {
	var n1, n2 = 18, 24
	var fac1, fac2, commonFactors = []int{}, []int{}, []int{}
	for i :=1;i <= n1; i++ {
		if n1 % i == 0 {
			fac1 = append(fac1, i)
		}
	}
	fmt.Printf("\nFactors of number %d are: %v", n1, fac1)

	for  j :=1; j <= n2; j++ {
		if n2 % j == 0 {
			fac2 = append(fac2, j)
		}
	}
	fmt.Printf("\nFactors of number %d are: %v", n2, fac2)
	if len(fac1) > len(fac2) {
		commonFactors = findCommonFactor(fac1, fac2)
	} else {
		commonFactors = findCommonFactor(fac2, fac1)
	}
	fmt.Printf("\nCommon factors of %d and %d are: %v\n", n1, n2, commonFactors)
}

func findCommonFactor(fac1 []int, fac2 []int) (cf []int) {
	for i:=0; i < len(fac1) ; i++ {
		for j :=0 ; j < len(fac2); j++ {
			if fac1[i] == fac2[j] {
				cf = append(cf, fac1[i])
			}
		}
	}
	return			
}
