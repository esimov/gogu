package gogu

import (
	"fmt"
	"strconv"
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
	fmt.Println(input2)
	assert.Equal([]string{"aa", "bb", "cc", "dd"}, input2)

	output3 := []string{}
	ForEachRight(input1, func(val int) {
		output3 = append(output3, strconv.Itoa(val))
	})
	assert.Equal([]string{"4", "3", "2", "1"}, output3)
}
