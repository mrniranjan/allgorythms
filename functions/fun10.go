/* using flag package and accessing flag arguments
using pointers 
./fun10 a bc def
./fun10 -s / a bc def ga
a/bc/def/ga
./fun10 --help
*/

package main

import (
	"fmt"
	"flag"
	"strings"
)


// n and sep are pointers which must be accessed using pointers
var (
	n = flag.Bool("n", false, "omit training newline")
	sep = flag.String("s", " ", "separator")
)

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if ! *n {
		fmt.Println()
	}
}
