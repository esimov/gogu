package list

import (
	"fmt"
)

// SingleNode has two components: the value and a pointer to the next node of the list.
type SingleNode[T comparable] struct {
	Value T
	next  *SingleNode[T] // no need to be exported
}

// SList contains the node elements of the singly linked list.
type SList[T comparable] struct {
	SingleNode[T]
}

// newNode creates a new singly linked list node element.
// It holds a pointer to the next node (which is nil on initialization) and the node value.
func newNode[T comparable](value T) *SingleNode[T] {
	return &SingleNode[T]{
		Value: value,
		next:  nil,
	}
}

// Init initializes a new singly linked list with one node.
// Because this is the only node in the list its next pointer will be nil.
func Init[T comparable](value T) *SList[T] {
	return &SList[T]{
		*newNode(value),
	}
}

// Unshift inserts a new node at the beginning of the list.
func (l *SList[T]) Unshift(value T) {
	newNode := newNode(value)

	firstNode := l.SingleNode
	newNode.next = &firstNode
	l.SingleNode = *newNode
}

// Append inserts a new node at the end of the list.
func (l *SList[T]) Append(value T) {
	newNode := newNode(value)
	head := &l.SingleNode

	if l.next == nil {
		l.SingleNode = *head
	}

	for head.next != nil {
		head = head.next
	}

	head.next = newNode
	newNode.next = nil
}

// InsertAfter inserts a new node after the current node.
// It returns an error in case the requested node does not exists.
func (l *SList[T]) InsertAfter(prev *SingleNode[T], value T) error {
	if prev == nil {
		return fmt.Errorf("the provided node does not exists")
	}

	if _, found := l.Find(prev.Value); !found {
		return fmt.Errorf("the node to be deleted does not exists")
	}

	newNode := newNode(value)
	newNode.next = prev.next
	prev.next = newNode

	return nil
}

// Replace replaces a node's value with a new one.
// It returns an error in case the requested node does not exists.
func (l *SList[T]) Replace(oldVal, newVal T) error {
	head := &l.SingleNode

	// Go through the list until the requested node is reached.
	for {
		if head.next == nil { // if this is the last node
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
func (l *SList[T]) Delete(node *SingleNode[T]) error {
	head := &l.SingleNode

	if _, found := l.Find(node.Value); !found {
		return fmt.Errorf("the node to be deleted does not exists")
	}

	// Check if the node we want to delete is the first one.
	if head == node {
		if head.next == nil {
			return fmt.Errorf("cannot remove the node if there is only one element in the list")
		}
		l.SingleNode = *head.next
		return nil
	}

	prev := SingleNode[T]{}
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

// Shift retrieves and removes the first node from the list.
func (l *SList[T]) Shift() {
	head := &l.SingleNode

	if head.next != nil {
		head = head.next
		l.SingleNode = *head
	}
}

// Pop removes the last node from the list.
func (l *SList[T]) Pop() {
	head := &l.SingleNode

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
func (l *SList[T]) Find(val T) (*SingleNode[T], bool) {
	head := l.SingleNode

	for n := &l.SingleNode; n != nil; n = n.next {
		if n.Value == val {
			l.SingleNode = head
			return n, true
		}
	}

	// Move the pointer to the head of the linked list.
	l.SingleNode = head

	return nil, false
}

// Each iterates over the elements of the linked list and invokes
// the callback function, having as parameter the nodes' value.
func (l *SList[T]) Each(fn func(value T)) {
	head := &l.SingleNode
	node := l.SingleNode

	for {
		fn(l.Value)
		if head.next == nil {
			break
		}
		l.SingleNode = *head.next
	}

	// Move the pointer back to the first node.
	l.SingleNode = node
}
