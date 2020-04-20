// basic for loop
package main

import "fmt"

func main() {
	s := "Go Programming Language"
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}
}
