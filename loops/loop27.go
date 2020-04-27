// product of 2 numbers is whole number
package main

import "fmt"

func main() {
	arr1 := []int{3,4,5,6}
	arr2 := []int{3,6,7,8}
	for i:=1; i < len(arr1); i++ {
		for j :=1; j < len(arr2); j++ {
			if i * j  >= 0 {
				fmt.Printf("\n%d * %d = %d a whole number\n", i, j, i * j)
			}
		}
	}
}
