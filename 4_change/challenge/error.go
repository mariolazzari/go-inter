package stacked

import (
	"bytes"
	"fmt"
	"path"
	"runtime"
)

type Error struct {
	cause error
	stack string
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s\n\n", err)
	fmt.Fprintf(&buf, callStack())

	s := Error{
		cause: err,
		stack: buf.String(),
	}

	return &s
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

func (e *Error) Error() string {
	return e.cause.Error()
}

func (e *Error) Format(f fmt.State, verb rune) {
	if verb == 'v' && f.Flag('+') {
		fmt.Fprint(f, e.stack)
		return
	}

	fmt.Fprint(f, e.Error())
}
