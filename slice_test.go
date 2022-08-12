package gogu

import (
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_Sum(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(6, Sum([]int{1, 2, 3}))
	assert.Equal(12, SumBy([]int{1, 2, 3}, func(val int) int {
		return val * 2
	}))
	assert.Equal(6, SumBy([]string{"one", "two"}, func(val string) int {
		return len(val)
	}))
}

func TestSlice_Mean(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, Mean([]int{1, 2, 3}))
}

func TestSlice_Map(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{2, 4, 6}, Map([]int{1, 2, 3}, func(val int) int {
		return val * 2
	}))
	assert.Len(Map([]int{2, 4}, func(val int) int {
		return val * val
	}), 2)
}

func TestSlice_ForEach(t *testing.T) {
	assert := assert.New(t)

	idx := 0
	input1 := []int{1, 2, 3, 4}
	output1 := make([]int, 4)

	ForEach(input1, func(val int) {
		output1[idx] = val
		idx++
	})
	assert.Equal(output1, input1)
	assert.IsIncreasing(output1)

	idx = 0
	input2 := []string{"a", "b", "c", "d"}
	output2 := make([]string, len(input2)-1)

	ForEach(input2, func(val string) {
		if idx != len(input1)-1 {
			output2[idx] = val
		}
		idx++
	})

	assert.Equal([]string{"a", "b", "c"}, output2)
	assert.Len(output2, 3)

	idx = 0
	ForEach(input2, func(val string) {
		input2[idx] = val + val
		idx++
	})
	assert.Equal([]string{"aa", "bb", "cc", "dd"}, input2)

	output3 := []string{}
	ForEachRight(input1, func(val int) {
		output3 = append(output3, strconv.Itoa(val))
	})
	assert.Equal([]string{"4", "3", "2", "1"}, output3)
}

func TestSlice_Reduce(t *testing.T) {
	assert := assert.New(t)

	input1 := []int{1, 2, 3, 4}
	assert.Equal(10, Reduce(input1, func(a, b int) int {
		return a + b
	}, 0))

	input2 := []string{"a", "b", "c", "d"}
	assert.Equal("abcd", Reduce(input2, func(a, b string) string {
		return b + a
	}, ""))

	res := Reduce(input2, func(a, b string) string {
		return a + b
	}, "")
	res1 := []byte(res)
	sort.Slice(res1, func(i, j int) bool { return res[i] < res[j] })

	assert.Equal("abcd", string(res1))
}

func TestSlice_Reverse(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{4, 3, 2, 1}, Reverse([]int{1, 2, 3, 4}))
	assert.Equal([]string{"a", "b", "c"}, Reverse([]string{"c", "b", "a"}))

	assert.Equal("abcd", Reduce(Reverse([]string{"a", "b", "c", "d"}), func(a, b string) string {
		return a + b
	}, ""))

	assert.NotEqual("abcd", Reverse([]string{"a", "b", "c", "d"}))
}

func TestSlice_Unique(t *testing.T) {
	assert := assert.New(t)

	input := []int{1, 2, 4, 3, 1, 4, 5}
	res := Unique(input)

	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	assert.Equal([]int{1, 2, 3, 4, 5}, res)

	assert.Equal([]float64{2.1, 1.2}, UniqueBy([]float64{2.1, 1.2, 2.3}, func(v float64) float64 {
		return math.Floor(v)
	}))

	assert.Equal([]string{"a", "b", "c"}, UniqueBy([]string{"a", "b", "c", "B", "c", "A"}, func(v string) string {
		return strings.ToUpper(v)
	}))
}

func TestSlice_Every(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, Every([]int{2, 4, 6, 8, 10}, func(val int) bool {
		return val%2 == 0
	}))
	assert.NotEqual(true, Every([]int{-1, -2, 6, 8, 10}, func(val int) bool {
		return val > 0
	}))
	assert.Equal(false, Every([]any{"1", 1, 10, false}, func(val any) bool {
		return reflect.TypeOf(val).Kind() == reflect.Int
	}))
	assert.Equal(true, Every([]string{"1", "2", "3", "4"}, func(val string) bool {
		v, _ := strconv.Atoi(val)
		return reflect.TypeOf(v).Kind() == reflect.Int
	}))
}

