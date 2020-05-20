//Arrays can be used in a way of key value pair
package main

import (
	"fmt"
)

type Attributes int

const (
	uidNumber Attributes = iota
	gidNumber
	cn
	sn
	homeDirectory
)


func main() {

	posixValues := [5]string{uidNumber : "10000", gidNumber : "12000",
		cn: "Foobar", sn: "Foobar", homeDirectory: "/home/foobar"}
	for index, value := range posixValues {
		fmt.Println("index = ", index, "value =", value)
	}
	fmt.Println("uidNumber: ", uidNumber)
	fmt.Println("gidNumber: ", gidNumber)
	fmt.Println("cn: ",  cn)
	fmt.Println("sn: ", sn)
	fmt.Println("homeDirectory: ", homeDirectory)
	fmt.Println(posixValues[uidNumber])
	fmt.Println(posixValues[gidNumber])
	fmt.Println(posixValues[cn])
	fmt.Println(posixValues[sn])
	fmt.Println(posixValues[homeDirectory])
}
	
