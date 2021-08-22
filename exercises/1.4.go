package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fd := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fd)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fd)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\t%s\n", fd[line], n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, fd map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		fd[input.Text()] = f.Name()
		counts[input.Text()]++
	}
}
