package gogu

type Heap[T any] struct {
	data []T
	comp func(i, j T) bool
}

type CompFn[T any] func(i, j T) bool

// NewHeap creates a new heap data structure having two components:
// a data slice holding the concrete values and a comparator function returning a boolean value.
// The comparision sign decides if the heap is a max heap or min heap.
func NewHeap[T any](comp CompFn[T]) *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0),
		comp: comp,
	}
}

// Size returns the heap size.
func (h *Heap[T]) Size() int {
	return len(h.data)
}

// IsEmpty returns true if a heap is empty, otherwise false.
func (h *Heap[T]) IsEmpty() bool {
	if h.Size() > 0 {
		return false
	}
	return true
}

// Clear removes all the elements from the heap.
func (h *Heap[T]) Clear() {
	if h.Size() == 0 {
		return
	}
	h.data = h.data[:0]
}

// Insert stores a new value at the end of the heap and calls the heapify algorithm to reorder
// the existing values in ascending or descending order, depending on the comparator function.
func (h *Heap[T]) Insert(val T) {
	h.data = append(h.data, val)
	h.moveUp(h.Size() - 1)
}

// Delete removes the peek element from the heap and reorder the existing elements.
// The removed element can be the minimum or maximum depending on the comparator function.
func (h *Heap[T]) Delete() T {
	var val T
	if h.Size() == 0 {
		return val
	}

	val = h.Peek()

	h.data[0] = h.data[h.Size()-1]
	h.data = h.data[:h.Size()-1]
	h.moveDown(0)

	return val
}

// Peek returns the first element of the heap.
// This can be the maximum or minimum depending if the heap is a min heap or max heap.
func (h *Heap[T]) Peek() T {
	if h.Size() == 0 {
		var t T
		return t
	}

	return h.data[0]
}

// Convert a min heap to max heap and vice versa.
func (h *Heap[T]) Convert() {
	// Start from bottom-rightmost internal mode and reorder all internal nodes.
	for i := (h.Size() - 2) / 2; i >= 0; i-- {
		h.moveDown(i)
	}
}

// Merge joins two heaps into a new one preserving the original heaps.
func (h *Heap[T]) Merge(h2 *Heap[T]) *Heap[T] {
	newHeap := NewHeap(h.comp)

	for i := 0; i < h.Size(); i++ {
		newHeap.Insert(h.data[i])
	}

	for i := 0; i < h2.Size(); i++ {
		newHeap.Insert(h2.data[i])
	}

	return newHeap
}

// Meld merge two heaps into a new one containing all the elements of both and destroys the original heaps.
func (h *Heap[T]) Meld(h2 *Heap[T]) *Heap[T] {
	newHeap := NewHeap(h.comp)

	for i := 0; i < h.Size(); i++ {
		newHeap.Insert(h.data[i])
	}

	for i := 0; i < h2.Size(); i++ {
		newHeap.Insert(h2.data[i])
	}
	h.data = nil
	h2.data = nil

	return newHeap
}

// moveDown moves the element at the position i down to its
// correct position in the heap following the heap rules.
func (h *Heap[T]) moveDown(i int) {
	left := h.leftChild(i)
	right := h.rightChild(i)
	current := i

	if left < h.Size() && h.comp(h.data[left], h.data[i]) {
		current = left
	}

	if right < h.Size() && h.comp(h.data[right], h.data[current]) {
		current = right
	}

	if current != i {
		h.swap(i, current)
		h.moveDown(current)
	}
}

// moveDown moves the element at the position i up to its
// correct position in the heap following the heap rules.
func (h *Heap[T]) moveUp(i int) {
	if h.comp(h.data[h.parent(i)], h.data[i]) {
		h.swap(i, h.parent(i))
		i = h.parent(i)

		h.moveUp(i)
	}
}

// leftChild returns the index of the left child of node at index i.
func (h *Heap[T]) leftChild(i int) int {
	return 2*i + 1
}

// rightChild returns the index of the right child of node at index i.
func (h *Heap[T]) rightChild(i int) int {
	return 2*i + 2
}

// parent returns the index of the child node parent at index i.
func (h *Heap[T]) parent(i int) int {
	return (i - 1) / 2
}

// swap swaps the position of elements at index i and j.
func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
