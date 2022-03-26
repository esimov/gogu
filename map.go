package main

import (
	"fmt"
	"strconv"
)

func main() {
	mapVals := map[int]int64{1: 1, 2: 2, 3: 3}
	mapKeys := map[string]int{"a": 1, "b": 2, "c": 3}

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

	fmt.Println("==================FilterMap")
	mp := []map[string]int{{"user": 1}}

	res2 := FilterMap[string, int](mp, func(v int) bool {
		return v < 10
	})
	fmt.Println(res2)

	fmt.Println("==================Keys")
	fmt.Println(Keys(mapKeys))

	fmt.Println("==================Values")
	fmt.Println(Values(mapKeys))

	fmt.Println("==================Filter2DMap")
	users := []map[string]map[string]int{
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

	fmt.Println(users)
	res3 := Filter2DMap[string, int](users, func(v map[string]int) bool {
		return v["age"] >= 10
	})
	fmt.Println(res3)
}

func Keys[K comparable, V any](s map[K]V) []K {
	keys := make([]K, len(s))

	idx := 0
	for k, _ := range s {
		keys[idx] = k
		idx++
	}

	return keys
}

func Values[K comparable, V any](s map[K]V) []V {
	values := make([]V, len(s))

	idx := 0
	for _, v := range s {
		values[idx] = v
		idx++
	}

	return values
}

func MapValues[K comparable, V, R any](s map[K]V, fn func(V) R) map[K]R {
	newMap := map[K]R{}

	for k, v := range s {
		newMap[k] = fn(v)
	}

	return newMap
}

func MapKeys[K comparable, V any, R comparable](s map[K]V, fn func(K, V) R) map[R]V {
	newMap := map[R]V{}

	for k, v := range s {
		newMap[fn(k, v)] = v
	}

	return newMap
}

func FilterMap[K comparable, V any](s []map[K]V, fn func(V) bool) []map[K]V {
	filtered := []map[K]V{}

	for _, v := range s {
		for _, v1 := range v {
			if fn(v1) {
				filtered = append(filtered, v)
			}
		}
	}

	return filtered
}

func Filter2DMap[K comparable, V any](s []map[K]map[K]V, fn func(map[K]V) bool) []map[K]map[K]V {
	filtered := []map[K]map[K]V{}

	for _, v := range s {
		for _, v1 := range v {
			if fn(v1) {
				filtered = append(filtered, v)
			}
		}
	}

	return filtered
}
