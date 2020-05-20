// Passing arrays by value and by reference

package main

import (
	"fmt"
)

func main() {
	a := [3]int{1,2,3}
	modArray(a)
	fmt.Println("After passing array to function", a)
	modArrayPtr(&a)
	fmt.Println("After passing array by reference, value of a is:", a)
}

func modArray(a [3]int) {
	a[2] = 100
	fmt.Println("Value of array a when passed by value: ", a)
}

func modArrayPtr(ptr *[3]int) {
	ptr[2] = 999
	fmt.Println("value array when passed by reference: ", ptr)
}
