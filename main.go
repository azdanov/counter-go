package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, Counter!")

	bytes, _ := os.ReadFile("./words.txt")
	contents := string(bytes)

	fmt.Println("file:", contents)
}
