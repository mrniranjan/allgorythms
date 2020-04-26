// find the largest number in an array

package main

import (
	"fmt"
)

func main() {
	numbers := []int{15,32,21,41,52,19,10}
	curLargestNumber := numbers[0]
	for i := 0; i < len(numbers); i ++  {
		if  numbers[i] > curLargestNumber {
			curLargestNumber = numbers[i]
		}
	}
	fmt.Printf("%d is largest of number of %v\n", curLargestNumber, numbers)
}
