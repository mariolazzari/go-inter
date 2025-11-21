package auth

import "fmt"

func ExamplePermission_String() {
	p := Write
	fmt.Println(p)
	fmt.Printf(" v: %v\n", p)
	fmt.Printf("+v: %+v\n", p)
	fmt.Printf("#v: %#v\n", p)

	// Output:
	// write
	//  v: write
	// +v: write
	// #v: 0x2
}
