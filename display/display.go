package display

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/azdanov/counter-go/stats"
)

type Options struct {
	ShowHeader bool
	ShowLines  bool
	ShowWords  bool
	ShowBytes  bool
}

func (d Options) IsEmpty() bool {
	return !d.ShowLines && !d.ShowWords && !d.ShowBytes
}

func (d Options) ShouldShowLines() bool {
	return d.ShowLines || d.IsEmpty()
}

func (d Options) ShouldShowWords() bool {
	return d.ShowWords || d.IsEmpty()
}

func (d Options) ShouldShowBytes() bool {
	return d.ShowBytes || d.IsEmpty()
}

func PrintHeaders(w io.Writer, opts Options) {
	if !opts.ShowHeader {
		return
	}

	h := []string{}
	if opts.ShouldShowLines() {
		h = append(h, "lines")
	}
	if opts.ShouldShowWords() {
		h = append(h, "words")
	}
	if opts.ShouldShowBytes() {
		h = append(h, "bytes")
	}

	headers := strings.Join(h, "\t") + "\t"
	fmt.Fprintln(w, headers)
}

func Print(w io.Writer, c stats.Counts, opts Options, suffix ...string) {
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

	stats := strings.Join(s, "\t") + "\t"
	fmt.Fprintf(w, "%s", stats)

	suffixes := strings.Join(suffix, " ")
	if suffixes != "" {
		fmt.Fprintf(w, " %s", suffixes)
	}

	fmt.Fprintln(w)
}
