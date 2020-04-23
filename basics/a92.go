// Another program on slices

package main

import (
	"fmt"
)
func main() {
	myarr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	if len(myarr1) > 5 {
		slice1 := myarr1[len(myarr1)-5:] // will print last 5 elements
		fmt.Println(slice1)
	}
	// we can't print slice here as it's outside of the scope
	//fmt.Println(slice1) // will throw undefined
}
