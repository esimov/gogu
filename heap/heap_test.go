package heap

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap_MinHeap(t *testing.T) {
	assert := assert.New(t)

	heap := NewHeap(func(a, b int) bool { return a < b })
	assert.Empty(heap.Size())
	assert.True(heap.IsEmpty())

	heap.Push(10)
	assert.Equal(len(heap.GetValues()), heap.Size())
	assert.Equal(1, heap.Size())
	assert.Equal(10, heap.Pop())
	assert.True(heap.IsEmpty())

	values := []int{2, 5, 1, 4, 3}
	for _, v := range values {
		heap.Push(v)
	}
	assert.Equal([]int{1, 3, 2, 5, 4}, heap.GetValues())

	heap.Push(0)
	assert.Equal([]int{0, 3, 1, 5, 4, 2}, heap.GetValues())

	heap.Clear()
	assert.Empty(heap.GetValues())

	v := heap.Pop()
	assert.Equal(0, v)

	heap.Push(10, 4, 2, 5, 3)
	expected := []int{2, 3, 4, 10, 5}
	assert.Equal(expected, heap.GetValues())

	for range heap.GetValues() {
		heap.Pop()
	}
	assert.Empty(heap.Size())
}

func Example_minHeap() {
	heap := NewHeap(func(a, b int) bool { return a < b })
	fmt.Println(heap.IsEmpty())

	heap.Push(10)
	fmt.Println(heap.Size())
	heap.Pop()
	fmt.Println(heap.IsEmpty())

	values := []int{2, 5, 1, 4, 3}
	for _, v := range values {
		heap.Push(v)
	}
	fmt.Println(heap.GetValues())

	heap.Push(0)
	fmt.Println(heap.GetValues())

	heap.Clear()
	fmt.Println(heap.GetValues())

	fmt.Println(heap.Pop())

	heap.Push(10, 4, 2, 5, 3)
	for range heap.GetValues() {
		heap.Pop()
	}
	fmt.Println(heap.Size())

	// Output:
	// true
	// 1
	// true
	// [1 3 2 5 4]
	// [0 3 1 5 4 2]
	// []
	// 0
	// 0
}

func TestHeap_MaxHeap(t *testing.T) {
	assert := assert.New(t)

	values := []int{9, 3, 20, 8, 6, 5, 12, 10, 9, 18}
	heap := FromSlice(values, func(a, b int) bool { return a > b })

	assert.Equal([]int{20, 18, 12, 10, 6, 5, 9, 8, 9, 3}, heap.GetValues())

	ok, err := heap.Delete(12)
	assert.True(ok)
	assert.NoError(err)
	assert.Len(heap.GetValues(), 9)
	assert.Equal([]int{20, 18, 3, 10, 6, 5, 9, 8, 9}, heap.GetValues())

	heap.Pop()
	assert.Len(heap.GetValues(), 8)

	heap.Clear()
	assert.Len(heap.GetValues(), 0)

	input := []int{20, 18, 10, 9, 9, 8, 6, 5, 3}
	heap.Push(input...)

	for idx := range heap.GetValues() {
		val := heap.Pop()
		assert.Equal(val, input[idx])
	}
}

func Example_maxHeap() {
	values := []int{9, 3, 20, 8, 6, 5, 12, 10, 9, 18}
	heap := FromSlice(values, func(a, b int) bool { return a > b })
	fmt.Println(heap.GetValues())

	ok, _ := heap.Delete(12)
	fmt.Println(ok)
	fmt.Println(heap.Size())

	heap.Clear()
	fmt.Println(heap.GetValues())

	input := []int{20, 18, 10, 9, 9, 8, 6, 5, 3}
	heap.Push(input...)

	popSlice := []int{}
	for range heap.GetValues() {
		val := heap.Pop()
		popSlice = append(popSlice, val)
	}
	fmt.Println(popSlice)

	// Output:
	// [20 18 12 10 6 5 9 8 9 3]
	// true
	// 9
	// []
	// [20 18 10 9 9 8 6 5 3]
}

func TestHeap_Struct(t *testing.T) {
	assert := assert.New(t)

	type person struct {
		name string
		age  int
	}

	persons := []person{
		{name: "John", age: 23},
		{name: "Eveline", age: 32},
		{name: "Rick", age: 34},
		{name: "Tommy", age: 43},
		{name: "Jack", age: 21},
		{name: "Kim", age: 18},
	}

	comp := func(a, b person) bool { return a.age < b.age }
	heap := NewHeap(comp)
	for _, p := range persons {
		heap.Push(p)
	}
	assert.Len(heap.GetValues(), 6)
	first := heap.Pop()
	assert.Equal("Kim", first.name)

	heap.Push(person{name: "Liza", age: 50})
	assert.Equal("Liza", heap.GetValues()[heap.Size()-1].name)

	pers := person{name: "John", age: 23}
	ok, err := heap.Delete(pers)
	assert.NoError(err)
	assert.True(ok)

	pers = person{name: "John", age: 23}
	ok, err = heap.Delete(pers)
	assert.Error(err)
	assert.False(ok)
}

