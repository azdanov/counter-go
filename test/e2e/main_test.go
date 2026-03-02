package e2e

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var binName = "counter-test"

func TestMain(m *testing.M) {
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	cmd := exec.Command("go", "build", "-o", binName, "../..")

	errBuf := &bytes.Buffer{}
	cmd.Stderr = errBuf

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to build the test binary: %s: %s", err, errBuf.String())
		os.Exit(1)
	}

	result := m.Run()

	os.Remove(binName)
	os.Exit(result)
}

func getCmd(args ...string) (*exec.Cmd, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(dir, binName)
	cmd := exec.Command(path, args...)

	return cmd, nil
}

func createFile(content string) (*os.File, error) {
	file, err := os.CreateTemp("", "counter-test-*.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't create temp file: %w", err)
	}

	if _, err := file.WriteString(content); err != nil {
		defer os.Remove(file.Name())
		defer file.Close()
		return nil, fmt.Errorf("couldn't write to temp file: %w", err)
	}

	if err := file.Close(); err != nil {
		defer os.Remove(file.Name())
		return nil, fmt.Errorf("couldn't close temp file: %w", err)
	}

	return file, nil
}
