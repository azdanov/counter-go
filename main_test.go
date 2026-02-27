package main_test

import (
	"bytes"
	"testing"

	"github.com/azdanov/counter-go"
)

func TestPrint(t *testing.T) {
	tests := []struct {
		name   string
		c      main.Counts
		opts   main.DisplayOptions
		suffix []string
		want   string
	}{
		{
			name: "show all counts with suffix",
			c: main.Counts{
				Lines: 10,
				Words: 20,
				Bytes: 30,
			},
			opts: main.DisplayOptions{
				ShowLines: false,
				ShowWords: false,
				ShowBytes: false,
			},
			suffix: []string{"file.txt"},
			want:   "10 20 30 file.txt\n",
		},
		{
			name: "show only lines and words without suffix",
			c: main.Counts{
				Lines: 5,
				Words: 15,
				Bytes: 25,
			},
			opts: main.DisplayOptions{
				ShowLines: true,
				ShowWords: true,
				ShowBytes: false,
			},
			suffix: []string{},
			want:   "5 15\n",
		},
		{
			name: "show only bytes with suffix",
			c: main.Counts{
				Lines: 3,
				Words: 6,
				Bytes: 9,
			},
			opts: main.DisplayOptions{
				ShowLines: false,
				ShowWords: false,
				ShowBytes: true,
			},
			suffix: []string{"data.bin"},
			want:   "9 data.bin\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := bytes.NewBuffer(nil)
			main.Print(w, tt.c, tt.opts, tt.suffix...)
			if got := w.String(); got != tt.want {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}
