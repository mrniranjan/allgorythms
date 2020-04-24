// Tagless Switch 

package main

import "fmt"

func main() {
	var x int
	x = 10
	fmt.Println(Signum(x))
	x = -5
	fmt.Println(Signum(x))
}
/* This form is called tagless switch also called switch true */
func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}



