package gogu

import (
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind_Index(t *testing.T) {
	assert := assert.New(t)

	input1 := []int{1, 2, 3, 4, 2, -2, -1, 10, 100}
	item1 := FindIndex(input1, func(v int) bool {
		return v == 2
	})
	assert.NotEqual(-1, item1)
	assert.Equal(1, item1)

	input2 := []string{"a", "b", "c", "d", "A", "B", "C", "D"}
	item2 := FindIndex(input2, func(v string) bool {
		return v == "A"
	})
	assert.NotEqual(-1, item2)

	item3 := FindIndex(input2, func(v string) bool {
		return v == "f"
	})
	assert.Equal(-1, item3)
}

func TestFind_LastIndex(t *testing.T) {
	assert := assert.New(t)

	input := []int{1, 2, 3, 4, 2, -2, -1, 10, 100}
	item := FindLastIndex(input, func(v int) bool {
		return v == 2
	})
	assert.NotEqual(-1, item)
	assert.Equal(4, item)
}

func TestFind_All(t *testing.T) {
	assert := assert.New(t)

	input1 := []int{1, 2, 3, 4, 2, -2, -1, 2}
	items1 := FindAll(input1, func(v int) bool {
		return v == 2
	})
	assert.NotEmpty(items1)
	assert.Equal(2, items1[1])
	assert.Equal(3, len(items1))

	input2 := []string{"foo", "bar", "baz"}
	items2 := FindAll(input2, func(v string) bool {
		return v == "foo2"
	})
	assert.Empty(items2)
}

func TestFind_IndexOf(t *testing.T) {
	input := []int{1, 2, 3, 4, 2, -2, -1, 2}

	assert := assert.New(t)
	assert.Equal(0, IndexOf(input, 1))
	assert.Equal(-1, IndexOf(input, 5))
}

func TestFind_LastIndexOf(t *testing.T) {
	input := []int{1, 2, -1, 4, 2, -2, -1, 2}

	assert := assert.New(t)
	assert.Equal(6, LastIndexOf(input, -1))
	assert.Equal(-1, IndexOf(input, 5))
}

func TestFind_Min(t *testing.T) {
	assert := assert.New(t)

	input1 := []int{1, 2, 3}
	assert.Equal(1, FindMin(input1))

	input2 := []float64{1.0, 0.001, -0.001}
	assert.Equal(-0.001, FindMin(input2))

	input3 := []string{"c", "b", "a"}
	assert.Equal("a", FindMin(input3))
}

func TestFind_MinBy(t *testing.T) {
	assert := assert.New(t)

	input1 := []float64{1.2, 1.4, 2.2, 4.8}
	assert.Equal(1.2, FindMinBy(input1, func(val float64) float64 {
		return math.Ceil(val)
	}))

	input2 := []string{"a", "b", "c"}
	assert.Equal("b", FindMinBy(input2, func(val string) string {
		return strings.Map(func(r rune) rune {
			// Because a string is effectively a slice of bytes
			// we are altering the underlying rune of codepoint "a"
			// to make sure that the resulted rune is greater than the original one.
			if string(r) == "a" {
				return r + r
			}
			return r
		}, val)
	}))
}
