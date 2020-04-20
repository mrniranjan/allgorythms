// Appending to a slice
package main

import (
	"fmt"
)

func main() {
	a := []int{3, 5, 11}
	b := append(a[:1], 7)
	fmt.Printf("a =%v, b =%v\n", a, b)
	c := []int{2, 4, 6}
	d := append(c[:2], 8)
	fmt.Printf("c =%v, d =%v\n", c, d)
	//dig deeper
	e := []int{2, 4, 6}
	/*[:1] is basically [0:1] , ie i = 0, j= 1,  
        so creates a slice that refers from 0th element to j -1 which is 1 - 1 =0
       so will refer to 0 element of array , i.e slice e will have only 1 element 2 */
	f := append(e[:1], 10)
	/* as we say the slice [:1] will have only 1 element 2, 
        and to that we are appending 10 after 2 , so the new slice e will be appended 10
        but will replace the existing 4. But slice f will see only 2 elements 2, 10*/
	fmt.Printf("e =%v, f =%v\n", e, f)
}

