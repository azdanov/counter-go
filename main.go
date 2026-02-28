package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"text/tabwriter"

	"github.com/azdanov/counter-go/display"
	"github.com/azdanov/counter-go/stats"
)

type FileCounts struct {
	counts   stats.Counts
	filename string
	err      error
}

func main() {
	args := display.NewOptionsArgs{}
	flag.BoolVar(&args.ShowHeaders, "headers", false, "Show header for each column")
	flag.BoolVar(&args.ShowLines, "l", false, "Show line count")
	flag.BoolVar(&args.ShowWords, "w", false, "Show word count")
	flag.BoolVar(&args.ShowBytes, "c", false, "Show byte count")
	flag.Parse()
	do := display.NewOptions(args)

	log.SetFlags(0)

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)

	binName := filepath.Base(os.Args[0])
	filenames := flag.Args()
	total := stats.Counts{}
	hadErr := false

	display.PrintHeaders(tw, do)

	ch := CountFiles(filenames)
	for fc := range ch {
		if fc.err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", binName, fc.err)
			hadErr = true
			continue
		}
		display.Print(tw, fc.counts, do, fc.filename)
		total = total.Add(fc.counts)
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

func CountFiles(filenames []string) <-chan FileCounts {
	ch := make(chan FileCounts)
	wg := sync.WaitGroup{}

	for _, filename := range filenames {
		wg.Go(func() {
			counts, err := HandleFileCount(filename)
			if err != nil {
				ch <- FileCounts{err: err, filename: filename}
				return
			}
			ch <- FileCounts{counts: counts, filename: filename}
		})
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func HandleFileCount(filename string) (stats.Counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return stats.Counts{}, err
	}
	defer file.Close()

	return stats.Count(file), nil
}
