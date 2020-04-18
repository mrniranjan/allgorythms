// program using signed number int8
package main

import "fmt"

func main() {
	var a1 int8
	a1 = -5/3
	fmt.Printf("value of a1 is %d and type is %T\n", a1, a1)
	a2 := -5 // this is integer which is signed integer 
	fmt.Printf("value of a2 is %d and type of a2 is: %T\n", a2, a2)
	a3 := 4294967295 * 429496799 // largest an integer can hold
	fmt.Printf("value of a2 is %d and type of a2 is: %T\n", a3, a3)
}
