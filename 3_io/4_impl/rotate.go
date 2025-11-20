package rotate

import (
	"fmt"
	"os"
	"path"
)

type Rotator struct {
	rootPath string
	n        int
	maxSize  int
	size     int
	out      *os.File
}

func New(rootPath string, maxSize int) (*Rotator, error) {
	if err := os.MkdirAll(rootPath, 0700); err != nil {
		return nil, err
	}

	r := Rotator{
		rootPath: rootPath,
		maxSize:  maxSize,
	}
	if err := r.rotate(); err != nil {
		return nil, err
	}

	return &r, nil
}

func (r *Rotator) Write(data []byte) (int, error) {
	if n, err := r.out.Write(data); err != nil {
		return n, err
	}

	r.size += len(data)
	if r.size > r.maxSize {
		if err := r.rotate(); err != nil {
			return len(data), err
		}
	}

	return len(data), nil
}

func (r *Rotator) Close() error {
	if r.out == nil {
		return nil
	}

	return r.out.Close()
}

func (r *Rotator) rotate() error {
	if r.out != nil {
		r.out.Close()
	}

	r.n++
	fileName := path.Join(r.rootPath, fmt.Sprintf("log-%02d.txt", r.n))
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	r.size = 0
	r.out = file
	return nil
}
