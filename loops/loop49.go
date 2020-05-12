/* Matrix Addition */
package main

import (
	"fmt"
)

func main() {
	m1 := [2][2]int{ {1, 2}, {3, 4}	}
	m2 := [2][2]int{ {3, 1}, {4, 2} }
	m3 := [2][2]int{}
	fmt.Println(m1)
	fmt.Println(m2)
	for i:=0; i < 2; i++ {
		for j:=0; j < 2; j++ {
			m3[i][j] = m1[i][j] + m2[i][j]
			fmt.Printf("%d ", m3[i][j])
		}
		fmt.Println()
	}
}
