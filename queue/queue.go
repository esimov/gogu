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
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.size() == 0 {
		return item, fmt.Errorf("queue is empty")
	}

	item = q.items[0]
	q.items = q.items[1:]

	return
}

// Peek returns the first element of the queue without removing it.
func (q *Queue[T]) Peek() (item T) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if q.size() == 0 {
		return
	}

	return q.items[0]
}

// Search searches for an element in the queue.
func (q *Queue[T]) Search(item T) bool {
	q.mu.RLock()
	defer q.mu.RUnlock()

	for i := 0; i < q.size(); i++ {
		if q.items[i] == item {
			return true
		}
	}

	return false
}

// Size returns the FIFO queue size.
func (q *Queue[T]) Size() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.size()
}

// size has a local scope only to avoid blocking the thread when trying to acquire the lock.
func (q *Queue[T]) size() int {
	return len(q.items)
}

// Clear erase all the items from the queue.
func (q *Queue[T]) Clear() {
	q.mu.Lock()
	q.items = nil
	q.mu.Unlock()
}
