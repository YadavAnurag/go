// if statement examples

package main 

import (
	"fmt"
)

func main(){
	x := 10

	if x > 5 {
		fmt.Println("x is more than 5")
	}

	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("x is not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}
	
	if x < 20 || x > 30 {
		fmt.Println("x is out of range")
	}

	a := 2.0
	b := 3.0
	if fraction := b/a + 1; fraction > a && fraction < b{
		fmt.Println("fraction is in range a and b")
	}
}