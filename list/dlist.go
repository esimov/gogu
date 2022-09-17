// Package list implements a linked list data structure.
// This is the doubly Linked list implementation,
// which compared to the singly linked list variant has
// an additional previous pointer to the node before.
package list

import (
	"fmt"
)

// doubleNode is a doubly linked list node which has aditional
// pointer holding the memory address of the previous node.
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

// InitDoubleList initializes a doubly linked list with one node.
// Because this is the only node curently existing in the list its next and prev pointers are nil.
func InitDoubleList[T comparable](data T) *DList[T] {
	return &DList[T]{
		*newDNode(data),
	}
}

// Append inserts a new node at the beginning of the doubly linked list.
func (l *DList[T]) Push(data T) *doubleNode[T] {
	node := newDNode(data)

	firstNode := l.doubleNode
	node.next = &firstNode
	l.prev = node

	// Move the pointer to the new node.
	l.doubleNode = *node

	return node
}

// Append inserts a new node at the end of the doubly linked list.
func (l *DList[T]) Append(data T) *doubleNode[T] {
	node := newDNode(data)
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

	node.next = head.next
	head.next = node
	node.prev = head

	return node
}

// InsertBefore inserts a new node before the provided node in the doubly linked list.
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

// InsertAfter inserts a new node after the provided node in the doubly linked list.
// In case the requested node is not in the list it returns an error.
func (l *DList[T]) InsertAfter(prev *doubleNode[T], data T) error {
	if prev == nil {
		return fmt.Errorf("the previous node does not exists")
	}

	if _, found := l.Find(prev.data); !found {
		return fmt.Errorf("the node to be deleted does not exists")
	}

	node := newDNode(data)
	node.next = prev.next
	prev.next = node
	node.prev = prev

	if node.next != nil {
		node.next.prev = node
	}

	return nil
}

// Replace replaces a node's value with the new one.
// It returns an error in case the requested value does not exists.
func (l *DList[T]) Replace(oldVal, newVal T) (*doubleNode[T], error) {
	node := &l.doubleNode

	// Go through the list until the requested node is reached.
	for {
		if node.next == nil {
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

	if head.next != nil {
		head = head.next
		l.doubleNode = *head
	}

	return &node
}

// Pop removes the last node from the list.
func (l *DList[T]) Pop() *doubleNode[T] {
	head := l.doubleNode
	tmp := &l.doubleNode

	node := &doubleNode[T]{}
	for tmp.next.next != nil {
		node = tmp
		tmp = tmp.next
	}
	tmp.next = nil
	l.doubleNode = head

	return node
}

// Find search for a node element in the linked list.
// It returns the node in case the element is found otherwise nil.
func (l *DList[T]) Find(val T) (*doubleNode[T], bool) {
	var node *doubleNode[T]
	head := l.doubleNode
	found := false

	for n := &l.doubleNode; n != nil && !found; n = n.next {
		if n.data == val {
			l.doubleNode = head
			return n, true
		}
	}

	// Move the pointer to the head of the linked list.
	l.doubleNode = head

	return node, false
}

// Each iterates over the elements of the linked list and invokes
// the callback function, having as parameter the nodes data.
func (l *DList[T]) Each(fn func(data T)) {
	node := &l.doubleNode
	tmp := l.doubleNode
	for {
		fn(l.data)
		if node.next == nil {
			break
		}
		l.doubleNode = *node.next
	}
	// Move the pointer back to the first node.
	l.doubleNode = tmp
}
