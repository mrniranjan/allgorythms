//This code uses some examples of displaying strings
package main

import "fmt"

func main() {
	fmt.Println("/home" + "/foobar")
	delimiter := "/"
	fmt.Println(delimiter  + "home" + delimiter + "foobar")
	var a int
	a = 25
	fmt.Println("25 / 10 = ", a/10) // answer should be 2 
	fmt.Println("25 % 10 = ", a % 10) // answer should be 5
}
