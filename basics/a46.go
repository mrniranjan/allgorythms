//understand untyped constants
package main

import "fmt"

func main() {
	var (
		cores = 8
		threads = 16
	)
	fmt.Println("before assigning untyped constant cores, threads are =", cores, threads)
	const untypedconstant = 8
	// assign a typed constant value to declared non constant variables. 
	cores = untypedconstant
	threads = untypedconstant
	fmt.Println(cores)
	fmt.Println(threads)
}
