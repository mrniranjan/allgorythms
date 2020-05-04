/* 
Passing string to functions 

./fun9 Go Programming Language 
output: Go_Programming_Language
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(mystr(os.Args))
}

func mystr(args []string) string {
	return strings.Join(args[1:], "_")
}
