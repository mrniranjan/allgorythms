/*This program illustrates the scope of variables
if statement below creates an implied block 
 */
package main

import (
	"fmt"
)

func f() {}
var g = "g"

func main() {
	f := "f"
	fmt.Println(f) // "f": local var f shadows package level func f
	fmt.Println(g) // "g": package level variable
	//fmt.Println(h) // compile error: undefined h
	if x := a(); x == 0 {
		fmt.Println(x)
	} else if y := b(x); x == y {    // here if is nested within first if statement
		fmt.Println(x, y)        //so variables declared in first if block are visible here
	} else {
		fmt.Println(x, y)
	}
	//fmt.Println(x, y) // here they are not visible
}

func a() int {
	return 0
}

func b(x int ) int {
	return x
}


