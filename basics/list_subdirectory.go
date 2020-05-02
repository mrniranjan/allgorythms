/* This program usins filepath/walkfunc to list 
all the files in a directory */
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := "/tmp"
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		fmt.Println("files :", path)
		return nil
	})
	if err != nil {
		fmt.Println("Something bad happened")
	}
}
