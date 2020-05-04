/*Using Continue statement in loops. This programs takes arguments
passed through command line and Word Hello is prefix to those args*/

package main

import (
	"fmt"
	"os"
)

func main() {
	s := "Hello "
	for index, arg := range os.Args {
		if index < 1 {
			continue
		}
		s += fmt.Sprintf("%s ",  arg)
	}
	fmt.Println(s)
}
