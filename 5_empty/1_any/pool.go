package main

import (
	"fmt"
	"sync"
)

func main() {
	var pool sync.Pool

	for i := 0; i < 5; i++ {
		buf := make([]byte, 1024)
		pool.Put(buf)
	}

	out := pool.Get()
	buf, ok := out.([]byte)
	if !ok {
		fmt.Printf("ERROR: not a []byte - %#v (%T)\n", out, out)
		return
	}

	fmt.Println(len(buf))
}
