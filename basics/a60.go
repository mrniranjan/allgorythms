// introduction to Arrays
// more array examples
package main

import "fmt"

func main() {
	// not all is to be declared in which case the last element is zero
	a := [3]int{1, 2}
	for i, v := range a {
		fmt.Printf("index=%d, value=%d\n", i, v)
	}

	b := [...]int{1, 2, 3}
	fmt.Printf("b is of type : %T\n", b)

	// another method of array declaration
	// defines an array of 10 elements, all zeros , except last element
	// is -1
	r :=[...]int{9: -1}
	for i, v := range r {
		fmt.Printf("index=%d, value=%d\n", i, v)
	}
}
