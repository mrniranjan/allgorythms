// Program exploring signed and unsigned integers
package main

import "fmt"

func main() {
	var a uint8
	a = 255
	fmt.Printf("var a has value %d and type %T\n", a, a)
	a = 256 // this overflows and will not compile 
	fmt.Printf("var a has value %d\n", a)
	a = 999/10
	fmt.Printf("var a has value %d\n", a)
	a = 9999/9  // this overflows and will not compile
	fmt.Printf("var a has value %d\n", a)
}
