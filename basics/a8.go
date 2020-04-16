package main

import "fmt"

func main() {
	name := "sssd"
	version := "2.0.0-43"
	release := "el8_0.3"
	arch := "x86_64"
	pkg := name + "-" + version + "-" + release + "-" + arch + ".rpm"
	fmt.Println(pkg)
}
