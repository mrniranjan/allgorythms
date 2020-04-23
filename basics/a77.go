/* Constants and Type Declarations
*/
package main

import (
	"fmt"
	"time"
)

const fiveSeconds = 5 * time.Second

func main() {
	now := time.Now()
	FiveNanoSeconds1 := now.Add(-5) 
	FiveNanoSeconds2 := now.Add(-fiveSeconds)
	// if we declare integer -5
	myvar := -5
	//FiveNanoSeconds3 := now.Add(myvar) // will give compiler erro
	// as there is an implicit converstion of typed variable to type
	// duration

	fmt.Printf("Now :%v\n", now)
	fmt.Printf("FiveNanoSeconds1 :%v\n", FiveNanoSeconds1)
	fmt.Printf("FiveNanoSeconds2 :%v\n", FiveNanoSeconds2)
}
	
