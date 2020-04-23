/*Using time package to check time taken to execute the program
Execute the program by providing large Arguments
Example: ./a103 Joy, beautiful spark of divinity, Daughter from Elysium, We enter, burning with fervour, heavenly being, your sanctuary!
*/


package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s , sep := "", ""
	fmt.Println("*** Printing os.Args using Loop ***")
	start1 := time.Now()
	for _, arg  := range os.Args[1:] {
		s +=sep + arg
		sep = " "
		fmt.Printf("%.7fs elapsed\n", time.Since(start1).Seconds())
	}
	fmt.Println(s)
	fmt.Println()
	fmt.Println("*** Printing os.Args using Join operation of strings ***")
	start2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.7fs elapsed\n", time.Since(start2).Seconds())
}
