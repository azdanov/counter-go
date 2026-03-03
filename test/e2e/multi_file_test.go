package e2e

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMultipleFiles(t *testing.T) {
	fileA, err := createFile("Hello, World!\n")
	if err != nil {
		t.Fatal("couldn't create file A:", err)
	}
	defer os.Remove(fileA.Name())

	fileB, err := createFile(strings.Repeat("Hello,\tWorld!!\n", 150))
	if err != nil {
		t.Fatal("couldn't create file B:", err)
	}
	defer os.Remove(fileB.Name())

	fileC, err := createFile(strings.Repeat("Hello\nWorld!!\n", 16))
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

	want := fmt.Sprintf(
		"   1   2   14 %s\n 150 300 2250 %s\n  32  32  224 %s\n 183 334 2488 total\n",
		fileA.Name(),
		fileB.Name(),
		fileC.Name(),
	)

	got := stdoutBuf.String()
	if got != want {
		t.Fatalf("unexpected: got %q, want %q", got, want)
	}
}