func TestSlice_Some(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, Some([]int{1, 2, 3, 4, 5, 6}, func(val int) bool {
		return val%2 == 0
	}))
	assert.Equal(false, Some([]int{1, 3, 5}, func(val int) bool {
		return val%2 == 0
	}))
	assert.Equal(true, Some([]string{"1", "2", "3", "a"}, func(val string) bool {
		v, _ := strconv.Atoi(val)
		return reflect.TypeOf(v).Kind() == reflect.Int
	}))

	assert.Equal(false, Some([]string{"a", "b", "c"}, func(val string) bool {
		return reflect.TypeOf(val).Kind() == reflect.Int
	}))
}

func TestSlice_Partition(t *testing.T) {
	assert := assert.New(t)

	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Len(Partition(input, func(val int) bool {
		return val >= 5
	}), 2)

	res1 := Partition(input, func(val int) bool {
		return val < 5
	})
	assert.Equal([]int{0, 1, 2, 3, 4}, res1[0])

	res2 := Partition(input, func(val int) bool {
		return val < 0
	})
	assert.Empty(res2[0])
	assert.NotEmpty(res2[1])
}

func TestSlice_Contains(t *testing.T) {
	assert := assert.New(t)

	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Empty(Contains(input, -1))
	assert.Equal(true, Contains(input, 0))
	assert.NotEqual(true, Contains(input, 100))
}

func TestSlice_Duplicate(t *testing.T) {
	assert := assert.New(t)

	input1 := []int{-1, -1, 0, 1, 2, 3, 2, 5, 1, 6}
	assert.NotEmpty(Duplicate(input1))
	assert.Len(Duplicate(input1), 3)
	assert.ElementsMatch([]int{-1, 1, 2}, Duplicate(input1))

	input2 := []string{"One", "Two", "Three", "two", "One"}
	assert.ElementsMatch([]string{"One"}, Duplicate(input2))
	assert.ElementsMatch([]string{"one", "two"}, Duplicate(Map(input2, func(val string) string {
		return strings.ToLower(val)
	})))

	assert.Len(DuplicateWithIndex(input1), 3)
	res := DuplicateWithIndex(input1)

	indices := make([]int, 0, len(input1))
	for k := range res {
		indices = append(indices, k)
	}
	assert.ElementsMatch([]int{-1, 1, 2}, indices)
}

func TestSlice_Merge(t *testing.T) {
	assert := assert.New(t)

	sl1 := []int{1, 2, 3, 4}
	sl2 := []int{5, 6, 7, 8}

	assert.Len(Merge(sl1, sl2), len(sl1)+len(sl2))
	assert.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8}, Merge(sl1, sl2))
	assert.Equal([]int{1, 2, 3}, Merge([]int{1}, []int{2}, []int{3}))
}

func TestSlice_Flatten(t *testing.T) {
	assert := assert.New(t)

	input1 := []any{[]float64{1.0, 2.0}, 1.1}
	result1, err := Flatten[float64](input1)
	assert.Equal([]float64{1.0, 2.0, 1.1}, result1)
	assert.NotNil(result1)
	assert.NoError(err)
	assert.Len(result1, 3)

	input2 := []any{[]float32{1.0, 2.0}, 3.0}
	result2, err := Flatten[float32](input2)
	assert.Error(err)
	assert.Nil(result2) // result is nil, because the last element in the slice is of type float64

	input3 := []string{"a", "b", "c"}
	result3, err := Flatten[string](Merge(input3, []string{"d", "e"}))
	assert.Equal([]string{"a", "b", "c", "d", "e"}, result3)
	assert.NotNil(result3)
	assert.NoError(err)

	input4 := []any{[]int{1, 2, 3}, []any{[]int{4}, 5}}
	result4, _ := Flatten[int](input4)
	assert.Equal([]int{1, 2, 3, 4, 5}, result4)

	int1 := []int{1, 2}
	res1 := Map([]string{"3", "4"}, func(val string) int {
		res, _ := strconv.Atoi(val)
		return res
	})
	result5, err := Flatten[int]([]any{int1, res1})
	assert.Equal([]int{1, 2, 3, 4}, result5)
	assert.NoError(err)
}

