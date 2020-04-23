// Creating Empty Maps
package main

import (
	"fmt"
)

func main() {
	var m map[string]int
	fmt.Println(m["errors"]) // contains value 0
	fmt.Println(len(m)) // will return 0
	/* To determine if the value we got from map is the
value specified or the Nil value is to use comma, ok */
	val, ok := m["errors"] // ok will be true if key exists
	fmt.Println("value stored in errors keys is", val)
	fmt.Println(ok)
	// m["errors"] = 9  will result in Panic
}
