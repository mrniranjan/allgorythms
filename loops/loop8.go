/* iterating through loop using loop7 concept */

package main

import "fmt"

func main() {
	var myStr1 = []string{"sssd", "example.test", "pam", "autofs"}
	var myStr2 = []string{}
	/* Method1 of loop7 program */
	for i:=0; i < len(myStr1); i++ {
		myStr2 = append(myStr2, myStr1[i])
	}
	fmt.Println(myStr2)
	/* Comparing slices using method2 of loop7
        because we can't do myStr1 == myStr2 */
	for i, v := range myStr2 {
		fmt.Println(v == myStr1[i])
	}
}
