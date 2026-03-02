package e2e

import (
	"bytes"
	"fmt"
	"testing"
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
	if outBuf.String() != want {
		t.Fatalf("unexpected: got %q, want %q", outBuf.String(), want)
	}
}

func TestSingleFile(t *testing.T) {
	file, err := createFile("Hello, World!\n")
	if err != nil {
		t.Fatal(err)
	}

	cmd, err := getCmd(file.Name())
	if err != nil {
		t.Fatal("couldn't create command:", err)
	}

	outBuf := &bytes.Buffer{}
	errBuf := &bytes.Buffer{}
	cmd.Stdout = outBuf
	cmd.Stderr = errBuf

	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to run the test binary: %s: %s", err, errBuf.String())
	}

	want := fmt.Sprintf(" 1 2 14 %s\n", file.Name())
	if outBuf.String() != want {
		t.Fatalf("unexpected: got %q, want %q", outBuf.String(), want)
	}
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
	if err.Error() != wantErrMsg {
		t.Fatalf("unexpected: got %q, want %q", err.Error(), wantErrMsg)
	}

	wantStdout := ""
	if outBuf.String() != wantStdout {
		t.Fatalf("unexpected: got %q, want %q", outBuf.String(), wantStdout)
	}

	wantErr := fmt.Sprintf("%s: open %s: no such file or directory\n", binName, fileName)
	if !bytes.Contains(errBuf.Bytes(), []byte(wantErr)) {
		t.Fatalf("unexpected: got %q, want it to contain %q", errBuf.String(), wantErr)
	}
}
