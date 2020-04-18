//maps: Deleting values from maps

package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	fmt.Println("map ages:", ages)
	delete(ages, "alice")
	fmt.Println("after deleting slice", ages)
	// check what happens when we try to delete a key which doesn't exit
	// so no error is thrown out when trying to delete a key which doesn't exist
	_, status := ages["bob"]
	if status {
		fmt.Println("key bob exists")
		delete(ages, "bob")
	} else {
		fmt.Println("key bob doesn't exist")
	}
}
