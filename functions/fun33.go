// similar to fun32.go, Here we use appendInt to append
// to integer array
package main

import (
	"fmt"
)

func main() {
	var x, y []int
	for i:=0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var p []int
	p = appendInt2(p, 1)
	fmt.Println(p)
	p = appendInt2(p, 2, 3)
	fmt.Println(p)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		//There is room to grow, Exten the slice
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array
		// growing by doubling
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
//makeing the appendInt2 variadic 
func appendInt2(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int , zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
