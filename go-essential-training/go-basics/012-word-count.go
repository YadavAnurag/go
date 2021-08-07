// count how many times a word apperas in a tex

package main

import (
	"fmt"
	"strings"
)

func main(){
	// create a text
	// poem := `
	// 	Twinkle twinkle little star,
	// 	How I wonder what you are,
	// 	Up above the world so high,
	// 	Like a diamond in the sky
	// `
	anotherPoem := `
		Needles and pins
		Needles and pins
		Sew me a sail
		To catch me the wind
	`

	// split text on space and store it in a slice
	words := strings.Fields(strings.ToLower(anotherPoem))

	// create a map[string]int64 to store word count
	wordsCount := map[string]int64 {}

	// iterate over splitted text
	
	for _, word := range words{
		// store first word
		// create count initialize to 0
		var wordCount int64
		wordCount = 0
		// iterate over splitted text
			for _, anotherWord := range words{
				// match if word is found

				if fmt.Sprintf("%s", word[len(word)-1]) == ","{
					if fmt.Sprintf("%s", word[len(word)-2]) == anotherWord{
						wordCount++
					}
				} else {
					if word == anotherWord{
						// count++
						wordCount++
					}
				}
				
				// assign key and value to map
				if fmt.Sprintf("%s", word[len(word)-1]) == ","{
					wordsCount[fmt.Sprintf("%s", word[len(word)-1])] = wordCount
				}else{
					wordsCount[word] = wordCount
				}
			}
	}
	// iterate over map and print 	
	// for word, wordCount := range wordsCount{
	// 	fmt.Println(word, " = ", wordCount)
	// }
	fmt.Println(wordsCount)

}