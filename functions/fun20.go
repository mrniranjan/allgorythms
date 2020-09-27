//This program illustrates the use of defer
package main

import (
	"fmt"
)
// Defer cause the loop to be printed in reverse order
func d1() {
	for i:=3;i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}
// Same output as above, below is the better method to call defer
func d2() {
	for k :=3; k > 0; k-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(k)
	}
}
func main() {
	d1()
	fmt.println()
	d2()
	fmt.Println()
}
	
