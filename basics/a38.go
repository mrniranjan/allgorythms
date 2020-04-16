// this program is differentiate in size when
// string is declared in single and double quotes

package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := 'a'
	s2 := "a"
	s3 := `a`

	fmt.Printf("Type of s1 is: %T\n", s1)
	fmt.Printf("Size of s1 variable is: %d\n", reflect.TypeOf(s1).Size())
	fmt.Printf("Type of s2 is: %T\n", s2)
	fmt.Printf("Size of s2 variable is: %d\n", reflect.TypeOf(s2).Size())
	fmt.Printf("Type of s3 is: %T\n", s3)
	fmt.Printf("Size of s3 variable is: %d\n", reflect.TypeOf(s3).Size())
}
