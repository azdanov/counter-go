package stats_test

import (
	"io"
	"strings"
	"testing"

	"github.com/azdanov/counter-go/stats"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name string
		f    io.ReadSeeker
		want stats.Counts
	}{
		{
			name: "empty input",
			f:    strings.NewReader(""),
			want: stats.Counts{Lines: 0, Words: 0, Bytes: 0},
		},
		{
			name: "single line with words",
			f:    strings.NewReader("hello world"),
			want: stats.Counts{Lines: 0, Words: 2, Bytes: 11},
		},
		{
			name: "multiple lines with words",
			f:    strings.NewReader("hello world\nhello world\nhello world"),
			want: stats.Counts{Lines: 2, Words: 6, Bytes: 35},
		},
		{
			name: "string with tabs and multiple spaces",
			f:    strings.NewReader("hello\tworld  \nhello\tworld"),
			want: stats.Counts{Lines: 1, Words: 4, Bytes: 25},
		},
		{
			name: "string with unicode characters",
			f:    strings.NewReader("привет мир"),
			want: stats.Counts{Lines: 0, Words: 2, Bytes: 19},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stats.Count(tt.f)
			if got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCounts_Add(t *testing.T) {
	tests := []struct {
		name string
		c    stats.Counts
		o    stats.Counts
		want stats.Counts
	}{
		{
			name: "add empty counts",
			c:    stats.Counts{Lines: 0, Words: 0, Bytes: 0},
			o:    stats.Counts{Lines: 0, Words: 0, Bytes: 0},
			want: stats.Counts{Lines: 0, Words: 0, Bytes: 0},
		},
		{
			name: "add counts with zero values",
			c:    stats.Counts{Lines: 1, Words: 2, Bytes: 3},
			o:    stats.Counts{Lines: 0, Words: 0, Bytes: 0},
			want: stats.Counts{Lines: 1, Words: 2, Bytes: 3},
		},
		{
			name: "add counts with values",
			c:    stats.Counts{Lines: 1, Words: 2, Bytes: 3},
			o:    stats.Counts{Lines: 4, Words: 5, Bytes: 6},
			want: stats.Counts{Lines: 5, Words: 7, Bytes: 9},
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

var benchData = []io.ReadSeeker{
	strings.NewReader("hello world\thello world\nhello world\n"),
	strings.NewReader("привет мир\nпривет мир\nпривет мир"),
	strings.NewReader(strings.Repeat("hello world ", 1000) + "\n" + strings.Repeat("привет мир ", 1000)),
}

func BenchmarkCount(b *testing.B) {
	for i := range b.N {
		stats.Count(benchData[i%len(benchData)])
	}
}
