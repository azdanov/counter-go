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

type FileCount struct {
	counts   stats.Counts
	filename string
}

func main() {
	ch := make(chan FileCount)
	wg := sync.WaitGroup{}

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

	for _, filename := range filenames {
		wg.Go(func() {
			counts, err := HandleFileCount(filename)
			if err != nil {
				hadErr = true
				fmt.Fprintf(os.Stderr, "%s: %v\n", binName, err)
				return
			}
			ch <- FileCount{counts: counts, filename: filename}
		})
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for fc := range ch {
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

func HandleFileCount(filename string) (stats.Counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return stats.Counts{}, err
	}
	defer file.Close()

	return stats.Count(file), nil
}
