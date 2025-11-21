package net

import "fmt"

func ExampleAddress_Format() {
	addr := Address{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("%H\n", addr) // host
	fmt.Printf("%P\n", addr) // port
	fmt.Printf("%#v\n", addr)

	// Output:
	// localhost
	// 8080
	// net.Address{Host: "localhost" Port: 8080}

}
