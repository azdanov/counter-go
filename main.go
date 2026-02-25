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
