// Linked-list implementation of the queue (FIFO) data structure.
package queue

import (
	"sync"

	"github.com/esimov/gogu/list"
)

// LQueue implements the linked-list version of the FIFO queue.
type LQueue[T comparable] struct {
	list *list.DList[T]
	mu   *sync.RWMutex
	n    int
}

// NewLinked creates a new FIFO queue where the items are stored in a linked-list.
func NewLinked[T comparable](t T) *LQueue[T] {
	return &LQueue[T]{
		list: list.InitDList(t),
		mu:   &sync.RWMutex{},
		n:    1,
	}
}

// Enqueue inserts a new element at the end of the queue.
func (l *LQueue[T]) Enqueue(item T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.n++
	l.list.Append(item)
}

// Dequeue retrieves and removes the first element from the queue.
// The queue size will be decreased by one.
func (l *LQueue[T]) Dequeue() (item T, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var t T
	node, err := l.list.Shift()
	if err != nil {
		return t, err
	}
	l.n--
	return l.list.Data(node), nil
}

// Peek returns the first element of the queue. It does not remove it.
func (l *LQueue[T]) Peek() T {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.list.First()
}

// Search searches for an element in the queue.
func (l *LQueue[T]) Search(item T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _, ok := l.list.Find(item); ok {
		return true
	}

	return false
}

// Size returns the queue size.
func (l *LQueue[T]) Size() int {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.n
}
