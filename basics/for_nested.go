package main

import (
	"fmt"
)

func main() {
	n := 5
	//outer for loop iterates i from 0 to 4
	for i := 0; i < n; i++ {
		//inner for loop iterates j from 0 to the current value of i.
		for j := 0; j <= i; j++ {
			//inner loop prints j for each iteration
			fmt.Print(j)
		}
		//outer loop prints a new line at the end of each iteration
		fmt.Println()
	}
}
