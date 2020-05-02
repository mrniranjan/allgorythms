/*Using Basename */
package main

import (
	"fmt"
	"path"
	"path/filepath"
)


func main() {
	cachePath := path.Base("/var/lib/sss/db")
	fmt.Println(cachePath)
	absPath, _ := filepath.Abs("db")
	fmt.Println(absPath)
}
