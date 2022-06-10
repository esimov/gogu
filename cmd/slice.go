package main

import (
	"fmt"
	"math"

	"github.com/esimov/gogu"
)

func main() {

	fmt.Println("==================Sum")
	fmt.Println(gogu.Sum([]int{1, 2, 3}))

	fmt.Println("==================Mean")
	fmt.Println(gogu.Mean([]int{4, 2, 8, 6}))

	ints := []int{2, 1, 4, 12, 8, 10, 22, 2, 10, 2, 13, 10, 4, 13}

	fmt.Println("==================Map")
	maps := gogu.Map(ints, func(a int) int {
		return a * 2
	})
	fmt.Println(maps)

	fmt.Println("==================ForEach")
	gogu.ForEach(ints, func(v int) {
		fmt.Printf("%v ", v)
	})
	fmt.Println()

	fmt.Println("==================ForEachRight")
	gogu.ForEachRight(ints, func(v int) {
		fmt.Printf("%v ", v)
	})
	fmt.Println()

	fmt.Println("==================Reduce")
	reduce := gogu.Reduce(ints, func(a, b int) int {
		return a + b
	}, 0)
	fmt.Println(reduce)

	fmt.Println("==================Reverse")
	fmt.Println(gogu.Reverse(ints))

	fmt.Println("==================Unique")
	fmt.Println(gogu.Unique[int](ints))

	fmt.Println("==================UniqueBy")
	fmt.Println(gogu.UniqueBy[float64]([]float64{2.1, 1.2, 2.3}, func(v float64) float64 {
		return math.Floor(v)
	}))

	fmt.Println("==================Every")
	fmt.Println(gogu.Every[int](ints, func(v int) bool {
		return v == 1
	}))

	fmt.Println("==================Some")
	fmt.Println(gogu.Some[int](ints, func(v int) bool {
		return v >= 1
	}))

	fmt.Println("==================Partition")
	fmt.Println(gogu.Partition[int](ints, func(v int) bool {
		return v == 2
	}))

	fmt.Println("==================Contains")
	fmt.Println(gogu.Contains[int](ints, 10))

	fmt.Println("==================Duplicate")
	fmt.Println(gogu.Duplicate[int](ints))

	fmt.Println("==================Duplicate With Index")
	fmt.Println(gogu.DuplicateWithIndex[int](ints))

	fmt.Println("==================Duplicate Strings")
	strs := []string{"One", "Two", "Foo", "Bar", "Baz", "Foo", "Foo", "One"}
	fmt.Println(gogu.Duplicate(strs))

	fmt.Println("==================Merge")
	fmt.Println(gogu.Merge(ints, []int{2, 10, 4}, []int{2, 23, 2}))

	fmt.Println("==================Without")
	fmt.Println(gogu.Without[int, int](ints, 2, 1, 12))

	fmt.Println("==================Difference")
	fmt.Println(gogu.Difference[int]([]int{1, 2, 3, 4, 5}, []int{5, 2, 10}))

	fmt.Println("==================DifferenceBy")
	fmt.Println(gogu.DifferenceBy[float64]([]float64{2.1, 1.2}, []float64{2.3, 3.4}, func(v float64) float64 {
		return math.Floor(v)
	}))

	fmt.Println("==================Flatten")
	sl1 := []any{[]any{1.0, 2.0, []any{3.0, []float64{4, 5, 6}}}, 7.0}
	fl, _ := gogu.Flatten[float64](sl1)
	fmt.Println(fl)

	fmt.Println("==================Union")
	sl2 := []any{[]any{1, 2, []any{3, []int{4, 5, 6}}}, 7, []int{1, 2}, 3, []int{4, 7}, 10, 10}
	un, _ := gogu.Union[int](sl2)
	fmt.Println(un)

	fmt.Println("==================Union Strings")
	str2 := []any{[]any{"One", "Two", []any{"Foo", []string{"Bar", "Baz", "Qux"}}}, "Foo", []string{"Foo", "Two"}, "Baz", "bar"}
	sl3, _ := gogu.Union[string](str2)
	fmt.Println(sl3)

	fmt.Println("==================Intersection")
	sl4 := []any{[]int{1, 2, 3}, []int{101, 2, 1, 10}, []int{2, 1}}
	in, _ := gogu.Intersection[int](sl4)
	fmt.Println(in)

	fmt.Println("==================IntersectionBy")
	fl4, _ := gogu.IntersectionBy(func(v float64) float64 {
		return math.Floor(v)
	}, []float64{2.1, 1.2, 5.09}, []float64{2.3, 2.2, 3.04, 3.1, 4.8, 4.1})
	fmt.Println(fl4)

	fmt.Println("==================Chunk")
	fmt.Println(gogu.Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 5))
}
