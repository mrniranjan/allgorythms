/*Passing maps to functions. Maps types are pointers by default s*/
package main

import (
	"fmt"
	"os"
)

func main() {
	var m map[string]int
	if m == nil {
		fmt.Println("nil map ", m)
	}
	/* make function allocates and initializes a hash map data 
        structure and returns a map value that points to it. so data
        here is a pointer to data structure. 
        */
	data := make(map[string]int) 
	fun(data) // when passing to function, data is already pointer so no need to use &data
}

func fun(m map[string]int) {
	for i :=1; i < len(os.Args); i++ {
		s := os.Args[i]
		m[s]++
	}
	fmt.Println(m)
}
