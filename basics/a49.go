// using iota

package main

import "fmt"

type Weekday int

const (
	sun Weekday = iota + 1 // we are adding iota to increment other constants by 1
	mon // 2
	tue // 3
	wed // 4
	thu // 5
	fri // 6
	sat // 7

)

const (
	hours = 10
)

func main() {

	a := mon * 15 // 2 * 15
	b := sun * hours  // 10 * 1
	c := sat * 8 // 7 * 8
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
