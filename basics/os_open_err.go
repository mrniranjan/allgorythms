package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("unknown.go") // This file doesn't exist
	if err != nil {
		log.Fatal(err) // If the open fails, the error string will be self-explanatory, like

		// log.Fatal prints the following:
		// 2020/04/18 16:00:37 open unknown.go: no such file or directory
		// exit status 1
	}
}
