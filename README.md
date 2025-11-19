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

```go

```
