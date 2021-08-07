// calculate the mean of three numbers

package main

import (
	"fmt"
)

func main(){
	var x float32
	var y float32

	x = 1.0
	y = 2.0
	z := 0.0 // inline declaration and assignment uses type as of value provided
	// z type if float64

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)
	fmt.Printf("z=%v, type of %T\n", z, z)

	mean := (x + y + float32(z)) / 3

	fmt.Printf("mean=%v, type of %T\n", mean, mean)

}