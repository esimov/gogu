package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	assert := assert.New(t)

	// sorting in ascending order
	data := []int{1, 3, 2, 8, 7, 6, 4, 9, 5, 10}
	res := Sort(data, func(a, b int) bool { return a > b })
	assert.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, res)

	// sorting in descending order
	res = Sort(data, func(a, b int) bool { return a < b })
	assert.Equal([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, res)
}
