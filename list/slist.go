package gogu

import (
	"fmt"
)

// singleNode has two components: the data and a pointer to the next singleNode of the list.
type singleNode[T comparable] struct {
	data T
	next *singleNode[T]
}

// SList holds the individual nodes of the list.
type SList[T comparable] struct {
	singleNode[T]
}

// newNode creates a new node. It requires the data, but the pointer to the next node should be empty (nil)
// This will be updated on the linked list basic operations like unshift, append and insert after.
func newNode[T comparable](data T) *singleNode[T] {
	return &singleNode[T]{
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

	firstNode := l.singleNode
	node.next = &firstNode
	l.singleNode = *node
}

// InsertAfter inserts a new node after the provided node.
// In case the requested node is not in the list it returns an error.
func (l *SList[T]) InsertAfter(prev *singleNode[T], data T) error {
	if prev == nil {
		return fmt.Errorf("the provided node does not exists")
	}

	node := newNode(data)
	node.next = prev.next
	prev.next = node

	return nil
}

// Append inserts a new node at the end of the list.
func (l *SList[T]) Append(data T) *singleNode[T] {
	node := newNode(data)
	lastNode := &l.singleNode

	if l.next == nil {
		l.singleNode = *lastNode
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

// Replace changes the node old value with the new one.
// It returns an error in case the requested value does not exists.
func (l *SList[T]) Replace(oldVal, newVal T) (*singleNode[T], error) {
	node := &l.singleNode

	// Go through the list until the requested node is reached.
	for {
		if node.next == nil { // if this is the last node
			if node.data == oldVal {
				node.data = newVal
				break
			}
			return nil, fmt.Errorf("requested node does not exists")
		}
		if node.data == oldVal {
			node.data = newVal
			break
		}
		node = node.next
	}

	return node, nil
}

// Delete deletes the specified node from the list.
func (l *SList[T]) Delete(n *singleNode[T]) error {
	tmp := &l.singleNode
	// Check if the node we want to delete is the first one.
	if tmp.data == n.data {
		l.singleNode = *tmp.next
		return nil
	}

	prev := singleNode[T]{}
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
func (l *SList[T]) DeleteFirst() *singleNode[T] {
	head := &l.singleNode
	node := l.singleNode

	if head.next != nil {
		head = head.next
		l.singleNode = *head
	}

	return &node
}

// DeleteLast deletes the last node from the list.
func (l *SList[T]) DeleteLast() *singleNode[T] {
	head := l.singleNode
	tmp := &l.singleNode

	node := &singleNode[T]{}
	for tmp.next.next != nil {
		node = tmp
		tmp = tmp.next
	}
	tmp.next = nil
	l.singleNode = head

	return node
}

// Each iterates over the elements of the linked list and invokes
// the callback function, having as parameter the nodes data.
func (l *SList[T]) Each(fn func(data T)) {
	node := &l.singleNode
	tmp := l.singleNode

	for {
		fn(l.data)
		if node.next == nil {
			break
		}
		l.singleNode = *node.next
	}

	// Move the pointer back to the first node.
	l.singleNode = tmp
}
