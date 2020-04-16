//another method to declare variables of different type in single statement

package main

import "fmt"

func main() {
	var family, cores, threads, frequency = "Comet lakey", 8, 16, 2.40
	fmt.Println(family)
	fmt.Printf("Type of varilable family = %T\n", family)
	fmt.Println(cores)
	fmt.Printf("Type of varilable cores = %T\n", cores)
	fmt.Println(threads)
	fmt.Printf("Type of varilable threads = %T\n", threads)
	fmt.Println(frequency)
	fmt.Printf("Type of varilable frequency = %T\n", frequency)
}
