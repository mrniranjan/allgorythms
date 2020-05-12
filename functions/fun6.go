/* Passing string pointers to functions */
package main

import (
	"fmt"
)

func main() {
	var s1, s2 = "Hello World", "Hello"
	compare(&s1, &s2)
}

func compare (s1 * string, s2 * string) {
	fmt.Printf("\nLength of %s is %d\n", *s1, len(*s1))
	fmt.Printf("\nLength of %s is %d\n", *s2, len(*s2))
}

