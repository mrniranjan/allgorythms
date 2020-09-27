package main

import (
	"fmt"
)

func main() {
	s1 := []byte("\u0938 \u0930\u0947 \u0917 \u092e ")
	s2 := []byte("\u0938 \u0930\u0947 \u092e \u0917 ")
	s3 := []byte("\u0938 \u0917 \u092e \u0930\u0947 ")
	s4 := []byte("\u0930\u0947 \u0917 \u092e \u0938 ")
	fmt.Println(string(s1),string(s2), string(s3), string(s4))
}
