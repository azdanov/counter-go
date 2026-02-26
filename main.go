package main

import (
	"fmt"
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
		counts, err := HandleFileCount(filename)
		if err != nil {
			hadErr = true
			fmt.Fprintf(os.Stderr, "%s: %v\n", binName, err)
			continue
		}

		fmt.Printf("%d %d %d %s\n", counts.Lines, counts.Words, counts.Bytes, filename)
		total += counts.Words
	}

	if len(filenames) == 0 {
		counts := Count(os.Stdin)
		fmt.Printf("%d %d %d\n", counts.Lines, counts.Words, counts.Bytes)
	}

	if len(filenames) > 1 {
		fmt.Printf("%d total\n", total)
	}

	if hadErr {
		os.Exit(1)
	}
}

func HandleFileCount(filename string) (Counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Counts{}, err
	}
	defer file.Close()

	return Count(file), nil
}
