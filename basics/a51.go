package main
import "fmt"

// Define type Celsius and Fahrenheit of type float64 
type (
	Celsius float64
	Fahrenheit float64
)

func main() {
	//calculate Celsius to Fahrenheit
	var c Celsius
	var f Fahrenheit
	c = 25
	// converting c to type Fahrenheit
	f = Fahrenheit(c) * 9/5 + 32  // celsius to fahrenheit
	fmt.Printf("%g Degree Celsius is %g fahrenheit\n", c, f)
	// converting f to type Celsius
	c = Celsius((f - 32) * 5 / 9)
	fmt.Printf("%g Fahrenheit is %g celsius\n", f, c)
}
