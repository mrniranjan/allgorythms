// Simple Input scanner. Reads input from stdin
// Press Ctrl-D to exit 
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("After entering press ctrl-d to exit")
	str := "Enter Name: "
	name := ""
	fmt.Println(str)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "" {
			break
		} else {
			name = input.Text()
		}
	}
	fmt.Println(name)
}
