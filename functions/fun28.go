package main

import (
	"fmt"
)

func main() {
	var number, bits uint8 = 1, 1
	shiftLeft(number, bits)
	shiftRight(number, bits)
}
func shiftLeft(num uint8, bits uint8) {
	var a uint8 = 1
	for i :=1;i < 10;i++ {
		fmt.Printf("Before shift %d =  %08b\n", num, num)
		// Left shift 
		fmt.Printf("After Left shift by %d %08b which is %d\n", bits, num << bits, num << bits)
		a = a + 1
	}
}

func shiftRight(num uint8, bits uint8) {
	for j :=1; j < 10; j++ {
		fmt.Printf("Before shift %d = %08b\n", num, num)
		// Right shift  1 * bits 
		fmt.Printf("After right shift by %d %08b which is %d\n", bits, byte(num >> (1*bits)), byte(num >> (1*bits)))
		num = num + 1
	}
}
