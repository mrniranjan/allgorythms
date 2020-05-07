/* Similar to fun14.go , we read the contents of /etc/resolv.conf
and prints comments in the file
*/

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	defaultResolvConf = "/etc/resolv.conf"
)

func main() {
	contents, err := ioutil.ReadFile(defaultResolvConf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	printComments(contents)
}

//function which prints line which has comments
func printComments(contents []byte) {
	lines := bytes.Split(contents, []byte("\n"))
	for _, curLine := range lines {
		cIndex := bytes.Index(curLine, []byte("#"))
		if cIndex != -1 {
			fmt.Println(string(curLine))
		}
	}
}

