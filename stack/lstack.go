// Linked-list implementation of the stack (LIFO) data structure.
package stack

import (
	"sync"

	"github.com/esimov/gogu/list"
)

// LStack implements the linked-list version of the LIFO stack.
type LStack[T comparable] struct {
	list *list.DList[T]
	mu   *sync.RWMutex
	n    int
}

// NewLinked creates a new LIFO stack where the items are stored in a linked-list.
func NewLinked[T comparable](t T) *LStack[T] {
	return &LStack[T]{
		list: list.InitDList(t),
		mu:   &sync.RWMutex{},
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
func (s *LStack[T]) Pop() (item T, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var t T
	node, err := s.list.Pop()
	if err != nil {
		return t, err
	}
	s.n--

	return s.list.Data(node), nil
}

// Peek returns the last element of the stack. It does not remove it.
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
