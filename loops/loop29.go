// Getting a chunk from an array

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for len(numbers) > 0 {
		n := 2
		if n > len(numbers) {
			n = len(numbers)
		}
		chunk := numbers[:n]
		numbers = numbers[n:]
		fmt.Println("Batch", chunk)
	}
}
