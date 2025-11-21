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

	err = Unmarshal(data, evt) // Not a pointer
	fmt.Println("err:", err)

	var s string
	err = Unmarshal(data, &s) // Unknown type
	fmt.Println("err:", err)

	// Output:
	// evt: {2024-02-11 20:27:31.275620016 +0000 UTC joe}, err: <nil>
	// err: event.LoginEvent - not a pointer
	// err: *string - unsupported type

}
