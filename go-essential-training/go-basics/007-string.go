// string example

package main

import (
	"fmt"
)

func main(){
	book := "India is a beautiful country"
	fmt.Println(book)

	// index access, it returns byte(uint8 is a byte in go)
	fmt.Printf("book[0] = %v and type = %T\n", book[0], book[0])

	// len access in byte
	fmt.Println(len(book))

	// slice range
	fmt.Println(book[3:10])

	// slice range only start
	fmt.Println(book[6:])

	// slice range only end
	fmt.Println(book[:15])

	// string on the go assignment
	// strings on the go is immutable
	// book[0] = "A" 

	// concatenate string but will return the string no assignment
	fmt.Println("Uttar Pradesh" + book[5: len(book)-7] + "state")

	// Go is based on Unicode Character, so it supports any unicode character

	// multi line
	slogan := `
		Saare jahan se achha
		Hindostan hamara
		Hum bulbule hain iski
		Ye ghulsitan hamara hamara...
	`
	fmt.Println(slogan)

}