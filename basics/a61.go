// using constants and Arrays
package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)


func main() {
	symbol := [...]string{USD: "$", EUR: "\u20AC", GBP : "\u00A3", RMB: "\u00A5"}
	fmt.Println(RMB, symbol[RMB])
	fmt.Println(USD, symbol[USD])
	fmt.Println(GBP, symbol[GBP])
	fmt.Println(EUR, symbol[EUR])
}
