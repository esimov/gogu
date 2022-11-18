package stack

import "sync"

// Stack implements the LIFO Stack.
type Stack[T comparable] struct {
	mu    *sync.RWMutex
	items []T
}

// New creates a new LIFO stack where the items are stored in a plain slice.
func New[T comparable]() *Stack[T] {
	return &Stack[T]{
		mu: &sync.RWMutex{},
	}
}

// Push inserts a new element at the end of the stack.
func (s *Stack[T]) Push(item T) {
	s.mu.Lock()
	s.items = append(s.items, item)
	s.mu.Unlock()
}

// Pop retrieves and removes the last element pushed into the stack.
// The stack size will be decreased by one.
func (s *Stack[T]) Pop() (item T) {
	if s.Size() == 0 {
		return
	}
	len := s.Size()

	s.mu.Lock()
	item = s.items[len-1]
	s.items = s.items[:len-1]
	s.mu.Unlock()

	return
}

// Peek returns the last element of the stack without removing it.
func (s *Stack[T]) Peek() (item T) {
	len := s.Size()

	s.mu.RLock()
	defer s.mu.RUnlock()

	if len == 0 {
		return
	}
	return s.items[len-1]
}

// Search searches for an element in the stack.
func (s *Stack[T]) Search(item T) bool {
	len := s.Size()

	s.mu.RLock()
	for i := 0; i < len; i++ {
		if s.items[i] == item {
			s.mu.RUnlock()
			return true
		}
	}
	s.mu.RUnlock()

	return false
}

// Size returns the LIFO stack size.
func (s *Stack[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.items)
}
