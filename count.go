package main

import (
	"bufio"
	"io"
	"log"
	"unicode"
)

type Counts struct {
	Lines int
	Words int
	Bytes int
}

func (c Counts) Add(o Counts) Counts {
	return Counts{
		Lines: c.Lines + o.Lines,
		Words: c.Words + o.Words,
		Bytes: c.Bytes + o.Bytes,
	}
}

func Count(f io.ReadSeeker) Counts {
	c := Counts{}

	r := bufio.NewReader(f)
	inWord := false

	for {
		r, size, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln("Error reading file:", err)
		}

		if r == '\n' {
			c.Lines++
		}

		isSpace := unicode.IsSpace(r)
		if !isSpace && !inWord {
			c.Words++
		}
		inWord = !isSpace

		c.Bytes += size
	}

	return c
}
