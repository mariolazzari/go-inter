package event

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type LoginEvent struct {
	Time  time.Time
	Login string
}

type AccessEvent struct {
	Time   time.Time
	Login  string
	URI    string
	Action string
}

func Unmarshal(data []byte, v any) error {
	if reflect.TypeOf(v).Kind() != reflect.Pointer {
		return fmt.Errorf("%T - not a pointer", v)
	}

	switch v.(type) {
	case *AccessEvent, *LoginEvent:
		// OK
	default:
		return fmt.Errorf("%T - unsupported type", v)
	}

	return json.Unmarshal(data, v)
}
