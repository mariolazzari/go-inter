package email

import "fmt"

func ExampleEmail_String() {
	e := Email{"Frodo Baggins", "frodo@shire.org"}
	fmt.Println(e)

	// Output:
	// Frodo Baggins <frodo@shire.org>
}
