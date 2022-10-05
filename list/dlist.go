// Package list provides an implementation of the linked list data structure.
// This is the doubly linked list version, which compared to the
// singly linked list variant has an additional pointer to the previous node.
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

// InitDoubly initializes a doubly linked list with one node.
// Because this is the only node curently existing in the list, its next and prev pointers are nil.
func InitDoubly[T comparable](data T) *DList[T] {
	return &DList[T]{
		*newDNode(data),
	}
}

// Unshift inserts a new node at the beginning of the doubly linked list.
func (l *DList[T]) Unshift(data T) *doubleNode[T] {
	newNode := newDNode(data)
	head := l.doubleNode

	newNode.next = &head
	l.prev = newNode

	// Move the pointer to the new node.
	l.doubleNode = *newNode

	return newNode
}

// Append inserts a new node at the end of the doubly linked list.
func (l *DList[T]) Append(data T) *doubleNode[T] {
	newNode := newDNode(data)
	head := &l.doubleNode

	if l.next == nil {
		l.doubleNode = *head
	}

	for {
		if head.next == nil {
			break
		}
		head = head.next
	}

	newNode.next = head.next
	head.next = newNode
	newNode.prev = head

	return newNode
}

// InsertBefore inserts a new node before the current node.
// In case the requested node is not in the list it returns an error.
func (l *DList[T]) InsertBefore(node *doubleNode[T], data T) error {
	if node == nil {
		return fmt.Errorf("the previous node does not exists")
	}

	if _, found := l.Find(node.data); !found {
		return fmt.Errorf("the node to be deleted does not exists")
	}
	newNode := newDNode(data)

	if node.prev != nil {
		newNode = node.prev
	}

	newNode.next = node
	l.doubleNode = *newNode

	return nil
}

// InsertAfter inserts a new node after the existing node.
// In case the requested node is not in the list it returns an error.
func (l *DList[T]) InsertAfter(prev *doubleNode[T], data T) error {
	if prev == nil {
		return fmt.Errorf("the previous node does not exists")
	}

	if _, found := l.Find(prev.data); !found {
		return fmt.Errorf("the node to be deleted does not exists")
	}

	newNode := newDNode(data)
	newNode.next = prev.next
	prev.next = newNode
	newNode.prev = prev

	if newNode.next != nil {
		newNode.next.prev = newNode
	}

	return nil
}

// Replace replaces a node's value with a new one.
// It returns an error in case the requested node does not exists.
func (l *DList[T]) Replace(oldVal, newVal T) (*doubleNode[T], error) {
	head := &l.doubleNode

	// Go through the list until the requested node is reached.
	for {
		if head.next == nil {
			if head.data == oldVal {
				head.data = newVal
				break
			}
			return nil, fmt.Errorf("requested node does not exists")
		}
		if head.data == oldVal {
			head.data = newVal
			break
		}
		head = head.next
	}

	return head, nil
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
func (l *DList[T]) Shift() (*doubleNode[T], error) {
	head := &l.doubleNode
	node := l.doubleNode

	if head.next == nil {
		return nil, fmt.Errorf("cannot remove the node if there is only one element in the list")
	} else {
		head = head.next
		l.doubleNode = *head
	}

	return &node, nil
}

// Pop removes the last node from the list.
func (l *DList[T]) Pop() (*doubleNode[T], error) {
	head := &l.doubleNode
	node := &doubleNode[T]{}

	if head.next == nil {
		return nil, fmt.Errorf("cannot remove the node if there is only one element in the list")
	} else {
		tmp := head
		for tmp.next.next != nil {
			tmp = tmp.next
		}
		tmp.next = nil
	}

	return node, nil
}

// Find search for a node element in the linked list.
// It returns the node in case the element is found otherwise nil.
func (l *DList[T]) Find(val T) (*doubleNode[T], bool) {
	head := l.doubleNode
	var found bool

	for n := &l.doubleNode; n != nil; n = n.next {
		if found {
			break
		}
		if n.data == val {
			l.doubleNode = head
			found = true

			return n, true
		}
	}

	// Move the pointer to the head of the linked list.
	l.doubleNode = head

	return nil, false
}

// First retrieves the first element of the double linked list.
func (l *DList[T]) First() T {
	head := l.doubleNode

	return head.data
}

// Last retrieves the last element of the double linked list.
func (l *DList[T]) Last() T {
	head := l.doubleNode
	var data T
	for {
		if l.doubleNode.next == nil {
			break
		}
		l.doubleNode = *l.doubleNode.next
	}
	data = l.doubleNode.data

	// Move the pointer to the head of the linked list.
	l.doubleNode = head

	return data
}

// Each iterates over the elements of the linked list and invokes
// the callback function, having as parameter the nodes data.
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
