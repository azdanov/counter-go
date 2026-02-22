package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")
	fmt.Println(CountWords(data))
}

func CountWords(data []byte) int {
	return len(bytes.Fields(data))
}
