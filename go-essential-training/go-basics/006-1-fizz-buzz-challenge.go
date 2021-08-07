/*
	print the number between 1 and 20, however, 
	1. if the number is divisible by 3(say 9), print the word "fizz" instead of the number
	1. if the number is divisible by 5(say 10), print the word "buzz" instead of the number
	1. if the number is divisible by 3 and 5(say 15), print the word "fizzbuzz" instead of the number
*/
// better solution

package main

import (
	"fmt"
)

func main(){
	for i := 1; i <= 20; i++{
		switch{
			case i%3 == 0 && i%5 == 0:
				fmt.Println("fizzbuzz")
			case i%5 == 0:
				fmt.Println("buzz")
			case i%3 == 0:
				fmt.Println("fizz")
			case i==i:
				fmt.Println(i)
			default:
				fmt.Println()
		}
	}
}


