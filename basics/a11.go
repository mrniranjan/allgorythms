package main

import "fmt"

const sssd_ver = "sssd-2.0.0-43.el8_0.3.x86_64"

func main() {
	var s string
	s = "Version of sssd on RHEL8 is %s"
	fmt.Printf(s, sssd_ver)
}
