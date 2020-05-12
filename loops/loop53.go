//2D Slices
package main

import (
	"fmt"
)

func main() {
	n := 5
	m := 2
	matrix := make([][]int, n)
	rows := make([]int, n * m)
	for i:=0; i < n; i++ {
		matrix[i] = rows[i*m : (i+1)*m]
	}
	fmt.Println(matrix)
	
}
