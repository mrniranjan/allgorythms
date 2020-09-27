package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	contents, err := ioutil.ReadFile("/etc/shadow")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("File Contents: %s", contents)
	}
}
