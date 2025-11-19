package wc

import (
	"fmt"
	"os"
	"strings"
)

var data = `Baa, baa, black sheep,
Have you any wool?
Yes, sir, yes, sir,
Three bags full.
`

func ExampleLineCount() {
	r := strings.NewReader(data)
	fmt.Println(LineCount(r))

	file, err := os.Open("testdata/poem.txt")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	defer file.Close()
	fmt.Println(LineCount(file))

	// Output:
	// 4 <nil>
	// 4 <nil>
}
