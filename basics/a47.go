//understanding untype constants
package main

import "fmt"

func main() {
	var (
		h int
		f float64
		s string
	)
	// Here constants are used to give a ddefault value
	// example of TypedConstant. were the variables
	// declared above have typed constants. 
	const (
		defaultValue = 3
		defaultString = "Foobar"
	)
	h = defaultValue
	f = defaultValue
	s = defaultString
	fmt.Println(h)
	fmt.Println(f)
	fmt.Println(s)
}
