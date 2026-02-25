package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func CountWordsInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return CountWords(file), nil
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
