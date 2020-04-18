// bit wise operation

package main

import "fmt"

func main() {
	var a uint8 = 1 & 0
	var b uint8 = 1 | 0
	var c uint8 = 4<<1 // 4 + 4
	fmt.Println("bit wise operation of 1 & 0 is : x", a)
	fmt.Println("bit wise operation of 1 | 0 is : y", b)
	fmt.Println(c)
	var d uint8 = 4<<2 // 4 * 4
	fmt.Println("bit wise operation of 4 << 2 is: ", d)
	var e uint8 = 1 
	fmt.Printf("%08b\n", e)
	var f uint8 = e | 1
	fmt.Printf("%08b\n", f)
	fmt.Printf("%08b\n", e & f)
	fmt.Printf("%08b\n", e ^ f)
}
