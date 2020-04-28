//Insertion Sort

package main

import (
	"fmt"
)
func main() {
	var j, a =0, []int{55, 11, 45, 22, 2, 16, 13, 43, 19, 21, 98}
	fmt.Println("Unsorted Array: ", a)
	for i :=0; i < len(a); i++ {
		key := a[i]
		for j = i-1; j >= 0 && a[j] > key; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = key
	}
	fmt.Println("Sorted Array: ", a)
}
