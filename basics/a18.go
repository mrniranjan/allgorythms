package main

import "fmt"

func main() {
	const version, distro = "sssd-2.0.0-43.el8_0.3.x86_64", 8
	s := fmt.Sprintf("On RHEL %d sssd version is %s", distro, version)
	fmt.Println(s)
}
