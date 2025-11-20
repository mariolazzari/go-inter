/*
	Count how many lines are in roads.txt.gz

You should do the Go equivalent of

	$ cat roads.txt.gz| gunzip | wc -l

You *must* use io.Copy in your code.
*/
package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

type LineCount struct {
	n int
}

func (l *LineCount) Write(data []byte) (int, error) {
	for _, c := range data {
		if c == '\n' {
			l.n++
		}
	}

	return len(data), nil
}

func (l *LineCount) Len() int {
	return l.n
}

func main() {
	const fileName = "roads.txt.gz"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	r, err := gzip.NewReader(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %q - %s\n", fileName, err)
		os.Exit(1)
	}

	var w LineCount
	if _, err := io.Copy(&w, r); err != nil {
		fmt.Fprintf(os.Stderr, "error: %q - %s\n", fileName, err)
		os.Exit(1)
	}

	fmt.Println(w.Len())
}
