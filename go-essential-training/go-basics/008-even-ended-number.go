// even ended number - a number with the same first and last digits for example 1, 11, 121, 1331, 12231
// Find how many even-ended numbers result from multiplying two four-digits numbers


package main

import (
	"fmt"
)

func main(){
	count := 0	
	for firstNumber := 1000; firstNumber <= 9999; firstNumber++{
		for secondNumber := firstNumber; secondNumber <= 9999; secondNumber++{ // don't count twice please initialize second number as first number because 1x3 === 3x1
			// multiply both number and store it in a const
			multiplicatedNumber := firstNumber * secondNumber

			// convert multiplication int to a string using Sprintf that returns string
			multiplicatedString := fmt.Sprintf("%d", multiplicatedNumber)

			// check if first index and last index are same
			// if same then count++
			if multiplicatedString[0] == multiplicatedString[len(multiplicatedString)-1]{
				count++
			}
		}
	}

	fmt.Println("Total ",  count)
}