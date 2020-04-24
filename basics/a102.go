//program to pring Args[0], the name of the command that invoked it
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("The Name of the command is: ", os.Args[0])
	fmt.Println("Arguments passed are: ", strings.Join(os.Args[1:], " "))
}
