// more example of strings
package main

import "fmt"

func main() {
	s := "Programming Language"
	fmt.Println("Go " + s[:])
	a := "Go"
	b := a
	a += " Programming Language"
	fmt.Println(b)
	fmt.Println(a)
}
