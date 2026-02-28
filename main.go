package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/azdanov/counter-go/display"
	"github.com/azdanov/counter-go/stats"
)

func main() {
	do := display.Options{}
	flag.BoolVar(&do.ShowHeader, "headers", false, "Show header for each column")
	flag.BoolVar(&do.ShowLines, "l", false, "Show line count")
	flag.BoolVar(&do.ShowWords, "w", false, "Show word count")
	flag.BoolVar(&do.ShowBytes, "c", false, "Show byte count")
	flag.Parse()

	log.SetFlags(0)

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)

	binName := filepath.Base(os.Args[0])
	filenames := flag.Args()
	total := stats.Counts{}
	hadErr := false

	display.PrintHeaders(tw, do)

	for _, filename := range filenames {
		counts, err := HandleFileCount(filename)
		if err != nil {
			hadErr = true
			fmt.Fprintf(tw, "%s: %v\n", binName, err)
			continue
		}

		display.Print(tw, counts, do, filename)
		total = total.Add(counts)
	}

	if len(filenames) == 0 {
		counts := stats.Count(os.Stdin)
		display.Print(tw, counts, do)
	}

	if len(filenames) > 1 {
		display.Print(tw, total, do, "total")
	}

	tw.Flush()
	if hadErr {
		os.Exit(1)
	}
}

func HandleFileCount(filename string) (stats.Counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return stats.Counts{}, err
	}
	defer file.Close()

	return stats.Count(file), nil
}
