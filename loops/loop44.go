// Convert lowercase string to upper case 
package main

import (
	"fmt"
	"strings"
)

func main() {
	source := "hello world"
	destination := ""
	for _, v := range source {
		if string(v) != " " {
			/* we get v as rune which is integer so
			integer subtraction and coverting back to string*/
			destination += string(v - 32) 
		} else {
			destination +=" "
		}
	}
	fmt.Println(destination)
	fmt.Println(strings.ToUpper(source))
}
