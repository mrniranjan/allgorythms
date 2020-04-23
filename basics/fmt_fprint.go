// Package fmt examples

package main

import (
	"fmt"
	"os"
)

func main() {
	/*
	Fprint formats using the default formats for its operands and writes to w.
	Spaces are added between operands when neither is a string.
	It returns the number of bytes written and any write error encountered.
	*/
	const name, age = "Kim", 22
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")

	// The n and err returns values from Fprint are
	// those returned by the underlying io.writer.
	if err != nil {
		fmt.Fprint(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")
}
