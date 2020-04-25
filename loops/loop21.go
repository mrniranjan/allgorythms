//finding a perfect number
package main

import "fmt"

func main() {
	var number, sum int = 6, 0
	for i :=1 ; i < number; i++ {
		if (number % i )  == 0 {
			sum += i
		}
	}
	if sum == number {
		fmt.Printf("%d is a perfect number\n", number)
	}
}
