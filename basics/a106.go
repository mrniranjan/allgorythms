// Simple Input program

package main

import (
	"fmt"
)
func main() {
	var answer string
	fmt.Println("Is zero a whole number (Y/N):")
	_, err := fmt.Scan(&answer)
	if err != nil {
		fmt.Printf("Something went wrong : %s\n", err)
	}
	fmt.Println(answer)
}
