package stack

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	assert := assert.New(t)

	s := New[int]()
	s.Push(1)
	assert.Equal(1, s.Peek())
	s.Push(2)
	assert.Equal(2, s.Peek())
	s.Pop()
	assert.Equal(1, s.Peek())
	assert.True(s.Search(1))
	s.Pop()
	assert.False(s.Search(1))
	assert.Empty(s.Size())

	s.Push(1)
	s.Pop()
	assert.False(s.Search(1))
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
