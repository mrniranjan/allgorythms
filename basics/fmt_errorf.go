package main

import (
	"fmt"
)

func main() {
	const name, id = "bueller", 17
	fmt.Printf("user %q (id %d) not found\n", name, id)

	// Errorf formats according to a format specifier and returns the string as a value that satisfies error.
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())
}
