/* Traversing different kind of arrays using loops */
package main

import (
	"fmt"
)

func main() {
	count := []int{1, 2, 3, 4, 5}
	proc := []string{"nahelem", "sandybridge", "Haswell", "Broadwell", "kabylake"}
	/* Method 1 */
	fmt.Println("--------- Method 1 -----------")
	for i:=0;i < len(count); i++ {
		fmt.Println(count[i])
	}
	/* Method 2 */
	fmt.Println("--------- Method 2 -----------") 
	for index, element := range count {
		fmt.Printf("Index = %d, value = %d\n", index, element)
	}
	/* Method 3 similar to method 2, but we can ignore index using _*/
	fmt.Println("--------- Method 3 -----------") 
	for _, value := range count {
		fmt.Println("value = ", value)
	}
	/* Method 4 */
	fmt.Println("--------- Method 4 -----------") 
	j := 0
	for range count {
		fmt.Println(count[j])
		j++
	}
	/* Strings  method 1 */
	fmt.Println("--------- Strings Method 1 -----------") 
	for i :=0; i < len(proc); i++ {
		fmt.Println(proc[i])
	}
    	/* Strings  method 2 */
	fmt.Println("--------- Strings Method 2 -----------")
	for index, element := range proc {
		fmt.Printf("Index = %d, value = %s\n", index, element)
	}
	/* Method 3 similar to method 2, but we can ignore index using _*/
	fmt.Println("--------- Strings Method 3 -----------") 
	for _, value := range proc {
		fmt.Println("value = ", value)
	}
	/* Method 4 */
	fmt.Println("--------- Strings Method 4 -----------") 
	k := 0
	for range proc {
		fmt.Println(proc[k])
		k++
	}	
}
