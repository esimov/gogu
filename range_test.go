package gogu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	assert := assert.New(t)

	// Testing range with 1 argument.
	// In this case the start argument defaults to 0 and the step argument to 1.
	range1, _ := Range(5)
	assert.Equal(range1, []int{0, 1, 2, 3, 4})

	// Testing range with 2 arguments.
	range2, _ := Range(1, 5)
	assert.Equal(range2, []int{1, 2, 3, 4})

	// Testing range with 3 arguments.
	range3, _ := Range(0, 2, 10)
	assert.Equal(range3, []int{0, 2, 4, 6, 8})

	// Testing range with more than 3 arguments. In this case it should return an error.
	_, err := Range(0, 1, 2, 3)
	assert.Error(err)

	// The end value should be greater than the start value.
	_, err = Range(12, 2, 10)
	assert.Error(err)

	// Testing range with negative values using 1 argument.
	range4, _ := Range(-4)
	assert.Equal(range4, []int{0, -1, -2, -3})

	// Testing range with negative values using 2 arguments.
	range5, _ := Range(-1, -4)
	assert.Equal(range5, []int{-1, -2, -3})

	// Testing range with negative values using 3 arguments.
	range6, _ := Range(0, -1, -4)
	assert.Equal(range6, []int{0, -1, -2, -3})

	range7, _ := Range[float64](0, 0.12, 0.9)
	assert.Equal(range7, []float64{0, 0.12, 0.24, 0.36, 0.48, 0.6, 0.72, 0.84})
}

func TestRangeRight(t *testing.T) {
	assert := assert.New(t)

	rangeR1, _ := RangeRight(5)
	assert.Equal(rangeR1, []int{4, 3, 2, 1, 0})

	rangeR2, _ := RangeRight(1, 2, 6)
	assert.Equal(rangeR2, []int{5, 3, 1})
}
