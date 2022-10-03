// Package queue implements a basic FIFO (First-In-First-Out)
// data structure using as storage a resizing array,
// where the first element added to the queue is processed first.
package queue

import "sync"

// queue implements a FIFO queue data structure.
type queue[T comparable] struct {
	mu    *sync.RWMutex
	items []T
}

// New creates a new FIFO queue where the items are stored in a plain slice.
func New[T comparable]() *queue[T] {
	return &queue[T]{
		mu: &sync.RWMutex{},
	}
}

// Enqueue inserts a new element at the end of the queue.
func (q *queue[T]) Enqueue(item T) {
	q.mu.Lock()
	q.items = append(q.items, item)
	q.mu.Unlock()
}

// Dequeue retrieves and removes the first element from the queue.
// The queue size will be decreased by one.
func (q *queue[T]) Dequeue() (item T) {
	if q.Size() == 0 {
		return
	}

	q.mu.Lock()
	item = q.items[0]
	q.items = q.items[1:]
	q.mu.Unlock()

	return
}

// Peek returns the first element of the queue without removing it.
func (q *queue[T]) Peek() (item T) {
	len := q.Size()

	q.mu.RLock()
	defer q.mu.RUnlock()

	if len == 0 {
		return
	}
	return q.items[0]
}

// Search searches for an element in the queue.
func (q *queue[T]) Search(item T) bool {
	len := q.Size()

	q.mu.RLock()
	for i := 0; i < len; i++ {
		if q.items[i] == item {
			q.mu.RUnlock()
			return true
		}
	}
	q.mu.RUnlock()

	return false
}

// Size returns the FIFO queue size.
func (q *queue[T]) Size() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return len(q.items)
}
