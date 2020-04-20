/* comparing and assigning  nil 
The program will not compile
*/
package main

import (
	"fmt"
)

func main() {
	n := nil
	const T = nil != nil // nil cannot be compared 
	fmt.Println(T)
	fmt.Println(nil == nil) // nill cannot be compared
	fmt.Println(n)
}
