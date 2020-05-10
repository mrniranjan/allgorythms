// checking the sign

package main

import (
	"fmt"
)

func main() {
	numbers := []int{2, 3, -2, -1, -6}
	check := -6
	for _, value := range numbers {
		fmt.Printf("\n%d and %d have same Sign? %t\n", value, check, (value ^ check) >= 0)
	}
}
