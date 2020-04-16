package main

import "fmt"

func main() {
	var version int
	version = 8
	s1 := "sssd-2.0.0-43.el8_0.3.x86_64"
	s2 := ("sssd version is %s on RHEL%d\n")
	fmt.Printf(s2, s1, version)
}
