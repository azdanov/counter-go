package e2e

import (
	"bufio"
	"bytes"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
	"testing"
)

func TestMultipleFiles(t *testing.T) {
	fileA, err := createFile("Hello, World!\n")
	if err != nil {
		t.Fatal("couldn't create file A:", err)
	}
	defer os.Remove(fileA.Name())

	fileB, err := createFile("Hello, World!!\n")
	if err != nil {
		t.Fatal("couldn't create file B:", err)
	}
	defer os.Remove(fileB.Name())

	fileC, err := createFile("Hello, World!!!\n")
	if err != nil {
		t.Fatal("couldn't create file C:", err)
	}
	defer os.Remove(fileC.Name())

	cmd, err := getCmd(fileA.Name(), fileB.Name(), fileC.Name())
	if err != nil {
		t.Fatal("couldn't create command:", err)
	}

	stdoutBuf := &bytes.Buffer{}
	cmd.Stdout = stdoutBuf
	stderrBuf := &bytes.Buffer{}
	cmd.Stderr = stderrBuf

	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to run the test binary: %s: %s", err, stderrBuf.String())
	}

	wants := map[string]string{
		fileA.Name(): fmt.Sprintf(" 1 2 14 %s", fileA.Name()),
		fileB.Name(): fmt.Sprintf(" 1 2 15 %s", fileB.Name()),
		fileC.Name(): fmt.Sprintf(" 1 2 16 %s", fileC.Name()),
		"total":      " 3 6 45 total",
	}
	scanner := bufio.NewScanner(stdoutBuf)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) == 0 {
			t.Fatalf("unexpected empty line")
		}

		lastField := fields[len(fields)-1]
		want, ok := wants[lastField]
		if !ok {
			t.Fatalf("unexpected last field: got %q, want one of %q", lastField, slices.Collect(maps.Keys(wants)))
		}

		if line != want {
			t.Fatalf("unexpected line: got %q, want %q", line, want)
		}
	}
}
