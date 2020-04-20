package main
import "fmt"

type Celsius float64
type Fahrenheit float64

func main() {
	c := Celsius((212 - 32) * 5 / 9)
	a := fmt.Sprintf("%g degrees", c)
	fmt.Printf("Type of a is : %T\n", a)
	fmt.Printf("Type of c is : %T\n", c)
	fmt.Println(a)
}
