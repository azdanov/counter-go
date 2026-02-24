package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(0)

	binName := filepath.Base(os.Args[0])

	filenames := os.Args[1:]
	total := 0
	hadErr := false

	for _, filename := range filenames {
		count, err := CountWordsInFile(filename)
		if err != nil {
			hadErr = true
			fmt.Fprintf(os.Stderr, "%s: %v\n", binName, err)
			continue
		}

		fmt.Printf("%d %s\n", count, filename)
		total += count
	}

	if len(filenames) == 0 {
		count := CountWords(os.Stdin)
		fmt.Printf("%d\n", count)
	}

	if len(filenames) > 1 {
		fmt.Printf("%d total\n", total)
	}

	if hadErr {
		os.Exit(1)
	}
}

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
