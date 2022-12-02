package list

import (
	"fmt"
)

// singleNode has two components: the data and a pointer to the next node of the list.
type singleNode[T comparable] struct {
	data T
	next *singleNode[T]
}

// SList contains the node elements of the singly linked list.
type SList[T comparable] struct {
	singleNode[T]
}

// newNode creates a new singly linked list node element.
// It holds a pointer to the next node (which is nil on initialization) and the node data.
func newNode[T comparable](data T) *singleNode[T] {
	return &singleNode[T]{
		data: data,
		next: nil,
	}
}

// Init initializes a new singly linked list with one node.
// Because this is the only node in the list its next pointer will be nil.
func Init[T comparable](data T) *SList[T] {
	return &SList[T]{
		*newNode(data),
	}
}

// Unshift inserts a new node at the beginning of the list.
func (l *SList[T]) Unshift(data T) {
	newNode := newNode(data)

	firstNode := l.singleNode
	newNode.next = &firstNode
	l.singleNode = *newNode
}

// Append inserts a new node at the end of the list.
func (l *SList[T]) Append(data T) {
	newNode := newNode(data)
	head := &l.singleNode

	if l.next == nil {
		l.singleNode = *head
	}

	for head.next != nil {
		head = head.next
	}

	head.next = newNode
	newNode.next = nil
}

// InsertAfter inserts a new node after the current node.
// It returns an error in case the requested node does not exists.
func (l *SList[T]) InsertAfter(prev *singleNode[T], data T) error {
	if prev == nil {
		return fmt.Errorf("the provided node does not exists")
	}

	if _, found := l.Find(prev.data); !found {
		return fmt.Errorf("the node to be deleted does not exists")
	}

	newNode := newNode(data)
	newNode.next = prev.next
	prev.next = newNode

	return nil
}

// Replace replaces a node's value with a new one.
// It returns an error in case the requested node does not exists.
func (l *SList[T]) Replace(oldVal, newVal T) error {
	head := &l.singleNode

	// Go through the list until the requested node is reached.
	for {
		if head.next == nil { // if this is the last node
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
func (l *SList[T]) Delete(node *singleNode[T]) error {
	head := &l.singleNode

	if _, found := l.Find(node.data); !found {
		return fmt.Errorf("the node to be deleted does not exists")
	}

	// Check if the node we want to delete is the first one.
	if head == node {
		if head.next == nil {
			return fmt.Errorf("cannot remove the node if there is only one element in the list")
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
func (l *SList[T]) Shift() {
	head := &l.singleNode

	if head.next != nil {
		head = head.next
		l.singleNode = *head
	}
}

// Pop removes the last node from the list.
func (l *SList[T]) Pop() {
	head := &l.singleNode

	if head.next == nil {
		head = nil
	} else {
		tmp := head
		for tmp.next.next != nil {
			tmp = tmp.next
		}
		tmp.next = nil
	}
}

// Find search for a node element in the linked list.
// It returns the node in case the element is found otherwise nil.
func (l *SList[T]) Find(val T) (*singleNode[T], bool) {
	head := l.singleNode

	for n := &l.singleNode; n != nil; n = n.next {
		if n.data == val {
			l.singleNode = head
			return n, true
		}
	}

	// Move the pointer to the head of the linked list.
	l.singleNode = head

	return nil, false
}

// Each iterates over the elements of the linked list and invokes
// the callback function, having as parameter the nodes' data.
func (l *SList[T]) Each(fn func(data T)) {
	head := &l.singleNode
	node := l.singleNode

	for {
		fn(l.data)
		if head.next == nil {
			break
		}
		l.singleNode = *head.next
	}

	// Move the pointer back to the first node.
	l.singleNode = node
}
