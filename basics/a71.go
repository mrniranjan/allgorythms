// basic for loop example to calculate sume of digits 
package main

import "fmt"

func main() {
	num := 2324
	var sum int
	for i :=0; num != 0; i++ {
		digit := num%10 // get last digit
		num = num/10
		sum += digit
	}
	s := "Sum of %d digits is %d\n"
	fmt.Printf(s, 2324, sum)
}
