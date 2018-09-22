package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	seen := make(map[string]int)

	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		line := strings.TrimSpace(input.Text())
		line = strings.TrimPrefix(line, "\n")
		seen[line]++
	}

	for k, v := range seen {
		fmt.Printf("%s %d\n", k, v)
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}

}
