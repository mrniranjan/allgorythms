//slices = variable length sequences : array without a size
// slices have 3 components pointer, length, capacity
// pointer points to first element of array reachable by slice
// length is number of elements length <= capacity
package main
import "fmt"
func main() {
	// declar array
	months := []string{1:"Jan",2:"Feb",3: "Mar", 4:"Apr", 5:"May",
		6:"jun", 7:"jul", 8:"Aug", 9:"Sep", 10:"Oct", 11:"Nov", 12:"Dec"}
	fmt.Println(months)
	fmt.Println("Length of array months is: ", len(months))
	// slices:
	fmt.Println(months[1:13]) // slice months[1:13] refers to whole range
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)
	fmt.Printf("\nLength of slice Q2 is: %d", len(Q2))
	fmt.Printf("\nCapacity of lslice summer is: %d\n", cap(summer))
}
