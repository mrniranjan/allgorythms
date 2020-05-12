/* 2-D Array */

package main

import (
	"fmt"
)

func main() {
	twodarray := [2][2]int{}
	for i :=0; i < 2; i++ {
		for j:=0; j<2; j++ {
			twodarray[i][j] = i + 1
			fmt.Printf("%d ", twodarray[i][j])
		}
		fmt.Println()
	}
}
