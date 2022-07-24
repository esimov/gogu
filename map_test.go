package gogu

import (
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_Keys(t *testing.T) {
	assert := assert.New(t)

	items1 := map[string]int{"a": 1, "b": 2, "c": 3}
	// retrieve keys
	assert.ElementsMatch([]string{"a", "b", "c"}, Keys(items1))

	keys := Keys(items1)
	sort.Strings(keys)
	assert.Equal([]string{"a", "b", "c"}, keys)

	// retrieve values
	assert.ElementsMatch([]int{1, 2, 3}, Values(items1))

	vals := Values(items1)
	sort.Ints(vals)
	assert.Equal([]int{1, 2, 3}, vals)

	assert.Len(MapKeys(items1, func(k string, v int) string {
		return k
	}), len(items1))

	expected1 := map[string]int{"a1": 1, "b2": 2, "c3": 3}
	expected2 := map[string]string{"a": "10", "b": "20", "c": "30"}

	// map keys
	assert.Equal(expected1, MapKeys(items1, func(k string, v int) string {
		return k + strconv.FormatInt(int64(v), 10)
	}))

	// map values
	assert.Equal(expected2, MapValues(items1, func(v int) string {
		return strconv.FormatInt(int64(v*10), 10)
	}))

	// find
	items2 := map[int]string{1: "John", 2: "Doe", 3: "Fred", 4: "John"}
	assert.Equal(map[int]string{1: "John"}, Find(items2, func(v string) bool {
		return v == "John"
	}))

	items3 := map[string]string{"one": "John", "two": "Doe", "three": "Fred", "four": "John"}
	// On sorted strings "four' is the first item.
	assert.Equal(map[string]string{"four": "John"}, Find(items3, func(v string) bool {
		return v == "John"
	}))
	assert.Len(Find(items2, func(v string) bool {
		return v == "Fred"
	}), 1)

	assert.Empty(Find(items2, func(v string) bool {
		return v == "Jane"
	}))

	// find key
	assert.Equal(2, FindKey(items2, func(v string) bool {
		return v == "Doe"
	}))
	assert.Empty(FindKey(items2, func(v string) bool {
		return v == "Jane"
	}))

	// find by key
	assert.Equal(map[int]string{1: "John"}, FindByKey(items2, func(v int) bool {
		return v == 1
	}))
}
