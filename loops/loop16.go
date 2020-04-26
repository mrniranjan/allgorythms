// check if a number exists in an array

package main

import (
	"fmt"
)

func main() {
	var myArray =  []int{2, 3, 8, 4, 1}
	fmt.Println(ifExists(4, myArray))
	fmt.Println(ifExists(5, myArray))
}
// Returns true if the number found in array
func ifExists(x int, y []int) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}

