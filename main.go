package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type DisplayOptions struct {
	ShowLines bool
	ShowWords bool
	ShowBytes bool
}

func (d DisplayOptions) IsEmpty() bool {
	return !d.ShowLines && !d.ShowWords && !d.ShowBytes
}

func (d DisplayOptions) ShouldShowLines() bool {
	return d.ShowLines || d.IsEmpty()
}

func (d DisplayOptions) ShouldShowWords() bool {
	return d.ShowWords || d.IsEmpty()
}

func (d DisplayOptions) ShouldShowBytes() bool {
	return d.ShowBytes || d.IsEmpty()
}

func main() {
	opts := DisplayOptions{}
	flag.BoolVar(&opts.ShowLines, "l", false, "Show line count")
	flag.BoolVar(&opts.ShowWords, "w", false, "Show word count")
	flag.BoolVar(&opts.ShowBytes, "c", false, "Show byte count")
	flag.Parse()

	log.SetFlags(0)

	binName := filepath.Base(os.Args[0])
	filenames := flag.Args()
	total := Counts{}
	hadErr := false

	for _, filename := range filenames {
		counts, err := HandleFileCount(filename)
		if err != nil {
			hadErr = true
			fmt.Fprintf(os.Stderr, "%s: %v\n", binName, err)
			continue
		}

		Print(os.Stdout, counts, opts, filename)
		total = total.Add(counts)
	}

	if len(filenames) == 0 {
		counts := Count(os.Stdin)
		Print(os.Stdout, counts, opts)
	}

	if len(filenames) > 1 {
		Print(os.Stdout, total, opts, "total")
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

func Print(w io.Writer, c Counts, opts DisplayOptions, suffix ...string) {
	s := []string{}
	if opts.ShouldShowLines() {
		s = append(s, strconv.Itoa(c.Lines))
	}
	if opts.ShouldShowWords() {
		s = append(s, strconv.Itoa(c.Words))
	}
	if opts.ShouldShowBytes() {
		s = append(s, strconv.Itoa(c.Bytes))
	}
	s = append(s, suffix...)

	fmt.Fprintln(w, strings.Join(s, " "))
}
