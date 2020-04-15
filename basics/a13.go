package main

import "fmt"

func main() {

	s1 := fmt.Sprint("sssd-2.0.0-43")
	s2 := fmt.Sprint("-el8_0.3")
	s3 := fmt.Sprint(".x86_64")
	s4 := fmt.Sprint(s1, s2, s3, " is the version for RHEL", 8)
	fmt.Print(s4)
}
