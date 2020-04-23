//Parsing IP Address

package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	valid_ip_addr := "192.168.122.132"
	fmt.Println(parseIpAddr(valid_ip_addr))
	invalide_ip_addr := "1968.168.122.132"
	fmt.Println(parseIpAddr(invalide_ip_addr))
}

// Parses ip address
func parseIpAddr(ip_addr string) bool {
	s := strings.Split(ip_addr, ".")
	for _, value := range s {
		if s, err := strconv.Atoi(value); err == nil {
			if s > 255 {
				return false
			}
		}
	}
	return true
}
