// finding all even divisors of a number with remainder 0

package main

import "fmt"

func main() {
	number := 424
	fmt.Printf("All the even divisors of %d are: ", number)
	for i:=1; i < number; i++ {
		if (number % i== 0) && (i % 2 == 0)  {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
