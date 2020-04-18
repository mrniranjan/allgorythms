//strings and booleans

package main

import "fmt"

func main() {
	sssd_rhel80_ver := "sssd-2.0.0-43"
	sssd_rhel81_ver := "sssd-2.2.0-19"
	ret1 := sssd_rhel80_ver == sssd_rhel81_ver
	ret2 := sssd_rhel81_ver == "sssd-2.2.0-19"
	fmt.Printf("The type of result ret1 is %T\n", ret1)
	fmt.Printf("The type of result ret2 is %T\n", ret2)
	fmt.Println(ret1)
	fmt.Println(ret2)
}
