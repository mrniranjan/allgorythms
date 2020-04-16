// program using uint8

package main

import "fmt"

func main() {
	var a1, a2, a3, a4 uint8
	a1 = 99
	a2 = 't'
	fmt.Printf("variable a1 has value %d and type %T\n", a1, a1)
	// the ascii value of t which is 116 is printed 
	fmt.Printf("variable a2 has value %d and type %T\n", a2, a2)
	a3 = 's'
	fmt.Printf("variable a3 has value %d and type %T\n", a3, a3)
	a4 = a2 + a3 // here ascii value of t is added with ascii value of s
	fmt.Printf("variable a4 has value %d and type %T\n", a4, a4)
	fmt.Println(string(a2) + string(a3) + string(a4))
}
