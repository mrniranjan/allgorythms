// Slice program
package main

import (
	"fmt"
)

func main() {
	myarr1 := make([]int, 12) // 12 zeros
	myarr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	slice1 := append(myarr1[:12],0)
	// creating slice similar to myarr2 
	for i, v := range myarr2 {
		fmt.Printf("%d %d\n", i, v)
		slice1 = append(myarr1[:i], v)
	}
	fmt.Println(slice1)
}
