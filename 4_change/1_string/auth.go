package auth

import "fmt"

type Permission byte

const (
	Read Permission = iota + 1
	Write
	Admin
)

// String implements fmt.Stringer
func (p Permission) String() string {
	switch p {
	case Read:
		return "read"
	case Write:
		return "write"
	case Admin:
		return "admin"
	}

	return fmt.Sprintf("<Permission: %d>", p)
}
