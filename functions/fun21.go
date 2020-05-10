/*This program breaks the string in to smaller strings or tokens using
delimiters */
package main

import (
	"fmt"
)

func main() {
	src := "Hello, World!"
	delimiters := []string{",", "!"}
	for _, v := range delimiters {
		strtok(src, v)
	}
}
/*This function take string and delimiter and splits
the string based on delimiter*/
func strtok(s1 string, delimiter string) {
	toknum := 0
	for i, v := range s1 {
		if string(v) == delimiter {
			toknum = i
		}
	}
	fmt.Println(s1[:toknum])
}
