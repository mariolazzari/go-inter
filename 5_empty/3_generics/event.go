package event

import (
	"encoding/json"
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

type Event interface {
	*AccessEvent | *LoginEvent
}

func Unmarshal[T Event](data []byte, v T) error {
	return json.Unmarshal(data, v)
}
