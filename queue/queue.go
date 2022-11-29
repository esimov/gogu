// Package queue implements a concurrent safe FIFO (First-In-First-Out)
// data structure where the first element added to the queue is processed first.
// It's implemented in two versions:
//
// 1.) where the storage system is a resizing array,
//
// 2.) where the storage system is a doubly linked list.
package queue

import (
	"fmt"
	"sync"
)

// Queue implements a FIFO Queue data structure.
type Queue[T comparable] struct {
	mu    sync.RWMutex
	items []T
}

// New creates a new FIFO queue where the items are stored in a plain slice.
func New[T comparable]() *Queue[T] {
	return &Queue[T]{
		mu: sync.RWMutex{},
	}
}

// Enqueue inserts a new element at the end of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.mu.Lock()
	q.items = append(q.items, item)
	q.mu.Unlock()
}

// Dequeue retrieves and removes the first element from the queue.
// The queue size will be decreased by one.
func (q *Queue[T]) Dequeue() (item T, err error) {
	if q.Size() == 0 {
		return item, fmt.Errorf("queue is empty")
	}

	q.mu.Lock()
	item = q.items[0]
	q.items = q.items[1:]
	q.mu.Unlock()

	return
}

// Peek returns the first element of the queue without removing it.
func (q *Queue[T]) Peek() (item T) {
	len := q.Size()

	q.mu.RLock()
	defer q.mu.RUnlock()

	if len == 0 {
		return
	}
	return q.items[0]
}

// Search searches for an element in the queue.
func (q *Queue[T]) Search(item T) bool {
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
func (q *Queue[T]) Size() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return len(q.items)
}

// Clear erase all the items from the queue.
func (q *Queue[T]) Clear() {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.items = nil
}
