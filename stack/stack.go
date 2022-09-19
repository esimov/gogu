// Package stack implements a basic LIFO (Last-In-First-Out)
// data structure using as storage a resizing array,
// where the last element added to the stack is processed first.
package stack

// Stack implements the LIFO stack.
type Stack[T comparable] struct {
	items []T
}

// New creates a new LIFO stack where the items are stored in a plain slice.
func New[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

// Push appends a new element at the end of the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop pull out the last element added to the stack.
func (s *Stack[T]) Pop() (item T) {
	if s.Size() == 0 {
		return
	}

	item = s.items[s.Size()-1]
	s.items = s.items[:s.Size()-1]

	return
}

// Peek returns the last element of the stack without removing it.
func (s *Stack[T]) Peek() (item T) {
	if s.Size() == 0 {
		return
	}
	return s.items[s.Size()-1]
}

// Search searches for an element in the stack.
func (s *Stack[T]) Search(item T) bool {
	for i := 0; i < s.Size(); i++ {
		if s.items[i] == item {
			return true
		}
	}

	return false
}

// Size returns the LIFO stack size.
func (s *Stack[T]) Size() int {
	return len(s.items)
}
