//Simple loop program for reverse of a number

package main

import (
	"fmt"
	"math"
)

func main() {
	var num, temp_num, count int = 4827, 0, 0
	temp := num
	for temp != 0 {
		temp = temp / 10
		count++
	}
	place := math.Pow(float64(10), float64(count -1))
	for num > 0 {
		temp_num += num % 10 * int(place)
		place = place / 10
		num = num / 10
	}
	fmt.Printf("Reverse of number %d is %d\n", 4827, temp_num)
}
