// Storing factors of number in map

package main

import "fmt"

func main() {
	numbers := []int{75, 24}
	// create a map of integer and array
	factors := make(map[int][]int)
	for index := range numbers {
		var temp []int 
		num := numbers[index]
		for i :=1; i <= num; i++ {
			if num%i == 0 {
				temp = append(temp, i)
			}
		}
		factors[num] = temp
	}
	for key, value := range factors {
		fmt.Println("number: ", key, "Factors: ", value)
	}
	
}
