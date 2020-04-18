package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR")) // This should return emtpy
	// To see the value of BAR, you could also run as: BAR=2 go run os_env.go

	// This prints all the environ variables available
	fmt.Println(os.Environ())
}
