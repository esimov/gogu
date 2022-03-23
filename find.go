package main

import "fmt"

func main() {
	ints := []int{12, 23, 1, 643, 99, 2, 2}
	// strings := []string{"Hello", "Bernie"}

	// var users = map[string]int{
	// 	"barney": 36,
	// 	"fred": 40,
	// 	""
	// 	'barney':  { 'age': 36, 'active': true },
	// 	'fred':    { 'age': 40, 'active': false },
	// 	'pebbles': { 'age': 1,  'active': true }
	//       };

	fmt.Println("==================FindAll")
	res1 := FindAll(ints, func(v int) bool {
		return v == 2
	})
	fmt.Println(res1)

	fmt.Println("==================FindAllFromLast")
	res2 := FindAllFromLast(ints, func(v int) bool {
		return v == 12
	})
	fmt.Println(res2)

	fmt.Println("==================FindIndex")
	res3 := FindIndex(ints, func(v int) bool {
		return v == 23
	})
	fmt.Println(res3)

	fmt.Println("==================FindLastIndex")
	res4 := FindLastIndex(ints, func(v int) bool {
		return v == 23
	})
	fmt.Println(res4)

	mp := []map[string]map[string]int{
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

	fmt.Println(mp)
	fmt.Println("==================FilterMap")
	res5 := Filter2DMap[string, int](mp, func(v map[string]int) bool {
		return v["age"] >= 10
	})
	fmt.Println(res5)
}

func FindAll[T any](s []T, fn func(T) bool) map[int]T {
	m := make(map[int]T, len(s))

	for k, v := range s {
		if fn(v) {
			m[k] = v
		}
	}

	return m
}

func FindAllFromLast[T any](s []T, fn func(T) bool) map[int]T {
	m := make(map[int]T, len(s))

	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		if fn(s[i]) {
			m[i] = s[j]
		}
	}

	return m
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

func FindIndex[T any](s []T, fn func(T) bool) int {
	for k, v := range s {
		if fn(v) {
			return k
		}
	}
	return -1
}

func FindLastIndex[T any](s []T, fn func(T) bool) int {
	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		if fn(s[j]) {
			return i
		}
	}
	return -1
}
