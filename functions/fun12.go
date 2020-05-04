/* Using new to create pointers */
package main

import (
	"fmt"
)

func main() {
	number := 61809
	p := new(int) // p, of type *int points to an unnamed int variable
	fmt.Println("uninitialized pointer will contain ", *p) // will print 0
	*p = number
	fmt.Println("pointer now contains", *p)
	/* Here the most important thing is *p is not a pointer to number
        instead its pointer to another integer type containing value which 
        is equivalent to value of number
        */
	retValue := funPtr(p) //Note here we are not passing *p, instead we are passing integer
	fmt.Println("After executing function funPtr value of number is", number) 
	fmt.Println("Value returned by function funPtr is", retValue)
}

func funPtr(p * int) int {
	fmt.Println("Value in p is:", p)
	fmt.Println("value of *p is:", *p) 
	*p = 25000
	return *p
}
