package torx

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_Keys(t *testing.T) {
	assert := assert.New(t)

	items := map[string]int{"a": 1, "b": 2, "c": 3}
	assert.ElementsMatch([]string{"a", "b", "c"}, Keys(items))

	keys := Keys(items)
	sort.Strings(keys)
	assert.Equal([]string{"a", "b", "c"}, keys)
}

func TestMap_Values(t *testing.T) {
	assert := assert.New(t)

	items := map[string]int{"a": 1, "b": 2, "c": 3}
	assert.ElementsMatch([]int{1, 2, 3}, Values(items))

	vals := Values(items)
	sort.Ints(vals)
	assert.Equal([]int{1, 2, 3}, vals)
}

func TestMap_MapKeys(t *testing.T) {
	assert := assert.New(t)

	items := map[string]int{"a": 1, "b": 2, "c": 3}
	assert.Len(MapKeys(items, func(k string, v int) string {
		return k
	}), len(items))

	expected := map[string]int{"a1": 1, "b2": 2, "c3": 3}
	assert.Equal(expected, MapKeys(items, func(k string, v int) string {
		return k + strconv.FormatInt(int64(v), 10)
	}))
}

func TestMap_MapValues(t *testing.T) {
	assert := assert.New(t)

	items := map[string]int{"a": 1, "b": 2, "c": 3}
	expected := map[string]string{"a": "10", "b": "20", "c": "30"}

	assert.Len(MapKeys(items, func(k string, v int) string {
		return k
	}), len(items))
	assert.Equal(expected, MapValues(items, func(v int) string {
		return strconv.FormatInt(int64(v*10), 10)
	}))
}

func TestMap_MapInvert(t *testing.T) {
	assert := assert.New(t)

	items := map[string]int{"a": 1, "b": 2, "c": 3}
	inv := map[int]string{1: "a", 2: "b", 3: "c"}

	assert.Equal(inv, Invert(items))

}
func TestMap_MapFunc(t *testing.T) {
	assert := assert.New(t)

	items1 := map[string]int{"a": 1, "b": 2, "c": 3}
	assert.Equal(true, MapEvery(items1, func(val int) bool {
		return val >= 1
	}))

	assert.Equal(true, MapSome(items1, func(val int) bool {
		return val > 2
	}))

	items2 := map[int]string{1: "John", 2: "Doe", 3: "Fred", 4: "John"}
	item := FindByKey(items2, func(v int) bool {
		return v == 1
	})
	assert.Equal(true, MapContains(item, "John"))

	expected := []int{2, 4, 6}
	assert.ElementsMatch(expected, MapCollection(items1, func(val int) int {
		return val * 2
	}))

	assert.Len(MapUnique(items2), 3)

	res2 := []string{"john", "doe", "fred"}
	assert.ElementsMatch(res2, MapCollection(MapUnique(items2), func(val string) string {
		return strings.ToLower(val)
	}))
}

func TestMap_Find(t *testing.T) {
	assert := assert.New(t)

	items1 := map[int]string{1: "John", 2: "Doe", 3: "Fred", 4: "John"}
	assert.Equal(map[int]string{1: "John"}, Find(items1, func(v string) bool {
		return v == "John"
	}))

	items2 := map[string]string{"one": "John", "two": "Doe", "three": "Fred", "four": "John"}
	// 'four' is the first item on the sorted strings.
	assert.Equal(map[string]string{"four": "John"}, Find(items2, func(v string) bool {
		return v == "John"
	}))
	assert.Len(Find(items2, func(v string) bool {
		return v == "Fred"
	}), 1)

	assert.Empty(Find(items2, func(v string) bool {
		return v == "Jane"
	}))
}

func TestMap_FindKey(t *testing.T) {
	assert := assert.New(t)

	items := map[int]string{1: "John", 2: "Doe", 3: "Fred", 4: "John"}
	assert.Equal(2, FindKey(items, func(v string) bool {
		return v == "Doe"
	}))
	assert.Empty(FindKey(items, func(v string) bool {
		return v == "Jane"
	}))
}

