/* similar to fun15.go, Here we add new nameserver 
This program doesn't update resolv.conf yet. just
creates a slice with updated nameservers. 
*/
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	defaultResolvConf = "resolv.conf"
)

func main() {
	contents, err := ioutil.ReadFile(defaultResolvConf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	newIPAddr := "192.168.122.123"
	newContents := make([]string, 0, 1)
	newDNS := fmt.Sprintf("nameserver %s", newIPAddr)
	newContents = append(newContents, newDNS)
	addToNameServers(contents, &newContents)
}

func addToNameServers(contents []byte, newContents *[]string) {
	lines := bytes.Split(contents, []byte("\n"))
	for _, curLine := range lines {
		cIndex := bytes.Index(curLine, []byte("#"))
		if cIndex != -1 {
			break
		}
		*newContents = append(*newContents, string(curLine))
	}
}
