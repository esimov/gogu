// Linked-list implementation of the stack (LIFO) data structure.
package queue

import (
	"github.com/esimov/gogu/list"
)

// Stack implements the LIFO stack.
type LinkedStack[T comparable] struct {
	items *list.DList[T]
}

// NewStack creates a new LIFO stack where the items are stored in a plain slice.
func NewLinked[T comparable](t T) *LinkedStack[T] {
	list := list.InitDoubly(t)

	return &LinkedStack[T]{
		items: list,
	}
}

// Push appends a new element at the end of the stack.
func (s *LinkedStack[T]) Push(item T) {
	s.items.Append(item)
}

// Pop pull out the last element added to the stack.
func (s *LinkedStack[T]) Pop() (item T, err error) {
	var t T
	node, err := s.items.Pop()
	if err != nil {
		return t, err
	}
	return s.items.Data(node), nil
}

// Peek returns the last element of the stack. It does not remove it.
func (s *LinkedStack[T]) Peek() T {
	return s.items.Last()
}

// Search searches for an element in the stack.
func (s *LinkedStack[T]) Search(item T) bool {
	if _, ok := s.items.Find(item); ok {
		return true
	}

	return false
}
