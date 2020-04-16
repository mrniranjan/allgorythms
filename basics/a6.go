// Ode to Joy

package main

import "fmt"

func main() {
	var s, s1, s2, s3, s4 string
	s1 = "\tJoy, beautiful spark of divinity"
	s2 = "\tDaughter of Elysium"
	s3 = "\tWe enter, drunk with fire,"
	s4 = `
        Heavenly one, thy sanctuary! 
        Thy magic binds again
        What custom strictly divided;*
        All people become brothers,*
        Where thy gentle wing abides.`

	s = s1 + "\n" + s2 + "\n" + s3 + s4 + "\n"
	fmt.Println(s)
}
