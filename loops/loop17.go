// Is the number divisible by 7
package main

import (
	"fmt"
)

func main() {
	// This code is from loop15.go
	var count, n int =  0, 534
	temp := n
	digits := make([]int, 0, 1)
	divisible := make([]int, 0, 1)
	for temp != 0 {
		digits = append(digits, temp % 10)
		temp = temp / 10
	}
	for i:=2 ; i <= n; i++ {
		if i == 2 {
			if ifExists(digits[0], []int{0,2,4,6,8}) {
				divisible = append(divisible, 2)
				fmt.Println("Divisible by 2")
				count++
			}
		}
		if i == 7 {
			if divBy7(n) {
				divisible = append(divisible, 7)
				fmt.Println("Divisible by 7")
				count++
			}
		}
	}
}

// check if a number is divisible by 7 
func divBy7(n int) bool {
	lastDigit := n % 10
	remainingDigits := n / 10
	value := remainingDigits - (lastDigit * 2)
	if value % 7 == 0 {
		return true
	}
	return false
}

// Returns true if number found in array (from loop16.go)
func ifExists(x int, y []int) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}
