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

func CountWords(r io.Reader) int {
	scanner := bufio.NewScanner(r)
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

func CountLines(r io.Reader) int {
	reader := bufio.NewReader(r)

	count := 0
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln("Error reading file:", err)
		}
		if r == '\n' {
			count++
		}
	}

	return count
}

func CountBytes(r io.Reader) int {
	count, err := io.Copy(io.Discard, r)
	if err != nil {
		log.Fatalln("Error counting bytes:", err)
	}
	return int(count)
}
