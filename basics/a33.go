// program using uint16 
package main

import "fmt"

func main() {
	var distance, time, speed uint16
	distance = 8000
	time = 160
	speed = distance / time
	fmt.Printf("speed is %d when distance is %d km and time is %d ms\n", speed, distance, time)
}
