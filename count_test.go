package main_test

import (
	"io"
	"strings"
	"testing"

	counter "github.com/azdanov/counter-go"
)

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
			name: "string with tabs and multiple spaces",
			f:    strings.NewReader("hello\tworld  \nhello\tworld"),
			want: counter.Counts{Lines: 1, Words: 4, Bytes: 25},
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

func TestCounts_Add(t *testing.T) {
	tests := []struct {
		name string
		c    counter.Counts
		o    counter.Counts
		want counter.Counts
	}{
		{
			name: "add empty counts",
			c:    counter.Counts{Lines: 0, Words: 0, Bytes: 0},
			o:    counter.Counts{Lines: 0, Words: 0, Bytes: 0},
			want: counter.Counts{Lines: 0, Words: 0, Bytes: 0},
		},
		{
			name: "add counts with zero values",
			c:    counter.Counts{Lines: 1, Words: 2, Bytes: 3},
			o:    counter.Counts{Lines: 0, Words: 0, Bytes: 0},
			want: counter.Counts{Lines: 1, Words: 2, Bytes: 3},
		},
		{
			name: "add counts with values",
			c:    counter.Counts{Lines: 1, Words: 2, Bytes: 3},
			o:    counter.Counts{Lines: 4, Words: 5, Bytes: 6},
			want: counter.Counts{Lines: 5, Words: 7, Bytes: 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.Add(tt.o)
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
