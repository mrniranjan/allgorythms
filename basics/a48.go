// understanding constants

package main
import "fmt"

// here constants are declared which have explicit
// values 
const (
	mon = 2
	tue = 3
	wed = 5
	thu = 6
	fri = 7
	sat = 8
)

const (
	hours = 10
)

func main() {
	a := mon * 15
	b := tue * hours
	c := sat * 8
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
