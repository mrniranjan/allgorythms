// program using unsigned 8 and 16 bits . This program will give error
package main

import "fmt"

func main() {
	var a uint8 // unsigned 8 bits
	var b uint16 //unsigned 16 bits
	a = 256 -1
	fmt.Printf("var a has value %d and type %T\n", a, a)
	b = 9999/9
	fmt.Printf("var b has value %d and type %T\n", b, b)
	/*
	b = a + b // this is currently not possible
	fmt.Printf("var b has value %d and type %T\n", b, b)
        */
	b = uint16(a) + b
	fmt.Printf("var b has value %d and type %T\n", b, b)
}

