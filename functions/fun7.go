/* Example of passing multiple integer pointers */
package main

import (
	"fmt"
)

func main() {
	num1, num2 := 32, 555
	ret := compare(&num1, &num2)
	if ret {
		fmt.Printf("\n%d is greater then %d\n", num1, num2)
	} else {
		fmt.Printf("\n%d is les than %d\n", num1, num2)
	}
}

func compare(p1 * int, p2 * int) bool {
	temp1, temp2, c1, c2 := *p1, *p2, 0, 0
	for temp1 != 0 {
		temp1 = temp1/10
		c1++
	}
	for temp2 != 0 {
		temp2 = temp2 / 10
		c2++
	}
	if c1 > c2 {
		return true
	}
	return false

}
