// Using os.Args
package main

import (
	"fmt"
	"os"
)

func main() {
	var str string
	// What ever is typed after ./a85.go will
	// be appended to string str
	for i:=1; i < len(os.Args); i++ {
		str += os.Args[i]
	}
	fmt.Println(str)
}
