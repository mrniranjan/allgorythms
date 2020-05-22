// This is based on popCount https://godoc.org/github.com/tmthrgd/go-popcount#Count64
// Also based on  Donovan & Kernighan's book, The Go Programming Language.
package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	var x uint64 = 3
	fmt.Printf("Binary representation of %d is : %08b\n", x, x)
	fmt.Println("Popcount of x is: ", popCount(x))
}

func popCount(x uint64) int {
	s := 0
	for i :=uint(0); i < 8; i++ {
		s += int(pc[byte(x>>(i*8))])
	}
	return s
}
