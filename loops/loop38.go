//Selection Sort

package main

import (
	"fmt"
)

func main() {
	var t, a = 0,  []int{98, 55, 11, 45, 22, 2, 16, 13, 43, 19, 21}
	fmt.Printf("\nUnsorted Array: %v", a)
	for i :=0; i < len(a); i++ {
		for j:=i+1; j <= len(a)-1; j++ {
			if a[i] > a[j] {
				t = a[i]
				a[i] = a[j]
				a[j] = t
			}
		}
	}
	fmt.Printf("\nSorted Array is : %v\n", a)
}
