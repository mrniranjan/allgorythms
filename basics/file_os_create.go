package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	l, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	
}
