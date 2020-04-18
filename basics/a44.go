//multiple constant declarations

package main

import "fmt"

func main() {
	const (
		Base_Frequency = 2.40
		Turbo_Frequency = 5.30
		Tdp = 45
	)
	fmt.Println(Base_Frequency)
	fmt.Println(Turbo_Frequency)
	fmt.Println(Tdp)
}
