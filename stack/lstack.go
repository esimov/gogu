// Package stack Package queue implements a concurrent safe LIFO (Last-In-First-Out)
// data structure where the last element added to the stack is processed first.
// It's implemented in two versions:
// 1.) where the storage system is a resizing array,
// 2.) where the storage system is a doubly linked list.
package stack

import (
	"sync"

	"github.com/esimov/gogu/list"
)

// LStack implements the linked-list version of the LIFO stack.
type LStack[T comparable] struct {
	list *list.DList[T]
	mu   sync.RWMutex
	n    int
}

// NewLinked creates a new LIFO stack where the items are stored in a linked-list.
func NewLinked[T comparable](t T) *LStack[T] {
	return &LStack[T]{
		list: list.InitDList(t),
		mu:   sync.RWMutex{},
		n:    1,
	}
}

// Push inserts a new element at the end of the stack.
func (s *LStack[T]) Push(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.n++
	s.list.Append(item)
}

// Pop retrieves and removes the last element pushed into the stack.
// The stack size will be decreased by one.
func (s *LStack[T]) Pop() (item T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	node := s.list.Pop()
	if s.n > 0 {
		s.n--
	}

	return s.list.Data(node)
}

// Peek returns the last element of the stack without removing it.
func (s *LStack[T]) Peek() T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.list.Last()
}

// Search searches for an element in the stack.
func (s *LStack[T]) Search(item T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, ok := s.list.Find(item); ok {
		return true
	}

	return false
}

// Size returns the stack size.
func (s *LStack[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.n
}
