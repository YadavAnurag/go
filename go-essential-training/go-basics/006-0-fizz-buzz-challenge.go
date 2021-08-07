/*
	print the number between 1 and 20, however, 
	1. if the number is divisible by 3(say 9), print the word "fizz" instead of the number
	1. if the number is divisible by 5(say 10), print the word "buzz" instead of the number
	1. if the number is divisible by 3 and 5(say 15), print the word "fizzbuzz" instead of the number
*/

package main

import (
	"fmt"
)

// better solution is to use %(modulo operator) to check divisibility
func main(){
	for i := 1; i <= 20; i++{
		number := i
		divisibleBy3 := false
		for number > 0{
			number = number - 3
		}
		if number == 0 {divisibleBy3 = true}
		 
		number = i
		divisibleBy5 := false
		for number > 0{
			number = number - 5
		}
		if number == 0 {divisibleBy5 = true}
		

		switch{
			case divisibleBy3 && divisibleBy5:
				fmt.Println("fizz buzz")
			case divisibleBy5:
				fmt.Println("buzz")
			case divisibleBy3:
				fmt.Println("fizz")
			default:
				fmt.Println(i)
		}
	}
}