/* Basic Switch */

package main

import "fmt"

func main() {
	var heads, tails int
	switch coinflip() {
	case "heads":
		heads++
		fmt.Println("Heads: ", heads)
	case "tails":
		tails++
		fmt.Println("tails: ", tails)
	default:
		fmt.Println("landed on edge!")
	}
}


func coinflip() string {
	return "heads"
}
