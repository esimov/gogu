// Package heap provides a thread-safe implementation of the binary heap data structure.
// A common implementation of the heap is the binary tree, where each node of the subtree
// satisfies the heap property:
// each node of the subtree is greather or equal then the parent node in case of min heap,
// and less or equal than the parent node in case of max heap.
// The conditional operator used on the heap initialization defines the heap type
package heap

import (
	"sync"

	"github.com/esimov/torx"
)

// Sort sorts the heap in ascending or descening order depending on the heap type.
// In case the heap is a max heap, the heap is sorted in ascending order, otherwise in descending order.
func Sort[T comparable](data []T, comp torx.CompFn[T]) []T {
	mu := &sync.Mutex{}
	heap := FromSlice(data, comp)

	for i := heap.Size() - 1; i > 0; i-- {
		mu.Lock()
		swap(data, 0, i)
		mu.Unlock()

		heap.moveDown(i, 0)
	}

	return heap.GetValues()
}
