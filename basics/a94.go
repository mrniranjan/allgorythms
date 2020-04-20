// Use range on string to loop through the charactres

package main

import "fmt"

func main() {
	s := "Go Programing Language"
	// Print the character and ascii values
	for i ,r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	n := 0
	for _, _= range s {
		n++
	}
	fmt.Println("Total Number of characters",n)
	n = 0
	for range s {
		n++
	}
	fmt.Println("n = ", n)
}
