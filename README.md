# Deep Dive Into Go Lang Interfaces

## Interfaces in Go

### When should you use interfaces

```go
package email

import "fmt"

type Email struct {
	Name    string
	Address string
}

// String implements fmt.Stringer
func (e Email) String() string {
	return fmt.Sprintf("%s <%s>", e.Name, e.Address)
}
```

### Interface implementation

```go
package db

import (
	"fmt"
)

type DBError struct {
	Address string
	Reason  string
}

func (e *DBError) Error() string {
	return fmt.Sprintf("%s (address=%q)", e.Reason, e.Address)
}

type DB struct{}

func Open(host string) (*DB, error) {
	var err *DBError

	// TODO: Connect
	return nil, err
}

```

### Cost of interfaces

```go
package conn

import (
	"io"
	"testing"
)

func BenchmarkMethod(b *testing.B) {
	var c Conn
	for i := 0; i < b.N; i++ {
		err := c.Close()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIface(b *testing.B) {
	var c io.Closer = &Conn{}
	for i := 0; i < b.N; i++ {
		err := c.Close()
		if err != nil {
			b.Fatal(err)
		}
	}
}
```

## Iterface design

### Say what you need, not what you provide

[copy](https://pkg.go.dev/builtin#copy)

- Small: max 5 methods
- Accept interfaces, return types

### Sort interface

[Sort](https://pkg.go.dev/sort#Interface)

```go
package sort

type Sortable interface {
	Less(i, j int) bool
	Swao(i, j int)
	Len() int
}

func Sort(s Sortable) {
	// TODO: sort
}
```

### Embedding interfaces

[ReadClosers](https://pkg.go.dev/storj.io/common/readcloser)

```go
package open

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// OpenURI opens a URI.
// Supported schemes are: file, http & https.
func OpenURI(uri string) (io.ReadCloser, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "http", "https":
		resp, err := http.Get(uri)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("%q: bad status - %s", uri, resp.Status)
		}
		return resp.Body, nil
	case "file":
		file, err := os.Open(u.Path)
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	return nil, fmt.Errorf("unknown scheme: %s", u.Scheme)
}
```

### Type assertions

[Proverb](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=5m17s&themeRefresh=1)

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Event struct {
	Time    time.Time `json:"time,omitempty"`
	Message string    `json:"message,omitempty"`
}

type syncer interface {
	Sync() error
}

type Encoder struct {
	w io.Writer
	s syncer
}

type nosincer struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	if s, ok := w.(syncer); ok {
		return &Encoder{w: w, s: s}
	}

	e := Encoder{w: w}
	return &e
}

func (e *Encoder) Encode(evt Event) error {
	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	n, err := e.w.Write(data)
	if err != nil {
		return err
	}

	if n != len(data) {
		return fmt.Errorf("partial write (%d out of %d bytes)", n, len(data))
	}

	if s, ok := e.w.(syncer); ok {
		s.Sync()
	}

	return nil
}

func main() {
	enc := NewEncoder(os.Stdout)
	evt := Event{
		Time:    time.Now().UTC(),
		Message: "elliot login",
	}
	enc.Encode(evt)
}
```

### Design for performance

[Reader](https://pkg.go.dev/io#Reader)

```go
package main

import "fmt"

type User struct {
	Login string
	// TODO: More fields
}

type UserIter interface {
	Next(*User) bool
}

type Query struct {
	n int
}

func (q *Query) Next(u *User) bool {
	if q.n == 5 {
		return false
	}

	q.n++
	u.Login = fmt.Sprintf("user-%d", q.n)

	return true
}

func PrintUsers(ui UserIter) {
	var u User
	for ui.Next(&u) {
		fmt.Println(u)
	}
}

func main() {
	var q Query
	PrintUsers(&q)
}
```

### Challenge: puller

```go
package main

type Record struct {
	Key  uint
	Data []byte
}

// Puller pulls a single record from the database.
type Puller interface {
	Pull(r *Record) error
}
```

## I/O interfaces

### Reader and Writer

[Reader](https://pkg.go.dev/io#Reader)
[Writer](https://pkg.go.dev/io#Writer)
[Copy](https://pkg.go.dev/io#example-Copy)

### Composing reader and writer

```go
package main

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var data = []byte(`
“Hope” is the thing with feathers -
That perches in the soul -
And sings the tune without the words -
And never stops - at all -
`)

func poemHandler(w http.ResponseWriter, r *http.Request) {
	var wtr io.Writer = w
	accept := r.Header.Get("Accept-Encoding")
	if strings.Contains(accept, "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		wtr = gz
	}

	n, err := wtr.Write(data)
	if err != nil || n < len(data) {
		log.Printf("ERROR: bad write size=%d, written=%d, error=%s", len(data), n, err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/poem", poemHandler)

	addr := ":8080"
	log.Printf("INFO: server starting on %s", addr)
	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("ERROR: can't run - %s", err)
		os.Exit(1)
	}
}
```

```sh
curl -H 'Accept-Encoding: gzip' http://localhost:8080/poem | gunzip
```

### In memory readers and writes

[Buffer](https://pkg.go.dev/bytes#example-Buffer)

```go
package markdown

import (
	"bytes"
	"fmt"
)

// List renders a slice of item to a markdown list.
func List(items []string) string {
	var buf bytes.Buffer

	for _, item := range items {
		fmt.Fprintf(&buf, "- %s\n", item)
	}
	return buf.String()
}
```

### Implementing reader and writer

```go
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
```

### Challenge: counting lines

```go
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
```

## Interfaces that change behavior

### String representation

[Stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
[Gostringer](https://pkg.go.dev/github.com/sourcegraph/gostringer)

```go
package auth

import "fmt"

type Permission byte

const (
	Read Permission = iota + 1
	Write
	Admin
)

// String implements fmt.Stringer
func (p Permission) String() string {
	switch p {
	case Read:
		return "read"
	case Write:
		return "write"
	case Admin:
		return "admin"
	}

	return fmt.Sprintf("<Permission: %d>", p)
}
```

### Formatter print flags

[format](https://pkg.go.dev/fmt#example-package-Formats)

```go
package net

import "fmt"

type Address struct {
	Host string
	Port int
}

// Format implements fmt.Formatter
func (a Address) Format(f fmt.State, verb rune) {
	switch verb {
	case 'H':
		fmt.Fprintf(f, a.Host)
		return
	case 'P':
		fmt.Fprintf(f, "%d", a.Port)
		return
	case 'v':
		switch {
		case f.Flag('+'):
			fmt.Fprintf(f, "{Host: %s Port: %d}", a.Host, a.Port)
			return
		case f.Flag('#'):
			fmt.Fprintf(f, "%T{Host: %q Port: %d}", a, a.Host, a.Port)
			return
		}
	}

	fmt.Printf("{%s %d}", a.Host, a.Port)
}
```

### Marshaler and Unmarshaler

[Marshaler](https://pkg.go.dev/encoding/json#Marshaler)
[Unmarshaler](https://pkg.go.dev/encoding/json#Unmarshaler)

```go
package stack

import (
	"encoding/json"
	"errors"
)

type node struct {
	value rune
	next  *node
}

type Stack struct {
	head *node
}

// Push pushes a value to the stack.
func (s *Stack) Push(r rune) {
	s.head = &node{r, s.head}
}

var ErrEmpty = errors.New("empty stack")

// Pop pops an value from the stack.
func (s *Stack) Pop() (rune, error) {
	if s.head == nil {
		return 0, ErrEmpty
	}

	v := s.head.value
	s.head = s.head.next
	return v, nil
}

// Len returns the number of elements in the stack.
func (s *Stack) Len() int {
	count := 0

	for n := s.head; n != nil; n = n.next {
		count++
	}

	return count
}

// MarshalJSON implements json.Marshaler
func (s Stack) MarshalJSON() ([]byte, error) {
	values := make([]string, s.Len())

	i := 0
	for n := s.head; n != nil; n = n.next {
		values[i] = string(n.value)
		i++
	}

	return json.Marshal(values)
}
```

### Challenge: custom error

```go
package stacked

import (
	"bytes"
	"fmt"
	"path"
	"runtime"
)

type Error struct {
	cause error
	stack string
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s\n\n", err)
	fmt.Fprintf(&buf, callStack())

	s := Error{
		cause: err,
		stack: buf.String(),
	}

	return &s
}

func callStack() string {
	var buf bytes.Buffer
	pcs := make([]uintptr, 20)

	n := runtime.Callers(3, pcs)
	if n > 0 {
		frames := runtime.CallersFrames(pcs[:n])

		for {
			fr, more := frames.Next()
			if fr.Function == "runtime.main" || fr.Function == "testing.runExample" {
				break
			}

			fmt.Fprintln(&buf, fr.Function)
			fmt.Fprintf(&buf, "\t%s:%d\n", path.Base(fr.File), fr.Line)

			if !more {
				break
			}
		}
	}

	return buf.String()
}

func (e *Error) Error() string {
	return e.cause.Error()
}

func (e *Error) Format(f fmt.State, verb rune) {
	if verb == 'v' && f.Flag('+') {
		fmt.Fprint(f, e.stack)
		return
	}

	fmt.Fprint(f, e.Error())
}
```
