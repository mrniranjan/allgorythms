// UTF-8
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := "去" // chinese of Go
	fmt.Println(a)
	b := "\xe4\xb8"
	fmt.Println(b) // will print ?
	c := "\u4e16" // 16-bit value
	fmt.Println(c)
	d := "\U00004e16" // 32 bit value
	fmt.Println(d)
	fmt.Println(len(a))
	fmt.Println(utf8.RuneCountInString(a)) // will print 1
	a += " Programming Language"
	fmt.Println(utf8.RuneCountInString(a))
}
