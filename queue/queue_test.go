package queue

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	assert := assert.New(t)

	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(3, q.Size())
	assert.Equal(1, q.Peek())
	q.Dequeue()
	assert.Equal(2, q.Peek())
	q.Dequeue()
	assert.Equal(3, q.Peek())
	assert.True(q.Search(3))
	q.Dequeue()
	assert.Equal(0, q.Size())
	assert.False(q.Search(3))
}

func TestQueue_Concurrency(t *testing.T) {
	assert := assert.New(t)
	wg := &sync.WaitGroup{}

	q := New[int]()
	ch := make(chan int)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			q.Enqueue(i)
			ch <- i
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for range ch {
		_, err := q.Dequeue()
		assert.NoError(err)
	}
	assert.Equal(0, q.Size())
}
