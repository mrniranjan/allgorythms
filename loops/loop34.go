// looping through strings

package main

import "fmt"

func main() {
	myStr := "Go Lang"
	for i, r := range myStr {
		fmt.Println(i, r, string(r))
	}
}
