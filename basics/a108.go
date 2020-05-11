//Another example of switch 
package main

import (
	"fmt"
)

func main() {
	mychar, ans := "a", "vowel"
	switch mychar {
	case "a":
	case "e":
	case "i":
	case "o":
	case "u":
	default:
		ans = "Consonant"
	}
	fmt.Println(ans)
}
