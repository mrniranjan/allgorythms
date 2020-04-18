// using iota

package main


import "fmt"

const (
	one = 1
	two = 2
	three = iota + 1
	four = iota + 2
)
func main() {
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	fmt.Println(four)
}
