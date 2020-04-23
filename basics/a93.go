// example of for loop where there is no initialization
// and increment
package main

import "fmt"

func main() {
	num := 1234
	temp := 0
	for ; num != 0 ; {
		num = num / 10
		temp++
	}
	fmt.Printf("Number of digits in number %d is %d\n", 1234, temp)
}
