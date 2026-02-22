package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, _ := os.ReadFile("./words.txt")

	wordCount := 0

	for _, b := range bytes {
		if b == ' ' || b == '\n' {
			wordCount++
		}
	}

	fmt.Println(wordCount)
}
