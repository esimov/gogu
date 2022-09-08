package gogu

// HeapSort sorts the heap in ascending or descening order, depending on the heap type.
// If the heap is a max heap, the heap is sorted in ascending order,
// otherwise if the heap is a min heap, then is sorted in descending order.
func HeapSort[T comparable](data []T, comp CompFn[T]) []T {
	heap := FromSlice(data, comp)

	for i := heap.Size() - 1; i > 0; i-- {
		swap(data, 0, i)
		heap.moveDown(i, 0)
	}

	return heap.GetValues()
}
