//Basics of arithmetic operators
package main


import (
	"fmt"
	"unsafe"
)

func main() {
	a := 25
	fmt.Printf("Variable a having value %d is of type %T and size of a is %d\n", a, a, unsafe.Sizeof(a)) // a is int
	b := 25.5
	fmt.Printf("Variable b having value %g is of type %T and size of b is %d\n", b, b, unsafe.Sizeof(b))
	c := int32(25)
	fmt.Printf("Variable c having value %d is of type %T and size of c is %d\n", c, c, unsafe.Sizeof(c))
	a = a / 5.0 // a is of type int, int / float will be be int here 
	fmt.Printf("a is of type %T\n", a)
	//a = a % 5.5 // this will fail because unlike "/" operator which can work on int and float 
	//fmt.Printf("a is of type %T\n", a) // % operator works only on integers.
	a = 13
	a = a % -5 // also the behaviour of % on negative number is is always same as dividend 
	fmt.Println("value of a % -5 is: ", a)
}
