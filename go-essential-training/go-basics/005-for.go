// for example

package main

import (
	"fmt"
)

func main(){
	for i := 0; i < 3; i++{
		fmt.Println(i)
	}

	fmt.Println("------")
	for i := 0; i < 3; i++{
		if i > 1{
			break
		}
		fmt.Println(i)
	}

	fmt.Println("------")
	for i := 0; i < 3; i++{
		if i < 2{
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("------")
	// like a while loop
	a := 0
	for a < 4{
		fmt.Println(a)
		a++
	}

	fmt.Println("------")
	// like a while true loop
	b := 0
	for {
		if b > 2{
			break
		}
		fmt.Println(b)
		b++
	}
}