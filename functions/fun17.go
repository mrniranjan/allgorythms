// Reverse of an array

package main

import (
	"fmt"
)

func main() {
	Numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12}
	fmt.Println(Numbers)
	reverse(Numbers)
	reverse(Numbers[:])
	reverse(Numbers[:2])
	reverse(Numbers[2:4])
	reverse(Numbers[5:])
}

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}
