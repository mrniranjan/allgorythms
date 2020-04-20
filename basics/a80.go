// strings and unicode 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	czCity := "Břno"
	fmt.Println(len(czCity))
	/* The number of characters in the variable cz_city is 7
        Since in Go strings are UTF-8 encoded called runes, ř takes 2 bytes.
        To get the number of runes in a string use utf8.RuneCountInString() */
	fmt.Println(utf8.RuneCountInString(czCity))
}
