package stack

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExampleStack_MarshalJSON() {
	var s Stack
	for i := 0; i < 5; i++ {
		s.Push('a' + rune(i))
	}
	if err := json.NewEncoder(os.Stdout).Encode(s); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	// Output:
	// ["e","d","c","b","a"]

}
