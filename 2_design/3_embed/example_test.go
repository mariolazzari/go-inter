package open

import (
	"fmt"
	"io"
	"os"
)

func ExampleOpenURI() {
	url := "https://httpbin.org/base64/R28="
	file, err := OpenURI(url)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	defer file.Close()

	if _, err := io.Copy(os.Stdout, file); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	// Output:
	// Go
}
