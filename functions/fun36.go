//implementing stack using slice

package main

import (
	"fmt"
)

func main() {
	var n []int // empty slice
	fmt.Println("Empty stack ", n)
	n = push(n, 1)
	fmt.Printf("After pushing %d stack is %v\n", 1, n)
	n = push(n, 2)
	fmt.Printf("After pushing %d stack is %v\n", 2, n)
	n = pop(n)
	fmt.Printf("After pop stack is %v\n", n)
}
func push(x []int, v int) []int {
	x = append(x, v)
	return x
}

func pop(x []int) []int {
	//get top of the stack
	top := x[len(x) -1]
	fmt.Println("Top of the stack is :", top)
	x = x[:len(x)-1]
	return x
}



