package event

import (
	"fmt"
)

func ExampleDecoder_Decode() {
	data := []byte(`{
		"Time": "2024-02-11T20:27:31.275620016Z",
		"Login": "joe"
	}`)

	var evt LoginEvent
	err := Unmarshal(data, &evt)
	fmt.Printf("evt: %v, err: %v\n", evt, err)

	// These won't compile
	// Unmarshal(data, evt)
	// var s string
	// err = Unmarshal(data, &s) // Unknown type

	// Output:
	// evt: {2024-02-11 20:27:31.275620016 +0000 UTC joe}, err: <nil>
}
