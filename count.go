package main

import (
	"bufio"
	"io"
	"log"
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
	lines := CountLines(f)
	f.Seek(0, io.SeekStart)
	words := CountWords(f)
	f.Seek(0, io.SeekStart)
	bytes := CountBytes(f)

	return Counts{
		Lines: lines,
		Words: words,
		Bytes: bytes,
	}
}

func CountLines(r io.Reader) int {
	reader := bufio.NewReader(r)

	count := 0
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln("Error reading file:", err)
		}
		if r == '\n' {
			count++
		}
	}

	return count
}

func CountWords(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Error scanning file:", err)
	}

	return count
}

func CountBytes(r io.Reader) int {
	count, err := io.Copy(io.Discard, r)
	if err != nil {
		log.Fatalln("Error counting bytes:", err)
	}
	return int(count)
}
