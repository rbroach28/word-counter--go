package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	filename := "words.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Failed to read file:", err)
	}

	wordCount :=CountWordsInFile(file)
	fmt.Println(wordCount)

}

func CountWordsInFile(file *os.File) int {
	wordCount := 0
	isInsideWord := false

	const bufferSize = 8192
	buffer := make([]byte, bufferSize)
	
	for {
		size, err := file.Read(buffer)
		if err != nil {
			break
		}

		isInsideWord = !unicode.IsSpace(rune(buffer[0])) && isInsideWord // If true last buffer threshold was inside a word
		bufferCount := CountWords(buffer[:size])
		if isInsideWord {
			bufferCount--
		}

		wordCount += bufferCount
		isInsideWord = !unicode.IsSpace(rune(buffer[size-1]))
	}

	return wordCount

}

func CountWords(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}
