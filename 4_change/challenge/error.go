package stacked

import (
	"bytes"
	"fmt"
	"path"
	"runtime"
)

type Error struct {
	// TODO
}

func Wrap(err error) error {
	// TODO
	return nil
}

func (e *Error) Error() string {
	// TODO
	return ""
}

func (e *Error) Format(f fmt.State, verb rune) {
	// TODO
}

func callStack() string {
	var buf bytes.Buffer
	pcs := make([]uintptr, 20)

	n := runtime.Callers(3, pcs)
	if n > 0 {
		frames := runtime.CallersFrames(pcs[:n])

		for {
			fr, more := frames.Next()
			if fr.Function == "runtime.main" || fr.Function == "testing.runExample" {
				break
			}

			fmt.Fprintln(&buf, fr.Function)
			fmt.Fprintf(&buf, "\t%s:%d\n", path.Base(fr.File), fr.Line)

			if !more {
				break
			}
		}
	}

	return buf.String()
}
