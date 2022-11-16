package heap

import (
	"sync"

	"github.com/esimov/torx"
)

// Sort sorts the heap in ascending or descening order depending on the heap type.
// If the heap is a max heap, the heap is sorted in ascending order,
// otherwise in case the heap is a min heap, it is sorted in descending order.
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
