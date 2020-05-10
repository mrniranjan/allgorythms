// Reverse of a string 
package main

import (
	"fmt"
)

func main() {
	s1 := "Hello World"
	fmt.Printf("Reverse of string %s is %s\n", s1, myReverse(s1))
}

func myReverse(s string) (rString string) {
	for i :=len(s)-1; i >= 0 ; i-- {
		rString += string(s[i])
	}
	return 
}
