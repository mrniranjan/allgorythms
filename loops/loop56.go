/* Based on basics/list_subdirectory.go using for loop to 
to walk the directory given from os.Args
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	for _, filePaths := range os.Args[1:] {
		err := filepath.Walk(filePaths, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("Something wrong with accessing path %q, %v\n", filePaths, err)
			}
			fmt.Println("files: ", path)
			return nil
		})
		if err != nil {
			fmt.Println("Something bad happened")
		}
		
	}
}
