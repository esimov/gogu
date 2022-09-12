package gogu

import (
	"fmt"
)

// node has two components: the data and a pointer to the next node of the list.
type node[T any] struct {
	data T
	next *node[T]
}

// SList holds the nodes of the simple linked list.
type SList[T any] struct {
	node[T]
}

// newNode creates a new node. It requires the data, but the pointer to the next node should be empty (nil)
// This will be updated on the linked list basic operations like unshift, append and insert after.
func newNode[T any](data T) *node[T] {
	return &node[T]{
		data: data,
		next: nil,
	}
}

// InitList initializes a new single linked list with a single node
// Since this is the only node its next pointer will be nil.
func InitList[T any](data T) *SList[T] {
	return &SList[T]{
		*newNode(data),
	}
}

// Unshift insert a new node at the beginning of the linked list.
func (l *SList[T]) Unshift(data T) {
	node := newNode(data)

	firstNode := l.node
	node.next = &firstNode
	l.node = *node
}

// InsertAfter inserts a new node after the provided node.
// In case the node does not exists it returns an error.
func (l *SList[T]) InsertAfter(prev *node[T], data T) error {
	if prev.next == nil {
		return fmt.Errorf("the provided node does not exists")
	}

	node := newNode(data)
	node.next = prev.next
	prev.next = node

	return nil
}

// Append inserts a new node at the end of the linked list.
func (l *SList[T]) Append(data T) *node[T] {
	node := newNode(data)
	lastNode := &l.node

	if l.node.next == nil {
		l.node = *lastNode
	}

	for {
		if lastNode.next == nil {
			break
		}
		lastNode = lastNode.next
	}

	lastNode.next = node
	node.next = nil

	return lastNode
}

// Each iterates over the elements of the linked list and calls
// the callback function, having as parameter the nodes data.
func (l *SList[T]) Each(fn func(data T)) {
	node := &l.node
	for {
		fn(l.node.data)
		if node.next == nil {
			break
		}
		l.node = *node.next
	}
}
