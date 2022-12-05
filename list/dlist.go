// Package list provides an implementation of the linked list value structure.
// It comes with two version: singly and doubly linked list.
// The singly linked list version has a value element storing the node value
// and a pointer to the next element of the list.
// The doubly linked list version has an additional pointer to previous node.
package list

import (
	"fmt"
)

// DoubleNode holds an additional prev pointer to the node before.
type DoubleNode[T comparable] struct {
	Value T
	next  *DoubleNode[T]
	prev  *DoubleNode[T]
}

// DList contains the node elements of the doubly linked list.
type DList[T comparable] struct {
	DoubleNode[T]
}

// newDNode creates a new doubly linked list node element.
func newDNode[T comparable](value T) *DoubleNode[T] {
	return &DoubleNode[T]{
		Value: value,
		next:  nil,
		prev:  nil,
	}
}

// InitDList initializes a doubly linked list with one node.
// Because this is the only node in the list, its next and prev pointers are nil.
func InitDList[T comparable](value T) *DList[T] {
	return &DList[T]{
		*newDNode(value),
	}
}

// Unshift inserts a new node at the beginning of the doubly linked list.
func (l *DList[T]) Unshift(value T) {
	newNode := newDNode(value)
	head := l.DoubleNode

	newNode.next = &head
	l.prev = newNode

	// Move the pointer to the new node.
	l.DoubleNode = *newNode
}

// Append inserts a new node at the end of the doubly linked list.
func (l *DList[T]) Append(value T) {
	newNode := newDNode(value)
	head := &l.DoubleNode

	if l.next == nil {
		l.DoubleNode = *head
	}

	for head.next != nil {
		head = head.next
	}

	newNode.next = head.next
	head.next = newNode
	newNode.prev = head
}

// InsertBefore inserts a new node before the current node.
// It returns an error in case the requested node does not exists.
func (l *DList[T]) InsertBefore(node *DoubleNode[T], value T) error {
	head := l.DoubleNode
	if node == nil {
		return fmt.Errorf("the previous node does not exists")
	}

	if _, found := l.Find(node.Value); !found {
		return fmt.Errorf("the node to be deleted does not exists")
	}
	newNode := newDNode(value)

	newNode.prev = node.prev
	node.prev = newNode
	newNode.next = node

	if newNode.prev != nil {
		newNode.prev.next = newNode
	} else {
		newNode.next = &head
		// Move the pointer to the new node.
		l.DoubleNode = *newNode
	}

	return nil
}

// InsertAfter inserts a new node after the existing node.
// It returns an error in case the requested node does not exists.
func (l *DList[T]) InsertAfter(node *DoubleNode[T], value T) error {
	if node == nil {
		return fmt.Errorf("the previous node does not exists")
	}

	if _, found := l.Find(node.Value); !found {
		return fmt.Errorf("the node to be deleted does not exists")
	}

	newNode := newDNode(value)
	newNode.next = node.next
	node.next = newNode
	newNode.prev = node

	if newNode.next != nil {
		newNode.next.prev = newNode
	}

	return nil
}

// Replace replaces a node's value with the new one.
// It returns an error in case the requested node does not exist.
func (l *DList[T]) Replace(oldVal, newVal T) error {
	head := &l.DoubleNode

	// Go through the list until the requested node is reached.
	for {
		if head.next == nil {
			if head.Value == oldVal {
				head.Value = newVal
				break
			}
			return fmt.Errorf("requested node does not exists")
		}
		if head.Value == oldVal {
			head.Value = newVal
			break
		}
		head = head.next
	}

	return nil
}

// Delete removes the specified node from the list.
func (l *DList[T]) Delete(node *DoubleNode[T]) error {
	head := &l.DoubleNode

	if _, found := l.Find(node.Value); !found {
		return fmt.Errorf("the node to be deleted does not exists")
	}

	if head.next == nil && head.prev == nil {
		return fmt.Errorf("cannot delete the node if there is only one element in the list")
	}

	// Check if the node to be deleted is the head node.
	if head.Value == node.Value {
		l.DoubleNode = *head.next
		return nil
	}

	// Replace the next pointer of the node to be deleted
	// only if it's not the last element of the list.
	if node.next != nil {
		node.next.prev = node.prev
	}

	// Replace the prev pointer of the node to be deleted
	// only if it's not the first element of the list.
	if node.prev != nil {
		node.prev.next = node.next
	}

	return nil
}

// Shift retrieves and removes the first node from the list.
func (l *DList[T]) Shift() *DoubleNode[T] {
	head := &l.DoubleNode
	node := l.DoubleNode

	if head.next == nil {
		var value T
		head.next = nil
		head.prev = nil
		head.Value = value

		l.DoubleNode = *head
	} else {
		head = head.next
		l.DoubleNode = *head
	}

	return &node
}

// Pop removes the last node from the list.
func (l *DList[T]) Pop() *DoubleNode[T] {
	head := &l.DoubleNode
	node := DoubleNode[T]{}

	if head.next == nil {
		head = nil
	} else {
		tmp := head
		node = *tmp
		for tmp.next.next != nil {
			tmp = tmp.next
			node = *tmp
		}
		tmp.next = nil
	}
	return &node
}

// Find searches for a node element in the linked list.
// It returns the node in case the element is found otherwise nil.
func (l *DList[T]) Find(val T) (*DoubleNode[T], bool) {
	head := &l.DoubleNode

	for n := &l.DoubleNode; n != nil; n = n.next {
		if n.Value == val {
			l.DoubleNode = *head
			return n, true
		}
	}

	// Move the pointer to the head of the linked list.
	l.DoubleNode = *head

	return nil, false
}

// First retrieves the first element of the doubly linked list.
func (l *DList[T]) First() T {
	head := l.DoubleNode

	return head.Value
}

// Last retrieves the last element of the doubly linked list.
func (l *DList[T]) Last() T {
	head := l.DoubleNode
	var value T

	for l.DoubleNode.next != nil {
		l.DoubleNode = *l.DoubleNode.next
	}
	value = l.DoubleNode.Value

	// Move the pointer to the head of the linked list.
	l.DoubleNode = head

	return value
}

// Each iterates over the elements of the linked list and invokes
// the callback function having as parameter the nodes' value.
func (l *DList[T]) Each(fn func(value T)) {
	head := &l.DoubleNode
	node := l.DoubleNode
	for {
		fn(l.Value)
		if head.next == nil {
			break
		}
		l.DoubleNode = *head.next
	}
	// Move the pointer back to the first node.
	l.DoubleNode = node
}

// Val retrieves the node value.
func (l *DList[T]) Val(node *DoubleNode[T]) T {
	return node.Value
}

// Clear deletes all the nodes from the list.
func (l *DList[T]) Clear() {
	head := &l.DoubleNode
	head.next = nil
	head.prev = nil
}
