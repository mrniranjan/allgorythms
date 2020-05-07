//check prefix

package main

import (
	"fmt"
)

func main() {
	s := "dis-engage"
	fmt.Println(hasPrefix(s, "dis"))
	s = "sssd.conf"
	fmt.Println(hasSuffix(s, "conf"))
	s = "192.168.122.252"
	fmt.Println(contains(s, "168.122"))
}

func hasPrefix(s, prefix  string)  bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s) - len(suffix):] == suffix
}

func contains(s, substr string) bool {
	for i :=0; i < len(s); i++ {
		if hasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
