package stack

import (
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

	_, err := l.Pop()
	assert.Error(err)

	l.Push(1)
	l.Push(2)
	_, err = l.Pop()
	assert.NoError(err)
	assert.True(l.Search(1))

	l.Pop()
	l.Pop()
	// the stack cannot be cleared out totally, it should have at least one element
	assert.Equal(1, l.Size())
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

	for {
		item, err := l.Pop()
		assert.Equal(tmp[item], item)
		if err != nil {
			break
		}
	}
}
