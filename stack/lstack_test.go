package stack

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedStack(t *testing.T) {
	assert := assert.New(t)

	l := NewLinked(1)
	assert.Equal(1, l.Size())
	assert.Equal(1, l.Peek())

	l.Push(2)
	assert.Equal(2, l.Size())
	assert.Equal(2, l.Peek())

	l.Pop()
	assert.Equal(1, l.Size())
	assert.Equal(1, l.Peek())
	assert.True(l.Search(1))

	l.Pop()
	l.Push(1)
	l.Push(2)
	l.Pop()
	assert.True(l.Search(1))

	l.Pop()
	l.Pop()
	assert.Equal(0, l.Size())
}

func TestLinkedStack_Concurrency(t *testing.T) {
	assert := assert.New(t)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	l := NewLinked(0)

	n := 100
	tmp := make([]int, n)
	tmp[0] = 0

	for i := 1; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			l.Push(i)

			mu.Lock()
			tmp[i] = i
			mu.Unlock()

			wg.Done()
		}(i)
	}
	wg.Wait()
	assert.Equal(n, l.Size())

	for l.Size() > 0 {
		item := l.Pop()
		assert.Equal(tmp[item], item)
	}
}

func Example_LinkedList() {
	l := NewLinked("foo")
	fmt.Println(l.Size())
	fmt.Println(l.Peek())

	l.Push("bar")
	fmt.Println(l.Peek())

	fmt.Println(l.Pop())
	fmt.Println(l.Peek())
	fmt.Println(l.Search("foo"))

	// Output:
	// 1
	// foo
	// bar
	// foo
	// foo
	// true
}
