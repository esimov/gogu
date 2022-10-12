package main

import (
	"fmt"
	"math"

	"github.com/esimov/torx"
)

func main() {

	fmt.Println("==================Sum")
	fmt.Println(torx.Sum([]int{1, 2, 3}))

	fmt.Println("==================SumBy")
	fmt.Println(torx.SumBy[string, int]([]string{"one", "two"}, func(elem string) int {
		return len(elem)
	}))

	fmt.Println("==================Mean")
	fmt.Println(torx.Mean([]int{4, 2, 8, 6}))

	ints := []int{2, 1, 4, 12, 8, 10, 22, 2, 10, 2, 13, 10, 4, 13}

	fmt.Println("==================Map")
	maps := torx.Map(ints, func(a int) int {
		return a * 2
	})
	fmt.Println(maps)

	fmt.Println("==================ForEach")
	torx.ForEach(ints, func(v int) {
		fmt.Printf("%v ", v)
	})
	fmt.Println()

	fmt.Println("==================ForEachRight")
	torx.ForEachRight(ints, func(v int) {
		fmt.Printf("%v ", v)
	})
	fmt.Println()

	fmt.Println("==================Reduce")
	reduce := torx.Reduce(ints, func(a, b int) int {
		return a + b
	}, 0)
	fmt.Println(reduce)

	fmt.Println("==================Reverse")
	fmt.Println(torx.Reverse(ints))

	fmt.Println("==================Unique")
	fmt.Println(torx.Unique[int](ints))

	fmt.Println("==================UniqueBy")
	fmt.Println(torx.UniqueBy[float64]([]float64{2.1, 1.2, 2.3}, func(v float64) float64 {
		return math.Floor(v)
	}))

	fmt.Println("==================Every")
	fmt.Println(torx.Every[int](ints, func(v int) bool {
		return v == 1
	}))

	fmt.Println("==================Some")
	fmt.Println(torx.Some[int](ints, func(v int) bool {
		return v >= 1
	}))

	fmt.Println("==================Partition")
	fmt.Println(torx.Partition[int](ints, func(v int) bool {
		return v == 2
	}))

	fmt.Println("==================Contains")
	fmt.Println(torx.Contains[int](ints, 10))

	fmt.Println("==================Duplicate")
	fmt.Println(torx.Duplicate[int](ints))

	fmt.Println("==================Duplicate With Index")
	fmt.Println(torx.DuplicateWithIndex[int](ints))

	fmt.Println("==================Duplicate Strings")
	strs := []string{"One", "Two", "Foo", "Bar", "Baz", "Foo", "Foo", "One"}
	fmt.Println(torx.Duplicate(strs))

	fmt.Println("==================Merge")
	fmt.Println(torx.Merge(ints, []int{2, 10, 4}, []int{2, 23, 2}))

	fmt.Println("==================Without")
	fmt.Println(torx.Without[int, int](ints, 2, 1, 12))

	fmt.Println("==================Difference")
	fmt.Println(torx.Difference[int]([]int{1, 2, 3, 4, 5}, []int{5, 2, 10}))

	fmt.Println("==================DifferenceBy")
	fmt.Println(torx.DifferenceBy[float64]([]float64{2.1, 1.2}, []float64{2.3, 3.4}, func(v float64) float64 {
		return math.Floor(v)
	}))

	fmt.Println("==================Flatten")
	sl1 := []any{[]any{1.0, 2.0, []any{3.0, []float64{4, 5, 6}}}, 7.0}
	fl, _ := torx.Flatten[float64](sl1)
	fmt.Println(fl)

	fmt.Println("==================Union")
	sl2 := []any{[]any{1, 2, []any{3, []int{4, 5, 6}}}, 7, []int{1, 2}, 3, []int{4, 7}, 10, 10}
	un, _ := torx.Union[int](sl2)
	fmt.Println(un)

	fmt.Println("==================Union Strings")
	str2 := []any{[]any{"One", "Two", []any{"Foo", []string{"Bar", "Baz", "Qux"}}}, "Foo", []string{"Foo", "Two"}, "Baz", "bar"}
	sl3, _ := torx.Union[string](str2)
	fmt.Println(sl3)

	fmt.Println("==================Intersection")
	in := torx.Intersection[int]([]int{1, 2, 3}, []int{101, 2, 1, 10}, []int{2, 1})
	fmt.Println(in)

	fmt.Println("==================IntersectionBy")
	fl4 := torx.IntersectionBy(func(v float64) float64 {
		return math.Floor(v)
	}, []float64{2.1, 1.2, 5.09}, []float64{2.3, 2.2, 3.04, 3.1, 4.8, 4.1})
	fmt.Println(fl4)

	fmt.Println("==================Chunk")
	fmt.Println(torx.Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 5))

	fmt.Println("==================Drop")
	fmt.Println(torx.Drop([]int{1, 2, 3}, 2))

	fmt.Println("==================DropWhile")
	fmt.Println(torx.DropWhile([]string{"AAA", "AA", "A", "AAAA"}, func(elem string) bool {
		return len(elem) > 2
	}))

	fmt.Println("==================DropRightWhile")
	fmt.Println(torx.DropRightWhile([]string{"AAA", "AA", "A", "AAAA"}, func(elem string) bool {
		return len(elem) > 2
	}))

	fmt.Println("==================Shuffle")
	fmt.Println(torx.Shuffle([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

	fmt.Println("==================GroupBy")
	fmt.Println(torx.GroupBy[float64, float64]([]float64{1.3, 2.1, 2.4}, func(val float64) float64 {
		return math.Floor(val)
	}))
	fmt.Println(torx.GroupBy[string, int]([]string{"one", "two", "three"}, func(val string) int {
		return len(val)
	}))

	fmt.Println("==================Zip")
	fmt.Println(torx.Zip[any]([]any{"one", "two", "three"}, []any{10, 20, 30}, []any{true, true, false}))

	fmt.Println("==================UnZip")
	fmt.Println(torx.Unzip[any]([]any{"one", 10}, []any{"two", 20}))

	fmt.Println("==================ToSlice")
	fmt.Println(torx.ToSlice[int](1, 2, 3, 4))
}
