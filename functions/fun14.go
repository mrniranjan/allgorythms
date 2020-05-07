/* Similar to fun13.go, instead using ioutil.Readfile, we use 
bufio.ScanLines to read each line of /etc/resolv.conf 
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	defaultResolvConf = "/etc/resolv.conf"
)

func main() {
	var nameServers = []string{}
	getNameServers(&nameServers)
	for _, v := range nameServers {
		fmt.Println(v)
	}
}

func getNameServers(p *[]string) {
	readFile, err := os.Open(defaultResolvConf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			*p = append(*p, fileScanner.Text())
		}
		readFile.Close()
	}
}
