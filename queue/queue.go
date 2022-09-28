// Package queue implements a basic FIFO (First-In-First-Out)
// data structure using as storage a resizing array,
// where the first element added to the stack is processed first.
package queue

import "sync"

// Queue implements the FIFO stack.
type Queue[T comparable] struct {
	mu    *sync.RWMutex
	items []T
}

// New creates a new FIFO stack where the items are stored in a plain slice.
func New[T comparable]() *Queue[T] {
	return &Queue[T]{
		mu: &sync.RWMutex{},
	}
}

// Enqueue inserts a new element at the end of the queue.
func (s *Queue[T]) Enqueue(item T) {
	s.mu.Lock()
	s.items = append(s.items, item)
	s.mu.Unlock()
}

// Dequeue retrieves and removes the first element from the queue.
// The queue size will be decreased by one.
func (s *Queue[T]) Dequeue() (item T) {
	if s.Size() == 0 {
		return
	}

	s.mu.Lock()
	item = s.items[0]
	s.items = s.items[1:]
	s.mu.Unlock()

	return
}

// Peek returns the first element of the queue without removing it.
func (s *Queue[T]) Peek() (item T) {
	len := s.Size()

	s.mu.RLock()
	defer s.mu.RUnlock()

	if len == 0 {
		return
	}
	return s.items[0]
}

// Search searches for an element in the queue.
func (s *Queue[T]) Search(item T) bool {
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

// Size returns the FIFO stack size.
func (s *Queue[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.items)
}
