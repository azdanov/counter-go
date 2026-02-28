package display_test

import (
	"bytes"
	"testing"

	"github.com/azdanov/counter-go/display"
	"github.com/azdanov/counter-go/stats"
)

func TestPrint(t *testing.T) {
	tests := []struct {
		name   string
		c      stats.Counts
		opts   display.Options
		suffix []string
		want   string
	}{
		{
			name: "show all counts with suffix",
			c: stats.Counts{
				Lines: 10,
				Words: 20,
				Bytes: 30,
			},
			opts: display.Options{
				ShowLines: false,
				ShowWords: false,
				ShowBytes: false,
			},
			suffix: []string{"file.txt"},
			want:   "10\t20\t30\t file.txt\n",
		},
		{
			name: "show only lines and words without suffix",
			c: stats.Counts{
				Lines: 5,
				Words: 15,
				Bytes: 25,
			},
			opts: display.Options{
				ShowLines: true,
				ShowWords: true,
				ShowBytes: false,
			},
			suffix: []string{},
			want:   "5\t15\t\n",
		},
		{
			name: "show only bytes with suffix",
			c: stats.Counts{
				Lines: 3,
				Words: 6,
				Bytes: 9,
			},
			opts: display.Options{
				ShowLines: false,
				ShowWords: false,
				ShowBytes: true,
			},
			suffix: []string{"data.bin"},
			want:   "9\t data.bin\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := bytes.NewBuffer(nil)
			display.Print(w, tt.c, tt.opts, tt.suffix...)
			if got := w.String(); got != tt.want {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintHeaders(t *testing.T) {
	tests := []struct {
		name string
		opts display.Options
		want string
	}{
		{
			name: "show all headers",
			opts: display.Options{
				ShowLines:  true,
				ShowWords:  true,
				ShowBytes:  true,
				ShowHeader: true,
			},
			want: "lines\twords\tbytes\t\n",
		},
		{
			name: "show only lines and words headers",
			opts: display.Options{
				ShowLines:  true,
				ShowWords:  true,
				ShowBytes:  false,
				ShowHeader: true,
			},
			want: "lines\twords\t\n",
		},
		{
			name: "show only bytes header",
			opts: display.Options{
				ShowLines:  false,
				ShowWords:  false,
				ShowBytes:  true,
				ShowHeader: true,
			},
			want: "bytes\t\n",
		},
		{
			name: "do not show headers",
			opts: display.Options{
				ShowLines:  true,
				ShowWords:  true,
				ShowBytes:  true,
				ShowHeader: false,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := bytes.NewBuffer(nil)
			display.PrintHeaders(w, tt.opts)
			if got := w.String(); got != tt.want {
				t.Errorf("PrintHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}
