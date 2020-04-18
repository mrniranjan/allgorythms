package main
import "fmt"

//global constants

const pi = 3.14
const radius = 45
const area = pi * radius * radius

func main() {
	fmt.Println(area)
	const radius = 3 // overrides the global constants
	const area = pi * radius * radius // overrides the global constants
	fmt.Println(area)
}
