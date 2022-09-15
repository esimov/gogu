package gogu

import "fmt"

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
	lastNode := &l.doubleNode

	if l.next == nil {
		l.doubleNode = *lastNode
	}

	for {
		if lastNode.next == nil {
			break
		}
		lastNode = lastNode.next
	}

	node.next = lastNode.next
	lastNode.next = node
	node.prev = lastNode

	return node
}

// InsertBefore inserts a new node before the provided node in the doubly linked list.
// In case the requested node is not in the list it returns an error.
func (l *DList[T]) InsertBefore(node *doubleNode[T], data T) error {
	if node == nil {
		return fmt.Errorf("the previous node does not exists")
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

	node := newDNode(data)
	node.next = prev.next
	prev.next = node
	node.prev = prev

	if node.next != nil {
		node.next.prev = node
	}
	l.doubleNode = *node.prev

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
func (l *DList[T]) Delete(n *doubleNode[T]) error {
	tmp := &l.doubleNode
	// Check if the node we want to delete is the first one.
	if tmp.data == n.data {
		l.doubleNode = *tmp.next
		return nil
	}

	prev := doubleNode[T]{}
	// Go through the list until the requested node is reached.
	for tmp.next != nil && tmp.data != n.data {
		prev = *tmp
		tmp = tmp.next
	}

	// Check if the node we want to delete is the last one.
	if tmp.next == nil {
		l.Pop()
		return nil
	}
	*prev.next = *tmp.next

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
