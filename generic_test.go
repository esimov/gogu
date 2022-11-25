package torx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Compare compares two values using as comparator the callback function argument.
func TestCompare(t *testing.T) {
	assert := assert.New(t)
	res := Compare(1, 2, func(a, b int) bool {
		return a < b
	})
	assert.Equal(1, res)
}

func Example_compare() {
	res1 := Compare(1, 2, func(a, b int) bool {
		return a < b
	})
	fmt.Println(res1)

	res2 := Compare("a", "b", func(a, b string) bool {
		return a > b
	})
	fmt.Println(res2)

	// Output:
	// 1
	// -1
}

// Equal checks if two values are equal.
func TestEqual(t *testing.T) {
	assert := assert.New(t)

	assert.True(Equal(1, 1))
	assert.False(Equal("a", "b"))
	assert.False(Equal("a", "A"))
}

func TestLess(t *testing.T) {
	assert := assert.New(t)

	assert.True(Less(1, 2))
	assert.False(Less("b", "a"))
}
