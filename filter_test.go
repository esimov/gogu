package gogu

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	assert := assert.New(t)

	// Obtain all the values which satisfies the callback function condition.
	items1 := []int{1, 2, 3, 4, 5, 10, 20, 30, 40, 50}
	assert.Equal([]int{10, 20, 30, 40, 50}, Filter(items1, func(val int) bool {
		return val >= 10
	}))

	items2 := []float64{12.2, 22.1, 10.01, 1, 20, 50}
	assert.Equal([]float64{1}, Filter(items2, func(val float64) bool {
		return val == 1.0
	}))

	items3 := []float64{2.0, 4.0, 6.0}
	assert.Equal([]float64{4.0, 6.0}, Filter(items3, func(val float64) bool {
		return math.Sqrt(val) >= 2
	}))

	// Reject all the values which doesn't satisfies the callback function condition.
	items4 := []int{1, 2, 3, 4, 5, 10, 20, 30, 40, 50}
	assert.Equal([]int{10, 20, 30, 40, 50}, Reject(items4, func(val int) bool {
		return val < 10
	}))

	assert.NotEmpty(Filter(items4, func(val int) bool {
		return val%2 == 0
	}))

	// Map filter test cases.
	items5 := map[int]string{1: "John", 2: "Doe", 3: "Fred"}
	assert.Equal(map[int]string{1: "John"}, FilterMap(items5, func(v string) bool {
		return v == "John"
	}))

	input1 := []map[string]int{
		{"bernie": 22},
		{"robert": 30},
	}
	expected1 := []map[string]int{
		{"robert": 30},
	}

	assert.Equal(expected1, FilterMapCollection(input1, func(val int) bool {
		return val > 22
	}))

	input2 := []map[string]map[string]int{
		{"bernie": {"age": 30, "ranking": 1}},
		{"robert": {"age": 20, "ranking": 5}},
	}
	expected := []map[string]map[string]int{
		{"bernie": {"age": 30, "ranking": 1}},
	}

	assert.Equal(expected, Filter2DMapCollection(input2, func(v map[string]int) bool {
		return v["age"] > 20 && v["ranking"] < 5
	}))
	assert.NotEqual(expected, Filter2DMapCollection(input2, func(v map[string]int) bool {
		return v["age"] > 20 && v["ranking"] > 1
	}))
}
