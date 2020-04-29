/* Count the number of characters entered through stdin 
To run the program pass string or file through stdin as show
below 

echo "Hello World" | ./loop42
cat <file> | ./loop42
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	myStr := ""
	var nChar, nLines, nWords int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		myStr = scanner.Text()
		nChar += len(myStr)
		words := strings.Split(myStr, " ")
		nWords += len(words)
		nLines++
	}
	fmt.Printf("Number of characters entered are : %d\n", nChar)
	fmt.Printf("Number of Lines are : %d\n", nLines)
	fmt.Printf("Number of words are : %d\n", nWords)
}
