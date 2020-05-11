/*This program illustrates that package level variable cwd
is used to save the output of os.Getwd() . since neither
cwd nor err is already declared in in the init functions
block := declares both of them as local variables. Which 
causes the inner declaration the outer one inaccessible */
package main

import (
	"log"
	"os"
)

var cwd string

func init() {
	cwd, err := os.Getwd() // will result in error: unused 
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}
