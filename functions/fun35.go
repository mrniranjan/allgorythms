//inplace slice using append

package main

import (
	"fmt"
)

func main() {
	data := []string{"red", "white", "", "", "", "", "", "pink"}
	fmt.Println("Actual data ", data)
	fmt.Printf("non empty strings :%q\n", nonEmpty(data))
	data = nonEmpty(data)
	fmt.Println("modified data ", data)
}

func nonEmpty(s1 []string) []string {
	out := s1[:0]
	for _, s := range s1 {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
