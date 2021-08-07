// map data structure
// all keys must be of same type, and all values must be of same type

package main

import (
	"fmt"
)

func main(){
	statesLiteracyPercentage := map[string]float64{
		"UP": 90.2,
		"MP": 91,
		"BR": 89.3,
		"KL": 94.6, // must have the trailling comma in multiline
	}

	fmt.Printf("%v \n", statesLiteracyPercentage)

	// number of items
	fmt.Printf("%v \n", len(statesLiteracyPercentage))

	// access using keys
	fmt.Printf("Literacy Percentage of UP = %v\n", statesLiteracyPercentage["UP"])
	// if there is no key will return 0
	fmt.Printf("Literacy Percentage of OD = %v\n", statesLiteracyPercentage["OD"])

	// using key and found 
	literacy, found := statesLiteracyPercentage["MP"]
	if !found{
		fmt.Printf("MP literacy data not found \n")
	}else{
		fmt.Printf("MP literacy = %v\n", literacy)
	}

	// set value on the go
	statesLiteracyPercentage["AS"] = 89.31
	fmt.Printf("AS literacy rate = %v\n", statesLiteracyPercentage["AS"])

	// delete any key
	delete(statesLiteracyPercentage, "MP")
	fmt.Println("Literacy = ", statesLiteracyPercentage)

	// iteration over map using key
	for state := range statesLiteracyPercentage{
		fmt.Printf("%v ", state)
	}

	// iteration over map using key and value
	for state, literacyPercentage := range statesLiteracyPercentage{
		fmt.Println(state, "literacy", literacyPercentage)
	} 
}