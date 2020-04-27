// divisiblity by 4

package main

import "fmt"

func main() {
	number := 1936
	temp := number
	t1 := 0
	for temp > 0 {
		t1 := temp % 100
		temp = temp / 10
		if t1 > 10 && t1 < 100 {
			break
		}
	}
	if t1 % 4 == 0 {
		fmt.Printf("\n%d is divisible by 4\n", number)
	}
}
