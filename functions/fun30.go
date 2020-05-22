// This program uses popCount in fun29 and calcuates
// the difference in number of bits that are different
// in 2 sha256 hashes 
package main


import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n%T\n", c1,c2, c1==c2, c1, c2)
	count := 0
	for i:=uint8(0); i < 32; i ++ {
		if popCount(c1[i]) != popCount(c2[i]) {
			count++
		}
	}
	fmt.Println("The number of bits that are different in 2 hashes are", count)
}

func popCount(x uint8) int {
	s := 0
	for i:=uint(0); i < 8; i++ {
		s += int(pc[byte(x>>(i*8))])
	}
	return s
}
