package e2e

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/azdanov/counter-go/test/e2e/assert"
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

	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("failed to run the test binary: %s", err)
	}

	want := fmt.Sprintf(
		"   1   2   14 %s\n 150 300 2250 %s\n  32  32  224 %s\n 183 334 2488 total\n",
		fileA.Name(),
		fileB.Name(),
		fileC.Name(),
	)
	assert.Equal(t, string(out), want)
}
