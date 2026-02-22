package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")
	fmt.Println(countWords(data))
}

func countWords(data []byte) int {
	return len(bytes.Fields(data))
}
