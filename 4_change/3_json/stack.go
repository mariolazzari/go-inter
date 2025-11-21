package stack

import (
	"encoding/json"
	"errors"
)

type node struct {
	value rune
	next  *node
}

type Stack struct {
	head *node
}

// Push pushes a value to the stack.
func (s *Stack) Push(r rune) {
	s.head = &node{r, s.head}
}

var ErrEmpty = errors.New("empty stack")

// Pop pops an value from the stack.
func (s *Stack) Pop() (rune, error) {
	if s.head == nil {
		return 0, ErrEmpty
	}

	v := s.head.value
	s.head = s.head.next
	return v, nil
}

// Len returns the number of elements in the stack.
func (s *Stack) Len() int {
	count := 0

	for n := s.head; n != nil; n = n.next {
		count++
	}

	return count
}

// MarshalJSON implements json.Marshaler
func (s Stack) MarshalJSON() ([]byte, error) {
	values := make([]string, s.Len())

	i := 0
	for n := s.head; n != nil; n = n.next {
		values[i] = string(n.value)
		i++
	}

	return json.Marshal(values)
}
