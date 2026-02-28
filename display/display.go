package display

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/azdanov/counter-go/stats"
)

type Options struct {
	opts NewOptionsArgs
}

type NewOptionsArgs struct {
	ShowHeaders bool
	ShowLines   bool
	ShowWords   bool
	ShowBytes   bool
}

func NewOptions(args NewOptionsArgs) Options {
	return Options{
		opts: args,
	}
}

func (d Options) IsEmpty() bool {
	return !d.opts.ShowLines && !d.opts.ShowWords && !d.opts.ShowBytes
}

func (d Options) ShouldShowLines() bool {
	return d.opts.ShowLines || d.IsEmpty()
}

func (d Options) ShouldShowWords() bool {
	return d.opts.ShowWords || d.IsEmpty()
}

func (d Options) ShouldShowBytes() bool {
	return d.opts.ShowBytes || d.IsEmpty()
}

func PrintHeaders(w io.Writer, d Options) {
	if !d.opts.ShowHeaders {
		return
	}

	h := []string{}
	if d.ShouldShowLines() {
		h = append(h, "lines")
	}
	if d.ShouldShowWords() {
		h = append(h, "words")
	}
	if d.ShouldShowBytes() {
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
