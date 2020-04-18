// Strings: accessing individual bytes of string
package main

import "fmt"

func main() {
	s := "Go Programming Language"
	fmt.Printf("Length of string '%s' is : %d\n", s, len(s))
	fmt.Println(s[0], s[7]) // 0th and 7th byte which is L, o
	// substring operation s[i:j] will yield a new string consisting
	// of bytes of original string starting at index i and continuing up to
	// but not including the byte at index j, i.e j-i bytes
	fmt.Println(s[0:6]) // Go Pro
	fmt.Println(s[:6]) // same as above
	fmt.Println(s[3:]) // Rest of the characters from space.
	fmt.Println(s[:]) //0 - 0 so the string as is
}
