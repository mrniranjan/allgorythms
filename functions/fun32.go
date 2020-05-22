/* Append function for slices */
package main

import (
	"fmt"
)

func main() {
	var mystr1 []string
	for _, v := range "All Go Rythms" {
		mystr1 = append(mystr1, string(v))
	}
	fmt.Printf("%q\n", mystr1)
	var mystr2 []string 
	for _, v := range "Go Lang" {
		mystr2 = appendStr(mystr2, string(v))
	}
	fmt.Printf("%q\n", mystr2)
	var x, y []string
	z := "Hello World"
	for i:=0; i<len(z); i++ {
		y = appendStr(x, string(z[i]))
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
/*Each call to appendStr must check whether slice has 
sufficient capacity to hold new elements. if yes
extend the slice by defining larger slice , copy
the string 's' in to tempStr and return slice

if there is insufficient space. allocate a big array
by doubling its slice. 
 */
//appendStr is equivalent to append function above.
func appendStr(str []string, s string) []string {
	var tempStr []string
	strlen := len(str) + 1
	if strlen <= cap(str) {
		// There is room to grow. Extend the slice
		tempStr = str[:strlen]
	} else {
		// there is insufficient space. Allocate a new array
		// Grow by doubling
		strcap := strlen
		if strcap < 2*len(str) {
			strcap = 2 * len(str)
		}
		tempStr = make([]string, strlen, strcap)
		copy(tempStr, str)
	}
	tempStr[len(str)] = s
	return tempStr
}
