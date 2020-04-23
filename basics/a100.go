// Comparing 2 sha256 digests 
package main

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	// cryptographic hash is stored in byte slice
	hash1 := sha256.Sum256([]byte("X"))
	hash2 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n%x\n%t\n%T\n", hash1, hash2, hash1 == hash2, hash1)
}
