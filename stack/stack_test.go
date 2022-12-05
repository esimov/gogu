package stack

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	assert := assert.New(t)

	stack := New[int]()
	stack.Push(1)
	assert.Equal(1, stack.Peek())
	stack.Push(2)
	assert.Equal(2, stack.Peek())
	stack.Pop()
	assert.Equal(1, stack.Peek())
	assert.True(stack.Search(1))
	stack.Pop()
	assert.False(stack.Search(1))
	assert.Empty(stack.Size())

	stack.Push(1)
	stack.Pop()
	assert.False(stack.Search(1))
}

func TestStack_Concurrency(t *testing.T) {
	assert := assert.New(t)
	wg := &sync.WaitGroup{}

	s := New[int]()
	n := 100

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			s.Push(i)
			wg.Done()
		}(i)
		wg.Wait()
	}

	assert.Equal(n, s.Size())
	for i := s.Size() - 1; i > 0; i-- {
		item := s.Pop()
		assert.Equal(i, item)
	}
}

func Example() {
	stack := New[string]()

	stack.Push("foo")
	fmt.Println(stack.Size())
	fmt.Println(stack.Peek())
	stack.Push("bar")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Search("foo"))
	fmt.Println(stack.Peek())

	// Output:
	// 1
	// foo
	// bar
	// true
	// foo
}

func TestStack_Race(t *testing.T) {
	const count = 1000
	stack := New[int]()
	for i := 0; i < 64; i++ {
		go func() {
			for i := 0; i < count; i++ {
				stack.Push(i)
				stack.Peek()
				stack.Size()
				stack.Search(10)
			}
		}()
	}
	for i := 0; i < count; i++ {
		stack.Push(0)
		stack.Pop()
	}
}
