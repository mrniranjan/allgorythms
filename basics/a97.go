// Simple Functions
// Prints even numbers from an array
package main

import (
	"fmt"
)

func main() {
	num := []int{1,2,3,4,5,6,7,8,9}
	checkEven(num)
}
func checkEven(arr []int) {
	for i:=0; i < len(arr); i++ {
		if arr[i] % 2 == 0 {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()
}
