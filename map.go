package main

import (
	"fmt"
	"strconv"
)

func main() {
	mapVals := map[int]int64{1: 1, 2: 2, 3: 3}

	fmt.Println("==================MapValues")
	newVals := MapValues[int, int64, string](mapVals, func(v int64) string {
		v = v * 10
		return strconv.FormatInt(v, 10)
	})
	fmt.Println(newVals)

	fmt.Println("==================MapKeys")
	mapKeys := map[string]int{"a": 1, "b": 2, "c": 3}
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

func MapValues[K comparable, V, R any](s map[K]V, fn func(V) R) map[K]R {
	newMap := map[K]R{}

	for idx, v := range s {
		newMap[idx] = fn(v)
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

func FilterMap[T1 comparable, T2 any](s []map[T1]T2, fn func(T2) bool) []map[T1]T2 {
	filtered := []map[T1]T2{}

	for _, v := range s {
		for _, v1 := range v {
			if fn(v1) {
				filtered = append(filtered, v)
			}
		}
	}

	return filtered
}

func Filter2DMap[T1 comparable, T2 any](s []map[T1]map[T1]T2, fn func(map[T1]T2) bool) []map[T1]map[T1]T2 {
	filtered := []map[T1]map[T1]T2{}

	for _, v := range s {
		for _, v1 := range v {
			if fn(v1) {
				filtered = append(filtered, v)
			}
		}
	}

	return filtered
}
