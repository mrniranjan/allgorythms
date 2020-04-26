// checking if a number is prime

package main

import "fmt"

func main() {
	var num , count int = 11, 0
	for i :=1 ; i <= num; i++ {
		if num % i == 0 {
			count ++
		}
	}
	if count == 2 {
		fmt.Printf("\n%d is a prime number\n", num)
	}
}

