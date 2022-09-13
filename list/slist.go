package gogu

import (
	"fmt"
)

// node has two components: the data and a pointer to the next node of the list.
type node[T comparable] struct {
	data T
	next *node[T]
}

// SList holds the individual nodes of the list.
type SList[T comparable] struct {
	node[T]
}

// newNode creates a new node. It requires the data, but the pointer to the next node should be empty (nil)
// This will be updated on the linked list basic operations like unshift, append and insert after.
func newNode[T comparable](data T) *node[T] {
	return &node[T]{
		data: data,
		next: nil,
	}
}

// InitList initializes a new single linked list with one node.
// Because this is the only node currently existing in the list its next pointer will be nil.
func InitList[T comparable](data T) *SList[T] {
	return &SList[T]{
		*newNode(data),
	}
}

// Push insert a new node at the beginning of the list.
func (l *SList[T]) Push(data T) {
	node := newNode(data)

	firstNode := l.node
	node.next = &firstNode
	l.node = *node
}

// InsertAfter inserts a new node after the provided node.
// In case the requested node is not in the list it returns an error.
func (l *SList[T]) InsertAfter(prev *node[T], data T) error {
	if prev == nil {
		return fmt.Errorf("the provided node does not exists")
	}

	node := newNode(data)
	node.next = prev.next
	prev.next = node

	return nil
}

// Delete deletes the specified node from the list.
func (l *SList[T]) Delete(n *node[T]) error {
	tmp := &l.node
	// Check if the node we want to delete is the first one.
	if tmp.data == n.data {
		l.node = *tmp.next
		return nil
	}

	prev := node[T]{}
	// Go through the list until the requested node is reached.
	for tmp.next != nil && tmp.data != n.data {
		prev = *tmp
		tmp = tmp.next
	}

	// Check if the node we want to delete is the last one.
	if tmp.next == nil {
		l.DeleteLast()
		return nil
	}
	*prev.next = *tmp.next

	return nil
}

// DeleteFirst deletes the first node from the list.
func (l *SList[T]) DeleteFirst() *node[T] {
	firstNode := &l.node
	node := l.node

	if firstNode.next != nil {
		firstNode = firstNode.next
		l.node = *firstNode
	}

	return &node
}

// DeleteLast deletes the last node from the list.
func (l *SList[T]) DeleteLast() *node[T] {
	firstNode := l.node
	tmp := &l.node
	node := &node[T]{}

	for tmp.next.next != nil {
		node = tmp
		tmp = tmp.next
	}
	tmp.next = nil
	l.node = firstNode

	return node
}

// Append inserts a new node at the end of the list.
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

	return node
}

// Each iterates over the elements of the linked list and invokes
// the callback function, having as parameter the nodes data.
func (l *SList[T]) Each(fn func(data T)) {
	node := &l.node
	tmp := l.node
	for {
		fn(l.node.data)
		if node.next == nil {
			break
		}
		l.node = *node.next
	}
	// Move the pointer back to the first node.
	l.node = tmp
}
