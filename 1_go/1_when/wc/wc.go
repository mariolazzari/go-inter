package wc

import (
	"bufio"
	"io"
)

// LineCount return how many lines in r.
func LineCount(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	count := 0
	for s.Scan() {
		count++
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return count, nil
}
