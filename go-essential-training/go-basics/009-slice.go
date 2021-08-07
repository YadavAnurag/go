// slice - sequence of elements in go, and all elements must be of same size

package main

import (
	"fmt"
)

func main(){
	states := []string {"UP", "MP", "JK", "LD", "GJ"}
	
	fmt.Printf("states = %v, type = %T \n", states, states)

	// length
	fmt.Printf("Length of states = %v \n", len(states))

	// index access
	fmt.Printf("First state = %v \n", states[0])

	// slices of slice
	fmt.Printf("First two states = %v \n", states[:2])

	// iterate over slice
	for iterator := 0; iterator < len(states); iterator++{
		fmt.Printf("%v ", states[iterator])
	}

	// iteration over slice using single value range
	fmt.Println("\nIteration over slice using single value range")
	for iterator := range states{
		fmt.Printf("%v ", states[iterator])
	}

	// iteration over slice using double value range
	fmt.Println("\nIteration over slice using double value range")
	for index, state := range states{
		fmt.Printf("%v = %v \n", index, state)
	}

	// iteration over slice using double value range and index ignore
	fmt.Printf("\nIteration over slice using double value range, and index ignore\n")
	for _, state := range states{
		fmt.Println(state)
	}

	// slide append
	fmt.Println("\nslice append")
	states = append(states, "SK")
	fmt.Printf("%v \n", states)

}