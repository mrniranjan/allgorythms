/* Print Duplicate Lines: In standard output 
Example:
./loop52
Go Lang Programming Language
GoLang Programming Language
GoLang Programming Language
GoLang Programming Language 
ctrl-D
<output>
2       Go Lang Programming Language
2       GoLang Programming Language
map[Go Lang Programming Language:2 GoLang Programming Language:2]
</output>
*/
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	fmt.Println("Press CTRL-D To Exit")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	fmt.Println(counts)
}
