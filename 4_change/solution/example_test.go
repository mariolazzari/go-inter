package stacked

import "fmt"

func ExampleError() {
	a := func() error {
		return Wrap(fmt.Errorf("oops"))
	}

	b := func() error {
		return a()
	}

	err := b()
	fmt.Println(err)
	fmt.Println("---------------")
	fmt.Printf("%+v\n", err)

	// Output:
	// oops
	// ---------------
	// oops
	//
	//goiface/Ch04/solution.ExampleError.func1
	//	example_test.go:7
	//goiface/Ch04/solution.ExampleError.func2
	//	example_test.go:11
	// goiface/Ch04/solution.ExampleError
	//	example_test.go:14

}
