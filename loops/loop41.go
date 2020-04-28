// Using bitwise operators

package main

import (
	"fmt"
)

func main() {
	power := make([]int, 5)
	for i := range power {
		power[i] = 1 << uint(i)  // 2 ** i
		}
	for _,  value := range power {
		fmt.Printf("%d\n", value)
	}
}
