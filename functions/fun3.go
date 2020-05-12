/* Using arrays and passing pointer to array and incrementing array */
package main

import (
	"fmt"
)

func main() {
	myArray := []int{0,1,2,3,4}
	fmt.Printf("Original Array: %d\n", myArray)
	incrArray(&myArray)
	fmt.Printf("Incremented Array: %d\n", myArray)
}

func incrArray(p *[]int) {
	for i :=0; i < len(*p); i++ {
		//start incrementing
		(* p)[i]++
	}
}
