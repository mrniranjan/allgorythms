// exchanging values

package main

import "fmt"

func main() {

	cores, threads := 8, 16
	fmt.Println("cores =", cores)
	fmt.Println("threads =", threads)
	cores, threads  = threads, cores
	fmt.Println("cores =", cores)
	fmt.Println("threads =", threads)
}
