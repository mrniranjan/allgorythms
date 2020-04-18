package main
import "fmt"

// Define type Celsius and Fahrenheit of type float64
type (
	Celsius float64
	Fahrenheit float64
)

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)
func main() {
	var c Celsius
	var f Fahrenheit
	// Doing arithmetic operation
	fmt.Printf("%g\n", BoilingC - FreezingC)
	// Adding comparision operators
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	//fmt.Println(c == f)  this compiler err
	// due to type mismatch
	fmt.Println(c == Celsius(f))
}
