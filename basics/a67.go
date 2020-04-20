// maps: Program to check when key doesn't exist
package main

import "fmt"

func main() {
	proc := make(map[string]string)
	proc["CodeName"] = "Whiskeylake"
	proc["Processor Number"] = "i7-8665UE"
	// when key exists , ret is value
	// and status is boolean value true
	ret, status := proc["CodeName"]
	fmt.Println("ret = ", ret)
	fmt.Println("status =", status)
	// when key doesn't exit , result is null
	// stat is false
	result, stat := proc["Cores"]
	fmt.Println("ret =", result)
	fmt.Println("status =", stat)
	fmt.Println(proc)
}
