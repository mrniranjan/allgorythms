package main

import (
	"fmt"
)

func main() {
	s := "This is a string"
	fmt.Printf("%v, %T\n", s, s)

	/*
	Alot of the functions we use in Go actually work with byte slices.
	That makes it much more flexible than a hardcoded string.
	 */
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b ) // returns %T as []uint8

	backToString := string(b)
	fmt.Printf("%v, %T\n", backToString, backToString)
}
