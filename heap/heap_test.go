package gogu

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	assert := assert.New(t)

	// Min Heap
	comp := func(a, b int) bool {
		return a > b
	}
	h := NewHeap(comp)
	assert.Empty(h.Size())
	assert.True(h.IsEmpty())
	h.Push(10)
	assert.Equal(len(h.GetValues()), h.Size())
	assert.Equal(1, h.Size())
	assert.Equal(10, h.Pop())
	assert.True(h.IsEmpty())

	values := []int{9, 3, 20, 8, 6, 5, 12, 10, 9, 18}
	for _, v := range values {
		h.Push(v)
	}
	assert.Equal([]int{3, 6, 5, 9, 8, 20, 12, 10, 9, 18}, h.GetValues())

	h.Push(7)
	assert.Equal([]int{3, 6, 5, 9, 7, 20, 12, 10, 9, 18, 8}, h.GetValues())

	h.Clear()
	assert.Empty(h.GetValues())

	v := h.Pop()
	assert.Equal(0, v)

	h.Push(10, 4, 2, 5, 3)
	assert.Equal([]int{2, 3, 4, 10, 5}, h.GetValues())

	for range h.GetValues() {
		h.Pop()
	}
	assert.Empty(h.Size())

	// Max Heap
	comp = func(a, b int) bool {
		return a < b
	}

	h2 := NewHeap(comp)

	values = []int{9, 3, 20, 8, 6, 5, 12, 10, 9, 18}
	for _, v := range values {
		h2.Push(v)
	}
	assert.Equal([]int{20, 18, 12, 9, 10, 5, 9, 3, 8, 6}, h2.GetValues())

	// Import from Slice
	h3 := FromSlice(values, func(a, b int) bool { return a < b })
	assert.NotEmpty(h3.GetValues())
	assert.Equal([]int{20, 18, 12, 9, 10, 5, 9, 3, 8, 6}, h3.GetValues())

	// Insert

	h4 := NewHeap(comp)
	h4.Push(3, 5, 1, 4)
	fmt.Println(h4.GetValues())
	ok, err := h4.Delete(10)
	assert.Error(err)
	assert.False(ok)
	ok, err = h4.Delete(1)
	assert.NoError(err)
	assert.True(ok)
	assert.Equal(3, h4.Size())
	fmt.Println(h4.GetValues())

	// h3 := h.Merge(h2)
	fmt.Println("H2:", h2.data)
	// fmt.Println("H3:", h3.data)

	// h4 := h.Meld(h2)
	// fmt.Println("H4:", h4.data)
	// fmt.Println("H:", h.data)
	// fmt.Println("H2:", h2.data)
}
