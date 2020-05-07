/*Function to read /etc/resolv.conf */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)
const (
	defaultResolvConf = "/etc/resolv.conf"
)

func main() {
	contents := getCurrentContents()
	fmt.Println(contents)
}

func getCurrentContents() (string) {
	contents, err := ioutil.ReadFile(defaultResolvConf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} 
	return string(contents)
}
