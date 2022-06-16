package main

import (
	"fmt"
	"strconv"

	"github.com/esimov/gogu"
)

func main() {
	mapVals := map[int]int64{1: 1, 2: 2, 3: 3}
	mapKeys := map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Println("==================Keys")
	fmt.Println(gogu.Keys(mapKeys))

	fmt.Println("==================Values")
	fmt.Println(gogu.Values(mapKeys))

	fmt.Println("==================MapValues")
	newVals := gogu.MapValues[int, int64, string](mapVals, func(v int64) string {
		v = v * 10
		return strconv.FormatInt(v, 10)
	})
	fmt.Println(newVals)

	fmt.Println("==================MapKeys")
	newKeys := gogu.MapKeys[string, int, string](mapKeys, func(k string, v int) string {
		return k + strconv.Itoa(v)
	})
	fmt.Println(newKeys)

	mp := map[int]string{1: "John", 2: "Doe", 3: "Fred"}

	fmt.Println("==================Find")
	res0 := gogu.Find[int, string](mp, func(v string) bool {
		return v == "John"
	})
	fmt.Println(res0)

	fmt.Println("==================FindKey")
	res1 := gogu.FindKey[int, string](mp, func(v string) bool {
		return v == "John"
	})
	fmt.Println(res1)

	fmt.Println("==================Invert")
	inverted := gogu.Invert(mp)
	fmt.Println(inverted)

	input := map[string]int{"John": 2, "Doe": 1, "Fred": 3}

	fmt.Println("==================MapEvery")
	every := gogu.MapEvery[string, int](input, func(v int) bool {
		return v > 1
	})
	fmt.Println(every)

	fmt.Println("==================MapSome")
	some := gogu.MapSome[string, int](input, func(v int) bool {
		return v > 1
	})
	fmt.Println(some)

	fmt.Println("==================MapContains")
	contains := gogu.MapContains[string, int](input, 3)
	fmt.Println(contains)

	fmt.Println("==================MapCollection")
	mapcol := map[string]int{"one": 1, "two": 2, "three": 3}
	col := gogu.MapCollection[string, int](mapcol, func(val int) int {
		return val * 2
	})
	fmt.Println(col)

	fmt.Println("==================Pluck")
	in := []map[string]any{
		{"name": "moe", "age": 40, "active": false},
		{"name": "larry", "age": 50, "active": true},
		{"name": "curly", "age": 60, "active": false},
		{"name": "moly", "age": 60, "active": false},
	}
	pl := gogu.Pluck[string, any](in, "name")
	fmt.Println(pl)

	fmt.Println("==================Pick")
	fmt.Println(gogu.Pick(map[string]any{"name": "moe", "age": 40, "active": false}, "name", "age"))

	fmt.Println("==================PickBy")
	fmt.Println(gogu.PickBy(map[string]int{"aa": 1, "b": 2, "c": 3}, func(key string, val int) bool {
		return len(key) == 1
	}))

	fmt.Println("==================Omit")
	fmt.Println(gogu.Omit(map[string]any{"name": "moe", "age": 40, "active": false}, "name", "age"))

	fmt.Println("==================OmitBy")
	fmt.Println(gogu.OmitBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(key string, val int) bool {
		return val%2 == 1
	}))

	fmt.Println("==================PartitionMap")
	pm := gogu.PartitionMap[string, any](in, func(m map[string]any) bool {
		return m["age"] == false
	})
	fmt.Println(pm)

	fmt.Println("==================SliceToMap")
	fmt.Println(gogu.SliceToMap[string, int]([]string{"a", "b", "c"}, []int{1, 2, 3}))
}
