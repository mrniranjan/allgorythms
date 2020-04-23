/* Iterating over a Map */

package main

import "fmt"

func main() {
	change := map[int]string {
		1 : "pennies",
		2 : "diemes",
		3 : "quarter",
	}
	for k, v := range change {
		fmt.Printf("k = %d v = %s\n", k, v)
	}
}
