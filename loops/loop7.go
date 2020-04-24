/* Iterating through loop and appending */
package main

import "fmt"

func main() {
	var count = []int{}
	/* Method 1 */
	for i:=0; i < 5; i++ {
		count = append(count, i)
	}
	fmt.Println(count)
	/* method 2 */
	for _, value := range count {
		value = value +1
		count = append(count, value)
	}
	fmt.Println(count)
	
}