func TestSlice_Union(t *testing.T) {
	assert := assert.New(t)

	input1 := []any{[]any{1, 2, []any{3, []int{4, 5, 6}}}, 7, []int{1, 2}, 3, []int{4, 7}, 8, 9, 9}
	result1, err := Union[int](input1)
	assert.Len(result1, 9)
	assert.NoError(err)
	assert.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, result1)

	input2 := []any{[]any{"One", "Two", []any{"Foo", []string{"Bar", "Baz", "Qux"}}}, "Foo", []string{"Foo", "Two"}, "Baz", "bar"}
	result2, err := Union[string](input2)
	assert.Len(result2, 7)
	assert.NoError(err)
	assert.Equal([]string{"One", "Two", "Foo", "Bar", "Baz", "Qux", "bar"}, result2)

	resMap := Map(result2, func(val string) string {
		return strings.ToLower(val)
	})
	result3, err := Union[string](resMap)
	assert.Equal([]string{"one", "two", "foo", "bar", "baz", "qux"}, result3)
	assert.NoError(err)
}

func TestSlice_Intersection(t *testing.T) {
	assert := assert.New(t)

	result1 := Intersection([]int{1, 2, 4}, []int{0, 2, 1}, []int{2, 1, -2})
	assert.Equal([]int{1, 2}, result1)

	result2 := Intersection([]int{-1, 0}, []int{2, 3})
	assert.Empty(result2)
	assert.Equal([]int{}, result2)

	result3 := Intersection([]int{0, 1, 2}, []int{2, 0, 1}, []int{2, 1, 0})
	assert.Equal([]int{0, 1, 2}, result3)

	result4 := Intersection([]string{"a", "b"}, []string{"a", "a", "a"}, []string{"b", "a", "e"})
	assert.Equal([]string{"a"}, result4)
}

func TestSlice_IntersectionBy(t *testing.T) {
	assert := assert.New(t)

	result1 := IntersectionBy(func(v float64) float64 {
		return math.Floor(v)
	}, []float64{2.1, 1.2}, []float64{2.3, 3.4}, []float64{1.0, 2.3})
	assert.Equal([]float64{2.1}, result1)

	result2 := IntersectionBy(func(v int) int {
		return v % 2
	}, []int{1, 2}, []int{2, 1})
	assert.Equal([]int{1, 2}, result2)

	result3 := IntersectionBy(func(v float64) float64 {
		return math.Floor(v)
	}, []float64{1.1, 2.0, 3.2}, []float64{4.0})
	assert.Equal([]float64{}, result3)
}

func TestSlice_Without(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{3}, Without[int, int]([]int{2, 1, 2, 3}, 1, 2))
	assert.Equal([]int{1, 2}, Without[int, int]([]int{1, 2, 3, 4}, 3, 4))
	assert.Equal([]int{1, 2}, Without[int, int]([]int{0, 1, 2, 3, 4, 5}, 0, 3, 4, 5))
	assert.Equal([]float64{1.0, 2.2}, Without[float64, float64]([]float64{1.0, 2.2, 3.0, 4.2}, 3.0, 4.2))

	assert.Empty(Without[int, int]([]int{}, 1, 2, 3, 4))
	assert.Empty(Without[int, int]([]int{0, 1, 2}, 0, 1, 2))
	assert.Empty(Without[int, int]([]int{}, 0, 1, 2))
	assert.Empty(Without[int, int]([]int{}))
}

func TestSlice_Difference(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{1, 3}, Difference([]int{1, 2, 3, 4}, []int{2, 4}))
	assert.Equal([]int{1, 2, 3}, Difference([]int{1, 2, 3, 4}, []int{4, 5, 6, 7}))
	assert.Equal([]int{1, 2, 3, 4}, Difference([]int{1, 2, 3, 4}, []int{}))

	assert.Empty(Difference([]int{}, []int{1, 2, 3, 4}))
	assert.Empty(Difference([]int{}, []int{-1}))
}
