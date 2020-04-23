package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter text: \n")
	var input string
	// Scan scans text read from standard input, storing successive space-separated values into successive arguments.
	// Newlines count as space.
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("input string is: %v\n", input)
}
