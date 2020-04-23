// Using loops to find factors of a Number
package main

import "fmt"

const num = 50

func main() {
	var i int = 1
	fmt.Printf("Factors of %d are: ", num)
	for ;i<num;i++ {
		if (num % i) == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
