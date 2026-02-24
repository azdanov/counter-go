package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	filename := "./words.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	defer file.Close()

	fmt.Println(CountWords(file))
}

func CountWords(handle io.Reader) int {
	scanner := bufio.NewScanner(handle)
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Error scanning file:", err)
	}

	return count
}
