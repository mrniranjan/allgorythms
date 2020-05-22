//inplace slice
package main

import (
	"fmt"
)

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("non empty strings :%q\n", nonempty(data))
	fmt.Printf("Actual data : %q\n", data)
	// override
	data = nonempty(data)
	fmt.Printf("modified data : %q\n", data)
}

func nonempty(s1 []string) []string {
	i := 0
	for _, s := range s1 {
		if s != "" {
			s1[i] = s
			i++
		}
	}
	return s1[:i]
}
