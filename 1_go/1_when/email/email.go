package email

import "fmt"

type Email struct {
	Name    string
	Address string
}

// String implements fmt.Stringer
func (e Email) String() string {
	return fmt.Sprintf("%s <%s>", e.Name, e.Address)
}
