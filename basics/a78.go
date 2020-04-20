// loop examples
package main

import (
	"fmt"
)

const myNum = 123456789

func main() {
	startNum  := 18
	for i:=9; startNum <= 81; startNum += i {
		ret := myNum * startNum
		s := fmt.Sprintf(" %d * %d = %d", myNum, startNum, ret)
		fmt.Println(s)
	}
	/* the scope of variable i only in the for loop */
	fmt.Println("value of i is: ", i) // will give undefined error. 
}
