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
	total := Counts{}
	hadErr := false

	for _, filename := range filenames {
		counts, err := HandleFileCount(filename)
		if err != nil {
			hadErr = true
			fmt.Fprintf(os.Stderr, "%s: %v\n", binName, err)
			continue
		}

		fmt.Printf("%s %s\n", counts, filename)
		total.Lines += counts.Lines
		total.Words += counts.Words
		total.Bytes += counts.Bytes
	}

	if len(filenames) == 0 {
		counts := Count(os.Stdin)
		fmt.Printf("%s\n", counts)
	}

	if len(filenames) > 1 {
		fmt.Printf("%s total\n", total)
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
