// maps: operations using keys 
package main

import "fmt"

func main() {
	data := make(map[string]int)
	fmt.Println("Length of map after initiating is: ", len(data))
	s1 := "Go Programming Lang"
	fmt.Println("value of key 'Go Programming Lang' is : ", data[s1])
	fmt.Println(len(data))
	// default value (v) os key "Go programming Lang" is 0  which incremented to 1 
	data[s1]++
	fmt.Println("length of map data is: ", len(data))
	fmt.Println(data[s1])
	a := data[s1]
	fmt.Println("a = ", a)
}
