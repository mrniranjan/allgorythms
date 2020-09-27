/* In this program we take user input of 2 numbers 
and also the operation to be performed on the numbers. 
We use swtich statements and call the functions which
actually do the operations on the numbers 
*/
package  main

import (
	"fmt"
	"os"
)


func main() {
	var operator string
	var n1, n2 int 
	operatorList := []string{"+","-","/","*"}
	
	fmt.Printf("\nEnter an operator(+,-,*,/): ")
	if _, err := fmt.Scan(&operator); err != nil {
		fmt.Println(err)
		os.Exit(1)
		
	}
	if verifyOperator(operator, operatorList) {
		fmt.Printf("\nEnter 2 numbers: ")
		if _, err2 := fmt.Scan(&n1, &n2); err2 != nil {
			fmt.Println(err2)
			os.Exit(1)
		}
		fmt.Println(numericalOperation(&n1, &n2, operator))
	} else {
		fmt.Printf("Did not enter right operator\n")
	}
}
// verify operators are valid 
func verifyOperator(op string, opList []string) bool {
	for _, v := range opList {
		if op == v {
			return true
		} else {
			continue
		}
	}
	return false
}
// Does numerical operation on the numbers
func numericalOperation(n1 *int, n2 *int, operator string) (result int) {
	result := 0
	switch operator {
	case "+" :
		result = *n1 + *n2
	case "-":
		result = *n1 - *n2
	case "*":
		result = *n1 * *n2
	case "/":
		result = *n1 / *n2
	default:
		result = *n1 + *n2
	}
	return 
}
