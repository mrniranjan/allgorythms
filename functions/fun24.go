/* This program is extension of fun23.go. Below is the right
method to use the package level variable */
package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func init() {
	fmt.Println("init function")
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}

func main() {}
