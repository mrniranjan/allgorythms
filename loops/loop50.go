/* Making 2D Slices */
package main

import (
	"fmt"
)

func main() {
	n := 2 // 2 rows
	m := 3 // 3 colums
	m1 := make([][]uint, n, m)
	for i :=0; i < n; i++ {
		m1[i] = make([]uint, m)
		for j :=0; j < m; j++ {
			m1[i][j] =  uint(i * j)
		}
	}
	fmt.Println(m1)
}
