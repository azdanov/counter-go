package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	data, err := os.ReadFile("./words.txt")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	fmt.Println(CountWords(data))
}

func CountWords(data []byte) int {
	return len(bytes.Fields(data))
}
