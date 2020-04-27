package main

import (
	"fmt"
)

func main() {
	myString := "a"
	myRune := 'a'

	omg := '\U0001F602'

	fmt.Printf("%v is of type: %T\n", myString, myString)
	fmt.Printf("%v is of type: %T\n", myRune, myRune)
	fmt.Printf("%c is of type: %T\n", omg, omg)
}
