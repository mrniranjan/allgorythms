package main

import "fmt"

func main() {
	num1 := 382
	num2 := 49
	num3 := 18
	num4 := 5
	num5 := 750
	s1 := `Can you instantly identify the greatest
and smallest numbers in the below list`
	fmt.Printf(s1)
	fmt.Printf("%d, %d, %d, %d %d\n", num1, num2, num3, num4, num5)
	fmt.Printf("%d is greater and %d is smallest\n", num4, num3)
}
