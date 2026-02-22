package main_test

import "testing"
import counter "github.com/azdanov/counter-go"

func Test_countWords(t *testing.T) {
	tests := []struct {
		name  string
		bytes []byte
		want  int
	}{
		{
			name:  "empty input",
			bytes: []byte(""),
			want:  0,
		},
		{
			name:  "single space",
			bytes: []byte(" "),
			want:  0,
		},
		{
			name:  "single word",
			bytes: []byte("one"),
			want:  1,
		},
		{
			name:  "multiple words",
			bytes: []byte("one two three four five"),
			want:  5,
		},
		{
			name:  "words with newlines",
			bytes: []byte("one\n\ntwo three\nfour\nfive"),
			want:  5,
		},
		{
			name:  "words with multiple spaces",
			bytes: []byte("one   two    three     four      five"),
			want:  5,
		},
		{
			name:  "words with tabs and newlines",
			bytes: []byte("one\t\two\nthree\tfour\n\nfive"),
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := counter.CountWords(tt.bytes)
			if got != tt.want {
				t.Errorf("countWords() = %d, want %d", got, tt.want)
			}
		})
	}
}
