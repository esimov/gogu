// Package list provides an implementation of the linked list data structure.
// It comes with two version: singly and doubly linked list.
// The singly linked list version has a data element storing the node value
// and a pointer to the next element of the list.
// The doubly linked list version has an additional pointer to previous node.
package list

import (
	"fmt"
)

// doubleNode holds an additional prev pointer to the node before.
type doubleNode[T comparable] struct {
	data T
	next *doubleNode[T]
	prev *doubleNode[T]
}

// DList contains the node elements of the doubly linked list.
type DList[T comparable] struct {
	doubleNode[T]
}

// newDNode creates a new doubly linked list node element.
func newDNode[T comparable](data T) *doubleNode[T] {
	return &doubleNode[T]{
		data: data,
		next: nil,
		prev: nil,
	}
}

// InitDList initializes a doubly linked list with one node.
// Because this is the only node in the list, its next and prev pointers are nil.
func InitDList[T comparable](data T) *DList[T] {
	return &DList[T]{
		*newDNode(data),
	}
}

// Unshift inserts a new node at the beginning of the doubly linked list.
func (l *DList[T]) Unshift(data T) {
	newNode := newDNode(data)
	head := l.doubleNode

	newNode.next = &head
	l.prev = newNode

	// Move the pointer to the new node.
	l.doubleNode = *newNode
}

// Append inserts a new node at the end of the doubly linked list.
func (l *DList[T]) Append(data T) {
	newNode := newDNode(data)
	head := &l.doubleNode

	if l.next == nil {
		l.doubleNode = *head
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
func (l *DList[T]) InsertBefore(node *doubleNode[T], data T) error {
	head := l.doubleNode
	if node == nil {
		return fmt.Errorf("the previous node does not exists")
	}

	if _, found := l.Find(node.data); !found {
		return fmt.Errorf("the node to be deleted does not exists")
	}
	newNode := newDNode(data)

	newNode.prev = node.prev
	node.prev = newNode
	newNode.next = node

	if newNode.prev != nil {
		newNode.prev.next = newNode
	} else {
		newNode.next = &head
		// Move the pointer to the new node.
		l.doubleNode = *newNode
	}

	return nil
}

// InsertAfter inserts a new node after the existing node.
// It returns an error in case the requested node does not exists.
func (l *DList[T]) InsertAfter(node *doubleNode[T], data T) error {
	if node == nil {
		return fmt.Errorf("the previous node does not exists")
	}

	if _, found := l.Find(node.data); !found {
		return fmt.Errorf("the node to be deleted does not exists")
	}

	newNode := newDNode(data)
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
	head := &l.doubleNode

	// Go through the list until the requested node is reached.
	for {
		if head.next == nil {
			if head.data == oldVal {
				head.data = newVal
				break
			}
			return fmt.Errorf("requested node does not exists")
		}
		if head.data == oldVal {
			head.data = newVal
			break
		}
		head = head.next
	}

	return nil
}

// Delete removes the specified node from the list.
func (l *DList[T]) Delete(node *doubleNode[T]) error {
	head := &l.doubleNode

	if _, found := l.Find(node.data); !found {
		return fmt.Errorf("the node to be deleted does not exists")
	}

	if head.next == nil && head.prev == nil {
		return fmt.Errorf("cannot delete the node if there is only one element in the list")
	}

	// Check if the node to be deleted is the head node.
	if head.data == node.data {
		l.doubleNode = *head.next
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

// Shift removes the first node from the list.
func (l *DList[T]) Shift() *doubleNode[T] {
	head := &l.doubleNode
	node := l.doubleNode

	if head.next == nil {
		var data T
		head.next = nil
		head.prev = nil
		head.data = data

		l.doubleNode = *head
	} else {
		head = head.next
		l.doubleNode = *head
	}

	return &node
}

// Pop removes the last node from the list.
func (l *DList[T]) Pop() *doubleNode[T] {
	head := &l.doubleNode
	node := doubleNode[T]{}

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
func (l *DList[T]) Find(val T) (*doubleNode[T], bool) {
	head := &l.doubleNode

	for n := &l.doubleNode; n != nil; n = n.next {
		if n.data == val {
			l.doubleNode = *head
			return n, true
		}
	}

	// Move the pointer to the head of the linked list.
	l.doubleNode = *head

	return nil, false
}

// First retrieves the first element of the doubly linked list.
func (l *DList[T]) First() T {
	head := l.doubleNode

	return head.data
}

// Last retrieves the last element of the doubly linked list.
func (l *DList[T]) Last() T {
	head := l.doubleNode
	var data T

	for l.doubleNode.next != nil {
		l.doubleNode = *l.doubleNode.next
	}
	data = l.doubleNode.data

	// Move the pointer to the head of the linked list.
	l.doubleNode = head

	return data
}

// Each iterates over the elements of the linked list and invokes
// the callback function having as parameter the nodes' data.
func (l *DList[T]) Each(fn func(data T)) {
	head := &l.doubleNode
	node := l.doubleNode
	for {
		fn(l.data)
		if head.next == nil {
			break
		}
		l.doubleNode = *head.next
	}
	// Move the pointer back to the first node.
	l.doubleNode = node
}

// Data retrieves the node value.
func (l *DList[T]) Data(node *doubleNode[T]) T {
	return node.data
}

// Clear deletes all the nodes from the list.
func (l *DList[T]) Clear() {
	head := &l.doubleNode
	head.next = nil
	head.prev = nil
}
