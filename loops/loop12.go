//fibonacci series
package main

import (
	"fmt"
)

func main() {
	var n, a, b, counter int = 10, 0, 1, 0
	fmt.Printf("Fibonacci series for number %d are: \n", n)
	for  {
		if counter > n {
			break
		} else {
			a, b = b, a+b
			counter += 1
			fmt.Printf("%d ", a)
		}
	}
	fmt.Println()
}
