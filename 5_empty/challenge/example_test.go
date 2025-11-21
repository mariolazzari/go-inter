package pool

import "fmt"

func ExamplePool() {
	p := New[[]byte](nil)
	n := 3
	for i := 0; i < n; i++ {
		p.Put(make([]byte, 1024))
	}

	for i := 0; i < n; i++ {
		b, ok := p.Get()
		fmt.Printf("len: %d, ok=%v\n", len(b), ok)
	}

	// Empty pool
	b, ok := p.Get()
	fmt.Printf("len: %d, ok=%v\n", len(b), ok)

	// Output:
	// len: 1024, ok=true
	// len: 1024, ok=true
	// len: 1024, ok=true
	// len: 0, ok=false
}
