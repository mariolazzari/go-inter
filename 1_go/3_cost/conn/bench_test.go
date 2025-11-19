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
