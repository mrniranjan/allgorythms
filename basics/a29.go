// alias declarations

package main

import "fmt"

func main() {
	type A = string // defining a new data type A 
	var a A
	a = "foo"
	a += "bar"
	fmt.Println(a)
	fmt.Printf("Type of a is %T\n", a) // this prints string
	type B string
	var b B
	b = "foo"
	b += "baz"
	fmt.Println(b)
	fmt.Printf("Type of b is %T\n", b) // this should print B
}
