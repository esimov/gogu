package gogu

import (
	"fmt"
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

func Example() {
	input := []int{1, 2, 3, 4, 2, -2, -1, 2}
	items := FindAll(input, func(v int) bool {
		return v == 2
	})
	fmt.Println(items)

	// Output:
	// map[1:2 4:2 7:2]
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

func TestFind_MinByKey(t *testing.T) {
	assert := assert.New(t)

	input1 := []map[string]int{
		{"number": 1},
		{"number": 2},
		{"number": 3},
		{"number": 10},
	}
	// Testing with wrong key.
	_, err := FindMinByKey(input1, "n")
	assert.Error(err)

	min1, _ := FindMinByKey(input1, "number")
	assert.Equal(1, min1)

	input2 := []map[int]string{
		{1: "number0"},
		{1: "number1"},
		{1: "number2"},
		{1: "number3"},
	}

	min2, _ := FindMinByKey(input2, 1)
	assert.Equal("number0", min2)
}

func TestFind_Max(t *testing.T) {
	assert := assert.New(t)

	input1 := []int{1, 2, 3}
	assert.Equal(3, FindMax(input1))

	input2 := []float64{1.0, 0.001, -0.001}
	assert.Equal(-0.001, FindMin(input2))

	input3 := []string{"a", "b", "c"}
	assert.Equal("c", FindMax(input3))
}

func TestFind_MaxBy(t *testing.T) {
	assert := assert.New(t)

	input1 := []float64{1.2, 1.4, 2.2, 4.8}
	assert.Equal(4.8, FindMaxBy(input1, func(val float64) float64 {
		return math.Ceil(val)
	}))

	input2 := []string{"a", "b", "c"}
	assert.Equal("b", FindMaxBy(input2, func(val string) string {
		return strings.Map(func(r rune) rune {
			// Normally "c" should be the biggest value, but we are altering
			// the rune of string "b" in case this is found.
			if string(r) == "b" {
				return r + r
			}
			return r
		}, val)
	}))
}

func TestFind_MaxByKey(t *testing.T) {
	assert := assert.New(t)

	input := []map[string]int{
		{"number": 1},
		{"number": 2},
		{"number": 3},
		{"number": 10},
	}
	// Testing with wrong key.
	_, err := FindMaxByKey(input, "n")
	assert.Error(err)

	max, _ := FindMaxByKey(input, "number")
	assert.Equal(10, max)
}

func TestFind_Nth(t *testing.T) {
	assert := assert.New(t)
	type numbers struct {
		min, max int
	}

	nth1, _ := Nth([]int{1, 2, 3, 4}, 1)
	assert.Equal(2, nth1)

	nth2, _ := Nth([]int{1, 2, 3, 4}, -2)
	assert.Equal(3, nth2)

	sample := []numbers{
		{min: 1, max: 10},
		{min: 2, max: 100},
	}

	_, err := Nth(sample, 2)
	assert.Error(err)

	nth3, _ := Nth(sample, 1)
	assert.Equal(numbers{min: 2, max: 100}, nth3)

	_, err = Nth(sample, -3)
	assert.Error(err)

	nth4, _ := Nth(sample, -2)
	assert.Equal(numbers{min: 1, max: 10}, nth4)
}
