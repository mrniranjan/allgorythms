// This program illustrate simple encryption and decryption
// which does simple xor 
package main

import (
	"fmt"
)

func main() {
	plainText := "Hello World"
	e := encryptString(plainText)
	fmt.Printf("Encrypted %s to %s ", s, string(e))
	d := decryptString(string(e))
	fmt.Println(string(d) == plainText)
}
// encryptString encrypts string 
func encryptString(s string) []byte {
	b := make([]byte, len(s))
	for i:=0; i < len(s); i++ {
		a := s[i] ^ 23
		b[i] = a 
	}
	return b
}
// decryptString decrypts string 
func decryptString(s string) []byte {
	c := make([]byte, len(s))
	for j:=0; j < len(s); j++ {
		a := s[j] ^ 23
		c[j] = a
	}
	return c
}
