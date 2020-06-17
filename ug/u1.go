//This program looks up users and groups
package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, _ := user.Current()
	fmt.Println(*user)
	fmt.Printf("%T", *user)
}
