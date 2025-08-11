package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "words.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Failed to read file:", err)
	}

	PrintFileContents(file)

	// data, _ := io.ReadAll(file)
	// wordCount := CountWords(data)

	// fmt.Println(wordCount)
}

func PrintFileContents(file *os.File) {
	const bufferSize = 8192
	buffer := make([]byte, bufferSize)
	totalSize := 0
	for {
		size, err := file.Read(buffer)
		if err != nil {
			break
		}

		totalSize += size

	}

	fmt.Println("read from file:", string(buffer))
	fmt.Println("total bytes read:", totalSize)

}

func CountWords(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}
