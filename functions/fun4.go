package main

import (
	"fmt"
)

func main() {
	var num int = 53121
	fmt.Printf("%d has %d digits\n", numberOfDigits(&num), num)
}

func numberOfDigits(p * int) int {
	count := 0
	temp := *p
	for temp != 0 {
		digit := temp % 10
		count++
		temp = temp / 10
	}
	return count
}
