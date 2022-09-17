// Package list implements a linked list data structure.
// This is the singly Linked list implementation, which has
// a data element storing the node value
// and a pointer to the next element of the list.
package list

import (
	"fmt"
)

// singleNode has two components: the data and a pointer to the next singleNode of the list.
type singleNode[T comparable] struct {
	data T
	next *singleNode[T]
}

// SList contains the individual nodes of the list.
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

// Append inserts a new node at the end of the list.
func (l *SList[T]) Append(data T) *singleNode[T] {
	node := newNode(data)
	head := &l.singleNode

	if l.next == nil {
		l.singleNode = *head
	}

	for {
		if head.next == nil {
			break
		}
		head = head.next
	}

	head.next = node
	node.next = nil

	return node
}

// InsertAfter inserts a new node after the provided node.
// In case the requested node is not in the list it returns an error.
func (l *SList[T]) InsertAfter(prev *singleNode[T], data T) error {
	if prev == nil {
		return fmt.Errorf("the provided node does not exists")
	}

	if _, found := l.Find(prev.data); !found {
		return fmt.Errorf("the node to be deleted does not exists")
	}

	node := newNode(data)
	node.next = prev.next
	prev.next = node

	return nil
}

// Replace replaces a node value with the new one.
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

// Delete removes the specified node from the list.
func (l *SList[T]) Delete(node *singleNode[T]) error {
	head := &l.singleNode

	if _, found := l.Find(node.data); !found {
		return fmt.Errorf("the node to be deleted does not exists")
	}

	// Check if the node we want to delete is the first one.
	if head == node {
		if head.next == nil {
			return fmt.Errorf("cannot delete the node if there is only one element in the list")
		}
		l.singleNode = *head.next
		return nil
	}

	prev := singleNode[T]{}
	// Go through the list until the requested node is reached.
	for head.next != nil && head != node {
		prev = *head
		head = head.next
	}

	// Check if the node we want to delete is the last one.
	if head.next == nil {
		l.Pop()
		return nil
	}
	*prev.next = *head.next

	return nil
}

// Shift removes the first node from the list.
func (l *SList[T]) Shift() *singleNode[T] {
	head := &l.singleNode
	node := l.singleNode

	if head.next != nil {
		head = head.next
		l.singleNode = *head
	}

	return &node
}

// Pop removes the last node from the list.
func (l *SList[T]) Pop() *singleNode[T] {
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

// Find search for a node element in the linked list.
// It returns the node in case the element is found otherwise nil.
func (l *SList[T]) Find(val T) (*singleNode[T], bool) {
	var node *singleNode[T]
	head := l.singleNode
	found := false

	for n := &l.singleNode; n != nil && !found; n = n.next {
		if n.data == val {
			l.singleNode = head
			return n, true
		}
	}

	// Move the pointer to the head of the linked list.
	l.singleNode = head

	return node, false
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
