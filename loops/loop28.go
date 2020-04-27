// Another method to traverse a loop
package main

import "fmt"

func main() {
	numbers :=[]int{2, 8, 5, 3, 15}
	for i:= range numbers {
		numbers[i]++
	}
	fmt.Println(numbers)
}
