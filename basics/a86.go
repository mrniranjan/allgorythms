package main

import (
	"fmt"
	"os"
)

func main() {
	var sep, str string
	// Get the length of each string
	for i:=1; i < len(os.Args); i++ {
		fmt.Printf("Length of '%s' is %d\n", os.Args[i], len(os.Args[i]))
		str += sep + os.Args[i]
		sep = " "
	}
	/* ./a86 Hello World should print
	 length of string "Hello" and "World" and length of 
         combind string "Hello World"*/
	fmt.Printf("Length of string %s is %d\n", str, len(str))
}
