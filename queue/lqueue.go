// Linked-list implementation of the queue (FIFO) data structure.
package queue

import (
	"github.com/esimov/gogu/list"
)

// lQueue implements the linked-list version of the FIFO queue.
type lQueue[T comparable] struct {
	items *list.DList[T]
}

// NewLinked creates a new FIFO queue where the items are stored in a linked-list.
func NewLinked[T comparable](t T) *lQueue[T] {
	list := list.InitDoubly(t)

	return &lQueue[T]{
		items: list,
	}
}

// Enqueue inserts a new element at the end of the queue.
func (l *lQueue[T]) Enqueue(item T) {
	l.items.Append(item)
}

// Dequeue retrieves and removes the first element from the queue.
// The queue size will be decreased by one.
func (l *lQueue[T]) Dequeue() (item T, err error) {
	var t T
	node, err := l.items.Shift()
	if err != nil {
		return t, err
	}
	return l.items.Data(node), nil
}

// Peek returns the first element of the queue. It does not remove it.
func (l *lQueue[T]) Peek() T {
	return l.items.First()
}

// Search searches for an element in the queue.
func (l *lQueue[T]) Search(item T) bool {
	if _, ok := l.items.Find(item); ok {
		return true
	}

	return false
}
