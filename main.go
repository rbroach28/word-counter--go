package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("words.txt")
	
	wordCount := 0

	for _, x := range data {
		if x ==  ' ' {
			wordCount++
		}	
	}

	wordCount++ // for the last word

	fmt.Println(wordCount)

}
