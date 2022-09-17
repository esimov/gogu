// Package stack implements a basic LIFO (Last-In-First-Out)
// data structure using as storage a resizing array,
// where the last element added to the stack is processed first.
package stack

// Stack implements the LIFO stack.
type Stack[T comparable] struct {
	Items []T
}

// NewStack creates a new LIFO stack where the items are stored in a plain slice.
func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

// Push appends a new element at the end of the stack.
func (s *Stack[T]) Push(item T) {
	s.Items = append(s.Items, item)
}

// Pop pull out the last element added to the stack.
func (s *Stack[T]) Pop() (item T) {
	if s.Size() == 0 {
		return
	}

	item = s.Items[s.Size()-1]
	s.Items = s.Items[:s.Size()-1]

	return
}

// Peek returns the last element of the stack. It does not remove it.
func (s *Stack[T]) Peek() (item T) {
	if s.Size() == 0 {
		return
	}
	return s.Items[s.Size()-1]
}

// Search searches for an element in the stack.
func (s *Stack[T]) Search(item T) bool {
	for i := 0; i < s.Size(); i++ {
		if s.Items[i] == item {
			return true
		}
	}

	return false
}

// Size returns the LIFO stack size.
func (s *Stack[T]) Size() int {
	return len(s.Items)
}
