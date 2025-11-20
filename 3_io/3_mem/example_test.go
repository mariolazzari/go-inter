package markdown

import "fmt"

func ExampleList() {
	cart := []string{"bread", "butter", "orange juice"}
	fmt.Println(List(cart))

	// Output:
	// - bread
	// - butter
	// - orange juice
}
