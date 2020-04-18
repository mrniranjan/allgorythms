//slices : Appending to slices 
package main

import "fmt"

func main() {
	s := make([]string, 5)
	fmt.Println("empty string:", s)
	s[0] = "H"
	s[1] = "E"
	s[2] = "L"
	s[3] = "L"
	s[4] = "O"
	fmt.Println(s)
	fmt.Printf("Length of string %s is %d\n",s, len(s))
	s = append(s, " ")
	s = append(s, "W", "O", "R", "L", "D")
	fmt.Println("New string is: ", s)
	a := s[:5]
	fmt.Println(a)
	b := s[2:5]
	fmt.Println(b)
}
