// switch example

package main

import (
	"fmt"
)

func main(){
	salary := 100.0

	switch salary{
		case 20:
			fmt.Println("20")
		case 50:
			fmt.Println("50")
		case 100:
			fmt.Println("100")
		default:
			fmt.Println("so good")
	}

	// switch without expression
	switch {
		case salary > 100:
			fmt.Println("is so good")
		case salary > 50:
			fmt.Println("more than 50")
		default:
			fmt.Println("its ok")	
	}
}