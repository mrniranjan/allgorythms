//LCM

package main

import "fmt"

func main() {
	numbers := []int{9, 12}
	multiples := make(map[int][]int)
	commonMultiples := []int{}
	for index := range numbers {
		for i:=1; i < 10; i++ {
			multiples[numbers[index]] = append(multiples[numbers[index]], numbers[index] * i)
		}
	}
	fmt.Println(multiples)
	// Get the first key, map , i.e values of 78 from map 
	curDataSet := multiples[numbers[0]]
	for key, value := range multiples {
		if key != numbers[0] { // we ignore the first key since it's already in currentdataset
			for i :=0; i < len(value); i++ {
				for j :=0; j < len(curDataSet); j++ {
					if value[i] == curDataSet[j]  {
						commonMultiples = append(commonMultiples, value[i])
					}
				}
			}
		}
	}
	fmt.Printf("\nLCM of %d and %d is %d\n", numbers[0], numbers[1], commonMultiples[0])
}
	
