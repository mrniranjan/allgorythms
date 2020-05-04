//Bubble Sort

package main

import (
	"fmt"
)

func main() {
	var t, a =  0, []int{25,33,12,9,7,5,15}

	for i :=0; i < len(a); i++ {
		for j :=0; j < len(a) -1; j++ {
			if a[j] > a[j+1] {
				t = a[j]
				a[j] = a[j+1]
				a[j+1] = t
			}
		}
	}
	fmt.Println(a)
}
