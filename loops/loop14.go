// Generate a range of numbers and find factors of those numbers
package main

import (
	"fmt"
)

func main() {
	var startNum, endNum int =  5, 20
	// Initialy our slice will have No elements
	myRange := make([]int, 0, 1)
	/* we populate and increase the length and capability
	of the slice */
	for startNum <= endNum {
		myRange = append(myRange, startNum)
		startNum = startNum + 1
	}
	for _,v := range myRange {
		myFac := []int{}
		for i:=1; i <= v; i++ {
			if v % i == 0 {
				myFac = append(myFac, i)
			}
		}
		if len(myFac) == 2 {
			fmt.Printf("%d is Prime Number\n", v)
		}
		fmt.Printf("Factors of %d are: %v\n",v, myFac)
	}
}
