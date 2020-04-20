// multiple else if
package main

import "fmt"

func main() {
	num := 40
	if num % 3 == 0 {
		fmt.Printf("%d is divided by 3\n", num)
	} else if num % 4 == 0 {
		fmt.Printf("%d is divided by 4\n", num)
	} else if num % 8 == 0 {
		fmt.Printf("%d is divided by 8\n", num)
	}
}
