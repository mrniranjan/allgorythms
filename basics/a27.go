// Understanding basic variable data types

package main

import "fmt"

func main() {
	var a = "/home"
	fmt.Printf("Data type of variable a is %T\n", a)

	var aa = 'a'
	fmt.Printf("Data type of variable 'aa' is %T\n", aa)

	var b, c = "/home", "/foobar"
	fmt.Print(b, c)
	fmt.Println() // to print new line

	var d = 2 // this will be an integer when not declared
	fmt.Printf("Data type of variable variable d is %T\n", d)

	var e int
	e = 25
	fmt.Println(e)

	f := "foobar" // another form of string declaration
	fmt.Println(f)

	s1 := ""
	fmt.Println(s1)

	var s2 string
	s2 = "Golang"
	fmt.Println(s2)

	var s3 string = "Go Programming Language"
	fmt.Println(s3)
}
