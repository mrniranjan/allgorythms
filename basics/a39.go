// This program is to know the default values
// stored in a variable 
package main

import "fmt"

func main() {
	var i int
	var f float32
	var s string

	fmt.Println(i)
	fmt.Println(f) // expected 0.0 instead prints 0
	fmt.Println(s)
}
