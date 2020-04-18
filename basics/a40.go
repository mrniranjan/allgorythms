// defining a bunch of a different type of variables

package main

import "fmt"

func main() {
	// case:1
	// we are declaring multiple variables with
	// explicitly mentioning the type
	var (
		i int
		f float64
		s string
	)
	//case:2 
	// Here we are declaring multiple variables
	// where the value of the variable defines the
	// type of variable. 
	var (
		processor_family = "Comet_lake"
		cores = 8
		base_frequency = 2.4
	)
	// in case:1 we are not specifying any value
	// so we are printing to see what are the default
	//value assigned
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(processor_family)
	fmt.Printf("Type of varilable processor_family = %T\n", processor_family)
	fmt.Println(cores)
	fmt.Printf("Type of varilable cores = %T\n", cores)
	fmt.Println(base_frequency)
	fmt.Printf("Type of base_frequency = %T\n", base_frequency)
}
