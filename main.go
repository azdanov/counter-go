package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, _ := os.ReadFile("./words.txt")

	fmt.Println(countWords(bytes))
}

func countWords(bytes []byte) int {
	wordCount := 0

	for _, b := range bytes {
		if b == ' ' || b == '\n' {
			wordCount++
		}
	}

	return wordCount
}
