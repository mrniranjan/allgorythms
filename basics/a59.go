// introduction to Arrays
// This programs looks in to various array declarations
package main

import "fmt"

func main() {
	var a [3]int // initialized with array of 3 integers (0)
	fmt.Println(a[0]) // will print 0
	// print indices and elements
	for i, v := range a {
		fmt.Printf("index=%d, value=%d\n", i, v)
	}
	fmt.Println()
	fmt.Println("Strings")
	var b [3]string // an array of 3 string characters
	for i, v := range b {
		fmt.Printf("index=%d, value=%s\n", i, v)
	}
	b[0] = "a"
	b[1] = "b"
	b[2] = "c"
	for i, v := range b {
		fmt.Printf("index=%d, value=%s\n", i, v)
	}
	fmt.Println()
	fmt.Println("Integers")
	var c = [3]int{1, 2, 3} // another form of array declaration
	for i, v := range c {
		fmt.Printf("index=%d, value=%d\n", i, v)
	}
	fmt.Println()
	fmt.Println("Float")
	d := [3]float64{3.0, 5.5, 7.9} // another for array declaration
	for i, v := range d {
		fmt.Printf("index=%d, value=%g\n", i, v)
	}
}
