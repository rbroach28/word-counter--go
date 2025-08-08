package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("words.txt")
	
	wordCount := countWords(data)

	fmt.Printf("The file contains %d words.\n", wordCount)
}

func countWords(data []byte) int {
	wordCount := 0

	for _, x := range data {
		if x == ' ' {
			wordCount++
		}
	}

	wordCount++ // for the last word

	return wordCount
}
