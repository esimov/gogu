package main

import (
	"fmt"
	"strconv"
)

func main() {
	mapVals := map[int]int64{1: 1, 2: 2, 3: 3}
	mapKeys := map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Println("==================Keys")
	fmt.Println(Keys(mapKeys))

	fmt.Println("==================Values")
	fmt.Println(Values(mapKeys))

	fmt.Println("==================MapValues")
	newVals := MapValues[int, int64, string](mapVals, func(v int64) string {
		v = v * 10
		return strconv.FormatInt(v, 10)
	})
	fmt.Println(newVals)

	fmt.Println("==================MapKeys")
	newKeys := MapKeys[string, int, string](mapKeys, func(k string, v int) string {
		return k + strconv.Itoa(v)
	})
	fmt.Println(newKeys)

	mp := map[int]string{1: "John", 2: "Doe", 3: "Fred"}

	fmt.Println("==================FindKey")
	res1 := FindKey[int, string](mp, func(v string) bool {
		return v == "John"
	})
	fmt.Println(res1)

	fmt.Println("==================FilterMap")
	res2 := FilterMap[int, string](mp, func(v string) bool {
		return v == "John"
	})
	fmt.Println(res2)

	usersMap := map[string]map[string]int{
		"bernie": {
			"age":     30,
			"ranking": 1,
		},

		"robert": {
			"age":     20,
			"ranking": 5,
		},
	}

	fmt.Println("==================Filter2DMap")
	res3 := Filter2DMap[string, int](usersMap, func(v int) bool {
		return v > 20
	})
	fmt.Println(res3)

	usersSlice := []map[string]map[string]int{
		{
			"bernie": {
				"age":     30,
				"ranking": 1,
			},
		},
		{
			"robert": {
				"age":     20,
				"ranking": 5,
			},
		},
	}

	fmt.Println("==================Filter2DMapSlice")
	res4 := Filter2DMapSlice[string, int](usersSlice, func(v map[string]int) bool {
		return v["age"] > 20 && v["ranking"] < 5
	})
	fmt.Println(res4)

	fmt.Println("==================Invert")
	inverted := Invert(mp)
	fmt.Println(inverted)
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))

	idx := 0
	for k, _ := range m {
		keys[idx] = k
		idx++
	}

	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))

	idx := 0
	for _, v := range m {
		values[idx] = v
		idx++
	}

	return values
}

func MapValues[K comparable, V, R any](m map[K]V, fn func(V) R) map[K]R {
	newMap := map[K]R{}

	for k, v := range m {
		newMap[k] = fn(v)
	}

	return newMap
}

func MapKeys[K comparable, V any, R comparable](m map[K]V, fn func(K, V) R) map[R]V {
	newMap := map[R]V{}

	for k, v := range m {
		newMap[fn(k, v)] = v
	}

	return newMap
}

func FindKey[K comparable, V any](m map[K]V, fn func(V) bool) K {
	var res K
	for k, v := range m {
		if fn(v) {
			res = k
			break
		}
	}
	return res
}

func FilterMap[K comparable, V any](m map[K]V, fn func(V) bool) map[K]V {
	filtered := map[K]V{}

	for k, v := range m {
		if fn(v) {
			filtered[k] = v
		}
	}

	return filtered
}

func Filter2DMap[K comparable, V any](m map[K]map[K]V, fn func(V) bool) map[K]map[K]V {
	filtered := map[K]map[K]V{}

	for k, v := range m {
		for _, v1 := range v {
			if fn(v1) {
				filtered[k] = v
			}
		}
	}

	return filtered
}

// TODO consider to remove
func Filter2DMapSlice[K comparable, V any](mapSlice []map[K]map[K]V, fn func(map[K]V) bool) []map[K]map[K]V {
	filtered := []map[K]map[K]V{}

	for _, s := range mapSlice {
		for _, v := range s {
			if fn(v) {
				filtered = append(filtered, s)
			}
		}
	}

	return filtered
}

// Returns a copy of the map where the keys have become the values and the values the keys.
// For this to work, all of your map's values should be unique.
func Invert[K, V comparable](m map[K]V) map[V]K {
	inverted := map[V]K{}
	keys := Keys(m)

	for i := 0; i < len(keys); i++ {
		inverted[m[keys[i]]] = keys[i]
	}

	return inverted
}
