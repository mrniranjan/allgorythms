/* Difference between Raw strings and interpreted 
string literals */
package main

import "fmt"

func main() {
	/*Raw String are enclosed using backticks. In raw
        strings \t doesn't have any special meaning so \t is 2 characters*/
	a := `Go\tProgramming\tLanguage`
	/* Below string is an Interpreted string \t is interpreted*/
	b := "Go\tProgramming\tLanguage"
	fmt.Println(a)
	fmt.Println(b)
}
