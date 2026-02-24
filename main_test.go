package main_test

import (
	"io"
	"strings"
	"testing"

	counter "github.com/azdanov/counter-go"
)

func Test_countWords(t *testing.T) {
	tests := []struct {
		name   string
		handle io.Reader
		want   int
	}{
		{
			name:   "empty input",
			handle: strings.NewReader(""),
			want:   0,
		},
		{
			name:   "single space",
			handle: strings.NewReader(" "),
			want:   0,
		},
		{
			name:   "single word",
			handle: strings.NewReader("one"),
			want:   1,
		},
		{
			name:   "multiple words",
			handle: strings.NewReader("one two three four five"),
			want:   5,
		},
		{
			name:   "words with newlines",
			handle: strings.NewReader("one\n\ntwo three\nfour\nfive"),
			want:   5,
		},
		{
			name:   "words with multiple spaces",
			handle: strings.NewReader("one   two    three     four      five"),
			want:   5,
		},
		{
			name:   "words with tabs and newlines",
			handle: strings.NewReader("one\t\two\nthree\tfour\n\nfive"),
			want:   5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := counter.CountWords(tt.handle)
			if got != tt.want {
				t.Errorf("countWords() = %d, want %d", got, tt.want)
			}
		})
	}
}
