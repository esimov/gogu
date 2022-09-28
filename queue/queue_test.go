package queue

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	assert := assert.New(t)

	s := New[int]()
	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)
	assert.Equal(3, s.Size())
	assert.Equal(1, s.Peek())
	s.Dequeue()
	assert.Equal(2, s.Peek())
	s.Dequeue()
	assert.Equal(3, s.Peek())
	assert.True(s.Search(3))
	s.Dequeue()
	assert.Equal(0, s.Size())
	assert.False(s.Search(3))
}

func TestQueue_Concurrency(t *testing.T) {
	assert := assert.New(t)
	wg := &sync.WaitGroup{}

	s := New[int]()
	ch := make(chan int)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			s.Enqueue(i)
			ch <- i
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		assert.Equal(i, s.Dequeue())
	}
}
