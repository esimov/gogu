package stack

import "sync"

// Stack implements the LIFO Stack.
type Stack[T comparable] struct {
	mu    sync.RWMutex
	items []T
}

// New creates a new LIFO stack where the items are stored in a plain slice.
func New[T comparable]() *Stack[T] {
	return &Stack[T]{
		mu: sync.RWMutex{},
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
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.size() == 0 {
		return
	}

	item = s.items[s.size()-1]
	s.items = s.items[:s.size()-1]

	return
}

// Peek returns the last element of the stack without removing it.
func (s *Stack[T]) Peek() (item T) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	len := s.size()
	if len == 0 {
		return
	}
	return s.items[len-1]
}

// Search searches for an element in the stack.
func (s *Stack[T]) Search(item T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for i := 0; i < s.size(); i++ {
		if s.items[i] == item {
			return true
		}
	}

	return false
}

// Size returns the LIFO stack size.
func (s *Stack[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.size()
}

// size has a local scope only to avoid blocking the thread when trying to acquire the lock.
func (s *Stack[T]) size() int {
	return len(s.items)
}
