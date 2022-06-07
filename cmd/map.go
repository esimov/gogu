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

	fmt.Println("==================FilterMap")
	res2 := gogu.FilterMap[int, string](mp, func(v string) bool {
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
	res3 := gogu.Filter2DMap[string, int](usersMap, func(v int) bool {
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
	res4 := gogu.Filter2DMapSlice[string, int](usersSlice, func(v map[string]int) bool {
		return v["age"] > 20 && v["ranking"] < 5
	})
	fmt.Println(res4)

	fmt.Println("==================Invert")
	inverted := gogu.Invert(mp)
	fmt.Println(inverted)

	fmt.Println("==================MapEvery")
	input := map[string]int{"John": 2, "Doe": 1, "Fred": 3}
	every := gogu.MapEvery[string, int](input, func(v int) bool {
		return v > 1
	})
	fmt.Println(every)
}
