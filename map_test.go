package gogu

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_Keys(t *testing.T) {
	assert := assert.New(t)

	items1 := map[string]int{"a": 1, "b": 2, "c": 3}
	// retrieve keys
	assert.ElementsMatch([]string{"a", "b", "c"}, Keys(items1))
	// retrieve values
	assert.ElementsMatch([]int{1, 2, 3}, Values(items1))

	expected1 := map[string]int{"a1": 1, "b2": 2, "c3": 3}
	expected2 := map[string]string{"a": "10", "b": "20", "c": "30"}

	//map keys
	assert.Equal(expected1, MapKeys(items1, func(k string, v int) string {
		return k + strconv.FormatInt(int64(v), 10)
	}))

	// map values
	assert.Equal(expected2, MapValues(items1, func(v int) string {
		return strconv.FormatInt(int64(v*10), 10)
	}))
}
