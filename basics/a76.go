/* Go can have Unicode variables
Refer: https://golang.org/ref/spec#Identifiers
*/
package main

import (
	"fmt"
)

func main() {
	u03a0 := 3.14
	fmt.Printf("value of \u03a0 is %g\n", u03a0)
	π := 22 / 7.0
	fmt.Printf("value of π is %g\n", π)
}
