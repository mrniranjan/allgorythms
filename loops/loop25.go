//divisibility by 3

package main

import (
	"fmt"
)

func main() {
	var n int = 2532
	digits := getDigits(n)
	fmt.Printf("Digits of number %d are: %v\n", n, digits)
	sum := sumOfDigits(digits)
	fmt.Printf("Sum of digits of %d is %d\n", n, sum)
	if sum % 3 == 0 {
		fmt.Printf("%d is divisible by 3\n", n)
	} else {
		fmt.Printf("%d is not divisible by 3\n", n)
	}
}

func getDigits(n int) (digits []int) {
	digits = make([]int, 0, 1)
	for i :=0 ; n != 0; i++ {
		digits = append(digits, n % 10)
		n = n / 10
	}
	return
}

func sumOfDigits(arr []int) (sum int) {
	sum = 0
	for _, v := range arr {
		sum += v
	}
	return
}
