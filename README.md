# counter-go

`counter-go` is a fast Go utility for counting lines, words, and bytes in files or from standard input. Similar to [wc](https://man7.org/linux/man-pages/man1/wc.1.html) command.

## Installation

Ensure you have [Go](https://go.dev/) installed. Clone the repository and build:

```bash
git clone https://github.com/azdanov/counter-go.git
cd counter-go
go build -o counter-go .
```

## Usage

```bash
./counter-go [flags] [file...]
```

If no files are provided, `counter-go` reads from standard input.

### Flags

| Flag       | Description                 |
| ---------- | --------------------------- |
| `-l`       | Show line count             |
| `-w`       | Show word count             |
| `-c`       | Show byte count             |
| `-headers` | Show header for each column |

_(If no counting flags are provided, lines, words, and bytes are displayed by default.)_

### Examples

**Count a single file:**

```bash
./counter-go file.txt
```

**Count multiple files and show column headers:**

```bash
./counter-go -headers file1.txt file2.txt
```

**Count only lines and words from standard input:**

```bash
echo "Hello World" | ./counter-go -l -w
```
