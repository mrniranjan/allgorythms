// program using float32 and float64
package main

import "fmt"

func main() {
	var f1 float32
	f1 = 12.2352523131 // precision up to 6 decimal points
	fmt.Printf("value of f1 is: %f\n", f1)
	fmt.Printf("value of f1 when rounded to 2 decimal points is : %.2f\n", f1)
	var f2 float64
	f2 = 12.3433543
	fmt.Printf("value of f2 when rounded to 2 decimal points is: %.2f\n", f2)
	var f3 int
	f3 = 57
	fmt.Printf("Value 57 can be converted to float using casting: %.2f\n", float32(f3))
}
