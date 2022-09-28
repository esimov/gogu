// Linked-list implementation of the queue (FIFO) data structure.
package queue

import (
	"github.com/esimov/gogu/list"
)

// LQueue implements the linked-list version of the FIFO queue.
type LQueue[T comparable] struct {
	items *list.DList[T]
}

// NewLinked creates a new FIFO queue where the items are stored in a linked-list.
func NewLinked[T comparable](t T) *LQueue[T] {
	list := list.InitDoubly(t)

	return &LQueue[T]{
		items: list,
	}
}

// Enqueue inserts a new element at the end of the queue.
func (l *LQueue[T]) Enqueue(item T) {
	l.items.Append(item)
}

// Dequeue retrieves and removes the first element from the queue.
// The queue size will be decreased by one.
func (l *LQueue[T]) Dequeue() (item T, err error) {
	var t T
	node, err := l.items.Shift()
	if err != nil {
		return t, err
	}
	return l.items.Data(node), nil
}

// Peek returns the first element of the queue. It does not remove it.
func (l *LQueue[T]) Peek() T {
	return l.items.First()
}

// Search searches for an element in the queue.
func (l *LQueue[T]) Search(item T) bool {
	if _, ok := l.items.Find(item); ok {
		return true
	}

	return false
}
