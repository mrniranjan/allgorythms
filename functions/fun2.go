/* Using array and pointer to array to function */

package main

import (
	"fmt"
)

func main() {
	numbers := []int{328, 4972, 18, 47281, 530}
	a := bigNum(&numbers)
	b := bigNum(&numbers)
	fmt.Printf("%d is bigger and %d is smaller among %d\n", a,b, numbers)
}

func bigNum(p *[]int) int  {
	curBig := (* p)[0]
	for _, v := range *p {
		if v > curBig {
			curBig = v
		}
	}
	return curBig
}

func curSmall(p * []int) int {
	curSmall := (* p)[0]
	for _, v := range *p {
		if v < curSmall  {
			curSmall = v 
		}
	}
	return curSmall
}
