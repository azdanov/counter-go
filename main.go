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

	if len(os.Args) < 2 {
		name := os.Args[0]
		log.Fatalf("Usage: %s <filename1> [<filename2> ...]", name)
	}

	filenames := os.Args[1:]
	total := 0

	for _, filename := range filenames {
		count := CountWordsInFile(filename)
		fmt.Printf("%d %s\n", count, filename)
		total += count
	}

	if len(filenames) > 1 {
		fmt.Printf("%d total\n", total)
	}
}

func CountWordsInFile(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	defer file.Close()

	return CountWords(file)
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
