// find the sum values in a map

package main

import (
	"fmt"
)

func main() {
	memGrowth := map[string]int {
		"sssd-2.2.3-18" : 5123,
		"sssd-2.3.3-19" : 6823,
	}
	sum := 0
	for _, v := range memGrowth {
		sum += v
	}
	fmt.Println(sum)
}
