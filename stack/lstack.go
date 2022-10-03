// Linked-list implementation of the stack (LIFO) data structure.
package stack

import (
	"github.com/esimov/gogu/list"
)

// lStack implements the linked-list version of the LIFO stack.
type lStack[T comparable] struct {
	items *list.DList[T]
}

// NewLinked creates a new LIFO stack where the items are stored in a linked-list.
func NewLinked[T comparable](t T) *lStack[T] {
	list := list.InitDoubly(t)

	return &lStack[T]{
		items: list,
	}
}

// Push inserts a new element at the end of the stack.
func (s *lStack[T]) Push(item T) {
	s.items.Append(item)
}

// Pop retrieves and removes the last element pushed into the stack.
// The stack size will be decreased by one.
func (s *lStack[T]) Pop() (item T, err error) {
	var t T
	node, err := s.items.Pop()
	if err != nil {
		return t, err
	}
	return s.items.Data(node), nil
}

// Peek returns the last element of the stack. It does not remove it.
func (s *lStack[T]) Peek() T {
	return s.items.Last()
}

// Search searches for an element in the stack.
func (s *lStack[T]) Search(item T) bool {
	if _, ok := s.items.Find(item); ok {
		return true
	}

	return false
}
