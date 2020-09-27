/*The purpose of this program is to illustate within a
function lexical blocks may be nested so one declaration
may shadow another. 

The below program has 3 different variables called x 
because each declaration appears in different lexical block.
*/
package main

import (
	"fmt"
)

func main() {
	x := "hello!"
	for i:=0;i<len(x);i++ {
		x := x[i]
		if x != '!' {
			x:= x + 'A' - 'a'
			fmt.Printf("%c", x) //HELLO one letter per iteration
		}
	}
}
