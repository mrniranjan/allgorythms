package main

import "fmt"

func main() {
	s1 := fmt.Sprint("sssd-2.0.0-43")
	s2 := fmt.Sprint("-el8_0.3")
	s3 := fmt.Sprint(".x86_64")
	fmt.Print(s1, s2, s3)
}
