package queue

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedQueue(t *testing.T) {
	assert := assert.New(t)

	ls := NewLinked(1)
	ls.Enqueue(2)
	ls.Enqueue(3)
	assert.Equal(1, ls.Peek())
	ls.Dequeue()
	assert.Equal(2, ls.Peek())
	ls.Dequeue()
	assert.Equal(3, ls.Peek())
	assert.True(ls.Search(3))
	_, err := ls.Dequeue()
	assert.Error(err)
}

func TestLinkedQueue_Concurrency(t *testing.T) {
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
