/* Basic For Loop 
for initialization condition
post {
    // zero or more statements
}
initialization is optional  and executed
before the loop starts , it should be short
variable declaration like a := 10

condition is evaluated at beginning of each loop iteration
*/

/* Below program is traditional while loop */
package main

import "fmt"

func main() {
	var sum int = 1 
	fmt.Println("**** Traditional while Loop *****")
	for sum < 10 {
		sum += sum
		fmt.Println("After each iteration value of sum :", sum)
	}
	fmt.Println("**** Classic For Loop *****")
	for x:=1; x<10; x++ {
		fmt.Printf("%d * %d = %d\n", 2, x, 2 *x )
	}
}
