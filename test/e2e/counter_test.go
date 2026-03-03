package e2e

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/azdanov/counter-go/test/e2e/assert"
)

func TestStdin(t *testing.T) {
	cmd, err := getCmd()
	if err != nil {
		t.Fatal("couldn't create command:", err)
	}

	cmd.Stdin = bytes.NewBufferString("Hello, World!\n")

	outBuf := &bytes.Buffer{}
	errBuf := &bytes.Buffer{}
	cmd.Stdout = outBuf
	cmd.Stderr = errBuf

	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to run the test binary: %s: %s", err, errBuf.String())
	}

	want := " 1 2 14\n"
	assert.Equal(t, outBuf.String(), want)
}

func TestSingleFile(t *testing.T) {
	file, err := createFile("Hello, World!\n")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	cmd, err := getCmd(file.Name())
	if err != nil {
		t.Fatal("couldn't create command:", err)
	}

	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("failed to run the test binary: %s", err)
	}

	want := fmt.Sprintf(" 1 2 14 %s\n", file.Name())
	assert.Equal(t, string(out), want)
}

func TestNonExistentFile(t *testing.T) {
	fileName := "nonexistent.txt"

	cmd, err := getCmd(fileName)
	if err != nil {
		t.Fatal("couldn't create command:", err)
	}

	outBuf := &bytes.Buffer{}
	errBuf := &bytes.Buffer{}
	cmd.Stdout = outBuf
	cmd.Stderr = errBuf

	err = cmd.Run()
	if err == nil {
		t.Fatal("unexpected: expected an error when trying to read a non-existent file")
	}

	wantErrMsg := "exit status 1"
	assert.Equal(t, err.Error(), wantErrMsg)

	wantStdout := ""
	assert.Equal(t, outBuf.String(), wantStdout)

	wantErr := fmt.Sprintf("%s: open %s: no such file or directory\n", binName, fileName)
	assert.Equal(t, errBuf.String(), wantErr)
}

func TestFlags(t *testing.T) {
	file, err := createFile("Hello, World!\n")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	tests := []struct {
		name     string
		flag     string
		expected string
	}{
		{"lines", "-l", fmt.Sprintf(" 1 %s\n", file.Name())},
		{"bytes", "-c", fmt.Sprintf(" 14 %s\n", file.Name())},
		{"words", "-w", fmt.Sprintf(" 2 %s\n", file.Name())},
		{"headers", "-headers", fmt.Sprintf(" lines words bytes\n     1     2    14 %s\n", file.Name())},
		{"all flags", "-l -w -c -headers", fmt.Sprintf(" lines words bytes\n     1     2    14 %s\n", file.Name())},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := append(strings.Fields(tt.flag), file.Name())
			cmd, err := getCmd(args...)
			if err != nil {
				t.Fatal("couldn't create command:", err)
			}

			out, err := cmd.Output()
			if err != nil {
				t.Fatalf("failed to run the test binary: %s", err)
			}

			assert.Equal(t, string(out), tt.expected)
		})
	}
}
