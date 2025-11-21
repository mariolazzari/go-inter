package net

import "fmt"

type Address struct {
	Host string
	Port int
}

// Format implements fmt.Formatter
func (a Address) Format(f fmt.State, verb rune) {
	switch verb {
	case 'H':
		fmt.Fprintf(f, a.Host)
		return
	case 'P':
		fmt.Fprintf(f, "%d", a.Port)
		return
	case 'v':
		switch {
		case f.Flag('+'):
			fmt.Fprintf(f, "{Host: %s Port: %d}", a.Host, a.Port)
			return
		case f.Flag('#'):
			fmt.Fprintf(f, "%T{Host: %q Port: %d}", a, a.Host, a.Port)
			return
		}
	}

	fmt.Printf("{%s %d}", a.Host, a.Port)
}
