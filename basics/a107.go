//basics of slices
package main

import (
	"fmt"
)

func main() {
	months := [13]string{"", "January", "Febuary", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"}
	//creating a slices for the above array
	/* slice operator s[i:j] where 0<=i<=j<=cap(s)
        refers to elements i though j - 1 of sequence s
        which may be array variable, pointer to an array, 
        if i is omitted, it's 0, if j is omitted its len(s)
         */
	q1 := months[4:7]
	summer  := months[4:8]
	fmt.Println("length of q1 slice is: ", len(q1))
	fmt.Println("Capacity of q1 slice is: ", cap(q1))
	fmt.Println("q1 slices is:", q1)
	fmt.Println("length of q1 slice is: ", len(summer))
	fmt.Println("Capacity of q1 slice is: ", cap(summer))
	fmt.Println("summer slices is:", summer)

	for _, s := range summer {
		for _, q := range q1 {
			if s == q {
				fmt.Printf("%s appears in both \n", s)
			}
		}
	}
	// slicing beyond cap(s) causes panic
	// fmt.Println(summer[:20]) // panic
	// slicing beyond len(s) extends slice
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
}    
