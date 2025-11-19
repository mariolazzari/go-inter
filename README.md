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
