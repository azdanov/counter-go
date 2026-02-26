package main_test

import (
	"io"
	"strings"
	"testing"

	counter "github.com/azdanov/counter-go"
)

func Test_countWords(t *testing.T) {
	tests := []struct {
		name string
		r    io.Reader
		want int
	}{
		{
			name: "empty input",
			r:    strings.NewReader(""),
			want: 0,
		},
		{
			name: "single space",
			r:    strings.NewReader(" "),
			want: 0,
		},
		{
			name: "single word",
			r:    strings.NewReader("one"),
			want: 1,
		},
		{
			name: "multiple words",
			r:    strings.NewReader("one two three four five"),
			want: 5,
		},
		{
			name: "words with newlines",
			r:    strings.NewReader("one\n\ntwo three\nfour\nfive"),
			want: 5,
		},
		{
			name: "words with multiple spaces",
			r:    strings.NewReader("one   two    three     four      five"),
			want: 5,
		},
		{
			name: "words with tabs and newlines",
			r:    strings.NewReader("one\t\two\nthree\tfour\n\nfive"),
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := counter.CountWords(tt.r)
			if got != tt.want {
				t.Errorf("countWords() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestCountLines(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		r    io.Reader
		want int
	}{
		{
			name: "empty input",
			r:    strings.NewReader(""),
			want: 0,
		},
		{
			name: "single line without newline",
			r:    strings.NewReader("one line"),
			want: 0,
		},
		{
			name: "multiple lines",
			r:    strings.NewReader("line one\nline two\nline three\nline four\nline five"),
			want: 4,
		},
		{
			name: "lines with empty lines",
			r:    strings.NewReader("line one\n\nline two\n\nline three\n\nline four\n\nline five"),
			want: 8,
		},
		{
			name: "lines with trailing newline",
			r:    strings.NewReader("line one\nline two\nline three\nline four\nline five\n"),
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := counter.CountLines(tt.r)
			if got != tt.want {
				t.Errorf("CountLines() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestCountBytes(t *testing.T) {
	tests := []struct {
		name string
		r    io.Reader
		want int
	}{
		{
			name: "empty input",
			r:    strings.NewReader(""),
			want: 0,
		},
		{
			name: "single character",
			r:    strings.NewReader("a"),
			want: 1,
		},
		{
			name: "multiple characters",
			r:    strings.NewReader("hello world"),
			want: 11,
		},
		{
			name: "string with newlines",
			r:    strings.NewReader("hello world\nhello world\nhello world"),
			want: 35,
		},
		{
			name: "string with unicode characters",
			r:    strings.NewReader("привет мир"),
			want: 19, // 18 bytes for the characters + 1 byte for the space
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := counter.CountBytes(tt.r)
			if got != tt.want {
				t.Errorf("CountBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		name string
		f    io.ReadSeeker
		want counter.Counts
	}{
		{
			name: "empty input",
			f:    strings.NewReader(""),
			want: counter.Counts{Lines: 0, Words: 0, Bytes: 0},
		},
		{
			name: "single line with words",
			f:    strings.NewReader("hello world"),
			want: counter.Counts{Lines: 0, Words: 2, Bytes: 11},
		},
		{
			name: "multiple lines with words",
			f:    strings.NewReader("hello world\nhello world\nhello world"),
			want: counter.Counts{Lines: 2, Words: 6, Bytes: 35},
		},
		{
			name: "string with unicode characters",
			f:    strings.NewReader("привет мир"),
			want: counter.Counts{Lines: 0, Words: 2, Bytes: 19},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := counter.Count(tt.f)
			if got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
