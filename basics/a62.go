//Arrays: Comparing 2 arrays
package main

import "fmt"

func main() {
	a := [2]int{1, 2}
	b :=[...]int{1, 2}
	fmt.Println(a == b) // will be true 
	c := [2]int{1, 3}
	 // will be false because the values of array don't match 
	fmt.Println(b == c)
}
