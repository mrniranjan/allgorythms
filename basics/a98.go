//function 
// Simple functions, Returning an integer
package main

import "fmt"

func main() {
	numbers := []int{4,8,9,15,23,12}
	a := GreatNum(numbers)
	fmt.Printf("%d is bigger amount %d\n", a, numbers)
}
// returns greatest number from the array given
func GreatNum(arr []int) int {
	curGreat := arr[0] // store the first number as current greatest
	for i:=0;i<len(arr);i++ {
		if arr[i] > curGreat {
			curGreat = arr[i]
		}
	}
	return curGreat
}
