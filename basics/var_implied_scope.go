/* This program illustrates the implicit block . In the below program
there are 3 variables named x :
1. function body (explicit)
2. one in for statement (implicit)
3. one in loop body (explicit)
*/
package main

import (
	"fmt"
)

func main() {
	x := "golang"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // GOLANG one letter per iteration
	}
}
