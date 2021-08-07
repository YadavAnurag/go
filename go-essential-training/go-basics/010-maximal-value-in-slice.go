// find maximal value in slice eg. take a slice of type numbers

package main

import (
	"fmt"
)

func main(){
	// create a slice of type number and initialize a maxValue with 0
	numbers := []int {12, 23, 632, 62, 6, 7, 12, 15}
	maxValue := numbers[0] // considering only positive value slice
	// iterate over slice
	for _, number := range numbers[1:]{
		// if element of slice is greater than maxValue then assign element to maxValue
		if maxValue < number{
			maxValue = number
		}
	}

	fmt.Printf("Max Value = %v", maxValue)
}