// divisibility by 2 and 6

package main

import "fmt"

func main() {
	numbers := []int{2,3,5,9,18,20,25,27,29,35,46,}
	for i := range numbers {
		if numbers[i] % 2 == 0 {
			if numbers[i] % 3 == 0 {
				fmt.Printf("%d is divisible by 2 and 6\n", numbers[i])
			}
		}
	}
	// another method

	for i := range numbers {
		if numbers[i] %2 == 0 && numbers[i] % 3 == 0 {
			fmt.Printf("%d is divisible by 2 and 6\n", numbers[i])
		}
	}
		
}
