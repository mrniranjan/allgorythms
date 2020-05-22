/*Program that prints the sha256 hash
of it's standard input by default
but supports command-line flag to print
sha384 or sha512 instead*/
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"flag"
)

var (
	size = flag.Int("size", 256 ,"Specify the secure hash size")
	str = flag.String("str", "", "Specify string to be hashed")
)
	
func main() {
	flag.Parse()
	s := myCrypto(size, str)
	fmt.Printf("hash of string '%s' is %s\n", *str, s)
}
//Function myCrypto takes hash size and string to be hashed
// and returns the secure hash
func myCrypto(size *int, str *string) (s string) {
	s = ""
	switch *size {
	case 384:
		hash := sha512.Sum384([]byte(*str))
		s = fmt.Sprintf("%x", hash)
	case 512:
		hash := sha512.Sum512([]byte(*str))
		s = fmt.Sprintf("%x", hash)
	default:
		hash := sha256.Sum256([]byte(*str))
		s = fmt.Sprintf("%x", hash)
	}
	return 
}
