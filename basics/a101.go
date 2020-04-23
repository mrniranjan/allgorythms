// Program to print the command line Arguments
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""
	/* In this case everytime the loop runs
        a new string is created when it encounters 
        +=, i.e if we give ./a101 Hello world 
        There will be a string ./a101 Hello and then 
         ./a101 Hello World, the earlier string ./a101 hello
        will be discarded or garbage collected 
	*/
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	/*Another method to do the above using strings package which
        is much more simpler and efficient */
	fmt.Println(strings.Join(os.Args[1:], " "))
}
