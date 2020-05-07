package main

import (
	"fmt"
	"os"
)

func main() {
	f, err1 := os.Create("a.txt")
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}
	l, err2 := f.WriteString("Hello World")
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(l, "bytes written successfully")
	f.Close()
}
