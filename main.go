package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("words.txt")
	
	wordCount := CountWords(data)

	fmt.Printf("The file contains %d words.\n", wordCount)
}

func CountWords(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	wordDetected := false
	wordCount := 0

	for _, x := range data {
		if x == ' ' {
			wordCount++
		} else {
			wordDetected = true
		}
	}

	if !wordDetected {
		return 0
	}

	wordCount++

	return wordCount
}
