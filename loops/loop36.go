//How to get the keys as string array from map in go lang?
//https://stackoverflow.com/questions/41690156/how-to-get-the-keys-as-string-array-from-map-in-go-lang

package main

import (
	"fmt"
	"strings"
)

func main() {
	data := map[string]int {
		"Na": 11,
		"Zn": 30,
	}
	myStr := make([]string, 0, len(data))
	for key, _ := range data {
		myStr = append(myStr, key)
	}
	fmt.Println(strings.Join(myStr, ""))
}