func TestMap_FindByKey(t *testing.T) {
	assert := assert.New(t)

	items1 := map[int]string{1: "John", 2: "Doe", 3: "Fred", 4: "John"}
	assert.Equal(map[int]string{1: "John"}, FindByKey(items1, func(v int) bool {
		return v == 1
	}))
}

func TestMap_Pluck(t *testing.T) {
	assert := assert.New(t)

	input := []map[string]string{
		{"name": "moe", "email": "moe@example.com"},
		{"name": "larry", "email": "larry@example.com"},
		{"name": "curly", "email": "curly@example.com"},
		{"name": "moly", "email": "moly@example.com"},
	}
	expected := []string{"moe", "larry", "curly", "moly"}
	assert.ElementsMatch(expected, Pluck(input, "name"))

	expected2 := []string{"Moe", "Larry", "Curly", "Moly"}
	assert.ElementsMatch(expected2, Map(Pluck(input, "name"), func(val string) string {
		return strings.ToUpper(string([]byte(val)[0])) + string([]byte(val)[1:])
	}))
	assert.NotEmpty(Pluck(input, "name"))
	assert.Empty(Pluck(input, "age"))
}

func TestMap_Pick(t *testing.T) {
	assert := assert.New(t)

	res1, err := Pick(map[string]any{"name": "moe", "age": 20, "active": true})
	assert.Error(err)
	assert.Empty(res1)

	expected := map[string]any{"name": "moe", "age": 20}
	res2, _ := Pick(map[string]any{"name": "moe", "age": 20, "active": true}, "name", "age")
	assert.Equal(expected, res2)

	assert.Equal(map[string]int{"b": 2, "c": 3}, PickBy(map[string]int{"aa": 1, "b": 2, "c": 3}, func(key string, val int) bool {
		return len(key) == 1
	}))
}

func TestMap_Omit(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(map[string]any{"active": false}, Omit(map[string]any{"name": "moe", "age": 40, "active": false}, "name", "age"))

	assert.Equal(map[string]int{"b": 2}, OmitBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(key string, val int) bool {
		return val%2 == 1
	}))
}

func Example_MapOmit() {
	res := Omit(map[string]any{"name": "moe", "age": 40, "active": false}, "name", "age")
	fmt.Println(res)

	// Output:
	// map[active:false]
}

func Example_MapOmitBy() {
	res := OmitBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(key string, val int) bool {
		return val%2 == 1
	})
	fmt.Println(res)

	// Output:
	// map[b:2]
}

func TestMap_Partition(t *testing.T) {
	assert := assert.New(t)

	input := []map[string]string{
		{"name": "moe", "email": "moe@example.com"},
		{"name": "larry", "email": "larry@example.com"},
		{"name": "curly", "email": "curly@example.com"},
		{"name": "moly", "email": "moly@example.com"},
	}

	res := PartitionMap(input, func(m map[string]string) bool {
		return len(m["name"]) == 3
	})
	assert.Len(res, 2)
	assert.Len(res[0], 1)
	assert.Len(res[1], 3)
	assert.Equal([]map[string]string{{"name": "moe", "email": "moe@example.com"}}, res[0])
}

func TestMap_SliceToMap(t *testing.T) {
	assert := assert.New(t)

	// slice to map
	sl1 := []string{"a", "b", "c"}
	sl2 := []int{1, 2, 3}
	assert.NotEmpty(SliceToMap(sl1, sl2))
	assert.Len(SliceToMap(sl1, sl2), 3)

	keys := Keys(SliceToMap(sl1, sl2))
	sort.Strings(keys)
	assert.Equal(sl1, keys)

	values := Values(SliceToMap(sl1, sl2))
	sort.Ints(values)
	assert.Equal(sl2, values)
}
