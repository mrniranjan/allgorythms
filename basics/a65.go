//maps: map is a reference to hash table
//map is written map[k]v, where k is key, v is value
// it's unordered 

package main
import "fmt"
func main() {
	var m map[string]int
	m = make(map[string]int)
	m["alice"]= 21
	m["bob"] = 22
	fmt.Println(m)
	// another method to create map
	ages := map[string]int{
		"eve": 31,
		"adam": 32,
	}
	fmt.Println(ages)
	fmt.Println("Age of eve is:", ages["eve"])
}