func TestHeap_Multiple(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		name   string
		cond   func(a, b int) bool
		actual []int
		sorted []int
	}{
		{
			name:   "MinHeap",
			cond:   func(a, b int) bool { return a < b },
			actual: []int{5, 3, 2, 4, 1},
			sorted: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "MaxHeap",
			cond:   func(a, b int) bool { return a > b },
			actual: []int{5, 3, 2, 4, 1},
			sorted: []int{5, 4, 3, 2, 1},
		},
	}

	for _, t := range testCases {
		heap := NewHeap(t.cond)
		heap.Push(t.actual...)

		for _, v := range t.sorted {
			elem := heap.Pop()
			assert.Equal(elem, v)
		}
	}
}

func TestHeap_Convert(t *testing.T) {
	assert := assert.New(t)

	input := []int{1, 4, 2, 3, 5}

	heap := NewHeap(func(a, b int) bool { return a < b })
	heap.Push(input...)
	heap.Convert(func(a, b int) bool { return a > b })
	assert.Equal([]int{5, 4, 2, 1, 3}, heap.GetValues())
}

func Example_convert() {
	input := []int{1, 4, 2, 3, 5}

	heap := NewHeap(func(a, b int) bool { return a < b })
	heap.Push(input...)
	heap.Convert(func(a, b int) bool { return a > b })
	fmt.Println(heap.GetValues())

	// Output:
	// [5 4 2 1 3]
}

func TestHeap_Merge(t *testing.T) {
	assert := assert.New(t)

	slice1 := []int{1, 4, 2, 3, 5}
	slice2 := []int{8, 6, 9, 10, 7}

	heap1 := FromSlice(slice1, func(a, b int) bool { return a < b })
	assert.Len(heap1.GetValues(), 5)
	heap2 := FromSlice(slice2, func(a, b int) bool { return a < b })
	assert.Len(heap2.GetValues(), 5)

	mergedHeap := heap1.Merge(heap2)
	assert.Len(mergedHeap.GetValues(), 10)
	assert.Len(heap1.GetValues(), 5)
	assert.Len(heap2.GetValues(), 5)
}

func Example_merge() {
	slice1 := []int{1, 4, 2, 3, 5}
	slice2 := []int{8, 6, 9, 10, 7}

	heap1 := FromSlice(slice1, func(a, b int) bool { return a < b })
	heap2 := FromSlice(slice2, func(a, b int) bool { return a < b })

	mergedHeap := heap1.Merge(heap2)
	fmt.Println(mergedHeap.Size())
	fmt.Println(heap1.Size())
	fmt.Println(heap2.Size())

	// Output:
	// 10
	// 5
	// 5
}

func TestHeap_Meld(t *testing.T) {
	assert := assert.New(t)

	slice1 := []int{1, 4, 2, 3, 5}
	slice2 := []int{8, 6, 9, 10, 7}

	heap1 := FromSlice(slice1, func(a, b int) bool { return a < b })
	assert.Len(heap1.GetValues(), 5)
	heap2 := FromSlice(slice2, func(a, b int) bool { return a < b })
	assert.Len(heap2.GetValues(), 5)

	mergedHeap := heap1.Meld(heap2)
	assert.Len(mergedHeap.GetValues(), 10)
	assert.Len(heap1.GetValues(), 0)
	assert.Len(heap2.GetValues(), 0)
}

func Example_meld() {
	slice1 := []int{1, 4, 2, 3, 5}
	slice2 := []int{8, 6, 9, 10, 7}

	heap1 := FromSlice(slice1, func(a, b int) bool { return a < b })
	heap2 := FromSlice(slice2, func(a, b int) bool { return a < b })

	mergedHeap := heap1.Meld(heap2)
	fmt.Println(mergedHeap.Size())
	fmt.Println(heap1.Size())
	fmt.Println(heap2.Size())

	// Output:
	// 10
	// 0
	// 0
}

func TestHeap_Concurrency(t *testing.T) {
	assert := assert.New(t)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	heap := NewHeap(func(a, b int) bool { return a < b })

	slice := []int{1, 4, 2, 3, 5}
	for i := 0; i < len(slice); i++ {
		wg.Add(1)
		go func(i int) {
			mu.Lock()
			heap.Push(slice[i])
			mu.Unlock()

			wg.Done()
		}(i)
	}
	wg.Wait()

	assert.Equal(1, heap.Peek())
	assert.Equal(5, heap.Size())

	heap.Pop()
	assert.Equal(2, heap.Peek())
	assert.Equal(4, heap.Size())

	heap.Clear()
	assert.Empty(heap.Size())
}
