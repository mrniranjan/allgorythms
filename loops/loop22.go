//finding a perfect number from an array of numbers
package main

import "fmt"

func main() {
	for number :=2; number < 100; number++ {
		sum := 0
		for i :=1; i < number; i++ {
			if (number % i) == 0 {
				sum += i
			}
		}
		if sum == number {
			fmt.Printf("%d ", number)
		}
	}
	fmt.Println()
}
